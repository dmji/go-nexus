// Code generated by go-swagger; DO NOT EDIT.

package script

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

// GetScriptReader is a Reader for the GetScript structure.
type GetScriptReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetScriptReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetScriptOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/script] GetScript", response, response.Code())
	}
}

// NewGetScriptOK creates a GetScriptOK with default headers values
func NewGetScriptOK() *GetScriptOK {
	return &GetScriptOK{}
}

/*
GetScriptOK describes a response with status code 200, with default header values.

successful operation
*/
type GetScriptOK struct {
	Payload []*models.ScriptXO
}

// IsSuccess returns true when this get script o k response has a 2xx status code
func (o *GetScriptOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get script o k response has a 3xx status code
func (o *GetScriptOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get script o k response has a 4xx status code
func (o *GetScriptOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get script o k response has a 5xx status code
func (o *GetScriptOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get script o k response a status code equal to that given
func (o *GetScriptOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get script o k response
func (o *GetScriptOK) Code() int {
	return 200
}

func (o *GetScriptOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/script][%d] getScriptOK %s", 200, payload)
}

func (o *GetScriptOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/script][%d] getScriptOK %s", 200, payload)
}

func (o *GetScriptOK) GetPayload() []*models.ScriptXO {
	return o.Payload
}

func (o *GetScriptOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
