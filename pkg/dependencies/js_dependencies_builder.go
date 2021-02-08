package dependencies

import "os/exec"

func DependenciesJsExpress(p string, t string) ([]byte, error) {
	cmd := exec.Command("null")
	if t == "javascript-express" {
		cmd = exec.Command(
			"yarn",
			"add",
			"express",
			"morgan",
			"cors",
			"dotenv",
		)
	} else if t == "javascript-mongoose" {
		cmd = exec.Command(
			"yarn",
			"add",
			"express",
			"morgan",
			"mongoose",
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

func DevDependenciesJsExpress(p string, t string) ([]byte, error) {
	cmd := exec.Command("null")
	if t == "javascript-express" {
		cmd = exec.Command(
			"yarn",
			"add",
			"nodemon",
			"--dev",
		)
	} else if t == "javascript-mongoose" {
		cmd = exec.Command(
			"yarn",
			"add",
			"nodemon",
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
