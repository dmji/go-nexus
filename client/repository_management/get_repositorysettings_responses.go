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

// GetRepositorysettingsReader is a Reader for the GetRepositorysettings structure.
type GetRepositorysettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRepositorysettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRepositorysettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetRepositorysettingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRepositorysettingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/repositorySettings] GetRepositorysettings", response, response.Code())
	}
}

// NewGetRepositorysettingsOK creates a GetRepositorysettingsOK with default headers values
func NewGetRepositorysettingsOK() *GetRepositorysettingsOK {
	return &GetRepositorysettingsOK{}
}

/*
GetRepositorysettingsOK describes a response with status code 200, with default header values.

Repositories list returned
*/
type GetRepositorysettingsOK struct {
	Payload []*models.AbstractAPIRepository
}

// IsSuccess returns true when this get repositorysettings o k response has a 2xx status code
func (o *GetRepositorysettingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get repositorysettings o k response has a 3xx status code
func (o *GetRepositorysettingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get repositorysettings o k response has a 4xx status code
func (o *GetRepositorysettingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get repositorysettings o k response has a 5xx status code
func (o *GetRepositorysettingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get repositorysettings o k response a status code equal to that given
func (o *GetRepositorysettingsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get repositorysettings o k response
func (o *GetRepositorysettingsOK) Code() int {
	return 200
}

func (o *GetRepositorysettingsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/repositorySettings][%d] getRepositorysettingsOK %s", 200, payload)
}

func (o *GetRepositorysettingsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/repositorySettings][%d] getRepositorysettingsOK %s", 200, payload)
}

func (o *GetRepositorysettingsOK) GetPayload() []*models.AbstractAPIRepository {
	return o.Payload
}

func (o *GetRepositorysettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRepositorysettingsUnauthorized creates a GetRepositorysettingsUnauthorized with default headers values
func NewGetRepositorysettingsUnauthorized() *GetRepositorysettingsUnauthorized {
	return &GetRepositorysettingsUnauthorized{}
}

/*
GetRepositorysettingsUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type GetRepositorysettingsUnauthorized struct {
}

// IsSuccess returns true when this get repositorysettings unauthorized response has a 2xx status code
func (o *GetRepositorysettingsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get repositorysettings unauthorized response has a 3xx status code
func (o *GetRepositorysettingsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get repositorysettings unauthorized response has a 4xx status code
func (o *GetRepositorysettingsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get repositorysettings unauthorized response has a 5xx status code
func (o *GetRepositorysettingsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get repositorysettings unauthorized response a status code equal to that given
func (o *GetRepositorysettingsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get repositorysettings unauthorized response
func (o *GetRepositorysettingsUnauthorized) Code() int {
	return 401
}

func (o *GetRepositorysettingsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/repositorySettings][%d] getRepositorysettingsUnauthorized", 401)
}

func (o *GetRepositorysettingsUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/repositorySettings][%d] getRepositorysettingsUnauthorized", 401)
}

func (o *GetRepositorysettingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRepositorysettingsForbidden creates a GetRepositorysettingsForbidden with default headers values
func NewGetRepositorysettingsForbidden() *GetRepositorysettingsForbidden {
	return &GetRepositorysettingsForbidden{}
}

/*
GetRepositorysettingsForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type GetRepositorysettingsForbidden struct {
}

// IsSuccess returns true when this get repositorysettings forbidden response has a 2xx status code
func (o *GetRepositorysettingsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get repositorysettings forbidden response has a 3xx status code
func (o *GetRepositorysettingsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get repositorysettings forbidden response has a 4xx status code
func (o *GetRepositorysettingsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get repositorysettings forbidden response has a 5xx status code
func (o *GetRepositorysettingsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get repositorysettings forbidden response a status code equal to that given
func (o *GetRepositorysettingsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get repositorysettings forbidden response
func (o *GetRepositorysettingsForbidden) Code() int {
	return 403
}

func (o *GetRepositorysettingsForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/repositorySettings][%d] getRepositorysettingsForbidden", 403)
}

func (o *GetRepositorysettingsForbidden) String() string {
	return fmt.Sprintf("[GET /v1/repositorySettings][%d] getRepositorysettingsForbidden", 403)
}

func (o *GetRepositorysettingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
