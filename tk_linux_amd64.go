// Copyright 2024 The tk9.0-go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tk9_0 // import "modernc.org/tk9.0"

import _ "embed"

const (
	tclBin = "libtcl9.0.so"
	tkBin  = "libtcl9tk9.0.so"
)

//go:embed embed/linux/amd64/lib.zip
var libZip []byte

var shasig = map[string]string{
	"libtcl9.0.so":    "c2b1ff5ba362ff021ba89278e7ada8d9ddcaa1c5074b5d14ae1005474da298bf0",
	"libtcl9tk9.0.so": "0edf386314e75fee05d5fc6d41c42b420afa78c5cf474b5de5062e04fa72787b",
	"libtk9.0.0.zip":  "55e73014a533226cdd7ab3876c1e21cb14a92c8d3f38efa9f8971bcf81926596",
}
