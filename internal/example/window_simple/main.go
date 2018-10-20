package main

import (
	"fmt"
	"github.com/pekim/gobbi/lib/gdk"
	"github.com/pekim/gobbi/lib/gtk"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	gtk.Init([]string{})

	window := gtk.WindowNew(gtk.GTK_WINDOW_TOPLEVEL)
	window.SetTitle("A window title")
	window.SetDefaultSize(300, 300)

	container := gtk.BoxNew(gtk.GTK_ORIENTATION_VERTICAL, 10).Container()
	container.Add(gtk.LabelNew("Label 1").Widget())
	container.Add(gtk.LabelNew("Label 2 ++").Widget())
	window.Container().Add(container.Widget())

	window.Widget().ConnectDestroy(gtk.MainQuit)

	window.Widget().ConnectKeyPressEvent(func(event *gdk.EventKey) bool {
		fmt.Println(event.String, event.Keyval, event.State)
		return false
	})

	connectId2 := window.Widget().ConnectKeyPressEvent(func(event *gdk.EventKey) bool {
		fmt.Println("Should not see this message.")
		return false
	})
	window.Widget().DisconnectKeyPressEvent(connectId2)

	window.Widget().ShowAll()
	gtk.Main()
}