package main

import (
	"bufio"
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"os"

	"./filter"
)

func main() {
	infile := flag.String("infile", "", "please supply a url")
	flag.Parse()
	var in io.Reader
	if infile != nil && *infile != "" {
		x, downloadErr := DownloadURL(*infile)
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

	grayScale := filter.Gray(imgContent)

	if imgFormat == "jpeg" {
		jpeg.Encode(os.Stdout, grayScale, nil)
	} else if imgFormat == "png" {
		png.Encode(os.Stdout, grayScale)
	} else {
		fmt.Println("File formart must be png or jpg")
	}

}

//DownloadURL downloads from the specified url
func DownloadURL(url string) (io.ReadCloser, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}
