package main

import (
	"fmt"
	"log"
	"note_extension/pkg/config"
)

func main() {
	fmt.Println("Hello world")

	mySQLConf, err := config.NewMySQLConfig()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(mySQLConf.ConnectionString())
}
