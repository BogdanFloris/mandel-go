package main

import (
	"github.com/gorilla/mux"
	"log"
	config "mandel-go/internal/common"
	"net/http"
	"path/filepath"
	"time"
)

func test(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Muie Mihnea!"))
	if err != nil {
		panic("")
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

	router := mux.NewRouter()
	router.HandleFunc("/", test)

	server := &http.Server{
		Handler: router,
		Addr:    config.GetString("APP_HOST", "127.0.0.1") + ":" + config.GetString("APP_PORT", "8080"),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: config.GetDuration("WRITE_TIMEOUT", 15) * time.Second,
		ReadTimeout:  config.GetDuration("READ_TIMEOUT", 15) * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
