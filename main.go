package main

import (
	"github.com/sunfmin/gowebapp/routes"
	"net/http"
)

func main() {
	http.ListenAndServe("localhost:5000", routes.NewMux())
}
