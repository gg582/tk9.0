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
	"libtommath.dll": "2d760fefb452665b6af8c8d9d29f3a8378f10fc0847cdd9938ea0cb5edf1d573",
	"tcl90.dll":      "e9b862228e85434988f46d20ff11e1a5246797f02ff8b4420d5cf1d790be17fa",
	"tcl9tk90.dll":   "dfc74c0a60248e6cf6137253e15de320f88288c6b325a9d4bb3d139213f45d47",
	"libtk9.0.0.zip": "c40e06a91e46cbdd447dd0d3d2140a7a564ca749ceab964f26b8e21f7a3f45af",
}
