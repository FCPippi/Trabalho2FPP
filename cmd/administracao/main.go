package main

import (
	"TrabalhoDois/internal/rpc"
	"fmt"
	"log"
)

func main() {

	fmt.Println("Iniciando servidor de administração...")
	err := rpc.StartServer(":1234")

	if err != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}
}
