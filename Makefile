all: build

build:
	go build -o bin/administracao cmd/administracao/main.go
	go build -o bin/agencia cmd/agencia/main.go
	go build -o bin/caixa_automatico cmd/caixa_automatico/main.go
	
run-admin:
	./bin/administracao

run-agencia:
	./bin/agencia

run-caixa:
	./bin/caixa_automatico

clean:
	rm -rf bin/