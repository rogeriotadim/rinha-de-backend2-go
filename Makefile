SELECT="select * from clientes c join transacoes t on c.id=t.cliente_id where cliente_id=1;"
select:
	docker exec -it db2 psql -Uroot -drinha2 -c ${SELECT}

build:
	mkdir -p dist
	GOOS=darwin GOARCH=arm64 go build -o dist/rinha2 cmd/main.go
	GOOS=linux GOARCH=amd64 go build -o dist/rinha2-x64 main.go

build_x64:


.PHONY: select build build_x64
