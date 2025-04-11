// Code generated by go-swagger; DO NOT EDIT.

package components

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new components API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new components API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new components API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for components API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// This client is generated with a few options you might find useful for your swagger spec.
//
// Feel free to add you own set of options.

// WithContentType allows the client to force the Content-Type header
// to negotiate a specific Consumer from the server.
//
// You may use this option to set arbitrary extensions to your MIME media type.
func WithContentType(mime string) ClientOption {
	return func(r *runtime.ClientOperation) {
		r.ConsumesMediaTypes = []string{mime}
	}
}

// WithContentTypeApplicationJSON sets the Content-Type header to "application/json".
func WithContentTypeApplicationJSON(r *runtime.ClientOperation) {
	r.ConsumesMediaTypes = []string{"application/json"}
}

// WithContentTypeMultipartFormData sets the Content-Type header to "multipart/form-data".
func WithContentTypeMultipartFormData(r *runtime.ClientOperation) {
	r.ConsumesMediaTypes = []string{"multipart/form-data"}
}

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteComponent(params *DeleteComponentParams, opts ...ClientOption) (*DeleteComponentNoContent, error)

	GetComponentByID(params *GetComponentByIDParams, opts ...ClientOption) (*GetComponentByIDOK, error)

	GetComponents(params *GetComponentsParams, opts ...ClientOption) (*GetComponentsOK, error)

	UploadComponent(params *UploadComponentParams, opts ...ClientOption) error

	SetTransport(transport runtime.ClientTransport)
}

/*
DeleteComponent deletes a single component
*/
func (a *Client) DeleteComponent(params *DeleteComponentParams, opts ...ClientOption) (*DeleteComponentNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteComponentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteComponent",
		Method:             "DELETE",
		PathPattern:        "/v1/components/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteComponentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteComponentNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteComponent: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetComponentByID gets a single component
*/
func (a *Client) GetComponentByID(params *GetComponentByIDParams, opts ...ClientOption) (*GetComponentByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetComponentByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getComponentById",
		Method:             "GET",
		PathPattern:        "/v1/components/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetComponentByIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetComponentByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getComponentById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetComponents lists components
*/
func (a *Client) GetComponents(params *GetComponentsParams, opts ...ClientOption) (*GetComponentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetComponentsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getComponents",
		Method:             "GET",
		PathPattern:        "/v1/components",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetComponentsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetComponentsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getComponents: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UploadComponent uploads a single component
*/
func (a *Client) UploadComponent(params *UploadComponentParams, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUploadComponentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "uploadComponent",
		Method:             "POST",
		PathPattern:        "/v1/components",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UploadComponentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
