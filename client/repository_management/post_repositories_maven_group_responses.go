// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostRepositoriesMavenGroupReader is a Reader for the PostRepositoriesMavenGroup structure.
type PostRepositoriesMavenGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRepositoriesMavenGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostRepositoriesMavenGroupCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPostRepositoriesMavenGroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostRepositoriesMavenGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/repositories/maven/group] PostRepositoriesMavenGroup", response, response.Code())
	}
}

// NewPostRepositoriesMavenGroupCreated creates a PostRepositoriesMavenGroupCreated with default headers values
func NewPostRepositoriesMavenGroupCreated() *PostRepositoriesMavenGroupCreated {
	return &PostRepositoriesMavenGroupCreated{}
}

/*
PostRepositoriesMavenGroupCreated describes a response with status code 201, with default header values.

Repository created
*/
type PostRepositoriesMavenGroupCreated struct {
}

// IsSuccess returns true when this post repositories maven group created response has a 2xx status code
func (o *PostRepositoriesMavenGroupCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post repositories maven group created response has a 3xx status code
func (o *PostRepositoriesMavenGroupCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post repositories maven group created response has a 4xx status code
func (o *PostRepositoriesMavenGroupCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post repositories maven group created response has a 5xx status code
func (o *PostRepositoriesMavenGroupCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post repositories maven group created response a status code equal to that given
func (o *PostRepositoriesMavenGroupCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the post repositories maven group created response
func (o *PostRepositoriesMavenGroupCreated) Code() int {
	return 201
}

func (o *PostRepositoriesMavenGroupCreated) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/maven/group][%d] postRepositoriesMavenGroupCreated", 201)
}

func (o *PostRepositoriesMavenGroupCreated) String() string {
	return fmt.Sprintf("[POST /v1/repositories/maven/group][%d] postRepositoriesMavenGroupCreated", 201)
}

func (o *PostRepositoriesMavenGroupCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostRepositoriesMavenGroupUnauthorized creates a PostRepositoriesMavenGroupUnauthorized with default headers values
func NewPostRepositoriesMavenGroupUnauthorized() *PostRepositoriesMavenGroupUnauthorized {
	return &PostRepositoriesMavenGroupUnauthorized{}
}

/*
PostRepositoriesMavenGroupUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type PostRepositoriesMavenGroupUnauthorized struct {
}

// IsSuccess returns true when this post repositories maven group unauthorized response has a 2xx status code
func (o *PostRepositoriesMavenGroupUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post repositories maven group unauthorized response has a 3xx status code
func (o *PostRepositoriesMavenGroupUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post repositories maven group unauthorized response has a 4xx status code
func (o *PostRepositoriesMavenGroupUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post repositories maven group unauthorized response has a 5xx status code
func (o *PostRepositoriesMavenGroupUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post repositories maven group unauthorized response a status code equal to that given
func (o *PostRepositoriesMavenGroupUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the post repositories maven group unauthorized response
func (o *PostRepositoriesMavenGroupUnauthorized) Code() int {
	return 401
}

func (o *PostRepositoriesMavenGroupUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/maven/group][%d] postRepositoriesMavenGroupUnauthorized", 401)
}

func (o *PostRepositoriesMavenGroupUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/repositories/maven/group][%d] postRepositoriesMavenGroupUnauthorized", 401)
}

func (o *PostRepositoriesMavenGroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostRepositoriesMavenGroupForbidden creates a PostRepositoriesMavenGroupForbidden with default headers values
func NewPostRepositoriesMavenGroupForbidden() *PostRepositoriesMavenGroupForbidden {
	return &PostRepositoriesMavenGroupForbidden{}
}

/*
PostRepositoriesMavenGroupForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type PostRepositoriesMavenGroupForbidden struct {
}

// IsSuccess returns true when this post repositories maven group forbidden response has a 2xx status code
func (o *PostRepositoriesMavenGroupForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post repositories maven group forbidden response has a 3xx status code
func (o *PostRepositoriesMavenGroupForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post repositories maven group forbidden response has a 4xx status code
func (o *PostRepositoriesMavenGroupForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post repositories maven group forbidden response has a 5xx status code
func (o *PostRepositoriesMavenGroupForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post repositories maven group forbidden response a status code equal to that given
func (o *PostRepositoriesMavenGroupForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the post repositories maven group forbidden response
func (o *PostRepositoriesMavenGroupForbidden) Code() int {
	return 403
}

func (o *PostRepositoriesMavenGroupForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/maven/group][%d] postRepositoriesMavenGroupForbidden", 403)
}

func (o *PostRepositoriesMavenGroupForbidden) String() string {
	return fmt.Sprintf("[POST /v1/repositories/maven/group][%d] postRepositoriesMavenGroupForbidden", 403)
}

func (o *PostRepositoriesMavenGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
