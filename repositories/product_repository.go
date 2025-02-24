package repositories

import (
	"github.com/aronipurwanto/go-sample-pos/models"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *models.Product) error
	CreateBulk(product *[]models.Product) error
	GetAll() ([]models.Product, error)
	GetByID(id string) (*models.Product, error)
	Update(id string, product *models.Product) error
	Delete(id string) error
}

type productRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepositoryImpl{db}
}

func (r *productRepositoryImpl) Create(product *models.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepositoryImpl) CreateBulk(products *[]models.Product) error {
	return r.db.Create(products).Error
}

func (r *productRepositoryImpl) GetAll() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Find(&products).Error
	return products, err
}

func (r *productRepositoryImpl) GetByID(id string) (*models.Product, error) {
	var product models.Product
	err := r.db.First(&product, "product_id = ?", id).Error
	return &product, err
}

func (r *productRepositoryImpl) Update(id string, product *models.Product) error {
	return r.db.Model(&models.Product{}).Where("product_id = ?", id).Updates(product).Error
}

func (r *productRepositoryImpl) Delete(id string) error {
	return r.db.Delete(&models.Product{}, "product_id = ?", id).Error
}
