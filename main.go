package main

import (
	"log"
	"net/http"

	"github.com/sifatulrabbi/sifatul-api/configs"
	"github.com/sifatulrabbi/sifatul-api/mailer"

	"github.com/gin-gonic/gin"
)

func main() {
	configs.LoadConfigs()
	config := configs.GetConfigs()
	router := gin.Default()

	router.GET("/api/v1", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})
	// register routers
	mailer.RegisterRouter(router)

	if err := http.ListenAndServe(":"+config.PORT, router); err != nil {
		log.Fatal(err.Error())
	}
}
