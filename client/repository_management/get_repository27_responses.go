// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"nexus/models"
)

// GetRepository27Reader is a Reader for the GetRepository27 structure.
type GetRepository27Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRepository27Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRepository27OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/repositories/npm/group/{repositoryName}] getRepository_27", response, response.Code())
	}
}

// NewGetRepository27OK creates a GetRepository27OK with default headers values
func NewGetRepository27OK() *GetRepository27OK {
	return &GetRepository27OK{}
}

/*
GetRepository27OK describes a response with status code 200, with default header values.

successful operation
*/
type GetRepository27OK struct {
	Payload *models.SimpleAPIGroupDeployRepository
}

// IsSuccess returns true when this get repository27 o k response has a 2xx status code
func (o *GetRepository27OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get repository27 o k response has a 3xx status code
func (o *GetRepository27OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get repository27 o k response has a 4xx status code
func (o *GetRepository27OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get repository27 o k response has a 5xx status code
func (o *GetRepository27OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get repository27 o k response a status code equal to that given
func (o *GetRepository27OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get repository27 o k response
func (o *GetRepository27OK) Code() int {
	return 200
}

func (o *GetRepository27OK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/repositories/npm/group/{repositoryName}][%d] getRepository27OK %s", 200, payload)
}

func (o *GetRepository27OK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/repositories/npm/group/{repositoryName}][%d] getRepository27OK %s", 200, payload)
}

func (o *GetRepository27OK) GetPayload() *models.SimpleAPIGroupDeployRepository {
	return o.Payload
}

func (o *GetRepository27OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SimpleAPIGroupDeployRepository)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
