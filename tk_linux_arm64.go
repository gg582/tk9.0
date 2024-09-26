// Copyright 2024 The tk9.0-go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tk9_0 // import "modernc.org/tk9.0"

import _ "embed"

const (
	tclBin = "libtcl9.0.so"
	tkBin  = "libtcl9tk9.0.so"
)

//go:embed embed/linux/arm64/lib.zip
var libZip []byte

var shasig = map[string]string{
	"libtcl9.0.so":    "8d892cc2f014e1763b3c7d659f02ce16a64273f6bfb4e37e86531efc48e92688",
	"libtcl9tk9.0.so": "2a5b7a7729c48504ed2233af2ec4ef08e17a3996409597bf1006c7c99d176e4e",
	"libtk9.0.0.zip":  "b88a5a45bac776e74aedc8eef05e4085c710f7047031e4b37bed58e6614fe6da",
}
