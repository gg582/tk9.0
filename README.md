# tk9.0: The CGo-free cross platform GUI for Go

![photo](_examples/photo.png "photo")

Using embedded images (_examples/photo.go).

     1	package main
     2	
     3	import _ "embed"
     4	import . "modernc.org/tk9.0"
     5	
     6	//go:embed gopher.png
     7	var gopher []byte
     8	
     9	func main() {
    10		Pack(Label(Image(NewPhoto(Data(gopher)))), TExit(), Padx("1m"), Pady("2m"), Ipadx("1m"), Ipady("1m"))
    11		App.Configure(Padx("4m"), Pady("3m")).Center().Wait()
    12	}

![menu](_examples/menu.png "menu")

Cascading menus (_examples/menu.go)

     1	package main
     2	
     3	import (
     4		"fmt"
     5		"runtime"
     6		. "modernc.org/tk9.0"
     7	)
     8	
     9	func main() {
    10		menubar := Menu()
    11	
    12		fileMenu := menubar.Menu()
    13		fileMenu.AddCommand(Lbl("New"), Underline(0), Accelerator("Ctrl+N"))
    14		fileMenu.AddCommand(Lbl("Open..."), Underline(0), Accelerator("Ctrl+O"), Command(func() { GetOpenFile() }))
    15		Bind(App, "<Control-o>", Command(func() { fileMenu.Invoke(1) }))
    16		fileMenu.AddCommand(Lbl("Save"), Underline(0), Accelerator("Ctrl+S"))
    17		fileMenu.AddCommand(Lbl("Save As..."), Underline(5))
    18		fileMenu.AddCommand(Lbl("Close"), Underline(0), Accelerator("Crtl+W"))
    19		fileMenu.AddSeparator()
    20		fileMenu.AddCommand(Lbl("Exit"), Underline(1), Accelerator("Ctrl+Q"), ExitHandler())
    21		Bind(App, "<Control-q>", Command(func() { fileMenu.Invoke(6) }))
    22		menubar.AddCascade(Lbl("File"), Underline(0), Mnu(fileMenu))
    23	
    24		editMenu := menubar.Menu()
    25		editMenu.AddCommand(Lbl("Undo"))
    26		editMenu.AddSeparator()
    27		editMenu.AddCommand(Lbl("Cut"))
    28		editMenu.AddCommand(Lbl("Copy"))
    29		editMenu.AddCommand(Lbl("Paste"))
    30		editMenu.AddCommand(Lbl("Delete"))
    31		editMenu.AddCommand(Lbl("Select All"))
    32		menubar.AddCascade(Lbl("Edit"), Underline(0), Mnu(editMenu))
    33	
    34		helpMenu := menubar.Menu()
    35		helpMenu.AddCommand(Lbl("Help Index"))
    36		helpMenu.AddCommand(Lbl("About..."))
    37		menubar.AddCascade(Lbl("Help"), Underline(0), Mnu(helpMenu))
    38	
    39		App.WmTitle(fmt.Sprintf("%s on %s", App.WmTitle(""), runtime.GOOS))
    40		App.Configure(Mnu(menubar), Width("8c"), Height("6c"), Padx("4m"), Pady("3m")).Center().Wait()
    41	}

![svg](_examples/text.png "text")

