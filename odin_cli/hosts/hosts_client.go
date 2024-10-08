// Code generated by go-swagger; DO NOT EDIT.

package hosts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new hosts API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new hosts API client with basic auth credentials.
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

// New creates a new hosts API client with a bearer token for authentication.
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
Client for hosts API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetHostsCveIP(params *GetHostsCveIPParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetHostsCveIPOK, error)

	GetHostsExploitsIP(params *GetHostsExploitsIPParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetHostsExploitsIPOK, error)

	GetHostsIP(params *GetHostsIPParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetHostsIPOK, error)

	PostHostsCount(params *PostHostsCountParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostHostsCountOK, error)

	PostHostsFilteredSearch(params *PostHostsFilteredSearchParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostHostsFilteredSearchOK, error)

	PostHostsSearch(params *PostHostsSearchParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostHostsSearchOK, error)

	PostHostsSummary(params *PostHostsSummaryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostHostsSummaryOK, error)

	PostV2HostsCount(params *PostV2HostsCountParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV2HostsCountOK, error)

	PostV2HostsIP(params *PostV2HostsIPParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV2HostsIPOK, error)

	PostV2HostsSearch(params *PostV2HostsSearchParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV2HostsSearchOK, error)

	PostV2HostsSummary(params *PostV2HostsSummaryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV2HostsSummaryOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetHostsCveIP fetches the ip cve details

Returns the complete ip cve details
*/
func (a *Client) GetHostsCveIP(params *GetHostsCveIPParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetHostsCveIPOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHostsCveIPParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "hosts-cve-ip",
		Method:             "GET",
		PathPattern:        "/hosts/cve/{ip}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetHostsCveIPReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*GetHostsCveIPOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetHostsCveIP: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetHostsExploitsIP fetches the ip cve details

Returns the complete ip cve-exploits details
*/
func (a *Client) GetHostsExploitsIP(params *GetHostsExploitsIPParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetHostsExploitsIPOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHostsExploitsIPParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "hosts-exploits-ip",
		Method:             "GET",
		PathPattern:        "/hosts/exploits/{ip}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetHostsExploitsIPReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*GetHostsExploitsIPOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetHostsExploitsIP: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetHostsIP fetches the latest ip details

Returns the complete ip details
*/
func (a *Client) GetHostsIP(params *GetHostsIPParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetHostsIPOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHostsIPParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-hosts-ip",
		Method:             "GET",
		PathPattern:        "/hosts/{ip}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetHostsIPReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*GetHostsIPOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetHostsIP: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostHostsCount fetches the record count

Returns the total no of records based on query
*/
func (a *Client) PostHostsCount(params *PostHostsCountParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostHostsCountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostHostsCountParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "count",
		Method:             "POST",
		PathPattern:        "/hosts/count",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostHostsCountReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*PostHostsCountOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostHostsCount: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostHostsFilteredSearch fetches the record based on query

Returns the record based on query and blank query will return all records. It uses es searchafter for the pagination.
*/
func (a *Client) PostHostsFilteredSearch(params *PostHostsFilteredSearchParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostHostsFilteredSearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostHostsFilteredSearchParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "filtered-search",
		Method:             "POST",
		PathPattern:        "/hosts/filtered/search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostHostsFilteredSearchReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*PostHostsFilteredSearchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostHostsFilteredSearch: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostHostsSearch fetches the record based on query

Returns the record based on query and blank query will return all records. It uses es searchafter for the pagination.
*/
func (a *Client) PostHostsSearch(params *PostHostsSearchParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostHostsSearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostHostsSearchParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "search",
		Method:             "POST",
		PathPattern:        "/hosts/search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostHostsSearchReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*PostHostsSearchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostHostsSearch: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostHostsSummary creates the summary of the field based on query

Returns the summary of the field based on query
*/
func (a *Client) PostHostsSummary(params *PostHostsSummaryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostHostsSummaryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostHostsSummaryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "summary",
		Method:             "POST",
		PathPattern:        "/hosts/summary",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostHostsSummaryReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*PostHostsSummaryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostHostsSummary: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostV2HostsCount fetches the record count

Returns the total no of records based on query
*/
func (a *Client) PostV2HostsCount(params *PostV2HostsCountParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV2HostsCountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostV2HostsCountParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostV2HostsCount",
		Method:             "POST",
		PathPattern:        "/v2/hosts/count",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostV2HostsCountReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*PostV2HostsCountOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostV2HostsCount: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostV2HostsIP fetches the latest ip details

Returns the complete ip details
*/
func (a *Client) PostV2HostsIP(params *PostV2HostsIPParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV2HostsIPOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostV2HostsIPParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostV2HostsIP",
		Method:             "POST",
		PathPattern:        "/v2/hosts/{ip}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostV2HostsIPReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*PostV2HostsIPOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostV2HostsIP: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostV2HostsSearch fetches the record based on query

Returns the record based on query and blank query will return all records. It uses es searchafter for the pagination.
*/
func (a *Client) PostV2HostsSearch(params *PostV2HostsSearchParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV2HostsSearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostV2HostsSearchParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostV2HostsSearch",
		Method:             "POST",
		PathPattern:        "/hosts/search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostV2HostsSearchReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*PostV2HostsSearchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostV2HostsSearch: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostV2HostsSummary creates the summary of the field based on query

Returns the summary of the field based on query
*/
func (a *Client) PostV2HostsSummary(params *PostV2HostsSummaryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV2HostsSummaryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostV2HostsSummaryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostV2HostsSummary",
		Method:             "POST",
		PathPattern:        "/hosts/summary",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostV2HostsSummaryReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*PostV2HostsSummaryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostV2HostsSummary: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
