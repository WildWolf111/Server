package main

import (
	"StandartWebServer/internal/app/api"
	"flag"
	"log"

	"github.com/BurntSushi/toml"
)

var (
	configPath string = "configs/app/api.toml"
)

func init() {
	flag.StringVar(&configPath, "path", "configs/api.toml", "path to config file to .toml format")
}

func main() {
	flag.Parse()
	log.Println("It works")
	// servver instance initialisation
	config := api.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	server := api.New(config)
	if err != nil {
		log.Fatal(err)
	}

	//api server start
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
