// Code generated by go-swagger; DO NOT EDIT.

package security_management_privileges

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostSecurityPrivilegesRepositoryViewReader is a Reader for the PostSecurityPrivilegesRepositoryView structure.
type PostSecurityPrivilegesRepositoryViewReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSecurityPrivilegesRepositoryViewReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 400:
		result := NewPostSecurityPrivilegesRepositoryViewBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostSecurityPrivilegesRepositoryViewForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/security/privileges/repository-view] PostSecurityPrivilegesRepositoryView", response, response.Code())
	}
}

// NewPostSecurityPrivilegesRepositoryViewBadRequest creates a PostSecurityPrivilegesRepositoryViewBadRequest with default headers values
func NewPostSecurityPrivilegesRepositoryViewBadRequest() *PostSecurityPrivilegesRepositoryViewBadRequest {
	return &PostSecurityPrivilegesRepositoryViewBadRequest{}
}

/*
PostSecurityPrivilegesRepositoryViewBadRequest describes a response with status code 400, with default header values.

Privilege object not configured properly.
*/
type PostSecurityPrivilegesRepositoryViewBadRequest struct {
}

// IsSuccess returns true when this post security privileges repository view bad request response has a 2xx status code
func (o *PostSecurityPrivilegesRepositoryViewBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post security privileges repository view bad request response has a 3xx status code
func (o *PostSecurityPrivilegesRepositoryViewBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post security privileges repository view bad request response has a 4xx status code
func (o *PostSecurityPrivilegesRepositoryViewBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post security privileges repository view bad request response has a 5xx status code
func (o *PostSecurityPrivilegesRepositoryViewBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post security privileges repository view bad request response a status code equal to that given
func (o *PostSecurityPrivilegesRepositoryViewBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post security privileges repository view bad request response
func (o *PostSecurityPrivilegesRepositoryViewBadRequest) Code() int {
	return 400
}

func (o *PostSecurityPrivilegesRepositoryViewBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/security/privileges/repository-view][%d] postSecurityPrivilegesRepositoryViewBadRequest", 400)
}

func (o *PostSecurityPrivilegesRepositoryViewBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/security/privileges/repository-view][%d] postSecurityPrivilegesRepositoryViewBadRequest", 400)
}

func (o *PostSecurityPrivilegesRepositoryViewBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSecurityPrivilegesRepositoryViewForbidden creates a PostSecurityPrivilegesRepositoryViewForbidden with default headers values
func NewPostSecurityPrivilegesRepositoryViewForbidden() *PostSecurityPrivilegesRepositoryViewForbidden {
	return &PostSecurityPrivilegesRepositoryViewForbidden{}
}

/*
PostSecurityPrivilegesRepositoryViewForbidden describes a response with status code 403, with default header values.

The user does not have permission to perform the operation.
*/
type PostSecurityPrivilegesRepositoryViewForbidden struct {
}

// IsSuccess returns true when this post security privileges repository view forbidden response has a 2xx status code
func (o *PostSecurityPrivilegesRepositoryViewForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post security privileges repository view forbidden response has a 3xx status code
func (o *PostSecurityPrivilegesRepositoryViewForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post security privileges repository view forbidden response has a 4xx status code
func (o *PostSecurityPrivilegesRepositoryViewForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post security privileges repository view forbidden response has a 5xx status code
func (o *PostSecurityPrivilegesRepositoryViewForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post security privileges repository view forbidden response a status code equal to that given
func (o *PostSecurityPrivilegesRepositoryViewForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the post security privileges repository view forbidden response
func (o *PostSecurityPrivilegesRepositoryViewForbidden) Code() int {
	return 403
}

func (o *PostSecurityPrivilegesRepositoryViewForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/security/privileges/repository-view][%d] postSecurityPrivilegesRepositoryViewForbidden", 403)
}

func (o *PostSecurityPrivilegesRepositoryViewForbidden) String() string {
	return fmt.Sprintf("[POST /v1/security/privileges/repository-view][%d] postSecurityPrivilegesRepositoryViewForbidden", 403)
}

func (o *PostSecurityPrivilegesRepositoryViewForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
