// Copyright 2024 The tk9.0-go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build tk.eval

package tk9_0

// Eval exposes the underlying Tcl interpreter evaluation directly.
func Eval(tclCode string) (result string, err error) {
	return eval(tclCode)
}
