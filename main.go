package main

import "os"

// import (
// 	"github.com/AlghazHernanda/ecommerce-go/controllers"
// 	"github.com/AlghazHernanda/ecommerce-go/database"
// 	"github.com/AlghazHernanda/ecommerce-go/middleware"
// 	"github.com/AlghazHernanda/ecommerce-go/routes"
// 	"github.com/gin-gonic/gin"
// )

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductsData(database.Client, "Products"), database.UserData(database.Client, "Users"))
}
