package main

import (
	"github.com/bradleyjkemp/cupaloy"
	"github.com/stretchr/testify/assert"
	"github.com/Ernyoke/Imger/edgedetection"
	"github.com/Ernyoke/Imger/imgio"
	"testing"
)

func TestCannyEdge(t *testing.T) {
	// Read in image
	img, err := imgio.ImreadRGBA("images/iris01.jpeg")
	assert.Nil(t, err)

	// Create canny edges
	edgeImg, err := edgedetection.CannyRGBA(img, 15, 45, 5)
	assert.Nil(t, err)

	// Write image out to images for visual analysis
	if _, err := imgio.ImreadRGBA("images/irisedge01.jpeg"); err != nil {
		err = imgio.Imwrite(edgeImg, "images/irisedge01.jpeg")
		assert.Nil(t, err)
	}

	// Compare canny edges with snapshot
	cupaloy.SnapshotT(t, edgeImg)
}
