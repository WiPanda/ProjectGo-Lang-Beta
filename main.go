package main

import (
	"Project3/config"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error load env, use global env")
	}

	config.OpenDB()

	// close
	db, err := config.Mysql.DB.DB()
	if err != nil {
		log.Fatal("Gagal mendapatkan koneksi database")
	}
	defer db.Close()

	app := fiber.New()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fmt.Println("Server running on port: ", port)
	log.Fatal(app.Listen(":" + port))
}
