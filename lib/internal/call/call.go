package call

/*
	#cgo LDFLAGS: -ldl -lavcall

	#include "call.h"
*/
import "C"

import (
	"errors"
)

func Open() error {
	cString := C.open()
	if cString != nil {
		return errors.New(C.GoString(cString))
	}

	return nil
}

func Function(index int) {
	C.call_function(C.int(index))
}

func Close() {
	C.close()
}
