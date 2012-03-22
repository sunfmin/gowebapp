package routes

import (
	"github.com/bmizerany/pat"
	"github.com/sunfmin/gowebapp/accounts"
	"github.com/sunfmin/gowebapp/home"
	"github.com/sunfmin/gowebapp/layout"
	"net/http"
)

func NewMux() (r *pat.PatternServeMux) {
	m := pat.New()
	m.Get("/", http.HandlerFunc(layout.MainLayout(home.Index)))
	m.Get("/accounts", http.HandlerFunc(layout.MainLayout(accounts.Index)))
	m.Get("/accounts/:id/edit", http.HandlerFunc(layout.MainLayout(accounts.Edit)))
	m.Get("/accounts/:id/popup_edit", http.HandlerFunc(layout.SimpleLayout(accounts.Edit)))

	return m
}
