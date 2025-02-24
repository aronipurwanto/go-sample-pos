package controllers

import (
	"github.com/aronipurwanto/go-sample-pos/models"
	"github.com/aronipurwanto/go-sample-pos/services"
	"github.com/gofiber/fiber/v2"
)

type ProductController struct {
	productService services.ProductService
}

func NewProductController(service services.ProductService) *ProductController {
	return &ProductController{service}
}

func (c *ProductController) CreateProduct(ctx *fiber.Ctx) error {
	var product models.Product
	if err := ctx.BodyParser(&product); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := c.productService.CreateProduct(&product); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(201).JSON(product)
}

func (c *ProductController) CreateProductV2(ctx *fiber.Ctx) error {
	var product []models.Product
	if err := ctx.BodyParser(&product); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	type ProductError struct {
		Product models.Product `json:"product"`
		Error   error          `json:"error"`
	}

	var productErrors []ProductError

	for _, p := range product {
		if err := c.productService.CreateProduct(&p); err != nil {
			productErrors = append(productErrors, ProductError{
				Product: p,
				Error:   err,
			})
		}
	}

	if len(productErrors) > 0 {
		return ctx.Status(500).JSON(fiber.Map{"error": productErrors})
	}
	return ctx.Status(201).JSON(product)
}

// Bulk insert products
func (c *ProductController) CreateProductsBulk(ctx *fiber.Ctx) error {
	var products []models.Product

	// Parse JSON request
	if err := ctx.BodyParser(&products); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON format",
		})
	}

	// Insert products
	if err := c.productService.CreateProductBulk(&products); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to insert products",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Products added successfully",
	})
}

func (c *ProductController) GetAllProducts(ctx *fiber.Ctx) error {
	products, err := c.productService.GetAllProducts()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(products)
}

func (c *ProductController) GetProductByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	product, err := c.productService.GetProductByID(id)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": "Product not found"})
	}
	return ctx.JSON(product)
}

func (c *ProductController) UpdateProduct(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var product models.Product
	if err := ctx.BodyParser(&product); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := c.productService.UpdateProduct(id, &product); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(product)
}

func (c *ProductController) DeleteProduct(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if err := c.productService.DeleteProduct(id); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(204).Send(nil)
}
