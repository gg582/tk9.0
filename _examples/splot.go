package main

import . "modernc.org/tk9.0"

func main() {
	Pack(
		Label(Image(NewPhoto(Width(800), Height(600)).Graph("set grid; splot x**2+y**2, x**2-y**2"))),
		TExit(), Padx("1m"), Pady("2m"), Ipadx("1m"), Ipady("1m"),
	)
	App.Center().Wait()
}
