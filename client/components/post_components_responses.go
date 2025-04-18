// Code generated by go-swagger; DO NOT EDIT.

package components

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostComponentsReader is a Reader for the PostComponents structure.
type PostComponentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostComponentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 403:
		result := NewPostComponentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPostComponentsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/components] PostComponents", response, response.Code())
	}
}

// NewPostComponentsForbidden creates a PostComponentsForbidden with default headers values
func NewPostComponentsForbidden() *PostComponentsForbidden {
	return &PostComponentsForbidden{}
}

/*
PostComponentsForbidden describes a response with status code 403, with default header values.

Insufficient permissions to upload a component
*/
type PostComponentsForbidden struct {
}

// IsSuccess returns true when this post components forbidden response has a 2xx status code
func (o *PostComponentsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post components forbidden response has a 3xx status code
func (o *PostComponentsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post components forbidden response has a 4xx status code
func (o *PostComponentsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post components forbidden response has a 5xx status code
func (o *PostComponentsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post components forbidden response a status code equal to that given
func (o *PostComponentsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the post components forbidden response
func (o *PostComponentsForbidden) Code() int {
	return 403
}

func (o *PostComponentsForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/components][%d] postComponentsForbidden", 403)
}

func (o *PostComponentsForbidden) String() string {
	return fmt.Sprintf("[POST /v1/components][%d] postComponentsForbidden", 403)
}

func (o *PostComponentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostComponentsUnprocessableEntity creates a PostComponentsUnprocessableEntity with default headers values
func NewPostComponentsUnprocessableEntity() *PostComponentsUnprocessableEntity {
	return &PostComponentsUnprocessableEntity{}
}

/*
PostComponentsUnprocessableEntity describes a response with status code 422, with default header values.

Parameter 'repository' is required
*/
type PostComponentsUnprocessableEntity struct {
}

// IsSuccess returns true when this post components unprocessable entity response has a 2xx status code
func (o *PostComponentsUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post components unprocessable entity response has a 3xx status code
func (o *PostComponentsUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post components unprocessable entity response has a 4xx status code
func (o *PostComponentsUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this post components unprocessable entity response has a 5xx status code
func (o *PostComponentsUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this post components unprocessable entity response a status code equal to that given
func (o *PostComponentsUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the post components unprocessable entity response
func (o *PostComponentsUnprocessableEntity) Code() int {
	return 422
}

func (o *PostComponentsUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /v1/components][%d] postComponentsUnprocessableEntity", 422)
}

func (o *PostComponentsUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /v1/components][%d] postComponentsUnprocessableEntity", 422)
}

func (o *PostComponentsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
