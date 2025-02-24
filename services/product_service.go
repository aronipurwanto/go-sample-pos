package services

import (
	"github.com/aronipurwanto/go-sample-pos/models"
	"github.com/aronipurwanto/go-sample-pos/repositories"
)

type ProductService interface {
	CreateProduct(product *models.Product) error
	CreateProductBulk(product *[]models.Product) error
	GetAllProducts() ([]models.Product, error)
	GetProductByID(id string) (*models.Product, error)
	UpdateProduct(id string, product *models.Product) error
	DeleteProduct(id string) error
}

type productServiceImpl struct {
	Repo repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) ProductService {
	return &productServiceImpl{repo}
}
func (s *productServiceImpl) CreateProduct(product *models.Product) error {
	return s.Repo.Create(product)
}

func (s *productServiceImpl) CreateProductBulk(products *[]models.Product) error {
	return s.Repo.CreateBulk(products)
}

func (s *productServiceImpl) GetAllProducts() ([]models.Product, error) {
	return s.Repo.GetAll()
}

func (s *productServiceImpl) GetProductByID(id string) (*models.Product, error) {
	return s.Repo.GetByID(id)
}

func (s *productServiceImpl) UpdateProduct(id string, product *models.Product) error {
	return s.Repo.Update(id, product)
}

func (s *productServiceImpl) DeleteProduct(id string) error {
	return s.Repo.Delete(id)
}
