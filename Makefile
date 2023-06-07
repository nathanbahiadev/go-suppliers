init:
	@echo "***** Inicializando projeto *****"
	go mod tidy

run:
	@echo "***** Executando projeto *****"
	go run src/cmd/main.go

build:
	@echo "***** Executando o build *****"
	go build src/cmd/main.go

tests:
	@echo "***** Executando testes unitários e de integração *****"
	go test -coverprofile=coverage.out ./...

coverage:
	@echo "***** Executando testes e gerando relatório de cobertura *****"
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
