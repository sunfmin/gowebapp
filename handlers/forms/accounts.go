package forms

import (
	. "github.com/paulbellamy/mango"
	"github.com/sunfmin/goform"
)

func AccountRegisterForm() (f *goform.Builder) {
	f = goform.NewFormBuilder()
	f.TextField("Name").Placeholder("Type your name")

	f.RadioButtons("Gender").Collection(func(fo goform.FormObject, env Env) goform.Options {
		return goform.StringOptions([][]string{
			{"1", "Men"},
			{"2", "Women"},
		})
	})

	f.Select("Department").Collection(func(fo goform.FormObject, env Env) goform.Options {
		return goform.StringOptions([][]string{
			{"HR", "Human Resource"},
			{"IT", "IT Development"},
		})
	})
	return
}
