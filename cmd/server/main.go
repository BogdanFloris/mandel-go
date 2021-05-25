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

func indexHandler(width int, height int) func(http.ResponseWriter, *http.Request) {
	templateVars := make(map[string]interface{})
	templateVars["Width"] = width
	templateVars["Height"] = height
	tmpl, err := template.ParseFiles("./web/index.html")
	if err != nil {
		panic(err)
	}
	return func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := tmpl.Execute(w, templateVars); err != nil {
			panic(err)
		}
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
	width := config.GetInt("WIDTH", 1280)
	height := config.GetInt("HEIGHT", 1024)

	// Initialize router and run server
	router := mux.NewRouter()
	router.HandleFunc("/", indexHandler(width, height))
	router.HandleFunc("/mandelbrot", mandelbrot.HandleMandelbrot(width, height))

	server := &http.Server{
		Handler: router,
		Addr:    config.GetString("APP_HOST", "127.0.0.1") + ":" + config.GetString("APP_PORT", "8080"),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: config.GetDuration("WRITE_TIMEOUT", 15) * time.Second,
		ReadTimeout:  config.GetDuration("READ_TIMEOUT", 15) * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
