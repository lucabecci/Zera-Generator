package dependencies

import "os/exec"

func DependenciesTsExpress(p string, t string) ([]byte, error) {
	cmd := exec.Command("null")
	if t == "typescript-express" {
		cmd = exec.Command(
			"yarn",
			"add",
			"express",
			"morgan",
			"cors",
			"dotenv",
		)
	} else if t == "typescript-mongoose" {
		cmd = exec.Command(
			"yarn",
			"add",
			"express",
			"morgan",
			"mongoose",
			"cors",
			"dotenv",
		)
	} else if t == "typescript-typeORM" {
		cmd = exec.Command(
			"yarn",
			"add",
			"express",
			"morgan",
			"typeorm",
			"cors",
			"dotenv",
		)
	}
	cmd.Dir = "./" + p
	out, err := cmd.Output()
	if err != nil {
		return []byte{}, err
	}
	return out, nil
}

func DevDependenciesTsExpress(p string, t string) ([]byte, error) {
	cmd := exec.Command("null")
	if t == "typescript-express" {
		cmd = exec.Command(
			"yarn",
			"add",
			"@types/express",
			"@types/morgan",
			"@types/cors",
			"@types/node",
			"typescript",
			"ts-node-dev",
			"--dev",
		)
	} else if t == "typescript-mongoose" {
		cmd = exec.Command(
			"yarn",
			"add",
			"@types/express",
			"@types/morgan",
			"@types/cors",
			"@types/node",
			"@types/mongoose",
			"typescript",
			"ts-node-dev",
			"--dev",
		)
	} else if t == "typescript-typeORM" {
		cmd = exec.Command(
			"yarn",
			"add",
			"@types/express",
			"@types/morgan",
			"@types/cors",
			"@types/node",
			"typescript",
			"ts-node-dev",
			"--dev",
		)
	}
	cmd.Dir = "./" + p
	out, err := cmd.Output()
	if err != nil {
		return []byte{}, err
	}
	return out, nil
}
