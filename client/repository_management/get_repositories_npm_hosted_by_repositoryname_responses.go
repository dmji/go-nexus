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

// GetRepositoriesNpmHostedByRepositorynameReader is a Reader for the GetRepositoriesNpmHostedByRepositoryname structure.
type GetRepositoriesNpmHostedByRepositorynameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRepositoriesNpmHostedByRepositorynameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRepositoriesNpmHostedByRepositorynameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/repositories/npm/hosted/{repositoryName}] GetRepositoriesNpmHostedByRepositoryname", response, response.Code())
	}
}

// NewGetRepositoriesNpmHostedByRepositorynameOK creates a GetRepositoriesNpmHostedByRepositorynameOK with default headers values
func NewGetRepositoriesNpmHostedByRepositorynameOK() *GetRepositoriesNpmHostedByRepositorynameOK {
	return &GetRepositoriesNpmHostedByRepositorynameOK{}
}

/*
GetRepositoriesNpmHostedByRepositorynameOK describes a response with status code 200, with default header values.

successful operation
*/
type GetRepositoriesNpmHostedByRepositorynameOK struct {
	Payload *models.SimpleAPIHostedRepository
}

// IsSuccess returns true when this get repositories npm hosted by repositoryname o k response has a 2xx status code
func (o *GetRepositoriesNpmHostedByRepositorynameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get repositories npm hosted by repositoryname o k response has a 3xx status code
func (o *GetRepositoriesNpmHostedByRepositorynameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get repositories npm hosted by repositoryname o k response has a 4xx status code
func (o *GetRepositoriesNpmHostedByRepositorynameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get repositories npm hosted by repositoryname o k response has a 5xx status code
func (o *GetRepositoriesNpmHostedByRepositorynameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get repositories npm hosted by repositoryname o k response a status code equal to that given
func (o *GetRepositoriesNpmHostedByRepositorynameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get repositories npm hosted by repositoryname o k response
func (o *GetRepositoriesNpmHostedByRepositorynameOK) Code() int {
	return 200
}

func (o *GetRepositoriesNpmHostedByRepositorynameOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/repositories/npm/hosted/{repositoryName}][%d] getRepositoriesNpmHostedByRepositorynameOK %s", 200, payload)
}

func (o *GetRepositoriesNpmHostedByRepositorynameOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/repositories/npm/hosted/{repositoryName}][%d] getRepositoriesNpmHostedByRepositorynameOK %s", 200, payload)
}

func (o *GetRepositoriesNpmHostedByRepositorynameOK) GetPayload() *models.SimpleAPIHostedRepository {
	return o.Payload
}

func (o *GetRepositoriesNpmHostedByRepositorynameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SimpleAPIHostedRepository)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
