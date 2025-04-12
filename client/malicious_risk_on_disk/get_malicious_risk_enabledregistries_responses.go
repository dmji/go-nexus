// Code generated by go-swagger; DO NOT EDIT.

package malicious_risk_on_disk

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetMaliciousRiskEnabledregistriesReader is a Reader for the GetMaliciousRiskEnabledregistries structure.
type GetMaliciousRiskEnabledregistriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMaliciousRiskEnabledregistriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMaliciousRiskEnabledregistriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetMaliciousRiskEnabledregistriesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetMaliciousRiskEnabledregistriesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/malicious-risk/enabledRegistries] GetMaliciousRiskEnabledregistries", response, response.Code())
	}
}

// NewGetMaliciousRiskEnabledregistriesOK creates a GetMaliciousRiskEnabledregistriesOK with default headers values
func NewGetMaliciousRiskEnabledregistriesOK() *GetMaliciousRiskEnabledregistriesOK {
	return &GetMaliciousRiskEnabledregistriesOK{}
}

/*
GetMaliciousRiskEnabledregistriesOK describes a response with status code 200, with default header values.

RHC Enabled registries returned
*/
type GetMaliciousRiskEnabledregistriesOK struct {
}

// IsSuccess returns true when this get malicious risk enabledregistries o k response has a 2xx status code
func (o *GetMaliciousRiskEnabledregistriesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get malicious risk enabledregistries o k response has a 3xx status code
func (o *GetMaliciousRiskEnabledregistriesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get malicious risk enabledregistries o k response has a 4xx status code
func (o *GetMaliciousRiskEnabledregistriesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get malicious risk enabledregistries o k response has a 5xx status code
func (o *GetMaliciousRiskEnabledregistriesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get malicious risk enabledregistries o k response a status code equal to that given
func (o *GetMaliciousRiskEnabledregistriesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get malicious risk enabledregistries o k response
func (o *GetMaliciousRiskEnabledregistriesOK) Code() int {
	return 200
}

func (o *GetMaliciousRiskEnabledregistriesOK) Error() string {
	return fmt.Sprintf("[GET /v1/malicious-risk/enabledRegistries][%d] getMaliciousRiskEnabledregistriesOK", 200)
}

func (o *GetMaliciousRiskEnabledregistriesOK) String() string {
	return fmt.Sprintf("[GET /v1/malicious-risk/enabledRegistries][%d] getMaliciousRiskEnabledregistriesOK", 200)
}

func (o *GetMaliciousRiskEnabledregistriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetMaliciousRiskEnabledregistriesUnauthorized creates a GetMaliciousRiskEnabledregistriesUnauthorized with default headers values
func NewGetMaliciousRiskEnabledregistriesUnauthorized() *GetMaliciousRiskEnabledregistriesUnauthorized {
	return &GetMaliciousRiskEnabledregistriesUnauthorized{}
}

/*
GetMaliciousRiskEnabledregistriesUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type GetMaliciousRiskEnabledregistriesUnauthorized struct {
}

// IsSuccess returns true when this get malicious risk enabledregistries unauthorized response has a 2xx status code
func (o *GetMaliciousRiskEnabledregistriesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get malicious risk enabledregistries unauthorized response has a 3xx status code
func (o *GetMaliciousRiskEnabledregistriesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get malicious risk enabledregistries unauthorized response has a 4xx status code
func (o *GetMaliciousRiskEnabledregistriesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get malicious risk enabledregistries unauthorized response has a 5xx status code
func (o *GetMaliciousRiskEnabledregistriesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get malicious risk enabledregistries unauthorized response a status code equal to that given
func (o *GetMaliciousRiskEnabledregistriesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get malicious risk enabledregistries unauthorized response
func (o *GetMaliciousRiskEnabledregistriesUnauthorized) Code() int {
	return 401
}

func (o *GetMaliciousRiskEnabledregistriesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/malicious-risk/enabledRegistries][%d] getMaliciousRiskEnabledregistriesUnauthorized", 401)
}

func (o *GetMaliciousRiskEnabledregistriesUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/malicious-risk/enabledRegistries][%d] getMaliciousRiskEnabledregistriesUnauthorized", 401)
}

func (o *GetMaliciousRiskEnabledregistriesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetMaliciousRiskEnabledregistriesForbidden creates a GetMaliciousRiskEnabledregistriesForbidden with default headers values
func NewGetMaliciousRiskEnabledregistriesForbidden() *GetMaliciousRiskEnabledregistriesForbidden {
	return &GetMaliciousRiskEnabledregistriesForbidden{}
}

/*
GetMaliciousRiskEnabledregistriesForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type GetMaliciousRiskEnabledregistriesForbidden struct {
}

// IsSuccess returns true when this get malicious risk enabledregistries forbidden response has a 2xx status code
func (o *GetMaliciousRiskEnabledregistriesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get malicious risk enabledregistries forbidden response has a 3xx status code
func (o *GetMaliciousRiskEnabledregistriesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get malicious risk enabledregistries forbidden response has a 4xx status code
func (o *GetMaliciousRiskEnabledregistriesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get malicious risk enabledregistries forbidden response has a 5xx status code
func (o *GetMaliciousRiskEnabledregistriesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get malicious risk enabledregistries forbidden response a status code equal to that given
func (o *GetMaliciousRiskEnabledregistriesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get malicious risk enabledregistries forbidden response
func (o *GetMaliciousRiskEnabledregistriesForbidden) Code() int {
	return 403
}

func (o *GetMaliciousRiskEnabledregistriesForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/malicious-risk/enabledRegistries][%d] getMaliciousRiskEnabledregistriesForbidden", 403)
}

func (o *GetMaliciousRiskEnabledregistriesForbidden) String() string {
	return fmt.Sprintf("[GET /v1/malicious-risk/enabledRegistries][%d] getMaliciousRiskEnabledregistriesForbidden", 403)
}

func (o *GetMaliciousRiskEnabledregistriesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
