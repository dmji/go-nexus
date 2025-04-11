// Code generated by go-swagger; DO NOT EDIT.

package routing_rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"nexus/models"
)

// GetRoutingRuleReader is a Reader for the GetRoutingRule structure.
type GetRoutingRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoutingRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRoutingRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetRoutingRuleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRoutingRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/routing-rules/{name}] getRoutingRule", response, response.Code())
	}
}

// NewGetRoutingRuleOK creates a GetRoutingRuleOK with default headers values
func NewGetRoutingRuleOK() *GetRoutingRuleOK {
	return &GetRoutingRuleOK{}
}

/*
GetRoutingRuleOK describes a response with status code 200, with default header values.

successful operation
*/
type GetRoutingRuleOK struct {
	Payload *models.RoutingRuleXO
}

// IsSuccess returns true when this get routing rule o k response has a 2xx status code
func (o *GetRoutingRuleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get routing rule o k response has a 3xx status code
func (o *GetRoutingRuleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing rule o k response has a 4xx status code
func (o *GetRoutingRuleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing rule o k response has a 5xx status code
func (o *GetRoutingRuleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing rule o k response a status code equal to that given
func (o *GetRoutingRuleOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get routing rule o k response
func (o *GetRoutingRuleOK) Code() int {
	return 200
}

func (o *GetRoutingRuleOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/routing-rules/{name}][%d] getRoutingRuleOK %s", 200, payload)
}

func (o *GetRoutingRuleOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/routing-rules/{name}][%d] getRoutingRuleOK %s", 200, payload)
}

func (o *GetRoutingRuleOK) GetPayload() *models.RoutingRuleXO {
	return o.Payload
}

func (o *GetRoutingRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RoutingRuleXO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingRuleForbidden creates a GetRoutingRuleForbidden with default headers values
func NewGetRoutingRuleForbidden() *GetRoutingRuleForbidden {
	return &GetRoutingRuleForbidden{}
}

/*
GetRoutingRuleForbidden describes a response with status code 403, with default header values.

Insufficient permissions to read routing rules
*/
type GetRoutingRuleForbidden struct {
}

// IsSuccess returns true when this get routing rule forbidden response has a 2xx status code
func (o *GetRoutingRuleForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing rule forbidden response has a 3xx status code
func (o *GetRoutingRuleForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing rule forbidden response has a 4xx status code
func (o *GetRoutingRuleForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing rule forbidden response has a 5xx status code
func (o *GetRoutingRuleForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing rule forbidden response a status code equal to that given
func (o *GetRoutingRuleForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get routing rule forbidden response
func (o *GetRoutingRuleForbidden) Code() int {
	return 403
}

func (o *GetRoutingRuleForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/routing-rules/{name}][%d] getRoutingRuleForbidden", 403)
}

func (o *GetRoutingRuleForbidden) String() string {
	return fmt.Sprintf("[GET /v1/routing-rules/{name}][%d] getRoutingRuleForbidden", 403)
}

func (o *GetRoutingRuleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRoutingRuleNotFound creates a GetRoutingRuleNotFound with default headers values
func NewGetRoutingRuleNotFound() *GetRoutingRuleNotFound {
	return &GetRoutingRuleNotFound{}
}

/*
GetRoutingRuleNotFound describes a response with status code 404, with default header values.

Routing rule not found
*/
type GetRoutingRuleNotFound struct {
}

// IsSuccess returns true when this get routing rule not found response has a 2xx status code
func (o *GetRoutingRuleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing rule not found response has a 3xx status code
func (o *GetRoutingRuleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing rule not found response has a 4xx status code
func (o *GetRoutingRuleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing rule not found response has a 5xx status code
func (o *GetRoutingRuleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing rule not found response a status code equal to that given
func (o *GetRoutingRuleNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get routing rule not found response
func (o *GetRoutingRuleNotFound) Code() int {
	return 404
}

func (o *GetRoutingRuleNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/routing-rules/{name}][%d] getRoutingRuleNotFound", 404)
}

func (o *GetRoutingRuleNotFound) String() string {
	return fmt.Sprintf("[GET /v1/routing-rules/{name}][%d] getRoutingRuleNotFound", 404)
}

func (o *GetRoutingRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
