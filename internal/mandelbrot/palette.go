package mandelbrot

import (
	"image/color"
	"math/rand"
)

// Palette enum
type Palette int

const (
	Random Palette = iota
)

func (p Palette) String() string {
	return [...]string{"Random"}[p]
}

func (p Palette) GetPalette(maxEscape int) []color.Color {
	palette := make([]color.Color, maxEscape)

	switch p {
	case Random:
		for i := 0; i < maxEscape-1; i++ {
			palette[i] = color.RGBA{
				R: uint8(rand.Intn(256)),
				G: uint8(rand.Intn(256)),
				B: uint8(rand.Intn(256)),
				A: 255,
			}
		}
		palette[maxEscape-1] = color.RGBA{R: 0, G: 0, B: 0, A: 0}
	}

	return palette
}
