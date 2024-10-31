package tests

import (
	"TrabalhoDois/internal/rpc"
	"TrabalhoDois/pkg/utils"
	"sync"
	"testing"
	"time"
)

const testServerAddress = "localhost:12345"

func setupTestServer() (*rpc.BancoServer, *rpc.BancoClient, func()) {
	server := rpc.NewBancoServer()
	go rpc.StartServer(testServerAddress)
	time.Sleep(time.Second)

	client, err := rpc.NewBancoClient(testServerAddress)
	if err != nil {
		panic(err)
	}

	cleanup := func() {
		client.Close()
	}

	return server, client, cleanup
}

// Teste de Idempotência
func TestIdempotentDeposit(t *testing.T) {
	_, client, cleanup := setupTestServer()
	defer cleanup()

	// Abre uma conta
	success, err := client.AbrirConta(1)
	if err != nil || !success {
		t.Fatalf("Falha ao abrir conta: %v", err)
	}

	idTransacao := utils.GenerateTransactionID()

	// Primeira tentativa de depósito
	success, err = client.Deposito(1, 100, idTransacao)
	if err != nil || !success {
		t.Errorf("Primeira tentativa de depósito falhou: %v", err)
	}

	// Segunda tentativa de depósito (deve ser ignorada)
	success, err = client.Deposito(1, 100, idTransacao)
	if err == nil || success {
		t.Errorf("Segunda tentativa de depósito não deveria ter sucesso: %v", err)
	}

	// Verifica o saldo
	saldo, err := client.ConsultaSaldo(1)
	if err != nil || saldo != 100 {
		t.Errorf("Saldo esperado: 100, obtido: %f, erro: %v", saldo, err)
	}
}

// Teste de Concorrência
func TestConcurrentDeposits(t *testing.T) {
	_, client, cleanup := setupTestServer()
	defer cleanup()

	// Abre uma conta
	success, err := client.AbrirConta(1)
	if err != nil || !success {
		t.Fatalf("Falha ao abrir conta: %v", err)
	}

	var wg sync.WaitGroup
	numGoroutines := 100
	valorDeposito := 10.0

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			idTransacao := utils.GenerateTransactionID()
			_, err := client.Deposito(1, valorDeposito, idTransacao)
			if err != nil {
				t.Errorf("Erro no depósito concorrente: %v, %s", err, idTransacao)
			}
		}()
	}

	wg.Wait()

	// Verifica o saldo final
	saldoEsperado := float64(numGoroutines) * valorDeposito
	saldo, err := client.ConsultaSaldo(1)
	if err != nil || saldo != saldoEsperado {
		t.Errorf("Saldo esperado: %f, obtido: %f, erro: %v", saldoEsperado, saldo, err)
	}
}

// Teste de Integração
func TestIntegracaoOperacoesBancarias(t *testing.T) {
	_, client, cleanup := setupTestServer()
	defer cleanup()

	// Teste de abertura de conta
	success, err := client.AbrirConta(1)
	if err != nil || !success {
		t.Fatalf("Falha ao abrir conta: %v", err)
	}

	// Teste de depósito
	success, err = client.Deposito(1, 1000, utils.GenerateTransactionID())
	if err != nil || !success {
		t.Errorf("Falha no depósito: %v", err)
	}

	// Teste de saque
	success, err = client.Saque(1, 500, utils.GenerateTransactionID())
	if err != nil || !success {
		t.Errorf("Falha no saque: %v", err)
	}

	// Verifica o saldo
	saldo, err := client.ConsultaSaldo(1)
	if err != nil || saldo != 500 {
		t.Errorf("Saldo esperado: 500, obtido: %f, erro: %v", saldo, err)
	}

	// Teste de saque com saldo insuficiente
	success, err = client.Saque(1, 1000, utils.GenerateTransactionID())
	if err == nil || success {
		t.Errorf("Saque com saldo insuficiente deveria falhar")
	}

	// Teste de fechamento de conta
	success, err = client.FecharConta(1)
	if err != nil || !success {
		t.Errorf("Falha ao fechar conta: %v", err)
	}

	// Tenta operar em conta fechada
	_, err = client.ConsultaSaldo(1)
	if err == nil {
		t.Errorf("Consulta em conta fechada deveria falhar")
	}
}

// Teste de Retry e Falha de Rede
func TestRetryAndNetworkFailure(t *testing.T) {
	_, client, cleanup := setupTestServer()
	defer cleanup()

	// Abre uma conta
	success, err := client.AbrirConta(1)
	if err != nil || !success {
		t.Fatalf("Falha ao abrir conta: %v", err)
	}

	// Simula uma operação com falhas de rede
	falhasSimuladas := 0
	operacaoComFalha := func() (bool, error) {
		if falhasSimuladas < 2 {
			falhasSimuladas++
			return false, err
		}
		return client.Deposito(1, 100, utils.GenerateTransactionID())
	}

	success, err = client.RetryOperation(operacaoComFalha, 3)
	if err != nil || !success {
		t.Errorf("Operação com retry falhou: %v", err)
	}

	// Verifica se o depósito foi realizado após as falhas
	saldo, err := client.ConsultaSaldo(1)
	if err != nil || saldo != 100 {
		t.Errorf("Saldo esperado após retry: 100, obtido: %f, erro: %v", saldo, err)
	}
}
