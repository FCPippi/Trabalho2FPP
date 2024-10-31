package banco

import (
	"errors"
	"sync"
)

type Banco struct {
	Contas map[int]*Conta
	mutex  sync.RWMutex
}

func NovoBanco() *Banco {
	return &Banco{
		Contas: make(map[int]*Conta),
	}
}

func (b *Banco) AbrirConta(numero int) error {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if _, exists := b.Contas[numero]; exists {
		return errors.New("conta já existe")
	}

	b.Contas[numero] = NovaConta(numero)
	return nil
}

func (b *Banco) FecharConta(numero int) error {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if _, exists := b.Contas[numero]; !exists {
		return errors.New("conta não existe")
	}

	delete(b.Contas, numero)
	return nil
}

func (b *Banco) Deposito(numero int, valor float64, idTransacao string) error {
	b.mutex.RLock()
	conta, exists := b.Contas[numero]
	b.mutex.RUnlock()

	if !exists {
		return errors.New("conta não existe")
	}

	if !conta.Deposito(valor, idTransacao) {
		return errors.New("transação já processada ou falha no depósito")
	}

	return nil
}

func (b *Banco) Saque(numero int, valor float64, idTransacao string) error {
	b.mutex.RLock()
	conta, exists := b.Contas[numero]
	b.mutex.RUnlock()

	if !exists {
		return errors.New("conta não existe")
	}

	if !conta.Saque(valor, idTransacao) {
		return errors.New("transação já processada ou saldo insuficiente")
	}

	return nil
}

func (b *Banco) ConsultaSaldo(numero int) (float64, error) {
	b.mutex.RLock()
	conta, exists := b.Contas[numero]
	b.mutex.RUnlock()

	if !exists {
		return 0, errors.New("conta não existe")
	}

	return conta.ConsultaSaldo(), nil
}
