// repository.go

package repository

import (
	"context"

	"github.com/msegeya56/todo-kit-microservice/customer/pkg/connection" // Import the connection package
	"github.com/msegeya56/todo-kit-microservice/customer/pkg/domains/entity"
	"github.com/msegeya56/todo-kit-microservice/customer/pkg/domains/model"
	"gorm.io/gorm" // Import GORM package
)

type DatabaseCustomerRepository struct {
	DB *gorm.DB // Add GORM DB instance as a field
}

// Post implements CustomerRepository.
func (*DatabaseCustomerRepository) Post(ctx context.Context, data *model.Customer) (*entity.Customer, error) {
	panic("unimplemented")
}

// CustomerRepository interface defines methods to interact with the customer data in the database.
type CustomerRepository interface {
	Post(ctx context.Context, data *model.Customer) (*entity.Customer, error)
	GetAll(ctx context.Context) ([]*entity.Customer, error)
	GetByID(ctx context.Context, id uint) (*entity.Customer, error)
	Update(ctx context.Context, id uint, data *model.Customer) (*entity.Customer, error)
	GetByDate(ctx context.Context, date string) ([]entity.Customer, error)
	Delete(ctx context.Context, id uint) error
}

func NewDatabaseCustomerRepository() CustomerRepository {
	// Initialize and return the DatabaseCustomerRepository with the database connection
	db := connection.DB // Use the DB instance from the connection package
	return &DatabaseCustomerRepository{DB: db}
}

// Implement the methods of the CustomerRepository interface

func (r *DatabaseCustomerRepository) Add(ctx context.Context, data *model.Customer) (*entity.Customer, error) {
	// Your implementation to add a new customer to the database
	// Use the database connection and appropriate ORM (like GORM) to execute the insertion query
	// Return the added customer (replyData) and any potential error (replyError)

	// For example, with GORM:
	newCustomer := &entity.Customer{
		Username: data.Username,
		Email:    data.Email,
		Password: data.Password,
		FullName: data.FullName,
		Phone:    data.Phone,
		Address:  data.Address,
	}

	// Insert the new customer into the database
	result := r.DB.Create(newCustomer)
	if result.Error != nil {
		return nil, result.Error
	}

	return newCustomer, nil
}

func (r *DatabaseCustomerRepository) GetAll(ctx context.Context) ([]*entity.Customer, error) {
	var customers []*entity.Customer
	result := r.DB.Find(&customers)
	if result.Error != nil {
		return nil, result.Error
	}
	return customers, nil
}

func (r *DatabaseCustomerRepository) GetByID(ctx context.Context, id uint) (*entity.Customer, error) {
	var customer entity.Customer
	result := r.DB.First(&customer, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &customer, nil
}

func (r *DatabaseCustomerRepository) Update(ctx context.Context, id uint, data *model.Customer) (*entity.Customer, error) {
	// Get the existing customer
	customer, err := r.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Update the customer fields with the data
	customer.Username = data.Username
	customer.Email = data.Email
	customer.Password = data.Password
	customer.FullName = data.FullName
	customer.Phone = data.Phone
	customer.Address = data.Address

	// Save the updated customer back to the database
	result := r.DB.Save(&customer)
	if result.Error != nil {
		return nil, result.Error
	}
	return customer, nil
}

func (r *DatabaseCustomerRepository) GetByDate(ctx context.Context, date string) ([]entity.Customer, error) {
	var customers []entity.Customer
	result := r.DB.Where("date_column = ?", date).Find(&customers)
	if result.Error != nil {
		return nil, result.Error
	}
	return customers, nil
}

func (r *DatabaseCustomerRepository) Delete(ctx context.Context, id uint) error {
	result := r.DB.Delete(&entity.Customer{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
