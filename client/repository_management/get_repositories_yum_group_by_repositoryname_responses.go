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

// GetRepositoriesYumGroupByRepositorynameReader is a Reader for the GetRepositoriesYumGroupByRepositoryname structure.
type GetRepositoriesYumGroupByRepositorynameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRepositoriesYumGroupByRepositorynameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRepositoriesYumGroupByRepositorynameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/repositories/yum/group/{repositoryName}] GetRepositoriesYumGroupByRepositoryname", response, response.Code())
	}
}

// NewGetRepositoriesYumGroupByRepositorynameOK creates a GetRepositoriesYumGroupByRepositorynameOK with default headers values
func NewGetRepositoriesYumGroupByRepositorynameOK() *GetRepositoriesYumGroupByRepositorynameOK {
	return &GetRepositoriesYumGroupByRepositorynameOK{}
}

/*
GetRepositoriesYumGroupByRepositorynameOK describes a response with status code 200, with default header values.

successful operation
*/
type GetRepositoriesYumGroupByRepositorynameOK struct {
	Payload *models.SimpleAPIGroupRepository
}

// IsSuccess returns true when this get repositories yum group by repositoryname o k response has a 2xx status code
func (o *GetRepositoriesYumGroupByRepositorynameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get repositories yum group by repositoryname o k response has a 3xx status code
func (o *GetRepositoriesYumGroupByRepositorynameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get repositories yum group by repositoryname o k response has a 4xx status code
func (o *GetRepositoriesYumGroupByRepositorynameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get repositories yum group by repositoryname o k response has a 5xx status code
func (o *GetRepositoriesYumGroupByRepositorynameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get repositories yum group by repositoryname o k response a status code equal to that given
func (o *GetRepositoriesYumGroupByRepositorynameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get repositories yum group by repositoryname o k response
func (o *GetRepositoriesYumGroupByRepositorynameOK) Code() int {
	return 200
}

func (o *GetRepositoriesYumGroupByRepositorynameOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/repositories/yum/group/{repositoryName}][%d] getRepositoriesYumGroupByRepositorynameOK %s", 200, payload)
}

func (o *GetRepositoriesYumGroupByRepositorynameOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/repositories/yum/group/{repositoryName}][%d] getRepositoriesYumGroupByRepositorynameOK %s", 200, payload)
}

func (o *GetRepositoriesYumGroupByRepositorynameOK) GetPayload() *models.SimpleAPIGroupRepository {
	return o.Payload
}

func (o *GetRepositoriesYumGroupByRepositorynameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SimpleAPIGroupRepository)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
