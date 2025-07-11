// Copyright 2024 The tk9.0-go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build none

package main

import (
	"archive/zip"
	"bytes"
	"crypto/sha256"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"slices"
	"strings"

	"github.com/adrg/xdg"
	"golang.org/x/net/html"
	util "modernc.org/fileutil/ccgo"
	tcllib "modernc.org/libtcl9.0/library"
	libtk "modernc.org/libtk9.0"
	tklib "modernc.org/libtk9.0/library"
	ngrab "modernc.org/ngrab/lib"
	"modernc.org/rec/lib"
)

const (
	header = `// Code generated by generator.go, DO NOT EDIT.

package tk9_0 // import "github.com/gg582/tk9.0"

import "fmt"

`
	head = `<html>
<head>
<style>
div {
	margin-left: 50px;
}
p {
    padding : 0;
    margin : 0;
    background-color: #FDFDC9;
}
.comment {
	background-color: lightgray;
}
.TH {
	background-color: #9E9E9E;
}
.SH {
	background-color: #2196F3;
}
.SO .text {
	background-color: #00FFFF;
}
.OP {
	background-color: #FFEB3B;
}
.BS {
	background-color: #F44336;
}
.CS .text {
	background-color: #87CEEB;
}
.DS {
	background-color: #F44336;
}
.RS {
	background-color: #F44336;
}
.sso {
	background-color: #616161;
}
.SS {
	background-color: #8f8;
}
.LP {
	background-color: #8f8;
}
</style>
</head>
<body>
`
	goarch  = runtime.GOARCH
	goos    = runtime.GOOS
	ofn     = "generated.go"
	tempDir = "html"
)

type cmdOpts map[string][]string

type pageInfo struct {
	commands map[string][]string // Go command name: option names
	ignore   bool
	manual   bool
	widget   bool
}

