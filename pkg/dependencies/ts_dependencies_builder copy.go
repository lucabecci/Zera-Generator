package dependencies

import "os/exec"

func DependenciesTsExpress(p string) ([]byte, error) {
	cmd := exec.Command(
		"yarn",
		"add",
		"express",
		"morgan",
		"cors",
		"dotenv",
	)
	cmd.Dir = "./" + p
	out, err := cmd.Output()
	if err != nil {
		return []byte{}, err
	}
	return out, nil
}

func DevDependenciesTsExpress(p string) ([]byte, error) {
	cmd := exec.Command(
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
	cmd.Dir = "./" + p
	out, err := cmd.Output()
	if err != nil {
		return []byte{}, err
	}
	return out, nil
}
