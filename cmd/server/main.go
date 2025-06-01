package main

import (
	"fmt"

	"github.com/sunnygosdk/go-chi-fullcycle-api/configs"
)

func main() {
	config := configs.NewConfig()
	fmt.Println(config.GetEnvironment())
}
