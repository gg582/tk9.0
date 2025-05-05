
package main

import (
    "fmt"
    "log"
    "os"
    "runtime"
    "strings"
    "path"

    . "github.com/yoonjin67/tk9.0"
)

var pbuttons []*TButtonWidget
var extensions []FileType
var pbutton *TButtonWidget = nil
var listbox, listbox2 *ListboxWidget
var cur *LabelWidget = nil
var imagesLoaded []*LabelWidget
func PhotoName(fileName string) string {
	fileName = path.Base(fileName)
	return fileName[:len(fileName)-len(path.Ext(fileName))]
}

func handleFileOpen() {
    res := GetOpenFile(Multiple(true),Filetypes(extensions)) //다중 선택을 활성화하고 필터를 켭니다.
    s := make([]string,0,1000)
    for _, itm := range res {
        if itm != "" {
            tmp := strings.Split(itm," ")
            s = append(s,tmp...)
        }
    }

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
			imagesLoaded = append(imagesLoaded,imageLabel)
            var deleteTestButton *TButtonWidget
            deleteTestButton = TButton(
                Txt("Unshow Image"),
            Command(func() {
                GridForget(imageLabel.Window)
                GridForget(deleteTestButton.Window)
            }))

            pbuttons = append(pbuttons,deleteTestButton)

			listbox.AddItem(len(imagesLoaded)-1,PhotoName(photo))
			listbox2.AddItem(len(imagesLoaded)-1,PhotoName(photo))
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
        if len(imagesLoaded)-1>i {
            continue
        }
        if cur == imagesLoaded[i] {
            pbutton = nil
            cur = nil
        }
    	Destroy(imagesLoaded[i])
        Destroy(pbuttons[i])
		imagesLoaded = append(imagesLoaded[:i],imagesLoaded[i+1:]...)
        pbuttons = append(pbuttons[:i], pbuttons[i+1:]...)
    }
}

func SelectImage() {
	s:=listbox2.Selected()
	if len(s) == 0 {
		return
	}

    if len(imagesLoaded) -1 < s[0] {
        return
    }
    if imagesLoaded[s[0]] == nil {
        return 
    }
    if cur != nil {
	    GridForget(cur.Window)
    }
    if pbutton != nil {
        GridForget(pbutton.Window)
    }

	Grid(imagesLoaded[s[0]], Row(0), Column(2))
    Grid(pbuttons[s[0]], Row(0), Column(3))
    cur = imagesLoaded[s[0]]
    pbutton = pbuttons[s[0]]
}

func SelectIndex(index int) {

    if len(imagesLoaded) -1 <index {
        return
    }
    if imagesLoaded[index] == nil {
        return 
    }
    if cur != nil {
	    GridForget(cur.Window)
    }
    if pbutton != nil {
        GridForget(pbutton.Window)
    }

	Grid(imagesLoaded[index], Row(0), Column(2))
    Grid(pbuttons[index], Row(0), Column(3))
    cur = imagesLoaded[index]
    pbutton = pbuttons[index]
}

    




func main() {
    menubar := Menu()
    //DefaultTheme("awdark","themes/awthemes-10.4.0")
    //테마를 사용하고 싶을 때에는 테마 명과 경로를 지정해 줍니다.
    fileMenu := menubar.Menu()
    extensions = make([]FileType,0,1)
    extensions = append(extensions, FileType{ "Supported Images", []string {".png",".ico"}, "" } )
    //필터에 png와 ico를 넣어 줍니다.
    fileMenu.AddCommand(Lbl("Open..."), Underline(0), Accelerator("Ctrl+O"), Command(func () {
        handleFileOpen()
        SelectIndex(len(imagesLoaded)-1)
    } ))
    Bind(App, "<Control-o>", Command(func() { fileMenu.Invoke(0) }))
    fileMenu.AddCommand(Lbl("Exit"), Underline(1), Accelerator("Ctrl+Q"), ExitHandler())
    Bind(App, "<Control-q>", Command(func() { fileMenu.Invoke(1) }))
    menubar.AddCascade(Lbl("File"), Underline(0), Mnu(fileMenu))
	imagesLoaded = make([]*LabelWidget, 0, 10000)
    pbuttons = make([]*TButtonWidget,0,10000)
    var scrollx, scroll, scroll2, scrollx2 *TScrollbarWidget
	listbox = Listbox(Yscrollcommand(func(e *Event) { e.ScrollSet(scroll)}) , Xscrollcommand( func(e *Event) { e.ScrollSet(scrollx)}))
	listbox2 = Listbox(Yscrollcommand(func(e *Event) { e.ScrollSet(scroll2)}), Xscrollcommand(func(e *Event) { e.ScrollSet(scrollx2)}))
	listbox.SelectMode("multiple")
	listbox2 = Listbox()
	listbox.Background("white")
	listbox.SelectBackground("blue")
	listbox.SelectForeground("yellow")
	listbox2.Background("grey")
	listbox2.SelectBackground("green")
    listbox2.SelectForeground("blue")
    listbox2.SelectBackground("brown")
	listbox.Height(5)
	listbox.Width(4)
	listbox2.Height(5)
	listbox2.Width(4)
	delBtn := Button(Txt("Delete"), Command(func () { DeleteSelected() }))
	selBtn := Button(Txt("Select"), Command(func () { SelectImage() }))
	scroll = TScrollbar(Command(func(e *Event) { e.Yview(listbox) }))
	scrollx = TScrollbar(Orient("horizontal"),Command(func(e *Event) { e.Xview(listbox) }))
    scroll2 = TScrollbar(Command(func(e *Event) { e.Yview(listbox2) }))
	scrollx2 = TScrollbar(Orient("horizontal"),Command(func(e *Event) { e.Xview(listbox2) }))
	Grid(listbox,Row(1),Column(0), Sticky("nes"))
	Grid(scroll,Row(1),Column(1), Sticky("nes"))
    Grid(scrollx,Row(2),Column(0),  Sticky("nes"))
	Grid(delBtn,Row(3),Column(0), Sticky("nes"))
	Grid(listbox2,Row(1),Column(2), Sticky("nes"))
	Grid(scroll2,Row(1),Column(3), Sticky("nes"))
    Grid(scrollx2,Row(2),Column(2), Sticky("nes"))
	Grid(selBtn,Row(3),Column(2), Sticky("nes"))
    App.WmTitle(fmt.Sprintf("%s on %s", App.WmTitle(""), runtime.GOOS))
    App.Configure(Mnu(menubar), Width("80c"), Height("60c")).Wait()
}

