package home

import (
	"github.com/sunfmin/gowebapp/layout"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	layout.Templates.ExecuteTemplate(w, "home/index", nil)
}
