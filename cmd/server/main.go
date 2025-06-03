package main

import (
	"log"

	"github.com/sunnygosdk/go-chi-fullcycle-api/config"
	"github.com/sunnygosdk/go-chi-fullcycle-api/database"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/app"
)

func main() {
	log.Println("FullCycle API Starting...")
	config := config.LoadConfig()
	db, err := database.ConnectDB(config)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	port := config.GetWebServerPort()
	app.StartServer(port, db)
}
