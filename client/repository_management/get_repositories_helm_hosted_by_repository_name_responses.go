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

// GetRepositoriesHelmHostedByRepositoryNameReader is a Reader for the GetRepositoriesHelmHostedByRepositoryName structure.
type GetRepositoriesHelmHostedByRepositoryNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRepositoriesHelmHostedByRepositoryNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRepositoriesHelmHostedByRepositoryNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/repositories/helm/hosted/{repositoryName}] GetRepositoriesHelmHostedByRepositoryName", response, response.Code())
	}
}

// NewGetRepositoriesHelmHostedByRepositoryNameOK creates a GetRepositoriesHelmHostedByRepositoryNameOK with default headers values
func NewGetRepositoriesHelmHostedByRepositoryNameOK() *GetRepositoriesHelmHostedByRepositoryNameOK {
	return &GetRepositoriesHelmHostedByRepositoryNameOK{}
}

/*
GetRepositoriesHelmHostedByRepositoryNameOK describes a response with status code 200, with default header values.

successful operation
*/
type GetRepositoriesHelmHostedByRepositoryNameOK struct {
	Payload *models.SimpleAPIHostedRepository
}

// IsSuccess returns true when this get repositories helm hosted by repository name o k response has a 2xx status code
func (o *GetRepositoriesHelmHostedByRepositoryNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get repositories helm hosted by repository name o k response has a 3xx status code
func (o *GetRepositoriesHelmHostedByRepositoryNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get repositories helm hosted by repository name o k response has a 4xx status code
func (o *GetRepositoriesHelmHostedByRepositoryNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get repositories helm hosted by repository name o k response has a 5xx status code
func (o *GetRepositoriesHelmHostedByRepositoryNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get repositories helm hosted by repository name o k response a status code equal to that given
func (o *GetRepositoriesHelmHostedByRepositoryNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get repositories helm hosted by repository name o k response
func (o *GetRepositoriesHelmHostedByRepositoryNameOK) Code() int {
	return 200
}

func (o *GetRepositoriesHelmHostedByRepositoryNameOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/repositories/helm/hosted/{repositoryName}][%d] getRepositoriesHelmHostedByRepositoryNameOK %s", 200, payload)
}

func (o *GetRepositoriesHelmHostedByRepositoryNameOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/repositories/helm/hosted/{repositoryName}][%d] getRepositoriesHelmHostedByRepositoryNameOK %s", 200, payload)
}

func (o *GetRepositoriesHelmHostedByRepositoryNameOK) GetPayload() *models.SimpleAPIHostedRepository {
	return o.Payload
}

func (o *GetRepositoriesHelmHostedByRepositoryNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SimpleAPIHostedRepository)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
