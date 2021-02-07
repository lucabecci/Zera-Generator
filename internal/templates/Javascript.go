package templates

import (
	"fmt"
	"os"

	. "github.com/logrusorgru/aurora/v3"
	"github.com/lucabecci/project-generator/pkg"
)

func CreateJavascriptProject(p string, t string) {
	build, err := pkg.InitializeJS(p)
	if err != nil {
		fmt.Println(Red(string(err.Error())))
		fmt.Println(Red("Closing system..."))
		os.Exit(1)
	}
	fmt.Println(Green(string(build)))

	err = pkg.StructureJS(p)
	if err != nil {
		fmt.Println(Red(string(err.Error())))
		fmt.Println(Red("Closing system..."))
		os.Exit(1)
	}

	prod, err := pkg.DependenciesJS(p, t)
	if err != nil {
		fmt.Println(Red(string(err.Error())))
		fmt.Println(Red("Closing system..."))
		os.Exit(1)
	}
	fmt.Println(Green(string(prod)))

	dev, err := pkg.DevDependenciesJS(p, t)
	if err != nil {
		fmt.Println(Red(string(err.Error())))
		fmt.Println(Red("Closing system..."))
		os.Exit(1)
	}
	fmt.Println(Green(string(dev)))

	err = pkg.GitIgnoreData(p)
	if err != nil {
		fmt.Println(Red(string(err.Error())))
		fmt.Println(Red("Closing system..."))
		os.Exit(1)
	}
	err = pkg.AppDataJS(p, t)
	if err != nil {
		fmt.Println(Red(string(err.Error())))
		fmt.Println(Red("Closing system..."))
		os.Exit(1)
	}
	fmt.Println(Green("------------------------------------------------------"))
	fmt.Println(Green("------------------------------------------------------"))
	fmt.Println(Green("PROJECT CREATED, PLEASE READ THE README FOR CORRECT USAGE"))
	fmt.Println(Green("------------------------------------------------------"))
	fmt.Println(Green("------------------------------------------------------"))
}
