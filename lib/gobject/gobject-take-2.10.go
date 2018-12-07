// +build gobject_2.10 gobject_2.12 gobject_2.18 gobject_2.24 gobject_2.26 gobject_2.28 gobject_2.30 gobject_2.32 gobject_2.34 gobject_2.36 gobject_2.38 gobject_2.42 gobject_2.44 gobject_2.46 gobject_2.54

package gobject

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <stdlib.h>
import "C"
import (
	"fmt"
	"runtime"
)

func (recv *Object) Take() {
	//if recv.IsFloating() {
	//	fmt.Printf("%p RefSink \n", recv)
	//	recv.RefSink()
	//	runtime.SetFinalizer(recv, objectFinalizer)
	//}

	//if !recv.IsFloating() {
	//	fmt.Printf("%p Ref \n", recv)
	//	recv.Ref()
	//	runtime.SetFinalizer(recv, objectFinalizer)
	//}

	if recv.IsFloating() {
		fmt.Printf("%p RefSink\n", recv)
		recv.RefSink()
	} else {
		fmt.Printf("%p Ref\n", recv)
		recv.Ref()
	}

	runtime.SetFinalizer(recv, objectFinalizer)
}

func objectFinalizer(o *Object) {
	fmt.Printf("%p Unref (in finalizer)\n", o)
	o.Unref()
}
