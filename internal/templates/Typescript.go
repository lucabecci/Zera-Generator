package templates

import (
	"fmt"

	. "github.com/logrusorgru/aurora/v3"
	"github.com/lucabecci/project-generator/pkg"
)

func CreateTypescriptProject(p string, t string) {
	build, err := pkg.InitializeTS(p)
	if err != nil {
		fmt.Println(Red(string(err.Error())))
	}
	fmt.Println(Green(string(build)))

	err = pkg.StructureTS(p)
	if err != nil {
		fmt.Println(Red(string(err.Error())))
	}

	prod, err := pkg.DependenciesTS(p, t)
	if err != nil {
		fmt.Println(Red(string(err.Error())))
	}
	fmt.Println(Green(string(prod)))

	dev, err := pkg.DevDependenciesTS(p, t)
	if err != nil {
		fmt.Println(Red(string(err.Error())))
	}
	fmt.Println(Green(string(dev)))

	out, err := pkg.InitTsConfig(p)
	if err != nil {
		fmt.Println(Red(string(err.Error())))
	}
	fmt.Println(Green(string(out)))

	err = pkg.GitIgnoreData(p)
	if err != nil {
		fmt.Println(Red(string(err.Error())))
	}
	err = pkg.AppDataTS(p)
	if err != nil {
		fmt.Println(Red(string(err.Error())))
	}
}