Rich text using markup (_examples/text.go).

     1	// Adapted from tk library demos style.tcl
     2	//
     3	// See LICENSE-TCLTK for copyright info.
     4	
     5	package main
     6	
     7	import "strings"
     8	import . "modernc.org/tk9.0"
     9	
    10	func main() {
    11		family := "Helvetica"
    12		var scroll *TScrollbarWidget
    13		t := Text(Fnt(family+" 10"), Setgrid(true), Wrap("word"), Padx("2m"), Pady("2m"), Yscrollcommand(func(e *Event) { e.ScrollSet(scroll) }))
    14		scroll = TScrollbar(Command(func(e *Event) { e.Yview(t) }))
    15		Grid(t, Sticky("news"), Pady("2m"))
    16		Grid(scroll, Row(0), Column(1), Sticky("nes"), Pady("2m"))
    17		GridRowConfigure(App, 0, Weight(1))
    18		GridColumnConfigure(App, 0, Weight(1))
    19		Grid(TExit(), Padx("1m"), Pady("2m"), Ipadx("1m"), Ipady("1m"))
    20		t.TagConfigure("bgstipple", Background(Black), Borderwidth(0), Bgstipple(Gray12))
    21		t.TagConfigure("big", Fnt(family+" 12 bold"))
    22		t.TagConfigure("bold", Fnt(family+" 10 bold italic"))
    23		t.TagConfigure("center", Justify("center"))
    24		t.TagConfigure("color1", Background("#a0b7ce"))
    25		t.TagConfigure("color2", Foreground(Red))
    26		t.TagConfigure("margins", Lmargin1("12m"), Lmargin2("6m"), Rmargin("10m"))
    27		t.TagConfigure("overstrike", Overstrike(1))
    28		t.TagConfigure("raised", Relief("raised"), Borderwidth(1))
    29		t.TagConfigure("right", Justify("right"))
    30		t.TagConfigure("spacing", Spacing1("10p"), Spacing2("2p"), Lmargin1("12m"), Lmargin2("6m"), Rmargin("10m"))
    31		t.TagConfigure("sub", Offset("-2p"), Fnt(family+" 8"))
    32		t.TagConfigure("sunken", Relief("sunken"), Borderwidth(1))
    33		t.TagConfigure("super", Offset("4p"), Fnt(family+" 8"))
    34		t.TagConfigure("tiny", Fnt("Times 8 bold"))
    35		t.TagConfigure("underline", Underline(1))
    36		t.TagConfigure("verybig", Fnt("Courier 22 bold"))
    37		t.InsertML(`Text widgets like this one allow you to display information in a variety of styles. Display styles are controlled using a mechanism called
    38	<bold>tags</bold>. Tags are just textual names that you can apply to one or more ranges of characters within a text widget. You can configure tags
    39	with various display styles. If you do this, then the tagged characters will be displayed with the styles you chose. The available display styles are:
    40	
    41	<br><br><big>1. Font.</big> You can choose any system font, <verybig>large</verybig> or <tiny>small</tiny>.
    42	
    43	<br><br><big>2. Color.</big> You can change either the <color1>background</color1> or <color2>foreground</color2> color, or
    44	<color1><color2>both</color2></color1>.
    45	
    46	<br><br><big>3. Stippling.</big> You can cause the <bgstipple>background</bgstipple> information to be drawn with a stipple fill instead of a solid
    47	fill.
    48	
    49	<br><br><big>4. Underlining.</big> You can <underline>underline</underline> ranges of text.
    50	
    51	<br><br><big>5. Overstrikes.</big> You can <overstrike>draw lines through</overstrike> ranges of text.
    52	
    53	<br><br><big>6. 3-D effects.</big> You can arrange for the background to be drawn with a border that makes characters appear either
    54	<raised>raised</raised> or <sunken>sunken</sunken>.
    55	
    56	<br><br><big>7. Justification.</big> You can arrange for lines to be displayed <br>left-justified <br><right>right-justified, or</right>
    57	<br><center>centered.</center>
    58	
    59	<br><br><big>8. Superscripts and subscripts.</big> You can control the vertical position of text to generate superscript effects like
    60	10<super>n</super> or subscript effects like X<sub>i</sub>.
    61	
    62	<br><br><big>9. Margins.</big> You can control the amount of extra space left on each side of the text
    63	
    64	<br><br><margins>This paragraph is an example of the use of margins. It consists of a single line of text that wraps around on the screen. There are
    65	two separate left margin values, one for the first display line associated with the text line, and one for the subsequent display lines, which occur
    66	because of wrapping. There is also a separate specification for the right margin, which is used to choose wrap points for lines.</margins>
    67	
    68	<br><br><big>10. Spacing.</big> You can control the spacing of lines with three separate parameters. "Spacing1" tells how much extra space to leave
    69	above a line, "spacing3" tells how much space to leave below a line, and if a text line wraps, "spacing2" tells how much space to leave between the
    70	display lines that make up the text line.
    71	
    72	<br><spacing>These indented paragraphs illustrate how spacing can be used. Each paragraph is actually a single line in the text widget, which is
    73	word-wrapped by the widget.</spacing>
    74	
    75	<br><spacing>Spacing1 is set to 10 points for this text, which results in relatively large gaps between the paragraphs. Spacing2 is set to 2 points,
    76	which results in just a bit of extra space within a pararaph. Spacing3 isn't used in this example.</spacing>
    77	
    78	<br><spacing>To see where the space is, select ranges of text within these paragraphs. The selection highlight will cover the extra space.</spacing>`)
    79		App.Configure(Padx("4m"), Pady("3m")).Center().Wait()
    80	}
    81	
    82	func s(s string) string {
    83		return strings.ReplaceAll(s, "\n", " ")
    84	}

