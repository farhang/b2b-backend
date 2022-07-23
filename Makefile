include .env.dev
export

run:
	go run main.go
swagger:
	swag init -ot go
