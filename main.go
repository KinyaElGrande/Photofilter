package main

import (
	"bufio"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"./filter"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	//decode  Image
	imgContent, imgFormat, err := image.Decode(in)
	if err != nil {
		panic(err.Error())
	}

	grayScale := filter.Gray(imgContent)

	if imgFormat == "jpeg" {
		jpeg.Encode(os.Stdout, grayScale, nil)
	} else if imgFormat == "png" {
		png.Encode(os.Stdout, grayScale)
	} else {
		fmt.Println("File formart must be png or jpg")
	}

}
