package main

import (
	"fmt"
	"log"
	"server/db"
)

func main() {
	_, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("could not initialize database conntection: %v", err)
	}
	fmt.Println("Я живу!")
}
