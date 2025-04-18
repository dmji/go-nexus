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

// GetRepositoriesReader is a Reader for the GetRepositories structure.
type GetRepositoriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRepositoriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRepositoriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/repositories] GetRepositories", response, response.Code())
	}
}

// NewGetRepositoriesOK creates a GetRepositoriesOK with default headers values
func NewGetRepositoriesOK() *GetRepositoriesOK {
	return &GetRepositoriesOK{}
}

/*
GetRepositoriesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetRepositoriesOK struct {
	Payload []*models.RepositoryXO
}

// IsSuccess returns true when this get repositories o k response has a 2xx status code
func (o *GetRepositoriesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get repositories o k response has a 3xx status code
func (o *GetRepositoriesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get repositories o k response has a 4xx status code
func (o *GetRepositoriesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get repositories o k response has a 5xx status code
func (o *GetRepositoriesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get repositories o k response a status code equal to that given
func (o *GetRepositoriesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get repositories o k response
func (o *GetRepositoriesOK) Code() int {
	return 200
}

func (o *GetRepositoriesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/repositories][%d] getRepositoriesOK %s", 200, payload)
}

func (o *GetRepositoriesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/repositories][%d] getRepositoriesOK %s", 200, payload)
}

func (o *GetRepositoriesOK) GetPayload() []*models.RepositoryXO {
	return o.Payload
}

func (o *GetRepositoriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
