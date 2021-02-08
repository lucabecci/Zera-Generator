package pkg

import (
	"fmt"
	"io/ioutil"
	"os/exec"

	"github.com/lucabecci/zera-generator/pkg/dependencies"
)

func InitializeTS(p string) ([]byte, error) {
	fmt.Println("Yarn build...")
	cmd := exec.Command("yarn", "init", "-y")
	cmd.Dir = "./" + p
	out, err := cmd.Output()
	if err != nil {
		return []byte{}, err
	}
	return out, nil
}

func StructureTS(p string) error {
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
		"src/middlewares",
		"src/helpers",
	)
	cmd.Dir = "./" + p
	err = cmd.Run()
	if err != nil {
		return err
	}
	fmt.Println("API filesbuild...")
	//files ts
	cmd = exec.Command(
		"touch",
		"index.ts",
		"controllers/index.controller.ts",
		"routes/index.routes.ts",
		"middlewares/auth.ts",
		"helpers/checks.ts",
	)
	cmd.Dir = "./" + p + "/src"
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func DependenciesTS(p string, t string) ([]byte, error) {
	fmt.Println("Adding dependencies...")
	var out []byte
	var err error
	if t == "typescript-express" {
		out, err = dependencies.DependenciesTsExpress(p, t)
		if err != nil {
			return []byte{}, err
		}
	} else if t == "typescript-mongoose" {
		out, err = dependencies.DependenciesTsExpress(p, t)
		if err != nil {
			return []byte{}, err
		}
	} else if t == "typescript-typeORM" {
		out, err = dependencies.DependenciesTsExpress(p, t)
		if err != nil {
			return []byte{}, err
		}
	}
	return out, nil
}

func DevDependenciesTS(p string, t string) ([]byte, error) {
	fmt.Println("Adding dev-dependencies...")
	var out []byte
	var err error
	if t == "typescript-express" {
		out, err = dependencies.DevDependenciesTsExpress(p, t)
		if err != nil {
			return []byte{}, err
		}
	} else if t == "typescript-mongoose" {
		out, err = dependencies.DevDependenciesTsExpress(p, t)
		if err != nil {
			return []byte{}, err
		}
	} else if t == "typescript-typeORM" {
		out, err = dependencies.DevDependenciesTsExpress(p, t)
		if err != nil {
			return []byte{}, err
		}
	}
	return out, nil
}

func InitTsConfig(p string) ([]byte, error) {
	fmt.Println("Creating a ts config...")
	cmd := exec.Command("tsc", "--init")
	cmd.Dir = "./" + p
	out, err := cmd.Output()
	if err != nil {
		return []byte{}, err
	}
	return out, nil
}

func AppDataTS(p string, t string) error {
	var content []byte
	var err error

	fmt.Println("Configuring the app server...")
	if t == "typescript-express" {
		content, err = ioutil.ReadFile("./pkg/file_information/ts/index-express.txt")
		if err != nil {
			return err
		}
	} else if t == "typescript-mongoose" {
		content, err = ioutil.ReadFile("./pkg/file_information/ts/index-mongoose.txt")
		if err != nil {
			return err
		}
	} else if t == "typescript-typeORM" {
		content, err = ioutil.ReadFile("./pkg/file_information/ts/index-typeORM.txt")
		if err != nil {
			return err
		}
	}
	err = ioutil.WriteFile(p+"/src/index.ts", content, 0777)
	if err != nil {
		return err
	}
	return nil
}
