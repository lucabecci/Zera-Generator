package internal

import (
	"fmt"
	"os"

	. "github.com/logrusorgru/aurora/v3"
	"github.com/lucabecci/zera-generator/internal/templates"
)

//Template is a struct principal
type Template struct {
	project string
}

//TemplateGenerator of the template with the project folder path
func TemplateGenerator(p string) *Template {
	template := Template{
		project: p,
	}
	return &template
}

//TypescriptExpress is for generate ts-express template
func (t *Template) TypescriptExpress() {
	templates.CreateTypescriptProject(t.project, "typescript-express")
}

//TypescriptMongoose is for generate ts-mongoose template
func (t *Template) TypescriptMongoose() {
	templates.CreateTypescriptProject(t.project, "typescript-mongoose")
}

//TypescriptTypeOrm is for generate ts-typeORM template
func (t *Template) TypescriptTypeOrm() {
	templates.CreateTypescriptProject(t.project, "typescript-typeORM")
}

//JavascriptExpress is for generate js-express template
func (t *Template) JavascriptExpress() {
	templates.CreateJavascriptProject(t.project, "javascript-express")
}

//JavascriptMongoose is for generate js-mongoose template
func (t *Template) JavascriptMongoose() {
	templates.CreateJavascriptProject(t.project, "javascript-mongoose")
}

//Compare is for compare the user template
func (t *Template) Compare(s string) {
	if s == "ts-express" {
		t.TypescriptExpress()
	} else if s == "js-express" {
		t.JavascriptExpress()
	} else if s == "ts-gql" {
		fmt.Println(Blue("SOON"))
		os.Remove(t.project)
	} else if s == "ts-mongoose" {
		t.TypescriptMongoose()
	} else if s == "js-mongoose" {
		t.JavascriptMongoose()
	} else if s == "ts-typeORM" {
		t.TypescriptTypeOrm()
	} else {
		fmt.Println(Red("Error to create your template"))
		os.Remove(t.project)
	}
}
