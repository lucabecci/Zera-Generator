package templates

import (
	"fmt"

	. "github.com/logrusorgru/aurora/v3"
	"github.com/lucabecci/project-generator/pkg"
)

func CreateJavascriptProject(p string, t string) {
	build, err := pkg.InitializeJS(p)
	if err != nil {
		ErrorCreation(p, err)
		return
	}
	fmt.Println(Green(string(build)))

	err = pkg.StructureJS(p)
	if err != nil {
		ErrorCreation(p, err)
		return
	}

	prod, err := pkg.DependenciesJS(p, t)
	if err != nil {
		ErrorCreation(p, err)
		return
	}
	fmt.Println(Green(string(prod)))

	dev, err := pkg.DevDependenciesJS(p, t)
	if err != nil {
		ErrorCreation(p, err)
		return
	}
	fmt.Println(Green(string(dev)))

	err = pkg.GitIgnoreData(p)
	if err != nil {
		ErrorCreation(p, err)
		return
	}
	err = pkg.AppDataJS(p, t)
	if err != nil {
		ErrorCreation(p, err)
		return
	}
	err = pkg.ReadmeInfo(p)
	if err != nil {
		ErrorCreation(p, err)
		return
	}
	fmt.Println(Green("------------------------------------------------------"))
	fmt.Println(Green("------------------------------------------------------"))
	fmt.Println(Green("PROJECT CREATED, PLEASE READ THE README FOR CORRECT USAGE"))
	fmt.Println(Green("------------------------------------------------------"))
	fmt.Println(Green("------------------------------------------------------"))
}
