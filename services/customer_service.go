package services

import (
	"github.com/aronipurwanto/go-sample-pos/models"
	"github.com/aronipurwanto/go-sample-pos/repositories"
)

type CustomerService interface {
	CreateCustomer(customer *models.Customer) error
	GetAllCustomers() ([]models.Customer, error)
	GetCustomerByID(id string) (*models.Customer, error)
	UpdateCustomer(id string, customer *models.Customer) error
	DeleteCustomer(id string) error
}

type customerServiceImpl struct {
	customerRepo repositories.CustomerRepository
}

func NewCustomerService(repo repositories.CustomerRepository) CustomerService {
	return &customerServiceImpl{repo}
}

func (s *customerServiceImpl) CreateCustomer(customer *models.Customer) error {
	return s.customerRepo.Create(customer)
}

func (s *customerServiceImpl) GetAllCustomers() ([]models.Customer, error) {
	return s.customerRepo.GetAll()
}

func (s *customerServiceImpl) GetCustomerByID(id string) (*models.Customer, error) {
	return s.customerRepo.GetByID(id)
}

func (s *customerServiceImpl) UpdateCustomer(id string, customer *models.Customer) error {
	return s.customerRepo.Update(id, customer)
}

func (s *customerServiceImpl) DeleteCustomer(id string) error {
	return s.customerRepo.Delete(id)
}
