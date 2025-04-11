// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateRepository43Reader is a Reader for the CreateRepository43 structure.
type CreateRepository43Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRepository43Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateRepository43Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateRepository43Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateRepository43Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/repositories/yum/hosted] createRepository_43", response, response.Code())
	}
}

// NewCreateRepository43Created creates a CreateRepository43Created with default headers values
func NewCreateRepository43Created() *CreateRepository43Created {
	return &CreateRepository43Created{}
}

/*
CreateRepository43Created describes a response with status code 201, with default header values.

Repository created
*/
type CreateRepository43Created struct {
}

// IsSuccess returns true when this create repository43 created response has a 2xx status code
func (o *CreateRepository43Created) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create repository43 created response has a 3xx status code
func (o *CreateRepository43Created) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository43 created response has a 4xx status code
func (o *CreateRepository43Created) IsClientError() bool {
	return false
}

// IsServerError returns true when this create repository43 created response has a 5xx status code
func (o *CreateRepository43Created) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository43 created response a status code equal to that given
func (o *CreateRepository43Created) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create repository43 created response
func (o *CreateRepository43Created) Code() int {
	return 201
}

func (o *CreateRepository43Created) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/yum/hosted][%d] createRepository43Created", 201)
}

func (o *CreateRepository43Created) String() string {
	return fmt.Sprintf("[POST /v1/repositories/yum/hosted][%d] createRepository43Created", 201)
}

func (o *CreateRepository43Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRepository43Unauthorized creates a CreateRepository43Unauthorized with default headers values
func NewCreateRepository43Unauthorized() *CreateRepository43Unauthorized {
	return &CreateRepository43Unauthorized{}
}

/*
CreateRepository43Unauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type CreateRepository43Unauthorized struct {
}

// IsSuccess returns true when this create repository43 unauthorized response has a 2xx status code
func (o *CreateRepository43Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create repository43 unauthorized response has a 3xx status code
func (o *CreateRepository43Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository43 unauthorized response has a 4xx status code
func (o *CreateRepository43Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create repository43 unauthorized response has a 5xx status code
func (o *CreateRepository43Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository43 unauthorized response a status code equal to that given
func (o *CreateRepository43Unauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create repository43 unauthorized response
func (o *CreateRepository43Unauthorized) Code() int {
	return 401
}

func (o *CreateRepository43Unauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/yum/hosted][%d] createRepository43Unauthorized", 401)
}

func (o *CreateRepository43Unauthorized) String() string {
	return fmt.Sprintf("[POST /v1/repositories/yum/hosted][%d] createRepository43Unauthorized", 401)
}

func (o *CreateRepository43Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRepository43Forbidden creates a CreateRepository43Forbidden with default headers values
func NewCreateRepository43Forbidden() *CreateRepository43Forbidden {
	return &CreateRepository43Forbidden{}
}

/*
CreateRepository43Forbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type CreateRepository43Forbidden struct {
}

// IsSuccess returns true when this create repository43 forbidden response has a 2xx status code
func (o *CreateRepository43Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create repository43 forbidden response has a 3xx status code
func (o *CreateRepository43Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository43 forbidden response has a 4xx status code
func (o *CreateRepository43Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create repository43 forbidden response has a 5xx status code
func (o *CreateRepository43Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository43 forbidden response a status code equal to that given
func (o *CreateRepository43Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create repository43 forbidden response
func (o *CreateRepository43Forbidden) Code() int {
	return 403
}

func (o *CreateRepository43Forbidden) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/yum/hosted][%d] createRepository43Forbidden", 403)
}

func (o *CreateRepository43Forbidden) String() string {
	return fmt.Sprintf("[POST /v1/repositories/yum/hosted][%d] createRepository43Forbidden", 403)
}

func (o *CreateRepository43Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
