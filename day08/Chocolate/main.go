package main

/*
	#cgo CFLAGS: -x objective-c
	#cgo LDFLAGS: -framework Cocoa
	#include <window.h>
	#include <application.h>
*/
import "C"

func main() {
	C.InitApplication()
	C.Window_MakeKeyAndOrderFront(C.Window_Create(500, -200, 300, 200, C.CString("School 21")))
	C.RunApplication()
}
