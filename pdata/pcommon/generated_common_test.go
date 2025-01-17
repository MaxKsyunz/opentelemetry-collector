// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by "model/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "go run model/internal/cmd/pdatagen/main.go".

package pcommon

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.opentelemetry.io/collector/pdata/internal"
)

func TestInstrumentationScope_MoveTo(t *testing.T) {
	ms := InstrumentationScope(internal.GenerateTestInstrumentationScope())
	dest := NewInstrumentationScope()
	ms.MoveTo(dest)
	assert.Equal(t, NewInstrumentationScope(), ms)
	assert.Equal(t, InstrumentationScope(internal.GenerateTestInstrumentationScope()), dest)
}

func TestInstrumentationScope_CopyTo(t *testing.T) {
	ms := NewInstrumentationScope()
	orig := NewInstrumentationScope()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
	orig = InstrumentationScope(internal.GenerateTestInstrumentationScope())
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
}

func TestInstrumentationScope_Name(t *testing.T) {
	ms := NewInstrumentationScope()
	assert.Equal(t, "", ms.Name())
	ms.SetName("test_name")
	assert.Equal(t, "test_name", ms.Name())
}

func TestInstrumentationScope_Version(t *testing.T) {
	ms := NewInstrumentationScope()
	assert.Equal(t, "", ms.Version())
	ms.SetVersion("test_version")
	assert.Equal(t, "test_version", ms.Version())
}

func TestInstrumentationScope_Attributes(t *testing.T) {
	ms := NewInstrumentationScope()
	assert.Equal(t, NewMap(), ms.Attributes())
	internal.FillTestMap(internal.Map(ms.Attributes()))
	assert.Equal(t, Map(internal.GenerateTestMap()), ms.Attributes())
}

func TestInstrumentationScope_DroppedAttributesCount(t *testing.T) {
	ms := NewInstrumentationScope()
	assert.Equal(t, uint32(0), ms.DroppedAttributesCount())
	ms.SetDroppedAttributesCount(uint32(17))
	assert.Equal(t, uint32(17), ms.DroppedAttributesCount())
}
