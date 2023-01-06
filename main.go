package main

import (
	"log"
	"os"

	"github.com/AlghazHernanda/ecommerce-go/controllers"
	"github.com/AlghazHernanda/ecommerce-go/database"
	"github.com/AlghazHernanda/ecommerce-go/middleware"
	"github.com/AlghazHernanda/ecommerce-go/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductsData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRouters(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))
}
