package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	sifatulapi "github.com/sifatulrabbi/sifatul-api"
)

func main() {
	runCtx := context.Background()
	runCtx = prepareENV(runCtx)
	runCtx = prepareDB(runCtx)
	if err := sifatulapi.StartAPI(runCtx); err != nil {
		panic(err)
	}
}

func prepareENV(ctx context.Context) context.Context {
	fmt.Println("Preparing ENV vars...")
	var (
		GOENV = os.Getenv("GOENV")
		PORT  string
	)
	if GOENV != "production" {
		if err := godotenv.Load(".env"); err != nil {
			log.Panicln("No .env file found", err)
		}
	}
	PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "9876"
	}
	ctx = context.WithValue(ctx, sifatulapi.GOENV, GOENV)
	ctx = context.WithValue(ctx, sifatulapi.PORT, PORT)
	return ctx
}

func prepareDB(ctx context.Context) context.Context {
	fmt.Println("Preparing DB connection...")
	db, err := gorm.Open(sqlite.Open("sifatulapi.db"), &gorm.Config{})
	if err != nil {
		log.Panicln("Failed to connect with sqlite database:", err)
	}
	ctx = context.WithValue(ctx, sifatulapi.DB_CONN, db)
	return ctx
}
