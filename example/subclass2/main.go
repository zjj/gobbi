package main

import (
	"fmt"
	"github.com/pekim/gobbi/lib/gtk"
)

type MyWidget struct {
	instance *gtk.WidgetDerived
}

func (w *MyWidget) Init() {
	fmt.Println("my widget init")
}

func main() {
	gtk.Init([]string{})

	//sc.SubclassCreate()
	//gobject.SubclassCreate()
	//gobject.FfiClosure()

	class := gtk.WidgetDerive("test_widget")
	fmt.Println(class)

	myWidget := MyWidget{}
	myWidget.instance = class.New(&myWidget)

	fmt.Println(myWidget)
	//fmt.Println(myw.Widget().IsVisible())
	//myw.Widget().SetVisible(true)
	//fmt.Println(myw.Widget().IsVisible())
}
