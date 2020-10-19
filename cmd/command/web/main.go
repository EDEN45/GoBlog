package main

import (
	"fmt"
	"github.com/EDEN45/GoBlog/cmd/blueprint/config"
	"github.com/EDEN45/GoBlog/cmd/blueprint/routes"
)

func main() {
	routes.Run()

	// load application configurations
	if err := config.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("invalid application configuration: %s", err))
	}

	fmt.Println(config.Config.ConfigVar)
}
