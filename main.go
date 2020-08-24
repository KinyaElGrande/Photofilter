package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"math"
	"os"
)

func main() {
	imagefile := "test.jpg"
	in, err := os.Open(imagefile)

	if err != nil {
		log.Printf("Failed to open %s: %s", imagefile, err)
		panic(err.Error())
	}

	defer in.Close()

	imgContent, _, err := image.Decode(in)
	if err != nil {
		panic(err.Error())
	}

	//creating a blank Graysclae image
	bounds := imgContent.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	grayScale := image.NewGray(image.Rectangle{image.Point{0, 0}, image.Point{width, height}})
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			imageColor := imgContent.At(x, y)
			rr, gg, bb, _ := imageColor.RGBA()
			r := math.Pow(float64(rr), 2.2)
			g := math.Pow(float64(gg), 2.2)
			b := math.Pow(float64(bb), 2.2)

			m := math.Pow(0.2125*r+0.7154*g+0.0721*b, 1/2.2)
			Y := uint16(m + 0.5)
			grayColor := color.Gray{uint8(Y >> 8)}
			grayScale.Set(x, y, grayColor)
		}
	}

	newImage := "kinya.jpg"
	newfile, err := os.Create(newImage)
	if err != nil {
		log.Printf("failed creating %s:", err)
		panic(err.Error())
	}
	defer newfile.Close()
	jpeg.Encode(newfile, grayScale, nil)
}
