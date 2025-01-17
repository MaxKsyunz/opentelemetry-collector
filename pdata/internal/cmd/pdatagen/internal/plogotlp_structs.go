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

package internal // import "go.opentelemetry.io/collector/pdata/internal/cmd/pdatagen/internal"
import (
	"path/filepath"
)

var logOtlpFile = &File{
	Path:        filepath.Join("pdata", "plog"),
	Name:        "logs_otlp",
	PackageName: "plogotlp",
	imports: []string{
		`"sort"`,
		``,
		`"go.opentelemetry.io/collector/pdata/internal"`,
		`otlpcollectorlog "go.opentelemetry.io/collector/pdata/internal/data/protogen/collector/logs/v1"`,
	},
	testImports: []string{
		`"testing"`,
		``,
		`"github.com/stretchr/testify/assert"`,
		``,
		`"go.opentelemetry.io/collector/pdata/internal"`,
		`otlpcollectorlog "go.opentelemetry.io/collector/pdata/internal/data/protogen/collector/logs/v1"`,
	},
	structs: []baseStruct{
		exportLogsPartialSuccess,
	},
}

var exportLogsPartialSuccess = &messageValueStruct{
	structName:     "ExportPartialSuccess",
	description:    "// ExportPartialSuccess represents the details of a partially successful export request.",
	originFullName: "otlpcollectorlog.ExportLogsPartialSuccess",
	fields: []baseField{
		&primitiveField{
			fieldName:  "RejectedLogRecords",
			returnType: "int64",
			defaultVal: `int64(0)`,
			testVal:    `int64(13)`,
		},
		&primitiveField{
			fieldName:  "ErrorMessage",
			returnType: "string",
			defaultVal: `""`,
			testVal:    `"error message"`,
		},
	},
}
