package main

/*
#include "str.h"
*/
import "C"

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	str_res := C.make_string()
	stringHeader := reflect.StringHeader{
		Data: uintptr(unsafe.Pointer(str_res.str)),
		Len:  int(str_res.len),
	}
	s := (*string)(unsafe.Pointer(&stringHeader))
	fmt.Println(*s)
}
