// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateRepository12Reader is a Reader for the CreateRepository12 structure.
type CreateRepository12Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRepository12Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateRepository12Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateRepository12Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateRepository12Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/repositories/composer/proxy] createRepository_12", response, response.Code())
	}
}

// NewCreateRepository12Created creates a CreateRepository12Created with default headers values
func NewCreateRepository12Created() *CreateRepository12Created {
	return &CreateRepository12Created{}
}

/*
CreateRepository12Created describes a response with status code 201, with default header values.

Repository created
*/
type CreateRepository12Created struct {
}

// IsSuccess returns true when this create repository12 created response has a 2xx status code
func (o *CreateRepository12Created) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create repository12 created response has a 3xx status code
func (o *CreateRepository12Created) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository12 created response has a 4xx status code
func (o *CreateRepository12Created) IsClientError() bool {
	return false
}

// IsServerError returns true when this create repository12 created response has a 5xx status code
func (o *CreateRepository12Created) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository12 created response a status code equal to that given
func (o *CreateRepository12Created) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create repository12 created response
func (o *CreateRepository12Created) Code() int {
	return 201
}

func (o *CreateRepository12Created) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/composer/proxy][%d] createRepository12Created", 201)
}

func (o *CreateRepository12Created) String() string {
	return fmt.Sprintf("[POST /v1/repositories/composer/proxy][%d] createRepository12Created", 201)
}

func (o *CreateRepository12Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRepository12Unauthorized creates a CreateRepository12Unauthorized with default headers values
func NewCreateRepository12Unauthorized() *CreateRepository12Unauthorized {
	return &CreateRepository12Unauthorized{}
}

/*
CreateRepository12Unauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type CreateRepository12Unauthorized struct {
}

// IsSuccess returns true when this create repository12 unauthorized response has a 2xx status code
func (o *CreateRepository12Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create repository12 unauthorized response has a 3xx status code
func (o *CreateRepository12Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository12 unauthorized response has a 4xx status code
func (o *CreateRepository12Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create repository12 unauthorized response has a 5xx status code
func (o *CreateRepository12Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository12 unauthorized response a status code equal to that given
func (o *CreateRepository12Unauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create repository12 unauthorized response
func (o *CreateRepository12Unauthorized) Code() int {
	return 401
}

func (o *CreateRepository12Unauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/composer/proxy][%d] createRepository12Unauthorized", 401)
}

func (o *CreateRepository12Unauthorized) String() string {
	return fmt.Sprintf("[POST /v1/repositories/composer/proxy][%d] createRepository12Unauthorized", 401)
}

func (o *CreateRepository12Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRepository12Forbidden creates a CreateRepository12Forbidden with default headers values
func NewCreateRepository12Forbidden() *CreateRepository12Forbidden {
	return &CreateRepository12Forbidden{}
}

/*
CreateRepository12Forbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type CreateRepository12Forbidden struct {
}

// IsSuccess returns true when this create repository12 forbidden response has a 2xx status code
func (o *CreateRepository12Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create repository12 forbidden response has a 3xx status code
func (o *CreateRepository12Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository12 forbidden response has a 4xx status code
func (o *CreateRepository12Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create repository12 forbidden response has a 5xx status code
func (o *CreateRepository12Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository12 forbidden response a status code equal to that given
func (o *CreateRepository12Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create repository12 forbidden response
func (o *CreateRepository12Forbidden) Code() int {
	return 403
}

func (o *CreateRepository12Forbidden) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/composer/proxy][%d] createRepository12Forbidden", 403)
}

func (o *CreateRepository12Forbidden) String() string {
	return fmt.Sprintf("[POST /v1/repositories/composer/proxy][%d] createRepository12Forbidden", 403)
}

func (o *CreateRepository12Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
