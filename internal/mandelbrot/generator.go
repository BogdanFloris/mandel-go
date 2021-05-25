package mandelbrot

import "image"

type Generator struct {
	Width     int
	Height    int
	MaxEscape int
}

func (gen *Generator) generateMandelbrotImage() image.Image {
	return image.NewRGBA(image.Rect(0, 0, gen.Width, gen.Height))
}
