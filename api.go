package sifatulapi

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sifatulrabbi/sifatul-api/internals/blogs"
	"github.com/sifatulrabbi/sifatul-api/internals/controllers/emails"
)

func StartAPI(ctx context.Context) error {
	PORT, ok := ctx.Value(PORT).(string)
	if !ok {
		return fmt.Errorf("PORT not found in context")
	}

	r := setupRouter(ctx)
	v1 := r.Group("/api/v1")
	v1.POST("/emails/to-me", emails.HandleEmailToMe)
	blogs.RegisterBlogRoutes(v1)

	r.GET("/api/health", func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"message": "API is up and serving ginger bread.",
			"success": true,
		})
	})

	r.NoRoute(func(c *gin.Context) {
		errMsg := fmt.Sprintf("Not found: %s %s", c.Request.Method, c.Request.URL.Path)
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": errMsg, "success": false})
	})

	if err := r.Run(":" + PORT); err != nil {
		return err
	}
	return nil
}

func setupRouter(ctx context.Context) *gin.Engine {
	r := gin.Default()
	corsConfig := cors.DefaultConfig()
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
	if goEnv := ctx.Value(GOENV).(string); goEnv == "production" {
		corsConfig.AllowOrigins = []string{
			"https://sifatulrabbi.com",
			"https://www.sifatulrabbi.com",
			"https://api.sifatulrabbi.com",
		}
	} else {
		corsConfig.AllowOrigins = []string{"*"}
	}
	r.Use(cors.New(corsConfig))
	return r
}
