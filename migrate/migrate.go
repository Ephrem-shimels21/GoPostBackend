package main

import (
	"example.com/gocrud/Initializers"
	"example.com/gocrud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
