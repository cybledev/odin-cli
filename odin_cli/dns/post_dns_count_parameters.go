// Code generated by go-swagger; DO NOT EDIT.

package dns

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

	"github.com/cybledev/odin-cli/models"
)

// NewPostDNSCountParams creates a new PostDNSCountParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostDNSCountParams() *PostDNSCountParams {
	return &PostDNSCountParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostDNSCountParamsWithTimeout creates a new PostDNSCountParams object
// with the ability to set a timeout on a request.
func NewPostDNSCountParamsWithTimeout(timeout time.Duration) *PostDNSCountParams {
	return &PostDNSCountParams{
		timeout: timeout,
	}
}

// NewPostDNSCountParamsWithContext creates a new PostDNSCountParams object
// with the ability to set a context for a request.
func NewPostDNSCountParamsWithContext(ctx context.Context) *PostDNSCountParams {
	return &PostDNSCountParams{
		Context: ctx,
	}
}

// NewPostDNSCountParamsWithHTTPClient creates a new PostDNSCountParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostDNSCountParamsWithHTTPClient(client *http.Client) *PostDNSCountParams {
	return &PostDNSCountParams{
		HTTPClient: client,
	}
}

/*
PostDNSCountParams contains all the parameters to send to the API endpoint

	for the post DNS count operation.

	Typically these are written to a http.Request.
*/
type PostDNSCountParams struct {

	/* Query.

	   Query
	*/
	Query *models.DNSDNSCountRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post DNS count params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDNSCountParams) WithDefaults() *PostDNSCountParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post DNS count params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDNSCountParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post DNS count params
func (o *PostDNSCountParams) WithTimeout(timeout time.Duration) *PostDNSCountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post DNS count params
func (o *PostDNSCountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post DNS count params
func (o *PostDNSCountParams) WithContext(ctx context.Context) *PostDNSCountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post DNS count params
func (o *PostDNSCountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post DNS count params
func (o *PostDNSCountParams) WithHTTPClient(client *http.Client) *PostDNSCountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post DNS count params
func (o *PostDNSCountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithQuery adds the query to the post DNS count params
func (o *PostDNSCountParams) WithQuery(query *models.DNSDNSCountRequest) *PostDNSCountParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the post DNS count params
func (o *PostDNSCountParams) SetQuery(query *models.DNSDNSCountRequest) {
	o.Query = query
}

// WriteToRequest writes these params to a swagger request
func (o *PostDNSCountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Query != nil {
		if err := r.SetBodyParam(o.Query); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
