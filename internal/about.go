package internal

import (
	"fmt"

	. "github.com/logrusorgru/aurora/v3"
)

//AboutMessage is for the principal message
func AboutMessage() {
	var title string = "Zera Generator by Luca Becci"
	fmt.Println(BrightCyan(title))
	fmt.Println(Cyan("Check github.com/lucabecc/zera-generator for avalibles templates"))
	fmt.Println("---------------------")
	fmt.Println(BrightMagenta("Started 1: ts-express"))
	fmt.Println(BrightMagenta("Started 2: js-express"))
	fmt.Println("---------------------")
}
