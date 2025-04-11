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

// GetRepository31Reader is a Reader for the GetRepository31 structure.
type GetRepository31Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRepository31Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRepository31OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/repositories/nuget/hosted/{repositoryName}] getRepository_31", response, response.Code())
	}
}

// NewGetRepository31OK creates a GetRepository31OK with default headers values
func NewGetRepository31OK() *GetRepository31OK {
	return &GetRepository31OK{}
}

/*
GetRepository31OK describes a response with status code 200, with default header values.

successful operation
*/
type GetRepository31OK struct {
	Payload *models.SimpleAPIHostedRepository
}

// IsSuccess returns true when this get repository31 o k response has a 2xx status code
func (o *GetRepository31OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get repository31 o k response has a 3xx status code
func (o *GetRepository31OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get repository31 o k response has a 4xx status code
func (o *GetRepository31OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get repository31 o k response has a 5xx status code
func (o *GetRepository31OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get repository31 o k response a status code equal to that given
func (o *GetRepository31OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get repository31 o k response
func (o *GetRepository31OK) Code() int {
	return 200
}

func (o *GetRepository31OK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/repositories/nuget/hosted/{repositoryName}][%d] getRepository31OK %s", 200, payload)
}

func (o *GetRepository31OK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/repositories/nuget/hosted/{repositoryName}][%d] getRepository31OK %s", 200, payload)
}

func (o *GetRepository31OK) GetPayload() *models.SimpleAPIHostedRepository {
	return o.Payload
}

func (o *GetRepository31OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SimpleAPIHostedRepository)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
