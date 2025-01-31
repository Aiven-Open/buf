// Copyright 2020-2024 Buf Technologies, Inc.
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

// Package buflintv2 contains the VersionSpec for v2.
//
// It uses buflintcheck and buflintbuild.
//
// The only changes from v1 to v2 are:
// 1. New rule PACKAGE_NO_IMPORT_CYCLE was added to MINIMAL, BASIC, DEFAULT.
// 2. New rule FIELD_NO_REQUIRED was added to BASIC and DEFAULT.
// 3. New rule STABLE_PACKAGE_NO_IMPORT_UNSTABLE added with no category.
package buflintv2

import (
	"github.com/bufbuild/buf/private/bufpkg/bufcheck/internal"
	"github.com/bufbuild/buf/private/bufpkg/bufconfig"
)

// VersionSpec is the version specification for v2.
//
// PACKAGE_NO_IMPORT_CYCLE was added to MINIMAL, BASIC, DEFAULT.
var VersionSpec = &internal.VersionSpec{
	FileVersion:       bufconfig.FileVersionV2,
	RuleBuilders:      v2RuleBuilders,
	DefaultCategories: v2DefaultCategories,
	IDToCategories:    v2IDToCategories,
}
