package main

import (
	"log"

	"github.com/arshabbir/brokermod/api"
	"github.com/arshabbir/brokermod/config"
)

func main() {
	log.Println("Starting broker server")
	cfg := &config.Config{}
	cfg.LoadConfig()

	s := api.NewServer(cfg)
	if err := s.Start(); err != nil {
		log.Fatal("error starting the server")
	}

}
