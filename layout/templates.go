package layout

import (
	"net/http"
	"text/template"
)

var Templates *template.Template

func init() {
	var err error
	Templates, err = template.ParseGlob("templates/*/*.html")
	if err != nil {
		panic(err)
	}
}

type Header struct {
}

type Footer struct {
}

func MainLayout(handler http.HandlerFunc) (r http.HandlerFunc) {

	return func(w http.ResponseWriter, r *http.Request) {
		header := &Header{}
		footer := &Footer{}
		Templates.ExecuteTemplate(w, "mainHeader", header)
		handler(w, r)
		Templates.ExecuteTemplate(w, "mainFooter", footer)
	}

}

type SimpleHeader struct {
}

func SimpleLayout(handler http.HandlerFunc) (r http.HandlerFunc) {
	return func(w http.ResponseWriter, r *http.Request) {
		header := &SimpleHeader{}
		footer := &Footer{}
		Templates.ExecuteTemplate(w, "simpleHeader", header)
		handler(w, r)
		Templates.ExecuteTemplate(w, "simpleFooter", footer)
	}

}
