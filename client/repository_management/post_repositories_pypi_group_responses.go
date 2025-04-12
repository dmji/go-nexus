// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostRepositoriesPypiGroupReader is a Reader for the PostRepositoriesPypiGroup structure.
type PostRepositoriesPypiGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRepositoriesPypiGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostRepositoriesPypiGroupCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPostRepositoriesPypiGroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostRepositoriesPypiGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/repositories/pypi/group] PostRepositoriesPypiGroup", response, response.Code())
	}
}

// NewPostRepositoriesPypiGroupCreated creates a PostRepositoriesPypiGroupCreated with default headers values
func NewPostRepositoriesPypiGroupCreated() *PostRepositoriesPypiGroupCreated {
	return &PostRepositoriesPypiGroupCreated{}
}

/*
PostRepositoriesPypiGroupCreated describes a response with status code 201, with default header values.

Repository created
*/
type PostRepositoriesPypiGroupCreated struct {
}

// IsSuccess returns true when this post repositories pypi group created response has a 2xx status code
func (o *PostRepositoriesPypiGroupCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post repositories pypi group created response has a 3xx status code
func (o *PostRepositoriesPypiGroupCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post repositories pypi group created response has a 4xx status code
func (o *PostRepositoriesPypiGroupCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post repositories pypi group created response has a 5xx status code
func (o *PostRepositoriesPypiGroupCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post repositories pypi group created response a status code equal to that given
func (o *PostRepositoriesPypiGroupCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the post repositories pypi group created response
func (o *PostRepositoriesPypiGroupCreated) Code() int {
	return 201
}

func (o *PostRepositoriesPypiGroupCreated) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/pypi/group][%d] postRepositoriesPypiGroupCreated", 201)
}

func (o *PostRepositoriesPypiGroupCreated) String() string {
	return fmt.Sprintf("[POST /v1/repositories/pypi/group][%d] postRepositoriesPypiGroupCreated", 201)
}

func (o *PostRepositoriesPypiGroupCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostRepositoriesPypiGroupUnauthorized creates a PostRepositoriesPypiGroupUnauthorized with default headers values
func NewPostRepositoriesPypiGroupUnauthorized() *PostRepositoriesPypiGroupUnauthorized {
	return &PostRepositoriesPypiGroupUnauthorized{}
}

/*
PostRepositoriesPypiGroupUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type PostRepositoriesPypiGroupUnauthorized struct {
}

// IsSuccess returns true when this post repositories pypi group unauthorized response has a 2xx status code
func (o *PostRepositoriesPypiGroupUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post repositories pypi group unauthorized response has a 3xx status code
func (o *PostRepositoriesPypiGroupUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post repositories pypi group unauthorized response has a 4xx status code
func (o *PostRepositoriesPypiGroupUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post repositories pypi group unauthorized response has a 5xx status code
func (o *PostRepositoriesPypiGroupUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post repositories pypi group unauthorized response a status code equal to that given
func (o *PostRepositoriesPypiGroupUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the post repositories pypi group unauthorized response
func (o *PostRepositoriesPypiGroupUnauthorized) Code() int {
	return 401
}

func (o *PostRepositoriesPypiGroupUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/pypi/group][%d] postRepositoriesPypiGroupUnauthorized", 401)
}

func (o *PostRepositoriesPypiGroupUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/repositories/pypi/group][%d] postRepositoriesPypiGroupUnauthorized", 401)
}

func (o *PostRepositoriesPypiGroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostRepositoriesPypiGroupForbidden creates a PostRepositoriesPypiGroupForbidden with default headers values
func NewPostRepositoriesPypiGroupForbidden() *PostRepositoriesPypiGroupForbidden {
	return &PostRepositoriesPypiGroupForbidden{}
}

/*
PostRepositoriesPypiGroupForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type PostRepositoriesPypiGroupForbidden struct {
}

// IsSuccess returns true when this post repositories pypi group forbidden response has a 2xx status code
func (o *PostRepositoriesPypiGroupForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post repositories pypi group forbidden response has a 3xx status code
func (o *PostRepositoriesPypiGroupForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post repositories pypi group forbidden response has a 4xx status code
func (o *PostRepositoriesPypiGroupForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post repositories pypi group forbidden response has a 5xx status code
func (o *PostRepositoriesPypiGroupForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post repositories pypi group forbidden response a status code equal to that given
func (o *PostRepositoriesPypiGroupForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the post repositories pypi group forbidden response
func (o *PostRepositoriesPypiGroupForbidden) Code() int {
	return 403
}

func (o *PostRepositoriesPypiGroupForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/pypi/group][%d] postRepositoriesPypiGroupForbidden", 403)
}

func (o *PostRepositoriesPypiGroupForbidden) String() string {
	return fmt.Sprintf("[POST /v1/repositories/pypi/group][%d] postRepositoriesPypiGroupForbidden", 403)
}

func (o *PostRepositoriesPypiGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
