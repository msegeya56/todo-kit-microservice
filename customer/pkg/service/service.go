package service

import (
	"context"

	"github.com/msegeya56/todo-kit-microservice/customer/pkg/domains/entity"
	"github.com/msegeya56/todo-kit-microservice/customer/pkg/domains/model"
	"github.com/msegeya56/todo-kit-microservice/customer/pkg/repository"
)

// CustomerService describes the service.
type CustomerService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	Post(ctx context.Context, data *model.Customer) (replyData *entity.Customer, replyError error)
	GetAll(ctx context.Context) (replyData []*entity.Customer, replyError error)
	GetByID(ctx context.Context, param uint) (replySata *entity.Customer, reolyError error)
	Update(ctx context.Context, paparam uint, data *model.Customer) (replyData *entity.Customer, replyError error)
	GetByDate(ctx context.Context, param string) (replyData []entity.Customer, replyError error)
	Delete(ctx context.Context, param uint) (replyError error)
}

type basicCustomerService struct{
	customerRepo repository.CustomerRepository
}

func (b *basicCustomerService) Post(ctx context.Context, data *model.Customer) (replyData *entity.Customer, replyError error) {
	// TODO implement the business logic of Post
	return replyData, replyError
}
func (b *basicCustomerService) GetAll(ctx context.Context) (replyData []*entity.Customer, replyError error) {
	// TODO implement the business logic of GetAll
	return replyData, replyError
}
func (b *basicCustomerService) GetByID(ctx context.Context, param uint) (replySata *entity.Customer, reolyError error) {
	// TODO implement the business logic of GetByID
	return replySata, reolyError
}
func (b *basicCustomerService) Update(ctx context.Context, paparam uint, data *model.Customer) (replyData *entity.Customer, replyError error) {
	// TODO implement the business logic of Update
	return replyData, replyError
}
func (b *basicCustomerService) GetByDate(ctx context.Context, param string) (replyData []entity.Customer, replyError error) {
	// TODO implement the business logic of GetByDate
	return replyData, replyError
}
func (b *basicCustomerService) Delete(ctx context.Context, param uint) (replyError error) {
	// TODO implement the business logic of Delete
	return replyError
}

// NewBasicCustomerService returns a naive, stateless implementation of CustomerService.
func NewBasicCustomerService() CustomerService {
	return &basicCustomerService{}
}

// New returns a CustomerService with all of the expected middleware wired in.
func New(middleware []Middleware) CustomerService {
	var svc CustomerService = NewBasicCustomerService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
