package home

import (
	"github.com/sunfmin/gowebapp/layout"
	"net/http"
)

type HomeData struct {
	User string
}

func Index(w http.ResponseWriter, r *http.Request) {
	layout.Templates.ExecuteTemplate(w, "home/index", &HomeData{"Felix Sun"})
}
