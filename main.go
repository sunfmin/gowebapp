package gowebapp

import (
	"net/http"
)

func main() {
	http.ListenAndServe("localhost:5000", NewMux())
}
