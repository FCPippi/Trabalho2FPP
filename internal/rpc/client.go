package rpc

import (
	"errors"
	"net/rpc"
	"time"
)

type BancoClient struct {
	client *rpc.Client
}

func NewBancoClient(address string) (*BancoClient, error) {
	client, err := rpc.Dial("tcp", address)
	if err != nil {
		return nil, err
	}
	return &BancoClient{client: client}, nil
}

func (c *BancoClient) AbrirConta(numero int) (bool, error) {
	args := &AbrirContaArgs{Numero: numero}
	var reply bool
	err := c.client.Call("BancoServer.AbrirConta", args, &reply)
	return reply, err
}

func (c *BancoClient) FecharConta(numero int) (bool, error) {
	var reply bool
	err := c.client.Call("BancoServer.FecharConta", numero, &reply)
	return reply, err
}

func (c *BancoClient) Deposito(numero int, valor float64, idTransacao string) (bool, error) {
	args := &OperacaoArgs{Numero: numero, Valor: valor, IdTransacao: idTransacao}
	var reply bool
	err := c.client.Call("BancoServer.Deposito", args, &reply)
	return reply, err
}

func (c *BancoClient) Saque(numero int, valor float64, idTransacao string) (bool, error) {
	args := &OperacaoArgs{Numero: numero, Valor: valor, IdTransacao: idTransacao}
	var reply bool
	err := c.client.Call("BancoServer.Saque", args, &reply)
	return reply, err
}

func (c *BancoClient) ConsultaSaldo(numero int) (float64, error) {
	var saldo float64
	err := c.client.Call("BancoServer.ConsultaSaldo", numero, &saldo)
	return saldo, err
}

func (c *BancoClient) Close() error {
	return c.client.Close()
}

func (c *BancoClient) RetryOperation(operation func() (bool, error), maxRetries int) (bool, error) {
	backoff := time.Second
	for i := 0; i < maxRetries; i++ {
		success, err := operation()
		if err == nil {
			return success, nil
		}

		time.Sleep(backoff)
		backoff *= 2
	}
	return false, errors.New("máximo de tentativas excedido")
}
