package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"./controllers"

	_ "github.com/go-sql-driver/mysql"
)

const (
	defaultPort = "8008"

	idleTimeout       = 30 * time.Second
	writeTimeout      = 180 * time.Second
	readHeaderTimeout = 10 * time.Second
	readTimeout       = 10 * time.Second
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	log.Println("port:", port)

	// route
	handler := controllers.New()

	server := &http.Server{
		Addr:    "0.0.0.0:" + port,
		Handler: handler,

		IdleTimeout:       idleTimeout,
		WriteTimeout:      writeTimeout,
		ReadHeaderTimeout: readHeaderTimeout,
		ReadTimeout:       readTimeout,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Println("ERR ListenAndServe:", err)
	}
}
