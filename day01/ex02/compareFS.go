package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
)

func check_format(s string) bool {
	if path.Ext(s) == ".txt" {
		return true
	}
	return false
}

func compare(s1, s2 string) bool {
	m := map[string]bool{}
	file, err := os.Open(s1)
	if err != nil {
		fmt.Printf("Error when opening file: %s\n", err)
		return false
	}
	fileScanner := bufio.NewScanner(file)
	if err = fileScanner.Err(); err != nil {
		fmt.Printf("Error while reading file: %s\n", err)
		return false
	}
	for fileScanner.Scan() {
		m[fileScanner.Text()] = true
	}
	file.Close()

	file, err = os.Open(s2)
	if err != nil {
		fmt.Printf("Error when opening file: %s\n", err)
		return false
	}
	fileScanner = bufio.NewScanner(file)
	if err = fileScanner.Err(); err != nil {
		fmt.Printf("Error while reading file: %s\n", err)
		return false
	}
	for fileScanner.Scan() {
		if m[fileScanner.Text()] {
			delete(m, fileScanner.Text())
		} else {
			fmt.Printf("ADDED %s\n", fileScanner.Text())
		}
	}
	file.Close()
	for i := range m {
		fmt.Printf("REMOVED %s\n", i)
	}
	return true
}

func main() {
	var res bool
	if len(os.Args) == 5 && os.Args[1] == "--old" && check_format(os.Args[2]) && os.Args[3] == "--new" && check_format(os.Args[4]) {
		res = compare(os.Args[2], os.Args[4])
		if res == false {
			os.Exit(4)
		}
	} else {
		fmt.Println("Incorrect application launch")
		os.Exit(3)
	}

}
