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
	err = client.Call("AdministracaoService.Depositar", [2]interface{}{"12345", 100.0}, &sucesso)
	if err != nil || !sucesso {
		log.Println("Erro ao depositar:", err)
	} else {
		log.Println("Depósito realizado com sucesso")
	}

	var saldo float64
	err = client.Call("AdministracaoService.ConsultarSaldo", "12345", &saldo)
	if err != nil {
		log.Println("Erro ao consultar saldo:", err)
	} else {
		log.Printf("Saldo da conta 12345: %.2f\n", saldo)
	}
}
