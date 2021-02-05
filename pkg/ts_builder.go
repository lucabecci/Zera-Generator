package pkg

import (
	"fmt"
	"os/exec"
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
	//files ts
	cmd = exec.Command(
		"touch",
		"index.ts",
		"controllers/index.controller.ts",
		"routes/index.routes.ts",
		"database/database.ts",
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

func Dependencies(p string, t string) ([]byte, error) {
	fmt.Println("Adding dependencies...")
	var out []byte
	var err error
	if t == "ts-express" {
		out, err = DependenciesTsExpress(p)
		if err != nil {
			return []byte{}, err
		}
	}
	return out, nil
}

func DevDependencies(p string, t string) ([]byte, error) {
	fmt.Println("Adding dev-dependencies...")
	var out []byte
	var err error
	if t == "ts-express" {
		out, err = DevDependenciesTsExpress(p)
		if err != nil {
			return []byte{}, err
		}
	}
	return out, nil
}
