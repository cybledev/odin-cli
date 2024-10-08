// Code generated by go-swagger; DO NOT EDIT.

package hosts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetHostsIPParams creates a new GetHostsIPParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetHostsIPParams() *GetHostsIPParams {
	return &GetHostsIPParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetHostsIPParamsWithTimeout creates a new GetHostsIPParams object
// with the ability to set a timeout on a request.
func NewGetHostsIPParamsWithTimeout(timeout time.Duration) *GetHostsIPParams {
	return &GetHostsIPParams{
		timeout: timeout,
	}
}

// NewGetHostsIPParamsWithContext creates a new GetHostsIPParams object
// with the ability to set a context for a request.
func NewGetHostsIPParamsWithContext(ctx context.Context) *GetHostsIPParams {
	return &GetHostsIPParams{
		Context: ctx,
	}
}

// NewGetHostsIPParamsWithHTTPClient creates a new GetHostsIPParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetHostsIPParamsWithHTTPClient(client *http.Client) *GetHostsIPParams {
	return &GetHostsIPParams{
		HTTPClient: client,
	}
}

/*
GetHostsIPParams contains all the parameters to send to the API endpoint

	for the get hosts IP operation.

	Typically these are written to a http.Request.
*/
type GetHostsIPParams struct {

	/* IP.

	   get the ip
	*/
	IP string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get hosts IP params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetHostsIPParams) WithDefaults() *GetHostsIPParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get hosts IP params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetHostsIPParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get hosts IP params
func (o *GetHostsIPParams) WithTimeout(timeout time.Duration) *GetHostsIPParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get hosts IP params
func (o *GetHostsIPParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get hosts IP params
func (o *GetHostsIPParams) WithContext(ctx context.Context) *GetHostsIPParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get hosts IP params
func (o *GetHostsIPParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get hosts IP params
func (o *GetHostsIPParams) WithHTTPClient(client *http.Client) *GetHostsIPParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get hosts IP params
func (o *GetHostsIPParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIP adds the ip to the get hosts IP params
func (o *GetHostsIPParams) WithIP(ip string) *GetHostsIPParams {
	o.SetIP(ip)
	return o
}

// SetIP adds the ip to the get hosts IP params
func (o *GetHostsIPParams) SetIP(ip string) {
	o.IP = ip
}

// WriteToRequest writes these params to a swagger request
func (o *GetHostsIPParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ip
	if err := r.SetPathParam("ip", o.IP); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
