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

// GetRepositoriesRubygemsProxyByRepositorynameReader is a Reader for the GetRepositoriesRubygemsProxyByRepositoryname structure.
type GetRepositoriesRubygemsProxyByRepositorynameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRepositoriesRubygemsProxyByRepositorynameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRepositoriesRubygemsProxyByRepositorynameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/repositories/rubygems/proxy/{repositoryName}] GetRepositoriesRubygemsProxyByRepositoryname", response, response.Code())
	}
}

// NewGetRepositoriesRubygemsProxyByRepositorynameOK creates a GetRepositoriesRubygemsProxyByRepositorynameOK with default headers values
func NewGetRepositoriesRubygemsProxyByRepositorynameOK() *GetRepositoriesRubygemsProxyByRepositorynameOK {
	return &GetRepositoriesRubygemsProxyByRepositorynameOK{}
}

/*
GetRepositoriesRubygemsProxyByRepositorynameOK describes a response with status code 200, with default header values.

successful operation
*/
type GetRepositoriesRubygemsProxyByRepositorynameOK struct {
	Payload *models.SimpleAPIProxyRepository
}

// IsSuccess returns true when this get repositories rubygems proxy by repositoryname o k response has a 2xx status code
func (o *GetRepositoriesRubygemsProxyByRepositorynameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get repositories rubygems proxy by repositoryname o k response has a 3xx status code
func (o *GetRepositoriesRubygemsProxyByRepositorynameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get repositories rubygems proxy by repositoryname o k response has a 4xx status code
func (o *GetRepositoriesRubygemsProxyByRepositorynameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get repositories rubygems proxy by repositoryname o k response has a 5xx status code
func (o *GetRepositoriesRubygemsProxyByRepositorynameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get repositories rubygems proxy by repositoryname o k response a status code equal to that given
func (o *GetRepositoriesRubygemsProxyByRepositorynameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get repositories rubygems proxy by repositoryname o k response
func (o *GetRepositoriesRubygemsProxyByRepositorynameOK) Code() int {
	return 200
}

func (o *GetRepositoriesRubygemsProxyByRepositorynameOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/repositories/rubygems/proxy/{repositoryName}][%d] getRepositoriesRubygemsProxyByRepositorynameOK %s", 200, payload)
}

func (o *GetRepositoriesRubygemsProxyByRepositorynameOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/repositories/rubygems/proxy/{repositoryName}][%d] getRepositoriesRubygemsProxyByRepositorynameOK %s", 200, payload)
}

func (o *GetRepositoriesRubygemsProxyByRepositorynameOK) GetPayload() *models.SimpleAPIProxyRepository {
	return o.Payload
}

func (o *GetRepositoriesRubygemsProxyByRepositorynameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SimpleAPIProxyRepository)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
