package main

import (
	"log"
	"net/http"

	"github.com/sifatulrabbi/sifatul-api/configs"
	gql "github.com/sifatulrabbi/sifatul-api/graph_practice"
	"github.com/sifatulrabbi/sifatul-api/mailer"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	configs.LoadConfigs()
	config := configs.GetConfigs()
	router := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	// cors middleware
	router.Use(cors.New(corsConfig))
	router.GET("/api/v1", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})
	// register routers
	mailer.RegisterRouter(router)
	// graphql handler
	http.Handle("/graphql", gql.GqHandler)

	if err := http.ListenAndServe(":"+config.PORT, router); err != nil {
		log.Fatal(err.Error())
	}
}
