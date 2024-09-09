// Copyright 2024 The tk9.0-go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !windows

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
	_ "modernc.org/rec/lib"       // generator.go
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

func TestTokenizer(t *testing.T) {
	for i, test := range []struct {
		s    string
		ids  []int
		toks []string
	}{
		{},
		{"a", []int{0}, []string{"a"}},
		{"\\$", []int{0}, []string{"\\$"}},
		{"\\$\\$", []int{0}, []string{"\\$\\$"}},
		{"\\$\\$\\$", []int{0}, []string{"\\$\\$\\$"}},

		{"\\$\\$\\$\\$", []int{0}, []string{"\\$\\$\\$\\$"}},
		{"a\\$", []int{0}, []string{"a\\$"}},
		{"a\\$\\$", []int{0}, []string{"a\\$\\$"}},
		{"a\\$\\$\\$", []int{0}, []string{"a\\$\\$\\$"}},
		{"a\\$\\$\\$\\$", []int{0}, []string{"a\\$\\$\\$\\$"}},

		{"$a$", []int{1}, []string{"$a$"}},
		{"$$a$", []int{2}, []string{"$$a$"}},
		{"$$a$$", []int{2}, []string{"$$a$$"}},
		{"$a$$", []int{2}, []string{"$a$$"}},
		{"x$a$", []int{0, 1}, []string{"x", "$a$"}},

		{"x$$a$", []int{0, 2}, []string{"x", "$$a$"}},
		{"x$$a$$", []int{0, 2}, []string{"x", "$$a$$"}},
		{"x$a$$", []int{0, 2}, []string{"x", "$a$$"}},
		{"x$a$y", []int{0, 1, 0}, []string{"x", "$a$", "y"}},
		{"x$$a$y", []int{0, 2, 0}, []string{"x", "$$a$", "y"}},

		{"x$$a$$y", []int{0, 2, 0}, []string{"x", "$$a$$", "y"}},
		{"x$a$$y", []int{0, 2, 0}, []string{"x", "$a$$", "y"}},
		{"x\\$0$a\\$1b$$\\$y", []int{0, 2, 0}, []string{"x\\$0", "$a\\$1b$$", "\\$y"}},
	} {
		var ids []int
		var toks []string
		//trc(" test.s=%q", test.s)
		for s := test.s; ; {
			id, len := mlToken(s)
			//trc("\ts=%q id=%v len=%v", s[len:], id, len)
			if len == 0 {
				break
			}

			ids = append(ids, id)
			toks = append(toks, s[:len])
			if len == 0 {
				break
			}
			s = s[len:]
		}
		if g, e := fmt.Sprintf("%v %q", ids, toks), fmt.Sprintf("%v %q", test.ids, test.toks); g != e {
			t.Errorf("#%3v: `%s`\ngot %s\nexp %s", i, test.s, g, e)
		}
	}
}
