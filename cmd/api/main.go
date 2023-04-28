package main

import (
	"fmt"
	"note_extension/pkg/config"
)

func main() {
	fmt.Println("Hello world")
	mySqlConfig := config.MySQL{
		Host:     "localhost",
		Port:     3306,
		Username: "user",
		Password: "password",
		Database: "note",
	}
	fmt.Println(mySqlConfig.ConnectionString())

}
