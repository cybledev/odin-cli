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

// NewPostDomainSubdomainSearchParams creates a new PostDomainSubdomainSearchParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostDomainSubdomainSearchParams() *PostDomainSubdomainSearchParams {
	return &PostDomainSubdomainSearchParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostDomainSubdomainSearchParamsWithTimeout creates a new PostDomainSubdomainSearchParams object
// with the ability to set a timeout on a request.
func NewPostDomainSubdomainSearchParamsWithTimeout(timeout time.Duration) *PostDomainSubdomainSearchParams {
	return &PostDomainSubdomainSearchParams{
		timeout: timeout,
	}
}

// NewPostDomainSubdomainSearchParamsWithContext creates a new PostDomainSubdomainSearchParams object
// with the ability to set a context for a request.
func NewPostDomainSubdomainSearchParamsWithContext(ctx context.Context) *PostDomainSubdomainSearchParams {
	return &PostDomainSubdomainSearchParams{
		Context: ctx,
	}
}

// NewPostDomainSubdomainSearchParamsWithHTTPClient creates a new PostDomainSubdomainSearchParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostDomainSubdomainSearchParamsWithHTTPClient(client *http.Client) *PostDomainSubdomainSearchParams {
	return &PostDomainSubdomainSearchParams{
		HTTPClient: client,
	}
}

/*
PostDomainSubdomainSearchParams contains all the parameters to send to the API endpoint

	for the post domain subdomain search operation.

	Typically these are written to a http.Request.
*/
type PostDomainSubdomainSearchParams struct {

	/* Query.

	   Query
	*/
	Query *models.DNSDomainRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post domain subdomain search params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDomainSubdomainSearchParams) WithDefaults() *PostDomainSubdomainSearchParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post domain subdomain search params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDomainSubdomainSearchParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post domain subdomain search params
func (o *PostDomainSubdomainSearchParams) WithTimeout(timeout time.Duration) *PostDomainSubdomainSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post domain subdomain search params
func (o *PostDomainSubdomainSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post domain subdomain search params
func (o *PostDomainSubdomainSearchParams) WithContext(ctx context.Context) *PostDomainSubdomainSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post domain subdomain search params
func (o *PostDomainSubdomainSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post domain subdomain search params
func (o *PostDomainSubdomainSearchParams) WithHTTPClient(client *http.Client) *PostDomainSubdomainSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post domain subdomain search params
func (o *PostDomainSubdomainSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithQuery adds the query to the post domain subdomain search params
func (o *PostDomainSubdomainSearchParams) WithQuery(query *models.DNSDomainRequest) *PostDomainSubdomainSearchParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the post domain subdomain search params
func (o *PostDomainSubdomainSearchParams) SetQuery(query *models.DNSDomainRequest) {
	o.Query = query
}

// WriteToRequest writes these params to a swagger request
func (o *PostDomainSubdomainSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
