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
	"libtommath.dll": "7ff97843cde97215fcf4f087d61044cda01286630b486398117967e577e039e3",
	"tcl90.dll":      "47a9fc2781c4baf75bb52ae5966617a84c8e8871b5e6f79b27c1c156f4268010",
	"tcl9tk90.dll":   "b8272954427be784cac674ed7388694b0d915f19a19af242fab0db5b9b9e464f",
	"libtk9.0.0.zip": "5f7a7241a40e1533368bba5cb1b47b532680359dfb17df11517a543445106a6c",
}
