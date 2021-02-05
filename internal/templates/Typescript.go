package templates

import (
	"fmt"

	. "github.com/logrusorgru/aurora/v3"
	"github.com/lucabecci/project-generator/pkg"
)

func CreateTypescriptExpress(p string) {
	build, err := pkg.InitializeTS(p)
	if err != nil {
		fmt.Println(Red(string(err.Error())))
	}
	fmt.Println(Green(string(build)))

	err = pkg.StructureTS(p)
	if err != nil {
		fmt.Println(Red(string(err.Error())))
	}

	prod, err := pkg.Dependencies(p, "ts-express")
	if err != nil {
		fmt.Println(Red(string(err.Error())))
	}
	fmt.Println(Green(string(prod)))

	dev, err := pkg.DevDependencies(p, "ts-express")
	if err != nil {
		fmt.Println(Red(string(err.Error())))
	}
	fmt.Println(Green(string(dev)))

}

func CreateTypescriptGQL() {
	fmt.Println("TS-GQL")
}
