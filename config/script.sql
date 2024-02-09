CREATE TABLE clientes
	(
		id SMALLINT,
		limite BIGINT,
		saldo BIGINT,
		PRIMARY KEY (id)
	);

INSERT INTO clientes (id, limite, saldo)
VALUES (1,100000,0),
       (2,80000,0),
       (3,1000000,0),
       (4,10000000,0),
       (5,500000,0);

CREATE TABLE transacoes
	(
		id SERIAL PRIMARY KEY,
		cliente_id SMALLINT,
		valor BIGINT,
		tipo CHAR(1),
		descricao TEXT,
		realizada_em TIMESTAMP,
		FOREIGN KEY (cliente_id) REFERENCES clientes(id)

	);
CREATE INDEX realizada_em_idx ON transacoes (realizada_em DESC);

