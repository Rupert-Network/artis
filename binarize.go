package main

import (
	"image"
	"image/color"
)

// Binarize applies a linear threshold to a grayscale image
func Binarize(img *image.Gray, threshold uint8) *image.Gray {
	result := image.NewGray(img.Bounds())

	for x := 0; x < img.Bounds().Max.X; x++ {
		for y := 0; y < img.Bounds().Max.Y; y++ {
			var c uint8
			c = 0
			if img.GrayAt(x, y).Y > threshold {
				c = 1
			}
			result.SetGray(x, y, color.Gray{Y: c})
		}
	}

	return result
}
