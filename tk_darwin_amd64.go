// Copyright 2024 The tk9.0-go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tk9_0 // import "modernc.org/tk9.0"

import _ "embed"

const (
	tclBin = "libtcl9.0.dylib"
	tkBin  = "libtcl9tk9.0.dylib"
)

//go:embed embed/darwin/amd64/lib.zip
var libZip []byte

var shasig = map[string]string{
	"libtcl9.0.dylib":    "0edef85666838ab2abcced12c0ea7f64bbcd13cb51a9a42933c048774cfa49ba",
	"libtcl9tk9.0.dylib": "8a93e30f8f2819dce68b9c9cbb585910941b43381cedb4f90197e85814dcaa97",
	"libtk9.0.0.zip":     "a9d1f2d0688572f4202914f283a3d202cb290ac2bcc9a4a06f465249dce1bce1",
}
