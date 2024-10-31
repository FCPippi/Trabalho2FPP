package test

import (
	"TrabalhoDois/pkg/administracao"
	"sync"
	"testing"
)

func TestConcorrencia(t *testing.T) {
	service := administracao.NewService()
	var sucesso bool
	service.AbrirConta("12345", &sucesso)

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			service.Depositar([2]interface{}{"12345", 1.0}, &sucesso)
		}()
	}
	wg.Wait()

	var saldo float64
	service.ConsultarSaldo("12345", &saldo)
	if saldo != 100.0 {
		t.Errorf("Saldo incorreto após operações concorrentes: esperado 100.0, obteve %.2f", saldo)
	}
}
