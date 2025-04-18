// Code generated by go-swagger; DO NOT EDIT.

package support

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostSupportSupportzippathReader is a Reader for the PostSupportSupportzippath structure.
type PostSupportSupportzippathReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSupportSupportzippathReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostSupportSupportzippathOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPostSupportSupportzippathForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/support/supportzippath] PostSupportSupportzippath", response, response.Code())
	}
}

// NewPostSupportSupportzippathOK creates a PostSupportSupportzippathOK with default headers values
func NewPostSupportSupportzippathOK() *PostSupportSupportzippathOK {
	return &PostSupportSupportzippathOK{}
}

/*
PostSupportSupportzippathOK describes a response with status code 200, with default header values.

successful operation
*/
type PostSupportSupportzippathOK struct {
}

// IsSuccess returns true when this post support supportzippath o k response has a 2xx status code
func (o *PostSupportSupportzippathOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post support supportzippath o k response has a 3xx status code
func (o *PostSupportSupportzippathOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post support supportzippath o k response has a 4xx status code
func (o *PostSupportSupportzippathOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post support supportzippath o k response has a 5xx status code
func (o *PostSupportSupportzippathOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post support supportzippath o k response a status code equal to that given
func (o *PostSupportSupportzippathOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post support supportzippath o k response
func (o *PostSupportSupportzippathOK) Code() int {
	return 200
}

func (o *PostSupportSupportzippathOK) Error() string {
	return fmt.Sprintf("[POST /v1/support/supportzippath][%d] postSupportSupportzippathOK", 200)
}

func (o *PostSupportSupportzippathOK) String() string {
	return fmt.Sprintf("[POST /v1/support/supportzippath][%d] postSupportSupportzippathOK", 200)
}

func (o *PostSupportSupportzippathOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSupportSupportzippathForbidden creates a PostSupportSupportzippathForbidden with default headers values
func NewPostSupportSupportzippathForbidden() *PostSupportSupportzippathForbidden {
	return &PostSupportSupportzippathForbidden{}
}

/*
PostSupportSupportzippathForbidden describes a response with status code 403, with default header values.

Insufficient permissions to generate support zip
*/
type PostSupportSupportzippathForbidden struct {
}

// IsSuccess returns true when this post support supportzippath forbidden response has a 2xx status code
func (o *PostSupportSupportzippathForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post support supportzippath forbidden response has a 3xx status code
func (o *PostSupportSupportzippathForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post support supportzippath forbidden response has a 4xx status code
func (o *PostSupportSupportzippathForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post support supportzippath forbidden response has a 5xx status code
func (o *PostSupportSupportzippathForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post support supportzippath forbidden response a status code equal to that given
func (o *PostSupportSupportzippathForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the post support supportzippath forbidden response
func (o *PostSupportSupportzippathForbidden) Code() int {
	return 403
}

func (o *PostSupportSupportzippathForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/support/supportzippath][%d] postSupportSupportzippathForbidden", 403)
}

func (o *PostSupportSupportzippathForbidden) String() string {
	return fmt.Sprintf("[POST /v1/support/supportzippath][%d] postSupportSupportzippathForbidden", 403)
}

func (o *PostSupportSupportzippathForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
