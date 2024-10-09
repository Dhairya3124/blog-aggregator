package main

import (
	"fmt"
	"log"

	config "github.com/Dhairya3124/blog-aggregator/internal/config"
)

func main() {
	configDetails, err := config.Read()
	if err != nil {
		log.Fatalf(" Following Error %v", err)
	}
	configDetails.SetUser("Dhairya")
	configDetails, err = config.Read()
	if err != nil {
		log.Fatalf(" Following Error %v", err)
	}
	fmt.Println(configDetails)

}
