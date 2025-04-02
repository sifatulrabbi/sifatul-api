package personal_finance

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterPersonalFinanceRouters(r *gin.RouterGroup) {
	r.GET("/personal-finance/records", getAllRecords)
	r.POST("/personal-finance/records", createNewRecord)
	r.GET("/personal-finance/records/:id", getArticleById)
	r.PATCH("/personal-finance/records/:id", updateArticleById)
	r.DELETE("/personal-finance/records/:id", deleteArticleById)
}

func getAllRecords(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"success": true,
		"message": "This route has not been initialized yet.",
	})
}

func createNewRecord(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"success": true,
		"message": "This route has not been initialized yet.",
	})
}

func getArticleById(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"success": true,
		"message": "This route has not been initialized yet.",
	})
}

func updateArticleById(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"success": true,
		"message": "This route has not been initialized yet.",
	})
}

func deleteArticleById(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"success": true,
		"message": "This route has not been initialized yet.",
	})
}
