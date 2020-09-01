package filter

import (
	"image"
)

//Mosaic function implements mosaic algorithm in an image
func averageColor(im image.Image) [3]float64 {
	bounds := im.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	r, g, b := 0.0, 0.0, 0.0
	for y := bounds.Min.Y; y < height; y++ {
		for x := bounds.Min.X; x < width; x++ {
			r1, g1, b1, _ := im.At(x, y).RGBA()
			r, g, b = r+float64(r1), g+float64(g1), b+float64(b1)
		}
	}
	totalPixels := float64(width * height)
	return [3]float64{r / totalPixels, g / totalPixels, b / totalPixels}
}
