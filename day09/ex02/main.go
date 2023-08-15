package main

import (
	"fmt"
	"sync"
)

func multiplex(inputs ...chan interface{}) chan interface{} {
	output := make(chan interface{})
	var wg sync.WaitGroup
	for _, in := range inputs {
		wg.Add(1)
		go func(in chan interface{}) {
			for inn := range in {
				output <- inn
			}
			wg.Done()
		}(in)
	}
	go func() {
		wg.Wait()
		close(output)
	}()
	return output
}

func main() {

	a := make(chan interface{})
	b := make(chan interface{})
	go func() {
		for i := 0; i < 3; i++ {
			a <- i
			b <- "TEST"
		}
		close(a)
		close(b)
	}()
	out := multiplex(a, b)
	for x := range out {
		fmt.Println(x)
	}

}
