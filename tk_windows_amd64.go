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
	// embed/windows/amd64/lib.zip
	"libtommath.dll": "2d760fefb452665b6af8c8d9d29f3a8378f10fc0847cdd9938ea0cb5edf1d573",
	"tcl90.dll":      "e10dd6983753a5c4639822bcedb9d0791281c2404e2bf6cbd5fd7ab8aefe47a7",
	"tcl9tk90.dll":   "9018bcc9159b0d015f748c260486a1a15ff45f7ae000eede57f86dcb18c8f08d",
	"libtk9.0.0.zip": "b80665940639ac6b345131cd8f0710e8e2efec71771607e614a963515505da48",
}
