package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sifatulrabbi/sifatul-api/controllers/emails"
)

var (
	GOENV = os.Getenv("GOENV")
	PORT  string
)

func main() {
	prepareENV()
	r := gin.Default()
	v1 := r.Group("/v1")
	v1.POST("/emails/to-me", emails.HandleEmailTome)

	if err := r.Run(":" + PORT); err != nil {
		panic(err)
	}
}

func prepareENV() {
	if GOENV != "production" {
		if err := godotenv.Load(); err != nil {
			panic(err)
		}
	}
	PORT = os.Getenv("PORT")
}
