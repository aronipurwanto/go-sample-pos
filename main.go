package main

import (
	"fmt"
	"github.com/aronipurwanto/go-sample-pos/controllers"
	"github.com/aronipurwanto/go-sample-pos/database"
	"github.com/aronipurwanto/go-sample-pos/repositories"
	"github.com/aronipurwanto/go-sample-pos/routes"
	"github.com/aronipurwanto/go-sample-pos/services"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	// Initialize Fiber app
	app := fiber.New()

	// Connect to database
	database.ConnectDatabase()

	// Setup Dependencies
	productController, customerController, employeeController := setupDependencies()

	// Setup Routes
	route := routes.NewRouteSetup(app, productController, customerController, employeeController)
	route.Setup()

	// Start server
	port := ":8080" // Bisa dipindahkan ke config jika ingin lebih fleksibel
	fmt.Printf("🚀 Server running on http://localhost%s\n", port)

	if err := app.Listen(port); err != nil {
		log.Fatal(err) // Menggunakan log.Fatal agar program langsung berhenti jika error
	}
}

func setupDependencies() (*controllers.ProductController, *controllers.CustomerController, *controllers.EmployeeController) {
	// Product
	productRepo := repositories.NewProductRepository(database.DB)
	productService := services.NewProductService(productRepo)
	productController := controllers.NewProductController(productService)

	// Customer
	customerRepo := repositories.NewCustomerRepository(database.DB)
	customerService := services.NewCustomerService(customerRepo)
	customerController := controllers.NewCustomerController(customerService)

	// Employee
	employeeRepo := repositories.NewEmployeeRepository(database.DB)
	employeeService := services.NewEmployeeService(employeeRepo)
	employeeController := controllers.NewEmployeeController(employeeService)

	return productController, customerController, employeeController
}
