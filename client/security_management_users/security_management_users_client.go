// Code generated by go-swagger; DO NOT EDIT.

package security_management_users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new security management users API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new security management users API client with basic auth credentials.
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

// New creates a new security management users API client with a bearer token for authentication.
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
Client for security management users API
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

// WithContentTypeTextPlain sets the Content-Type header to "text/plain".
func WithContentTypeTextPlain(r *runtime.ClientOperation) {
	r.ConsumesMediaTypes = []string{"text/plain"}
}

// ClientService is the interface for Client methods
type ClientService interface {
	ChangePassword(params *ChangePasswordParams, opts ...ClientOption) error

	CreateUser(params *CreateUserParams, opts ...ClientOption) (*CreateUserOK, error)

	DeleteUser(params *DeleteUserParams, opts ...ClientOption) error

	GetUsers(params *GetUsersParams, opts ...ClientOption) (*GetUsersOK, error)

	UpdateUser(params *UpdateUserParams, opts ...ClientOption) error

	SetTransport(transport runtime.ClientTransport)
}

/*
ChangePassword changes a user s password
*/
func (a *Client) ChangePassword(params *ChangePasswordParams, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangePasswordParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "changePassword",
		Method:             "PUT",
		PathPattern:        "/v1/security/users/{userId}/change-password",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"text/plain"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ChangePasswordReader{formats: a.formats},
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

/*
CreateUser creates a new user in the default source
*/
func (a *Client) CreateUser(params *CreateUserParams, opts ...ClientOption) (*CreateUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateUserParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createUser",
		Method:             "POST",
		PathPattern:        "/v1/security/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateUserReader{formats: a.formats},
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
	success, ok := result.(*CreateUserOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createUser: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteUser deletes a user
*/
func (a *Client) DeleteUser(params *DeleteUserParams, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteUserParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteUser",
		Method:             "DELETE",
		PathPattern:        "/v1/security/users/{userId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteUserReader{formats: a.formats},
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

/*
GetUsers retrieves a list of users
*/
func (a *Client) GetUsers(params *GetUsersParams, opts ...ClientOption) (*GetUsersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUsersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getUsers",
		Method:             "GET",
		PathPattern:        "/v1/security/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetUsersReader{formats: a.formats},
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
	success, ok := result.(*GetUsersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getUsers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateUser updates an existing user
*/
func (a *Client) UpdateUser(params *UpdateUserParams, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateUserParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateUser",
		Method:             "PUT",
		PathPattern:        "/v1/security/users/{userId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateUserReader{formats: a.formats},
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
