package dependencies

import "os/exec"

func DependenciesJsExpress(p string) ([]byte, error) {
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

func DevDependenciesJsExpress(p string) ([]byte, error) {
	cmd := exec.Command(
		"yarn",
		"add",
		"nodemon",
		"--dev",
	)
	cmd.Dir = "./" + p
	out, err := cmd.Output()
	if err != nil {
		return []byte{}, err
	}
	return out, nil
}
