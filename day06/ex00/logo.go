package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main() {
	upLeft := image.Point{0, 0}
	lowRight := image.Point{300, 300}
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	lynch := color.RGBA{108, 122, 137, 0xff}
	snuff := color.RGBA{227, 218, 231, 0xff}
	for x := 0; x < 300; x++ {
		for y := 0; y < 300; y++ {
			switch {
			case x < 300/2 && y < 300/2:
				img.Set(x, y, lynch)
			case x >= 300/2 && y >= 300/2:
				img.Set(x, y, lynch)
			case x < 300/2 && y >= 300/2:
				img.Set(x, y, snuff)
			case x >= 300/2 && y < 300/2:
				img.Set(x, y, snuff)

			default:
			}
		}
	}
	f, err := os.Create("amazing_logo.png")
	if err != nil {
		log.Fatal(err)
	}
	err = png.Encode(f, img)
	if err != nil {
		log.Fatal(err)
	}
}
