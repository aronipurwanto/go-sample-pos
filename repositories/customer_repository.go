package repositories

import (
	"github.com/aronipurwanto/go-sample-pos/models"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	Create(customer *models.Customer) error
	GetAll() ([]models.Customer, error)
	GetByID(id string) (*models.Customer, error)
	Update(id string, customer *models.Customer) error
	Delete(id string) error
}

type customerRepositoryImpl struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepositoryImpl{db}
}

func (r *customerRepositoryImpl) Create(customer *models.Customer) error {
	return r.db.Create(customer).Error
}

func (r *customerRepositoryImpl) GetAll() ([]models.Customer, error) {
	var customers []models.Customer
	err := r.db.Find(&customers).Error
	return customers, err
}

func (r *customerRepositoryImpl) GetByID(id string) (*models.Customer, error) {
	var customer models.Customer
	err := r.db.First(&customer, "customer_id = ?", id).Error
	return &customer, err
}

func (r *customerRepositoryImpl) Update(id string, customer *models.Customer) error {
	return r.db.Model(&models.Customer{}).Where("customer_id = ?", id).Updates(customer).Error
}

func (r *customerRepositoryImpl) Delete(id string) error {
	return r.db.Delete(&models.Customer{}, "customer_id = ?", id).Error
}
