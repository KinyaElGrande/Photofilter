package filter

import (
	"image"
	"image/color"
	"math"
)

//Gray transforms an image to Grayscale
func Gray(im image.Image) image.Image {
	//creating a blank Graysclae image
	imgbounds := im.Bounds()
	width, height := imgbounds.Max.X, imgbounds.Max.Y

	grayScale := image.NewGray(image.Rectangle{image.Point{0, 0}, image.Point{width, height}})
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			imageColor := im.At(x, y)
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
	return grayScale
}
