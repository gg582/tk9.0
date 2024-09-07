// Copyright ©2021 The star-tex Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE-STAR-TEX file.

// Modifications are
//
// Copyright 2024 The tk9.0-go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tk9_0 // import "modernc.org/tk9.0"

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
	"strings"

	"github.com/disintegration/imaging"
	"golang.org/x/image/font"
	fixedmath "golang.org/x/image/math/fixed"
	"modernc.org/knuth/dvi"
	"modernc.org/knuth/font/fixed"
	"modernc.org/knuth/font/pkf"
	"modernc.org/knuth/kpath"
	"modernc.org/knuth/tex"
)

const (
	shrink = 1
)

var (
	_ dvi.Renderer = (*renderer)(nil)
)

type fntkey struct {
	name string
	size fixed.Int12_20
}

type renderer struct {
	bkg   color.Color
	bound image.Rectangle
	ctx   kpath.Context
	err   error
	faces map[fntkey]font.Face
	img   *image.RGBA
	out   bytes.Buffer
	page  int
	post  dvi.CmdPost
	pre   dvi.CmdPre
	scale float64

	conv  float32 // converts DVI units to pixels
	dpi   float32
	tconv float32 // converts unmagnified DVI units to pixels

	bounded bool
}

func newRenderer(ctx kpath.Context, scale float64) *renderer {
	return &renderer{ctx: ctx, faces: make(map[fntkey]font.Face), scale: scale}
}

func (pr *renderer) Init(pre *dvi.CmdPre, post *dvi.CmdPost) {
	pr.pre = *pre
	pr.post = *post
	if pr.dpi == 0 {
		pr.dpi = 600
	}
	res := pr.dpi
	conv := float32(pr.pre.Num) / 254000.0 * (res / float32(pr.pre.Den))
	pr.tconv = conv
	pr.conv = conv * float32(pr.pre.Mag) / 1000.0
	conv = 1/(float32(pre.Num)/float32(pre.Den)*(float32(pre.Mag)/1000.0)*(pr.dpi*shrink/254000.0)) + 0.5
	if pr.bkg == nil {
		pr.bkg = color.Transparent
	}
}

func (pr *renderer) BOP(bop *dvi.CmdBOP) {
	if pr.err != nil {
		return
	}

	pr.page = int(bop.C0)
	bnd := image.Rect(0, 0, int(pr.pixels(int32(pr.post.Width))), int(pr.pixels(int32(pr.post.Height))))
	pr.img = image.NewRGBA(bnd)
	draw.Draw(pr.img, bnd, image.NewUniform(pr.bkg), image.Point{}, draw.Over)
}

func (pr *renderer) DrawGlyph(x, y int32, font dvi.Font, glyph rune, c color.Color) {
	if pr.err != nil {
		return
	}

	dot := fixedmath.Point26_6{X: fixedmath.I(int(pr.pixels(x))), Y: fixedmath.I(int(pr.pixels(y)))}
	face, ok := pr.face(font)
	if !ok {
		return
	}

	dr, mask, maskp, _, ok := face.Glyph(dot, glyph)
	if !ok {
		pr.setErr(fmt.Errorf("could not find glyph 0x%02x", glyph))
		return
	}

	draw.DrawMask(pr.img, dr, image.NewUniform(c), image.Point{}, mask, maskp, draw.Over)
	pr.union(dr)
}

func (pr *renderer) union(r image.Rectangle) {
	switch pr.bounded {
	case true:
		pr.bound = pr.bound.Union(r)
	default:
		pr.bound = r
		pr.bounded = true
	}
}

func (pr *renderer) DrawRule(x, y, w, h int32, c color.Color) {
	if pr.err != nil {
		return
	}

	r := image.Rect(int(pr.pixels(x+0)), int(pr.pixels(y+0)), int(pr.pixels(x+w)), int(pr.pixels(y-h)))
	draw.Draw(pr.img, r, image.NewUniform(c), image.Point{}, draw.Over)
	pr.union(r)
}

