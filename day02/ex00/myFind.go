package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path"
	"path/filepath"
)

func read(dir string, f, d, sl bool, ext string) {
	dh, _ := os.Open(dir)
	defer dh.Close()
	for {
		fis, err := dh.Readdir(10)
		if err != nil {
			break
		}
		for _, fi := range fis {
			if f {
				if fi.Mode().IsRegular() {
					if len(ext) != 0 {
						if path.Ext(fi.Name()) == ext {
							fmt.Printf("%s/%s\n", dir, fi.Name())
						}
					} else {
						fmt.Printf("%s/%s\n", dir, fi.Name())
					}
				}
			}
			if sl {
				if fi.Mode()&fs.ModeSymlink != 0 {
					x, err := filepath.EvalSymlinks(fi.Name())
					if err != nil {
						fmt.Printf("%s/%s -> [broken]\n", dir, fi.Name())
					} else {
						fmt.Printf("%s/%s -> %s\n", dir, fi.Name(), x)
					}
				}
			}
			if fi.IsDir() {
				if d {
					fmt.Printf("%s/%s\n", dir, fi.Name())
				}
				read(dir+"/"+fi.Name(), f, d, sl, ext)
			}
		}
	}
}

func main() {
	var SL, D, F, EXT bool
	flag.BoolVar(&SL, "sl", false, "Print only symlinks")
	flag.BoolVar(&D, "d", false, "Print only directories")
	flag.BoolVar(&F, "f", false, "Print only files")
	flag.BoolVar(&EXT, "ext", false, "Print only files with a certain extension")
	flag.Parse()
	var ext string
	if EXT {
		if len(os.Args) >= 5 && F {
			ext = os.Args[len(os.Args)-2]
			ext = "." + ext
		} else {
			fmt.Println("Works only when -f is specified")
			os.Exit(3)
		}
	}
	_, err := os.ReadDir(os.Args[len(os.Args)-1])
	if err != nil {
		log.Fatal(err)
	}
	if !SL && !D && !F {
		read(os.Args[len(os.Args)-1], true, true, true, ext)
	} else {
		read(os.Args[len(os.Args)-1], F, D, SL, ext)
	}
}
