package main

import (
	"fmt"
	"log"

	"gopkg.in/ini.v1"
)

func main() {

	ArquivoINI, err := ini.Load("Informacoes.INI")
	if err != nil {
		log.Printf("Erro ao ler arquivo: %v", err)
	}

	fmt.Print(ArquivoINI)
}
