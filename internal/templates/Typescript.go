package templates

import (
	"fmt"

	. "github.com/logrusorgru/aurora/v3"
	"github.com/lucabecci/zera-generator/pkg"
)

func CreateTypescriptProject(p string, t string) {
	build, err := pkg.InitializeTS(p)
	if err != nil {
		ErrorCreation(p, err)
		return
	}
	fmt.Println(Green(string(build)))

	err = pkg.StructureTS(p)
	if err != nil {
		ErrorCreation(p, err)
		return
	}

	prod, err := pkg.DependenciesTS(p, t)
	if err != nil {
		ErrorCreation(p, err)
		return
	}
	fmt.Println(Green(string(prod)))

	dev, err := pkg.DevDependenciesTS(p, t)
	if err != nil {
		ErrorCreation(p, err)
		return
	}
	fmt.Println(Green(string(dev)))

	out, err := pkg.InitTsConfig(p)
	if err != nil {
		ErrorCreation(p, err)
		return
	}
	fmt.Println(Green(string(out)))

	err = pkg.GitIgnoreData(p)
	if err != nil {
		ErrorCreation(p, err)
		return
	}
	err = pkg.AppDataTS(p, t)
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
