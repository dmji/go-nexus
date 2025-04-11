// Code generated by go-swagger; DO NOT EDIT.

package content_selectors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/dmji/go-nexus/models"
)

// GetContentSelectorReader is a Reader for the GetContentSelector structure.
type GetContentSelectorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetContentSelectorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetContentSelectorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetContentSelectorForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/security/content-selectors/{name}] getContentSelector", response, response.Code())
	}
}

// NewGetContentSelectorOK creates a GetContentSelectorOK with default headers values
func NewGetContentSelectorOK() *GetContentSelectorOK {
	return &GetContentSelectorOK{}
}

/*
GetContentSelectorOK describes a response with status code 200, with default header values.

successful operation
*/
type GetContentSelectorOK struct {
	Payload *models.ContentSelectorAPIResponse
}

// IsSuccess returns true when this get content selector o k response has a 2xx status code
func (o *GetContentSelectorOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get content selector o k response has a 3xx status code
func (o *GetContentSelectorOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get content selector o k response has a 4xx status code
func (o *GetContentSelectorOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get content selector o k response has a 5xx status code
func (o *GetContentSelectorOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get content selector o k response a status code equal to that given
func (o *GetContentSelectorOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get content selector o k response
func (o *GetContentSelectorOK) Code() int {
	return 200
}

func (o *GetContentSelectorOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/security/content-selectors/{name}][%d] getContentSelectorOK %s", 200, payload)
}

func (o *GetContentSelectorOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/security/content-selectors/{name}][%d] getContentSelectorOK %s", 200, payload)
}

func (o *GetContentSelectorOK) GetPayload() *models.ContentSelectorAPIResponse {
	return o.Payload
}

func (o *GetContentSelectorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContentSelectorAPIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentSelectorForbidden creates a GetContentSelectorForbidden with default headers values
func NewGetContentSelectorForbidden() *GetContentSelectorForbidden {
	return &GetContentSelectorForbidden{}
}

/*
GetContentSelectorForbidden describes a response with status code 403, with default header values.

Insufficient permissions to read the content selector
*/
type GetContentSelectorForbidden struct {
}

// IsSuccess returns true when this get content selector forbidden response has a 2xx status code
func (o *GetContentSelectorForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get content selector forbidden response has a 3xx status code
func (o *GetContentSelectorForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get content selector forbidden response has a 4xx status code
func (o *GetContentSelectorForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get content selector forbidden response has a 5xx status code
func (o *GetContentSelectorForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get content selector forbidden response a status code equal to that given
func (o *GetContentSelectorForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get content selector forbidden response
func (o *GetContentSelectorForbidden) Code() int {
	return 403
}

func (o *GetContentSelectorForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/security/content-selectors/{name}][%d] getContentSelectorForbidden", 403)
}

func (o *GetContentSelectorForbidden) String() string {
	return fmt.Sprintf("[GET /v1/security/content-selectors/{name}][%d] getContentSelectorForbidden", 403)
}

func (o *GetContentSelectorForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
