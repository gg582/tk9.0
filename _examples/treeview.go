package main

import . "modernc.org/tk9.0"

func main() {
    
	tr := TreeView()
	tr.SelectMode("extended")
	Set("itm1", "{1 2 3 4}")
	Set("itm2", "{5 6 7 8}")
	tr.Insert("", "lbl1","$itm1",0)
	tr.Insert("", "lbl2","$itm2",1)
	b := TButton(Txt("Delete Selected"), Command(func () {
		sel := tr.Selected()
		tr.Delete(sel)
	}))
	Pack(TExit(),tr,b)
	App.Wait()
}
