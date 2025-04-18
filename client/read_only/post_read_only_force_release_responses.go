// Code generated by go-swagger; DO NOT EDIT.

package read_only

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostReadOnlyForceReleaseReader is a Reader for the PostReadOnlyForceRelease structure.
type PostReadOnlyForceReleaseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostReadOnlyForceReleaseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPostReadOnlyForceReleaseNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPostReadOnlyForceReleaseForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostReadOnlyForceReleaseNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/read-only/force-release] PostReadOnlyForceRelease", response, response.Code())
	}
}

// NewPostReadOnlyForceReleaseNoContent creates a PostReadOnlyForceReleaseNoContent with default headers values
func NewPostReadOnlyForceReleaseNoContent() *PostReadOnlyForceReleaseNoContent {
	return &PostReadOnlyForceReleaseNoContent{}
}

/*
PostReadOnlyForceReleaseNoContent describes a response with status code 204, with default header values.

Database is no longer read-only
*/
type PostReadOnlyForceReleaseNoContent struct {
}

// IsSuccess returns true when this post read only force release no content response has a 2xx status code
func (o *PostReadOnlyForceReleaseNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post read only force release no content response has a 3xx status code
func (o *PostReadOnlyForceReleaseNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post read only force release no content response has a 4xx status code
func (o *PostReadOnlyForceReleaseNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this post read only force release no content response has a 5xx status code
func (o *PostReadOnlyForceReleaseNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this post read only force release no content response a status code equal to that given
func (o *PostReadOnlyForceReleaseNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the post read only force release no content response
func (o *PostReadOnlyForceReleaseNoContent) Code() int {
	return 204
}

func (o *PostReadOnlyForceReleaseNoContent) Error() string {
	return fmt.Sprintf("[POST /v1/read-only/force-release][%d] postReadOnlyForceReleaseNoContent", 204)
}

func (o *PostReadOnlyForceReleaseNoContent) String() string {
	return fmt.Sprintf("[POST /v1/read-only/force-release][%d] postReadOnlyForceReleaseNoContent", 204)
}

func (o *PostReadOnlyForceReleaseNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostReadOnlyForceReleaseForbidden creates a PostReadOnlyForceReleaseForbidden with default headers values
func NewPostReadOnlyForceReleaseForbidden() *PostReadOnlyForceReleaseForbidden {
	return &PostReadOnlyForceReleaseForbidden{}
}

/*
PostReadOnlyForceReleaseForbidden describes a response with status code 403, with default header values.

Authentication required
*/
type PostReadOnlyForceReleaseForbidden struct {
}

// IsSuccess returns true when this post read only force release forbidden response has a 2xx status code
func (o *PostReadOnlyForceReleaseForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post read only force release forbidden response has a 3xx status code
func (o *PostReadOnlyForceReleaseForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post read only force release forbidden response has a 4xx status code
func (o *PostReadOnlyForceReleaseForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post read only force release forbidden response has a 5xx status code
func (o *PostReadOnlyForceReleaseForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post read only force release forbidden response a status code equal to that given
func (o *PostReadOnlyForceReleaseForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the post read only force release forbidden response
func (o *PostReadOnlyForceReleaseForbidden) Code() int {
	return 403
}

func (o *PostReadOnlyForceReleaseForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/read-only/force-release][%d] postReadOnlyForceReleaseForbidden", 403)
}

func (o *PostReadOnlyForceReleaseForbidden) String() string {
	return fmt.Sprintf("[POST /v1/read-only/force-release][%d] postReadOnlyForceReleaseForbidden", 403)
}

func (o *PostReadOnlyForceReleaseForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostReadOnlyForceReleaseNotFound creates a PostReadOnlyForceReleaseNotFound with default headers values
func NewPostReadOnlyForceReleaseNotFound() *PostReadOnlyForceReleaseNotFound {
	return &PostReadOnlyForceReleaseNotFound{}
}

/*
PostReadOnlyForceReleaseNotFound describes a response with status code 404, with default header values.

No change to read-only state
*/
type PostReadOnlyForceReleaseNotFound struct {
}

// IsSuccess returns true when this post read only force release not found response has a 2xx status code
func (o *PostReadOnlyForceReleaseNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post read only force release not found response has a 3xx status code
func (o *PostReadOnlyForceReleaseNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post read only force release not found response has a 4xx status code
func (o *PostReadOnlyForceReleaseNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post read only force release not found response has a 5xx status code
func (o *PostReadOnlyForceReleaseNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post read only force release not found response a status code equal to that given
func (o *PostReadOnlyForceReleaseNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the post read only force release not found response
func (o *PostReadOnlyForceReleaseNotFound) Code() int {
	return 404
}

func (o *PostReadOnlyForceReleaseNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/read-only/force-release][%d] postReadOnlyForceReleaseNotFound", 404)
}

func (o *PostReadOnlyForceReleaseNotFound) String() string {
	return fmt.Sprintf("[POST /v1/read-only/force-release][%d] postReadOnlyForceReleaseNotFound", 404)
}

func (o *PostReadOnlyForceReleaseNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
