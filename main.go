package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	// if err != nil {
	// 	log.Printf("Failed to open %s", err)
	// 	panic(err.Error())
	// }

	// defer in.Close()

	//decode jpeg Image
	imgContent, imgFormat, err := image.Decode(in)
	if err != nil {
		panic(err.Error())
	}

	//creating a blank Graysclae image
	imgbounds := imgContent.Bounds()
	width, height := imgbounds.Max.X, imgbounds.Max.Y

	grayScale := image.NewGray(image.Rectangle{image.Point{0, 0}, image.Point{width, height}})
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			imageColor := imgContent.At(x, y)
			rr, gg, bb, _ := imageColor.RGBA()
			r := math.Pow(float64(rr), 2.2)
			g := math.Pow(float64(gg), 2.2)
			b := math.Pow(float64(bb), 2.2)

			m := math.Pow(0.2126*r+0.7152*g+0.0722*b, 1/2.2)
			Y := uint16(m + 0.5)
			grayColor := color.Gray{uint8(Y >> 8)}

			//color conversion function
			grayScale.Set(x, y, grayColor)
		}
	}

	// newImage := "kinya2.jpg"
	// newfile, err := os.Create(newImage)
	// if err != nil {
	// 	log.Printf("failed creating %s:", err)
	// 	panic(err.Error())
	// }
	// defer newfile.Close()

	// outfile, err := os.Create(os.Args[2])
	// if err != nil {
	// 	log.Printf("failed creating %s:", err)
	// 	panic(err.Error())
	// }
	// defer outfile.Close()

	if imgFormat == "jpeg" {
		jpeg.Encode(os.Stdout, grayScale, nil)
	} else if imgFormat == "png" {
		png.Encode(os.Stdout, grayScale)
	} else {
		fmt.Println("File formart must be png or jpg")
	}

}
