package routes

import (
	"github.com/aronipurwanto/go-sample-pos/controllers"
	"github.com/gofiber/fiber/v2"
)

type RouteSetup struct {
	App                *fiber.App
	ProductController  *controllers.ProductController
	CustomerController *controllers.CustomerController
	EmployeeController *controllers.EmployeeController
}

func NewRouteSetup(app *fiber.App, productController *controllers.ProductController,
	customerController *controllers.CustomerController,
	employeeController *controllers.EmployeeController) *RouteSetup {
	return &RouteSetup{
		App:                app,
		ProductController:  productController,
		CustomerController: customerController,
		EmployeeController: employeeController,
	}
}

func (r *RouteSetup) Setup() {
	// Product Routes
	productGroup := r.App.Group("/products")
	{
		productGroup.Post("/", r.ProductController.CreateProduct)
		productGroup.Post("/bulk", r.ProductController.CreateProductsBulk)
		productGroup.Get("/", r.ProductController.GetAllProducts)
		productGroup.Get("/:id", r.ProductController.GetProductByID)
		productGroup.Put("/:id", r.ProductController.UpdateProduct)
		productGroup.Delete("/:id", r.ProductController.DeleteProduct)
	}

	// Customer Routes
	customerGroup := r.App.Group("/customers")
	{
		customerGroup.Post("/", r.CustomerController.CreateCustomer)
		customerGroup.Get("/", r.CustomerController.GetAllCustomers)
		customerGroup.Get("/:id", r.CustomerController.GetCustomerByID)
		customerGroup.Put("/:id", r.CustomerController.UpdateCustomer)
		customerGroup.Delete("/:id", r.CustomerController.DeleteCustomer)
	}

	// Employee Routes
	employeeGroup := r.App.Group("/employees")
	{
		employeeGroup.Post("/", r.EmployeeController.CreateEmployee)
		employeeGroup.Get("/", r.EmployeeController.GetAllEmployees) // FIXED
		employeeGroup.Get("/:id", r.EmployeeController.GetEmployeeByID)
		employeeGroup.Put("/:id", r.EmployeeController.UpdateEmployee)
		employeeGroup.Delete("/:id", r.EmployeeController.DeleteEmployee)
	}
}
