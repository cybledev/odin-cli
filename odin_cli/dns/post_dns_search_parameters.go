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

// NewPostDNSSearchParams creates a new PostDNSSearchParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostDNSSearchParams() *PostDNSSearchParams {
	return &PostDNSSearchParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostDNSSearchParamsWithTimeout creates a new PostDNSSearchParams object
// with the ability to set a timeout on a request.
func NewPostDNSSearchParamsWithTimeout(timeout time.Duration) *PostDNSSearchParams {
	return &PostDNSSearchParams{
		timeout: timeout,
	}
}

// NewPostDNSSearchParamsWithContext creates a new PostDNSSearchParams object
// with the ability to set a context for a request.
func NewPostDNSSearchParamsWithContext(ctx context.Context) *PostDNSSearchParams {
	return &PostDNSSearchParams{
		Context: ctx,
	}
}

// NewPostDNSSearchParamsWithHTTPClient creates a new PostDNSSearchParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostDNSSearchParamsWithHTTPClient(client *http.Client) *PostDNSSearchParams {
	return &PostDNSSearchParams{
		HTTPClient: client,
	}
}

/*
PostDNSSearchParams contains all the parameters to send to the API endpoint

	for the post DNS search operation.

	Typically these are written to a http.Request.
*/
type PostDNSSearchParams struct {

	/* Query.

	   Query
	*/
	Query *models.DNSDNSRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post DNS search params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDNSSearchParams) WithDefaults() *PostDNSSearchParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post DNS search params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDNSSearchParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post DNS search params
func (o *PostDNSSearchParams) WithTimeout(timeout time.Duration) *PostDNSSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post DNS search params
func (o *PostDNSSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post DNS search params
func (o *PostDNSSearchParams) WithContext(ctx context.Context) *PostDNSSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post DNS search params
func (o *PostDNSSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post DNS search params
func (o *PostDNSSearchParams) WithHTTPClient(client *http.Client) *PostDNSSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post DNS search params
func (o *PostDNSSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithQuery adds the query to the post DNS search params
func (o *PostDNSSearchParams) WithQuery(query *models.DNSDNSRequest) *PostDNSSearchParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the post DNS search params
func (o *PostDNSSearchParams) SetQuery(query *models.DNSDNSRequest) {
	o.Query = query
}

// WriteToRequest writes these params to a swagger request
func (o *PostDNSSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
