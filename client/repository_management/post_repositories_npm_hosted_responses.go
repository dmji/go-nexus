// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostRepositoriesNpmHostedReader is a Reader for the PostRepositoriesNpmHosted structure.
type PostRepositoriesNpmHostedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRepositoriesNpmHostedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostRepositoriesNpmHostedCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPostRepositoriesNpmHostedUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostRepositoriesNpmHostedForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/repositories/npm/hosted] PostRepositoriesNpmHosted", response, response.Code())
	}
}

// NewPostRepositoriesNpmHostedCreated creates a PostRepositoriesNpmHostedCreated with default headers values
func NewPostRepositoriesNpmHostedCreated() *PostRepositoriesNpmHostedCreated {
	return &PostRepositoriesNpmHostedCreated{}
}

/*
PostRepositoriesNpmHostedCreated describes a response with status code 201, with default header values.

Repository created
*/
type PostRepositoriesNpmHostedCreated struct {
}

// IsSuccess returns true when this post repositories npm hosted created response has a 2xx status code
func (o *PostRepositoriesNpmHostedCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post repositories npm hosted created response has a 3xx status code
func (o *PostRepositoriesNpmHostedCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post repositories npm hosted created response has a 4xx status code
func (o *PostRepositoriesNpmHostedCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post repositories npm hosted created response has a 5xx status code
func (o *PostRepositoriesNpmHostedCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post repositories npm hosted created response a status code equal to that given
func (o *PostRepositoriesNpmHostedCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the post repositories npm hosted created response
func (o *PostRepositoriesNpmHostedCreated) Code() int {
	return 201
}

func (o *PostRepositoriesNpmHostedCreated) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/npm/hosted][%d] postRepositoriesNpmHostedCreated", 201)
}

func (o *PostRepositoriesNpmHostedCreated) String() string {
	return fmt.Sprintf("[POST /v1/repositories/npm/hosted][%d] postRepositoriesNpmHostedCreated", 201)
}

func (o *PostRepositoriesNpmHostedCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostRepositoriesNpmHostedUnauthorized creates a PostRepositoriesNpmHostedUnauthorized with default headers values
func NewPostRepositoriesNpmHostedUnauthorized() *PostRepositoriesNpmHostedUnauthorized {
	return &PostRepositoriesNpmHostedUnauthorized{}
}

/*
PostRepositoriesNpmHostedUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type PostRepositoriesNpmHostedUnauthorized struct {
}

// IsSuccess returns true when this post repositories npm hosted unauthorized response has a 2xx status code
func (o *PostRepositoriesNpmHostedUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post repositories npm hosted unauthorized response has a 3xx status code
func (o *PostRepositoriesNpmHostedUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post repositories npm hosted unauthorized response has a 4xx status code
func (o *PostRepositoriesNpmHostedUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post repositories npm hosted unauthorized response has a 5xx status code
func (o *PostRepositoriesNpmHostedUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post repositories npm hosted unauthorized response a status code equal to that given
func (o *PostRepositoriesNpmHostedUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the post repositories npm hosted unauthorized response
func (o *PostRepositoriesNpmHostedUnauthorized) Code() int {
	return 401
}

func (o *PostRepositoriesNpmHostedUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/npm/hosted][%d] postRepositoriesNpmHostedUnauthorized", 401)
}

func (o *PostRepositoriesNpmHostedUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/repositories/npm/hosted][%d] postRepositoriesNpmHostedUnauthorized", 401)
}

func (o *PostRepositoriesNpmHostedUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostRepositoriesNpmHostedForbidden creates a PostRepositoriesNpmHostedForbidden with default headers values
func NewPostRepositoriesNpmHostedForbidden() *PostRepositoriesNpmHostedForbidden {
	return &PostRepositoriesNpmHostedForbidden{}
}

/*
PostRepositoriesNpmHostedForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type PostRepositoriesNpmHostedForbidden struct {
}

// IsSuccess returns true when this post repositories npm hosted forbidden response has a 2xx status code
func (o *PostRepositoriesNpmHostedForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post repositories npm hosted forbidden response has a 3xx status code
func (o *PostRepositoriesNpmHostedForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post repositories npm hosted forbidden response has a 4xx status code
func (o *PostRepositoriesNpmHostedForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post repositories npm hosted forbidden response has a 5xx status code
func (o *PostRepositoriesNpmHostedForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post repositories npm hosted forbidden response a status code equal to that given
func (o *PostRepositoriesNpmHostedForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the post repositories npm hosted forbidden response
func (o *PostRepositoriesNpmHostedForbidden) Code() int {
	return 403
}

func (o *PostRepositoriesNpmHostedForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/npm/hosted][%d] postRepositoriesNpmHostedForbidden", 403)
}

func (o *PostRepositoriesNpmHostedForbidden) String() string {
	return fmt.Sprintf("[POST /v1/repositories/npm/hosted][%d] postRepositoriesNpmHostedForbidden", 403)
}

func (o *PostRepositoriesNpmHostedForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
