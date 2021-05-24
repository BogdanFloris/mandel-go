package mandelbrot

import (
	"image"
	"image/color"
	"image/png"
	"net/http"
)

var Width int
var Height int

func HandleMandelbrot(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "image/png")
	img := image.NewRGBA(image.Rect(0, 0, Width, Height))

	for x := 0; x < Width; x++ {
		for y := 0; y < Height; y++ {
			img.Set(x, y, color.Black)
		}
	}

	err := png.Encode(w, img)
	if err != nil {
		panic(err)
	}
}
