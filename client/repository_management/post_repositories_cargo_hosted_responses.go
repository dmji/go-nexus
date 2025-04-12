// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostRepositoriesCargoHostedReader is a Reader for the PostRepositoriesCargoHosted structure.
type PostRepositoriesCargoHostedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRepositoriesCargoHostedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostRepositoriesCargoHostedCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPostRepositoriesCargoHostedUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostRepositoriesCargoHostedForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/repositories/cargo/hosted] PostRepositoriesCargoHosted", response, response.Code())
	}
}

// NewPostRepositoriesCargoHostedCreated creates a PostRepositoriesCargoHostedCreated with default headers values
func NewPostRepositoriesCargoHostedCreated() *PostRepositoriesCargoHostedCreated {
	return &PostRepositoriesCargoHostedCreated{}
}

/*
PostRepositoriesCargoHostedCreated describes a response with status code 201, with default header values.

Repository created
*/
type PostRepositoriesCargoHostedCreated struct {
}

// IsSuccess returns true when this post repositories cargo hosted created response has a 2xx status code
func (o *PostRepositoriesCargoHostedCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post repositories cargo hosted created response has a 3xx status code
func (o *PostRepositoriesCargoHostedCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post repositories cargo hosted created response has a 4xx status code
func (o *PostRepositoriesCargoHostedCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post repositories cargo hosted created response has a 5xx status code
func (o *PostRepositoriesCargoHostedCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post repositories cargo hosted created response a status code equal to that given
func (o *PostRepositoriesCargoHostedCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the post repositories cargo hosted created response
func (o *PostRepositoriesCargoHostedCreated) Code() int {
	return 201
}

func (o *PostRepositoriesCargoHostedCreated) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/cargo/hosted][%d] postRepositoriesCargoHostedCreated", 201)
}

func (o *PostRepositoriesCargoHostedCreated) String() string {
	return fmt.Sprintf("[POST /v1/repositories/cargo/hosted][%d] postRepositoriesCargoHostedCreated", 201)
}

func (o *PostRepositoriesCargoHostedCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostRepositoriesCargoHostedUnauthorized creates a PostRepositoriesCargoHostedUnauthorized with default headers values
func NewPostRepositoriesCargoHostedUnauthorized() *PostRepositoriesCargoHostedUnauthorized {
	return &PostRepositoriesCargoHostedUnauthorized{}
}

/*
PostRepositoriesCargoHostedUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type PostRepositoriesCargoHostedUnauthorized struct {
}

// IsSuccess returns true when this post repositories cargo hosted unauthorized response has a 2xx status code
func (o *PostRepositoriesCargoHostedUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post repositories cargo hosted unauthorized response has a 3xx status code
func (o *PostRepositoriesCargoHostedUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post repositories cargo hosted unauthorized response has a 4xx status code
func (o *PostRepositoriesCargoHostedUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post repositories cargo hosted unauthorized response has a 5xx status code
func (o *PostRepositoriesCargoHostedUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post repositories cargo hosted unauthorized response a status code equal to that given
func (o *PostRepositoriesCargoHostedUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the post repositories cargo hosted unauthorized response
func (o *PostRepositoriesCargoHostedUnauthorized) Code() int {
	return 401
}

func (o *PostRepositoriesCargoHostedUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/cargo/hosted][%d] postRepositoriesCargoHostedUnauthorized", 401)
}

func (o *PostRepositoriesCargoHostedUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/repositories/cargo/hosted][%d] postRepositoriesCargoHostedUnauthorized", 401)
}

func (o *PostRepositoriesCargoHostedUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostRepositoriesCargoHostedForbidden creates a PostRepositoriesCargoHostedForbidden with default headers values
func NewPostRepositoriesCargoHostedForbidden() *PostRepositoriesCargoHostedForbidden {
	return &PostRepositoriesCargoHostedForbidden{}
}

/*
PostRepositoriesCargoHostedForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type PostRepositoriesCargoHostedForbidden struct {
}

// IsSuccess returns true when this post repositories cargo hosted forbidden response has a 2xx status code
func (o *PostRepositoriesCargoHostedForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post repositories cargo hosted forbidden response has a 3xx status code
func (o *PostRepositoriesCargoHostedForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post repositories cargo hosted forbidden response has a 4xx status code
func (o *PostRepositoriesCargoHostedForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post repositories cargo hosted forbidden response has a 5xx status code
func (o *PostRepositoriesCargoHostedForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post repositories cargo hosted forbidden response a status code equal to that given
func (o *PostRepositoriesCargoHostedForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the post repositories cargo hosted forbidden response
func (o *PostRepositoriesCargoHostedForbidden) Code() int {
	return 403
}

func (o *PostRepositoriesCargoHostedForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/cargo/hosted][%d] postRepositoriesCargoHostedForbidden", 403)
}

func (o *PostRepositoriesCargoHostedForbidden) String() string {
	return fmt.Sprintf("[POST /v1/repositories/cargo/hosted][%d] postRepositoriesCargoHostedForbidden", 403)
}

func (o *PostRepositoriesCargoHostedForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
