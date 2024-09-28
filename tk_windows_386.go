// Copyright 2024 The tk9.0-go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tk9_0 // import "modernc.org/tk9.0"

import _ "embed"

const (
	tclBin = "tcl90.dll"
	tkBin  = "tcl9tk90.dll"
)

//go:embed embed/windows/386/lib.zip
var libZip []byte

var shasig = map[string]string{
	// embed/windows/386/lib.zip
	"libtommath.dll": "7ff97843cde97215fcf4f087d61044cda01286630b486398117967e577e039e3",
	"tcl90.dll":      "77a76f3027cc5c9b1819ec07bc80c19ef441eab229715f2490b2f648cd22ddd4",
	"tcl9tk90.dll":   "87fc7826fab4cd532fa6b449cb267412efd30b0252fc34fad9cdea7014b3dded",
	"libtk9.0.0.zip": "834e810be3bc388f594f5b735c5fd5250711cfeb1cf84ada98323e67bb9439a7",
}
