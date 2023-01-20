package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/shashank/ecommerce-yt/controllers"
	"github.com/shashank/ecommerce-yt/database"
	"github.com/shashank/ecommerce-yt/middleware"
	"github.com/shashank/ecommerce-yt/routes"
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
