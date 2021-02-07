package internal

import "github.com/lucabecci/project-generator/internal/templates"

type Template struct {
	project string
}

func TemplateGenerator(p string) *Template {
	template := Template{
		project: p,
	}
	return &template
}

//typescript
func (t *Template) TypescriptExpress() {
	templates.CreateTypescriptExpress(t.project)
}

//javascript
func (t *Template) JavascriptExpress() {
	templates.CreateJavascriptExpress(t.project)
}

func (t *Template) Compare(s string) {
	if s == "typescript-express" {
		t.TypescriptExpress()
	} else if s == "javascript-express" {
		t.JavascriptExpress()
	} else if s == "typescript-gql" {

	} else if s == "typescript-mongoose" {

	} else if s == "javascript-mongoose" {

	} else if s == "typescript-typeORM" {

	}
}
