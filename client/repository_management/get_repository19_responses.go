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

// GetRepository19Reader is a Reader for the GetRepository19 structure.
type GetRepository19Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRepository19Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRepository19OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/repositories/docker/hosted/{repositoryName}] getRepository_19", response, response.Code())
	}
}

// NewGetRepository19OK creates a GetRepository19OK with default headers values
func NewGetRepository19OK() *GetRepository19OK {
	return &GetRepository19OK{}
}

/*
GetRepository19OK describes a response with status code 200, with default header values.

successful operation
*/
type GetRepository19OK struct {
	Payload *models.DockerHostedAPIRepository
}

// IsSuccess returns true when this get repository19 o k response has a 2xx status code
func (o *GetRepository19OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get repository19 o k response has a 3xx status code
func (o *GetRepository19OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get repository19 o k response has a 4xx status code
func (o *GetRepository19OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get repository19 o k response has a 5xx status code
func (o *GetRepository19OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get repository19 o k response a status code equal to that given
func (o *GetRepository19OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get repository19 o k response
func (o *GetRepository19OK) Code() int {
	return 200
}

func (o *GetRepository19OK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/repositories/docker/hosted/{repositoryName}][%d] getRepository19OK %s", 200, payload)
}

func (o *GetRepository19OK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/repositories/docker/hosted/{repositoryName}][%d] getRepository19OK %s", 200, payload)
}

func (o *GetRepository19OK) GetPayload() *models.DockerHostedAPIRepository {
	return o.Payload
}

func (o *GetRepository19OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DockerHostedAPIRepository)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
