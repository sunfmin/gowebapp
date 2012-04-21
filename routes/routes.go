package routes

import (
	"github.com/binku87/pat"
	. "github.com/paulbellamy/mango"
	"github.com/sunfmin/gowebapp/handlers/home"
	"github.com/sunfmin/mangotemplate"
	"html/template"
)

type LayoutData struct {
	Name string
}

type provider struct{}

func (h *provider) LayoutData(env Env) (d interface{}) {
	d = &LayoutData{}
	return
}

func NewMux() (r *pat.PatternServeMux) {

	tpl, err := template.ParseGlob("templates/*/*.html")
	if err != nil {
		panic(err)
	}

	layout := mangotemplate.MakeLayout(tpl, "main", &provider{})
	renderer := mangotemplate.MakeRenderer(tpl)

	s := new(Stack)
	s.Middleware(layout, renderer)

	m := pat.New()
	m.Get("/", s.HandlerFunc(home.Index))
	return m
}
