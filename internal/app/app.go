package app

import (
	"github.com/sunnygosdk/go-chi-fullcycle-api/configs"
)

func StartServer() {
	db, err := configs.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
