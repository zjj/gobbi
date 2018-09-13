// +build glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Bytes is a wrapper around the C record GBytes.
type Bytes struct {
	native *C.GBytes
}

// Rwlock is a wrapper around the C record GRWLock.
type Rwlock struct {
	native *C.GRWLock
	P      int
	I      int
}

// Recmutex is a wrapper around the C record GRecMutex.
type Recmutex struct {
	native *C.GRecMutex
	P      int
	I      int
}
