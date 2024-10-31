package banco

import (
	"sync"
)

type Conta struct {
	Numero     int
	Saldo      float64
	mutex      sync.Mutex
	Transacoes map[string]bool
}

func NovaConta(numero int) *Conta {
	return &Conta{
		Numero:     numero,
		Saldo:      0,
		Transacoes: make(map[string]bool),
	}
}

func (c *Conta) Deposito(valor float64, idTransacao string) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Transacoes[idTransacao] {
		return false
	}

	c.Saldo += valor
	c.Transacoes[idTransacao] = true
	return true
}

func (c *Conta) Saque(valor float64, idTransacao string) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Transacoes[idTransacao] {
		return false
	}

	if c.Saldo < valor {
		return false
	}

	c.Saldo -= valor
	c.Transacoes[idTransacao] = true
	return true
}

func (c *Conta) ConsultaSaldo() float64 {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.Saldo
}
