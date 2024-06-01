.PHONY: default run docs

default: run-with-docs

run: 
	@go run main.go

run-with-docs:
	@swag init
	@go run main.go

docs: 
	@swag init