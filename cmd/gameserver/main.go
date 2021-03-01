package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/yarovlad/fatalforce/intrnals/app/gameserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/gameserver.toml", "path to config file")
}

func main() {

	flag.Parse()

	config := gameserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	s := gameserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("It Alive!!")
}
