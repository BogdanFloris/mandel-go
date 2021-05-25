package mandelbrot

import (
	"image/png"
	"net/http"
)

func Handle(generator Generator) func(http.ResponseWriter, *http.Request) {
	img := generator.generateMandelbrotImage()

	return func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "image/png")
		err := png.Encode(w, img)
		if err != nil {
			panic(err)
		}
	}
}
