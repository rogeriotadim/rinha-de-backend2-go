build:
	mkdir -p dist
	GOOS=darwin GOARCH=arm64 go build -o dist/rinha2 cmd/main.go
	GOOS=linux GOARCH=amd64 go build -o dist/rinha2-x64 cmd/main.go

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

.PHONY: build transacoes extrato
