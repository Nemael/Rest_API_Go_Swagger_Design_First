// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetBookHandlerFunc turns a function with the right signature into a get book handler
type GetBookHandlerFunc func(GetBookParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetBookHandlerFunc) Handle(params GetBookParams) middleware.Responder {
	return fn(params)
}

// GetBookHandler interface for that can handle valid get book params
type GetBookHandler interface {
	Handle(GetBookParams) middleware.Responder
}

// NewGetBook creates a new http.Handler for the get book operation
func NewGetBook(ctx *middleware.Context, handler GetBookHandler) *GetBook {
	return &GetBook{Context: ctx, Handler: handler}
}

/*
	GetBook swagger:route GET /book/{id} getBook

Return a book from ID
*/
type GetBook struct {
	Context *middleware.Context
	Handler GetBookHandler
}

func (o *GetBook) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetBookParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
