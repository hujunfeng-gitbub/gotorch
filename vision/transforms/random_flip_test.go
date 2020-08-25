package transforms

import (
	"image"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomFlip(t *testing.T) {
	a := assert.New(t)
	i := generateRandImage(image.Rect(0, 0, 50, 50))
	width := i.Bounds().Max.X

	trans := RandomFlip()
	_, err := trans.Run(100)
	a.Error(err)

	// Run 10 times to cover random cases
	for j := 0; j < 10; j++ {
		o, err := trans.Run(i)
		a.NoError(err)
		outImage := o.(*image.NRGBA)
		inImage := i.(*image.NRGBA)
		a.True(inImage.At(0, 0) == outImage.At(0, 0) || inImage.At(0, 0) == outImage.At(width-1, 0))
	}
}
