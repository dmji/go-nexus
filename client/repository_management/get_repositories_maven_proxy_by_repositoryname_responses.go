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

// GetRepositoriesMavenProxyByRepositorynameReader is a Reader for the GetRepositoriesMavenProxyByRepositoryname structure.
type GetRepositoriesMavenProxyByRepositorynameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRepositoriesMavenProxyByRepositorynameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRepositoriesMavenProxyByRepositorynameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/repositories/maven/proxy/{repositoryName}] GetRepositoriesMavenProxyByRepositoryname", response, response.Code())
	}
}

// NewGetRepositoriesMavenProxyByRepositorynameOK creates a GetRepositoriesMavenProxyByRepositorynameOK with default headers values
func NewGetRepositoriesMavenProxyByRepositorynameOK() *GetRepositoriesMavenProxyByRepositorynameOK {
	return &GetRepositoriesMavenProxyByRepositorynameOK{}
}

/*
GetRepositoriesMavenProxyByRepositorynameOK describes a response with status code 200, with default header values.

successful operation
*/
type GetRepositoriesMavenProxyByRepositorynameOK struct {
	Payload *models.MavenProxyAPIRepository
}

// IsSuccess returns true when this get repositories maven proxy by repositoryname o k response has a 2xx status code
func (o *GetRepositoriesMavenProxyByRepositorynameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get repositories maven proxy by repositoryname o k response has a 3xx status code
func (o *GetRepositoriesMavenProxyByRepositorynameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get repositories maven proxy by repositoryname o k response has a 4xx status code
func (o *GetRepositoriesMavenProxyByRepositorynameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get repositories maven proxy by repositoryname o k response has a 5xx status code
func (o *GetRepositoriesMavenProxyByRepositorynameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get repositories maven proxy by repositoryname o k response a status code equal to that given
func (o *GetRepositoriesMavenProxyByRepositorynameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get repositories maven proxy by repositoryname o k response
func (o *GetRepositoriesMavenProxyByRepositorynameOK) Code() int {
	return 200
}

func (o *GetRepositoriesMavenProxyByRepositorynameOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/repositories/maven/proxy/{repositoryName}][%d] getRepositoriesMavenProxyByRepositorynameOK %s", 200, payload)
}

func (o *GetRepositoriesMavenProxyByRepositorynameOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/repositories/maven/proxy/{repositoryName}][%d] getRepositoriesMavenProxyByRepositorynameOK %s", 200, payload)
}

func (o *GetRepositoriesMavenProxyByRepositorynameOK) GetPayload() *models.MavenProxyAPIRepository {
	return o.Payload
}

func (o *GetRepositoriesMavenProxyByRepositorynameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MavenProxyAPIRepository)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
