// Code generated by go-swagger; DO NOT EDIT.

package security_management_l_d_a_p

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetLdapServerReader is a Reader for the GetLdapServer structure.
type GetLdapServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLdapServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLdapServerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetLdapServerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLdapServerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLdapServerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/security/ldap/{name}] getLdapServer", response, response.Code())
	}
}

// NewGetLdapServerOK creates a GetLdapServerOK with default headers values
func NewGetLdapServerOK() *GetLdapServerOK {
	return &GetLdapServerOK{}
}

/*
GetLdapServerOK describes a response with status code 200, with default header values.

LDAP server returned
*/
type GetLdapServerOK struct {
}

// IsSuccess returns true when this get ldap server o k response has a 2xx status code
func (o *GetLdapServerOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get ldap server o k response has a 3xx status code
func (o *GetLdapServerOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get ldap server o k response has a 4xx status code
func (o *GetLdapServerOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get ldap server o k response has a 5xx status code
func (o *GetLdapServerOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get ldap server o k response a status code equal to that given
func (o *GetLdapServerOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get ldap server o k response
func (o *GetLdapServerOK) Code() int {
	return 200
}

func (o *GetLdapServerOK) Error() string {
	return fmt.Sprintf("[GET /v1/security/ldap/{name}][%d] getLdapServerOK", 200)
}

func (o *GetLdapServerOK) String() string {
	return fmt.Sprintf("[GET /v1/security/ldap/{name}][%d] getLdapServerOK", 200)
}

func (o *GetLdapServerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLdapServerUnauthorized creates a GetLdapServerUnauthorized with default headers values
func NewGetLdapServerUnauthorized() *GetLdapServerUnauthorized {
	return &GetLdapServerUnauthorized{}
}

/*
GetLdapServerUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type GetLdapServerUnauthorized struct {
}

// IsSuccess returns true when this get ldap server unauthorized response has a 2xx status code
func (o *GetLdapServerUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get ldap server unauthorized response has a 3xx status code
func (o *GetLdapServerUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get ldap server unauthorized response has a 4xx status code
func (o *GetLdapServerUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get ldap server unauthorized response has a 5xx status code
func (o *GetLdapServerUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get ldap server unauthorized response a status code equal to that given
func (o *GetLdapServerUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get ldap server unauthorized response
func (o *GetLdapServerUnauthorized) Code() int {
	return 401
}

func (o *GetLdapServerUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/security/ldap/{name}][%d] getLdapServerUnauthorized", 401)
}

func (o *GetLdapServerUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/security/ldap/{name}][%d] getLdapServerUnauthorized", 401)
}

func (o *GetLdapServerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLdapServerForbidden creates a GetLdapServerForbidden with default headers values
func NewGetLdapServerForbidden() *GetLdapServerForbidden {
	return &GetLdapServerForbidden{}
}

/*
GetLdapServerForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type GetLdapServerForbidden struct {
}

// IsSuccess returns true when this get ldap server forbidden response has a 2xx status code
func (o *GetLdapServerForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get ldap server forbidden response has a 3xx status code
func (o *GetLdapServerForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get ldap server forbidden response has a 4xx status code
func (o *GetLdapServerForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get ldap server forbidden response has a 5xx status code
func (o *GetLdapServerForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get ldap server forbidden response a status code equal to that given
func (o *GetLdapServerForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get ldap server forbidden response
func (o *GetLdapServerForbidden) Code() int {
	return 403
}

func (o *GetLdapServerForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/security/ldap/{name}][%d] getLdapServerForbidden", 403)
}

func (o *GetLdapServerForbidden) String() string {
	return fmt.Sprintf("[GET /v1/security/ldap/{name}][%d] getLdapServerForbidden", 403)
}

func (o *GetLdapServerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLdapServerNotFound creates a GetLdapServerNotFound with default headers values
func NewGetLdapServerNotFound() *GetLdapServerNotFound {
	return &GetLdapServerNotFound{}
}

/*
GetLdapServerNotFound describes a response with status code 404, with default header values.

LDAP server not found
*/
type GetLdapServerNotFound struct {
}

// IsSuccess returns true when this get ldap server not found response has a 2xx status code
func (o *GetLdapServerNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get ldap server not found response has a 3xx status code
func (o *GetLdapServerNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get ldap server not found response has a 4xx status code
func (o *GetLdapServerNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get ldap server not found response has a 5xx status code
func (o *GetLdapServerNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get ldap server not found response a status code equal to that given
func (o *GetLdapServerNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get ldap server not found response
func (o *GetLdapServerNotFound) Code() int {
	return 404
}

func (o *GetLdapServerNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/security/ldap/{name}][%d] getLdapServerNotFound", 404)
}

func (o *GetLdapServerNotFound) String() string {
	return fmt.Sprintf("[GET /v1/security/ldap/{name}][%d] getLdapServerNotFound", 404)
}

func (o *GetLdapServerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