func (pr *renderer) EOP() {
	if pr.err != nil && pr.out.Len() != 0 {
		return
	}

	img := pr.img.SubImage(pr.bound)
	if pr.scale != 1.0 {
		img = imaging.Resize(img, int(float64(pr.bound.Min.X)*pr.scale+0.5), 0, imaging.Lanczos)
	}
	if err := png.Encode(&pr.out, img); err != nil {
		pr.setErr(fmt.Errorf("could not encode PNG image: %w", err))
		return
	}
}

func (pr *renderer) setErr(err error) {
	if pr.err == nil {
		pr.err = err
	}
}

func (pr *renderer) face(fnt dvi.Font) (font.Face, bool) {
	key := fntkey{
		name: fnt.Name(),
		size: fnt.Size(),
	}
	if f, ok := pr.faces[key]; ok {
		return f, ok
	}

	fname, err := pr.ctx.Find(key.name + ".pk")
	if err != nil {
		pr.setErr(fmt.Errorf("could not find font face %q: %+v", key.name, err))
		return nil, false
	}

	f, err := pr.ctx.Open(fname)
	if err != nil {
		pr.setErr(fmt.Errorf("could not open font face %q: %+v", key.name, err))
		return nil, false
	}

	defer f.Close()

	pk, err := pkf.Parse(f)
	if err != nil {
		pr.setErr(fmt.Errorf("could not parse font face %q: %+v", key.name, err))
		return nil, false
	}

	tfm := fnt.Metrics()
	if tfm.Checksum() != pk.Checksum() {
		pr.setErr(fmt.Errorf(
			"TFM and PK checksum do not match for %q: tfm=0x%x, pk=0x%x",
			key.name,
			tfm.Checksum(),
			pk.Checksum(),
		))
		return nil, false
	}

	face := pkf.NewFace(pk, tfm, &pkf.FaceOptions{
		Size: tfm.DesignSize().Float64(),
		DPI:  float64(pr.dpi),
	})
	pr.faces[key] = face
	return face, true
}

func (pr *renderer) pixels(v int32) int32 {
	x := pr.conv * float32(v)
	return roundF32(x / shrink)
}

func roundF32(v float32) int32 {
	if v > 0 {
		return int32(v + 0.5)
	}
	return int32(v - 0.5)
}

func dvi2png(r io.Reader, scale float64) (png []byte, err error) {
	ctx := kpath.New()
	renderer := newRenderer(ctx, scale)
	vm := dvi.NewMachine(
		dvi.WithContext(ctx),
		dvi.WithRenderer(renderer),
		dvi.WithHandlers(dvi.NewColorHandler(ctx)),
		dvi.WithOffsetX(0),
		dvi.WithOffsetY(0),
	)

	raw, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("could not read DVI program file: %w", err)
	}

	prog, err := dvi.Compile(raw)
	if err != nil {
		return nil, fmt.Errorf("could not compile DVI program: %w", err)
	}

	err = vm.Run(prog)
	if err != nil {
		return nil, fmt.Errorf("could not interpret DVI program: %w", err)
	}

	if renderer.err != nil {
		return nil, fmt.Errorf("could not render DVI program: %w", renderer.err)
	}

	return renderer.out.Bytes(), nil
}

// TeX renders TeX 'src' as a png file that shows the TeX "snippet" in a fixed
// 600 dpi resolution. The result is afterwards resized by 'scale' factor.
//
// Only plain Tex and the default Computer Modern fonts are supported. To get
// rid of the page number rendered by default, the function prepends
// "\footline={}\n" to src. Also, "\n\bye\n" is appended to 'src' to make it a
// complete TeX document.
//
// The function has no runtime dependencies.
func TeX(src string, scale float64) (png []byte) {
	//TODO do not write plain.log to current directory which may not even be writable.
	var stdout, stderr, dvi bytes.Buffer
	if err := tex.Main(
		strings.NewReader("\\input plain \\input x"),
		&stdout,
		&stderr,
		tex.WithInputFile("x.tex", strings.NewReader(fmt.Sprintf("\\footline={}\n%s\n\\bye\n", src))),
		tex.WithDVIFile(&dvi),
	); err != nil {
		fail(fmt.Errorf("FAIL err=%v\nstdout=%s\nstderr=%s", err, stdout.Bytes(), stderr.Bytes()))
		return nil
	}

	b, err := dvi2png(&dvi, scale)
	if err != nil {
		fail(err)
		return nil
	}

	return b
}

