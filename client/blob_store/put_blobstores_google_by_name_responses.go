// Code generated by go-swagger; DO NOT EDIT.

package blob_store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PutBlobstoresGoogleByNameReader is a Reader for the PutBlobstoresGoogleByName structure.
type PutBlobstoresGoogleByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutBlobstoresGoogleByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutBlobstoresGoogleByNameNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPutBlobstoresGoogleByNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutBlobstoresGoogleByNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutBlobstoresGoogleByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /v1/blobstores/google/{name}] PutBlobstoresGoogleByName", response, response.Code())
	}
}

// NewPutBlobstoresGoogleByNameNoContent creates a PutBlobstoresGoogleByNameNoContent with default headers values
func NewPutBlobstoresGoogleByNameNoContent() *PutBlobstoresGoogleByNameNoContent {
	return &PutBlobstoresGoogleByNameNoContent{}
}

/*
PutBlobstoresGoogleByNameNoContent describes a response with status code 204, with default header values.

Google Cloud blob store updated
*/
type PutBlobstoresGoogleByNameNoContent struct {
}

// IsSuccess returns true when this put blobstores google by name no content response has a 2xx status code
func (o *PutBlobstoresGoogleByNameNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put blobstores google by name no content response has a 3xx status code
func (o *PutBlobstoresGoogleByNameNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put blobstores google by name no content response has a 4xx status code
func (o *PutBlobstoresGoogleByNameNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this put blobstores google by name no content response has a 5xx status code
func (o *PutBlobstoresGoogleByNameNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this put blobstores google by name no content response a status code equal to that given
func (o *PutBlobstoresGoogleByNameNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the put blobstores google by name no content response
func (o *PutBlobstoresGoogleByNameNoContent) Code() int {
	return 204
}

func (o *PutBlobstoresGoogleByNameNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/blobstores/google/{name}][%d] putBlobstoresGoogleByNameNoContent", 204)
}

func (o *PutBlobstoresGoogleByNameNoContent) String() string {
	return fmt.Sprintf("[PUT /v1/blobstores/google/{name}][%d] putBlobstoresGoogleByNameNoContent", 204)
}

func (o *PutBlobstoresGoogleByNameNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutBlobstoresGoogleByNameUnauthorized creates a PutBlobstoresGoogleByNameUnauthorized with default headers values
func NewPutBlobstoresGoogleByNameUnauthorized() *PutBlobstoresGoogleByNameUnauthorized {
	return &PutBlobstoresGoogleByNameUnauthorized{}
}

/*
PutBlobstoresGoogleByNameUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type PutBlobstoresGoogleByNameUnauthorized struct {
}

// IsSuccess returns true when this put blobstores google by name unauthorized response has a 2xx status code
func (o *PutBlobstoresGoogleByNameUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put blobstores google by name unauthorized response has a 3xx status code
func (o *PutBlobstoresGoogleByNameUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put blobstores google by name unauthorized response has a 4xx status code
func (o *PutBlobstoresGoogleByNameUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put blobstores google by name unauthorized response has a 5xx status code
func (o *PutBlobstoresGoogleByNameUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put blobstores google by name unauthorized response a status code equal to that given
func (o *PutBlobstoresGoogleByNameUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the put blobstores google by name unauthorized response
func (o *PutBlobstoresGoogleByNameUnauthorized) Code() int {
	return 401
}

func (o *PutBlobstoresGoogleByNameUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/blobstores/google/{name}][%d] putBlobstoresGoogleByNameUnauthorized", 401)
}

func (o *PutBlobstoresGoogleByNameUnauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/blobstores/google/{name}][%d] putBlobstoresGoogleByNameUnauthorized", 401)
}

func (o *PutBlobstoresGoogleByNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutBlobstoresGoogleByNameForbidden creates a PutBlobstoresGoogleByNameForbidden with default headers values
func NewPutBlobstoresGoogleByNameForbidden() *PutBlobstoresGoogleByNameForbidden {
	return &PutBlobstoresGoogleByNameForbidden{}
}

/*
PutBlobstoresGoogleByNameForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type PutBlobstoresGoogleByNameForbidden struct {
}

// IsSuccess returns true when this put blobstores google by name forbidden response has a 2xx status code
func (o *PutBlobstoresGoogleByNameForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put blobstores google by name forbidden response has a 3xx status code
func (o *PutBlobstoresGoogleByNameForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put blobstores google by name forbidden response has a 4xx status code
func (o *PutBlobstoresGoogleByNameForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put blobstores google by name forbidden response has a 5xx status code
func (o *PutBlobstoresGoogleByNameForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put blobstores google by name forbidden response a status code equal to that given
func (o *PutBlobstoresGoogleByNameForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the put blobstores google by name forbidden response
func (o *PutBlobstoresGoogleByNameForbidden) Code() int {
	return 403
}

func (o *PutBlobstoresGoogleByNameForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/blobstores/google/{name}][%d] putBlobstoresGoogleByNameForbidden", 403)
}

func (o *PutBlobstoresGoogleByNameForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/blobstores/google/{name}][%d] putBlobstoresGoogleByNameForbidden", 403)
}

func (o *PutBlobstoresGoogleByNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutBlobstoresGoogleByNameNotFound creates a PutBlobstoresGoogleByNameNotFound with default headers values
func NewPutBlobstoresGoogleByNameNotFound() *PutBlobstoresGoogleByNameNotFound {
	return &PutBlobstoresGoogleByNameNotFound{}
}

/*
PutBlobstoresGoogleByNameNotFound describes a response with status code 404, with default header values.

Repository not found
*/
type PutBlobstoresGoogleByNameNotFound struct {
}

// IsSuccess returns true when this put blobstores google by name not found response has a 2xx status code
func (o *PutBlobstoresGoogleByNameNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put blobstores google by name not found response has a 3xx status code
func (o *PutBlobstoresGoogleByNameNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put blobstores google by name not found response has a 4xx status code
func (o *PutBlobstoresGoogleByNameNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put blobstores google by name not found response has a 5xx status code
func (o *PutBlobstoresGoogleByNameNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put blobstores google by name not found response a status code equal to that given
func (o *PutBlobstoresGoogleByNameNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the put blobstores google by name not found response
func (o *PutBlobstoresGoogleByNameNotFound) Code() int {
	return 404
}

func (o *PutBlobstoresGoogleByNameNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/blobstores/google/{name}][%d] putBlobstoresGoogleByNameNotFound", 404)
}

func (o *PutBlobstoresGoogleByNameNotFound) String() string {
	return fmt.Sprintf("[PUT /v1/blobstores/google/{name}][%d] putBlobstoresGoogleByNameNotFound", 404)
}

func (o *PutBlobstoresGoogleByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
