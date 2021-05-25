package mandelbrot

import (
	"image"
	"image/color"
	"image/png"
	"net/http"
)

func HandleMandelbrot(width int, height int) func(http.ResponseWriter, *http.Request) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, color.Black)
		}
	}

	return func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "image/png")
		err := png.Encode(w, img)
		if err != nil {
			panic(err)
		}
	}
}
