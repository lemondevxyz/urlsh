package main

import (
	"fmt"
	"log"

	"github.com/toms1441/urlsh/internal/config"
)

func main() {
	c, err := config.NewConfig()
	if err != nil {
		log.Fatalf("config.NewConfig: %v", err)
	}

	fmt.Println(c)
}
