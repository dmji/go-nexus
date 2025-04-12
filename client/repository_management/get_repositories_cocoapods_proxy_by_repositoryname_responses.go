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

// GetRepositoriesCocoapodsProxyByRepositorynameReader is a Reader for the GetRepositoriesCocoapodsProxyByRepositoryname structure.
type GetRepositoriesCocoapodsProxyByRepositorynameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRepositoriesCocoapodsProxyByRepositorynameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRepositoriesCocoapodsProxyByRepositorynameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/repositories/cocoapods/proxy/{repositoryName}] GetRepositoriesCocoapodsProxyByRepositoryname", response, response.Code())
	}
}

// NewGetRepositoriesCocoapodsProxyByRepositorynameOK creates a GetRepositoriesCocoapodsProxyByRepositorynameOK with default headers values
func NewGetRepositoriesCocoapodsProxyByRepositorynameOK() *GetRepositoriesCocoapodsProxyByRepositorynameOK {
	return &GetRepositoriesCocoapodsProxyByRepositorynameOK{}
}

/*
GetRepositoriesCocoapodsProxyByRepositorynameOK describes a response with status code 200, with default header values.

successful operation
*/
type GetRepositoriesCocoapodsProxyByRepositorynameOK struct {
	Payload *models.SimpleAPIProxyRepository
}

// IsSuccess returns true when this get repositories cocoapods proxy by repositoryname o k response has a 2xx status code
func (o *GetRepositoriesCocoapodsProxyByRepositorynameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get repositories cocoapods proxy by repositoryname o k response has a 3xx status code
func (o *GetRepositoriesCocoapodsProxyByRepositorynameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get repositories cocoapods proxy by repositoryname o k response has a 4xx status code
func (o *GetRepositoriesCocoapodsProxyByRepositorynameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get repositories cocoapods proxy by repositoryname o k response has a 5xx status code
func (o *GetRepositoriesCocoapodsProxyByRepositorynameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get repositories cocoapods proxy by repositoryname o k response a status code equal to that given
func (o *GetRepositoriesCocoapodsProxyByRepositorynameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get repositories cocoapods proxy by repositoryname o k response
func (o *GetRepositoriesCocoapodsProxyByRepositorynameOK) Code() int {
	return 200
}

func (o *GetRepositoriesCocoapodsProxyByRepositorynameOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/repositories/cocoapods/proxy/{repositoryName}][%d] getRepositoriesCocoapodsProxyByRepositorynameOK %s", 200, payload)
}

func (o *GetRepositoriesCocoapodsProxyByRepositorynameOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/repositories/cocoapods/proxy/{repositoryName}][%d] getRepositoriesCocoapodsProxyByRepositorynameOK %s", 200, payload)
}

func (o *GetRepositoriesCocoapodsProxyByRepositorynameOK) GetPayload() *models.SimpleAPIProxyRepository {
	return o.Payload
}

func (o *GetRepositoriesCocoapodsProxyByRepositorynameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SimpleAPIProxyRepository)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