var (
	pageInfos = map[string]*pageInfo{
		"bell": {
			manual: true, // done
			commands: cmdOpts{"Bell": []string{
				"-displayof",
				"-nice",
			}},
		},
		"bind":     {manual: true}, // done
		"bindtags": {ignore: true},
		"bitmap": {
			manual: true, // done
			commands: cmdOpts{"NewBitmap": []string{
				"-background",
				"-data",
				"-file",
				"-foreground",
				"-maskdata",
				"-maskfile",
			}},
		},
		"busy": {
			manual: true, // done
			commands: cmdOpts{".Busy": []string{
				"-cursor",
			}},
		},
		"chooseColor": {
			manual: true, // done
			commands: cmdOpts{"ChooseColor": []string{
				"-initialcolor",
				"-parent",
				"-title",
			}},
		},
		"chooseDirectory": {
			manual: true, // done
			commands: cmdOpts{"ChooseDirectory": []string{
				//TODO? "-command",
				"-initialdir",
				"-message",
				"-mustexist",
				"-parent",
				"-title",
			}},
		},
		"clipboard": {
			manual: true, // done
			commands: cmdOpts{
				"ClipboardAppend": []string{
					"-displayof",
					"-format",
					"-type",
				},
				"ClipboardClear": []string{
					"-displayof",
				},
				"ClipboardGet": []string{
					"-displayof",
					"-type",
				},
			},
		},
		"colors":  {manual: true}, // done
		"console": {ignore: true},
		"cursors": {manual: true}, // done
		"destroy": {manual: true}, // done
		"dialog":  {ignore: true}, // ... largely deprecated by the tk_messageBox
		"event": {
			ignore: true, //MAYBE later
			commands: cmdOpts{"EventGenerate": []string{
				"-above",
				"-borderwidth",
				"-button",
				"-count",
				"-data",
				"-delta",
				"-detail",
				"-focus",
				"-height",
				"-keycode",
				"-keysym",
				"-mode",
				"-override",
				"-place",
				"-root",
				"-rootx",
				"-rooty",
				"-sendevent",
				"-serial",
				"-state",
				"-subwindow",
				"-time",
				"-warp",
				"-width",
				"-when",
				"-x",
				"-y",
			}},
		},
		"focus": {
			manual: true, // done
			commands: cmdOpts{"Focus": []string{
				"-displayof",
				"-force",
				"-lastfor",
			}},
		},
		"focusNext": {ignore: true}, //MAYBE later
		"font": {
			manual: true, // done
			commands: cmdOpts{
				"NewFont": []string{
					"-family",
					"-size",
					"-weight",
					"-slant",
					"-underline",
					"-overstrike",
				},
				"FontFamilies": []string{
					"-displayof",
				},
				// "FontMetrics": []string{
				// 	"-ascent",
				// 	"-descent",
				// 	"-linespace",
				// 	"-fixed",
				// },
			},
		},
		"fontchooser": {
			manual: true, // done
			commands: cmdOpts{"Fontchooser": []string{
				"-parent",
				"-title",
				"-font",
				"-command",
				// "-visible",
			}},
		},
		"getOpenFile": {
			manual: true, // done
			commands: cmdOpts{
				"GetOpenFile": []string{
					// "-command", // macOS only
					"-defaultextension",
					"-filetypes",
					"-initialdir",
					"-initialfile",
					// "-message", // macOS only
					"-multiple",
					"-parent",
					"-title",
					// "-typevariable",
				},
				"GetSaveFile": []string{
					// "-command",
					"-confirmoverwrite",
					"-defaultextension",
					"-filetypes",
					"-initialdir",
					"-initialfile",
					// "-message",
					"-parent",
					"-title",
					// "-typevariable",
				},
			},
		},
		"grab": {
			ignore: true, //MAYBE later
			// commands: cmdOpts{
			// 	"Grab": []string{
			// 		"-global",
			// 	},
			// 	"GrabSet": []string{
			// 		"-global",
			// 	},
			// },
		},
		"grid": {
			manual: true, // done
			commands: cmdOpts{"Grid": []string{
				"-column",
				"-columnspan",
				"-in",
				"-ipadx",
				"-ipady",
				"-padx",
				"-pady",
				"-row",
				"-rowspan",
				"-sticky",
			}},
		},
		"image":   {ignore: true}, //MAYBE later
		"keysyms": {ignore: true}, //MAYBE later
		"loadTk":  {ignore: true},
		"lower":   {manual: true}, // done
		"menu": {widget: true,
			commands: cmdOpts{
				"MenuWidget.AddCommand":   menuOptions,
				"MenuWidget.AddCascade":   menuOptions,
				"MenuWidget.AddSeparator": menuOptions,
				"MenuWidget.Invoke":       nil,
			},
		},
		"messageBox": {
			manual: true, // done
			commands: cmdOpts{"MessageBox": []string{
				"-command",
				"-default",
				"-detail",
				"-icon",
				"-message",
				"-parent",
				"-title",
				"-type",
			}},
		},
		"nsimage":    {ignore: true}, //TODO
		"option":     {ignore: true},
		"optionMenu": {manual: true}, //TODO
		"options":    {ignore: true},
		"pack": {
			manual: true, // done
			commands: cmdOpts{"Pack": []string{
				"-after",
				"-anchor",
				"-before",
				"-expand",
				"-fill",
				"-in",
				"-ipadx",
				"-ipady",
				"-padx",
				"-pady",
				"-side",
			}},
		},
		"palette": {ignore: true}, //MAYBE later
		"photo": {
			manual: true, // done
			commands: cmdOpts{"NewPhoto": []string{
				"-data",
				"-format",
				"-file",
				"-gamma",
				"-height",
				"-metadata",
				"-palette",
				"-width",
			}},
		},
		"place": {
			manual: true, // done
			commands: cmdOpts{"Place": []string{
				"-anchor",
				"-bordermode",
				"-height",
				"-in",
				"-relheight",
				"-relwidth",
				"-relx",
				"-rely",
				"-width",
				"-x",
				"-y",
			}},
		},
		"popup": {manual: true}, //TODO
		"print": {manual: true}, //TODO
		"raise": {manual: true}, // done
		"selection": {
			ignore: true, //MAYBE later
			// commands: cmdOpts{
			// 	"SelectionClear": []string{
			// 		"-displayof",
			// 		"-selection",
			// 	},
			// 	"SelectionGet": []string{
			// 		"-displayof",
			// 		"-selection",
			// 		"-type",
			// 	},
			// 	"SelectionHandle": []string{
			// 		"-selection",
			// 		"-type",
			// 		"-format",
			// 	},
			// 	"SelectionOwn": []string{
			// 		"-command",
			// 		"-displayof",
			// 		"-selection",
			// 	},
			// },
		},
		"send":      {ignore: true},
		"sysnotify": {manual: true}, //TODO
		"systray":   {manual: true}, //TODO
		"text": {widget: true, //MAYBE more
			commands: cmdOpts{
				"TextWidget.TagConfigure": []string{
					"-background",
					"-bgstipple",
					"-borderwidth",
					"-elide",
					"-fgstipple",
					"-font",
					"-foreground",
					"-justify",
					"-lmargin1",
					"-lmargin2",
					"-lmargincolor",
					"-offset",
					"-overstrike",
					"-overstrikefg",
					"-relief",
					"-rmargin",
					"-rmargincolor",
					"-selectbackground",
					"-selectforeground",
					"-spacing1",
					"-spacing2",
					"-spacing3",
					"-tabs",
					"-tabstyle",
					"-underline",
					"-underlinefg",
					"-wrap",
				},
			},
		},
		"tk": {
			manual: true, //MAYBE later
			// commands: cmdOpts{
			// 	"Caret": []string{
			// 		"-x",
			// 		"-y",
			// 		"-height",
			// 	},
			// 	"Inactive": []string{
			// 		"-displayof",
			// 		"-reset",
			// 	},
			// 	"Scaling": []string{
			// 		"-displayof",
			// 	},
			// 	"UseInputMethods": []string{
			// 		"-displayof",
			// 	},
			// },
		},
		"tk_mac":     {ignore: true},
		"tkerror":    {ignore: true},
		"tkvars":     {ignore: true},
		"tkwait":     {manual: true}, // done
		"ttk_image":  {ignore: true},
		"ttk_intro":  {ignore: true},
		"ttk_style":  {manual: true}, //TODO
		"ttk_vsapi":  {ignore: true},
		"ttk_widget": {ignore: true},
		"winfo":      {manual: true}, //TODO
		"wm":         {manual: true}, //TODO
	}

	menuOptions = []string{
		"-activebackground",
		"-activeforeground",
		"-accelerator",
		"-background",
		"-bitmap",
		"-columnbreak",
		"-command",
		"-compound",
		"-font",
		"-foreground",
		"-hidemargin",
		"-image",
		"-indicatoron",
		"-label",
		"-menu",
		"-offvalue",
		"-onvalue",
		"-selectcolor",
		"-selectimage",
		"-state",
		"-underline",
		"-value",
		"-variable",
	}

	handlers = map[string]bool{
		"Command":         true,
		"Invalidcommand":  true,
		"Postcommand":     true,
		"Tearoffcommand":  true,
		"Validatecommand": true,
		"Xscrollcommand":  true,
		"Yscrollcommand":  true,
	}

	hideOpts = map[string]bool{
		"Data": true,
		"Font": true,
		"From": true,
		"To":   true,
	}

	hideOptMethods = map[string]bool{
		"Textvariable": true,
	}

	replaceOpt = map[string]string{
		"Button":  "Btn",
		"Label":   "Lbl",
		"Menu":    "Mnu",
		"Message": "Msg",
		"Text":    "Txt",
	}
)

