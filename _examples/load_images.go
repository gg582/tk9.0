
package main

import (
    "fmt"
    "log"
    "os"
    "runtime"
    "strings"
    "path"

    . "modernc.org/tk9.0"
)

var listbox, listbox2 *ListboxWidget
var imagesLen int = 0
var cur *LabelWidget
var imagesLoaded []*LabelWidget
func PhotoName(fileName string) string {
	fileName = path.Base(fileName)
	return fileName[:len(fileName)-len(path.Ext(fileName))]
}

func handleFileOpen() {
    s := GetOpenFile()
    for _, photo := range s {
        formatCheck := strings.Split(photo, ".")
        format := formatCheck[len(formatCheck)-1]
        
        if (strings.Compare(format, "png") == 0) || (strings.Compare(format, "ico") == 0) {
            picFile, err := os.Open(photo)
            if err != nil {
                log.Println("Error while opening photo, error is: ", err)
            }

            pictureRawData := make([]byte, 10000*10000)
            picFile.Read(pictureRawData)

            imageLabel := Label(Image(NewPhoto(Data(pictureRawData))))
            cur = imageLabel
			imagesLen+=1
			imagesLoaded = make([]*LabelWidget, imagesLen, imagesLen)
			imagesLoaded = append(imagesLoaded,imageLabel)
            var deleteTestButton *TButtonWidget
            deleteTestButton = TButton(
                Txt("Unshow Image"),
            Command(func() {
                Forget(imageLabel.Window)
                Forget(deleteTestButton.Window)
            }))
            Pack(imageLabel,
                deleteTestButton, //to test delete function 
                Padx("1m"), Pady("2m"), Ipadx("1m"), Ipady("1m"),
                Side("left"))


			listbox.AddItem(imagesLen-1,PhotoName(photo))
			listbox2.AddItem(imagesLen-1,PhotoName(photo))
            picFile.Close()
        }
    }
}

func DeleteSelected () {
    s:=listbox.Selected()
    if len(s) == 0 {
    	return
	}
    for _, i := range s {
    	listbox.DeleteOne(i)
    	listbox2.DeleteOne(i)
    	Destroy(imagesLoaded[i])
		imagesLoaded = append(imagesLoaded[:i],imagesLoaded[i+1:]...)
		imagesLen = len(imagesLoaded)
		imagesLoaded = make([](*LabelWidget), imagesLen,imagesLen)
    }
}

func SelectImage() {
	s:=listbox2.Selected()
	if len(s) == 0 {
		return
	}
	Forget(cur.Window)

	var deleteTestButton *TButtonWidget
	deleteTestButton = TButton(
        Txt("Unshow Image"),
    Command(func() {
        Forget(imagesLoaded[s[0]].Window)
        Forget(deleteTestButton.Window)
    }))
	Grid(imagesLoaded[s[0]], Row(0), Column(1),deleteTestButton, Padx("1m"),Pady("2m"), Ipadx("1m"), Ipady("1m"), Side("left"))
}




func main() {
    menubar := Menu()

    fileMenu := menubar.Menu()
    fileMenu.AddCommand(Lbl("Open..."), Underline(0), Accelerator("Ctrl+O"), Command(handleFileOpen))
    Bind(App, "<Control-o>", Command(func() { fileMenu.Invoke(0) }))
    fileMenu.AddCommand(Lbl("Exit"), Underline(1), Accelerator("Ctrl+Q"), ExitHandler())
    Bind(App, "<Control-q>", Command(func() { fileMenu.Invoke(1) }))
    menubar.AddCascade(Lbl("File"), Underline(0), Mnu(fileMenu))
	listbox = Listbox()
	listbox.SelectMode("multiple")
	listbox2 = Listbox()
	listbox.Background("grey")
	listbox.SelectBackground("blue")
	listbox.SelectForeground("yellow")
	listbox2.Background("white")
	listbox2.SelectBackground("green")
	listbox2.SelectForeground("grey")
	listbox.Height(5)
	listbox.Width(4)
	listbox2.Height(5)
	listbox2.Width(4)
	delBtn := Button(Txt("Delete Images"), Command(func () { DeleteSelected() }))
	selBtn := Button(Txt("Select Images"), Command(func () { SelectImage() }))
	scroll := TScrollbar(Command(func(e *Event) { e.Yview(listbox) }))
	scroll2 := TScrollbar(Command(func(e *Event) { e.Yview(listbox2) }))
	Grid(listbox,Row(0),Column(0))
	Grid(scroll,Row(0),Column(1))
	Grid(delBtn,Row(1),Column(0))
	Grid(listbox2,Row(2),Column(0))
	Grid(scroll2,Row(2),Column(1))
	Grid(selBtn,Row(3),Column(0))
    App.WmTitle(fmt.Sprintf("%s on %s", App.WmTitle(""), runtime.GOOS))
    App.Configure(Mnu(menubar), Width("10c"), Height("10c")).Wait()
}

