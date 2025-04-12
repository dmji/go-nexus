// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostRepositoriesRubygemsProxyReader is a Reader for the PostRepositoriesRubygemsProxy structure.
type PostRepositoriesRubygemsProxyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRepositoriesRubygemsProxyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostRepositoriesRubygemsProxyCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPostRepositoriesRubygemsProxyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostRepositoriesRubygemsProxyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/repositories/rubygems/proxy] PostRepositoriesRubygemsProxy", response, response.Code())
	}
}

// NewPostRepositoriesRubygemsProxyCreated creates a PostRepositoriesRubygemsProxyCreated with default headers values
func NewPostRepositoriesRubygemsProxyCreated() *PostRepositoriesRubygemsProxyCreated {
	return &PostRepositoriesRubygemsProxyCreated{}
}

/*
PostRepositoriesRubygemsProxyCreated describes a response with status code 201, with default header values.

Repository created
*/
type PostRepositoriesRubygemsProxyCreated struct {
}

// IsSuccess returns true when this post repositories rubygems proxy created response has a 2xx status code
func (o *PostRepositoriesRubygemsProxyCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post repositories rubygems proxy created response has a 3xx status code
func (o *PostRepositoriesRubygemsProxyCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post repositories rubygems proxy created response has a 4xx status code
func (o *PostRepositoriesRubygemsProxyCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post repositories rubygems proxy created response has a 5xx status code
func (o *PostRepositoriesRubygemsProxyCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post repositories rubygems proxy created response a status code equal to that given
func (o *PostRepositoriesRubygemsProxyCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the post repositories rubygems proxy created response
func (o *PostRepositoriesRubygemsProxyCreated) Code() int {
	return 201
}

func (o *PostRepositoriesRubygemsProxyCreated) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/rubygems/proxy][%d] postRepositoriesRubygemsProxyCreated", 201)
}

func (o *PostRepositoriesRubygemsProxyCreated) String() string {
	return fmt.Sprintf("[POST /v1/repositories/rubygems/proxy][%d] postRepositoriesRubygemsProxyCreated", 201)
}

func (o *PostRepositoriesRubygemsProxyCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostRepositoriesRubygemsProxyUnauthorized creates a PostRepositoriesRubygemsProxyUnauthorized with default headers values
func NewPostRepositoriesRubygemsProxyUnauthorized() *PostRepositoriesRubygemsProxyUnauthorized {
	return &PostRepositoriesRubygemsProxyUnauthorized{}
}

/*
PostRepositoriesRubygemsProxyUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type PostRepositoriesRubygemsProxyUnauthorized struct {
}

// IsSuccess returns true when this post repositories rubygems proxy unauthorized response has a 2xx status code
func (o *PostRepositoriesRubygemsProxyUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post repositories rubygems proxy unauthorized response has a 3xx status code
func (o *PostRepositoriesRubygemsProxyUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post repositories rubygems proxy unauthorized response has a 4xx status code
func (o *PostRepositoriesRubygemsProxyUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post repositories rubygems proxy unauthorized response has a 5xx status code
func (o *PostRepositoriesRubygemsProxyUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post repositories rubygems proxy unauthorized response a status code equal to that given
func (o *PostRepositoriesRubygemsProxyUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the post repositories rubygems proxy unauthorized response
func (o *PostRepositoriesRubygemsProxyUnauthorized) Code() int {
	return 401
}

func (o *PostRepositoriesRubygemsProxyUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/rubygems/proxy][%d] postRepositoriesRubygemsProxyUnauthorized", 401)
}

func (o *PostRepositoriesRubygemsProxyUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/repositories/rubygems/proxy][%d] postRepositoriesRubygemsProxyUnauthorized", 401)
}

func (o *PostRepositoriesRubygemsProxyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostRepositoriesRubygemsProxyForbidden creates a PostRepositoriesRubygemsProxyForbidden with default headers values
func NewPostRepositoriesRubygemsProxyForbidden() *PostRepositoriesRubygemsProxyForbidden {
	return &PostRepositoriesRubygemsProxyForbidden{}
}

/*
PostRepositoriesRubygemsProxyForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type PostRepositoriesRubygemsProxyForbidden struct {
}

// IsSuccess returns true when this post repositories rubygems proxy forbidden response has a 2xx status code
func (o *PostRepositoriesRubygemsProxyForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post repositories rubygems proxy forbidden response has a 3xx status code
func (o *PostRepositoriesRubygemsProxyForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post repositories rubygems proxy forbidden response has a 4xx status code
func (o *PostRepositoriesRubygemsProxyForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post repositories rubygems proxy forbidden response has a 5xx status code
func (o *PostRepositoriesRubygemsProxyForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post repositories rubygems proxy forbidden response a status code equal to that given
func (o *PostRepositoriesRubygemsProxyForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the post repositories rubygems proxy forbidden response
func (o *PostRepositoriesRubygemsProxyForbidden) Code() int {
	return 403
}

func (o *PostRepositoriesRubygemsProxyForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/rubygems/proxy][%d] postRepositoriesRubygemsProxyForbidden", 403)
}

func (o *PostRepositoriesRubygemsProxyForbidden) String() string {
	return fmt.Sprintf("[POST /v1/repositories/rubygems/proxy][%d] postRepositoriesRubygemsProxyForbidden", 403)
}

func (o *PostRepositoriesRubygemsProxyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
