package emails

import (
	"errors"
	"fmt"
	"net/http"
	"net/smtp"
	"os"

	"github.com/gin-gonic/gin"
)

type SendEmailBody struct {
	SenderEmail string `json:"sender_email"`
	SenderName  string `json:"sender_name"`
	Subject     string `json:"subject"`
	Message     string `json:"message"`
}

const (
	receiver = "mdsifatulislam.rabbi@gmail.com"
)

func HandleEmailTome(c *gin.Context) {
	defer c.Abort()

	p := SendEmailBody{}
	if err := c.BindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "success": false})
		return
	}

	if p.SenderEmail == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "required field not 'sender_email' found", "success": false})
		return
	}
	if p.SenderName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "required field not 'sender_name' found", "success": false})
		return
	}
	if p.Subject == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "required field not 'subject' found", "success": false})
		return
	}
	if p.Message == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "required field not 'message' found", "success": false})
		return
	}
	if err := sendEmailToMe(p.SenderEmail, p.SenderName, p.Subject, p.Message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "success": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Email sent successfully!", "success": true})
}

func sendEmailToMe(senderEmail, senderName, senderSubject, senderMsg string) error {
	var (
		// SMTP_PORT_SSL  = os.Getenv("SMTP_PORT_SSL")
		SMTP_PORT_TLS  = os.Getenv("SMTP_PORT_TLS")
		SMTP_HOST      = os.Getenv("SMTP_HOST")
		EMAIL_ACCOUNT  = os.Getenv("EMAIL_ACCOUNT")
		EMAIL_PASSWORD = os.Getenv("EMAIL_PASSWORD")
		smtpHost       = SMTP_HOST + ":" + SMTP_PORT_TLS
		to             = []string{"mdsifatulislam.rabbi@gmail.com"}
		body           []byte
		auth           smtp.Auth
	)

	if SMTP_HOST == "" {
		return errors.New("smtp_host env not found")
	}
	if SMTP_PORT_TLS == "" {
		return errors.New("smtp_port_ssl env not found")
	}
	if EMAIL_ACCOUNT == "" {
		return errors.New("email_account env not found")
	}
	if EMAIL_PASSWORD == "" {
		return errors.New("email_password env not found")
	}

	body = []byte(fmt.Sprintf("From: %s <%s>\r\n", senderName, senderEmail) +
		fmt.Sprintf("Subject: Portfolio | %s\r\n", senderSubject) +
		fmt.Sprintf("To: %s\r\n", receiver) +
		senderMsg +
		fmt.Sprintf("Reply to - %s\r\n", senderEmail),
	)

	auth = smtp.PlainAuth("", EMAIL_ACCOUNT, EMAIL_PASSWORD, SMTP_HOST)
	err := smtp.SendMail(smtpHost, auth, EMAIL_ACCOUNT, to, body)
	return err
}
