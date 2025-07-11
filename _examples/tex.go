package main

import . "github.com/gg582/tk9.0"

func main() {
	tex := `$$\int _0 ^\infty {{\sin ax \sin bx}\over{x^2}}\,dx = {\pi a\over 2}$$`
	Pack(Label(Relief("sunken"), Image(NewPhoto(Data(TeX(tex, 2*TkScaling()*72/600))))),
		TExit(),
		Padx("1m"), Pady("2m"), Ipadx("1m"), Ipady("1m"))
	App.Center().Wait()
}
