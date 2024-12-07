package main

import (
	"marketplace-gin/database"
	"marketplace-gin/models"
)

func init() {
	database.ConnectDB()

}

func main() {
	database.DB.AutoMigrate(&models.User{})
}
