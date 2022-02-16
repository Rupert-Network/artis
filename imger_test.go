package main

import (
	"github.com/bradleyjkemp/cupaloy"
	"github.com/stretchr/testify/assert"
	"github.com/Ernyoke/Imger/edgedetection"
	"github.com/Ernyoke/Imger/imgio"
	"testing"
)

func TestCannyEdge(t *testing.T) {
	img, err := imgio.ImreadRGBA("images/iris01.jpeg")
	assert.Nil(t, err)

	edgeImg, err := edgedetection.CannyRGBA(img, 15, 45, 5)
	assert.Nil(t, err)

	err = imgio.Imwrite(edgeImg, "images/irisedge01.jpeg")
	assert.Nil(t, err)

	cupaloy.SnapshotT(t, edgeImg)
}
