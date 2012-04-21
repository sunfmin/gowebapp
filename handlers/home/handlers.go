package home

import (
	. "github.com/paulbellamy/mango"
	"github.com/sunfmin/gowebapp/accounts"
	"github.com/sunfmin/gowebapp/handlers/forms"
	"github.com/sunfmin/mangotemplate"
)

func Index(env Env) (status Status, headers Headers, body Body) {
	a := &accounts.Account{
		Name:   "Felix",
		Gender: "2",
	}
	mangotemplate.ForRender(env, "home/index", forms.AccountRegisterForm().Render(a, env))
	return
}
