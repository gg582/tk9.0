// Adapted from tk library demos style.tcl
//
// See LICENSE-TCLTK for copyright info.

package main

import (
	"fmt"
	"strings"

	. "modernc.org/tk9.0"
)

func main() {
	family := "Helvetica"
	var scroll *TScrollbarWidget
	// text $w.text -yscrollcommand "$w.scroll set" -setgrid true -width 70 -height 32 -wrap word -font "$family 12"
	t := Text(Fnt(family+" 10"), Setgrid(true), Wrap("word"), Padx("2m"), Pady("2m"), Yscrollcommand(
		func(*Window, any) (any, error) { return fmt.Sprintf("%v set", scroll), nil }),
	)
	// ttk::scrollbar $w.scroll -command "$w.text yview"
	scroll = TScrollbar(Command(func(*Window, any) (any, error) { return fmt.Sprintf("%v yview", t), nil }))
	Grid(t, Sticky("news"), Pady("2m"))
	Grid(scroll, Row(0), Column(1), Sticky("nes"), Pady("2m"))
	GridRowConfigure(App, 0, Weight(1))
	GridColumnConfigure(App, 0, Weight(1))
	Grid(TExit(), Padx("1m"), Pady("2m"), Ipadx("1m"), Ipady("1m"))

	t.TagConfigure("bgstipple", Background(Black), Borderwidth(0), Bgstipple(Gray12))
	t.TagConfigure("big", Fnt(family+" 12 bold"))
	t.TagConfigure("bold", Fnt(family+" 10 bold italic"))
	t.TagConfigure("center", Justify("center"))
	t.TagConfigure("color1", Background("#a0b7ce"))
	t.TagConfigure("color2", Foreground(Red))
	t.TagConfigure("margins", Lmargin1("12m"), Lmargin2("6m"), Rmargin("10m"))
	t.TagConfigure("overstrike", Overstrike(1))
	t.TagConfigure("raised", Relief("raised"), Borderwidth(1))
	t.TagConfigure("right", Justify("right"))
	t.TagConfigure("spacing", Spacing1("10p"), Spacing2("2p"), Lmargin1("12m"), Lmargin2("6m"), Rmargin("10m"))
	t.TagConfigure("sub", Offset("-2p"), Fnt(family+" 8"))
	t.TagConfigure("sunken", Relief("sunken"), Borderwidth(1))
	t.TagConfigure("super", Offset("4p"), Fnt(family+" 8"))
	t.TagConfigure("tiny", Fnt("Times 8 bold"))
	t.TagConfigure("underline", Underline(1))
	t.TagConfigure("veryBig", Fnt("Courier 22 bold"))

	t.Insert("end", s(`Text widgets like this one allow you to display information in a
variety of styles.  Display styles are controlled using a mechanism
called `))
	t.Insert("end", "tags", "bold")
	t.Insert("end", s(`.  Tags are just textual names that you can apply to one
or more ranges of characters within a text widget.  You can configure
tags with various display styles.  If you do this, then the tagged
characters will be displayed with the styles you chose.  The
available display styles are:
`))

	t.Insert("end", "\n\n1.  Font.", "big")
	t.Insert("end", "  You can choose any system font, ")
	t.Insert("end", "large", "veryBig")
	t.Insert("end", " or ")
	t.Insert("end", "small", "tiny", ".\n")

	t.Insert("end", "\n2.  Color.", "big")
	t.Insert("end", "  You can change either the ")
	t.Insert("end", "background", "color1")
	t.Insert("end", " or ")
	t.Insert("end", "foreground", "color2")
	t.Insert("end", " color, or ")
	t.Insert("end", "both", "color1 color2")
	t.Insert("end", ".\n")

	t.Insert("end", "\n3.  Stippling.", "big")
	t.Insert("end", "  You can cause the ")
	t.Insert("end", "background", "bgstipple")
	t.Insert("end", " information to be drawn with a stipple fill instead of a solid fill.\n")

	t.Insert("end", "\n4.  Underlining.", "big")
	t.Insert("end", "  You can ")
	t.Insert("end", "underline", "underline")
	t.Insert("end", " ranges of text.\n")

	t.Insert("end", "\n5.  Overstrikes.", "big")
	t.Insert("end", "  You can ")
	t.Insert("end", "draw lines through", "overstrike")
	t.Insert("end", " ranges of text.\n")

	t.Insert("end", "\n6.  3-D effects.", "big")
	t.Insert("end", "  You can arrange for the background to be drawn with a border that makes characters appear either ")
	t.Insert("end", "raised", "raised")
	t.Insert("end", " or ")
	t.Insert("end", "sunken", "sunken")
	t.Insert("end", ".\n")

	t.Insert("end", "\n7.  Justification.", "big")
	t.Insert("end", " You can arrange for lines to be displayed\n")
	t.Insert("end", "left-justified,\n")
	t.Insert("end", "right-justified, or\n", "right")
	t.Insert("end", "centered.\n", "center")

	t.Insert("end", "\n8.  Superscripts and subscripts.", "big")
	t.Insert("end", " You can control the vertical position of text to generate superscript effects like 10")
	t.Insert("end", "n", "super")
	t.Insert("end", " or subscript effects like X")
	t.Insert("end", "i", "sub")
	t.Insert("end", ".\n")

	t.Insert("end", "\n9.  Margins.", "big")
	t.Insert("end", " You can control the amount of extra space left on each side of the text:\n\n")
	t.Insert("end", s(`This paragraph is an example of the use of margins.  It
consists of a single line of text that wraps around on the screen.  There are two
separate left margin values, one for the first display line associated with the text line,
and one for the subsequent display lines, which occur because of wrapping.  There is also a
separate specification for the right margin, which is used to choose wrap points for lines.`), "margins")

	t.Insert("end", "\n\n10.  Spacing.", "big")
	t.Insert("end", s(` You can control the spacing of lines with three separate parameters.  "Spacing1"
tells how much extra space to leave above a line, "spacing3" tells how much space to leave below a line, and
if a text line wraps, "spacing2" tells how much space to leave between the display lines that
make up the text line.`)+"\n")
	t.Insert("end", s(`These indented paragraphs illustrate how spacing can be
used.  Each paragraph is actually a single line in the text widget, which is word-wrapped by the widget.`)+"\n", "spacing")
	t.Insert("end", s(`Spacing1 is set to 10 points for this text, which results
in relatively large gaps between the paragraphs.  Spacing2 is set to 2 points,
which results in just a bit of extra space within a pararaph.  Spacing3 isn't used
in this example.`)+"\n", "spacing")
	t.Insert("end", s(`To see where the space is, select ranges of
text within these paragraphs.  The selection
highlight will cover the extra space.`), "spacing")

	App.Configure(Padx("4m"), Pady("3m")).Center().Wait()
}

func s(s string) string {
	return strings.ReplaceAll(s, "\n", " ")
}
