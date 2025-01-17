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

package ptraceotlp

import (
	otlpcollectortrace "go.opentelemetry.io/collector/pdata/internal/data/protogen/collector/trace/v1"
)

// ExportPartialSuccess represents the details of a partially successful export request.
//
// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewExportPartialSuccess function to create new instances.
// Important: zero-initialized instance is not valid for use.
type ExportPartialSuccess struct {
	orig *otlpcollectortrace.ExportTracePartialSuccess
}

func newExportPartialSuccess(orig *otlpcollectortrace.ExportTracePartialSuccess) ExportPartialSuccess {
	return ExportPartialSuccess{orig}
}

// NewExportPartialSuccess creates a new empty ExportPartialSuccess.
//
// This must be used only in testing code. Users should use "AppendEmpty" when part of a Slice,
// OR directly access the member if this is embedded in another struct.
func NewExportPartialSuccess() ExportPartialSuccess {
	return newExportPartialSuccess(&otlpcollectortrace.ExportTracePartialSuccess{})
}

// MoveTo moves all properties from the current struct overriding the destination and
// resetting the current instance to its zero value
func (ms ExportPartialSuccess) MoveTo(dest ExportPartialSuccess) {
	*dest.orig = *ms.orig
	*ms.orig = otlpcollectortrace.ExportTracePartialSuccess{}
}

// RejectedSpans returns the rejectedspans associated with this ExportPartialSuccess.
func (ms ExportPartialSuccess) RejectedSpans() int64 {
	return ms.orig.RejectedSpans
}

// SetRejectedSpans replaces the rejectedspans associated with this ExportPartialSuccess.
func (ms ExportPartialSuccess) SetRejectedSpans(v int64) {
	ms.orig.RejectedSpans = v
}

// ErrorMessage returns the errormessage associated with this ExportPartialSuccess.
func (ms ExportPartialSuccess) ErrorMessage() string {
	return ms.orig.ErrorMessage
}

// SetErrorMessage replaces the errormessage associated with this ExportPartialSuccess.
func (ms ExportPartialSuccess) SetErrorMessage(v string) {
	ms.orig.ErrorMessage = v
}

// CopyTo copies all properties from the current struct overriding the destination.
func (ms ExportPartialSuccess) CopyTo(dest ExportPartialSuccess) {
	dest.SetRejectedSpans(ms.RejectedSpans())
	dest.SetErrorMessage(ms.ErrorMessage())
}
