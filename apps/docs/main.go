package main

import (
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	http.HandleFunc("/swagger/doc.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "doc.json")
	})

	http.Handle("/swagger/", httpSwagger.WrapHandler)

	http.ListenAndServe(":8081", nil)
}
