package main

import (
	. "github.com/paulbellamy/mango"
	"github.com/sunfmin/gowebapp/routes"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", routes.NewMux())
	mux.HandleFunc("/assets/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public"+r.URL.Path)
	})

	err := http.ListenAndServe("localhost:8000", mux)
	if err != nil {
		panic(err)
	}
}
