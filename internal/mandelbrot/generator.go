package mandelbrot

import (
	"image"
	"math/cmplx"
	"sync"
)

type Generator struct {
	Width     int
	Height    int
	MaxEscape int
	Radius    *float64
	StartX    *float64
	StartY    *float64
}

func (gen *Generator) generateMandelbrotImage() image.Image {
	img := image.NewRGBA(image.Rect(0, 0, gen.Width, gen.Height))
	palette := Random.GetPalette(gen.MaxEscape)

	viewCenter := complex(*gen.StartX, *gen.StartY)
	zoomWidth := *gen.Radius * 2
	pixelWidth := zoomWidth / float64(gen.Width)
	pixelHeight := pixelWidth
	viewHeight := (float64(gen.Height) / float64(gen.Width)) * zoomWidth
	left := (real(viewCenter) - (zoomWidth / 2)) + pixelWidth/2
	top := (imag(viewCenter) - (viewHeight / 2)) + pixelHeight/2

	var waitGroup sync.WaitGroup
	waitGroup.Add(gen.Width)

	for x := 0; x < gen.Width; x++ {
		go func(xx int) {
			defer waitGroup.Done()
			for y := 0; y < gen.Height; y++ {
				coordinate := complex(left+float64(xx)*pixelWidth, top+float64(y)*pixelHeight)
				escapeIter := gen.escape(coordinate)
				img.Set(xx, y, palette[escapeIter])
			}
		}(x)
	}
	waitGroup.Wait()

	return img
}

func (gen *Generator) escape(c complex128) int {
	z := c
	for i := 0; i < gen.MaxEscape-1; i++ {
		if cmplx.Abs(z) > 2 {
			return i
		}
		z = cmplx.Pow(z, 2) + c
	}
	return gen.MaxEscape - 1
}
