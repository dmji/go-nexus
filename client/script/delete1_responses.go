// Code generated by go-swagger; DO NOT EDIT.

package script

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// Delete1Reader is a Reader for the Delete1 structure.
type Delete1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Delete1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDelete1NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDelete1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /v1/script/{name}] delete_1", response, response.Code())
	}
}

// NewDelete1NoContent creates a Delete1NoContent with default headers values
func NewDelete1NoContent() *Delete1NoContent {
	return &Delete1NoContent{}
}

/*
Delete1NoContent describes a response with status code 204, with default header values.

Script was deleted
*/
type Delete1NoContent struct {
}

// IsSuccess returns true when this delete1 no content response has a 2xx status code
func (o *Delete1NoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete1 no content response has a 3xx status code
func (o *Delete1NoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete1 no content response has a 4xx status code
func (o *Delete1NoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete1 no content response has a 5xx status code
func (o *Delete1NoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete1 no content response a status code equal to that given
func (o *Delete1NoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete1 no content response
func (o *Delete1NoContent) Code() int {
	return 204
}

func (o *Delete1NoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/script/{name}][%d] delete1NoContent", 204)
}

func (o *Delete1NoContent) String() string {
	return fmt.Sprintf("[DELETE /v1/script/{name}][%d] delete1NoContent", 204)
}

func (o *Delete1NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDelete1NotFound creates a Delete1NotFound with default headers values
func NewDelete1NotFound() *Delete1NotFound {
	return &Delete1NotFound{}
}

/*
Delete1NotFound describes a response with status code 404, with default header values.

No script with the specified name
*/
type Delete1NotFound struct {
}

// IsSuccess returns true when this delete1 not found response has a 2xx status code
func (o *Delete1NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete1 not found response has a 3xx status code
func (o *Delete1NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete1 not found response has a 4xx status code
func (o *Delete1NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete1 not found response has a 5xx status code
func (o *Delete1NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete1 not found response a status code equal to that given
func (o *Delete1NotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete1 not found response
func (o *Delete1NotFound) Code() int {
	return 404
}

func (o *Delete1NotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/script/{name}][%d] delete1NotFound", 404)
}

func (o *Delete1NotFound) String() string {
	return fmt.Sprintf("[DELETE /v1/script/{name}][%d] delete1NotFound", 404)
}

func (o *Delete1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
