// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package plog

import (
	"testing"

	gogoproto "github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/assert"
	goproto "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	otlpcollectorlog "go.opentelemetry.io/collector/pdata/internal/data/protogen/collector/logs/v1"
	otlplogs "go.opentelemetry.io/collector/pdata/internal/data/protogen/logs/v1"
)

func TestLogRecordCount(t *testing.T) {
	logs := NewLogs()
	assert.EqualValues(t, 0, logs.LogRecordCount())

	rl := logs.ResourceLogs().AppendEmpty()
	assert.EqualValues(t, 0, logs.LogRecordCount())

	ill := rl.ScopeLogs().AppendEmpty()
	assert.EqualValues(t, 0, logs.LogRecordCount())

	ill.LogRecords().AppendEmpty()
	assert.EqualValues(t, 1, logs.LogRecordCount())

	rms := logs.ResourceLogs()
	rms.EnsureCapacity(3)
	rms.AppendEmpty().ScopeLogs().AppendEmpty()
	illl := rms.AppendEmpty().ScopeLogs().AppendEmpty().LogRecords()
	for i := 0; i < 5; i++ {
		illl.AppendEmpty()
	}
	// 5 + 1 (from rms.At(0) initialized first)
	assert.EqualValues(t, 6, logs.LogRecordCount())
}

func TestLogRecordCountWithEmpty(t *testing.T) {
	assert.Zero(t, NewLogs().LogRecordCount())
	assert.Zero(t, newLogs(&otlpcollectorlog.ExportLogsServiceRequest{
		ResourceLogs: []*otlplogs.ResourceLogs{{}},
	}).LogRecordCount())
	assert.Zero(t, newLogs(&otlpcollectorlog.ExportLogsServiceRequest{
		ResourceLogs: []*otlplogs.ResourceLogs{
			{
				ScopeLogs: []*otlplogs.ScopeLogs{{}},
			},
		},
	}).LogRecordCount())
	assert.Equal(t, 1, newLogs(&otlpcollectorlog.ExportLogsServiceRequest{
		ResourceLogs: []*otlplogs.ResourceLogs{
			{
				ScopeLogs: []*otlplogs.ScopeLogs{
					{
						LogRecords: []*otlplogs.LogRecord{{}},
					},
				},
			},
		},
	}).LogRecordCount())
}

func TestToFromLogOtlp(t *testing.T) {
	otlp := &otlpcollectorlog.ExportLogsServiceRequest{}
	logs := newLogs(otlp)
	assert.EqualValues(t, NewLogs(), logs)
	assert.EqualValues(t, otlp, logs.getOrig())
}

func TestResourceLogsWireCompatibility(t *testing.T) {
	// This test verifies that OTLP ProtoBufs generated using goproto lib in
	// opentelemetry-proto repository OTLP ProtoBufs generated using gogoproto lib in
	// this repository are wire compatible.

	// Generate ResourceLogs as pdata struct.
	logs := NewLogs()
	fillTestResourceLogsSlice(logs.ResourceLogs())

	// Marshal its underlying ProtoBuf to wire.
	wire1, err := gogoproto.Marshal(logs.getOrig())
	assert.NoError(t, err)
	assert.NotNil(t, wire1)

	// Unmarshal from the wire to OTLP Protobuf in goproto's representation.
	var goprotoMessage emptypb.Empty
	err = goproto.Unmarshal(wire1, &goprotoMessage)
	assert.NoError(t, err)

	// Marshal to the wire again.
	wire2, err := goproto.Marshal(&goprotoMessage)
	assert.NoError(t, err)
	assert.NotNil(t, wire2)

	// Unmarshal from the wire into gogoproto's representation.
	var gogoprotoRS2 otlpcollectorlog.ExportLogsServiceRequest
	err = gogoproto.Unmarshal(wire2, &gogoprotoRS2)
	assert.NoError(t, err)

	// Now compare that the original and final ProtoBuf messages are the same.
	// This proves that goproto and gogoproto marshaling/unmarshaling are wire compatible.
	assert.EqualValues(t, logs.getOrig(), &gogoprotoRS2)
}

func TestLogsMoveTo(t *testing.T) {
	logs := NewLogs()
	fillTestResourceLogsSlice(logs.ResourceLogs())
	dest := NewLogs()
	logs.MoveTo(dest)
	assert.EqualValues(t, NewLogs(), logs)
	assert.EqualValues(t, generateTestResourceLogsSlice(), dest.ResourceLogs())
}

func TestLogsCopyTo(t *testing.T) {
	logs := NewLogs()
	fillTestResourceLogsSlice(logs.ResourceLogs())
	logsCopy := NewLogs()
	logs.CopyTo(logsCopy)
	assert.EqualValues(t, logs, logsCopy)
}

func TestSeverityNumberString(t *testing.T) {
	assert.EqualValues(t, "Unspecified", SeverityNumberUnspecified.String())
	assert.EqualValues(t, "Trace", SeverityNumberTrace.String())
	assert.EqualValues(t, "Trace2", SeverityNumberTrace2.String())
	assert.EqualValues(t, "Trace3", SeverityNumberTrace3.String())
	assert.EqualValues(t, "Trace4", SeverityNumberTrace4.String())
	assert.EqualValues(t, "Debug", SeverityNumberDebug.String())
	assert.EqualValues(t, "Debug2", SeverityNumberDebug2.String())
	assert.EqualValues(t, "Debug3", SeverityNumberDebug3.String())
	assert.EqualValues(t, "Debug4", SeverityNumberDebug4.String())
	assert.EqualValues(t, "Info", SeverityNumberInfo.String())
	assert.EqualValues(t, "Info2", SeverityNumberInfo2.String())
	assert.EqualValues(t, "Info3", SeverityNumberInfo3.String())
	assert.EqualValues(t, "Info4", SeverityNumberInfo4.String())
	assert.EqualValues(t, "Warn", SeverityNumberWarn.String())
	assert.EqualValues(t, "Warn2", SeverityNumberWarn2.String())
	assert.EqualValues(t, "Warn3", SeverityNumberWarn3.String())
	assert.EqualValues(t, "Warn4", SeverityNumberWarn4.String())
	assert.EqualValues(t, "Error", SeverityNumberError.String())
	assert.EqualValues(t, "Error2", SeverityNumberError2.String())
	assert.EqualValues(t, "Error3", SeverityNumberError3.String())
	assert.EqualValues(t, "Error4", SeverityNumberError4.String())
	assert.EqualValues(t, "Fatal", SeverityNumberFatal.String())
	assert.EqualValues(t, "Fatal2", SeverityNumberFatal2.String())
	assert.EqualValues(t, "Fatal3", SeverityNumberFatal3.String())
	assert.EqualValues(t, "Fatal4", SeverityNumberFatal4.String())
	assert.EqualValues(t, "", SeverityNumber(100).String())
}
