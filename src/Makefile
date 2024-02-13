build: build-x64
	go build -o dist/rinha2 cmd/main.go # GOOS=darwin GOARCH=arm64 

build-x64:
	mkdir -p dist
	GOOS=linux GOARCH=amd64 go build -o dist/rinha2-x64 -ldflags="-s -w" cmd/main.go

run:
	go run ./cmd/main.go --debug

transacoes:
	for  i in `seq 1 40`; do \
		curl -X POST \
			http://localhost:9999/clientes/1/transacoes \
			-H 'Content-Type: application/json' \
			-d '{ "valor": 1, "tipo": "c", "descricao": "teste local" }' ;\
	done

extrato:
	for  i in `seq 1 40`; do \
		curl -X GET \
			http://localhost:9999/clientes/1/extrato ;\
	done

.PHONY: build build-x64 transacoes extrato
