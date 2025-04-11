// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateRepository16Reader is a Reader for the CreateRepository16 structure.
type CreateRepository16Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRepository16Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateRepository16Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateRepository16Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateRepository16Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewCreateRepository16MethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/repositories/conda/proxy] createRepository_16", response, response.Code())
	}
}

// NewCreateRepository16Created creates a CreateRepository16Created with default headers values
func NewCreateRepository16Created() *CreateRepository16Created {
	return &CreateRepository16Created{}
}

/*
CreateRepository16Created describes a response with status code 201, with default header values.

Repository created
*/
type CreateRepository16Created struct {
}

// IsSuccess returns true when this create repository16 created response has a 2xx status code
func (o *CreateRepository16Created) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create repository16 created response has a 3xx status code
func (o *CreateRepository16Created) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository16 created response has a 4xx status code
func (o *CreateRepository16Created) IsClientError() bool {
	return false
}

// IsServerError returns true when this create repository16 created response has a 5xx status code
func (o *CreateRepository16Created) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository16 created response a status code equal to that given
func (o *CreateRepository16Created) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create repository16 created response
func (o *CreateRepository16Created) Code() int {
	return 201
}

func (o *CreateRepository16Created) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/conda/proxy][%d] createRepository16Created", 201)
}

func (o *CreateRepository16Created) String() string {
	return fmt.Sprintf("[POST /v1/repositories/conda/proxy][%d] createRepository16Created", 201)
}

func (o *CreateRepository16Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRepository16Unauthorized creates a CreateRepository16Unauthorized with default headers values
func NewCreateRepository16Unauthorized() *CreateRepository16Unauthorized {
	return &CreateRepository16Unauthorized{}
}

/*
CreateRepository16Unauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type CreateRepository16Unauthorized struct {
}

// IsSuccess returns true when this create repository16 unauthorized response has a 2xx status code
func (o *CreateRepository16Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create repository16 unauthorized response has a 3xx status code
func (o *CreateRepository16Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository16 unauthorized response has a 4xx status code
func (o *CreateRepository16Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create repository16 unauthorized response has a 5xx status code
func (o *CreateRepository16Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository16 unauthorized response a status code equal to that given
func (o *CreateRepository16Unauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create repository16 unauthorized response
func (o *CreateRepository16Unauthorized) Code() int {
	return 401
}

func (o *CreateRepository16Unauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/conda/proxy][%d] createRepository16Unauthorized", 401)
}

func (o *CreateRepository16Unauthorized) String() string {
	return fmt.Sprintf("[POST /v1/repositories/conda/proxy][%d] createRepository16Unauthorized", 401)
}

func (o *CreateRepository16Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRepository16Forbidden creates a CreateRepository16Forbidden with default headers values
func NewCreateRepository16Forbidden() *CreateRepository16Forbidden {
	return &CreateRepository16Forbidden{}
}

/*
CreateRepository16Forbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type CreateRepository16Forbidden struct {
}

// IsSuccess returns true when this create repository16 forbidden response has a 2xx status code
func (o *CreateRepository16Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create repository16 forbidden response has a 3xx status code
func (o *CreateRepository16Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository16 forbidden response has a 4xx status code
func (o *CreateRepository16Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create repository16 forbidden response has a 5xx status code
func (o *CreateRepository16Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository16 forbidden response a status code equal to that given
func (o *CreateRepository16Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create repository16 forbidden response
func (o *CreateRepository16Forbidden) Code() int {
	return 403
}

func (o *CreateRepository16Forbidden) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/conda/proxy][%d] createRepository16Forbidden", 403)
}

func (o *CreateRepository16Forbidden) String() string {
	return fmt.Sprintf("[POST /v1/repositories/conda/proxy][%d] createRepository16Forbidden", 403)
}

func (o *CreateRepository16Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRepository16MethodNotAllowed creates a CreateRepository16MethodNotAllowed with default headers values
func NewCreateRepository16MethodNotAllowed() *CreateRepository16MethodNotAllowed {
	return &CreateRepository16MethodNotAllowed{}
}

/*
CreateRepository16MethodNotAllowed describes a response with status code 405, with default header values.

Feature is disabled in High Availability
*/
type CreateRepository16MethodNotAllowed struct {
}

// IsSuccess returns true when this create repository16 method not allowed response has a 2xx status code
func (o *CreateRepository16MethodNotAllowed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create repository16 method not allowed response has a 3xx status code
func (o *CreateRepository16MethodNotAllowed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository16 method not allowed response has a 4xx status code
func (o *CreateRepository16MethodNotAllowed) IsClientError() bool {
	return true
}

// IsServerError returns true when this create repository16 method not allowed response has a 5xx status code
func (o *CreateRepository16MethodNotAllowed) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository16 method not allowed response a status code equal to that given
func (o *CreateRepository16MethodNotAllowed) IsCode(code int) bool {
	return code == 405
}

// Code gets the status code for the create repository16 method not allowed response
func (o *CreateRepository16MethodNotAllowed) Code() int {
	return 405
}

func (o *CreateRepository16MethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/conda/proxy][%d] createRepository16MethodNotAllowed", 405)
}

func (o *CreateRepository16MethodNotAllowed) String() string {
	return fmt.Sprintf("[POST /v1/repositories/conda/proxy][%d] createRepository16MethodNotAllowed", 405)
}

func (o *CreateRepository16MethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
