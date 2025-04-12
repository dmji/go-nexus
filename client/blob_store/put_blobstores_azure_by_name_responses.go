// Code generated by go-swagger; DO NOT EDIT.

package blob_store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PutBlobstoresAzureByNameReader is a Reader for the PutBlobstoresAzureByName structure.
type PutBlobstoresAzureByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutBlobstoresAzureByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutBlobstoresAzureByNameNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutBlobstoresAzureByNameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutBlobstoresAzureByNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutBlobstoresAzureByNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /v1/blobstores/azure/{name}] PutBlobstoresAzureByName", response, response.Code())
	}
}

// NewPutBlobstoresAzureByNameNoContent creates a PutBlobstoresAzureByNameNoContent with default headers values
func NewPutBlobstoresAzureByNameNoContent() *PutBlobstoresAzureByNameNoContent {
	return &PutBlobstoresAzureByNameNoContent{}
}

/*
PutBlobstoresAzureByNameNoContent describes a response with status code 204, with default header values.

Azure blob store updated
*/
type PutBlobstoresAzureByNameNoContent struct {
}

// IsSuccess returns true when this put blobstores azure by name no content response has a 2xx status code
func (o *PutBlobstoresAzureByNameNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put blobstores azure by name no content response has a 3xx status code
func (o *PutBlobstoresAzureByNameNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put blobstores azure by name no content response has a 4xx status code
func (o *PutBlobstoresAzureByNameNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this put blobstores azure by name no content response has a 5xx status code
func (o *PutBlobstoresAzureByNameNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this put blobstores azure by name no content response a status code equal to that given
func (o *PutBlobstoresAzureByNameNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the put blobstores azure by name no content response
func (o *PutBlobstoresAzureByNameNoContent) Code() int {
	return 204
}

func (o *PutBlobstoresAzureByNameNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/blobstores/azure/{name}][%d] putBlobstoresAzureByNameNoContent", 204)
}

func (o *PutBlobstoresAzureByNameNoContent) String() string {
	return fmt.Sprintf("[PUT /v1/blobstores/azure/{name}][%d] putBlobstoresAzureByNameNoContent", 204)
}

func (o *PutBlobstoresAzureByNameNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutBlobstoresAzureByNameBadRequest creates a PutBlobstoresAzureByNameBadRequest with default headers values
func NewPutBlobstoresAzureByNameBadRequest() *PutBlobstoresAzureByNameBadRequest {
	return &PutBlobstoresAzureByNameBadRequest{}
}

/*
PutBlobstoresAzureByNameBadRequest describes a response with status code 400, with default header values.

Specified Azure blob store doesn't exist
*/
type PutBlobstoresAzureByNameBadRequest struct {
}

// IsSuccess returns true when this put blobstores azure by name bad request response has a 2xx status code
func (o *PutBlobstoresAzureByNameBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put blobstores azure by name bad request response has a 3xx status code
func (o *PutBlobstoresAzureByNameBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put blobstores azure by name bad request response has a 4xx status code
func (o *PutBlobstoresAzureByNameBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put blobstores azure by name bad request response has a 5xx status code
func (o *PutBlobstoresAzureByNameBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put blobstores azure by name bad request response a status code equal to that given
func (o *PutBlobstoresAzureByNameBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the put blobstores azure by name bad request response
func (o *PutBlobstoresAzureByNameBadRequest) Code() int {
	return 400
}

func (o *PutBlobstoresAzureByNameBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/blobstores/azure/{name}][%d] putBlobstoresAzureByNameBadRequest", 400)
}

func (o *PutBlobstoresAzureByNameBadRequest) String() string {
	return fmt.Sprintf("[PUT /v1/blobstores/azure/{name}][%d] putBlobstoresAzureByNameBadRequest", 400)
}

func (o *PutBlobstoresAzureByNameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutBlobstoresAzureByNameUnauthorized creates a PutBlobstoresAzureByNameUnauthorized with default headers values
func NewPutBlobstoresAzureByNameUnauthorized() *PutBlobstoresAzureByNameUnauthorized {
	return &PutBlobstoresAzureByNameUnauthorized{}
}

/*
PutBlobstoresAzureByNameUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type PutBlobstoresAzureByNameUnauthorized struct {
}

// IsSuccess returns true when this put blobstores azure by name unauthorized response has a 2xx status code
func (o *PutBlobstoresAzureByNameUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put blobstores azure by name unauthorized response has a 3xx status code
func (o *PutBlobstoresAzureByNameUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put blobstores azure by name unauthorized response has a 4xx status code
func (o *PutBlobstoresAzureByNameUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put blobstores azure by name unauthorized response has a 5xx status code
func (o *PutBlobstoresAzureByNameUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put blobstores azure by name unauthorized response a status code equal to that given
func (o *PutBlobstoresAzureByNameUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the put blobstores azure by name unauthorized response
func (o *PutBlobstoresAzureByNameUnauthorized) Code() int {
	return 401
}

func (o *PutBlobstoresAzureByNameUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/blobstores/azure/{name}][%d] putBlobstoresAzureByNameUnauthorized", 401)
}

func (o *PutBlobstoresAzureByNameUnauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/blobstores/azure/{name}][%d] putBlobstoresAzureByNameUnauthorized", 401)
}

func (o *PutBlobstoresAzureByNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutBlobstoresAzureByNameForbidden creates a PutBlobstoresAzureByNameForbidden with default headers values
func NewPutBlobstoresAzureByNameForbidden() *PutBlobstoresAzureByNameForbidden {
	return &PutBlobstoresAzureByNameForbidden{}
}

/*
PutBlobstoresAzureByNameForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type PutBlobstoresAzureByNameForbidden struct {
}

// IsSuccess returns true when this put blobstores azure by name forbidden response has a 2xx status code
func (o *PutBlobstoresAzureByNameForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put blobstores azure by name forbidden response has a 3xx status code
func (o *PutBlobstoresAzureByNameForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put blobstores azure by name forbidden response has a 4xx status code
func (o *PutBlobstoresAzureByNameForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put blobstores azure by name forbidden response has a 5xx status code
func (o *PutBlobstoresAzureByNameForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put blobstores azure by name forbidden response a status code equal to that given
func (o *PutBlobstoresAzureByNameForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the put blobstores azure by name forbidden response
func (o *PutBlobstoresAzureByNameForbidden) Code() int {
	return 403
}

func (o *PutBlobstoresAzureByNameForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/blobstores/azure/{name}][%d] putBlobstoresAzureByNameForbidden", 403)
}

func (o *PutBlobstoresAzureByNameForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/blobstores/azure/{name}][%d] putBlobstoresAzureByNameForbidden", 403)
}

func (o *PutBlobstoresAzureByNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
