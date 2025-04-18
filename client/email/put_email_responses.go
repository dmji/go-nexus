// Code generated by go-swagger; DO NOT EDIT.

package email

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PutEmailReader is a Reader for the PutEmail structure.
type PutEmailReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutEmailReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutEmailNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutEmailBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutEmailForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /v1/email] PutEmail", response, response.Code())
	}
}

// NewPutEmailNoContent creates a PutEmailNoContent with default headers values
func NewPutEmailNoContent() *PutEmailNoContent {
	return &PutEmailNoContent{}
}

/*
PutEmailNoContent describes a response with status code 204, with default header values.

Email configuration was successfully updated
*/
type PutEmailNoContent struct {
}

// IsSuccess returns true when this put email no content response has a 2xx status code
func (o *PutEmailNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put email no content response has a 3xx status code
func (o *PutEmailNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put email no content response has a 4xx status code
func (o *PutEmailNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this put email no content response has a 5xx status code
func (o *PutEmailNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this put email no content response a status code equal to that given
func (o *PutEmailNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the put email no content response
func (o *PutEmailNoContent) Code() int {
	return 204
}

func (o *PutEmailNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/email][%d] putEmailNoContent", 204)
}

func (o *PutEmailNoContent) String() string {
	return fmt.Sprintf("[PUT /v1/email][%d] putEmailNoContent", 204)
}

func (o *PutEmailNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutEmailBadRequest creates a PutEmailBadRequest with default headers values
func NewPutEmailBadRequest() *PutEmailBadRequest {
	return &PutEmailBadRequest{}
}

/*
PutEmailBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type PutEmailBadRequest struct {
}

// IsSuccess returns true when this put email bad request response has a 2xx status code
func (o *PutEmailBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put email bad request response has a 3xx status code
func (o *PutEmailBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put email bad request response has a 4xx status code
func (o *PutEmailBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put email bad request response has a 5xx status code
func (o *PutEmailBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put email bad request response a status code equal to that given
func (o *PutEmailBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the put email bad request response
func (o *PutEmailBadRequest) Code() int {
	return 400
}

func (o *PutEmailBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/email][%d] putEmailBadRequest", 400)
}

func (o *PutEmailBadRequest) String() string {
	return fmt.Sprintf("[PUT /v1/email][%d] putEmailBadRequest", 400)
}

func (o *PutEmailBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutEmailForbidden creates a PutEmailForbidden with default headers values
func NewPutEmailForbidden() *PutEmailForbidden {
	return &PutEmailForbidden{}
}

/*
PutEmailForbidden describes a response with status code 403, with default header values.

Insufficient permissions to update the email configuration
*/
type PutEmailForbidden struct {
}

// IsSuccess returns true when this put email forbidden response has a 2xx status code
func (o *PutEmailForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put email forbidden response has a 3xx status code
func (o *PutEmailForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put email forbidden response has a 4xx status code
func (o *PutEmailForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put email forbidden response has a 5xx status code
func (o *PutEmailForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put email forbidden response a status code equal to that given
func (o *PutEmailForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the put email forbidden response
func (o *PutEmailForbidden) Code() int {
	return 403
}

func (o *PutEmailForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/email][%d] putEmailForbidden", 403)
}

func (o *PutEmailForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/email][%d] putEmailForbidden", 403)
}

func (o *PutEmailForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
