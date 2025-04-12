// Code generated by go-swagger; DO NOT EDIT.

package blob_store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new blob store API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new blob store API client with basic auth credentials.
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

// New creates a new blob store API client with a bearer token for authentication.
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
Client for blob store API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteBlobstoresByName(params *DeleteBlobstoresByNameParams, opts ...ClientOption) error

	GetBlobstores(params *GetBlobstoresParams, opts ...ClientOption) (*GetBlobstoresOK, error)

	GetBlobstoresAzureByName(params *GetBlobstoresAzureByNameParams, opts ...ClientOption) (*GetBlobstoresAzureByNameOK, error)

	GetBlobstoresByNameQuotaStatus(params *GetBlobstoresByNameQuotaStatusParams, opts ...ClientOption) (*GetBlobstoresByNameQuotaStatusOK, error)

	GetBlobstoresFileByName(params *GetBlobstoresFileByNameParams, opts ...ClientOption) (*GetBlobstoresFileByNameOK, error)

	GetBlobstoresGoogleByName(params *GetBlobstoresGoogleByNameParams, opts ...ClientOption) (*GetBlobstoresGoogleByNameOK, error)

	GetBlobstoresGoogleRegionsByProjectID(params *GetBlobstoresGoogleRegionsByProjectIDParams, opts ...ClientOption) (*GetBlobstoresGoogleRegionsByProjectIDOK, error)

	GetBlobstoresGroupByName(params *GetBlobstoresGroupByNameParams, opts ...ClientOption) (*GetBlobstoresGroupByNameOK, error)

	GetBlobstoresS3ByName(params *GetBlobstoresS3ByNameParams, opts ...ClientOption) (*GetBlobstoresS3ByNameOK, error)

	PostBlobstoresAzure(params *PostBlobstoresAzureParams, opts ...ClientOption) (*PostBlobstoresAzureCreated, error)

	PostBlobstoresFile(params *PostBlobstoresFileParams, opts ...ClientOption) (*PostBlobstoresFileNoContent, error)

	PostBlobstoresGoogle(params *PostBlobstoresGoogleParams, opts ...ClientOption) (*PostBlobstoresGoogleCreated, error)

	PostBlobstoresGroup(params *PostBlobstoresGroupParams, opts ...ClientOption) (*PostBlobstoresGroupNoContent, error)

	PostBlobstoresGroupConvertByNameByNewNameForOriginal(params *PostBlobstoresGroupConvertByNameByNewNameForOriginalParams, opts ...ClientOption) (*PostBlobstoresGroupConvertByNameByNewNameForOriginalOK, error)

	PostBlobstoresS3(params *PostBlobstoresS3Params, opts ...ClientOption) (*PostBlobstoresS3Created, error)

	PutBlobstoresAzureByName(params *PutBlobstoresAzureByNameParams, opts ...ClientOption) (*PutBlobstoresAzureByNameNoContent, error)

	PutBlobstoresFileByName(params *PutBlobstoresFileByNameParams, opts ...ClientOption) (*PutBlobstoresFileByNameNoContent, error)

	PutBlobstoresGoogleByName(params *PutBlobstoresGoogleByNameParams, opts ...ClientOption) (*PutBlobstoresGoogleByNameNoContent, error)

	PutBlobstoresGroupByName(params *PutBlobstoresGroupByNameParams, opts ...ClientOption) (*PutBlobstoresGroupByNameNoContent, error)

	PutBlobstoresS3ByName(params *PutBlobstoresS3ByNameParams, opts ...ClientOption) (*PutBlobstoresS3ByNameNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DeleteBlobstoresByName deletes a blob store by name
*/
func (a *Client) DeleteBlobstoresByName(params *DeleteBlobstoresByNameParams, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteBlobstoresByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteBlobstoresByName",
		Method:             "DELETE",
		PathPattern:        "/v1/blobstores/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteBlobstoresByNameReader{formats: a.formats},
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
GetBlobstores lists the blob stores
*/
func (a *Client) GetBlobstores(params *GetBlobstoresParams, opts ...ClientOption) (*GetBlobstoresOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBlobstoresParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetBlobstores",
		Method:             "GET",
		PathPattern:        "/v1/blobstores",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetBlobstoresReader{formats: a.formats},
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
	success, ok := result.(*GetBlobstoresOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetBlobstores: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetBlobstoresAzureByName gets an azure blob store configuration by name
*/
func (a *Client) GetBlobstoresAzureByName(params *GetBlobstoresAzureByNameParams, opts ...ClientOption) (*GetBlobstoresAzureByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBlobstoresAzureByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetBlobstoresAzureByName",
		Method:             "GET",
		PathPattern:        "/v1/blobstores/azure/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetBlobstoresAzureByNameReader{formats: a.formats},
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
	success, ok := result.(*GetBlobstoresAzureByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetBlobstoresAzureByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetBlobstoresByNameQuotaStatus gets quota status for a given blob store
*/
func (a *Client) GetBlobstoresByNameQuotaStatus(params *GetBlobstoresByNameQuotaStatusParams, opts ...ClientOption) (*GetBlobstoresByNameQuotaStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBlobstoresByNameQuotaStatusParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetBlobstoresByNameQuotaStatus",
		Method:             "GET",
		PathPattern:        "/v1/blobstores/{name}/quota-status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetBlobstoresByNameQuotaStatusReader{formats: a.formats},
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
	success, ok := result.(*GetBlobstoresByNameQuotaStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetBlobstoresByNameQuotaStatus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetBlobstoresFileByName gets a file blob store configuration by name
*/
func (a *Client) GetBlobstoresFileByName(params *GetBlobstoresFileByNameParams, opts ...ClientOption) (*GetBlobstoresFileByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBlobstoresFileByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetBlobstoresFileByName",
		Method:             "GET",
		PathPattern:        "/v1/blobstores/file/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetBlobstoresFileByNameReader{formats: a.formats},
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
	success, ok := result.(*GetBlobstoresFileByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetBlobstoresFileByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetBlobstoresGoogleByName gets the configuration for a google cloud blob store
*/
func (a *Client) GetBlobstoresGoogleByName(params *GetBlobstoresGoogleByNameParams, opts ...ClientOption) (*GetBlobstoresGoogleByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBlobstoresGoogleByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetBlobstoresGoogleByName",
		Method:             "GET",
		PathPattern:        "/v1/blobstores/google/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetBlobstoresGoogleByNameReader{formats: a.formats},
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
	success, ok := result.(*GetBlobstoresGoogleByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetBlobstoresGoogleByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetBlobstoresGoogleRegionsByProjectID gets the project regions by project s id
*/
func (a *Client) GetBlobstoresGoogleRegionsByProjectID(params *GetBlobstoresGoogleRegionsByProjectIDParams, opts ...ClientOption) (*GetBlobstoresGoogleRegionsByProjectIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBlobstoresGoogleRegionsByProjectIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetBlobstoresGoogleRegionsByProjectId",
		Method:             "GET",
		PathPattern:        "/v1/blobstores/google/regions/{projectId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetBlobstoresGoogleRegionsByProjectIDReader{formats: a.formats},
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
	success, ok := result.(*GetBlobstoresGoogleRegionsByProjectIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetBlobstoresGoogleRegionsByProjectId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetBlobstoresGroupByName gets a group blob store configuration by name
*/
func (a *Client) GetBlobstoresGroupByName(params *GetBlobstoresGroupByNameParams, opts ...ClientOption) (*GetBlobstoresGroupByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBlobstoresGroupByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetBlobstoresGroupByName",
		Method:             "GET",
		PathPattern:        "/v1/blobstores/group/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetBlobstoresGroupByNameReader{formats: a.formats},
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
	success, ok := result.(*GetBlobstoresGroupByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetBlobstoresGroupByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetBlobstoresS3ByName gets a s3 blob store configuration by name
*/
func (a *Client) GetBlobstoresS3ByName(params *GetBlobstoresS3ByNameParams, opts ...ClientOption) (*GetBlobstoresS3ByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBlobstoresS3ByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetBlobstoresS3ByName",
		Method:             "GET",
		PathPattern:        "/v1/blobstores/s3/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetBlobstoresS3ByNameReader{formats: a.formats},
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
	success, ok := result.(*GetBlobstoresS3ByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetBlobstoresS3ByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostBlobstoresAzure creates an azure blob store
*/
func (a *Client) PostBlobstoresAzure(params *PostBlobstoresAzureParams, opts ...ClientOption) (*PostBlobstoresAzureCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostBlobstoresAzureParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostBlobstoresAzure",
		Method:             "POST",
		PathPattern:        "/v1/blobstores/azure",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostBlobstoresAzureReader{formats: a.formats},
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
	success, ok := result.(*PostBlobstoresAzureCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostBlobstoresAzure: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostBlobstoresFile creates a file blob store
*/
func (a *Client) PostBlobstoresFile(params *PostBlobstoresFileParams, opts ...ClientOption) (*PostBlobstoresFileNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostBlobstoresFileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostBlobstoresFile",
		Method:             "POST",
		PathPattern:        "/v1/blobstores/file",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostBlobstoresFileReader{formats: a.formats},
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
	success, ok := result.(*PostBlobstoresFileNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostBlobstoresFile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostBlobstoresGoogle creates a google cloud blob store
*/
func (a *Client) PostBlobstoresGoogle(params *PostBlobstoresGoogleParams, opts ...ClientOption) (*PostBlobstoresGoogleCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostBlobstoresGoogleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostBlobstoresGoogle",
		Method:             "POST",
		PathPattern:        "/v1/blobstores/google",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostBlobstoresGoogleReader{formats: a.formats},
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
	success, ok := result.(*PostBlobstoresGoogleCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostBlobstoresGoogle: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostBlobstoresGroup creates a group blob store
*/
func (a *Client) PostBlobstoresGroup(params *PostBlobstoresGroupParams, opts ...ClientOption) (*PostBlobstoresGroupNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostBlobstoresGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostBlobstoresGroup",
		Method:             "POST",
		PathPattern:        "/v1/blobstores/group",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostBlobstoresGroupReader{formats: a.formats},
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
	success, ok := result.(*PostBlobstoresGroupNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostBlobstoresGroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostBlobstoresGroupConvertByNameByNewNameForOriginal converts a blob store to a group blob store
*/
func (a *Client) PostBlobstoresGroupConvertByNameByNewNameForOriginal(params *PostBlobstoresGroupConvertByNameByNewNameForOriginalParams, opts ...ClientOption) (*PostBlobstoresGroupConvertByNameByNewNameForOriginalOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostBlobstoresGroupConvertByNameByNewNameForOriginalParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostBlobstoresGroupConvertByNameByNewNameForOriginal",
		Method:             "POST",
		PathPattern:        "/v1/blobstores/group/convert/{name}/{newNameForOriginal}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostBlobstoresGroupConvertByNameByNewNameForOriginalReader{formats: a.formats},
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
	success, ok := result.(*PostBlobstoresGroupConvertByNameByNewNameForOriginalOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostBlobstoresGroupConvertByNameByNewNameForOriginal: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostBlobstoresS3 creates an s3 blob store
*/
func (a *Client) PostBlobstoresS3(params *PostBlobstoresS3Params, opts ...ClientOption) (*PostBlobstoresS3Created, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostBlobstoresS3Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostBlobstoresS3",
		Method:             "POST",
		PathPattern:        "/v1/blobstores/s3",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostBlobstoresS3Reader{formats: a.formats},
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
	success, ok := result.(*PostBlobstoresS3Created)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostBlobstoresS3: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PutBlobstoresAzureByName updates an azure blob store configuration by name
*/
func (a *Client) PutBlobstoresAzureByName(params *PutBlobstoresAzureByNameParams, opts ...ClientOption) (*PutBlobstoresAzureByNameNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutBlobstoresAzureByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutBlobstoresAzureByName",
		Method:             "PUT",
		PathPattern:        "/v1/blobstores/azure/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutBlobstoresAzureByNameReader{formats: a.formats},
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
	success, ok := result.(*PutBlobstoresAzureByNameNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PutBlobstoresAzureByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PutBlobstoresFileByName updates a file blob store configuration by name
*/
func (a *Client) PutBlobstoresFileByName(params *PutBlobstoresFileByNameParams, opts ...ClientOption) (*PutBlobstoresFileByNameNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutBlobstoresFileByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutBlobstoresFileByName",
		Method:             "PUT",
		PathPattern:        "/v1/blobstores/file/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutBlobstoresFileByNameReader{formats: a.formats},
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
	success, ok := result.(*PutBlobstoresFileByNameNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PutBlobstoresFileByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PutBlobstoresGoogleByName updates a google cloud blob store
*/
func (a *Client) PutBlobstoresGoogleByName(params *PutBlobstoresGoogleByNameParams, opts ...ClientOption) (*PutBlobstoresGoogleByNameNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutBlobstoresGoogleByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutBlobstoresGoogleByName",
		Method:             "PUT",
		PathPattern:        "/v1/blobstores/google/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutBlobstoresGoogleByNameReader{formats: a.formats},
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
	success, ok := result.(*PutBlobstoresGoogleByNameNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PutBlobstoresGoogleByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PutBlobstoresGroupByName updates a group blob store configuration by name
*/
func (a *Client) PutBlobstoresGroupByName(params *PutBlobstoresGroupByNameParams, opts ...ClientOption) (*PutBlobstoresGroupByNameNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutBlobstoresGroupByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutBlobstoresGroupByName",
		Method:             "PUT",
		PathPattern:        "/v1/blobstores/group/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutBlobstoresGroupByNameReader{formats: a.formats},
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
	success, ok := result.(*PutBlobstoresGroupByNameNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PutBlobstoresGroupByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PutBlobstoresS3ByName updates an s3 blob store configuration by name
*/
func (a *Client) PutBlobstoresS3ByName(params *PutBlobstoresS3ByNameParams, opts ...ClientOption) (*PutBlobstoresS3ByNameNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutBlobstoresS3ByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutBlobstoresS3ByName",
		Method:             "PUT",
		PathPattern:        "/v1/blobstores/s3/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutBlobstoresS3ByNameReader{formats: a.formats},
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
	success, ok := result.(*PutBlobstoresS3ByNameNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PutBlobstoresS3ByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
