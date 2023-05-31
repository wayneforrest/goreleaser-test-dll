// go:build (cgo)

package main

import "C"

func main() {
}

//export Wrapper
func Wrapper(cmd *C.char) *C.char {
	// Convert the C string to a Go string
	goCmd := C.GoString(cmd)
	// Call the wrapped function
	goResult := wrapped(goCmd)
	// Convert the Go string result to a C string
	cResult := C.CString(goResult)
	return cResult
}

func wrapped(cmd string) string {
	return "hello world : " + cmd
}
