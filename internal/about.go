package internal

import (
	"fmt"

	. "github.com/logrusorgru/aurora/v3"
)

//AboutMessage is for the principal message
func AboutMessage() {
	var title string = "JAVASCRIPT PROJECT GENERATOR by Luca Becci"
	var description string = "Use the diferents names for the avalibles templates "
	fmt.Println(BrightCyan(title))
	fmt.Println(Cyan(description))
	fmt.Println("---------------------")
	fmt.Println(Cyan("--typescript-express(TS, express, morgan, cors, ETC)"))
	fmt.Println("---------------------")
	fmt.Println(Cyan("--javascript-express(JS, express, morgan, cors, ETC)"))
	fmt.Println("---------------------")
	fmt.Println(Cyan("--typescript-gql(TS, express, typegraphql, Apollo-server, ETC)"))
	fmt.Println("---------------------")
	fmt.Println(Cyan("--typescript-mongoose(TS, express, mongoose, morgan, ETC)"))
	fmt.Println("---------------------")
	fmt.Println(Cyan("--javascript-mongoose(JS, express, mongoose, morgan, ETC)"))
	fmt.Println("---------------------")
	fmt.Println(Cyan("--typescript-typeORM(TS, express, TypeORM, Morgan, ETC.)"))
	fmt.Println("---------------------")
	fmt.Println("----IF YOU NEED SEE MORE TEMPLATES VISIT THE REPOSITORY OF THE PROJECT----")
}
