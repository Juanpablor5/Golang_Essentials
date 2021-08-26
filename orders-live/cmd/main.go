package main

import (
	"log"

	"github.com/joho/godotenv"
	"leal.co/orders/internal/pkg/database"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	err := godotenv.Load()
	if err != nil {
		log.Printf("error loading .env, %v", err)
	}
}

func main() {
	defer database.GetDB().Close()

	log.Println("ehlo world")
}
