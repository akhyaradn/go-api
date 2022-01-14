package main

import (
	"api/database"
	"api/router"
)

const PORT = ":8080"

func main() {
	database.StartDB()
	r := router.StartRouter()
	r.Run(PORT)
}
