package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"unicode/utf8"
)

func count(str string, a, b bool, c chan string) {
	file, err := os.Open(str)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fileScanner := bufio.NewScanner(file)
	lineCount := 0
	for fileScanner.Scan() {
		if a {
			lineCount++
		} else if b {
			lineCount += utf8.RuneCountInString(fileScanner.Text())
		} else {
			if len(fileScanner.Text()) != 0 {
				lineCount++
			}
			for i := 0; i < len(fileScanner.Text()); i++ {
				if fileScanner.Text()[i] == 32 {
					lineCount++
				}
			}
		}
	}
	fmt.Printf("%d\t%s\n", lineCount, str)
	c <- "OK"
}

func main() {
	var L, M, W bool
	c := make(chan string)
	flag.BoolVar(&L, "l", false, "counting lines")
	flag.BoolVar(&M, "m", false, "counting characters")
	flag.BoolVar(&W, "w", false, "counting words")
	flag.Parse()
	if !W && !M && !L {
		W = true
	}
	if (L && !M && !W) || (M && !L && !W) || (W && !M && !L) {
		for i := 0; i < flag.NArg(); i++ {
			go count(flag.Arg(i), L, M, c)
		}
	} else {
		fmt.Println("Only one flag possible")
		os.Exit(3)
	}
	for i := 0; i < flag.NArg(); i++ {
		<-c
	}
}
