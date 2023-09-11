.PHONY: run

run:
	PORT=9876 GOENV=development go run main.go

run-prod:
	PORT=9876 GOENV=production GIN_MODE=release go run main.go
