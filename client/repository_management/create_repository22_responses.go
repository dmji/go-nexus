// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateRepository22Reader is a Reader for the CreateRepository22 structure.
type CreateRepository22Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRepository22Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateRepository22Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateRepository22Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateRepository22Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewCreateRepository22MethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/repositories/go/proxy] createRepository_22", response, response.Code())
	}
}

// NewCreateRepository22Created creates a CreateRepository22Created with default headers values
func NewCreateRepository22Created() *CreateRepository22Created {
	return &CreateRepository22Created{}
}

/*
CreateRepository22Created describes a response with status code 201, with default header values.

Repository created
*/
type CreateRepository22Created struct {
}

// IsSuccess returns true when this create repository22 created response has a 2xx status code
func (o *CreateRepository22Created) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create repository22 created response has a 3xx status code
func (o *CreateRepository22Created) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository22 created response has a 4xx status code
func (o *CreateRepository22Created) IsClientError() bool {
	return false
}

// IsServerError returns true when this create repository22 created response has a 5xx status code
func (o *CreateRepository22Created) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository22 created response a status code equal to that given
func (o *CreateRepository22Created) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create repository22 created response
func (o *CreateRepository22Created) Code() int {
	return 201
}

func (o *CreateRepository22Created) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/go/proxy][%d] createRepository22Created", 201)
}

func (o *CreateRepository22Created) String() string {
	return fmt.Sprintf("[POST /v1/repositories/go/proxy][%d] createRepository22Created", 201)
}

func (o *CreateRepository22Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRepository22Unauthorized creates a CreateRepository22Unauthorized with default headers values
func NewCreateRepository22Unauthorized() *CreateRepository22Unauthorized {
	return &CreateRepository22Unauthorized{}
}

/*
CreateRepository22Unauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type CreateRepository22Unauthorized struct {
}

// IsSuccess returns true when this create repository22 unauthorized response has a 2xx status code
func (o *CreateRepository22Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create repository22 unauthorized response has a 3xx status code
func (o *CreateRepository22Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository22 unauthorized response has a 4xx status code
func (o *CreateRepository22Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create repository22 unauthorized response has a 5xx status code
func (o *CreateRepository22Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository22 unauthorized response a status code equal to that given
func (o *CreateRepository22Unauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create repository22 unauthorized response
func (o *CreateRepository22Unauthorized) Code() int {
	return 401
}

func (o *CreateRepository22Unauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/go/proxy][%d] createRepository22Unauthorized", 401)
}

func (o *CreateRepository22Unauthorized) String() string {
	return fmt.Sprintf("[POST /v1/repositories/go/proxy][%d] createRepository22Unauthorized", 401)
}

func (o *CreateRepository22Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRepository22Forbidden creates a CreateRepository22Forbidden with default headers values
func NewCreateRepository22Forbidden() *CreateRepository22Forbidden {
	return &CreateRepository22Forbidden{}
}

/*
CreateRepository22Forbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type CreateRepository22Forbidden struct {
}

// IsSuccess returns true when this create repository22 forbidden response has a 2xx status code
func (o *CreateRepository22Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create repository22 forbidden response has a 3xx status code
func (o *CreateRepository22Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository22 forbidden response has a 4xx status code
func (o *CreateRepository22Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create repository22 forbidden response has a 5xx status code
func (o *CreateRepository22Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository22 forbidden response a status code equal to that given
func (o *CreateRepository22Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create repository22 forbidden response
func (o *CreateRepository22Forbidden) Code() int {
	return 403
}

func (o *CreateRepository22Forbidden) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/go/proxy][%d] createRepository22Forbidden", 403)
}

func (o *CreateRepository22Forbidden) String() string {
	return fmt.Sprintf("[POST /v1/repositories/go/proxy][%d] createRepository22Forbidden", 403)
}

func (o *CreateRepository22Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRepository22MethodNotAllowed creates a CreateRepository22MethodNotAllowed with default headers values
func NewCreateRepository22MethodNotAllowed() *CreateRepository22MethodNotAllowed {
	return &CreateRepository22MethodNotAllowed{}
}

/*
CreateRepository22MethodNotAllowed describes a response with status code 405, with default header values.

Feature is disabled in High Availability
*/
type CreateRepository22MethodNotAllowed struct {
}

// IsSuccess returns true when this create repository22 method not allowed response has a 2xx status code
func (o *CreateRepository22MethodNotAllowed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create repository22 method not allowed response has a 3xx status code
func (o *CreateRepository22MethodNotAllowed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository22 method not allowed response has a 4xx status code
func (o *CreateRepository22MethodNotAllowed) IsClientError() bool {
	return true
}

// IsServerError returns true when this create repository22 method not allowed response has a 5xx status code
func (o *CreateRepository22MethodNotAllowed) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository22 method not allowed response a status code equal to that given
func (o *CreateRepository22MethodNotAllowed) IsCode(code int) bool {
	return code == 405
}

// Code gets the status code for the create repository22 method not allowed response
func (o *CreateRepository22MethodNotAllowed) Code() int {
	return 405
}

func (o *CreateRepository22MethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/go/proxy][%d] createRepository22MethodNotAllowed", 405)
}

func (o *CreateRepository22MethodNotAllowed) String() string {
	return fmt.Sprintf("[POST /v1/repositories/go/proxy][%d] createRepository22MethodNotAllowed", 405)
}

func (o *CreateRepository22MethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
