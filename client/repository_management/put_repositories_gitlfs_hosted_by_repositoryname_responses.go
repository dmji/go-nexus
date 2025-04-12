// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PutRepositoriesGitlfsHostedByRepositorynameReader is a Reader for the PutRepositoriesGitlfsHostedByRepositoryname structure.
type PutRepositoriesGitlfsHostedByRepositorynameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutRepositoriesGitlfsHostedByRepositorynameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutRepositoriesGitlfsHostedByRepositorynameNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutRepositoriesGitlfsHostedByRepositorynameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutRepositoriesGitlfsHostedByRepositorynameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutRepositoriesGitlfsHostedByRepositorynameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /v1/repositories/gitlfs/hosted/{repositoryName}] PutRepositoriesGitlfsHostedByRepositoryname", response, response.Code())
	}
}

// NewPutRepositoriesGitlfsHostedByRepositorynameNoContent creates a PutRepositoriesGitlfsHostedByRepositorynameNoContent with default headers values
func NewPutRepositoriesGitlfsHostedByRepositorynameNoContent() *PutRepositoriesGitlfsHostedByRepositorynameNoContent {
	return &PutRepositoriesGitlfsHostedByRepositorynameNoContent{}
}

/*
PutRepositoriesGitlfsHostedByRepositorynameNoContent describes a response with status code 204, with default header values.

Repository updated
*/
type PutRepositoriesGitlfsHostedByRepositorynameNoContent struct {
}

