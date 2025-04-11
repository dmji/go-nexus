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

	"github.com/dmji/go-nexus/models"
)

// GetRepository22Reader is a Reader for the GetRepository22 structure.
type GetRepository22Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRepository22Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRepository22OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/repositories/go/group/{repositoryName}] getRepository_22", response, response.Code())
	}
}

// NewGetRepository22OK creates a GetRepository22OK with default headers values
func NewGetRepository22OK() *GetRepository22OK {
	return &GetRepository22OK{}
}

/*
GetRepository22OK describes a response with status code 200, with default header values.

successful operation
*/
type GetRepository22OK struct {
	Payload *models.SimpleAPIGroupRepository
}

// IsSuccess returns true when this get repository22 o k response has a 2xx status code
func (o *GetRepository22OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get repository22 o k response has a 3xx status code
func (o *GetRepository22OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get repository22 o k response has a 4xx status code
func (o *GetRepository22OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get repository22 o k response has a 5xx status code
func (o *GetRepository22OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get repository22 o k response a status code equal to that given
func (o *GetRepository22OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get repository22 o k response
func (o *GetRepository22OK) Code() int {
	return 200
}

func (o *GetRepository22OK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/repositories/go/group/{repositoryName}][%d] getRepository22OK %s", 200, payload)
}

func (o *GetRepository22OK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/repositories/go/group/{repositoryName}][%d] getRepository22OK %s", 200, payload)
}

func (o *GetRepository22OK) GetPayload() *models.SimpleAPIGroupRepository {
	return o.Payload
}

func (o *GetRepository22OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SimpleAPIGroupRepository)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
