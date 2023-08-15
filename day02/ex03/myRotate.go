package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func name(s string) (string, error) {
	var res string
	filename := s
	ext := filepath.Ext(filename)
	res = strings.TrimSuffix(filename, ext)
	f, err := os.Open(s)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return res, err
	}
	defer f.Close()
	stat, _ := f.Stat()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return res, err
	}
	tim := stat.ModTime().Unix()
	res += "_" + strconv.Itoa(int(tim)) + ".tar.gz"
	return res, nil
}

func rate(s, path string, c chan string) {
	mas := []string{"-czf"}
	ar_name, err := name(s)
	if err != nil {
		return
	}
	mas = append(mas, ar_name)
	mas = append(mas, s)
	exec.Command("tar", mas...).Output()
	if len(path) > 0 {
		mas1 := []string{ar_name, path}
		exec.Command("mv", mas1...).Output()
	}
	c <- "OK"
}

func main() {
	var A bool
	flag.BoolVar(&A, "a", false, "File path")
	flag.Parse()
	i := 1
	var path string
	c := make(chan string)
	if A && len(os.Args) >= 4 {
		i = 3
		path = flag.Arg(0)
		A = false
	}
	if !A && len(os.Args) >= 2 {
		for ; i < len(os.Args); i++ {
			go rate(os.Args[i], path, c)
			<-c
		}
	} else {
		fmt.Println("Incorrect input")
		os.Exit(3)
	}
}
