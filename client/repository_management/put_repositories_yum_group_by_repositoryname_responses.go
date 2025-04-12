// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PutRepositoriesYumGroupByRepositorynameReader is a Reader for the PutRepositoriesYumGroupByRepositoryname structure.
type PutRepositoriesYumGroupByRepositorynameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutRepositoriesYumGroupByRepositorynameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutRepositoriesYumGroupByRepositorynameNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutRepositoriesYumGroupByRepositorynameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutRepositoriesYumGroupByRepositorynameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutRepositoriesYumGroupByRepositorynameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /v1/repositories/yum/group/{repositoryName}] PutRepositoriesYumGroupByRepositoryname", response, response.Code())
	}
}

// NewPutRepositoriesYumGroupByRepositorynameNoContent creates a PutRepositoriesYumGroupByRepositorynameNoContent with default headers values
func NewPutRepositoriesYumGroupByRepositorynameNoContent() *PutRepositoriesYumGroupByRepositorynameNoContent {
	return &PutRepositoriesYumGroupByRepositorynameNoContent{}
}

/*
PutRepositoriesYumGroupByRepositorynameNoContent describes a response with status code 204, with default header values.

Repository updated
*/
type PutRepositoriesYumGroupByRepositorynameNoContent struct {
}

// IsSuccess returns true when this put repositories yum group by repositoryname no content response has a 2xx status code
func (o *PutRepositoriesYumGroupByRepositorynameNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put repositories yum group by repositoryname no content response has a 3xx status code
func (o *PutRepositoriesYumGroupByRepositorynameNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories yum group by repositoryname no content response has a 4xx status code
func (o *PutRepositoriesYumGroupByRepositorynameNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this put repositories yum group by repositoryname no content response has a 5xx status code
func (o *PutRepositoriesYumGroupByRepositorynameNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories yum group by repositoryname no content response a status code equal to that given
func (o *PutRepositoriesYumGroupByRepositorynameNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the put repositories yum group by repositoryname no content response
func (o *PutRepositoriesYumGroupByRepositorynameNoContent) Code() int {
	return 204
}

func (o *PutRepositoriesYumGroupByRepositorynameNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/yum/group/{repositoryName}][%d] putRepositoriesYumGroupByRepositorynameNoContent", 204)
}

func (o *PutRepositoriesYumGroupByRepositorynameNoContent) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/yum/group/{repositoryName}][%d] putRepositoriesYumGroupByRepositorynameNoContent", 204)
}

func (o *PutRepositoriesYumGroupByRepositorynameNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutRepositoriesYumGroupByRepositorynameBadRequest creates a PutRepositoriesYumGroupByRepositorynameBadRequest with default headers values
func NewPutRepositoriesYumGroupByRepositorynameBadRequest() *PutRepositoriesYumGroupByRepositorynameBadRequest {
	return &PutRepositoriesYumGroupByRepositorynameBadRequest{}
}

/*
PutRepositoriesYumGroupByRepositorynameBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PutRepositoriesYumGroupByRepositorynameBadRequest struct {
}

// IsSuccess returns true when this put repositories yum group by repositoryname bad request response has a 2xx status code
func (o *PutRepositoriesYumGroupByRepositorynameBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put repositories yum group by repositoryname bad request response has a 3xx status code
func (o *PutRepositoriesYumGroupByRepositorynameBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories yum group by repositoryname bad request response has a 4xx status code
func (o *PutRepositoriesYumGroupByRepositorynameBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put repositories yum group by repositoryname bad request response has a 5xx status code
func (o *PutRepositoriesYumGroupByRepositorynameBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories yum group by repositoryname bad request response a status code equal to that given
func (o *PutRepositoriesYumGroupByRepositorynameBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the put repositories yum group by repositoryname bad request response
func (o *PutRepositoriesYumGroupByRepositorynameBadRequest) Code() int {
	return 400
}

func (o *PutRepositoriesYumGroupByRepositorynameBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/yum/group/{repositoryName}][%d] putRepositoriesYumGroupByRepositorynameBadRequest", 400)
}

func (o *PutRepositoriesYumGroupByRepositorynameBadRequest) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/yum/group/{repositoryName}][%d] putRepositoriesYumGroupByRepositorynameBadRequest", 400)
}

func (o *PutRepositoriesYumGroupByRepositorynameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutRepositoriesYumGroupByRepositorynameUnauthorized creates a PutRepositoriesYumGroupByRepositorynameUnauthorized with default headers values
func NewPutRepositoriesYumGroupByRepositorynameUnauthorized() *PutRepositoriesYumGroupByRepositorynameUnauthorized {
	return &PutRepositoriesYumGroupByRepositorynameUnauthorized{}
}

/*
PutRepositoriesYumGroupByRepositorynameUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type PutRepositoriesYumGroupByRepositorynameUnauthorized struct {
}

// IsSuccess returns true when this put repositories yum group by repositoryname unauthorized response has a 2xx status code
func (o *PutRepositoriesYumGroupByRepositorynameUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put repositories yum group by repositoryname unauthorized response has a 3xx status code
func (o *PutRepositoriesYumGroupByRepositorynameUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories yum group by repositoryname unauthorized response has a 4xx status code
func (o *PutRepositoriesYumGroupByRepositorynameUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put repositories yum group by repositoryname unauthorized response has a 5xx status code
func (o *PutRepositoriesYumGroupByRepositorynameUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories yum group by repositoryname unauthorized response a status code equal to that given
func (o *PutRepositoriesYumGroupByRepositorynameUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the put repositories yum group by repositoryname unauthorized response
func (o *PutRepositoriesYumGroupByRepositorynameUnauthorized) Code() int {
	return 401
}

func (o *PutRepositoriesYumGroupByRepositorynameUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/yum/group/{repositoryName}][%d] putRepositoriesYumGroupByRepositorynameUnauthorized", 401)
}

func (o *PutRepositoriesYumGroupByRepositorynameUnauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/yum/group/{repositoryName}][%d] putRepositoriesYumGroupByRepositorynameUnauthorized", 401)
}

func (o *PutRepositoriesYumGroupByRepositorynameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutRepositoriesYumGroupByRepositorynameForbidden creates a PutRepositoriesYumGroupByRepositorynameForbidden with default headers values
func NewPutRepositoriesYumGroupByRepositorynameForbidden() *PutRepositoriesYumGroupByRepositorynameForbidden {
	return &PutRepositoriesYumGroupByRepositorynameForbidden{}
}

/*
PutRepositoriesYumGroupByRepositorynameForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type PutRepositoriesYumGroupByRepositorynameForbidden struct {
}

// IsSuccess returns true when this put repositories yum group by repositoryname forbidden response has a 2xx status code
func (o *PutRepositoriesYumGroupByRepositorynameForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put repositories yum group by repositoryname forbidden response has a 3xx status code
func (o *PutRepositoriesYumGroupByRepositorynameForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put repositories yum group by repositoryname forbidden response has a 4xx status code
func (o *PutRepositoriesYumGroupByRepositorynameForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put repositories yum group by repositoryname forbidden response has a 5xx status code
func (o *PutRepositoriesYumGroupByRepositorynameForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put repositories yum group by repositoryname forbidden response a status code equal to that given
func (o *PutRepositoriesYumGroupByRepositorynameForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the put repositories yum group by repositoryname forbidden response
func (o *PutRepositoriesYumGroupByRepositorynameForbidden) Code() int {
	return 403
}

func (o *PutRepositoriesYumGroupByRepositorynameForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/yum/group/{repositoryName}][%d] putRepositoriesYumGroupByRepositorynameForbidden", 403)
}

func (o *PutRepositoriesYumGroupByRepositorynameForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/yum/group/{repositoryName}][%d] putRepositoriesYumGroupByRepositorynameForbidden", 403)
}

func (o *PutRepositoriesYumGroupByRepositorynameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
