// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostRepositoriesGoGroupReader is a Reader for the PostRepositoriesGoGroup structure.
type PostRepositoriesGoGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRepositoriesGoGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostRepositoriesGoGroupCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPostRepositoriesGoGroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostRepositoriesGoGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewPostRepositoriesGoGroupMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/repositories/go/group] PostRepositoriesGoGroup", response, response.Code())
	}
}

// NewPostRepositoriesGoGroupCreated creates a PostRepositoriesGoGroupCreated with default headers values
func NewPostRepositoriesGoGroupCreated() *PostRepositoriesGoGroupCreated {
	return &PostRepositoriesGoGroupCreated{}
}

/*
PostRepositoriesGoGroupCreated describes a response with status code 201, with default header values.

Repository created
*/
type PostRepositoriesGoGroupCreated struct {
}

// IsSuccess returns true when this post repositories go group created response has a 2xx status code
func (o *PostRepositoriesGoGroupCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post repositories go group created response has a 3xx status code
func (o *PostRepositoriesGoGroupCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post repositories go group created response has a 4xx status code
func (o *PostRepositoriesGoGroupCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post repositories go group created response has a 5xx status code
func (o *PostRepositoriesGoGroupCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post repositories go group created response a status code equal to that given
func (o *PostRepositoriesGoGroupCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the post repositories go group created response
func (o *PostRepositoriesGoGroupCreated) Code() int {
	return 201
}

func (o *PostRepositoriesGoGroupCreated) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/go/group][%d] postRepositoriesGoGroupCreated", 201)
}

func (o *PostRepositoriesGoGroupCreated) String() string {
	return fmt.Sprintf("[POST /v1/repositories/go/group][%d] postRepositoriesGoGroupCreated", 201)
}

func (o *PostRepositoriesGoGroupCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostRepositoriesGoGroupUnauthorized creates a PostRepositoriesGoGroupUnauthorized with default headers values
func NewPostRepositoriesGoGroupUnauthorized() *PostRepositoriesGoGroupUnauthorized {
	return &PostRepositoriesGoGroupUnauthorized{}
}

/*
PostRepositoriesGoGroupUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type PostRepositoriesGoGroupUnauthorized struct {
}

// IsSuccess returns true when this post repositories go group unauthorized response has a 2xx status code
func (o *PostRepositoriesGoGroupUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post repositories go group unauthorized response has a 3xx status code
func (o *PostRepositoriesGoGroupUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post repositories go group unauthorized response has a 4xx status code
func (o *PostRepositoriesGoGroupUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post repositories go group unauthorized response has a 5xx status code
func (o *PostRepositoriesGoGroupUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post repositories go group unauthorized response a status code equal to that given
func (o *PostRepositoriesGoGroupUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the post repositories go group unauthorized response
func (o *PostRepositoriesGoGroupUnauthorized) Code() int {
	return 401
}

func (o *PostRepositoriesGoGroupUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/go/group][%d] postRepositoriesGoGroupUnauthorized", 401)
}

func (o *PostRepositoriesGoGroupUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/repositories/go/group][%d] postRepositoriesGoGroupUnauthorized", 401)
}

func (o *PostRepositoriesGoGroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostRepositoriesGoGroupForbidden creates a PostRepositoriesGoGroupForbidden with default headers values
func NewPostRepositoriesGoGroupForbidden() *PostRepositoriesGoGroupForbidden {
	return &PostRepositoriesGoGroupForbidden{}
}

/*
PostRepositoriesGoGroupForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type PostRepositoriesGoGroupForbidden struct {
}

// IsSuccess returns true when this post repositories go group forbidden response has a 2xx status code
func (o *PostRepositoriesGoGroupForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post repositories go group forbidden response has a 3xx status code
func (o *PostRepositoriesGoGroupForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post repositories go group forbidden response has a 4xx status code
func (o *PostRepositoriesGoGroupForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post repositories go group forbidden response has a 5xx status code
func (o *PostRepositoriesGoGroupForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post repositories go group forbidden response a status code equal to that given
func (o *PostRepositoriesGoGroupForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the post repositories go group forbidden response
func (o *PostRepositoriesGoGroupForbidden) Code() int {
	return 403
}

func (o *PostRepositoriesGoGroupForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/go/group][%d] postRepositoriesGoGroupForbidden", 403)
}

func (o *PostRepositoriesGoGroupForbidden) String() string {
	return fmt.Sprintf("[POST /v1/repositories/go/group][%d] postRepositoriesGoGroupForbidden", 403)
}

func (o *PostRepositoriesGoGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostRepositoriesGoGroupMethodNotAllowed creates a PostRepositoriesGoGroupMethodNotAllowed with default headers values
func NewPostRepositoriesGoGroupMethodNotAllowed() *PostRepositoriesGoGroupMethodNotAllowed {
	return &PostRepositoriesGoGroupMethodNotAllowed{}
}

/*
PostRepositoriesGoGroupMethodNotAllowed describes a response with status code 405, with default header values.

Feature is disabled in High Availability
*/
type PostRepositoriesGoGroupMethodNotAllowed struct {
}

// IsSuccess returns true when this post repositories go group method not allowed response has a 2xx status code
func (o *PostRepositoriesGoGroupMethodNotAllowed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post repositories go group method not allowed response has a 3xx status code
func (o *PostRepositoriesGoGroupMethodNotAllowed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post repositories go group method not allowed response has a 4xx status code
func (o *PostRepositoriesGoGroupMethodNotAllowed) IsClientError() bool {
	return true
}

// IsServerError returns true when this post repositories go group method not allowed response has a 5xx status code
func (o *PostRepositoriesGoGroupMethodNotAllowed) IsServerError() bool {
	return false
}

// IsCode returns true when this post repositories go group method not allowed response a status code equal to that given
func (o *PostRepositoriesGoGroupMethodNotAllowed) IsCode(code int) bool {
	return code == 405
}

// Code gets the status code for the post repositories go group method not allowed response
func (o *PostRepositoriesGoGroupMethodNotAllowed) Code() int {
	return 405
}

func (o *PostRepositoriesGoGroupMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/go/group][%d] postRepositoriesGoGroupMethodNotAllowed", 405)
}

func (o *PostRepositoriesGoGroupMethodNotAllowed) String() string {
	return fmt.Sprintf("[POST /v1/repositories/go/group][%d] postRepositoriesGoGroupMethodNotAllowed", 405)
}

func (o *PostRepositoriesGoGroupMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
