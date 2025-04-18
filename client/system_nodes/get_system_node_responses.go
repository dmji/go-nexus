// Code generated by go-swagger; DO NOT EDIT.

package system_nodes

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

// GetSystemNodeReader is a Reader for the GetSystemNode structure.
type GetSystemNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSystemNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSystemNodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetSystemNodeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/system/node] GetSystemNode", response, response.Code())
	}
}

// NewGetSystemNodeOK creates a GetSystemNodeOK with default headers values
func NewGetSystemNodeOK() *GetSystemNodeOK {
	return &GetSystemNodeOK{}
}

/*
GetSystemNodeOK describes a response with status code 200, with default header values.

successful operation
*/
type GetSystemNodeOK struct {
	Payload *models.NodeInformation
}

// IsSuccess returns true when this get system node o k response has a 2xx status code
func (o *GetSystemNodeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get system node o k response has a 3xx status code
func (o *GetSystemNodeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get system node o k response has a 4xx status code
func (o *GetSystemNodeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get system node o k response has a 5xx status code
func (o *GetSystemNodeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get system node o k response a status code equal to that given
func (o *GetSystemNodeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get system node o k response
func (o *GetSystemNodeOK) Code() int {
	return 200
}

func (o *GetSystemNodeOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/system/node][%d] getSystemNodeOK %s", 200, payload)
}

func (o *GetSystemNodeOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/system/node][%d] getSystemNodeOK %s", 200, payload)
}

func (o *GetSystemNodeOK) GetPayload() *models.NodeInformation {
	return o.Payload
}

func (o *GetSystemNodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NodeInformation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSystemNodeForbidden creates a GetSystemNodeForbidden with default headers values
func NewGetSystemNodeForbidden() *GetSystemNodeForbidden {
	return &GetSystemNodeForbidden{}
}

/*
GetSystemNodeForbidden describes a response with status code 403, with default header values.

Insufficient permissions to update settings
*/
type GetSystemNodeForbidden struct {
}

// IsSuccess returns true when this get system node forbidden response has a 2xx status code
func (o *GetSystemNodeForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get system node forbidden response has a 3xx status code
func (o *GetSystemNodeForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get system node forbidden response has a 4xx status code
func (o *GetSystemNodeForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get system node forbidden response has a 5xx status code
func (o *GetSystemNodeForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get system node forbidden response a status code equal to that given
func (o *GetSystemNodeForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get system node forbidden response
func (o *GetSystemNodeForbidden) Code() int {
	return 403
}

func (o *GetSystemNodeForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/system/node][%d] getSystemNodeForbidden", 403)
}

func (o *GetSystemNodeForbidden) String() string {
	return fmt.Sprintf("[GET /v1/system/node][%d] getSystemNodeForbidden", 403)
}

func (o *GetSystemNodeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
