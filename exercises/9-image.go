package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	width, height int
	white         uint8
}

func (i Image) Bounds() image.Rectangle {
	return image.Rectangle{
		Min: image.Point{0, 0},
		Max: image.Point{i.width, i.height},
	}
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{i.white, i.white, 255, 255}
}

func main() {
	m := Image{width: 100, height: 100, white: 255}
	pic.ShowImage(m)
}
