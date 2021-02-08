package templates

import (
	"fmt"
	"os"

	. "github.com/logrusorgru/aurora/v3"
)

func ErrorCreation(p string, err error) {
	fmt.Println(Red(string(err.Error())))
	fmt.Println(Red("Closing system..."))
	os.Remove(p)
	os.Exit(1)
}
