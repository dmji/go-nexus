// Code generated by go-swagger; DO NOT EDIT.

package security_management_roles

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

// GetSecurityRolesReader is a Reader for the GetSecurityRoles structure.
type GetSecurityRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSecurityRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSecurityRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSecurityRolesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSecurityRolesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/security/roles] GetSecurityRoles", response, response.Code())
	}
}

// NewGetSecurityRolesOK creates a GetSecurityRolesOK with default headers values
func NewGetSecurityRolesOK() *GetSecurityRolesOK {
	return &GetSecurityRolesOK{}
}

/*
GetSecurityRolesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetSecurityRolesOK struct {
	Payload []*models.RoleXOResponse
}

// IsSuccess returns true when this get security roles o k response has a 2xx status code
func (o *GetSecurityRolesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get security roles o k response has a 3xx status code
func (o *GetSecurityRolesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get security roles o k response has a 4xx status code
func (o *GetSecurityRolesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get security roles o k response has a 5xx status code
func (o *GetSecurityRolesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get security roles o k response a status code equal to that given
func (o *GetSecurityRolesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get security roles o k response
func (o *GetSecurityRolesOK) Code() int {
	return 200
}

func (o *GetSecurityRolesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/security/roles][%d] getSecurityRolesOK %s", 200, payload)
}

func (o *GetSecurityRolesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/security/roles][%d] getSecurityRolesOK %s", 200, payload)
}

func (o *GetSecurityRolesOK) GetPayload() []*models.RoleXOResponse {
	return o.Payload
}

func (o *GetSecurityRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSecurityRolesBadRequest creates a GetSecurityRolesBadRequest with default headers values
func NewGetSecurityRolesBadRequest() *GetSecurityRolesBadRequest {
	return &GetSecurityRolesBadRequest{}
}

/*
GetSecurityRolesBadRequest describes a response with status code 400, with default header values.

The specified source does not exist
*/
type GetSecurityRolesBadRequest struct {
}

// IsSuccess returns true when this get security roles bad request response has a 2xx status code
func (o *GetSecurityRolesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get security roles bad request response has a 3xx status code
func (o *GetSecurityRolesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get security roles bad request response has a 4xx status code
func (o *GetSecurityRolesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get security roles bad request response has a 5xx status code
func (o *GetSecurityRolesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get security roles bad request response a status code equal to that given
func (o *GetSecurityRolesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get security roles bad request response
func (o *GetSecurityRolesBadRequest) Code() int {
	return 400
}

func (o *GetSecurityRolesBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/security/roles][%d] getSecurityRolesBadRequest", 400)
}

func (o *GetSecurityRolesBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/security/roles][%d] getSecurityRolesBadRequest", 400)
}

func (o *GetSecurityRolesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSecurityRolesForbidden creates a GetSecurityRolesForbidden with default headers values
func NewGetSecurityRolesForbidden() *GetSecurityRolesForbidden {
	return &GetSecurityRolesForbidden{}
}

/*
GetSecurityRolesForbidden describes a response with status code 403, with default header values.

Insufficient permissions to read roles
*/
type GetSecurityRolesForbidden struct {
}

// IsSuccess returns true when this get security roles forbidden response has a 2xx status code
func (o *GetSecurityRolesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get security roles forbidden response has a 3xx status code
func (o *GetSecurityRolesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get security roles forbidden response has a 4xx status code
func (o *GetSecurityRolesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get security roles forbidden response has a 5xx status code
func (o *GetSecurityRolesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get security roles forbidden response a status code equal to that given
func (o *GetSecurityRolesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get security roles forbidden response
func (o *GetSecurityRolesForbidden) Code() int {
	return 403
}

func (o *GetSecurityRolesForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/security/roles][%d] getSecurityRolesForbidden", 403)
}

func (o *GetSecurityRolesForbidden) String() string {
	return fmt.Sprintf("[GET /v1/security/roles][%d] getSecurityRolesForbidden", 403)
}

func (o *GetSecurityRolesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
