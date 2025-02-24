package repositories

import (
	"github.com/aronipurwanto/go-sample-pos/models"
	"gorm.io/gorm"
)

type EmployeeRepository interface {
	Create(Employee *models.Employee) error
	GetAll() ([]models.Employee, error)
	GetByID(id string) (*models.Employee, error)
	Update(id string, Employee *models.Employee) error
	Delete(id string) error
}

type employeeRepositoryImpl struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &employeeRepositoryImpl{db: db}
}

func (r *employeeRepositoryImpl) Create(Employee *models.Employee) error {
	return r.db.Create(Employee).Error
}

func (r *employeeRepositoryImpl) GetAll() ([]models.Employee, error) {
	var Employees []models.Employee
	err := r.db.Find(&Employees).Error
	return Employees, err
}

func (r *employeeRepositoryImpl) GetByID(id string) (*models.Employee, error) {
	var Employee models.Employee
	err := r.db.First(&Employee, "Employee_id = ?", id).Error
	return &Employee, err
}

func (r *employeeRepositoryImpl) Update(id string, Employee *models.Employee) error {
	return r.db.Model(&models.Employee{}).Where("Employee_id = ?", id).Updates(Employee).Error
}

func (r *employeeRepositoryImpl) Delete(id string) error {
	return r.db.Delete(&models.Employee{}, "Employee_id = ?", id).Error
}
