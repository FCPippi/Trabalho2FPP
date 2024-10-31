package rpc

import (
	"TrabalhoDois/internal/banco"
	"net"
	"net/rpc"
)

type BancoServer struct {
	banco *banco.Banco
}

type AbrirContaArgs struct {
	Numero int
}

type OperacaoArgs struct {
	Numero      int
	Valor       float64
	IdTransacao string
}

func NewBancoServer() *BancoServer {
	return &BancoServer{
		banco: banco.NovoBanco(),
	}
}

func (s *BancoServer) AbrirConta(args *AbrirContaArgs, reply *bool) error {
	err := s.banco.AbrirConta(args.Numero)
	*reply = (err == nil)
	return err
}

func (s *BancoServer) FecharConta(numero *int, reply *bool) error {
	err := s.banco.FecharConta(*numero)
	*reply = (err == nil)
	return err
}

func (s *BancoServer) Deposito(args *OperacaoArgs, reply *bool) error {
	err := s.banco.Deposito(args.Numero, args.Valor, args.IdTransacao)
	*reply = (err == nil)
	return err
}

func (s *BancoServer) Saque(args *OperacaoArgs, reply *bool) error {
	err := s.banco.Saque(args.Numero, args.Valor, args.IdTransacao)
	*reply = (err == nil)
	return err
}

func (s *BancoServer) ConsultaSaldo(numero *int, saldo *float64) error {
	var err error
	*saldo, err = s.banco.ConsultaSaldo(*numero)
	return err
}

func StartServer(address string) error {
	server := NewBancoServer()
	rpc.Register(server)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		go rpc.ServeConn(conn)
	}
}
