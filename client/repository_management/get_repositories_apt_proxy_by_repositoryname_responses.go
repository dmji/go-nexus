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

// GetRepositoriesAptProxyByRepositorynameReader is a Reader for the GetRepositoriesAptProxyByRepositoryname structure.
type GetRepositoriesAptProxyByRepositorynameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRepositoriesAptProxyByRepositorynameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRepositoriesAptProxyByRepositorynameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/repositories/apt/proxy/{repositoryName}] GetRepositoriesAptProxyByRepositoryname", response, response.Code())
	}
}

// NewGetRepositoriesAptProxyByRepositorynameOK creates a GetRepositoriesAptProxyByRepositorynameOK with default headers values
func NewGetRepositoriesAptProxyByRepositorynameOK() *GetRepositoriesAptProxyByRepositorynameOK {
	return &GetRepositoriesAptProxyByRepositorynameOK{}
}

/*
GetRepositoriesAptProxyByRepositorynameOK describes a response with status code 200, with default header values.

successful operation
*/
type GetRepositoriesAptProxyByRepositorynameOK struct {
	Payload *models.AptProxyAPIRepository
}

// IsSuccess returns true when this get repositories apt proxy by repositoryname o k response has a 2xx status code
func (o *GetRepositoriesAptProxyByRepositorynameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get repositories apt proxy by repositoryname o k response has a 3xx status code
func (o *GetRepositoriesAptProxyByRepositorynameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get repositories apt proxy by repositoryname o k response has a 4xx status code
func (o *GetRepositoriesAptProxyByRepositorynameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get repositories apt proxy by repositoryname o k response has a 5xx status code
func (o *GetRepositoriesAptProxyByRepositorynameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get repositories apt proxy by repositoryname o k response a status code equal to that given
func (o *GetRepositoriesAptProxyByRepositorynameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get repositories apt proxy by repositoryname o k response
func (o *GetRepositoriesAptProxyByRepositorynameOK) Code() int {
	return 200
}

func (o *GetRepositoriesAptProxyByRepositorynameOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/repositories/apt/proxy/{repositoryName}][%d] getRepositoriesAptProxyByRepositorynameOK %s", 200, payload)
}

func (o *GetRepositoriesAptProxyByRepositorynameOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/repositories/apt/proxy/{repositoryName}][%d] getRepositoriesAptProxyByRepositorynameOK %s", 200, payload)
}

func (o *GetRepositoriesAptProxyByRepositorynameOK) GetPayload() *models.AptProxyAPIRepository {
	return o.Payload
}

func (o *GetRepositoriesAptProxyByRepositorynameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AptProxyAPIRepository)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
