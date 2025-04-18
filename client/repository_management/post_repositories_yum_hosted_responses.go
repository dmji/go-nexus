// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostRepositoriesYumHostedReader is a Reader for the PostRepositoriesYumHosted structure.
type PostRepositoriesYumHostedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRepositoriesYumHostedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostRepositoriesYumHostedCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPostRepositoriesYumHostedUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostRepositoriesYumHostedForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/repositories/yum/hosted] PostRepositoriesYumHosted", response, response.Code())
	}
}

// NewPostRepositoriesYumHostedCreated creates a PostRepositoriesYumHostedCreated with default headers values
func NewPostRepositoriesYumHostedCreated() *PostRepositoriesYumHostedCreated {
	return &PostRepositoriesYumHostedCreated{}
}

/*
PostRepositoriesYumHostedCreated describes a response with status code 201, with default header values.

Repository created
*/
type PostRepositoriesYumHostedCreated struct {
}

// IsSuccess returns true when this post repositories yum hosted created response has a 2xx status code
func (o *PostRepositoriesYumHostedCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post repositories yum hosted created response has a 3xx status code
func (o *PostRepositoriesYumHostedCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post repositories yum hosted created response has a 4xx status code
func (o *PostRepositoriesYumHostedCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post repositories yum hosted created response has a 5xx status code
func (o *PostRepositoriesYumHostedCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post repositories yum hosted created response a status code equal to that given
func (o *PostRepositoriesYumHostedCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the post repositories yum hosted created response
func (o *PostRepositoriesYumHostedCreated) Code() int {
	return 201
}

func (o *PostRepositoriesYumHostedCreated) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/yum/hosted][%d] postRepositoriesYumHostedCreated", 201)
}

func (o *PostRepositoriesYumHostedCreated) String() string {
	return fmt.Sprintf("[POST /v1/repositories/yum/hosted][%d] postRepositoriesYumHostedCreated", 201)
}

func (o *PostRepositoriesYumHostedCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostRepositoriesYumHostedUnauthorized creates a PostRepositoriesYumHostedUnauthorized with default headers values
func NewPostRepositoriesYumHostedUnauthorized() *PostRepositoriesYumHostedUnauthorized {
	return &PostRepositoriesYumHostedUnauthorized{}
}

/*
PostRepositoriesYumHostedUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type PostRepositoriesYumHostedUnauthorized struct {
}

// IsSuccess returns true when this post repositories yum hosted unauthorized response has a 2xx status code
func (o *PostRepositoriesYumHostedUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post repositories yum hosted unauthorized response has a 3xx status code
func (o *PostRepositoriesYumHostedUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post repositories yum hosted unauthorized response has a 4xx status code
func (o *PostRepositoriesYumHostedUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post repositories yum hosted unauthorized response has a 5xx status code
func (o *PostRepositoriesYumHostedUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post repositories yum hosted unauthorized response a status code equal to that given
func (o *PostRepositoriesYumHostedUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the post repositories yum hosted unauthorized response
func (o *PostRepositoriesYumHostedUnauthorized) Code() int {
	return 401
}

func (o *PostRepositoriesYumHostedUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/yum/hosted][%d] postRepositoriesYumHostedUnauthorized", 401)
}

func (o *PostRepositoriesYumHostedUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/repositories/yum/hosted][%d] postRepositoriesYumHostedUnauthorized", 401)
}

func (o *PostRepositoriesYumHostedUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostRepositoriesYumHostedForbidden creates a PostRepositoriesYumHostedForbidden with default headers values
func NewPostRepositoriesYumHostedForbidden() *PostRepositoriesYumHostedForbidden {
	return &PostRepositoriesYumHostedForbidden{}
}

/*
PostRepositoriesYumHostedForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type PostRepositoriesYumHostedForbidden struct {
}

// IsSuccess returns true when this post repositories yum hosted forbidden response has a 2xx status code
func (o *PostRepositoriesYumHostedForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post repositories yum hosted forbidden response has a 3xx status code
func (o *PostRepositoriesYumHostedForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post repositories yum hosted forbidden response has a 4xx status code
func (o *PostRepositoriesYumHostedForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post repositories yum hosted forbidden response has a 5xx status code
func (o *PostRepositoriesYumHostedForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post repositories yum hosted forbidden response a status code equal to that given
func (o *PostRepositoriesYumHostedForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the post repositories yum hosted forbidden response
func (o *PostRepositoriesYumHostedForbidden) Code() int {
	return 403
}

func (o *PostRepositoriesYumHostedForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/yum/hosted][%d] postRepositoriesYumHostedForbidden", 403)
}

func (o *PostRepositoriesYumHostedForbidden) String() string {
	return fmt.Sprintf("[POST /v1/repositories/yum/hosted][%d] postRepositoriesYumHostedForbidden", 403)
}

func (o *PostRepositoriesYumHostedForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
