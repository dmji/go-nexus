// Code generated by go-swagger; DO NOT EDIT.

package email

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteEmailConfigurationReader is a Reader for the DeleteEmailConfiguration structure.
type DeleteEmailConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteEmailConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteEmailConfigurationNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /v1/email] deleteEmailConfiguration", response, response.Code())
	}
}

// NewDeleteEmailConfigurationNoContent creates a DeleteEmailConfigurationNoContent with default headers values
func NewDeleteEmailConfigurationNoContent() *DeleteEmailConfigurationNoContent {
	return &DeleteEmailConfigurationNoContent{}
}

/*
DeleteEmailConfigurationNoContent describes a response with status code 204, with default header values.

Email configuration was successfully cleared
*/
type DeleteEmailConfigurationNoContent struct {
}

// IsSuccess returns true when this delete email configuration no content response has a 2xx status code
func (o *DeleteEmailConfigurationNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete email configuration no content response has a 3xx status code
func (o *DeleteEmailConfigurationNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete email configuration no content response has a 4xx status code
func (o *DeleteEmailConfigurationNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete email configuration no content response has a 5xx status code
func (o *DeleteEmailConfigurationNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete email configuration no content response a status code equal to that given
func (o *DeleteEmailConfigurationNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete email configuration no content response
func (o *DeleteEmailConfigurationNoContent) Code() int {
	return 204
}

func (o *DeleteEmailConfigurationNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/email][%d] deleteEmailConfigurationNoContent", 204)
}

func (o *DeleteEmailConfigurationNoContent) String() string {
	return fmt.Sprintf("[DELETE /v1/email][%d] deleteEmailConfigurationNoContent", 204)
}

func (o *DeleteEmailConfigurationNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
