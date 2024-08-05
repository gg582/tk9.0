# tk9.0: CGo-free cross platform GUI for Go

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
     9		var b *Window
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
    11		App.Configure(Padx("4m"), Pady("4m")).Center().Wait()
    12	}

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
    11		App.Configure(Padx("4m"), Pady("4m")).Center().Wait()
    12	}

![canvas-plot](_examples/canvas_plot.png "canvas plot")

Graphing to a canvas using gnuplot  (_examples/canvas_plot.go).

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
    16	set urange [ -3.14159 : 3.14159 ] noreverse nowriteback
    17	set vrange [ -3.14159 : 3.14159 ] noreverse nowriteback
    18	set xrange [ * : * ] noreverse writeback
    19	set x2range [ * : * ] noreverse writeback
    20	set yrange [ * : * ] noreverse writeback
    21	set y2range [ * : * ] noreverse writeback
    22	set zrange [ * : * ] noreverse writeback
    23	set cbrange [ * : * ] noreverse writeback
    24	set rrange [ * : * ] noreverse writeback
    25	set colorbox vertical origin screen 0.9, 0.2 size screen 0.05, 0.6 front  noinvert bdefault
    26	NO_ANIMATION = 1
    27	splot cos(u)+.5*cos(u)*cos(v),sin(u)+.5*sin(u)*cos(v),.5*sin(v) with lines,1+cos(u)+.5*cos(u)*cos(v),.5*sin(v),sin(u)+.5*sin(u)*cos(v) with lines
    28	`
    29	
    30	func main() {
    31		Pack(
    32			Canvas(Width(640), Height(480), Background(White)).Plot(script),
    33			TExit(), Padx("1m"), Pady("2m"), Ipadx("1m"), Ipady("1m"),
    34		)
    35		App.Configure(Padx("4m"), Pady("4m")).Center().Wait()
    36	}


Documentation at [pkg.go.dev].

[pkg.go.dev]: https://pkg.go.dev/modernc.org/tk9.0
