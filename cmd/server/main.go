package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	config "mandel-go/internal/common"
	"mandel-go/internal/mandelbrot"
	"net/http"
	"path/filepath"
	"time"
)

func indexHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	templateVars := make(map[string]interface{})
	templateVars["Width"] = mandelbrot.Width
	templateVars["Height"] = mandelbrot.Height
	tmpl, err := template.ParseFiles("./web/index.html")
	if err != nil {
		panic(err)
	}
	if err := tmpl.Execute(w, templateVars); err != nil {
		panic(err)
	}
}

func main() {
	// Get configuration
	configPath, err := filepath.Abs("./config/.env")
	if err != nil {
		log.Fatal(err)
	}

	err = config.Load(configPath)
	if err != nil {
		log.Fatal(err)
	}

	// Load width and height
	mandelbrot.Width = config.GetInt("WIDTH", 1280)
	mandelbrot.Height = config.GetInt("HEIGHT", 1024)

	// Initialize router and run server
	router := mux.NewRouter()
	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/mandelbrot", mandelbrot.HandleMandelbrot)

	server := &http.Server{
		Handler: router,
		Addr:    config.GetString("APP_HOST", "127.0.0.1") + ":" + config.GetString("APP_PORT", "8080"),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: config.GetDuration("WRITE_TIMEOUT", 15) * time.Second,
		ReadTimeout:  config.GetDuration("READ_TIMEOUT", 15) * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