func main() {
	hashes()
	makeTokenizer()
	w := bytes.NewBuffer(nil)
	w.WriteString(header)
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("PANIC: %v\n%s", err, debug.Stack())
			return
		}

		fmt.Printf("writing %v len=%v\n", ofn, len(w.Bytes()))
		if err := os.WriteFile(ofn, w.Bytes(), 0660); err != nil {
			panic(err)
		}
	}()

	nFilesDir := filepath.Join(xdg.ConfigHome, "ccgo", "v4", "libtk9.0", goos, goarch, libtk.Version, "doc")
	fmt.Printf("nFilesDir=%s\n", nFilesDir)
	util.MustShell(true, nil, "sh", "-c", fmt.Sprintf("rm -rf %s", tempDir))
	util.MustShell(true, nil, "mkdir", "-p", tempDir)
	fmt.Printf("nFilesDir=%s tempDir=%s\n", nFilesDir, tempDir)
	m, err := filepath.Glob(filepath.Join(nFilesDir, "*.n"))
	if err != nil {
		panic(err)
	}

	slices.Sort(m)
	t, err := ngrab.NewTask(io.Discard, m)
	if err != nil {
		panic(err)
	}

	if err := t.Main(); err != nil {
		panic(err)
	}

	htmlFiles := makeHTML(t)
	j := newJob(w, htmlFiles)
	j.main()
}

