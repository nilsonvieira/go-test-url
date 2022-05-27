package main

import (
	"fmt"
	"log"
	"os"
	"tstnet/app"
)

func main() {
	fmt.Println("Executando Testes...")
	application := app.Tstnet()
	erro := application.Run(os.Args)
	if erro != nil {
		log.Fatal(erro)
	}
}
