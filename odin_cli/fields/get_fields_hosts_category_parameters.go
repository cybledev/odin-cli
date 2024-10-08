// Code generated by go-swagger; DO NOT EDIT.

package fields

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

// NewGetFieldsHostsCategoryParams creates a new GetFieldsHostsCategoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetFieldsHostsCategoryParams() *GetFieldsHostsCategoryParams {
	return &GetFieldsHostsCategoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetFieldsHostsCategoryParamsWithTimeout creates a new GetFieldsHostsCategoryParams object
// with the ability to set a timeout on a request.
func NewGetFieldsHostsCategoryParamsWithTimeout(timeout time.Duration) *GetFieldsHostsCategoryParams {
	return &GetFieldsHostsCategoryParams{
		timeout: timeout,
	}
}

// NewGetFieldsHostsCategoryParamsWithContext creates a new GetFieldsHostsCategoryParams object
// with the ability to set a context for a request.
func NewGetFieldsHostsCategoryParamsWithContext(ctx context.Context) *GetFieldsHostsCategoryParams {
	return &GetFieldsHostsCategoryParams{
		Context: ctx,
	}
}

// NewGetFieldsHostsCategoryParamsWithHTTPClient creates a new GetFieldsHostsCategoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetFieldsHostsCategoryParamsWithHTTPClient(client *http.Client) *GetFieldsHostsCategoryParams {
	return &GetFieldsHostsCategoryParams{
		HTTPClient: client,
	}
}

/*
GetFieldsHostsCategoryParams contains all the parameters to send to the API endpoint

	for the get fields hosts category operation.

	Typically these are written to a http.Request.
*/
type GetFieldsHostsCategoryParams struct {

	/* Category.

	   get the category
	*/
	Category string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get fields hosts category params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFieldsHostsCategoryParams) WithDefaults() *GetFieldsHostsCategoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get fields hosts category params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFieldsHostsCategoryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get fields hosts category params
func (o *GetFieldsHostsCategoryParams) WithTimeout(timeout time.Duration) *GetFieldsHostsCategoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get fields hosts category params
func (o *GetFieldsHostsCategoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get fields hosts category params
func (o *GetFieldsHostsCategoryParams) WithContext(ctx context.Context) *GetFieldsHostsCategoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get fields hosts category params
func (o *GetFieldsHostsCategoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get fields hosts category params
func (o *GetFieldsHostsCategoryParams) WithHTTPClient(client *http.Client) *GetFieldsHostsCategoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get fields hosts category params
func (o *GetFieldsHostsCategoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCategory adds the category to the get fields hosts category params
func (o *GetFieldsHostsCategoryParams) WithCategory(category string) *GetFieldsHostsCategoryParams {
	o.SetCategory(category)
	return o
}

// SetCategory adds the category to the get fields hosts category params
func (o *GetFieldsHostsCategoryParams) SetCategory(category string) {
	o.Category = category
}

// WriteToRequest writes these params to a swagger request
func (o *GetFieldsHostsCategoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param category
	if err := r.SetPathParam("category", o.Category); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
