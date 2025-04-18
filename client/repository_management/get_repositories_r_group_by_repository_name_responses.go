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

// GetRepositoriesRGroupByRepositoryNameReader is a Reader for the GetRepositoriesRGroupByRepositoryName structure.
type GetRepositoriesRGroupByRepositoryNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRepositoriesRGroupByRepositoryNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRepositoriesRGroupByRepositoryNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/repositories/r/group/{repositoryName}] GetRepositoriesRGroupByRepositoryName", response, response.Code())
	}
}

// NewGetRepositoriesRGroupByRepositoryNameOK creates a GetRepositoriesRGroupByRepositoryNameOK with default headers values
func NewGetRepositoriesRGroupByRepositoryNameOK() *GetRepositoriesRGroupByRepositoryNameOK {
	return &GetRepositoriesRGroupByRepositoryNameOK{}
}

/*
GetRepositoriesRGroupByRepositoryNameOK describes a response with status code 200, with default header values.

successful operation
*/
type GetRepositoriesRGroupByRepositoryNameOK struct {
	Payload *models.SimpleAPIGroupRepository
}

// IsSuccess returns true when this get repositories r group by repository name o k response has a 2xx status code
func (o *GetRepositoriesRGroupByRepositoryNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get repositories r group by repository name o k response has a 3xx status code
func (o *GetRepositoriesRGroupByRepositoryNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get repositories r group by repository name o k response has a 4xx status code
func (o *GetRepositoriesRGroupByRepositoryNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get repositories r group by repository name o k response has a 5xx status code
func (o *GetRepositoriesRGroupByRepositoryNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get repositories r group by repository name o k response a status code equal to that given
func (o *GetRepositoriesRGroupByRepositoryNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get repositories r group by repository name o k response
func (o *GetRepositoriesRGroupByRepositoryNameOK) Code() int {
	return 200
}

func (o *GetRepositoriesRGroupByRepositoryNameOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/repositories/r/group/{repositoryName}][%d] getRepositoriesRGroupByRepositoryNameOK %s", 200, payload)
}

func (o *GetRepositoriesRGroupByRepositoryNameOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/repositories/r/group/{repositoryName}][%d] getRepositoriesRGroupByRepositoryNameOK %s", 200, payload)
}

func (o *GetRepositoriesRGroupByRepositoryNameOK) GetPayload() *models.SimpleAPIGroupRepository {
	return o.Payload
}

func (o *GetRepositoriesRGroupByRepositoryNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SimpleAPIGroupRepository)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
