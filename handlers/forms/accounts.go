package forms

import (
	. "github.com/paulbellamy/mango"
	"github.com/sunfmin/goform"
)

func AccountRegisterForm() (f *goform.Builder) {
	f = goform.NewFormBuilder()
	f.TextField("Name")
	f.Select("Gender").Collection(func(fo goform.FormObject, env Env) goform.Options {
		return goform.StringOptions([][]string{
			{"1", "Men"},
			{"2", "Women"},
		})
	})
	f.RadioButtons("Gender").Collection(func(fo goform.FormObject, env Env) goform.Options {
		return goform.StringOptions([][]string{
			{"1", "Men"},
			{"2", "Women"},
		})
	})
	return
}
