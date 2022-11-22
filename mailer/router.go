package mailer

import (
	"log"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(router *gin.Engine) {
	router.POST("/mailer", sendEmail)
}

func sendEmail(c *gin.Context) {
	var m mailer
	err := c.ShouldBind(&m)
	if err != nil {
		log.Println("Invalid payload")
		c.JSON(400, gin.H{
			"message": "Invalid request payload",
		})
		return
	}
	// send the email
	c.JSON(200, gin.H{
		"message": "Email sent successfully",
	})
}
