package main

import (
	"github.com/joho/godotenv"
	"github.com/thgmagno/api-go/routes"
	"github.com/thgmagno/api-go/services"
)

func main() {
	_ = godotenv.Load()
	services.InitRedis()

	r := routes.SetupRouter()
	r.Run(":8080")
}
