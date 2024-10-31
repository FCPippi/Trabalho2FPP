package main

import (
	"TrabalhoDois/internal/rpc"
	"TrabalhoDois/internal/simulador"
	"TrabalhoDois/pkg/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	client, err := rpc.NewBancoClient("localhost:1234")
	if err != nil {
		log.Fatalf("Erro ao conectar ao servidor: %v", err)
	}
	defer client.Close()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nEscolha uma operação:")
		fmt.Println("1. Abrir Conta")
		fmt.Println("2. Depósito")
		fmt.Println("3. Consultar Saldo")
		fmt.Println("4. Sair")
		fmt.Print("Opção: ")

		scanner.Scan()
		opcao := scanner.Text()

		switch opcao {
		case "1":
			fmt.Print("Número da conta: ")
			scanner.Scan()
			numeroConta, _ := strconv.Atoi(scanner.Text())

			success, err := client.RetryOperation(func() (bool, error) {
				if simulador.SimularFalhaRede() {
					return false, fmt.Errorf("falha de rede simulada")
				}
				return client.AbrirConta(numeroConta)
			}, 3)

			if err != nil {
				fmt.Printf("Erro ao abrir conta: %v\n", err)
			} else if success {
				fmt.Printf("Conta %d aberta com sucesso\n", numeroConta)
			}

		case "2":
			fmt.Print("Número da conta: ")
			scanner.Scan()
			numeroConta, _ := strconv.Atoi(scanner.Text())

			fmt.Print("Valor: ")
			scanner.Scan()
			valor, _ := strconv.ParseFloat(scanner.Text(), 64)

			idTransacao := utils.GenerateTransactionID()

			success, err := client.RetryOperation(func() (bool, error) {
				if simulador.SimularFalhaRede() {
					return false, fmt.Errorf("falha de rede simulada")
				}
				return client.Deposito(numeroConta, valor, idTransacao)
			}, 3)

			if err != nil {
				fmt.Printf("Erro ao realizar depósito: %v\n", err)
			} else if success {
				fmt.Printf("Depósito de %.2f realizado com sucesso na conta %d (ID da transação: %s)\n", valor, numeroConta, idTransacao)
			}

		case "3":
			fmt.Print("Número da conta: ")
			scanner.Scan()
			numeroConta, _ := strconv.Atoi(scanner.Text())

			saldo, err := client.ConsultaSaldo(numeroConta)
			if err != nil {
				fmt.Printf("Erro ao consultar saldo: %v\n", err)
			} else {
				fmt.Printf("Saldo da conta %d: %.2f\n", numeroConta, saldo)
			}

		case "4":
			fmt.Println("Encerrando o programa.")
			return

		default:
			fmt.Println("Opção inválida.")
		}
	}
}
