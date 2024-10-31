package main

import (
	"TrabalhoDois/pkg/administracao"
	"log"
	"net"
	"net/rpc"
)

func main() {
	administracaoService := administracao.NewService()
	rpc.Register(administracaoService)

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Erro ao iniciar o servidor:", err)
	}
	defer listener.Close()

	log.Println("Servidor iniciado na porta 1234")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Erro ao aceitar conexão:", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}
