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

// NewPostDomainSubdomainCountParams creates a new PostDomainSubdomainCountParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostDomainSubdomainCountParams() *PostDomainSubdomainCountParams {
	return &PostDomainSubdomainCountParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostDomainSubdomainCountParamsWithTimeout creates a new PostDomainSubdomainCountParams object
// with the ability to set a timeout on a request.
func NewPostDomainSubdomainCountParamsWithTimeout(timeout time.Duration) *PostDomainSubdomainCountParams {
	return &PostDomainSubdomainCountParams{
		timeout: timeout,
	}
}

// NewPostDomainSubdomainCountParamsWithContext creates a new PostDomainSubdomainCountParams object
// with the ability to set a context for a request.
func NewPostDomainSubdomainCountParamsWithContext(ctx context.Context) *PostDomainSubdomainCountParams {
	return &PostDomainSubdomainCountParams{
		Context: ctx,
	}
}

// NewPostDomainSubdomainCountParamsWithHTTPClient creates a new PostDomainSubdomainCountParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostDomainSubdomainCountParamsWithHTTPClient(client *http.Client) *PostDomainSubdomainCountParams {
	return &PostDomainSubdomainCountParams{
		HTTPClient: client,
	}
}

/*
PostDomainSubdomainCountParams contains all the parameters to send to the API endpoint

	for the post domain subdomain count operation.

	Typically these are written to a http.Request.
*/
type PostDomainSubdomainCountParams struct {

	/* Query.

	   Query
	*/
	Query *models.DNSDNSCountRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post domain subdomain count params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDomainSubdomainCountParams) WithDefaults() *PostDomainSubdomainCountParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post domain subdomain count params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDomainSubdomainCountParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post domain subdomain count params
func (o *PostDomainSubdomainCountParams) WithTimeout(timeout time.Duration) *PostDomainSubdomainCountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post domain subdomain count params
func (o *PostDomainSubdomainCountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post domain subdomain count params
func (o *PostDomainSubdomainCountParams) WithContext(ctx context.Context) *PostDomainSubdomainCountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post domain subdomain count params
func (o *PostDomainSubdomainCountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post domain subdomain count params
func (o *PostDomainSubdomainCountParams) WithHTTPClient(client *http.Client) *PostDomainSubdomainCountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post domain subdomain count params
func (o *PostDomainSubdomainCountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithQuery adds the query to the post domain subdomain count params
func (o *PostDomainSubdomainCountParams) WithQuery(query *models.DNSDNSCountRequest) *PostDomainSubdomainCountParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the post domain subdomain count params
func (o *PostDomainSubdomainCountParams) SetQuery(query *models.DNSDNSCountRequest) {
	o.Query = query
}

// WriteToRequest writes these params to a swagger request
func (o *PostDomainSubdomainCountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
