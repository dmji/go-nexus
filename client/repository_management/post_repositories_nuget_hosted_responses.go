// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostRepositoriesNugetHostedReader is a Reader for the PostRepositoriesNugetHosted structure.
type PostRepositoriesNugetHostedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRepositoriesNugetHostedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostRepositoriesNugetHostedCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPostRepositoriesNugetHostedUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostRepositoriesNugetHostedForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/repositories/nuget/hosted] PostRepositoriesNugetHosted", response, response.Code())
	}
}

// NewPostRepositoriesNugetHostedCreated creates a PostRepositoriesNugetHostedCreated with default headers values
func NewPostRepositoriesNugetHostedCreated() *PostRepositoriesNugetHostedCreated {
	return &PostRepositoriesNugetHostedCreated{}
}

/*
PostRepositoriesNugetHostedCreated describes a response with status code 201, with default header values.

Repository created
*/
type PostRepositoriesNugetHostedCreated struct {
}

// IsSuccess returns true when this post repositories nuget hosted created response has a 2xx status code
func (o *PostRepositoriesNugetHostedCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post repositories nuget hosted created response has a 3xx status code
func (o *PostRepositoriesNugetHostedCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post repositories nuget hosted created response has a 4xx status code
func (o *PostRepositoriesNugetHostedCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post repositories nuget hosted created response has a 5xx status code
func (o *PostRepositoriesNugetHostedCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post repositories nuget hosted created response a status code equal to that given
func (o *PostRepositoriesNugetHostedCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the post repositories nuget hosted created response
func (o *PostRepositoriesNugetHostedCreated) Code() int {
	return 201
}

func (o *PostRepositoriesNugetHostedCreated) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/nuget/hosted][%d] postRepositoriesNugetHostedCreated", 201)
}

func (o *PostRepositoriesNugetHostedCreated) String() string {
	return fmt.Sprintf("[POST /v1/repositories/nuget/hosted][%d] postRepositoriesNugetHostedCreated", 201)
}

func (o *PostRepositoriesNugetHostedCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostRepositoriesNugetHostedUnauthorized creates a PostRepositoriesNugetHostedUnauthorized with default headers values
func NewPostRepositoriesNugetHostedUnauthorized() *PostRepositoriesNugetHostedUnauthorized {
	return &PostRepositoriesNugetHostedUnauthorized{}
}

/*
PostRepositoriesNugetHostedUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type PostRepositoriesNugetHostedUnauthorized struct {
}

// IsSuccess returns true when this post repositories nuget hosted unauthorized response has a 2xx status code
func (o *PostRepositoriesNugetHostedUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post repositories nuget hosted unauthorized response has a 3xx status code
func (o *PostRepositoriesNugetHostedUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post repositories nuget hosted unauthorized response has a 4xx status code
func (o *PostRepositoriesNugetHostedUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post repositories nuget hosted unauthorized response has a 5xx status code
func (o *PostRepositoriesNugetHostedUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post repositories nuget hosted unauthorized response a status code equal to that given
func (o *PostRepositoriesNugetHostedUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the post repositories nuget hosted unauthorized response
func (o *PostRepositoriesNugetHostedUnauthorized) Code() int {
	return 401
}

func (o *PostRepositoriesNugetHostedUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/nuget/hosted][%d] postRepositoriesNugetHostedUnauthorized", 401)
}

func (o *PostRepositoriesNugetHostedUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/repositories/nuget/hosted][%d] postRepositoriesNugetHostedUnauthorized", 401)
}

func (o *PostRepositoriesNugetHostedUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostRepositoriesNugetHostedForbidden creates a PostRepositoriesNugetHostedForbidden with default headers values
func NewPostRepositoriesNugetHostedForbidden() *PostRepositoriesNugetHostedForbidden {
	return &PostRepositoriesNugetHostedForbidden{}
}

/*
PostRepositoriesNugetHostedForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type PostRepositoriesNugetHostedForbidden struct {
}

// IsSuccess returns true when this post repositories nuget hosted forbidden response has a 2xx status code
func (o *PostRepositoriesNugetHostedForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post repositories nuget hosted forbidden response has a 3xx status code
func (o *PostRepositoriesNugetHostedForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post repositories nuget hosted forbidden response has a 4xx status code
func (o *PostRepositoriesNugetHostedForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post repositories nuget hosted forbidden response has a 5xx status code
func (o *PostRepositoriesNugetHostedForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post repositories nuget hosted forbidden response a status code equal to that given
func (o *PostRepositoriesNugetHostedForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the post repositories nuget hosted forbidden response
func (o *PostRepositoriesNugetHostedForbidden) Code() int {
	return 403
}

func (o *PostRepositoriesNugetHostedForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/nuget/hosted][%d] postRepositoriesNugetHostedForbidden", 403)
}

func (o *PostRepositoriesNugetHostedForbidden) String() string {
	return fmt.Sprintf("[POST /v1/repositories/nuget/hosted][%d] postRepositoriesNugetHostedForbidden", 403)
}

func (o *PostRepositoriesNugetHostedForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
