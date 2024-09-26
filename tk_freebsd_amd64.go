// Copyright 2024 The tk9.0-go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tk9_0 // import "modernc.org/tk9.0"

import _ "embed"

const (
	tclBin = "libtcl9.0.so"
	tkBin  = "libtcl9tk9.0.so"
)

//go:embed embed/freebsd/amd64/lib.zip
var libZip []byte

var shasig = map[string]string{
	"libtcl9.0.so":    "08a8f6c91baf89e6971e3ef2a078bed570d4eee0ddd0f675cf51421eb505a2f0",
	"libtcl9tk9.0.so": "836dbe2f223bd05ab35443ccf8e7631f64c101a51ef9e171f15ff5f89e66c2a8",
	"libtk9.0.0.zip":  "b2ad542fad02af3553add200dc63853044db814191b4d2178a8d1f758fafa658",
}
