package mandelbrot

import (
	"image/png"
	"net/http"
	"strconv"
)

const (
	defaultRadius = 2.0
	defaultStartX = 0.0
	defaultStartY = 0.0
)

func safeFloat64(s string, def float64) *float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return &def
	}
	return &f
}

func Handle(generator Generator) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/png")
		generator.Radius = safeFloat64(r.FormValue("radius"), defaultRadius)
		generator.StartX = safeFloat64(r.FormValue("startX"), defaultStartX)
		generator.StartY = safeFloat64(r.FormValue("startY"), defaultStartY)
		img := generator.generateMandelbrotImage()
		err := png.Encode(w, img)
		if err != nil {
			panic(err)
		}
	}
}