func makeTokenizer() {
	args := []string{
		"-lexstring", "mlToken",
		"-pkg", "tk9_0",
		`([^$]|\\\$)*`,              // Not TeX, incl. "\$"
		`\$([^$]|\\\$)*[^\\]\$`,     // $TeX$, incl. $Te\$X$
		`\$\$([^$]|\\\$)*[^\\]\$\$`, // $$TeX$$, incl. $Te\$X$
	}
	var b bytes.Buffer
	rc, err := rec.Main(args, &b, io.Discard)
	if err != nil {
		panic(err)
	}

	if rc != 0 {
		panic(rc)
	}

	if err = os.WriteFile("mltoken.go", b.Bytes(), 0660); err != nil {
		panic(err)
	}
}

func makeHTML(t *ngrab.Task) (htmlFiles []string) {
	var documentFN string
	var w *bytes.Buffer
	var path []string
	for _, v := range t.Nodes {
		switch v.Type {
		case "document":
			if documentFN != "" {
				base := filepath.Base(documentFN)
				base = base[:len(base)-len(".n")] + ".html"
				w.WriteString("<html>\n")
				ofn := filepath.Join(tempDir, base)
				if err := os.WriteFile(ofn, w.Bytes(), 0660); err != nil {
					panic(err)
				}

				htmlFiles = append(htmlFiles, ofn)
			}
			documentFN = v.Text
			w = bytes.NewBuffer(nil)
			w.WriteString(head)
		case
			"BS",
			"CS",
			"DS",
			"RS",
			"SO":

			path = append(path, v.Type)
			fmt.Fprintf(w, "<div class=%q title=%q>\n", v.Type, strings.Join(path, "/"))
		case
			"BE",
			"CE",
			"DE",
			"RE",
			"SE":

			path = path[:len(path)-1]
			fmt.Fprintf(w, "</div>\n")
		case "so":
			fmt.Fprintf(w, "<p class=%q title=%q>%s</p>", "sso", strings.Join(append(path, "sso"), "/"), html.EscapeString(v.Text))
		default:
			fmt.Fprintf(w, "<p class=%q title=%q>%s</p>\n", v.Type, strings.Join(append(path, v.Type), "/"), html.EscapeString(v.Text))
		}
	}
	if documentFN != "" {
		base := filepath.Base(documentFN)
		base = base[:len(base)-len(".n")] + ".html"
		w.WriteString("<html>\n")
		ofn := filepath.Join(tempDir, base)
		if err := os.WriteFile(ofn, w.Bytes(), 0660); err != nil {
			panic(err)
		}

		htmlFiles = append(htmlFiles, ofn)
	}
	return htmlFiles
}

func walk(lvl int, n *html.Node, visitor func(n *html.Node) (dive bool)) {
	for ; n != nil; n = n.NextSibling {
		if visitor(n) {
			walk(lvl+1, n.FirstChild, visitor)
		}
	}
}

func class(n *html.Node) string {
	for _, v := range n.Attr {
		if v.Key == "class" {
			return v.Val
		}
	}
	return ""
}

type document struct {
	fn   string
	root *html.Node
	sh   []*html.Node
	shx  map[string]*html.Node
	so   []*html.Node
}

type option struct {
	docs    []string
	goName  string
	tclName string
	xref    map[string]struct{}

	isWidgetOption bool
}

type job struct {
	documents        []*document
	files            []string
	o                *bytes.Buffer
	optionsDoc       *document
	optionsByTclName map[string]*option
	ttkOptionsDoc    *document
}

func newJob(o *bytes.Buffer, files []string) *job {
	return &job{
		files:            files,
		o:                o,
		optionsByTclName: map[string]*option{},
	}
}

func (j *job) registerOption(tclName string, docs []string, xref string) (r *option) {
	switch {
	case strings.HasPrefix(xref, "[."):
		xref = "[Window." + xref[len("[."):]
	}
	r = j.optionsByTclName[tclName]
	if r == nil {
		r = &option{
			tclName: tclName,
			goName:  tclOptName2GoName(tclName),
			docs:    docs,
			xref:    map[string]struct{}{},
		}
		j.optionsByTclName[tclName] = r
	}
	if xref != "" {
		r.xref[xref] = struct{}{}
	}
	return r
}

func (j *job) w(s string, args ...any) {
	fmt.Fprintf(j.o, s, args...)
}

