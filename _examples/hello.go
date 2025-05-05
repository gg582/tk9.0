package main

import . "github.com/yoonjin67/tk9.0"

func main() {
	Pack(Button(Txt("Hello"), Command(func() { Destroy(App) })))
	App.Wait()
}
