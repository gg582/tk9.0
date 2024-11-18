
package main

import (
    "fmt"
    "log"
    "os"
    "runtime"
    "strings"

    . "modernc.org/tk9.0"
)

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
            var deleteTestButton *TButtonWidget
            deleteTestButton = TButton(
                                    Txt("Delete Image"),
                                    Command(func() {
                                        Destroy(imageLabel.Window)
                                        Destroy(deleteTestButton.Window)

                                        }))
            Pack(imageLabel,
                deleteTestButton, //to test delete function 
                Padx("1m"), Pady("2m"), Ipadx("1m"), Ipady("1m"))

            picFile.Close()
        }
    }
}

func main() {
    menubar := Menu()

    fileMenu := menubar.Menu()
    fileMenu.AddCommand(Lbl("Open..."), Underline(0), Accelerator("Ctrl+O"), Command(handleFileOpen))
    Bind(App, "<Control-o>", Command(func() { fileMenu.Invoke(0) }))
    fileMenu.AddCommand(Lbl("Exit"), Underline(1), Accelerator("Ctrl+Q"), ExitHandler())
    Bind(App, "<Control-q>", Command(func() { fileMenu.Invoke(1) }))
    menubar.AddCascade(Lbl("File"), Underline(0), Mnu(fileMenu))

    App.WmTitle(fmt.Sprintf("%s on %s", App.WmTitle(""), runtime.GOOS))
    App.Configure(Mnu(menubar), Width("10c"), Height("10c")).Wait()
}

