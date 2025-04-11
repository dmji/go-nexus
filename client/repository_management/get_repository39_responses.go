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

// GetRepository39Reader is a Reader for the GetRepository39 structure.
type GetRepository39Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRepository39Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRepository39OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/repositories/r/proxy/{repositoryName}] getRepository_39", response, response.Code())
	}
}

// NewGetRepository39OK creates a GetRepository39OK with default headers values
func NewGetRepository39OK() *GetRepository39OK {
	return &GetRepository39OK{}
}

/*
GetRepository39OK describes a response with status code 200, with default header values.

successful operation
*/
type GetRepository39OK struct {
	Payload *models.SimpleAPIProxyRepository
}

// IsSuccess returns true when this get repository39 o k response has a 2xx status code
func (o *GetRepository39OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get repository39 o k response has a 3xx status code
func (o *GetRepository39OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get repository39 o k response has a 4xx status code
func (o *GetRepository39OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get repository39 o k response has a 5xx status code
func (o *GetRepository39OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get repository39 o k response a status code equal to that given
func (o *GetRepository39OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get repository39 o k response
func (o *GetRepository39OK) Code() int {
	return 200
}

func (o *GetRepository39OK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/repositories/r/proxy/{repositoryName}][%d] getRepository39OK %s", 200, payload)
}

func (o *GetRepository39OK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/repositories/r/proxy/{repositoryName}][%d] getRepository39OK %s", 200, payload)
}

func (o *GetRepository39OK) GetPayload() *models.SimpleAPIProxyRepository {
	return o.Payload
}

func (o *GetRepository39OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SimpleAPIProxyRepository)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
