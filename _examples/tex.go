package main

import . "modernc.org/tk9.0"

const tex = `$$\sin x = {{e^{ix}-e^{-ix}}\over 2i}$$`

func main() {
	Pack(
		Label(Relief("sunken"), Image(NewPhoto(Data(TeX(tex, 2*72/TkScaling()/600))))),
		TExit(), Padx("1m"), Pady("2m"), Ipadx("1m"), Ipady("1m"),
	)
	App.Configure(Padx("4m"), Pady("3m")).Center().Wait()
}
