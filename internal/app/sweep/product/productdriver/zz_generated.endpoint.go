// +build !ignore_autogenerated

// Code generated by mga tool. DO NOT EDIT.

package productdriver

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"github.com/prasetyowira/sweep/internal/app/sweep/product"
	kitxendpoint "github.com/sagikazarmark/kitx/endpoint"
)

// endpointError identifies an error that should be returned as an endpoint error.
type endpointError interface {
	EndpointError() bool
}

// serviceError identifies an error that should be returned as a service error.
type serviceError interface {
	ServiceError() bool
}

// Endpoints collects all of the endpoints that compose the underlying service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	CreateProduct   endpoint.Endpoint
	GetProduct      endpoint.Endpoint
	GetProductBySKU endpoint.Endpoint
	ListProducts    endpoint.Endpoint
}

// MakeEndpoints returns a(n) Endpoints struct where each endpoint invokes
// the corresponding method on the provided service.
func MakeEndpoints(service product.Service, middleware ...endpoint.Middleware) Endpoints {
	mw := kitxendpoint.Combine(middleware...)

	return Endpoints{
		CreateProduct:   kitxendpoint.OperationNameMiddleware("product.CreateProduct")(mw(MakeCreateProductEndpoint(service))),
		GetProduct:      kitxendpoint.OperationNameMiddleware("product.GetProduct")(mw(MakeGetProductEndpoint(service))),
		GetProductBySKU: kitxendpoint.OperationNameMiddleware("product.GetProductBySKU")(mw(MakeGetProductBySKUEndpoint(service))),
		ListProducts:    kitxendpoint.OperationNameMiddleware("product.ListProducts")(mw(MakeListProductsEndpoint(service))),
	}
}

// CreateProductRequest is a request struct for CreateProduct endpoint.
type CreateProductRequest struct {
	Sku       string
	Name      string
	Expirable bool
}

// CreateProductResponse is a response struct for CreateProduct endpoint.
type CreateProductResponse struct {
	Id  string
	Err error
}

func (r CreateProductResponse) Failed() error {
	return r.Err
}

// MakeCreateProductEndpoint returns an endpoint for the matching method of the underlying service.
func MakeCreateProductEndpoint(service product.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateProductRequest)

		id, err := service.CreateProduct(ctx, req.Sku, req.Name, req.Expirable)

		if err != nil {
			if serviceErr := serviceError(nil); errors.As(err, &serviceErr) && serviceErr.ServiceError() {
				return CreateProductResponse{
					Err: err,
					Id:  id,
				}, nil
			}

			return CreateProductResponse{
				Err: err,
				Id:  id,
			}, err
		}

		return CreateProductResponse{Id: id}, nil
	}
}

// GetProductRequest is a request struct for GetProduct endpoint.
type GetProductRequest struct {
	Id string
}

// GetProductResponse is a response struct for GetProduct endpoint.
type GetProductResponse struct {
	Product product.Product
	Err     error
}

func (r GetProductResponse) Failed() error {
	return r.Err
}

// MakeGetProductEndpoint returns an endpoint for the matching method of the underlying service.
func MakeGetProductEndpoint(service product.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetProductRequest)

		product, err := service.GetProduct(ctx, req.Id)

		if err != nil {
			if serviceErr := serviceError(nil); errors.As(err, &serviceErr) && serviceErr.ServiceError() {
				return GetProductResponse{
					Err:     err,
					Product: product,
				}, nil
			}

			return GetProductResponse{
				Err:     err,
				Product: product,
			}, err
		}

		return GetProductResponse{Product: product}, nil
	}
}

// GetProductBySKURequest is a request struct for GetProductBySKU endpoint.
type GetProductBySKURequest struct {
	Sku string
}

// GetProductBySKUResponse is a response struct for GetProductBySKU endpoint.
type GetProductBySKUResponse struct {
	Product product.Product
	Err     error
}

func (r GetProductBySKUResponse) Failed() error {
	return r.Err
}

// MakeGetProductBySKUEndpoint returns an endpoint for the matching method of the underlying service.
func MakeGetProductBySKUEndpoint(service product.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetProductBySKURequest)

		product, err := service.GetProductBySKU(ctx, req.Sku)

		if err != nil {
			if serviceErr := serviceError(nil); errors.As(err, &serviceErr) && serviceErr.ServiceError() {
				return GetProductBySKUResponse{
					Err:     err,
					Product: product,
				}, nil
			}

			return GetProductBySKUResponse{
				Err:     err,
				Product: product,
			}, err
		}

		return GetProductBySKUResponse{Product: product}, nil
	}
}

// ListProductsRequest is a request struct for ListProducts endpoint.
type ListProductsRequest struct{}

// ListProductsResponse is a response struct for ListProducts endpoint.
type ListProductsResponse struct {
	Products []product.Product
	Err      error
}

func (r ListProductsResponse) Failed() error {
	return r.Err
}

// MakeListProductsEndpoint returns an endpoint for the matching method of the underlying service.
func MakeListProductsEndpoint(service product.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		products, err := service.ListProducts(ctx)

		if err != nil {
			if serviceErr := serviceError(nil); errors.As(err, &serviceErr) && serviceErr.ServiceError() {
				return ListProductsResponse{
					Err:      err,
					Products: products,
				}, nil
			}

			return ListProductsResponse{
				Err:      err,
				Products: products,
			}, err
		}

		return ListProductsResponse{Products: products}, nil
	}
}