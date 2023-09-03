package emails

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestEmailToMe(t *testing.T) {
	if goenv := os.Getenv("GOENV"); goenv != "production" && goenv != "testing" {
		if err := godotenv.Load("../../.env"); err != nil {
			t.Error(err)
			return
		}
	}

	// if err := sendEmailToMe(
	// 	"example@example.com",
	// 	"Example",
	// 	"I want to have a chat with you",
	// 	"Hi there, how are you?\nAre you available today for a quick call? I want to discuss about a project with you.\nRegards,\nExample",
	// ); err != nil {
	// 	t.Error(err)
	// }
}
