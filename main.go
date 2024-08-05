package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

// @title Example API
// @version 1.0
// @description This is a sample server.
// @host localhost:8080
// @BasePath /

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/swagger/", serveSwagger)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// helloHandler godoc
// @Summary Show a Hello World message
// @Description get string "Hello World"
// @Tags example
// @Accept  json
// @Produce  json
// @Success 200 {string} string "Hello World"
// @Router /hello [get]
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func serveSwagger(w http.ResponseWriter, r *http.Request) {
	// Удаляем префикс /swagger/ из пути
	p := strings.TrimPrefix(r.URL.Path, "/swagger/")
	if p == "" {
		p = "index.html"
	}
	p = filepath.Join("swagger", "docs", p)
	http.ServeFile(w, r, p)
}
