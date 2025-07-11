# Makefile para filtro-pix

APP_NAME=filtro-pix
MAIN=main.go

INPUT?=entrada.log
OUTPUT?=pagamentos_pix.log
EXTRACT_OUTPUT?=txids.txt

.PHONY: build run clean filter extract version release snapshot

## Compila o binário localmente
build:
	go build -o $(APP_NAME) $(MAIN)

## Executa o CLI localmente
run:
	go run $(MAIN)

## Executa os testes
test:
	go test -v ./cmd/...

## Remove binários antigos
clean:
	rm -f $(APP_NAME)

## Filtra logs contendo "Pagamento Pix"
filter:
	go run $(MAIN) filter -i $(INPUT) -o $(OUTPUT)

## Extrai TXID : EndToEndId do arquivo filtrado
extract:
	go run $(MAIN) extract -i $(OUTPUT) -o $(EXTRACT_OUTPUT)

## Mostra a versão da ferramenta
version:
	go run $(MAIN) version

## Faz build de release real (requer GITHUB_TOKEN)
release:
	env $(cat .env | xargs) goreleaser release --clean

## Faz build snapshot (sem publicar no GitHub)
snapshot:
	env $(cat .env | xargs) goreleaser release --clean --snapshot
