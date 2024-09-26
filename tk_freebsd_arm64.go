// Copyright 2024 The tk9.0-go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tk9_0 // import "modernc.org/tk9.0"

import _ "embed"

const (
	tclBin = "libtcl9.0.so"
	tkBin  = "libtcl9tk9.0.so"
)

//go:embed embed/freebsd/arm64/lib.zip
var libZip []byte

var shasig = map[string]string{
	"libtcl9.0.so":    "808a1210f3e2dd98880e82d7f0eaa32dc4370e1018c237a3688941421d8214be",
	"libtcl9tk9.0.so": "aa9d28c5da6243bf36ca2987804a4cc79830809f3c02b6b91e79e7a086e09441",
	"libtk9.0.0.zip":  "d4119c0082b6bd039a469e280d9c1ab120e168949f705d5b37aa19fc029376f9",
}
