// Code generated by go-swagger; DO NOT EDIT.

package security_management_l_d_a_p

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PutSecurityLdapByNameReader is a Reader for the PutSecurityLdapByName structure.
type PutSecurityLdapByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutSecurityLdapByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutSecurityLdapByNameNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPutSecurityLdapByNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutSecurityLdapByNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutSecurityLdapByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /v1/security/ldap/{name}] PutSecurityLdapByName", response, response.Code())
	}
}

// NewPutSecurityLdapByNameNoContent creates a PutSecurityLdapByNameNoContent with default headers values
func NewPutSecurityLdapByNameNoContent() *PutSecurityLdapByNameNoContent {
	return &PutSecurityLdapByNameNoContent{}
}

/*
PutSecurityLdapByNameNoContent describes a response with status code 204, with default header values.

LDAP server updated
*/
type PutSecurityLdapByNameNoContent struct {
}

// IsSuccess returns true when this put security ldap by name no content response has a 2xx status code
func (o *PutSecurityLdapByNameNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put security ldap by name no content response has a 3xx status code
func (o *PutSecurityLdapByNameNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put security ldap by name no content response has a 4xx status code
func (o *PutSecurityLdapByNameNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this put security ldap by name no content response has a 5xx status code
func (o *PutSecurityLdapByNameNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this put security ldap by name no content response a status code equal to that given
func (o *PutSecurityLdapByNameNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the put security ldap by name no content response
func (o *PutSecurityLdapByNameNoContent) Code() int {
	return 204
}

func (o *PutSecurityLdapByNameNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/security/ldap/{name}][%d] putSecurityLdapByNameNoContent", 204)
}

func (o *PutSecurityLdapByNameNoContent) String() string {
	return fmt.Sprintf("[PUT /v1/security/ldap/{name}][%d] putSecurityLdapByNameNoContent", 204)
}

func (o *PutSecurityLdapByNameNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutSecurityLdapByNameUnauthorized creates a PutSecurityLdapByNameUnauthorized with default headers values
func NewPutSecurityLdapByNameUnauthorized() *PutSecurityLdapByNameUnauthorized {
	return &PutSecurityLdapByNameUnauthorized{}
}

/*
PutSecurityLdapByNameUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type PutSecurityLdapByNameUnauthorized struct {
}

// IsSuccess returns true when this put security ldap by name unauthorized response has a 2xx status code
func (o *PutSecurityLdapByNameUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put security ldap by name unauthorized response has a 3xx status code
func (o *PutSecurityLdapByNameUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put security ldap by name unauthorized response has a 4xx status code
func (o *PutSecurityLdapByNameUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put security ldap by name unauthorized response has a 5xx status code
func (o *PutSecurityLdapByNameUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put security ldap by name unauthorized response a status code equal to that given
func (o *PutSecurityLdapByNameUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the put security ldap by name unauthorized response
func (o *PutSecurityLdapByNameUnauthorized) Code() int {
	return 401
}

func (o *PutSecurityLdapByNameUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/security/ldap/{name}][%d] putSecurityLdapByNameUnauthorized", 401)
}

func (o *PutSecurityLdapByNameUnauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/security/ldap/{name}][%d] putSecurityLdapByNameUnauthorized", 401)
}

func (o *PutSecurityLdapByNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutSecurityLdapByNameForbidden creates a PutSecurityLdapByNameForbidden with default headers values
func NewPutSecurityLdapByNameForbidden() *PutSecurityLdapByNameForbidden {
	return &PutSecurityLdapByNameForbidden{}
}

/*
PutSecurityLdapByNameForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type PutSecurityLdapByNameForbidden struct {
}

// IsSuccess returns true when this put security ldap by name forbidden response has a 2xx status code
func (o *PutSecurityLdapByNameForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put security ldap by name forbidden response has a 3xx status code
func (o *PutSecurityLdapByNameForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put security ldap by name forbidden response has a 4xx status code
func (o *PutSecurityLdapByNameForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put security ldap by name forbidden response has a 5xx status code
func (o *PutSecurityLdapByNameForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put security ldap by name forbidden response a status code equal to that given
func (o *PutSecurityLdapByNameForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the put security ldap by name forbidden response
func (o *PutSecurityLdapByNameForbidden) Code() int {
	return 403
}

func (o *PutSecurityLdapByNameForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/security/ldap/{name}][%d] putSecurityLdapByNameForbidden", 403)
}

func (o *PutSecurityLdapByNameForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/security/ldap/{name}][%d] putSecurityLdapByNameForbidden", 403)
}

func (o *PutSecurityLdapByNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutSecurityLdapByNameNotFound creates a PutSecurityLdapByNameNotFound with default headers values
func NewPutSecurityLdapByNameNotFound() *PutSecurityLdapByNameNotFound {
	return &PutSecurityLdapByNameNotFound{}
}

/*
PutSecurityLdapByNameNotFound describes a response with status code 404, with default header values.

LDAP server not found
*/
type PutSecurityLdapByNameNotFound struct {
}

// IsSuccess returns true when this put security ldap by name not found response has a 2xx status code
func (o *PutSecurityLdapByNameNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put security ldap by name not found response has a 3xx status code
func (o *PutSecurityLdapByNameNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put security ldap by name not found response has a 4xx status code
func (o *PutSecurityLdapByNameNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put security ldap by name not found response has a 5xx status code
func (o *PutSecurityLdapByNameNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put security ldap by name not found response a status code equal to that given
func (o *PutSecurityLdapByNameNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the put security ldap by name not found response
func (o *PutSecurityLdapByNameNotFound) Code() int {
	return 404
}

func (o *PutSecurityLdapByNameNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/security/ldap/{name}][%d] putSecurityLdapByNameNotFound", 404)
}

func (o *PutSecurityLdapByNameNotFound) String() string {
	return fmt.Sprintf("[PUT /v1/security/ldap/{name}][%d] putSecurityLdapByNameNotFound", 404)
}

func (o *PutSecurityLdapByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
