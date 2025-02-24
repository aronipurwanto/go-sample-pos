package controllers

import (
	"github.com/aronipurwanto/go-sample-pos/models"
	"github.com/aronipurwanto/go-sample-pos/services"
	"github.com/gofiber/fiber/v2"
)

type EmployeeController struct {
	employeeService services.EmployeeService
}

func NewEmployeeController(service services.EmployeeService) *EmployeeController {
	return &EmployeeController{service}
}

func (c *EmployeeController) CreateEmployee(ctx *fiber.Ctx) error {
	var Employee models.Employee
	if err := ctx.BodyParser(&Employee); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := c.employeeService.CreateEmployee(&Employee); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(201).JSON(Employee)
}

func (c *EmployeeController) GetAllEmployees(ctx *fiber.Ctx) error {
	Employees, err := c.employeeService.GetAllEmployee()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(Employees)
}

func (c *EmployeeController) GetEmployeeByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	Employee, err := c.employeeService.GetEmployeeByID(id)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": "Employee not found"})
	}
	return ctx.JSON(Employee)
}

func (c *EmployeeController) UpdateEmployee(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var Employee models.Employee
	if err := ctx.BodyParser(&Employee); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := c.employeeService.UpdateEmployee(id, &Employee); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(Employee)
}

func (c *EmployeeController) DeleteEmployee(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if err := c.employeeService.DeleteEmployee(id); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(204).Send(nil)
}
