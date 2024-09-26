// Copyright 2024 The tk9.0-go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tk9_0 // import "modernc.org/tk9.0"

import _ "embed"

const (
	tclBin = "tcl90.dll"
	tkBin  = "tcl9tk90.dll"
)

//go:embed embed/windows/amd64/lib.zip
var libZip []byte

var shasig = map[string]string{
	"libtommath.dll":     "907b9c7860fc07231f1e238551715e5d813283807f52dc383dae0cb47a879d29",
	"tcl90.dll":          "3fc7a55e1abaddc50665ca7c4a9672a87a3235ad4cde6e294fb5609c07d86c6d",
	"tcl9dde14.dll":      "3b2ddbb0d4f6cae6f16de8d1340829db36f8ec8efe5d54e4f38d11dbb6c6498f",
	"tcl9registry13.dll": "a5efe5aefbca2e68af0dd8ed111f56348ccc2365dce61d650c3e2575ba8f260f",
	"tcl9tk90.dll":       "2822ff97299b19536922d692d631309150fedcf17c6b7b2e08c51381844b666e",
	"tcldde14.dll":       "8f85015cd4d094782f0531dc4cae117aa460b2b477c73ce065e8b6076b729ea6",
	"tclregistry13.dll":  "88b8ade32607978af416d04ce03cc00713f93c4df5b3301ed053e0e670b8405d",
	"zlib1.dll":          "6f10a76dcc2c831d1f08d98c0b345afa0911bec0238fcba357b612ccc6ab5d81",
	"libtk9.0.0.zip":     "afa5958c9c1653ff041bff6222b8283d4e3b37147ca428050890f10ddba490c2",
}
