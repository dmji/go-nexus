// Code generated by go-swagger; DO NOT EDIT.

package content_selectors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteSecurityContentSelectorsByNameReader is a Reader for the DeleteSecurityContentSelectorsByName structure.
type DeleteSecurityContentSelectorsByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSecurityContentSelectorsByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteSecurityContentSelectorsByNameNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteSecurityContentSelectorsByNameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteSecurityContentSelectorsByNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /v1/security/content-selectors/{name}] DeleteSecurityContentSelectorsByName", response, response.Code())
	}
}

// NewDeleteSecurityContentSelectorsByNameNoContent creates a DeleteSecurityContentSelectorsByNameNoContent with default headers values
func NewDeleteSecurityContentSelectorsByNameNoContent() *DeleteSecurityContentSelectorsByNameNoContent {
	return &DeleteSecurityContentSelectorsByNameNoContent{}
}

/*
DeleteSecurityContentSelectorsByNameNoContent describes a response with status code 204, with default header values.

Content selector deleted successfully
*/
type DeleteSecurityContentSelectorsByNameNoContent struct {
}

// IsSuccess returns true when this delete security content selectors by name no content response has a 2xx status code
func (o *DeleteSecurityContentSelectorsByNameNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete security content selectors by name no content response has a 3xx status code
func (o *DeleteSecurityContentSelectorsByNameNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete security content selectors by name no content response has a 4xx status code
func (o *DeleteSecurityContentSelectorsByNameNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete security content selectors by name no content response has a 5xx status code
func (o *DeleteSecurityContentSelectorsByNameNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete security content selectors by name no content response a status code equal to that given
func (o *DeleteSecurityContentSelectorsByNameNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete security content selectors by name no content response
func (o *DeleteSecurityContentSelectorsByNameNoContent) Code() int {
	return 204
}

func (o *DeleteSecurityContentSelectorsByNameNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/security/content-selectors/{name}][%d] deleteSecurityContentSelectorsByNameNoContent", 204)
}

func (o *DeleteSecurityContentSelectorsByNameNoContent) String() string {
	return fmt.Sprintf("[DELETE /v1/security/content-selectors/{name}][%d] deleteSecurityContentSelectorsByNameNoContent", 204)
}

func (o *DeleteSecurityContentSelectorsByNameNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSecurityContentSelectorsByNameBadRequest creates a DeleteSecurityContentSelectorsByNameBadRequest with default headers values
func NewDeleteSecurityContentSelectorsByNameBadRequest() *DeleteSecurityContentSelectorsByNameBadRequest {
	return &DeleteSecurityContentSelectorsByNameBadRequest{}
}

/*
DeleteSecurityContentSelectorsByNameBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type DeleteSecurityContentSelectorsByNameBadRequest struct {
}

// IsSuccess returns true when this delete security content selectors by name bad request response has a 2xx status code
func (o *DeleteSecurityContentSelectorsByNameBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete security content selectors by name bad request response has a 3xx status code
func (o *DeleteSecurityContentSelectorsByNameBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete security content selectors by name bad request response has a 4xx status code
func (o *DeleteSecurityContentSelectorsByNameBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete security content selectors by name bad request response has a 5xx status code
func (o *DeleteSecurityContentSelectorsByNameBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete security content selectors by name bad request response a status code equal to that given
func (o *DeleteSecurityContentSelectorsByNameBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete security content selectors by name bad request response
func (o *DeleteSecurityContentSelectorsByNameBadRequest) Code() int {
	return 400
}

func (o *DeleteSecurityContentSelectorsByNameBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /v1/security/content-selectors/{name}][%d] deleteSecurityContentSelectorsByNameBadRequest", 400)
}

func (o *DeleteSecurityContentSelectorsByNameBadRequest) String() string {
	return fmt.Sprintf("[DELETE /v1/security/content-selectors/{name}][%d] deleteSecurityContentSelectorsByNameBadRequest", 400)
}

func (o *DeleteSecurityContentSelectorsByNameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSecurityContentSelectorsByNameForbidden creates a DeleteSecurityContentSelectorsByNameForbidden with default headers values
func NewDeleteSecurityContentSelectorsByNameForbidden() *DeleteSecurityContentSelectorsByNameForbidden {
	return &DeleteSecurityContentSelectorsByNameForbidden{}
}

/*
DeleteSecurityContentSelectorsByNameForbidden describes a response with status code 403, with default header values.

Insufficient permissions to delete the content selector
*/
type DeleteSecurityContentSelectorsByNameForbidden struct {
}

// IsSuccess returns true when this delete security content selectors by name forbidden response has a 2xx status code
func (o *DeleteSecurityContentSelectorsByNameForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete security content selectors by name forbidden response has a 3xx status code
func (o *DeleteSecurityContentSelectorsByNameForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete security content selectors by name forbidden response has a 4xx status code
func (o *DeleteSecurityContentSelectorsByNameForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete security content selectors by name forbidden response has a 5xx status code
func (o *DeleteSecurityContentSelectorsByNameForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete security content selectors by name forbidden response a status code equal to that given
func (o *DeleteSecurityContentSelectorsByNameForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete security content selectors by name forbidden response
func (o *DeleteSecurityContentSelectorsByNameForbidden) Code() int {
	return 403
}

func (o *DeleteSecurityContentSelectorsByNameForbidden) Error() string {
	return fmt.Sprintf("[DELETE /v1/security/content-selectors/{name}][%d] deleteSecurityContentSelectorsByNameForbidden", 403)
}

func (o *DeleteSecurityContentSelectorsByNameForbidden) String() string {
	return fmt.Sprintf("[DELETE /v1/security/content-selectors/{name}][%d] deleteSecurityContentSelectorsByNameForbidden", 403)
}

func (o *DeleteSecurityContentSelectorsByNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