![svg](_examples/svg.png "svg")

Using svg (_examples/svg.go).

     1	package main
     2	
     3	import . "modernc.org/tk9.0"
     4	
     5	// https://en.wikipedia.org/wiki/SVG
     6	const svg = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
     7	<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
     8	<svg width="391" height="391" viewBox="-70.5 -70.5 391 391" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink">
     9	<rect fill="#fff" stroke="#000" x="-70" y="-70" width="390" height="390"/>
    10	<g opacity="0.8">
    11		<rect x="25" y="25" width="200" height="200" fill="lime" stroke-width="4" stroke="pink" />
    12		<circle cx="125" cy="125" r="75" fill="orange" />
    13		<polyline points="50,150 50,200 200,200 200,100" stroke="red" stroke-width="4" fill="none" />
    14		<line x1="50" y1="50" x2="200" y2="200" stroke="blue" stroke-width="4" />
    15	</g>
    16	</svg>`
    17	
    18	func main() {
    19		Pack(
    20			Label(Image(NewPhoto(Data(svg)))),
    21			TExit(), Padx("1m"), Pady("2m"), Ipadx("1m"), Ipady("1m"),
    22		)
    23		App.Configure(Padx("4m"), Pady("3m")).Center().Wait()
    24	}

![calc](_examples/calc.png "calc")

A simple calculator (_examples/calc.go).

     1	package main
     2	
     3	import "github.com/expr-lang/expr"
     4	import 	. "modernc.org/tk9.0"
     5	
     6	func main() {
     7		out := Label(Height(2), Anchor("e"), Txt("(123+232)/(123-10)"))
     8		Grid(out, Columnspan(4), Sticky("e"))
     9		var b *ButtonWidget
    10		for i, c := range "C()/789*456-123+0.=" {
    11			b = Button(Txt(string(c)), Command(
    12				func() {
    13					switch c {
    14					case 'C':
    15						out.Configure(Txt(""))
    16					case '=':
    17						x, err := expr.Eval(out.Txt(), nil)
    18						if err != nil {
    19							MessageBox(Icon("error"), Msg(err.Error()), Title("Error"))
    20							x = ""
    21						}
    22						out.Configure(Txt(x))
    23					default:
    24						out.Configure(Txt(out.Txt() + string(c)))
    25					}
    26				},
    27			))
    28			Grid(b, Row(i/4+1), Column(i%4), Sticky("news"), Ipadx("1.5m"), Ipady("2.6m"))
    29		}
    30		Grid(b, Columnspan(2))
    31		App.Wait()
    32	}

![font](_examples/font.png "font")

A font previewer (_examples/font.go).

     1	package main
     2	
     3	import . "modernc.org/tk9.0"
     4	
     5	func main() {
     6		t := Text(Wrap("none"))
     7		for i, v := range FontFamilies() {
     8			t.TagConfigure(t.TagAdd(t.Insert(LC{i + 1, 0}, v+"\n"), LC{i + 2, 0}), Fnt(NewFont(Family(v))))
     9		}
    10		Pack(t, TExit(), Padx("1m"), Pady("2m"), Ipadx("1m"), Ipady("1m"))
    11		App.Configure(Padx("4m"), Pady("3m")).Center().Wait()
    12	}

![splot](_examples/splot.png "surface plot")

Surface plot (_examples/splot.go).

     1	package main
     2	
     3	import . "modernc.org/tk9.0"
     4	
     5	func main() {
     6		Pack(
     7			Label(Image(NewPhoto(Width(800), Height(600)).Graph("set grid; splot x**2+y**2, x**2-y**2"))),
     8			TExit(), Padx("1m"), Pady("2m"), Ipadx("1m"), Ipady("1m"),
     9		)
    10		App.Configure(Padx("4m"), Pady("3m")).Center().Wait()
    11	}

![tori](_examples/tori.png "interlocked tori")

Interlocked tori plot (_examples/tori.go).

     1	package main
     2	
     3	import . "modernc.org/tk9.0"
     4	
     5	// https://gnuplot.sourceforge.net/demo_5.4/hidden2.html
     6	const script = `
     7	set multiplot title "Interlocking Tori"
     8	set title "PM3D surface\nno depth sorting"
     9	set parametric
    10	set urange [-pi:pi]
    11	set vrange [-pi:pi]
    12	set isosamples 50,20
    13	set origin -0.02,0.0
    14	set size 0.55, 0.9
    15	unset key
    16	unset xtics
    17	unset ytics
    18	unset ztics
    19	set border 0
    20	set view 60, 30, 1.5, 0.9
    21	unset colorbox
    22	set pm3d scansbackward
    23	splot cos(u)+.5*cos(u)*cos(v),sin(u)+.5*sin(u)*cos(v),.5*sin(v) with pm3d,1+cos(u)+.5*cos(u)*cos(v),.5*sin(v),sin(u)+.5*sin(u)*cos(v) with pm3d
    24	set title "PM3D surface\ndepth sorting"
    25	set origin 0.40,0.0
    26	set size 0.55, 0.9
    27	set colorbox vertical user origin 0.9, 0.15 size 0.02, 0.50
    28	set format cb "%.1f"
    29	set pm3d depthorder
    30	splot cos(u)+.5*cos(u)*cos(v),sin(u)+.5*sin(u)*cos(v),.5*sin(v) with pm3d,1+cos(u)+.5*cos(u)*cos(v),.5*sin(v),sin(u)+.5*sin(u)*cos(v) with pm3d
    31	unset multiplot
    32	`
    33	
    34	func main() {
    35		Pack(
    36			Label(Image(NewPhoto(Width(800), Height(600)).Graph(script))),
    37			TExit(), Padx("1m"), Pady("2m"), Ipadx("1m"), Ipady("1m"),
    38		)
    39		App.Configure(Padx("4m"), Pady("3m")).Center().Wait()
    40	}

![tori-canvas](_examples/tori_canvas.png "interlocked tori on canvas")

Interlocked tori plot on canvas (_examples/tori_canvas.go).

     1	package main
     2	
     3	import . "modernc.org/tk9.0"
     4	
     5	// https://gnuplot.sourceforge.net/demo_5.4/surface2.9.gnu
     6	const script = `
     7	set dummy u, v
     8	set key bmargin center horizontal Right noreverse enhanced autotitle nobox
     9	set parametric
    10	set view 50, 30, 1, 1
    11	set isosamples 50, 20
    12	set hidden3d back offset 1 trianglepattern 3 undefined 1 altdiagonal bentover
    13	set style data lines
    14	set xyplane relative 0
    15	set title "Interlocking Tori" 
    16	set grid
    17	set urange [ -3.14159 : 3.14159 ] noreverse nowriteback
    18	set vrange [ -3.14159 : 3.14159 ] noreverse nowriteback
    19	set xrange [ * : * ] noreverse writeback
    20	set x2range [ * : * ] noreverse writeback
    21	set yrange [ * : * ] noreverse writeback
    22	set y2range [ * : * ] noreverse writeback
    23	set zrange [ * : * ] noreverse writeback
    24	set cbrange [ * : * ] noreverse writeback
    25	set rrange [ * : * ] noreverse writeback
    26	set colorbox vertical origin screen 0.9, 0.2 size screen 0.05, 0.6 front  noinvert bdefault
    27	NO_ANIMATION = 1
    28	splot cos(u)+.5*cos(u)*cos(v),sin(u)+.5*sin(u)*cos(v),.5*sin(v) with lines,1+cos(u)+.5*cos(u)*cos(v),.5*sin(v),sin(u)+.5*sin(u)*cos(v) with lines
    29	`
    30	
    31	func main() {
    32		Pack(
    33			Canvas(Width(800), Height(600), Background(White)).Graph(script),
    34			TExit(), Padx("1m"), Pady("2m"), Ipadx("1m"), Ipady("1m"),
    35		)
    36		App.Configure(Padx("4m"), Pady("3m")).Center().Wait()
    37	}

Documentation at [pkg.go.dev].

[pkg.go.dev]: https://pkg.go.dev/modernc.org/tk9.0
