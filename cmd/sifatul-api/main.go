package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/sifatulrabbi/sifatul-api/internals/blogs"
	"github.com/sifatulrabbi/sifatul-api/internals/controllers/emails"
)

var (
	GOENV = os.Getenv("GOENV")
	PORT  string
)

func main() {
	prepareENV()
	r := setupRouter()
	v1 := r.Group("/v1")
	v1.POST("/emails/to-me", emails.HandleEmailToMe)
	blogs.RegisterBlogRoutes(v1)

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

func setupRouter() *gin.Engine {
	r := gin.Default()

	corsConfig := cors.DefaultConfig()
	// corsConfig.AllowOrigins = []string{"https://sifatul.com", "http://localhost:3000", "https://www.sifatul.com"}
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	corsConfig.AllowHeaders = []string{
		"Accept",
		"Accept-Encoding",
		"Accept-Language",
		"Access-Control-Request-Headers",
		"Access-Control-Request-Method",
		"Authorization",
		"Connection",
		"Content-Type",
		"Cookie",
		"Date",
		"If-Modified-Since",
		"If-None-Match",
		"Origin",
		"Referrer",
		"User-Agent",
		"X-Requested-With",
	}
	r.Use(cors.New(corsConfig))

	return r
}
