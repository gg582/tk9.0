package main

import . "github.com/gg582/tk9.0"

func main() {
	Pack(Button(Txt("Hello"), Command(func() { Destroy(App) })))
	App.Wait()
}
