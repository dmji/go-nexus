// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostRepositoriesDockerProxyReader is a Reader for the PostRepositoriesDockerProxy structure.
type PostRepositoriesDockerProxyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRepositoriesDockerProxyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostRepositoriesDockerProxyCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPostRepositoriesDockerProxyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostRepositoriesDockerProxyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/repositories/docker/proxy] PostRepositoriesDockerProxy", response, response.Code())
	}
}

// NewPostRepositoriesDockerProxyCreated creates a PostRepositoriesDockerProxyCreated with default headers values
func NewPostRepositoriesDockerProxyCreated() *PostRepositoriesDockerProxyCreated {
	return &PostRepositoriesDockerProxyCreated{}
}

/*
PostRepositoriesDockerProxyCreated describes a response with status code 201, with default header values.

Repository created
*/
type PostRepositoriesDockerProxyCreated struct {
}

// IsSuccess returns true when this post repositories docker proxy created response has a 2xx status code
func (o *PostRepositoriesDockerProxyCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post repositories docker proxy created response has a 3xx status code
func (o *PostRepositoriesDockerProxyCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post repositories docker proxy created response has a 4xx status code
func (o *PostRepositoriesDockerProxyCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post repositories docker proxy created response has a 5xx status code
func (o *PostRepositoriesDockerProxyCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post repositories docker proxy created response a status code equal to that given
func (o *PostRepositoriesDockerProxyCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the post repositories docker proxy created response
func (o *PostRepositoriesDockerProxyCreated) Code() int {
	return 201
}

func (o *PostRepositoriesDockerProxyCreated) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/docker/proxy][%d] postRepositoriesDockerProxyCreated", 201)
}

func (o *PostRepositoriesDockerProxyCreated) String() string {
	return fmt.Sprintf("[POST /v1/repositories/docker/proxy][%d] postRepositoriesDockerProxyCreated", 201)
}

func (o *PostRepositoriesDockerProxyCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostRepositoriesDockerProxyUnauthorized creates a PostRepositoriesDockerProxyUnauthorized with default headers values
func NewPostRepositoriesDockerProxyUnauthorized() *PostRepositoriesDockerProxyUnauthorized {
	return &PostRepositoriesDockerProxyUnauthorized{}
}

/*
PostRepositoriesDockerProxyUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type PostRepositoriesDockerProxyUnauthorized struct {
}

// IsSuccess returns true when this post repositories docker proxy unauthorized response has a 2xx status code
func (o *PostRepositoriesDockerProxyUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post repositories docker proxy unauthorized response has a 3xx status code
func (o *PostRepositoriesDockerProxyUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post repositories docker proxy unauthorized response has a 4xx status code
func (o *PostRepositoriesDockerProxyUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post repositories docker proxy unauthorized response has a 5xx status code
func (o *PostRepositoriesDockerProxyUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post repositories docker proxy unauthorized response a status code equal to that given
func (o *PostRepositoriesDockerProxyUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the post repositories docker proxy unauthorized response
func (o *PostRepositoriesDockerProxyUnauthorized) Code() int {
	return 401
}

func (o *PostRepositoriesDockerProxyUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/docker/proxy][%d] postRepositoriesDockerProxyUnauthorized", 401)
}

func (o *PostRepositoriesDockerProxyUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/repositories/docker/proxy][%d] postRepositoriesDockerProxyUnauthorized", 401)
}

func (o *PostRepositoriesDockerProxyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostRepositoriesDockerProxyForbidden creates a PostRepositoriesDockerProxyForbidden with default headers values
func NewPostRepositoriesDockerProxyForbidden() *PostRepositoriesDockerProxyForbidden {
	return &PostRepositoriesDockerProxyForbidden{}
}

/*
PostRepositoriesDockerProxyForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type PostRepositoriesDockerProxyForbidden struct {
}

// IsSuccess returns true when this post repositories docker proxy forbidden response has a 2xx status code
func (o *PostRepositoriesDockerProxyForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post repositories docker proxy forbidden response has a 3xx status code
func (o *PostRepositoriesDockerProxyForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post repositories docker proxy forbidden response has a 4xx status code
func (o *PostRepositoriesDockerProxyForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post repositories docker proxy forbidden response has a 5xx status code
func (o *PostRepositoriesDockerProxyForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post repositories docker proxy forbidden response a status code equal to that given
func (o *PostRepositoriesDockerProxyForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the post repositories docker proxy forbidden response
func (o *PostRepositoriesDockerProxyForbidden) Code() int {
	return 403
}

func (o *PostRepositoriesDockerProxyForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/docker/proxy][%d] postRepositoriesDockerProxyForbidden", 403)
}

func (o *PostRepositoriesDockerProxyForbidden) String() string {
	return fmt.Sprintf("[POST /v1/repositories/docker/proxy][%d] postRepositoriesDockerProxyForbidden", 403)
}

func (o *PostRepositoriesDockerProxyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
