package pkg

import (
	"fmt"
	"io/ioutil"
	"os/exec"

	"github.com/lucabecci/project-generator/pkg/dependencies"
)

func InitializeJS(p string) ([]byte, error) {
	fmt.Println("Yarn build...")
	cmd := exec.Command("yarn", "init", "-y")
	cmd.Dir = "./" + p
	out, err := cmd.Output()
	if err != nil {
		return []byte{}, err
	}
	return out, nil
}

func StructureJS(p string) error {
	fmt.Println("Structure build...")
	cmd := exec.Command(
		"touch",
		"README.md",
		".gitignore",
		"Dockerfile",
		".dockerignore",
		".env",
	)
	cmd.Dir = "./" + p
	err := cmd.Run()
	if err != nil {
		return err
	}
	//folders
	fmt.Println("API Structure build...")
	cmd = exec.Command(
		"mkdir",
		"src",
		"src/routes",
		"src/controllers",
		"src/database",
		"src/middlewares",
		"src/helpers",
	)
	cmd.Dir = "./" + p
	err = cmd.Run()
	if err != nil {
		return err
	}
	fmt.Println("API filesbuild...")
	//files js
	cmd = exec.Command(
		"touch",
		"index.js",
		"controllers/index.controller.js",
		"routes/index.routes.js",
		"database/database.js",
		"middlewares/auth.js",
		"helpers/checks.js",
	)
	cmd.Dir = "./" + p + "/src"
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func DependenciesJS(p string, t string) ([]byte, error) {
	fmt.Println("Adding dependencies...")
	var out []byte
	var err error
	if t == "js-express" {
		out, err = dependencies.DependenciesJsExpress(p)
		if err != nil {
			return []byte{}, err
		}
	}
	return out, nil
}

func DevDependenciesJS(p string, t string) ([]byte, error) {
	fmt.Println("Adding dev-dependencies...")
	var out []byte
	var err error
	if t == "js-express" {
		out, err = dependencies.DevDependenciesJsExpress(p)
		if err != nil {
			return []byte{}, err
		}
	}
	return out, nil
}

func AppDataJS(p string) error {
	fmt.Println("Configuring the app server...")
	content, err := ioutil.ReadFile("./pkg/file_information/indexJS.txt")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(p+"/src/index.js", content, 0777)
	if err != nil {
		return err
	}
	return nil
}
