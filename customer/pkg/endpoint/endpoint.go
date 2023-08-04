package endpoint

import (
	"context"

	endpoint "github.com/go-kit/kit/endpoint"
	"github.com/msegeya56/todo-kit-microservice/customer/pkg/domains/entity"
	"github.com/msegeya56/todo-kit-microservice/customer/pkg/domains/model"
	service "github.com/msegeya56/todo-kit-microsrvice/customer/pkg/service"
)

// PostRequest collects the request parameters for the Post method.
type PostRequest struct {
	Data *model.Customer `json:"data"`
}

// PostResponse collects the response parameters for the Post method.
type PostResponse struct {
	ReplyData  *entity.Customer `json:"reply_data"`
	ReplyError error            `json:"reply_error"`
}

// MakePostEndpoint returns an endpoint that invokes Post on the service.
func MakePostEndpoint(s service.CustomerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PostRequest)
		replyData, replyError := s.Post(ctx, req.Data)
		return PostResponse{
			ReplyData:  replyData,
			ReplyError: replyError,
		}, nil
	}
}

// Failed implements Failer.
func (r PostResponse) Failed() error {
	return r.ReplyError
}

// GetAllRequest collects the request parameters for the GetAll method.
type GetAllRequest struct{}

// GetAllResponse collects the response parameters for the GetAll method.
type GetAllResponse struct {
	ReplyData  []*entity.Customer `json:"reply_data"`
	ReplyError error              `json:"reply_error"`
}

// MakeGetAllEndpoint returns an endpoint that invokes GetAll on the service.
func MakeGetAllEndpoint(s service.CustomerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		replyData, replyError := s.GetAll(ctx)
		return GetAllResponse{
			ReplyData:  replyData,
			ReplyError: replyError,
		}, nil
	}
}

// Failed implements Failer.
func (r GetAllResponse) Failed() error {
	return r.ReplyError
}

// GetByIDRequest collects the request parameters for the GetByID method.
type GetByIDRequest struct {
	Param uint `json:"param"`
}

// GetByIDResponse collects the response parameters for the GetByID method.
type GetByIDResponse struct {
	ReplySata  *entity.Customer `json:"reply_sata"`
	ReolyError error            `json:"reoly_error"`
}

// MakeGetByIDEndpoint returns an endpoint that invokes GetByID on the service.
func MakeGetByIDEndpoint(s service.CustomerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetByIDRequest)
		replySata, reolyError := s.GetByID(ctx, req.Param)
		return GetByIDResponse{
			ReolyError: reolyError,
			ReplySata:  replySata,
		}, nil
	}
}

// Failed implements Failer.
func (r GetByIDResponse) Failed() error {
	return r.ReolyError
}

// UpdateRequest collects the request parameters for the Update method.
type UpdateRequest struct {
	Paparam uint            `json:"paparam"`
	Data    *model.Customer `json:"data"`
}

// UpdateResponse collects the response parameters for the Update method.
type UpdateResponse struct {
	ReplyData  *entity.Customer `json:"reply_data"`
	ReplyError error            `json:"reply_error"`
}

// MakeUpdateEndpoint returns an endpoint that invokes Update on the service.
func MakeUpdateEndpoint(s service.CustomerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateRequest)
		replyData, replyError := s.Update(ctx, req.Paparam, req.Data)
		return UpdateResponse{
			ReplyData:  replyData,
			ReplyError: replyError,
		}, nil
	}
}

// Failed implements Failer.
func (r UpdateResponse) Failed() error {
	return r.ReplyError
}

// GetByDateRequest collects the request parameters for the GetByDate method.
type GetByDateRequest struct {
	Param string `json:"param"`
}

// GetByDateResponse collects the response parameters for the GetByDate method.
type GetByDateResponse struct {
	ReplyData  []entity.Customer `json:"reply_data"`
	ReplyError error             `json:"reply_error"`
}

// MakeGetByDateEndpoint returns an endpoint that invokes GetByDate on the service.
func MakeGetByDateEndpoint(s service.CustomerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetByDateRequest)
		replyData, replyError := s.GetByDate(ctx, req.Param)
		return GetByDateResponse{
			ReplyData:  replyData,
			ReplyError: replyError,
		}, nil
	}
}

// Failed implements Failer.
func (r GetByDateResponse) Failed() error {
	return r.ReplyError
}

// DeleteRequest collects the request parameters for the Delete method.
type DeleteRequest struct {
	Param uint `json:"param"`
}

// DeleteResponse collects the response parameters for the Delete method.
type DeleteResponse struct {
	ReplyError error `json:"reply_error"`
}

// MakeDeleteEndpoint returns an endpoint that invokes Delete on the service.
func MakeDeleteEndpoint(s service.CustomerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteRequest)
		replyError := s.Delete(ctx, req.Param)
		return DeleteResponse{ReplyError: replyError}, nil
	}
}

// Failed implements Failer.
func (r DeleteResponse) Failed() error {
	return r.ReplyError
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Post implements Service. Primarily useful in a client.
func (e Endpoints) Post(ctx context.Context, data *model.Customer) (replyData *entity.Customer, replyError error) {
	request := PostRequest{Data: data}
	response, err := e.PostEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(PostResponse).ReplyData, response.(PostResponse).ReplyError
}

// GetAll implements Service. Primarily useful in a client.
func (e Endpoints) GetAll(ctx context.Context) (replyData []*entity.Customer, replyError error) {
	request := GetAllRequest{}
	response, err := e.GetAllEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetAllResponse).ReplyData, response.(GetAllResponse).ReplyError
}

// GetByID implements Service. Primarily useful in a client.
func (e Endpoints) GetByID(ctx context.Context, param uint) (replySata *entity.Customer, reolyError error) {
	request := GetByIDRequest{Param: param}
	response, err := e.GetByIDEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetByIDResponse).ReplySata, response.(GetByIDResponse).ReolyError
}

// Update implements Service. Primarily useful in a client.
func (e Endpoints) Update(ctx context.Context, paparam uint, data *model.Customer) (replyData *entity.Customer, replyError error) {
	request := UpdateRequest{
		Data:    data,
		Paparam: paparam,
	}
	response, err := e.UpdateEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UpdateResponse).ReplyData, response.(UpdateResponse).ReplyError
}

// GetByDate implements Service. Primarily useful in a client.
func (e Endpoints) GetByDate(ctx context.Context, param string) (replyData []entity.Customer, replyError error) {
	request := GetByDateRequest{Param: param}
	response, err := e.GetByDateEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetByDateResponse).ReplyData, response.(GetByDateResponse).ReplyError
}

// Delete implements Service. Primarily useful in a client.
func (e Endpoints) Delete(ctx context.Context, param uint) (replyError error) {
	request := DeleteRequest{Param: param}
	response, err := e.DeleteEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteResponse).ReplyError
}
