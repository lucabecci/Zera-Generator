package templates

import (
	"fmt"

	. "github.com/logrusorgru/aurora/v3"
	"github.com/lucabecci/project-generator/pkg"
)

func CreateJavascriptProject(p string, t string) {
	build, err := pkg.InitializeJS(p)
	if err != nil {
		fmt.Println(Red(string(err.Error())))
	}
	fmt.Println(Green(string(build)))

	err = pkg.StructureJS(p)
	if err != nil {
		fmt.Println(Red(string(err.Error())))
	}

	prod, err := pkg.DependenciesJS(p, t)
	if err != nil {
		fmt.Println(Red(string(err.Error())))
	}
	fmt.Println(Green(string(prod)))

	dev, err := pkg.DevDependenciesJS(p, t)
	if err != nil {
		fmt.Println(Red(string(err.Error())))
	}
	fmt.Println(Green(string(dev)))

	err = pkg.GitIgnoreData(p)
	if err != nil {
		fmt.Println(Red(string(err.Error())))
	}
	err = pkg.AppDataJS(p)
	if err != nil {
		fmt.Println(Red(string(err.Error())))
	}
}