func (j *job) analyze(file string) {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	n, err := html.Parse(f)
	if err != nil {
		panic(err)
	}

	d := &document{
		fn:   file,
		root: n,
		shx:  map[string]*html.Node{},
	}
	switch filepath.Base(file) {
	case "options.html":
		j.optionsDoc = d
	case "ttk_widget.html":
		j.ttkOptionsDoc = d
	}
	j.documents = append(j.documents, d)
	walk(0, n, func(n *html.Node) (dive bool) {
		switch n.Type {
		case html.ElementNode:
			switch n.Data {
			case "p":
				switch class(n) {
				case "SH":
					s := strings.TrimSpace(n.FirstChild.Data)
					if d.shx[s] != nil {
						panic("internal error")
					}

					d.shx[s] = n
					d.sh = append(d.sh, n)
				}
			case "div":
				switch class(n) {
				case "SO":
					d.so = append(d.so, n)
				}
			}
		}
		return true
	})
}

func base(fn string) (r string) {
	x := strings.Index(fn, "/")
	r = fn[x+1:]
	return r[:len(r)-len(".html")]
}

func (j *job) main() {
	for _, v := range j.files {
		j.analyze(v)
	}
	for _, v := range j.documents {
		var a []string
		for k := range v.shx {
			a = append(a, k)
		}
		slices.Sort(a)
		if len(v.so) != 0 {
			a = append(a, fmt.Sprintf("len(so)=%v", len(v.so)))
		}
		// fmt.Printf("%s %v\n", v.fn, a)
		base := base(v.fn)
		if nfo := pageInfos[base]; nfo == nil && len(v.so) != 0 {
			pageInfos[base] = &pageInfo{widget: true}
		}
	}
	var fail bool
	for _, fn := range j.files {
		base := base(fn)
		nfo := pageInfos[base]
		switch {
		case nfo == nil:
			fmt.Printf("%s: missing page info\n", base)
		case nfo.ignore, nfo.manual, nfo.widget:
			// ok
		default:
			fmt.Printf("%s: missing classificaiotn\n", base)
			fail = true
		}
	}
	if fail {
		return
	}

	for i, fn := range j.files {
		base := base(fn)
		nfo := pageInfos[base]
		switch {
		case nfo.widget:
			j.widget(fn, j.documents[i])
			if len(nfo.commands) != 0 {
				j.manual(fn, j.documents[i], nfo)
			}
		case nfo.manual:
			j.manual(fn, j.documents[i], nfo)
		}
	}
	j.stdOptions()
	j.ttkStdOptions()
	var a []string
	for k := range j.optionsByTclName {
		a = append(a, k)
	}
	slices.Sort(a)
	for _, k := range a {
		v := j.optionsByTclName[k]
		if hideOpts[v.goName] || hideOptMethods[v.goName] {
			continue
		}

		j.w("\n\n// %s option.", v.goName)
		if handlers[v.goName] {
			j.w("\n//\n// See also [Event handlers].")
		}
		if len(v.docs) != 0 {
			j.w("\n//\n// %s", strings.Join(v.docs, "\n//"))
		}
		if len(v.xref) != 0 {
			j.w("\n//\n// Known uses:")
			var a []string
			for k := range v.xref {
				a = append(a, k)
			}
			slices.Sort(a)
			for _, v := range a {
				j.w("\n//  - %s", v)
			}
		}
		switch ok := handlers[v.goName]; {
		case ok:
			j.w("\n//\n// [Event handlers]: https://pkg.go.dev/github.com/gg582/tk9.0#hdr-Event_handlers")
			j.w("\nfunc %s(handler any) Opt {", v.goName)
			j.w("\n\treturn newEventHandler(%q, handler)", v.tclName)
			j.w("\n}")
		default:
			j.w("\nfunc %s(val any) Opt {", v.goName)
			j.w("\nreturn rawOption(fmt.Sprintf(`%s %%s`, optionString(val)))", v.tclName)
			j.w("\n}")

			if !v.isWidgetOption {
				break
			}

			j.w("\n\n// %s — Get the configured option value.", v.goName)
			if len(v.xref) != 0 {
				j.w("\n//\n// Known uses:")
				var a []string
				for k := range v.xref {
					a = append(a, k)
				}
				slices.Sort(a)
				for _, v := range a {
					if !strings.Contains(v, "command specific") {
						j.w("\n//  - %s", v)
					}
				}
			}
			j.w("\nfunc (w *Window) %s() string {", v.goName)
			j.w("\nreturn evalErr(fmt.Sprintf(`%%s cget %s`, w))", v.tclName)
			j.w("\n}")
		}
	}
}

