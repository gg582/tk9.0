package main

import . "modernc.org/tk9.0"

func main() {
    var b *ButtonWidget
    buttonText := Txt("Button 1")
    b = Button( buttonText, Command(func() {
        buttonText = Txt("Changed")
        Destroy(b.Window)
        b = Button(buttonText, Command( func() {
            Destroy(App)
        }))
        Pack(b)
    }))
    l := Listbox()
    l.AddItem(0, "white red blue")
	Pack(b,TExit(),l)
	App.Wait()
}
