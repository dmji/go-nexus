// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostRepositoriesPypiHostedReader is a Reader for the PostRepositoriesPypiHosted structure.
type PostRepositoriesPypiHostedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRepositoriesPypiHostedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostRepositoriesPypiHostedCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPostRepositoriesPypiHostedUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostRepositoriesPypiHostedForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/repositories/pypi/hosted] PostRepositoriesPypiHosted", response, response.Code())
	}
}

// NewPostRepositoriesPypiHostedCreated creates a PostRepositoriesPypiHostedCreated with default headers values
func NewPostRepositoriesPypiHostedCreated() *PostRepositoriesPypiHostedCreated {
	return &PostRepositoriesPypiHostedCreated{}
}

/*
PostRepositoriesPypiHostedCreated describes a response with status code 201, with default header values.

Repository created
*/
type PostRepositoriesPypiHostedCreated struct {
}

// IsSuccess returns true when this post repositories pypi hosted created response has a 2xx status code
func (o *PostRepositoriesPypiHostedCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post repositories pypi hosted created response has a 3xx status code
func (o *PostRepositoriesPypiHostedCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post repositories pypi hosted created response has a 4xx status code
func (o *PostRepositoriesPypiHostedCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post repositories pypi hosted created response has a 5xx status code
func (o *PostRepositoriesPypiHostedCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post repositories pypi hosted created response a status code equal to that given
func (o *PostRepositoriesPypiHostedCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the post repositories pypi hosted created response
func (o *PostRepositoriesPypiHostedCreated) Code() int {
	return 201
}

func (o *PostRepositoriesPypiHostedCreated) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/pypi/hosted][%d] postRepositoriesPypiHostedCreated", 201)
}

func (o *PostRepositoriesPypiHostedCreated) String() string {
	return fmt.Sprintf("[POST /v1/repositories/pypi/hosted][%d] postRepositoriesPypiHostedCreated", 201)
}

func (o *PostRepositoriesPypiHostedCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostRepositoriesPypiHostedUnauthorized creates a PostRepositoriesPypiHostedUnauthorized with default headers values
func NewPostRepositoriesPypiHostedUnauthorized() *PostRepositoriesPypiHostedUnauthorized {
	return &PostRepositoriesPypiHostedUnauthorized{}
}

/*
PostRepositoriesPypiHostedUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type PostRepositoriesPypiHostedUnauthorized struct {
}

// IsSuccess returns true when this post repositories pypi hosted unauthorized response has a 2xx status code
func (o *PostRepositoriesPypiHostedUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post repositories pypi hosted unauthorized response has a 3xx status code
func (o *PostRepositoriesPypiHostedUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post repositories pypi hosted unauthorized response has a 4xx status code
func (o *PostRepositoriesPypiHostedUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post repositories pypi hosted unauthorized response has a 5xx status code
func (o *PostRepositoriesPypiHostedUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post repositories pypi hosted unauthorized response a status code equal to that given
func (o *PostRepositoriesPypiHostedUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the post repositories pypi hosted unauthorized response
func (o *PostRepositoriesPypiHostedUnauthorized) Code() int {
	return 401
}

func (o *PostRepositoriesPypiHostedUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/pypi/hosted][%d] postRepositoriesPypiHostedUnauthorized", 401)
}

func (o *PostRepositoriesPypiHostedUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/repositories/pypi/hosted][%d] postRepositoriesPypiHostedUnauthorized", 401)
}

func (o *PostRepositoriesPypiHostedUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostRepositoriesPypiHostedForbidden creates a PostRepositoriesPypiHostedForbidden with default headers values
func NewPostRepositoriesPypiHostedForbidden() *PostRepositoriesPypiHostedForbidden {
	return &PostRepositoriesPypiHostedForbidden{}
}

/*
PostRepositoriesPypiHostedForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type PostRepositoriesPypiHostedForbidden struct {
}

// IsSuccess returns true when this post repositories pypi hosted forbidden response has a 2xx status code
func (o *PostRepositoriesPypiHostedForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post repositories pypi hosted forbidden response has a 3xx status code
func (o *PostRepositoriesPypiHostedForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post repositories pypi hosted forbidden response has a 4xx status code
func (o *PostRepositoriesPypiHostedForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post repositories pypi hosted forbidden response has a 5xx status code
func (o *PostRepositoriesPypiHostedForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post repositories pypi hosted forbidden response a status code equal to that given
func (o *PostRepositoriesPypiHostedForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the post repositories pypi hosted forbidden response
func (o *PostRepositoriesPypiHostedForbidden) Code() int {
	return 403
}

func (o *PostRepositoriesPypiHostedForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/pypi/hosted][%d] postRepositoriesPypiHostedForbidden", 403)
}

func (o *PostRepositoriesPypiHostedForbidden) String() string {
	return fmt.Sprintf("[POST /v1/repositories/pypi/hosted][%d] postRepositoriesPypiHostedForbidden", 403)
}

func (o *PostRepositoriesPypiHostedForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
