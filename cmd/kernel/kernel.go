package main

import (
	"unsafe"
)

func main() {
	print("Hello from Go Kernel!")
}

func print(s string) {
	for _, c := range s {
		*(*uint16)(unsafe.Pointer(uintptr(0xb8000) + uintptr(len(s)*2))) = uint16(c) | 0x0F00 
	}
}

