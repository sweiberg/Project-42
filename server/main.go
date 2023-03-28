package main

import (
	"exert-shop/controller"
	"exert-shop/db"
	"exert-shop/middleware"
	"exert-shop/model"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDB()
	loadRoutes()
}

func loadDB() {
	db.Connect()
	db.Database.AutoMigrate(&model.User{})
	db.Database.AutoMigrate(&model.Product{})
	db.Database.AutoMigrate(&model.Category{})
}

func loadEnv() {
	err := godotenv.Load(".env.local")

	if err != nil {
		log.Fatal("Error! The .env file could not be loaded.")
	}
}

func loadRoutes() {
	router := gin.Default()
	router.Use(middleware.CORS())

	authAPI := router.Group("/auth")
	authAPI.POST("/register", controller.Register)
	authAPI.POST("/login", controller.Login)

	publicAPI := router.Group("/api")
	publicAPI.GET("/product/:id", controller.ViewProduct)
	publicAPI.GET("/category/:id", controller.ViewCategory)
	publicAPI.GET("/profile/:id", controller.ViewProfile)

	protectedAPI := router.Group("/api")
	protectedAPI.Use(middleware.VerifyJWT())
	protectedAPI.POST("/addproduct", controller.AddProduct)
	protectedAPI.POST("/addcategory", controller.AddCategory)

	router.Run(":4300")

	fmt.Println("Listen server successfully started on port 4300.")
}
