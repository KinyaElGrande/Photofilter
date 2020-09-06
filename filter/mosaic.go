package filter

import (
	"image"
	"image/color"
)

//Mosaic function implements mosaic algorithm in an image
func Mosaic(im image.Image) image.Image {
	bounds := im.Bounds()

	newImage := image.NewRGBA(image.Rect(bounds.Min.X, bounds.Min.Y, bounds.Max.X, bounds.Max.Y))
	// break into tiles
	tiles := toTiles(im, 20, 20)
	//get the average color for each tile
	for _, tile := range tiles {
		c := averageColor(im, tile)
		blendColor(im, newImage, c, tile)
	}
	//blend each tile with the average

	return newImage
}

func averageColor(im image.Image, bounds image.Rectangle) color.Color {
	width, height := bounds.Max.X-bounds.Min.X+1, bounds.Max.Y-bounds.Min.Y+1
	var r, g, b, a uint32
	for y := bounds.Min.Y; y <= bounds.Max.Y; y++ {
		for x := bounds.Min.X; x <= bounds.Max.X; x++ {
			r1, g1, b1, a1 := im.At(x, y).RGBA()
			r, g, b, a = r+(r1), g+(g1), b+(b1), a+(a1)
		}
	}
	totalPixels := uint32(width * height)
	// return [3]float64{r / totalPixels, g / totalPixels, b / totalPixels}
	return color.RGBA{uint8(r / totalPixels), uint8(g / totalPixels), uint8(b / totalPixels), uint8(a / totalPixels)}
}

func toTiles(in image.Image, tx, ty int) []image.Rectangle {
	minX := in.Bounds().Min.X
	maxX := in.Bounds().Max.X
	minY := in.Bounds().Min.Y
	maxY := in.Bounds().Max.Y
	width := maxX - minX + 1
	height := maxY - minY + 1
	dx := width / tx
	dy := height / ty

	tiles := []image.Rectangle{}
	for x := minX; x < maxX; x += dx {
		for y := minY; y < maxY; y += dy {
			tiles = append(tiles, image.Rect(x, y, x+dx, y+dy))
		}
	}

	return tiles
}

func blendColor(in image.Image, out *image.RGBA, c color.Color, bounds image.Rectangle) {
	// rc, gc, bc, ac := c.RGBA()

	var blender uint32 = 65535
	var horChangeIncrement uint32 = 65535 / uint32(in.Bounds().Max.X-in.Bounds().Min.X)

	for x := bounds.Min.X; x <= bounds.Max.X; x++ {

		// decrease our red-ness each time we move across horizontally
		blender -= horChangeIncrement
		if blender < 0 {
			blender = 0
		}

		for y := bounds.Min.Y; y <= bounds.Max.Y; y++ {

			r1, g1, b1, _ := in.At(x, y).RGBA()

			// calculate a new values that averages the current RGB
			newRed := (r1 + blender) / 2
			newGreen := (g1 + blender) / 2
			newBlue := (b1 + blender) / 2

			newColor := color.RGBA{R: uint8(newRed / 0x101), G: uint8(newGreen / 0x101), B: uint8(newBlue / 0x101), A: 255}
			out.Set(x, y, newColor)
		}
	}
}
