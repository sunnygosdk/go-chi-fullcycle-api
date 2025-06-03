package main

import (
	"log"

	"github.com/sunnygosdk/go-chi-fullcycle-api/configs"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/app"
)

func main() {
	log.Println("FullCycle API Starting...")
	config := configs.NewConfig()
	db, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	port := config.GetWebServerPort()
	app.StartServer(port, db)
}
