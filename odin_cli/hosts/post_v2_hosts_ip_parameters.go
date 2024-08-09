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

// NewPostV2HostsIPParams creates a new PostV2HostsIPParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV2HostsIPParams() *PostV2HostsIPParams {
	return &PostV2HostsIPParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV2HostsIPParamsWithTimeout creates a new PostV2HostsIPParams object
// with the ability to set a timeout on a request.
func NewPostV2HostsIPParamsWithTimeout(timeout time.Duration) *PostV2HostsIPParams {
	return &PostV2HostsIPParams{
		timeout: timeout,
	}
}

// NewPostV2HostsIPParamsWithContext creates a new PostV2HostsIPParams object
// with the ability to set a context for a request.
func NewPostV2HostsIPParamsWithContext(ctx context.Context) *PostV2HostsIPParams {
	return &PostV2HostsIPParams{
		Context: ctx,
	}
}

// NewPostV2HostsIPParamsWithHTTPClient creates a new PostV2HostsIPParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV2HostsIPParamsWithHTTPClient(client *http.Client) *PostV2HostsIPParams {
	return &PostV2HostsIPParams{
		HTTPClient: client,
	}
}

/*
PostV2HostsIPParams contains all the parameters to send to the API endpoint

	for the post v2 hosts IP operation.

	Typically these are written to a http.Request.
*/
type PostV2HostsIPParams struct {

	/* IP.

	   get the ip
	*/
	IP string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v2 hosts IP params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV2HostsIPParams) WithDefaults() *PostV2HostsIPParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v2 hosts IP params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV2HostsIPParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v2 hosts IP params
func (o *PostV2HostsIPParams) WithTimeout(timeout time.Duration) *PostV2HostsIPParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v2 hosts IP params
func (o *PostV2HostsIPParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v2 hosts IP params
func (o *PostV2HostsIPParams) WithContext(ctx context.Context) *PostV2HostsIPParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v2 hosts IP params
func (o *PostV2HostsIPParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v2 hosts IP params
func (o *PostV2HostsIPParams) WithHTTPClient(client *http.Client) *PostV2HostsIPParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v2 hosts IP params
func (o *PostV2HostsIPParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIP adds the ip to the post v2 hosts IP params
func (o *PostV2HostsIPParams) WithIP(ip string) *PostV2HostsIPParams {
	o.SetIP(ip)
	return o
}

// SetIP adds the ip to the post v2 hosts IP params
func (o *PostV2HostsIPParams) SetIP(ip string) {
	o.IP = ip
}

// WriteToRequest writes these params to a swagger request
func (o *PostV2HostsIPParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
