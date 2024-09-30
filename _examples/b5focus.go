package main

import (
	. "modernc.org/tk9.0"
	"modernc.org/tk9.0/b5"
)

func main() {
	StyleThemeUse("default")
	opts := Opts{Padx("1m"), Pady("2m"), Ipadx("1m"), Ipady("1m")}
	Grid(TButton(Txt("Primary"), Style(b5.FocusedButtonStyle("focused.primary.TButton", b5.ButtonColors{b5.ButtonText: "#fff", b5.ButtonFace: "#0d6efd", b5.ButtonFocus: "#98c1fe"}, "#fff"))),
		TButton(Txt("Secondary"), Style(b5.FocusedButtonStyle("secondary.TButton", b5.ButtonColors{b5.ButtonText: "#fff", b5.ButtonFace: "#6c757d", b5.ButtonFocus: "#c0c4c8"}, "#fff"))),
		TButton(Txt("Success"), Style(b5.FocusedButtonStyle("sucess.TButton", b5.ButtonColors{b5.ButtonText: "#fff", b5.ButtonFace: "#198754", b5.ButtonFocus: "#9dccb6"}, "#fff"))),
		opts)
	Grid(TButton(Txt("Danger"), Style(b5.FocusedButtonStyle("danger.TButton", b5.ButtonColors{b5.ButtonText: "#fff", b5.ButtonFace: "#dc3545", b5.ButtonFocus: "#f0a9b0"}, "#fff"))),
		TButton(Txt("Warning"), Style(b5.FocusedButtonStyle("warning.TButton", b5.ButtonColors{b5.ButtonText: "#000", b5.ButtonFace: "#ffc107", b5.ButtonFocus: "#ecd182"}, "#fff"))),
		TButton(Txt("Info"), Style(b5.FocusedButtonStyle("info.TButton", b5.ButtonColors{b5.ButtonText: "#000", b5.ButtonFace: "#0dcaf0", b5.ButtonFocus: "#85d5e5"}, "#fff"))),
		opts)
	Grid(TButton(Txt("Light"), Style(b5.FocusedButtonStyle("light.TButton", b5.ButtonColors{b5.ButtonText: "#000", b5.ButtonFace: "#f8f9fa", b5.ButtonFocus: "#e9e9ea"}, "#fff"))),
		TButton(Txt("Dark"), Style(b5.FocusedButtonStyle("dark.TButton", b5.ButtonColors{b5.ButtonText: "#fff", b5.ButtonFace: "#212529", b5.ButtonFocus: "#a0a2a4"}, "#fff"))),
		TButton(Txt("Link"), Style(b5.FocusedButtonStyle("link.TButton", b5.ButtonColors{b5.ButtonText: "#1774fd", b5.ButtonFace: "#fff", b5.ButtonFocus: "#c2dbfe"}, "#fff"))),
		opts)
	Grid(TExit(), Columnspan(3), opts)
	App.Configure(Background("#fff")).Wait()
}
