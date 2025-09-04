package main

import (
	"log"
	"os"

	"github.com/federus1105/daysatu/internals/configs"
	"github.com/federus1105/daysatu/internals/routers"
	"github.com/joho/godotenv"
)

func main() {

	// inisialisai
	if err := godotenv.Load(); err != nil {
		log.Println("Failed to load env\nCause:", err.Error())
		return
	}
	log.Println(os.Getenv("DBUSER"))

	// inisialisasi db
	db, err := configs.InitDB()
	if err != nil {
		log.Println("Failed to connect to database\nCause: ", err.Error())
		return
	}

	defer db.Close()
	if err := configs.TestDB(db); err != nil {
		log.Println("Ping to DB failed\nCause:", err.Error())
		return
	}
	log.Println("DB Connected")

	router := routers.InitRouter(db)

	router.Run("localhost:8080")
}
