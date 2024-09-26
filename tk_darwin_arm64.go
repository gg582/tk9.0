// Copyright 2024 The tk9.0-go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tk9_0 // import "modernc.org/tk9.0"

import _ "embed"

const (
	tclBin = "libtcl9.0.dylib"
	tkBin  = "libtcl9tk9.0.dylib"
)

//go:embed embed/darwin/arm64/lib.zip
var libZip []byte

var shasig = map[string]string{
	"libtcl9.0.dylib":    "6bb577fb7d0acd3ded57323c55a9e1509f428bdaf31ed1604a0748727df31c6b",
	"libtcl9tk9.0.dylib": "98c537b32eb14890e29d359f9f0c22aa5ccee707481352e216f51ee460c69863",
	"libtk9.0.0.zip":     "e9e6d065ea00eeb839cd405a4c3ebbc39332f61ffdc3e98fcf38326d42b8f678",
}
