// Code generated by go-swagger; DO NOT EDIT.

package azure_blob_store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostAzureblobstoreTestConnectionReader is a Reader for the PostAzureblobstoreTestConnection structure.
type PostAzureblobstoreTestConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAzureblobstoreTestConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPostAzureblobstoreTestConnectionNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostAzureblobstoreTestConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostAzureblobstoreTestConnectionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostAzureblobstoreTestConnectionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/azureblobstore/test-connection] PostAzureblobstoreTestConnection", response, response.Code())
	}
}

// NewPostAzureblobstoreTestConnectionNoContent creates a PostAzureblobstoreTestConnectionNoContent with default headers values
func NewPostAzureblobstoreTestConnectionNoContent() *PostAzureblobstoreTestConnectionNoContent {
	return &PostAzureblobstoreTestConnectionNoContent{}
}

/*
PostAzureblobstoreTestConnectionNoContent describes a response with status code 204, with default header values.

Azure Blob Store connection was successful
*/
type PostAzureblobstoreTestConnectionNoContent struct {
}

// IsSuccess returns true when this post azureblobstore test connection no content response has a 2xx status code
func (o *PostAzureblobstoreTestConnectionNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post azureblobstore test connection no content response has a 3xx status code
func (o *PostAzureblobstoreTestConnectionNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post azureblobstore test connection no content response has a 4xx status code
func (o *PostAzureblobstoreTestConnectionNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this post azureblobstore test connection no content response has a 5xx status code
func (o *PostAzureblobstoreTestConnectionNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this post azureblobstore test connection no content response a status code equal to that given
func (o *PostAzureblobstoreTestConnectionNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the post azureblobstore test connection no content response
func (o *PostAzureblobstoreTestConnectionNoContent) Code() int {
	return 204
}

func (o *PostAzureblobstoreTestConnectionNoContent) Error() string {
	return fmt.Sprintf("[POST /v1/azureblobstore/test-connection][%d] postAzureblobstoreTestConnectionNoContent", 204)
}

func (o *PostAzureblobstoreTestConnectionNoContent) String() string {
	return fmt.Sprintf("[POST /v1/azureblobstore/test-connection][%d] postAzureblobstoreTestConnectionNoContent", 204)
}

func (o *PostAzureblobstoreTestConnectionNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAzureblobstoreTestConnectionBadRequest creates a PostAzureblobstoreTestConnectionBadRequest with default headers values
func NewPostAzureblobstoreTestConnectionBadRequest() *PostAzureblobstoreTestConnectionBadRequest {
	return &PostAzureblobstoreTestConnectionBadRequest{}
}

/*
PostAzureblobstoreTestConnectionBadRequest describes a response with status code 400, with default header values.

Azure Blob Store connection failed
*/
type PostAzureblobstoreTestConnectionBadRequest struct {
}

// IsSuccess returns true when this post azureblobstore test connection bad request response has a 2xx status code
func (o *PostAzureblobstoreTestConnectionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post azureblobstore test connection bad request response has a 3xx status code
func (o *PostAzureblobstoreTestConnectionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post azureblobstore test connection bad request response has a 4xx status code
func (o *PostAzureblobstoreTestConnectionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post azureblobstore test connection bad request response has a 5xx status code
func (o *PostAzureblobstoreTestConnectionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post azureblobstore test connection bad request response a status code equal to that given
func (o *PostAzureblobstoreTestConnectionBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post azureblobstore test connection bad request response
func (o *PostAzureblobstoreTestConnectionBadRequest) Code() int {
	return 400
}

func (o *PostAzureblobstoreTestConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/azureblobstore/test-connection][%d] postAzureblobstoreTestConnectionBadRequest", 400)
}

func (o *PostAzureblobstoreTestConnectionBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/azureblobstore/test-connection][%d] postAzureblobstoreTestConnectionBadRequest", 400)
}

func (o *PostAzureblobstoreTestConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAzureblobstoreTestConnectionUnauthorized creates a PostAzureblobstoreTestConnectionUnauthorized with default headers values
func NewPostAzureblobstoreTestConnectionUnauthorized() *PostAzureblobstoreTestConnectionUnauthorized {
	return &PostAzureblobstoreTestConnectionUnauthorized{}
}

/*
PostAzureblobstoreTestConnectionUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type PostAzureblobstoreTestConnectionUnauthorized struct {
}

// IsSuccess returns true when this post azureblobstore test connection unauthorized response has a 2xx status code
func (o *PostAzureblobstoreTestConnectionUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post azureblobstore test connection unauthorized response has a 3xx status code
func (o *PostAzureblobstoreTestConnectionUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post azureblobstore test connection unauthorized response has a 4xx status code
func (o *PostAzureblobstoreTestConnectionUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post azureblobstore test connection unauthorized response has a 5xx status code
func (o *PostAzureblobstoreTestConnectionUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post azureblobstore test connection unauthorized response a status code equal to that given
func (o *PostAzureblobstoreTestConnectionUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the post azureblobstore test connection unauthorized response
func (o *PostAzureblobstoreTestConnectionUnauthorized) Code() int {
	return 401
}

func (o *PostAzureblobstoreTestConnectionUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/azureblobstore/test-connection][%d] postAzureblobstoreTestConnectionUnauthorized", 401)
}

func (o *PostAzureblobstoreTestConnectionUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/azureblobstore/test-connection][%d] postAzureblobstoreTestConnectionUnauthorized", 401)
}

func (o *PostAzureblobstoreTestConnectionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAzureblobstoreTestConnectionForbidden creates a PostAzureblobstoreTestConnectionForbidden with default headers values
func NewPostAzureblobstoreTestConnectionForbidden() *PostAzureblobstoreTestConnectionForbidden {
	return &PostAzureblobstoreTestConnectionForbidden{}
}

/*
PostAzureblobstoreTestConnectionForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type PostAzureblobstoreTestConnectionForbidden struct {
}

// IsSuccess returns true when this post azureblobstore test connection forbidden response has a 2xx status code
func (o *PostAzureblobstoreTestConnectionForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post azureblobstore test connection forbidden response has a 3xx status code
func (o *PostAzureblobstoreTestConnectionForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post azureblobstore test connection forbidden response has a 4xx status code
func (o *PostAzureblobstoreTestConnectionForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post azureblobstore test connection forbidden response has a 5xx status code
func (o *PostAzureblobstoreTestConnectionForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post azureblobstore test connection forbidden response a status code equal to that given
func (o *PostAzureblobstoreTestConnectionForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the post azureblobstore test connection forbidden response
func (o *PostAzureblobstoreTestConnectionForbidden) Code() int {
	return 403
}

func (o *PostAzureblobstoreTestConnectionForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/azureblobstore/test-connection][%d] postAzureblobstoreTestConnectionForbidden", 403)
}

func (o *PostAzureblobstoreTestConnectionForbidden) String() string {
	return fmt.Sprintf("[POST /v1/azureblobstore/test-connection][%d] postAzureblobstoreTestConnectionForbidden", 403)
}

func (o *PostAzureblobstoreTestConnectionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
