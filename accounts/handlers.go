package accounts

import (
	"github.com/sunfmin/gowebapp/layout"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	layout.Templates.ExecuteTemplate(w, "accounts/index", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	layout.Templates.ExecuteTemplate(w, "accounts/edit", nil)
}