// IsSuccess returns true when this put repositories gitlfs hosted by repositoryname no content response has a 2xx status code
func (o *PutRepositoriesGitlfsHostedByRepositorynameNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put repositories gitlfs hosted by repositoryname no content response has a 3xx status code
func (o *PutRepositoriesGitlfsHostedByRepositorynameNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories gitlfs hosted by repositoryname no content response has a 4xx status code
func (o *PutRepositoriesGitlfsHostedByRepositorynameNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this put repositories gitlfs hosted by repositoryname no content response has a 5xx status code
func (o *PutRepositoriesGitlfsHostedByRepositorynameNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories gitlfs hosted by repositoryname no content response a status code equal to that given
func (o *PutRepositoriesGitlfsHostedByRepositorynameNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the put repositories gitlfs hosted by repositoryname no content response
func (o *PutRepositoriesGitlfsHostedByRepositorynameNoContent) Code() int {
	return 204
}

func (o *PutRepositoriesGitlfsHostedByRepositorynameNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/gitlfs/hosted/{repositoryName}][%d] putRepositoriesGitlfsHostedByRepositorynameNoContent", 204)
}

func (o *PutRepositoriesGitlfsHostedByRepositorynameNoContent) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/gitlfs/hosted/{repositoryName}][%d] putRepositoriesGitlfsHostedByRepositorynameNoContent", 204)
}

func (o *PutRepositoriesGitlfsHostedByRepositorynameNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutRepositoriesGitlfsHostedByRepositorynameBadRequest creates a PutRepositoriesGitlfsHostedByRepositorynameBadRequest with default headers values
func NewPutRepositoriesGitlfsHostedByRepositorynameBadRequest() *PutRepositoriesGitlfsHostedByRepositorynameBadRequest {
	return &PutRepositoriesGitlfsHostedByRepositorynameBadRequest{}
}

/*
PutRepositoriesGitlfsHostedByRepositorynameBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PutRepositoriesGitlfsHostedByRepositorynameBadRequest struct {
}

// IsSuccess returns true when this put repositories gitlfs hosted by repositoryname bad request response has a 2xx status code
func (o *PutRepositoriesGitlfsHostedByRepositorynameBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put repositories gitlfs hosted by repositoryname bad request response has a 3xx status code
func (o *PutRepositoriesGitlfsHostedByRepositorynameBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories gitlfs hosted by repositoryname bad request response has a 4xx status code
func (o *PutRepositoriesGitlfsHostedByRepositorynameBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put repositories gitlfs hosted by repositoryname bad request response has a 5xx status code
func (o *PutRepositoriesGitlfsHostedByRepositorynameBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories gitlfs hosted by repositoryname bad request response a status code equal to that given
func (o *PutRepositoriesGitlfsHostedByRepositorynameBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the put repositories gitlfs hosted by repositoryname bad request response
func (o *PutRepositoriesGitlfsHostedByRepositorynameBadRequest) Code() int {
	return 400
}

func (o *PutRepositoriesGitlfsHostedByRepositorynameBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/gitlfs/hosted/{repositoryName}][%d] putRepositoriesGitlfsHostedByRepositorynameBadRequest", 400)
}

func (o *PutRepositoriesGitlfsHostedByRepositorynameBadRequest) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/gitlfs/hosted/{repositoryName}][%d] putRepositoriesGitlfsHostedByRepositorynameBadRequest", 400)
}

func (o *PutRepositoriesGitlfsHostedByRepositorynameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutRepositoriesGitlfsHostedByRepositorynameUnauthorized creates a PutRepositoriesGitlfsHostedByRepositorynameUnauthorized with default headers values
func NewPutRepositoriesGitlfsHostedByRepositorynameUnauthorized() *PutRepositoriesGitlfsHostedByRepositorynameUnauthorized {
	return &PutRepositoriesGitlfsHostedByRepositorynameUnauthorized{}
}

/*
PutRepositoriesGitlfsHostedByRepositorynameUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type PutRepositoriesGitlfsHostedByRepositorynameUnauthorized struct {
}

// IsSuccess returns true when this put repositories gitlfs hosted by repositoryname unauthorized response has a 2xx status code
func (o *PutRepositoriesGitlfsHostedByRepositorynameUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put repositories gitlfs hosted by repositoryname unauthorized response has a 3xx status code
func (o *PutRepositoriesGitlfsHostedByRepositorynameUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories gitlfs hosted by repositoryname unauthorized response has a 4xx status code
func (o *PutRepositoriesGitlfsHostedByRepositorynameUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put repositories gitlfs hosted by repositoryname unauthorized response has a 5xx status code
func (o *PutRepositoriesGitlfsHostedByRepositorynameUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories gitlfs hosted by repositoryname unauthorized response a status code equal to that given
func (o *PutRepositoriesGitlfsHostedByRepositorynameUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the put repositories gitlfs hosted by repositoryname unauthorized response
func (o *PutRepositoriesGitlfsHostedByRepositorynameUnauthorized) Code() int {
	return 401
}

func (o *PutRepositoriesGitlfsHostedByRepositorynameUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/gitlfs/hosted/{repositoryName}][%d] putRepositoriesGitlfsHostedByRepositorynameUnauthorized", 401)
}

func (o *PutRepositoriesGitlfsHostedByRepositorynameUnauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/gitlfs/hosted/{repositoryName}][%d] putRepositoriesGitlfsHostedByRepositorynameUnauthorized", 401)
}

func (o *PutRepositoriesGitlfsHostedByRepositorynameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutRepositoriesGitlfsHostedByRepositorynameForbidden creates a PutRepositoriesGitlfsHostedByRepositorynameForbidden with default headers values
func NewPutRepositoriesGitlfsHostedByRepositorynameForbidden() *PutRepositoriesGitlfsHostedByRepositorynameForbidden {
	return &PutRepositoriesGitlfsHostedByRepositorynameForbidden{}
}

/*
PutRepositoriesGitlfsHostedByRepositorynameForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type PutRepositoriesGitlfsHostedByRepositorynameForbidden struct {
}

// IsSuccess returns true when this put repositories gitlfs hosted by repositoryname forbidden response has a 2xx status code
func (o *PutRepositoriesGitlfsHostedByRepositorynameForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put repositories gitlfs hosted by repositoryname forbidden response has a 3xx status code
func (o *PutRepositoriesGitlfsHostedByRepositorynameForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories gitlfs hosted by repositoryname forbidden response has a 4xx status code
func (o *PutRepositoriesGitlfsHostedByRepositorynameForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put repositories gitlfs hosted by repositoryname forbidden response has a 5xx status code
func (o *PutRepositoriesGitlfsHostedByRepositorynameForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories gitlfs hosted by repositoryname forbidden response a status code equal to that given
func (o *PutRepositoriesGitlfsHostedByRepositorynameForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the put repositories gitlfs hosted by repositoryname forbidden response
func (o *PutRepositoriesGitlfsHostedByRepositorynameForbidden) Code() int {
	return 403
}

func (o *PutRepositoriesGitlfsHostedByRepositorynameForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/gitlfs/hosted/{repositoryName}][%d] putRepositoriesGitlfsHostedByRepositorynameForbidden", 403)
}

func (o *PutRepositoriesGitlfsHostedByRepositorynameForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/gitlfs/hosted/{repositoryName}][%d] putRepositoriesGitlfsHostedByRepositorynameForbidden", 403)
}

func (o *PutRepositoriesGitlfsHostedByRepositorynameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
