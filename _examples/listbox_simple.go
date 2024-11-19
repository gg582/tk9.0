package main

import . "modernc.org/tk9.0"

func main() {
    length := 3
    l := Listbox()
    l.AddItem(0, "item1 item2 item3")
    b1 := TButton(Txt("Delete Multiple Items, index (0-1)"), Command( func(){
        if length >= 2 {
            l.DeleteItems(0,1)
            length-=2
        }
    }))
    b2 := TButton(Txt("Delete One Item, index (0)"), Command( func () {
        if length > 0 {
            l.DeleteOne(0)
            length-=1
        }
    }))
    
	Pack(TExit(),l,b1,b2)
	App.Wait()
}
