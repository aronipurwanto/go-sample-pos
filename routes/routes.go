package routes

import (
	"github.com/aronipurwanto/go-sample-pos/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, productController *controllers.ProductController,
	customerController *controllers.CustomerController,
	employeeController *controllers.EmployeeController) {
	// Product Routes
	productGroup := app.Group("/products")
	productGroup.Post("/", productController.CreateProduct)
	productGroup.Post("/bulk", productController.CreateProductsBulk)
	productGroup.Get("/", productController.GetAllProducts)
	productGroup.Get("/:id", productController.GetProductByID)
	productGroup.Put("/:id", productController.UpdateProduct)
	productGroup.Delete("/:id", productController.DeleteProduct)

	// Customer Routes
	customerGroup := app.Group("/customers")
	customerGroup.Post("/", customerController.CreateCustomer)
	customerGroup.Get("/", customerController.GetAllCustomers)
	customerGroup.Get("/:id", customerController.GetCustomerByID)
	customerGroup.Put("/:id", customerController.UpdateCustomer)
	customerGroup.Delete("/:id", customerController.DeleteCustomer)

	// Customer Routes
	employeeGroup := app.Group("/employees")
	employeeGroup.Post("/", employeeController.CreateEmployee)
	employeeGroup.Get("/", employeeController.DeleteEmployee)
	employeeGroup.Get("/:id", employeeController.GetEmployeeByID)
	employeeGroup.Put("/:id", employeeController.UpdateEmployee)
	employeeGroup.Delete("/:id", employeeController.DeleteEmployee)
}
