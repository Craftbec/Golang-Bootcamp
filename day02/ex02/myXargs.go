package main

import (
	"bufio"
	"os"
	"os/exec"
)

func main() {
	inp := bufio.NewScanner(os.Stdin)
	str := []string{}
	for inp.Scan() {
		str = append(str, inp.Text())
	}
	for i := range str {
		tmp := os.Args[2:]
		tmp = append(tmp, str[i])
		out := exec.Command(os.Args[1], tmp...)
		out.Stdout = os.Stdout
		out.Stderr = os.Stderr
		_ = out.Run()
	}
}
