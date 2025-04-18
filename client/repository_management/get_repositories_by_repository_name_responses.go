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

// GetRepositoriesByRepositoryNameReader is a Reader for the GetRepositoriesByRepositoryName structure.
type GetRepositoriesByRepositoryNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRepositoriesByRepositoryNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRepositoriesByRepositoryNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetRepositoriesByRepositoryNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRepositoriesByRepositoryNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRepositoriesByRepositoryNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/repositories/{repositoryName}] GetRepositoriesByRepositoryName", response, response.Code())
	}
}

// NewGetRepositoriesByRepositoryNameOK creates a GetRepositoriesByRepositoryNameOK with default headers values
func NewGetRepositoriesByRepositoryNameOK() *GetRepositoriesByRepositoryNameOK {
	return &GetRepositoriesByRepositoryNameOK{}
}

/*
GetRepositoriesByRepositoryNameOK describes a response with status code 200, with default header values.

successful operation
*/
type GetRepositoriesByRepositoryNameOK struct {
	Payload *models.RepositoryXO
}

// IsSuccess returns true when this get repositories by repository name o k response has a 2xx status code
func (o *GetRepositoriesByRepositoryNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get repositories by repository name o k response has a 3xx status code
func (o *GetRepositoriesByRepositoryNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get repositories by repository name o k response has a 4xx status code
func (o *GetRepositoriesByRepositoryNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get repositories by repository name o k response has a 5xx status code
func (o *GetRepositoriesByRepositoryNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get repositories by repository name o k response a status code equal to that given
func (o *GetRepositoriesByRepositoryNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get repositories by repository name o k response
func (o *GetRepositoriesByRepositoryNameOK) Code() int {
	return 200
}

func (o *GetRepositoriesByRepositoryNameOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/repositories/{repositoryName}][%d] getRepositoriesByRepositoryNameOK %s", 200, payload)
}

func (o *GetRepositoriesByRepositoryNameOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/repositories/{repositoryName}][%d] getRepositoriesByRepositoryNameOK %s", 200, payload)
}

func (o *GetRepositoriesByRepositoryNameOK) GetPayload() *models.RepositoryXO {
	return o.Payload
}

func (o *GetRepositoriesByRepositoryNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RepositoryXO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRepositoriesByRepositoryNameUnauthorized creates a GetRepositoriesByRepositoryNameUnauthorized with default headers values
func NewGetRepositoriesByRepositoryNameUnauthorized() *GetRepositoriesByRepositoryNameUnauthorized {
	return &GetRepositoriesByRepositoryNameUnauthorized{}
}

/*
GetRepositoriesByRepositoryNameUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type GetRepositoriesByRepositoryNameUnauthorized struct {
}

// IsSuccess returns true when this get repositories by repository name unauthorized response has a 2xx status code
func (o *GetRepositoriesByRepositoryNameUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get repositories by repository name unauthorized response has a 3xx status code
func (o *GetRepositoriesByRepositoryNameUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get repositories by repository name unauthorized response has a 4xx status code
func (o *GetRepositoriesByRepositoryNameUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get repositories by repository name unauthorized response has a 5xx status code
func (o *GetRepositoriesByRepositoryNameUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get repositories by repository name unauthorized response a status code equal to that given
func (o *GetRepositoriesByRepositoryNameUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get repositories by repository name unauthorized response
func (o *GetRepositoriesByRepositoryNameUnauthorized) Code() int {
	return 401
}

func (o *GetRepositoriesByRepositoryNameUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/repositories/{repositoryName}][%d] getRepositoriesByRepositoryNameUnauthorized", 401)
}

func (o *GetRepositoriesByRepositoryNameUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/repositories/{repositoryName}][%d] getRepositoriesByRepositoryNameUnauthorized", 401)
}

func (o *GetRepositoriesByRepositoryNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRepositoriesByRepositoryNameForbidden creates a GetRepositoriesByRepositoryNameForbidden with default headers values
func NewGetRepositoriesByRepositoryNameForbidden() *GetRepositoriesByRepositoryNameForbidden {
	return &GetRepositoriesByRepositoryNameForbidden{}
}

/*
GetRepositoriesByRepositoryNameForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type GetRepositoriesByRepositoryNameForbidden struct {
}

// IsSuccess returns true when this get repositories by repository name forbidden response has a 2xx status code
func (o *GetRepositoriesByRepositoryNameForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get repositories by repository name forbidden response has a 3xx status code
func (o *GetRepositoriesByRepositoryNameForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get repositories by repository name forbidden response has a 4xx status code
func (o *GetRepositoriesByRepositoryNameForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get repositories by repository name forbidden response has a 5xx status code
func (o *GetRepositoriesByRepositoryNameForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get repositories by repository name forbidden response a status code equal to that given
func (o *GetRepositoriesByRepositoryNameForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get repositories by repository name forbidden response
func (o *GetRepositoriesByRepositoryNameForbidden) Code() int {
	return 403
}

func (o *GetRepositoriesByRepositoryNameForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/repositories/{repositoryName}][%d] getRepositoriesByRepositoryNameForbidden", 403)
}

func (o *GetRepositoriesByRepositoryNameForbidden) String() string {
	return fmt.Sprintf("[GET /v1/repositories/{repositoryName}][%d] getRepositoriesByRepositoryNameForbidden", 403)
}

func (o *GetRepositoriesByRepositoryNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRepositoriesByRepositoryNameNotFound creates a GetRepositoriesByRepositoryNameNotFound with default headers values
func NewGetRepositoriesByRepositoryNameNotFound() *GetRepositoriesByRepositoryNameNotFound {
	return &GetRepositoriesByRepositoryNameNotFound{}
}

/*
GetRepositoriesByRepositoryNameNotFound describes a response with status code 404, with default header values.

Repository not found
*/
type GetRepositoriesByRepositoryNameNotFound struct {
}

// IsSuccess returns true when this get repositories by repository name not found response has a 2xx status code
func (o *GetRepositoriesByRepositoryNameNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get repositories by repository name not found response has a 3xx status code
func (o *GetRepositoriesByRepositoryNameNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get repositories by repository name not found response has a 4xx status code
func (o *GetRepositoriesByRepositoryNameNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get repositories by repository name not found response has a 5xx status code
func (o *GetRepositoriesByRepositoryNameNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get repositories by repository name not found response a status code equal to that given
func (o *GetRepositoriesByRepositoryNameNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get repositories by repository name not found response
func (o *GetRepositoriesByRepositoryNameNotFound) Code() int {
	return 404
}

func (o *GetRepositoriesByRepositoryNameNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/repositories/{repositoryName}][%d] getRepositoriesByRepositoryNameNotFound", 404)
}

func (o *GetRepositoriesByRepositoryNameNotFound) String() string {
	return fmt.Sprintf("[GET /v1/repositories/{repositoryName}][%d] getRepositoriesByRepositoryNameNotFound", 404)
}

func (o *GetRepositoriesByRepositoryNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
