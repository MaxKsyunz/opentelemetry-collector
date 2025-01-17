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

package internal

import (
	otlpcommon "go.opentelemetry.io/collector/pdata/internal/data/protogen/common/v1"
)

type InstrumentationScope struct {
	orig *otlpcommon.InstrumentationScope
}

func GetOrigInstrumentationScope(ms InstrumentationScope) *otlpcommon.InstrumentationScope {
	return ms.orig
}

func NewInstrumentationScope(orig *otlpcommon.InstrumentationScope) InstrumentationScope {
	return InstrumentationScope{orig: orig}
}

func GenerateTestInstrumentationScope() InstrumentationScope {
	orig := otlpcommon.InstrumentationScope{}
	tv := NewInstrumentationScope(&orig)
	FillTestInstrumentationScope(tv)
	return tv
}

func FillTestInstrumentationScope(tv InstrumentationScope) {
	tv.orig.Name = "test_name"
	tv.orig.Version = "test_version"
	FillTestMap(NewMap(&tv.orig.Attributes))
	tv.orig.DroppedAttributesCount = uint32(17)
}
