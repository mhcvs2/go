package main

import "C"

//export add
func add(left, right *C.char) *C.char {
	// bytes对应ctypes的c_char_p类型,翻译成C类型就是 char *指针
	merge := C.GoString(left) + C.GoString(right)
	return C.CString(merge)
}

func main() {}

//go build -buildmode=c-shared -o libadd.so libadd.go