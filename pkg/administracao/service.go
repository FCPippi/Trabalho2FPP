package administracao

type Service struct {
	contas map[string]*Conta
}

func NewService() *Service {
	return &Service{contas: make(map[string]*Conta)}
}

func (s *Service) AbrirConta(numero string, reply *bool) error {
	if _, existe := s.contas[numero]; existe {
		*reply = false
		return nil
	}
	s.contas[numero] = &Conta{Numero: numero, Saldo: 0}
	*reply = true
	return nil
}

func (s *Service) FecharConta(numero string, reply *bool) error {
	if _, existe := s.contas[numero]; !existe {
		*reply = false
		return nil
	}
	delete(s.contas, numero)
	*reply = true
	return nil
}

func (s *Service) Depositar(args [2]interface{}, reply *bool) error {
	numero := args[0].(string)
	valor := args[1].(float64)
	if conta, existe := s.contas[numero]; existe {
		conta.Saldo += valor
		*reply = true
		return nil
	}
	*reply = false
	return nil
}

func (s *Service) Sacar(args [2]interface{}, reply *bool) error {
	numero := args[0].(string)
	valor := args[1].(float64)
	if conta, existe := s.contas[numero]; existe && conta.Saldo >= valor {
		conta.Saldo -= valor
		*reply = true
		return nil
	}
	*reply = false
	return nil
}

func (s *Service) ConsultarSaldo(numero string, saldo *float64) error {
	if conta, existe := s.contas[numero]; existe {
		*saldo = conta.Saldo
		return nil
	}
	*saldo = 0
	return nil
}
