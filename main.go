package main

import (
	"bufio"
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"

	"./filter"
	"./source"
)

func main() {
	infile := flag.String("infile", "", "please supply a url")
	flag.Parse()
	var in io.Reader
	if infile != nil && *infile != "" {
		x, downloadErr := source.DownloadURL(*infile)
		if downloadErr != nil {
			panic("download blew up: " + downloadErr.Error())
		}
		defer x.Close()
		in = x
	} else {
		in = bufio.NewReader(os.Stdin)
	}

	//decode  Image
	imgContent, imgFormat, err := image.Decode(in)
	if err != nil {
		panic(err.Error())
	}

	grayScale := filter.Mosaic(imgContent)

	if imgFormat == "jpeg" {
		jpeg.Encode(os.Stdout, grayScale, nil)
	} else if imgFormat == "png" {
		png.Encode(os.Stdout, grayScale)
	} else {
		fmt.Println("File formart must be png or jpg")
	}

}
