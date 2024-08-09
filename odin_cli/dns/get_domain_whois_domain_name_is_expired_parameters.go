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
)

// NewGetDomainWhoisDomainNameIsExpiredParams creates a new GetDomainWhoisDomainNameIsExpiredParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDomainWhoisDomainNameIsExpiredParams() *GetDomainWhoisDomainNameIsExpiredParams {
	return &GetDomainWhoisDomainNameIsExpiredParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDomainWhoisDomainNameIsExpiredParamsWithTimeout creates a new GetDomainWhoisDomainNameIsExpiredParams object
// with the ability to set a timeout on a request.
func NewGetDomainWhoisDomainNameIsExpiredParamsWithTimeout(timeout time.Duration) *GetDomainWhoisDomainNameIsExpiredParams {
	return &GetDomainWhoisDomainNameIsExpiredParams{
		timeout: timeout,
	}
}

// NewGetDomainWhoisDomainNameIsExpiredParamsWithContext creates a new GetDomainWhoisDomainNameIsExpiredParams object
// with the ability to set a context for a request.
func NewGetDomainWhoisDomainNameIsExpiredParamsWithContext(ctx context.Context) *GetDomainWhoisDomainNameIsExpiredParams {
	return &GetDomainWhoisDomainNameIsExpiredParams{
		Context: ctx,
	}
}

// NewGetDomainWhoisDomainNameIsExpiredParamsWithHTTPClient creates a new GetDomainWhoisDomainNameIsExpiredParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDomainWhoisDomainNameIsExpiredParamsWithHTTPClient(client *http.Client) *GetDomainWhoisDomainNameIsExpiredParams {
	return &GetDomainWhoisDomainNameIsExpiredParams{
		HTTPClient: client,
	}
}

/*
GetDomainWhoisDomainNameIsExpiredParams contains all the parameters to send to the API endpoint

	for the get domain whois domain name is expired operation.

	Typically these are written to a http.Request.
*/
type GetDomainWhoisDomainNameIsExpiredParams struct {

	/* DomainName.

	   domain
	*/
	DomainName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get domain whois domain name is expired params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDomainWhoisDomainNameIsExpiredParams) WithDefaults() *GetDomainWhoisDomainNameIsExpiredParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get domain whois domain name is expired params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDomainWhoisDomainNameIsExpiredParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get domain whois domain name is expired params
func (o *GetDomainWhoisDomainNameIsExpiredParams) WithTimeout(timeout time.Duration) *GetDomainWhoisDomainNameIsExpiredParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get domain whois domain name is expired params
func (o *GetDomainWhoisDomainNameIsExpiredParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get domain whois domain name is expired params
func (o *GetDomainWhoisDomainNameIsExpiredParams) WithContext(ctx context.Context) *GetDomainWhoisDomainNameIsExpiredParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get domain whois domain name is expired params
func (o *GetDomainWhoisDomainNameIsExpiredParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get domain whois domain name is expired params
func (o *GetDomainWhoisDomainNameIsExpiredParams) WithHTTPClient(client *http.Client) *GetDomainWhoisDomainNameIsExpiredParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get domain whois domain name is expired params
func (o *GetDomainWhoisDomainNameIsExpiredParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomainName adds the domainName to the get domain whois domain name is expired params
func (o *GetDomainWhoisDomainNameIsExpiredParams) WithDomainName(domainName string) *GetDomainWhoisDomainNameIsExpiredParams {
	o.SetDomainName(domainName)
	return o
}

// SetDomainName adds the domainName to the get domain whois domain name is expired params
func (o *GetDomainWhoisDomainNameIsExpiredParams) SetDomainName(domainName string) {
	o.DomainName = domainName
}

// WriteToRequest writes these params to a swagger request
func (o *GetDomainWhoisDomainNameIsExpiredParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param domain-name
	if err := r.SetPathParam("domain-name", o.DomainName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
