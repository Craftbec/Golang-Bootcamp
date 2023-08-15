package main

import (
	"errors"
	"fmt"
	"log"
	"unsafe"
)

func getElement(arr []int, idx int) (int, error) {
	if len(arr) == 0 {
		return 0, errors.New("empty array")
	} else if idx >= len(arr) || idx < 0 {
		return 0, errors.New("index outside")
	}
	res := *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&arr[0])) + (uintptr(idx) * unsafe.Sizeof(int(0)))))
	return res, nil
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	idx := 4
	res, err := getElement(arr, idx)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res)
}
