.PHONY: run, run-prod

run:
	PORT=9876 GOENV=development go run ./cmd/sifatul-api/main.go

run-prod:
	PORT=9876 GOENV=production GIN_MODE=release go run ./cmd/sifatul-api/main.go
