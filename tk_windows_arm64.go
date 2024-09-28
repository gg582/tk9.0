// Copyright 2024 The tk9.0-go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tk9_0 // import "modernc.org/tk9.0"

import _ "embed"

const (
	tclBin = "tcl90.dll"
	tkBin  = "tcl9tk90.dll"
)

//go:embed embed/windows/arm64/lib.zip
var libZip []byte

var shasig = map[string]string{
	// embed/windows/arm64/lib.zip
	"libtommath.dll":     "907b9c7860fc07231f1e238551715e5d813283807f52dc383dae0cb47a879d29",
	"tcl90.dll":          "015c50120ab877a6e4f8d44c9ebebd07deec5fa277dc674be48daaccd1b14b8e",
	"tcl9dde14.dll":      "e0f5fc4aa9fb86c73beaa56e66904879c3f70ccf8fa2f4eabb15efed6c0b6f62",
	"tcl9registry13.dll": "1533010edecd71fdd65eb7674b0dd3fd1d6c5aec5025a7d75c4b389c6eb436aa",
	"tcl9tk90.dll":       "9549c1082a846b5b663a49c0b9e51aeef889a4f45c537e214cd0928ec634b3e1",
	"tcldde14.dll":       "b464e702a40e5213441f87b87f8fc8c74e2cbcd4eefb6d33232d285bc9570640",
	"tclregistry13.dll":  "43b98260cef3ba8407f1ae7fe3881be832d43ae6a9e29cceb2a6209739e36e0e",
	"zlib1.dll":          "6f10a76dcc2c831d1f08d98c0b345afa0911bec0238fcba357b612ccc6ab5d81",
	"libtk9.0.0.zip":     "ba2032f799159979ab42b1c93583e7a1975095b07709886d6a88a612c302f03c",
}