func (j *job) ttkStdOptions() {
	walk(0, j.ttkOptionsDoc.root, func(n *html.Node) (dive bool) {
		if nodeIs(n, "OP") {
			tclName := strings.TrimSpace(n.FirstChild.Data)
			a := strings.Fields(tclName)
			tclName = strings.TrimLeft(a[0], `"\`)
			var docs []string
			for n := n.NextSibling; n != nil; n = n.NextSibling {
				if nodeIs(n, "OP") || nodeIs(n, "SH") {
					break
				}

				if !nodeIs(n, "text") {
					continue
				}

				s := strings.TrimSpace(plain(n.FirstChild.Data))
				docs = append(docs, s)
			}
			o := j.registerOption(tclName, nil, "")
			o.isWidgetOption = true
			if len(o.docs) == 0 {
				o.docs = docs
			}
		}
		return true
	})
}

func (j *job) stdOptions() {
	walk(0, j.optionsDoc.root, func(n *html.Node) (dive bool) {
		if nodeIs(n, "OP") {
			tclName := strings.TrimSpace(n.FirstChild.Data)
			a := strings.Fields(tclName)
			tclName = strings.TrimLeft(a[0], `"\`)
			var docs []string
			for n := n.NextSibling; n != nil; n = n.NextSibling {
				if nodeIs(n, "OP") || nodeIs(n, "SH") {
					break
				}

				if !nodeIs(n, "text") {
					continue
				}

				s := strings.TrimSpace(plain(n.FirstChild.Data))
				docs = append(docs, s)
			}
			o := j.registerOption(tclName, nil, "")
			o.isWidgetOption = true
			o.docs = docs
		}
		return true
	})
}

func nodeIs(n *html.Node, pathSuffix string) bool {
	for _, v := range n.Attr {
		if v.Key == "title" && strings.HasSuffix(v.Val, pathSuffix) {
			return true
		}
	}

	return false
}

func (j *job) widgetStdOpts(xref string, doc *document) (r []string) {
	m := map[string]struct{}{}
	walk(0, doc.root, func(n *html.Node) (dive bool) {
		if nodeIs(n, "SO/text") {
			s := strings.TrimSpace(n.FirstChild.Data)
			a := strings.Fields(s)
			for _, v := range a {
				s := strings.TrimSpace(v)
				if s[0] == '\\' {
					s = s[1:]
				}
				j.registerOption(s, nil, xref).isWidgetOption = true
				s = tclOptName2GoName(s)
				m[s] = struct{}{}
			}
		}
		return true
	})
	for k := range m {
		r = append(r, k)
	}
	slices.Sort(r)
	return r
}

func (j *job) widgetStylingOptions(doc *document) (r []string) {
	for _, n := range doc.sh {
		if ch := n.FirstChild; ch != nil && strings.Contains(ch.Data, "STYLING OPTIONS") {
			for n := n.NextSibling; n != nil; n = n.NextSibling {
				if hasClass(n, "SH") {
					break
				}

				if hasClass(n, "PP") {
					r = append(r, "")
					continue
				}

				if ch := n.FirstChild; ch != nil {
					s := strings.TrimSpace(plain(ch.Data))
					if strings.HasPrefix(s, "-") {
						s = s[1:]
						a := strings.Fields(s)
						s = fmt.Sprintf(" - [%s] %s", capitalize(a[0]), strings.Join(a[1:], " "))
						for len(r) != 0 && r[len(r)-1] == "" {
							r = r[:len(r)-1]
						}
					}
					r = append(r, s)
				}
			}

			break
		}
	}
	return r
}

func capitalize(s string) string {
	return strings.ToUpper(string(s[0])) + s[1:]
}

func (j *job) widgetStdStyles(doc *document) (r []string) {
	for _, n := range doc.sh {
		if ch := n.FirstChild; ch != nil && strings.Contains(ch.Data, "STANDARD STYLES") {
			for n := n.NextSibling; n != nil; n = n.NextSibling {
				if hasClass(n, "SH") {
					break
				}

				if hasClass(n, "PP") {
					r = append(r, "")
					continue
				}

				if ch := n.FirstChild; ch != nil {
					r = append(r, strings.TrimSpace(plain(ch.Data)))
				}
			}

			break
		}
	}
	return r
}

func (j *job) widgetSpecificOpts(xref string, doc *document) (r []*option) {
	walk(0, doc.root, func(n *html.Node) (dive bool) {
		if nodeIs(n, "OP") {
			tclName := strings.TrimSpace(n.FirstChild.Data)
			a := strings.Fields(tclName)
			tclName = strings.TrimLeft(a[0], `"\`)
			var docs []string
			for n := n.NextSibling; n != nil; n = n.NextSibling {
				if nodeIs(n, "OP") || nodeIs(n, "SH") {
					break
				}

				if !nodeIs(n, "text") {
					continue
				}

				s := strings.TrimSpace(plain(n.FirstChild.Data))
				docs = append(docs, s)
			}
			o := j.registerOption(tclName, nil, xref)
			o.isWidgetOption = true
			p := *o
			p.docs = docs
			r = append(r, &p)
		}
		return true
	})
	return r
}

func tclOptName2GoName(s string) (r string) {
	s = strings.TrimSpace(s)
	if strings.HasPrefix(s, "\\") {
		s = s[1:]
	}
	if strings.HasPrefix(s, "-") {
		s = s[1:]
	}
	r = export(s)
	if x := replaceOpt[r]; x != "" {
		r = x
	}
	return r
}

func (j *job) pageLink(page string) {
	if page == "" {
		return
	}

	page = page[:len(page)-len(".html")]
	j.w("\n//\n// More information might be available at the [Tcl/Tk %s] page.", page)
	j.w("\n//\n// [Tcl/Tk %s]: https://www.tcl.tk/man/tcl9.0/TkCmd/%[1]s.html", page)
}

func (j *job) manual(fn string, doc *document, nfo *pageInfo) {
	for nm, options := range nfo.commands {
		for _, option := range options {
			j.registerOption(option, nil, fmt.Sprintf("[%s] (command specific)", nm))
		}
	}
}

func (j *job) widget(fn string, doc *document) {
	shNameNode := doc.shx["NAME"]
	txtNode := shNameNode.NextSibling.NextSibling.FirstChild
	txt := txtNode.Data
	txt = strings.TrimSpace(txt[strings.Index(txt, "-")+1:])
	txt = strings.ToUpper(txt[:1]) + txt[1:]
	base := base(fn)
	gnm := tclName2GoName(base)
	doc0 := []string{fmt.Sprintf("\n\n// %s — %s", gnm, txt)}
	doc0 = append(doc0, j.description(doc.shx["DESCRIPTION"])...)
	j.w("%s", strings.Join(doc0, "\n// "))
	j.w("\n//\n// Use [Window.%s] to create a %[1]s with a particular parent.", gnm)
	j.pageLink(fn[len(tempDir)+1:])
	if opts := j.widgetStdOpts(fmt.Sprintf("[%s]", gnm), doc); len(opts) != 0 {
		j.w("\n//\n// # Standard options\n//")
		for _, v := range opts {
			if hideOpts[v] {
				continue
			}

			j.w("\n//  - [%s]", v)
		}
	}
	if sos := j.widgetSpecificOpts(fmt.Sprintf("[%s] (widget specific)", gnm), doc); len(sos) != 0 {
		j.w("\n//\n// # Widget specific options")
		for _, v := range sos {
			j.w("\n//\n// [%s]", v.goName)
			j.w("\n//\n// %s", strings.Join(v.docs, "\n// "))
		}
	}
	if ss := j.widgetStdStyles(doc); len(ss) != 0 {
		j.w("\n//\n// # Standard styles\n//")
		for _, v := range ss {
			j.w("\n// %s", v)
		}
	}
	if sos := j.widgetStylingOptions(doc); len(sos) != 0 {
		j.w("\n//\n// # Styling options\n//")
		for _, v := range sos {
			j.w("\n// %s", v)
		}
	}
	j.w("\nfunc %s(options ...Opt) *%[1]sWidget {", gnm)
	j.w("\nreturn App.%s(options...)", gnm)
	j.w("\n}")

	j.w("%s", doc0[0])
	j.w("\n//\n// The resulting [Window] is a child of 'w'")
	j.w("\n//\n// For details please see [%s]", gnm)
	j.w("\nfunc (w *Window) %s(options ...Opt) *%[1]sWidget {", gnm)
	cmd := strings.Replace(base, "ttk_", "ttk::", 1)
	j.w("\nreturn &%sWidget{w.newChild(%q, options...)}", gnm, cmd)
	j.w("\n}")
	j.w("\n\n// %sWidget represents the Tcl/Tk %s widget/window", gnm, base)
	j.w("\ntype %sWidget struct {", gnm)
	j.w("\n*Window")
	j.w("\n}")
}

func (j *job) description(n *html.Node) (r []string) {
	if n == nil {
		return nil
	}

	r = append(r, "", "# Description\n//")
	stop := false
	walk(0, n.NextSibling, func(n *html.Node) (dive bool) {
		if stop {
			return false
		}

		if hasClass(n, "SH") {
			stop = true
			return false
		}

		if hasClass(n, "PP") {
			r = append(r, "")
			return false
		}

		switch n.Type {
		case html.TextNode:
			if s := strings.TrimSpace(plain(n.Data)); s != "" &&
				!strings.HasPrefix(s, `\"`) &&
				!strings.HasPrefix(s, "-") {
				r = append(r, s)
			}
		}
		return true
	})
	return r
}

func hasClass(n *html.Node, cls string) bool {
	for _, v := range n.Attr {
		if v.Key == "class" && v.Val == cls {
			return true
		}
	}

	return false
}

func plain(s string) (r string) {
	r = s
	r = strings.ReplaceAll(r, "\\-", "-")
	r = strings.ReplaceAll(r, "\\fB", "")
	r = strings.ReplaceAll(r, "\\fI", "")
	r = strings.ReplaceAll(r, "\\fP", "")
	r = strings.ReplaceAll(r, "\\fR", "")
	return r
}

func tclName2GoName(s string) string {
	switch {
	case strings.HasPrefix(s, "ttk_"):
		return "T" + export(s[len("ttk_"):])
	default:
		return export(s)
	}
}

func export(s string) (r string) {
	a := strings.Split(s, "_")
	for i, v := range a {
		a[i] = strings.ToUpper(v[:1]) + v[1:]
	}
	return strings.Join(a, "")
}

// origin returns caller's short position, skipping skip frames.
//
//lint:ignore U1000 debug helper
func origin(skip int) string {
	pc, fn, fl, _ := runtime.Caller(skip)
	f := runtime.FuncForPC(pc)
	var fns string
	if f != nil {
		fns = f.Name()
		if x := strings.LastIndex(fns, "."); x > 0 {
			fns = fns[x+1:]
		}
		if strings.HasPrefix(fns, "func") {
			num := true
			for _, c := range fns[len("func"):] {
				if c < '0' || c > '9' {
					num = false
					break
				}
			}
			if num {
				return origin(skip + 2)
			}
		}
	}
	return fmt.Sprintf("%s:%d:%s", filepath.Base(fn), fl, fns)
}

// todo prints and return caller's position and an optional message tagged with TODO. Output goes to stderr.
//
//lint:ignore U1000 debug helper
func todo(s string, args ...interface{}) string {
	switch {
	case s == "":
		s = fmt.Sprintf(strings.Repeat("%v ", len(args)), args...)
	default:
		s = fmt.Sprintf(s, args...)
	}
	r := fmt.Sprintf("%s\n\tTODO %s", origin(2), s)
	// fmt.Fprintf(os.Stderr, "%s\n", r)
	// os.Stdout.Sync()
	return r
}

// trc prints and return caller's position and an optional message tagged with TRC. Output goes to stderr.
//
//lint:ignore U1000 debug helper
func trc(s string, args ...interface{}) string {
	switch {
	case s == "":
		s = fmt.Sprintf(strings.Repeat("%v ", len(args)), args...)
	default:
		s = fmt.Sprintf(s, args...)
	}
	r := fmt.Sprintf("%s: TRC %s", origin(2), s)
	fmt.Fprintf(os.Stderr, "%s\n", r)
	os.Stderr.Sync()
	return r
}

func hashes() {
	filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			panic(err)
		}

		base := filepath.Base(path)
		if d.IsDir() || !strings.HasSuffix(base, ".zip") {
			return nil
		}

		r, err := zip.OpenReader(path)
		if err != nil {
			panic(err)
		}

		defer r.Close()

		fmt.Printf("// %s\n", path)
		for _, f := range r.File {
			// fmt.Printf("Contents of %s:\n", f.Name)
			r, err := f.Open()
			if err != nil {
				panic(err)
			}

			h := sha256.New()
			_, err = io.Copy(h, r)
			r.Close()
			if err != nil {
				panic(err)
			}

			fmt.Printf("%q: \"%0x\",\n", f.Name, h.Sum(nil))
		}

		return nil
	})
	fmt.Printf("// other\n")
	fmt.Printf("%q: \"%0x\",\n", "tcl_library.zip", sha256.Sum256([]byte(tcllib.Zip)))
	fmt.Printf("%q: \"%0x\",\n", "tk_library.zip", sha256.Sum256([]byte(tklib.Zip)))
}
