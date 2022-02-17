package main

import (
	"testing"

	"github.com/Ernyoke/Imger/imgio"
	"github.com/Ernyoke/Imger/grayscale"
	"github.com/bradleyjkemp/cupaloy"
	"github.com/stretchr/testify/assert"
)

func TestBinarize(t *testing.T) {
	// Read in image
	img, err := imgio.ImreadRGBA("images/iris01.jpeg")
	assert.Nil(t, err)

	// Convert to grayscale
	gImg := grayscale.Grayscale(img)

	// Binarize
	bImg := Binarize(gImg, 4)

	// Write image out to images for visual analysis
	if _, err := imgio.ImreadRGBA("images/irisedge01.jpeg"); err != nil {
		err = imgio.Imwrite(bImg, "images/irisbinarized01.jpeg")
		assert.Nil(t, err)
	}

	// Compare canny edges with snapshot
	cupaloy.SnapshotT(t, bImg)
}
