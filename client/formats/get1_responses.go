// Code generated by go-swagger; DO NOT EDIT.

package formats

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

// Get1Reader is a Reader for the Get1 structure.
type Get1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Get1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGet1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/formats/{format}/upload-specs] get_1", response, response.Code())
	}
}

// NewGet1OK creates a Get1OK with default headers values
func NewGet1OK() *Get1OK {
	return &Get1OK{}
}

/*
Get1OK describes a response with status code 200, with default header values.

successful operation
*/
type Get1OK struct {
	Payload *models.UploadDefinitionXO
}

// IsSuccess returns true when this get1 o k response has a 2xx status code
func (o *Get1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get1 o k response has a 3xx status code
func (o *Get1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get1 o k response has a 4xx status code
func (o *Get1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get1 o k response has a 5xx status code
func (o *Get1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get1 o k response a status code equal to that given
func (o *Get1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get1 o k response
func (o *Get1OK) Code() int {
	return 200
}

func (o *Get1OK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/formats/{format}/upload-specs][%d] get1OK %s", 200, payload)
}

func (o *Get1OK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/formats/{format}/upload-specs][%d] get1OK %s", 200, payload)
}

func (o *Get1OK) GetPayload() *models.UploadDefinitionXO {
	return o.Payload
}

func (o *Get1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UploadDefinitionXO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
