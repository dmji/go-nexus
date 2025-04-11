// Code generated by go-swagger; DO NOT EDIT.

package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/dmji/go-nexus/models"
)

// GetTaskByIDReader is a Reader for the GetTaskByID structure.
type GetTaskByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTaskByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTaskByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetTaskByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/tasks/{id}] getTaskById", response, response.Code())
	}
}

// NewGetTaskByIDOK creates a GetTaskByIDOK with default headers values
func NewGetTaskByIDOK() *GetTaskByIDOK {
	return &GetTaskByIDOK{}
}

/*
GetTaskByIDOK describes a response with status code 200, with default header values.

successful operation
*/
type GetTaskByIDOK struct {
	Payload *models.TaskXO
}

// IsSuccess returns true when this get task by Id o k response has a 2xx status code
func (o *GetTaskByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get task by Id o k response has a 3xx status code
func (o *GetTaskByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get task by Id o k response has a 4xx status code
func (o *GetTaskByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get task by Id o k response has a 5xx status code
func (o *GetTaskByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get task by Id o k response a status code equal to that given
func (o *GetTaskByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get task by Id o k response
func (o *GetTaskByIDOK) Code() int {
	return 200
}

func (o *GetTaskByIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/tasks/{id}][%d] getTaskByIdOK %s", 200, payload)
}

func (o *GetTaskByIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/tasks/{id}][%d] getTaskByIdOK %s", 200, payload)
}

func (o *GetTaskByIDOK) GetPayload() *models.TaskXO {
	return o.Payload
}

func (o *GetTaskByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TaskXO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTaskByIDNotFound creates a GetTaskByIDNotFound with default headers values
func NewGetTaskByIDNotFound() *GetTaskByIDNotFound {
	return &GetTaskByIDNotFound{}
}

/*
GetTaskByIDNotFound describes a response with status code 404, with default header values.

Task not found
*/
type GetTaskByIDNotFound struct {
}

// IsSuccess returns true when this get task by Id not found response has a 2xx status code
func (o *GetTaskByIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get task by Id not found response has a 3xx status code
func (o *GetTaskByIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get task by Id not found response has a 4xx status code
func (o *GetTaskByIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get task by Id not found response has a 5xx status code
func (o *GetTaskByIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get task by Id not found response a status code equal to that given
func (o *GetTaskByIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get task by Id not found response
func (o *GetTaskByIDNotFound) Code() int {
	return 404
}

func (o *GetTaskByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/tasks/{id}][%d] getTaskByIdNotFound", 404)
}

func (o *GetTaskByIDNotFound) String() string {
	return fmt.Sprintf("[GET /v1/tasks/{id}][%d] getTaskByIdNotFound", 404)
}

func (o *GetTaskByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
