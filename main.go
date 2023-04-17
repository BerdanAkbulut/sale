package main

import (
	"github.com/BerdanAkbulut/sale/models"
	"github.com/BerdanAkbulut/sale/routes"
	"log"
)

func main() {
	models.SetupDatabase()
	r := routes.SetupRouter()

	log.Println("Server running")
	r.Run()
}
