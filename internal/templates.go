package internal

import (
	"fmt"
	"os"

	. "github.com/logrusorgru/aurora/v3"
	"github.com/lucabecci/project-generator/internal/templates"
)

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
	templates.CreateTypescriptProject(t.project, "typescript-express")
}
func (t *Template) TypescriptMongoose() {
	templates.CreateTypescriptProject(t.project, "typescript-mongoose")
}
func (t *Template) TypescriptTypeOrm() {
	templates.CreateTypescriptProject(t.project, "typescript-typeORM")
}

//javascript
func (t *Template) JavascriptExpress() {
	templates.CreateJavascriptProject(t.project, "javascript-express")
}
func (t *Template) JavascriptMongoose() {
	templates.CreateJavascriptProject(t.project, "javascript-mongoose")
}

func (t *Template) Compare(s string) {
	if s == "typescript-express" {
		t.TypescriptExpress()
	} else if s == "javascript-express" {
		t.JavascriptExpress()
	} else if s == "typescript-gql" {
		fmt.Println(Blue("SOON"))
		os.Remove(t.project)
	} else if s == "typescript-mongoose" {
		t.TypescriptMongoose()
	} else if s == "javascript-mongoose" {
		t.JavascriptMongoose()
	} else if s == "typescript-typeORM" {
		t.TypescriptTypeOrm()
	} else {
		fmt.Println(Red("Error to create your template"))
		os.Remove(t.project)
	}
}
