// Code generated by go-swagger; DO NOT EDIT.

package security_management_l_d_a_p

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetSecurityLdapByNameReader is a Reader for the GetSecurityLdapByName structure.
type GetSecurityLdapByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSecurityLdapByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSecurityLdapByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetSecurityLdapByNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSecurityLdapByNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSecurityLdapByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/security/ldap/{name}] GetSecurityLdapByName", response, response.Code())
	}
}

// NewGetSecurityLdapByNameOK creates a GetSecurityLdapByNameOK with default headers values
func NewGetSecurityLdapByNameOK() *GetSecurityLdapByNameOK {
	return &GetSecurityLdapByNameOK{}
}

/*
GetSecurityLdapByNameOK describes a response with status code 200, with default header values.

LDAP server returned
*/
type GetSecurityLdapByNameOK struct {
}

// IsSuccess returns true when this get security ldap by name o k response has a 2xx status code
func (o *GetSecurityLdapByNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get security ldap by name o k response has a 3xx status code
func (o *GetSecurityLdapByNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get security ldap by name o k response has a 4xx status code
func (o *GetSecurityLdapByNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get security ldap by name o k response has a 5xx status code
func (o *GetSecurityLdapByNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get security ldap by name o k response a status code equal to that given
func (o *GetSecurityLdapByNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get security ldap by name o k response
func (o *GetSecurityLdapByNameOK) Code() int {
	return 200
}

func (o *GetSecurityLdapByNameOK) Error() string {
	return fmt.Sprintf("[GET /v1/security/ldap/{name}][%d] getSecurityLdapByNameOK", 200)
}

func (o *GetSecurityLdapByNameOK) String() string {
	return fmt.Sprintf("[GET /v1/security/ldap/{name}][%d] getSecurityLdapByNameOK", 200)
}

func (o *GetSecurityLdapByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSecurityLdapByNameUnauthorized creates a GetSecurityLdapByNameUnauthorized with default headers values
func NewGetSecurityLdapByNameUnauthorized() *GetSecurityLdapByNameUnauthorized {
	return &GetSecurityLdapByNameUnauthorized{}
}

/*
GetSecurityLdapByNameUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type GetSecurityLdapByNameUnauthorized struct {
}

// IsSuccess returns true when this get security ldap by name unauthorized response has a 2xx status code
func (o *GetSecurityLdapByNameUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get security ldap by name unauthorized response has a 3xx status code
func (o *GetSecurityLdapByNameUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get security ldap by name unauthorized response has a 4xx status code
func (o *GetSecurityLdapByNameUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get security ldap by name unauthorized response has a 5xx status code
func (o *GetSecurityLdapByNameUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get security ldap by name unauthorized response a status code equal to that given
func (o *GetSecurityLdapByNameUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get security ldap by name unauthorized response
func (o *GetSecurityLdapByNameUnauthorized) Code() int {
	return 401
}

func (o *GetSecurityLdapByNameUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/security/ldap/{name}][%d] getSecurityLdapByNameUnauthorized", 401)
}

func (o *GetSecurityLdapByNameUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/security/ldap/{name}][%d] getSecurityLdapByNameUnauthorized", 401)
}

func (o *GetSecurityLdapByNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSecurityLdapByNameForbidden creates a GetSecurityLdapByNameForbidden with default headers values
func NewGetSecurityLdapByNameForbidden() *GetSecurityLdapByNameForbidden {
	return &GetSecurityLdapByNameForbidden{}
}

/*
GetSecurityLdapByNameForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type GetSecurityLdapByNameForbidden struct {
}

// IsSuccess returns true when this get security ldap by name forbidden response has a 2xx status code
func (o *GetSecurityLdapByNameForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get security ldap by name forbidden response has a 3xx status code
func (o *GetSecurityLdapByNameForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get security ldap by name forbidden response has a 4xx status code
func (o *GetSecurityLdapByNameForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get security ldap by name forbidden response has a 5xx status code
func (o *GetSecurityLdapByNameForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get security ldap by name forbidden response a status code equal to that given
func (o *GetSecurityLdapByNameForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get security ldap by name forbidden response
func (o *GetSecurityLdapByNameForbidden) Code() int {
	return 403
}

func (o *GetSecurityLdapByNameForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/security/ldap/{name}][%d] getSecurityLdapByNameForbidden", 403)
}

func (o *GetSecurityLdapByNameForbidden) String() string {
	return fmt.Sprintf("[GET /v1/security/ldap/{name}][%d] getSecurityLdapByNameForbidden", 403)
}

func (o *GetSecurityLdapByNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSecurityLdapByNameNotFound creates a GetSecurityLdapByNameNotFound with default headers values
func NewGetSecurityLdapByNameNotFound() *GetSecurityLdapByNameNotFound {
	return &GetSecurityLdapByNameNotFound{}
}

/*
GetSecurityLdapByNameNotFound describes a response with status code 404, with default header values.

LDAP server not found
*/
type GetSecurityLdapByNameNotFound struct {
}

// IsSuccess returns true when this get security ldap by name not found response has a 2xx status code
func (o *GetSecurityLdapByNameNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get security ldap by name not found response has a 3xx status code
func (o *GetSecurityLdapByNameNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get security ldap by name not found response has a 4xx status code
func (o *GetSecurityLdapByNameNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get security ldap by name not found response has a 5xx status code
func (o *GetSecurityLdapByNameNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get security ldap by name not found response a status code equal to that given
func (o *GetSecurityLdapByNameNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get security ldap by name not found response
func (o *GetSecurityLdapByNameNotFound) Code() int {
	return 404
}

func (o *GetSecurityLdapByNameNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/security/ldap/{name}][%d] getSecurityLdapByNameNotFound", 404)
}

func (o *GetSecurityLdapByNameNotFound) String() string {
	return fmt.Sprintf("[GET /v1/security/ldap/{name}][%d] getSecurityLdapByNameNotFound", 404)
}

func (o *GetSecurityLdapByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
