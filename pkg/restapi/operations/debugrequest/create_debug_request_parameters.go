// Code generated by go-swagger; DO NOT EDIT.

package debugrequest

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/solo-io/squash/pkg/models"
)

// NewCreateDebugRequestParams creates a new CreateDebugRequestParams object
// with the default values initialized.
func NewCreateDebugRequestParams() CreateDebugRequestParams {
	var ()
	return CreateDebugRequestParams{}
}

// CreateDebugRequestParams contains all the bound params for the create debug request operation
// typically these are obtained from a http.Request
//
// swagger:parameters createDebugRequest
type CreateDebugRequestParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*DebugRequest object
	  Required: true
	  In: body
	*/
	Body *models.DebugRequest
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *CreateDebugRequestParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.DebugRequest
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("body", "body"))
			} else {
				res = append(res, errors.NewParseError("body", "body", "", err))
			}

		} else {
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = &body
			}
		}

	} else {
		res = append(res, errors.Required("body", "body"))
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}