# Makefile
APPNAME =`basename ${PWD}`

-include .env

.PHONY: build run test protos
build:
	@go build
run: build
	MYSQL_HOST=$(MYSQL_HOST) \
	MYSQL_PORT=${MYSQL_PORT} \
	MYSQL_DB=$(MYSQL_DB) \
	MYSQL_USERNAME=$(MYSQL_USERNAME) \
	MYSQL_PASSWORD=$(MYSQL_PASSWORD) \
	./$(APPNAME)
ent-schema-update:
	@entc generate ./ent/schema
ent-cli-update:
	@go get github.com/facebookincubator/ent/cmd/entc
dep:
	@go mod tidy
	@go mod vendor