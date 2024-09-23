package main

import (
	. "modernc.org/tk9.0"
	"modernc.org/tk9.0/b5"
)

func main() {
	StyleThemeUse("default")
	opts := Opts{Padx("1m"), Pady("2m"), Ipadx("1m"), Ipady("1m")}
	Grid(TButton(Txt("Primary"), Style(b5.ButtonStyle("primary.TButton", b5.ButtonColors{b5.ButtonText: "#fff", b5.ButtonFace: "#0d6efd"}, "#fff"))),
		TButton(Txt("Secondary"), Style(b5.ButtonStyle("secondary.TButton", b5.ButtonColors{b5.ButtonText: "#fff", b5.ButtonFace: "#6c757d"}, "#fff"))),
		TButton(Txt("Success"), Style(b5.ButtonStyle("sucess.TButton", b5.ButtonColors{b5.ButtonText: "#fff", b5.ButtonFace: "#198754"}, "#fff"))),
		opts)
	Grid(TButton(Txt("Danger"), Style(b5.ButtonStyle("sanger.TButton", b5.ButtonColors{b5.ButtonText: "#fff", b5.ButtonFace: "#dc3545"}, "#fff"))),
		TButton(Txt("Warning"), Style(b5.ButtonStyle("warning.TButton", b5.ButtonColors{b5.ButtonText: "#000", b5.ButtonFace: "#ffc107"}, "#fff"))),
		TButton(Txt("Info"), Style(b5.ButtonStyle("info.TButton", b5.ButtonColors{b5.ButtonText: "#000", b5.ButtonFace: "#0dcaf0"}, "#fff"))),
		opts)
	Grid(TButton(Txt("Light"), Style(b5.ButtonStyle("light.TButton", b5.ButtonColors{b5.ButtonText: "#000", b5.ButtonFace: "#f8f9fa"}, "#fff"))),
		TButton(Txt("Dark"), Style(b5.ButtonStyle("dark.TButton", b5.ButtonColors{b5.ButtonText: "#fff", b5.ButtonFace: "#212529"}, "#fff"))),
		TButton(Txt("Link"), Style(b5.ButtonStyle("link.TButton", b5.ButtonColors{b5.ButtonText: "#1774fd", b5.ButtonFace: "#fff"}, "#fff"))),
		opts)
	Grid(TExit(), Columnspan(3), opts)
	App.Configure(Background("#fff")).Wait()
}
