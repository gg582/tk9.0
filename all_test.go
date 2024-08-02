// Copyright 2024 The tk9.0-go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tk9_0 // import "modernc.org/tk9.0"

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
	"testing"

	_ "github.com/adrg/xdg"       // generator.go
	_ "github.com/expr-lang/expr" // examples
	_ "golang.org/x/net/html"     // generator.go
	_ "modernc.org/ngrab/lib"     // generator.go
)

func TestMain(m *testing.M) {
	if isBuilder {
		os.Exit(0)
	}

	if Error != nil {
		fmt.Fprintln(os.Stderr, Error)
		os.Exit(1)
	}

	flag.Parse()
	rc := m.Run()
	Finalize()
	os.Exit(rc)
}

//go:embed _examples/gopher.png
var gopher []byte

func Test(t *testing.T) {
	img := NewPhoto(Data(gopher))
	if img == nil || Error != nil {
		t.Fatal(Error)
	}

	t.Log(img.String())
}
