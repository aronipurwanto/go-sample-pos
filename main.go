package main

import (
	"fmt"
	"github.com/aronipurwanto/go-sample-pos/controllers"
	"github.com/aronipurwanto/go-sample-pos/database"
	"github.com/aronipurwanto/go-sample-pos/repositories"
	"github.com/aronipurwanto/go-sample-pos/routes"
	"github.com/aronipurwanto/go-sample-pos/services"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Connect to database
	database.ConnectDatabase()

	// Dependency Injection
	productRepo := repositories.NewProductRepository(database.DB)
	productService := services.NewProductService(productRepo)
	productController := controllers.NewProductController(productService)

	// Customer
	customerRepo := repositories.NewCustomerRepository(database.DB)
	customerService := services.NewCustomerService(customerRepo)
	customerController := controllers.NewCustomerController(customerService)

	//Employee
	employeeRepo := repositories.NewEmployeeRepository(database.DB)
	employeeService := services.NewEmployeeService(employeeRepo)
	employeeController := controllers.NewEmployeeController(employeeService)

	// Setup Routes
	routes.SetupRoutes(app, productController, customerController, employeeController)

	// Start server
	err := app.Listen(":8080")
	if err != nil {
		fmt.Println(err)
	}
}
