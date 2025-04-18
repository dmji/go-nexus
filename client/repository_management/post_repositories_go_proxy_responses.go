// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostRepositoriesGoProxyReader is a Reader for the PostRepositoriesGoProxy structure.
type PostRepositoriesGoProxyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRepositoriesGoProxyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostRepositoriesGoProxyCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPostRepositoriesGoProxyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostRepositoriesGoProxyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewPostRepositoriesGoProxyMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/repositories/go/proxy] PostRepositoriesGoProxy", response, response.Code())
	}
}

// NewPostRepositoriesGoProxyCreated creates a PostRepositoriesGoProxyCreated with default headers values
func NewPostRepositoriesGoProxyCreated() *PostRepositoriesGoProxyCreated {
	return &PostRepositoriesGoProxyCreated{}
}

/*
PostRepositoriesGoProxyCreated describes a response with status code 201, with default header values.

Repository created
*/
type PostRepositoriesGoProxyCreated struct {
}

// IsSuccess returns true when this post repositories go proxy created response has a 2xx status code
func (o *PostRepositoriesGoProxyCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post repositories go proxy created response has a 3xx status code
func (o *PostRepositoriesGoProxyCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post repositories go proxy created response has a 4xx status code
func (o *PostRepositoriesGoProxyCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post repositories go proxy created response has a 5xx status code
func (o *PostRepositoriesGoProxyCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post repositories go proxy created response a status code equal to that given
func (o *PostRepositoriesGoProxyCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the post repositories go proxy created response
func (o *PostRepositoriesGoProxyCreated) Code() int {
	return 201
}

func (o *PostRepositoriesGoProxyCreated) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/go/proxy][%d] postRepositoriesGoProxyCreated", 201)
}

func (o *PostRepositoriesGoProxyCreated) String() string {
	return fmt.Sprintf("[POST /v1/repositories/go/proxy][%d] postRepositoriesGoProxyCreated", 201)
}

func (o *PostRepositoriesGoProxyCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostRepositoriesGoProxyUnauthorized creates a PostRepositoriesGoProxyUnauthorized with default headers values
func NewPostRepositoriesGoProxyUnauthorized() *PostRepositoriesGoProxyUnauthorized {
	return &PostRepositoriesGoProxyUnauthorized{}
}

/*
PostRepositoriesGoProxyUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type PostRepositoriesGoProxyUnauthorized struct {
}

// IsSuccess returns true when this post repositories go proxy unauthorized response has a 2xx status code
func (o *PostRepositoriesGoProxyUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post repositories go proxy unauthorized response has a 3xx status code
func (o *PostRepositoriesGoProxyUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post repositories go proxy unauthorized response has a 4xx status code
func (o *PostRepositoriesGoProxyUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post repositories go proxy unauthorized response has a 5xx status code
func (o *PostRepositoriesGoProxyUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post repositories go proxy unauthorized response a status code equal to that given
func (o *PostRepositoriesGoProxyUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the post repositories go proxy unauthorized response
func (o *PostRepositoriesGoProxyUnauthorized) Code() int {
	return 401
}

func (o *PostRepositoriesGoProxyUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/go/proxy][%d] postRepositoriesGoProxyUnauthorized", 401)
}

func (o *PostRepositoriesGoProxyUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/repositories/go/proxy][%d] postRepositoriesGoProxyUnauthorized", 401)
}

func (o *PostRepositoriesGoProxyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostRepositoriesGoProxyForbidden creates a PostRepositoriesGoProxyForbidden with default headers values
func NewPostRepositoriesGoProxyForbidden() *PostRepositoriesGoProxyForbidden {
	return &PostRepositoriesGoProxyForbidden{}
}

/*
PostRepositoriesGoProxyForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type PostRepositoriesGoProxyForbidden struct {
}

// IsSuccess returns true when this post repositories go proxy forbidden response has a 2xx status code
func (o *PostRepositoriesGoProxyForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post repositories go proxy forbidden response has a 3xx status code
func (o *PostRepositoriesGoProxyForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post repositories go proxy forbidden response has a 4xx status code
func (o *PostRepositoriesGoProxyForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post repositories go proxy forbidden response has a 5xx status code
func (o *PostRepositoriesGoProxyForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post repositories go proxy forbidden response a status code equal to that given
func (o *PostRepositoriesGoProxyForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the post repositories go proxy forbidden response
func (o *PostRepositoriesGoProxyForbidden) Code() int {
	return 403
}

func (o *PostRepositoriesGoProxyForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/go/proxy][%d] postRepositoriesGoProxyForbidden", 403)
}

func (o *PostRepositoriesGoProxyForbidden) String() string {
	return fmt.Sprintf("[POST /v1/repositories/go/proxy][%d] postRepositoriesGoProxyForbidden", 403)
}

func (o *PostRepositoriesGoProxyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostRepositoriesGoProxyMethodNotAllowed creates a PostRepositoriesGoProxyMethodNotAllowed with default headers values
func NewPostRepositoriesGoProxyMethodNotAllowed() *PostRepositoriesGoProxyMethodNotAllowed {
	return &PostRepositoriesGoProxyMethodNotAllowed{}
}

/*
PostRepositoriesGoProxyMethodNotAllowed describes a response with status code 405, with default header values.

Feature is disabled in High Availability
*/
type PostRepositoriesGoProxyMethodNotAllowed struct {
}

// IsSuccess returns true when this post repositories go proxy method not allowed response has a 2xx status code
func (o *PostRepositoriesGoProxyMethodNotAllowed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post repositories go proxy method not allowed response has a 3xx status code
func (o *PostRepositoriesGoProxyMethodNotAllowed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post repositories go proxy method not allowed response has a 4xx status code
func (o *PostRepositoriesGoProxyMethodNotAllowed) IsClientError() bool {
	return true
}

// IsServerError returns true when this post repositories go proxy method not allowed response has a 5xx status code
func (o *PostRepositoriesGoProxyMethodNotAllowed) IsServerError() bool {
	return false
}

// IsCode returns true when this post repositories go proxy method not allowed response a status code equal to that given
func (o *PostRepositoriesGoProxyMethodNotAllowed) IsCode(code int) bool {
	return code == 405
}

// Code gets the status code for the post repositories go proxy method not allowed response
func (o *PostRepositoriesGoProxyMethodNotAllowed) Code() int {
	return 405
}

func (o *PostRepositoriesGoProxyMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/go/proxy][%d] postRepositoriesGoProxyMethodNotAllowed", 405)
}

func (o *PostRepositoriesGoProxyMethodNotAllowed) String() string {
	return fmt.Sprintf("[POST /v1/repositories/go/proxy][%d] postRepositoriesGoProxyMethodNotAllowed", 405)
}

func (o *PostRepositoriesGoProxyMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
