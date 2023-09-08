package main

import (
	"fmt"
	"net/http"
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

	r.NoRoute(func(c *gin.Context) {
		errMsg := fmt.Sprintf("Not found: %s %s", c.Request.Method, c.Request.URL.Path)
		c.JSON(http.StatusNotFound, gin.H{"message": errMsg, "success": false})
		c.Abort()
	})

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
