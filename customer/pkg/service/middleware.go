package service

import (
	"context"

	log "github.com/go-kit/kit/log"
	"github.com/msegeya56/todo-kit-microservice/customer/pkg/domains/entity"
	"github.com/msegeya56/todo-kit-microservice/customer/pkg/domains/model"
)

// Middleware describes a service middleware.
type Middleware func(CustomerService) CustomerService

type loggingMiddleware struct {
	logger log.Logger
	next   CustomerService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a CustomerService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next CustomerService) CustomerService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Post(ctx context.Context, data *model.Customer) (replyData *entity.Customer, replyError error) {
	defer func() {
		l.logger.Log("method", "Post", "data", data, "replyData", replyData, "replyError", replyError)
	}()
	return l.next.Post(ctx, data)
}
func (l loggingMiddleware) GetAll(ctx context.Context) (replyData []*entity.Customer, replyError error) {
	defer func() {
		l.logger.Log("method", "GetAll", "replyData", replyData, "replyError", replyError)
	}()
	return l.next.GetAll(ctx)
}
func (l loggingMiddleware) GetByID(ctx context.Context, param uint) (replySata *entity.Customer, reolyError error) {
	defer func() {
		l.logger.Log("method", "GetByID", "param", param, "replySata", replySata, "reolyError", reolyError)
	}()
	return l.next.GetByID(ctx, param)
}
func (l loggingMiddleware) Update(ctx context.Context, paparam uint, data *model.Customer) (replyData *entity.Customer, replyError error) {
	defer func() {
		l.logger.Log("method", "Update", "paparam", paparam, "data", data, "replyData", replyData, "replyError", replyError)
	}()
	return l.next.Update(ctx, paparam, data)
}
func (l loggingMiddleware) GetByDate(ctx context.Context, param string) (replyData []entity.Customer, replyError error) {
	defer func() {
		l.logger.Log("method", "GetByDate", "param", param, "replyData", replyData, "replyError", replyError)
	}()
	return l.next.GetByDate(ctx, param)
}
func (l loggingMiddleware) Delete(ctx context.Context, param uint) (replyError error) {
	defer func() {
		l.logger.Log("method", "Delete", "param", param, "replyError", replyError)
	}()
	return l.next.Delete(ctx, param)
}
