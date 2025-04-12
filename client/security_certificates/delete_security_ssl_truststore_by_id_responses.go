// Code generated by go-swagger; DO NOT EDIT.

package security_certificates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteSecuritySslTruststoreByIDReader is a Reader for the DeleteSecuritySslTruststoreByID structure.
type DeleteSecuritySslTruststoreByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSecuritySslTruststoreByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 403:
		result := NewDeleteSecuritySslTruststoreByIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /v1/security/ssl/truststore/{id}] DeleteSecuritySslTruststoreById", response, response.Code())
	}
}

// NewDeleteSecuritySslTruststoreByIDForbidden creates a DeleteSecuritySslTruststoreByIDForbidden with default headers values
func NewDeleteSecuritySslTruststoreByIDForbidden() *DeleteSecuritySslTruststoreByIDForbidden {
	return &DeleteSecuritySslTruststoreByIDForbidden{}
}

/*
DeleteSecuritySslTruststoreByIDForbidden describes a response with status code 403, with default header values.

Insufficient permissions to remove certificate from the trust store
*/
type DeleteSecuritySslTruststoreByIDForbidden struct {
}

// IsSuccess returns true when this delete security ssl truststore by Id forbidden response has a 2xx status code
func (o *DeleteSecuritySslTruststoreByIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete security ssl truststore by Id forbidden response has a 3xx status code
func (o *DeleteSecuritySslTruststoreByIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete security ssl truststore by Id forbidden response has a 4xx status code
func (o *DeleteSecuritySslTruststoreByIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete security ssl truststore by Id forbidden response has a 5xx status code
func (o *DeleteSecuritySslTruststoreByIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete security ssl truststore by Id forbidden response a status code equal to that given
func (o *DeleteSecuritySslTruststoreByIDForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete security ssl truststore by Id forbidden response
func (o *DeleteSecuritySslTruststoreByIDForbidden) Code() int {
	return 403
}

func (o *DeleteSecuritySslTruststoreByIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /v1/security/ssl/truststore/{id}][%d] deleteSecuritySslTruststoreByIdForbidden", 403)
}

func (o *DeleteSecuritySslTruststoreByIDForbidden) String() string {
	return fmt.Sprintf("[DELETE /v1/security/ssl/truststore/{id}][%d] deleteSecuritySslTruststoreByIdForbidden", 403)
}

func (o *DeleteSecuritySslTruststoreByIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
