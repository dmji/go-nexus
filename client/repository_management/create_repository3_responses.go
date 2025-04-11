// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateRepository3Reader is a Reader for the CreateRepository3 structure.
type CreateRepository3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRepository3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateRepository3Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateRepository3Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateRepository3Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/repositories/maven/hosted] createRepository_3", response, response.Code())
	}
}

// NewCreateRepository3Created creates a CreateRepository3Created with default headers values
func NewCreateRepository3Created() *CreateRepository3Created {
	return &CreateRepository3Created{}
}

/*
CreateRepository3Created describes a response with status code 201, with default header values.

Repository created
*/
type CreateRepository3Created struct {
}

// IsSuccess returns true when this create repository3 created response has a 2xx status code
func (o *CreateRepository3Created) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create repository3 created response has a 3xx status code
func (o *CreateRepository3Created) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository3 created response has a 4xx status code
func (o *CreateRepository3Created) IsClientError() bool {
	return false
}

// IsServerError returns true when this create repository3 created response has a 5xx status code
func (o *CreateRepository3Created) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository3 created response a status code equal to that given
func (o *CreateRepository3Created) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create repository3 created response
func (o *CreateRepository3Created) Code() int {
	return 201
}

func (o *CreateRepository3Created) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/maven/hosted][%d] createRepository3Created", 201)
}

func (o *CreateRepository3Created) String() string {
	return fmt.Sprintf("[POST /v1/repositories/maven/hosted][%d] createRepository3Created", 201)
}

func (o *CreateRepository3Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRepository3Unauthorized creates a CreateRepository3Unauthorized with default headers values
func NewCreateRepository3Unauthorized() *CreateRepository3Unauthorized {
	return &CreateRepository3Unauthorized{}
}

/*
CreateRepository3Unauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type CreateRepository3Unauthorized struct {
}

// IsSuccess returns true when this create repository3 unauthorized response has a 2xx status code
func (o *CreateRepository3Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create repository3 unauthorized response has a 3xx status code
func (o *CreateRepository3Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository3 unauthorized response has a 4xx status code
func (o *CreateRepository3Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create repository3 unauthorized response has a 5xx status code
func (o *CreateRepository3Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository3 unauthorized response a status code equal to that given
func (o *CreateRepository3Unauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create repository3 unauthorized response
func (o *CreateRepository3Unauthorized) Code() int {
	return 401
}

func (o *CreateRepository3Unauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/maven/hosted][%d] createRepository3Unauthorized", 401)
}

func (o *CreateRepository3Unauthorized) String() string {
	return fmt.Sprintf("[POST /v1/repositories/maven/hosted][%d] createRepository3Unauthorized", 401)
}

func (o *CreateRepository3Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRepository3Forbidden creates a CreateRepository3Forbidden with default headers values
func NewCreateRepository3Forbidden() *CreateRepository3Forbidden {
	return &CreateRepository3Forbidden{}
}

/*
CreateRepository3Forbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type CreateRepository3Forbidden struct {
}

// IsSuccess returns true when this create repository3 forbidden response has a 2xx status code
func (o *CreateRepository3Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create repository3 forbidden response has a 3xx status code
func (o *CreateRepository3Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository3 forbidden response has a 4xx status code
func (o *CreateRepository3Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create repository3 forbidden response has a 5xx status code
func (o *CreateRepository3Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository3 forbidden response a status code equal to that given
func (o *CreateRepository3Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create repository3 forbidden response
func (o *CreateRepository3Forbidden) Code() int {
	return 403
}

func (o *CreateRepository3Forbidden) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/maven/hosted][%d] createRepository3Forbidden", 403)
}

func (o *CreateRepository3Forbidden) String() string {
	return fmt.Sprintf("[POST /v1/repositories/maven/hosted][%d] createRepository3Forbidden", 403)
}

func (o *CreateRepository3Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
