package cmd

import (
	"custom-blockchain-iot-protocol/config"
	"custom-blockchain-iot-protocol/internal/server"
	"fmt"
	"log"
	_ "math/big"
)

func Simulate() {
	conf, err := config.NewConfig()
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		return
	}

	srv := server.NewServer(conf)

	err = srv.Start()
	if err != nil {
		log.Panicln(err)
	}
}
