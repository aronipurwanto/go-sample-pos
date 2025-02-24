package services

import (
	"github.com/aronipurwanto/go-sample-pos/models"
	"github.com/aronipurwanto/go-sample-pos/repositories"
)

type EmployeeService interface {
	CreateEmployee(Employee *models.Employee) error
	GetAllEmployee() ([]models.Employee, error)
	GetEmployeeByID(id string) (*models.Employee, error)
	UpdateEmployee(id string, Employee *models.Employee) error
	DeleteEmployee(id string) error
}

type employeeServiceImpl struct {
	EmployeeRepo repositories.EmployeeRepository
}

func NewEmployeeService(repo repositories.EmployeeRepository) EmployeeService {
	return &employeeServiceImpl{repo}
}

func (s *employeeServiceImpl) CreateEmployee(Employee *models.Employee) error {
	return s.EmployeeRepo.Create(Employee)
}

func (s *employeeServiceImpl) GetAllEmployee() ([]models.Employee, error) {
	return s.EmployeeRepo.GetAll()
}

func (s *employeeServiceImpl) GetEmployeeByID(id string) (*models.Employee, error) {
	return s.EmployeeRepo.GetByID(id)
}

func (s *employeeServiceImpl) UpdateEmployee(id string, Employee *models.Employee) error {
	return s.EmployeeRepo.Update(id, Employee)
}

func (s *employeeServiceImpl) DeleteEmployee(id string) error {
	return s.EmployeeRepo.Delete(id)
}
