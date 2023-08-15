package main

import (
	"fmt"
	"time"
)

func sleepSort(arr []int) chan int {
	c := make(chan int, len(arr))
	for _, a := range arr {
		go func(a int) {
			time.Sleep(time.Duration(a) * time.Second)
			c <- a
		}(a)
	}
	return c
}

func main() {
	arr := []int{3, 5, 1, 4, 9, 2}
	c := sleepSort(arr)
	for i := 0; i < len(arr); i++ {
		fmt.Println(<-c)
	}
	close(c)
}
