package main

import (
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Erro ao conectar ao servidor:", err)
	}
	defer client.Close()

	var sucesso bool
	err = client.Call("AdministracaoService.AbrirConta", "12345", &sucesso)
	if err != nil || !sucesso {
		log.Println("Erro ao abrir conta:", err)
	} else {
		log.Println("Conta aberta com sucesso")
	}

	var saldo float64
	err = client.Call("AdministracaoService.ConsultarSaldo", "12345", &saldo)
	if err != nil {
		log.Println("Erro ao consultar saldo:", err)
	} else {
		log.Printf("Saldo da conta 12345: %.2f\n", saldo)
	}
}
