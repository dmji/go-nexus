// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateRepository8Reader is a Reader for the CreateRepository8 structure.
type CreateRepository8Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRepository8Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateRepository8Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateRepository8Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateRepository8Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/repositories/cargo/group] createRepository_8", response, response.Code())
	}
}

// NewCreateRepository8Created creates a CreateRepository8Created with default headers values
func NewCreateRepository8Created() *CreateRepository8Created {
	return &CreateRepository8Created{}
}

/*
CreateRepository8Created describes a response with status code 201, with default header values.

Repository created
*/
type CreateRepository8Created struct {
}

// IsSuccess returns true when this create repository8 created response has a 2xx status code
func (o *CreateRepository8Created) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create repository8 created response has a 3xx status code
func (o *CreateRepository8Created) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository8 created response has a 4xx status code
func (o *CreateRepository8Created) IsClientError() bool {
	return false
}

// IsServerError returns true when this create repository8 created response has a 5xx status code
func (o *CreateRepository8Created) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository8 created response a status code equal to that given
func (o *CreateRepository8Created) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create repository8 created response
func (o *CreateRepository8Created) Code() int {
	return 201
}

func (o *CreateRepository8Created) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/cargo/group][%d] createRepository8Created", 201)
}

func (o *CreateRepository8Created) String() string {
	return fmt.Sprintf("[POST /v1/repositories/cargo/group][%d] createRepository8Created", 201)
}

func (o *CreateRepository8Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRepository8Unauthorized creates a CreateRepository8Unauthorized with default headers values
func NewCreateRepository8Unauthorized() *CreateRepository8Unauthorized {
	return &CreateRepository8Unauthorized{}
}

/*
CreateRepository8Unauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type CreateRepository8Unauthorized struct {
}

// IsSuccess returns true when this create repository8 unauthorized response has a 2xx status code
func (o *CreateRepository8Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create repository8 unauthorized response has a 3xx status code
func (o *CreateRepository8Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository8 unauthorized response has a 4xx status code
func (o *CreateRepository8Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create repository8 unauthorized response has a 5xx status code
func (o *CreateRepository8Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository8 unauthorized response a status code equal to that given
func (o *CreateRepository8Unauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create repository8 unauthorized response
func (o *CreateRepository8Unauthorized) Code() int {
	return 401
}

func (o *CreateRepository8Unauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/cargo/group][%d] createRepository8Unauthorized", 401)
}

func (o *CreateRepository8Unauthorized) String() string {
	return fmt.Sprintf("[POST /v1/repositories/cargo/group][%d] createRepository8Unauthorized", 401)
}

func (o *CreateRepository8Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRepository8Forbidden creates a CreateRepository8Forbidden with default headers values
func NewCreateRepository8Forbidden() *CreateRepository8Forbidden {
	return &CreateRepository8Forbidden{}
}

/*
CreateRepository8Forbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type CreateRepository8Forbidden struct {
}

// IsSuccess returns true when this create repository8 forbidden response has a 2xx status code
func (o *CreateRepository8Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create repository8 forbidden response has a 3xx status code
func (o *CreateRepository8Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository8 forbidden response has a 4xx status code
func (o *CreateRepository8Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create repository8 forbidden response has a 5xx status code
func (o *CreateRepository8Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository8 forbidden response a status code equal to that given
func (o *CreateRepository8Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create repository8 forbidden response
func (o *CreateRepository8Forbidden) Code() int {
	return 403
}

func (o *CreateRepository8Forbidden) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/cargo/group][%d] createRepository8Forbidden", 403)
}

func (o *CreateRepository8Forbidden) String() string {
	return fmt.Sprintf("[POST /v1/repositories/cargo/group][%d] createRepository8Forbidden", 403)
}

func (o *CreateRepository8Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
