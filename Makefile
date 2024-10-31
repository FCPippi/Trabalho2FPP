# Nome do módulo Go
MODULE_NAME = sistema_bancario

# Diretórios dos comandos
SERVER_DIR = ./cmd/server
AGENCIA_DIR = ./cmd/agencia
CAIXA_DIR = ./cmd/caixa

# Binários
SERVER_BIN = server
AGENCIA_BIN = agencia
CAIXA_BIN = caixa

# Comandos Go
GO = go
GOBUILD = $(GO) build
GOTEST = $(GO) test
GORUN = $(GO) run

# Alvos
.PHONY: all clean test run-server run-agencia run-caixa

all: build

build: build-server build-agencia build-caixa

build-server:
	$(GOBUILD) -o $(SERVER_BIN) $(SERVER_DIR)

build-agencia:
	$(GOBUILD) -o $(AGENCIA_BIN) $(AGENCIA_DIR)

build-caixa:
	$(GOBUILD) -o $(CAIXA_BIN) $(CAIXA_DIR)

test:
	$(GOTEST) ./...

run-server: build-server
	./$(SERVER_BIN)

run-agencia: build-agencia
	./$(AGENCIA_BIN)

run-caixa: build-caixa
	./$(CAIXA_BIN)

clean:
	rm -f $(SERVER_BIN) $(AGENCIA_BIN) $(CAIXA_BIN)