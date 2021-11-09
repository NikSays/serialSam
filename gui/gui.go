package gui

import (
	_ "embed"
	"log"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

//go:embed gui.glade
var glade string;

type ErrorHandlingBuilder struct {*gtk.Builder}
func (b *ErrorHandlingBuilder) MustGetObject(id string) glib.IObject {
	object, err := b.GetObject(id)
	if err != nil {
		log.Fatalf("Couldn't get element %s: %s", id, err)
	}
	return object
}

func OpenWindow() {
    gtk.Init(nil)

    nativeBuilder, _ := gtk.BuilderNew()
		builder := ErrorHandlingBuilder{nativeBuilder}

    err := builder.AddFromString(glade)
    if err != nil {
        log.Fatal("Couldn't open gui.glade: ", err)
    }
		

    win := builder.MustGetObject("main").(*gtk.Window)
    win.Connect("destroy", func() {
        gtk.MainQuit()
    })

    win.ShowAll()

    gtk.Main()
}