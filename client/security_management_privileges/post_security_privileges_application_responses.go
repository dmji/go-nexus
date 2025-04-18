// Code generated by go-swagger; DO NOT EDIT.

package security_management_privileges

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostSecurityPrivilegesApplicationReader is a Reader for the PostSecurityPrivilegesApplication structure.
type PostSecurityPrivilegesApplicationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSecurityPrivilegesApplicationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 400:
		result := NewPostSecurityPrivilegesApplicationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostSecurityPrivilegesApplicationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/security/privileges/application] PostSecurityPrivilegesApplication", response, response.Code())
	}
}

// NewPostSecurityPrivilegesApplicationBadRequest creates a PostSecurityPrivilegesApplicationBadRequest with default headers values
func NewPostSecurityPrivilegesApplicationBadRequest() *PostSecurityPrivilegesApplicationBadRequest {
	return &PostSecurityPrivilegesApplicationBadRequest{}
}

/*
PostSecurityPrivilegesApplicationBadRequest describes a response with status code 400, with default header values.

Privilege object not configured properly.
*/
type PostSecurityPrivilegesApplicationBadRequest struct {
}

// IsSuccess returns true when this post security privileges application bad request response has a 2xx status code
func (o *PostSecurityPrivilegesApplicationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post security privileges application bad request response has a 3xx status code
func (o *PostSecurityPrivilegesApplicationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post security privileges application bad request response has a 4xx status code
func (o *PostSecurityPrivilegesApplicationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post security privileges application bad request response has a 5xx status code
func (o *PostSecurityPrivilegesApplicationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post security privileges application bad request response a status code equal to that given
func (o *PostSecurityPrivilegesApplicationBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post security privileges application bad request response
func (o *PostSecurityPrivilegesApplicationBadRequest) Code() int {
	return 400
}

func (o *PostSecurityPrivilegesApplicationBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/security/privileges/application][%d] postSecurityPrivilegesApplicationBadRequest", 400)
}

func (o *PostSecurityPrivilegesApplicationBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/security/privileges/application][%d] postSecurityPrivilegesApplicationBadRequest", 400)
}

func (o *PostSecurityPrivilegesApplicationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSecurityPrivilegesApplicationForbidden creates a PostSecurityPrivilegesApplicationForbidden with default headers values
func NewPostSecurityPrivilegesApplicationForbidden() *PostSecurityPrivilegesApplicationForbidden {
	return &PostSecurityPrivilegesApplicationForbidden{}
}

/*
PostSecurityPrivilegesApplicationForbidden describes a response with status code 403, with default header values.

The user does not have permission to perform the operation.
*/
type PostSecurityPrivilegesApplicationForbidden struct {
}

// IsSuccess returns true when this post security privileges application forbidden response has a 2xx status code
func (o *PostSecurityPrivilegesApplicationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post security privileges application forbidden response has a 3xx status code
func (o *PostSecurityPrivilegesApplicationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post security privileges application forbidden response has a 4xx status code
func (o *PostSecurityPrivilegesApplicationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post security privileges application forbidden response has a 5xx status code
func (o *PostSecurityPrivilegesApplicationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post security privileges application forbidden response a status code equal to that given
func (o *PostSecurityPrivilegesApplicationForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the post security privileges application forbidden response
func (o *PostSecurityPrivilegesApplicationForbidden) Code() int {
	return 403
}

func (o *PostSecurityPrivilegesApplicationForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/security/privileges/application][%d] postSecurityPrivilegesApplicationForbidden", 403)
}

func (o *PostSecurityPrivilegesApplicationForbidden) String() string {
	return fmt.Sprintf("[POST /v1/security/privileges/application][%d] postSecurityPrivilegesApplicationForbidden", 403)
}

func (o *PostSecurityPrivilegesApplicationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
