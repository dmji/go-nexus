// Code generated by go-swagger; DO NOT EDIT.

package security_management_privileges

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostSecurityPrivilegesScriptReader is a Reader for the PostSecurityPrivilegesScript structure.
type PostSecurityPrivilegesScriptReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSecurityPrivilegesScriptReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 400:
		result := NewPostSecurityPrivilegesScriptBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostSecurityPrivilegesScriptForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/security/privileges/script] PostSecurityPrivilegesScript", response, response.Code())
	}
}

// NewPostSecurityPrivilegesScriptBadRequest creates a PostSecurityPrivilegesScriptBadRequest with default headers values
func NewPostSecurityPrivilegesScriptBadRequest() *PostSecurityPrivilegesScriptBadRequest {
	return &PostSecurityPrivilegesScriptBadRequest{}
}

/*
PostSecurityPrivilegesScriptBadRequest describes a response with status code 400, with default header values.

Privilege object not configured properly.
*/
type PostSecurityPrivilegesScriptBadRequest struct {
}

// IsSuccess returns true when this post security privileges script bad request response has a 2xx status code
func (o *PostSecurityPrivilegesScriptBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post security privileges script bad request response has a 3xx status code
func (o *PostSecurityPrivilegesScriptBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post security privileges script bad request response has a 4xx status code
func (o *PostSecurityPrivilegesScriptBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post security privileges script bad request response has a 5xx status code
func (o *PostSecurityPrivilegesScriptBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post security privileges script bad request response a status code equal to that given
func (o *PostSecurityPrivilegesScriptBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post security privileges script bad request response
func (o *PostSecurityPrivilegesScriptBadRequest) Code() int {
	return 400
}

func (o *PostSecurityPrivilegesScriptBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/security/privileges/script][%d] postSecurityPrivilegesScriptBadRequest", 400)
}

func (o *PostSecurityPrivilegesScriptBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/security/privileges/script][%d] postSecurityPrivilegesScriptBadRequest", 400)
}

func (o *PostSecurityPrivilegesScriptBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSecurityPrivilegesScriptForbidden creates a PostSecurityPrivilegesScriptForbidden with default headers values
func NewPostSecurityPrivilegesScriptForbidden() *PostSecurityPrivilegesScriptForbidden {
	return &PostSecurityPrivilegesScriptForbidden{}
}

/*
PostSecurityPrivilegesScriptForbidden describes a response with status code 403, with default header values.

The user does not have permission to perform the operation.
*/
type PostSecurityPrivilegesScriptForbidden struct {
}

// IsSuccess returns true when this post security privileges script forbidden response has a 2xx status code
func (o *PostSecurityPrivilegesScriptForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post security privileges script forbidden response has a 3xx status code
func (o *PostSecurityPrivilegesScriptForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post security privileges script forbidden response has a 4xx status code
func (o *PostSecurityPrivilegesScriptForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post security privileges script forbidden response has a 5xx status code
func (o *PostSecurityPrivilegesScriptForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post security privileges script forbidden response a status code equal to that given
func (o *PostSecurityPrivilegesScriptForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the post security privileges script forbidden response
func (o *PostSecurityPrivilegesScriptForbidden) Code() int {
	return 403
}

func (o *PostSecurityPrivilegesScriptForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/security/privileges/script][%d] postSecurityPrivilegesScriptForbidden", 403)
}

func (o *PostSecurityPrivilegesScriptForbidden) String() string {
	return fmt.Sprintf("[POST /v1/security/privileges/script][%d] postSecurityPrivilegesScriptForbidden", 403)
}

func (o *PostSecurityPrivilegesScriptForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
