
build:
	docker build -t k8s-sample-go:v0.0.1 .

server:
	go run cmd/main.go

air:
	mysql.server start; air -c .air.toml

test:
	echo Hello World!!


.PHONY: server
