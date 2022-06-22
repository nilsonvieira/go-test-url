package main

import (
	"fmt"
	"log"
	"os"
	"tstnet/app"
)

func main() {
	fmt.Println("Executing Tests...")
	application := app.Tstnet()
	erro := application.Run(os.Args)
	if erro != nil {
		log.Fatal(erro)
	}
}
