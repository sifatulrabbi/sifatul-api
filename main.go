package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sifatulrabbi/sifatul-api/configs"

	"github.com/gin-contrib/cors"
)

func main() {
	configs.LoadConfigs()
	config := configs.GetConfigs()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true

	fmt.Println("Server is starting on port 8000")
	if err := http.ListenAndServe(":"+config.PORT, nil); err != nil {
		log.Fatal(err.Error())
	}
}
