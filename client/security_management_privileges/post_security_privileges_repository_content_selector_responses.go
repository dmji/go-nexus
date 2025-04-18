// Code generated by go-swagger; DO NOT EDIT.

package security_management_privileges

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostSecurityPrivilegesRepositoryContentSelectorReader is a Reader for the PostSecurityPrivilegesRepositoryContentSelector structure.
type PostSecurityPrivilegesRepositoryContentSelectorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSecurityPrivilegesRepositoryContentSelectorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 400:
		result := NewPostSecurityPrivilegesRepositoryContentSelectorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostSecurityPrivilegesRepositoryContentSelectorForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/security/privileges/repository-content-selector] PostSecurityPrivilegesRepositoryContentSelector", response, response.Code())
	}
}

// NewPostSecurityPrivilegesRepositoryContentSelectorBadRequest creates a PostSecurityPrivilegesRepositoryContentSelectorBadRequest with default headers values
func NewPostSecurityPrivilegesRepositoryContentSelectorBadRequest() *PostSecurityPrivilegesRepositoryContentSelectorBadRequest {
	return &PostSecurityPrivilegesRepositoryContentSelectorBadRequest{}
}

/*
PostSecurityPrivilegesRepositoryContentSelectorBadRequest describes a response with status code 400, with default header values.

Privilege object not configured properly.
*/
type PostSecurityPrivilegesRepositoryContentSelectorBadRequest struct {
}

// IsSuccess returns true when this post security privileges repository content selector bad request response has a 2xx status code
func (o *PostSecurityPrivilegesRepositoryContentSelectorBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post security privileges repository content selector bad request response has a 3xx status code
func (o *PostSecurityPrivilegesRepositoryContentSelectorBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post security privileges repository content selector bad request response has a 4xx status code
func (o *PostSecurityPrivilegesRepositoryContentSelectorBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post security privileges repository content selector bad request response has a 5xx status code
func (o *PostSecurityPrivilegesRepositoryContentSelectorBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post security privileges repository content selector bad request response a status code equal to that given
func (o *PostSecurityPrivilegesRepositoryContentSelectorBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post security privileges repository content selector bad request response
func (o *PostSecurityPrivilegesRepositoryContentSelectorBadRequest) Code() int {
	return 400
}

func (o *PostSecurityPrivilegesRepositoryContentSelectorBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/security/privileges/repository-content-selector][%d] postSecurityPrivilegesRepositoryContentSelectorBadRequest", 400)
}

func (o *PostSecurityPrivilegesRepositoryContentSelectorBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/security/privileges/repository-content-selector][%d] postSecurityPrivilegesRepositoryContentSelectorBadRequest", 400)
}

func (o *PostSecurityPrivilegesRepositoryContentSelectorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSecurityPrivilegesRepositoryContentSelectorForbidden creates a PostSecurityPrivilegesRepositoryContentSelectorForbidden with default headers values
func NewPostSecurityPrivilegesRepositoryContentSelectorForbidden() *PostSecurityPrivilegesRepositoryContentSelectorForbidden {
	return &PostSecurityPrivilegesRepositoryContentSelectorForbidden{}
}

/*
PostSecurityPrivilegesRepositoryContentSelectorForbidden describes a response with status code 403, with default header values.

The user does not have permission to perform the operation.
*/
type PostSecurityPrivilegesRepositoryContentSelectorForbidden struct {
}

// IsSuccess returns true when this post security privileges repository content selector forbidden response has a 2xx status code
func (o *PostSecurityPrivilegesRepositoryContentSelectorForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post security privileges repository content selector forbidden response has a 3xx status code
func (o *PostSecurityPrivilegesRepositoryContentSelectorForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post security privileges repository content selector forbidden response has a 4xx status code
func (o *PostSecurityPrivilegesRepositoryContentSelectorForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post security privileges repository content selector forbidden response has a 5xx status code
func (o *PostSecurityPrivilegesRepositoryContentSelectorForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post security privileges repository content selector forbidden response a status code equal to that given
func (o *PostSecurityPrivilegesRepositoryContentSelectorForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the post security privileges repository content selector forbidden response
func (o *PostSecurityPrivilegesRepositoryContentSelectorForbidden) Code() int {
	return 403
}

func (o *PostSecurityPrivilegesRepositoryContentSelectorForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/security/privileges/repository-content-selector][%d] postSecurityPrivilegesRepositoryContentSelectorForbidden", 403)
}

func (o *PostSecurityPrivilegesRepositoryContentSelectorForbidden) String() string {
	return fmt.Sprintf("[POST /v1/security/privileges/repository-content-selector][%d] postSecurityPrivilegesRepositoryContentSelectorForbidden", 403)
}

func (o *PostSecurityPrivilegesRepositoryContentSelectorForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
