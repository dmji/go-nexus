// Code generated by go-swagger; DO NOT EDIT.

package content_selectors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new content selectors API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new content selectors API client with basic auth credentials.
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

// New creates a new content selectors API client with a bearer token for authentication.
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
Client for content selectors API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteSecurityContentSelectorsByName(params *DeleteSecurityContentSelectorsByNameParams, opts ...ClientOption) (*DeleteSecurityContentSelectorsByNameNoContent, error)

	GetSecurityContentSelectors(params *GetSecurityContentSelectorsParams, opts ...ClientOption) (*GetSecurityContentSelectorsOK, error)

	GetSecurityContentSelectorsByName(params *GetSecurityContentSelectorsByNameParams, opts ...ClientOption) (*GetSecurityContentSelectorsByNameOK, error)

	PostSecurityContentSelectors(params *PostSecurityContentSelectorsParams, opts ...ClientOption) (*PostSecurityContentSelectorsNoContent, error)

	PutSecurityContentSelectorsByName(params *PutSecurityContentSelectorsByNameParams, opts ...ClientOption) (*PutSecurityContentSelectorsByNameNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DeleteSecurityContentSelectorsByName deletes a content selector
*/
func (a *Client) DeleteSecurityContentSelectorsByName(params *DeleteSecurityContentSelectorsByNameParams, opts ...ClientOption) (*DeleteSecurityContentSelectorsByNameNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSecurityContentSelectorsByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteSecurityContentSelectorsByName",
		Method:             "DELETE",
		PathPattern:        "/v1/security/content-selectors/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteSecurityContentSelectorsByNameReader{formats: a.formats},
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
	success, ok := result.(*DeleteSecurityContentSelectorsByNameNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteSecurityContentSelectorsByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetSecurityContentSelectors lists content selectors
*/
func (a *Client) GetSecurityContentSelectors(params *GetSecurityContentSelectorsParams, opts ...ClientOption) (*GetSecurityContentSelectorsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSecurityContentSelectorsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetSecurityContentSelectors",
		Method:             "GET",
		PathPattern:        "/v1/security/content-selectors",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSecurityContentSelectorsReader{formats: a.formats},
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
	success, ok := result.(*GetSecurityContentSelectorsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetSecurityContentSelectors: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetSecurityContentSelectorsByName gets a content selector by name
*/
func (a *Client) GetSecurityContentSelectorsByName(params *GetSecurityContentSelectorsByNameParams, opts ...ClientOption) (*GetSecurityContentSelectorsByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSecurityContentSelectorsByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetSecurityContentSelectorsByName",
		Method:             "GET",
		PathPattern:        "/v1/security/content-selectors/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSecurityContentSelectorsByNameReader{formats: a.formats},
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
	success, ok := result.(*GetSecurityContentSelectorsByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetSecurityContentSelectorsByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostSecurityContentSelectors creates a new content selector
*/
func (a *Client) PostSecurityContentSelectors(params *PostSecurityContentSelectorsParams, opts ...ClientOption) (*PostSecurityContentSelectorsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostSecurityContentSelectorsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostSecurityContentSelectors",
		Method:             "POST",
		PathPattern:        "/v1/security/content-selectors",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostSecurityContentSelectorsReader{formats: a.formats},
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
	success, ok := result.(*PostSecurityContentSelectorsNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostSecurityContentSelectors: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PutSecurityContentSelectorsByName updates a content selector
*/
func (a *Client) PutSecurityContentSelectorsByName(params *PutSecurityContentSelectorsByNameParams, opts ...ClientOption) (*PutSecurityContentSelectorsByNameNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutSecurityContentSelectorsByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutSecurityContentSelectorsByName",
		Method:             "PUT",
		PathPattern:        "/v1/security/content-selectors/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutSecurityContentSelectorsByNameReader{formats: a.formats},
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
	success, ok := result.(*PutSecurityContentSelectorsByNameNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PutSecurityContentSelectorsByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
