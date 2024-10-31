package test

import (
	"TrabalhoDois/pkg/administracao"
	"testing"
)

func TestAbrirConta(t *testing.T) {
	service := administracao.NewService()
	var sucesso bool
	err := service.AbrirConta("12345", &sucesso)
	if err != nil || !sucesso {
		t.Errorf("Erro ao abrir conta: %v", err)
	}
}

func TestDepositar(t *testing.T) {
	service := administracao.NewService()
	var sucesso bool
	service.AbrirConta("12345", &sucesso)
	err := service.Depositar([2]interface{}{"12345", 100.0}, &sucesso)
	if err != nil || !sucesso {
		t.Errorf("Erro ao depositar: %v", err)
	}
}

func TestConsultarSaldo(t *testing.T) {
	service := administracao.NewService()
	var sucesso bool
	service.AbrirConta("12345", &sucesso)
	service.Depositar([2]interface{}{"12345", 100.0}, &sucesso)

	var saldo float64
	err := service.ConsultarSaldo("12345", &saldo)
	if err != nil || saldo != 100.0 {
		t.Errorf("Saldo incorreto: esperado 100.0, obteve %.2f", saldo)
	}
}
