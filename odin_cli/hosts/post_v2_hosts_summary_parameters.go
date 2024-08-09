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

	"github.com/cybledev/odin-cli/models"
)

// NewPostV2HostsSummaryParams creates a new PostV2HostsSummaryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV2HostsSummaryParams() *PostV2HostsSummaryParams {
	return &PostV2HostsSummaryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV2HostsSummaryParamsWithTimeout creates a new PostV2HostsSummaryParams object
// with the ability to set a timeout on a request.
func NewPostV2HostsSummaryParamsWithTimeout(timeout time.Duration) *PostV2HostsSummaryParams {
	return &PostV2HostsSummaryParams{
		timeout: timeout,
	}
}

// NewPostV2HostsSummaryParamsWithContext creates a new PostV2HostsSummaryParams object
// with the ability to set a context for a request.
func NewPostV2HostsSummaryParamsWithContext(ctx context.Context) *PostV2HostsSummaryParams {
	return &PostV2HostsSummaryParams{
		Context: ctx,
	}
}

// NewPostV2HostsSummaryParamsWithHTTPClient creates a new PostV2HostsSummaryParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV2HostsSummaryParamsWithHTTPClient(client *http.Client) *PostV2HostsSummaryParams {
	return &PostV2HostsSummaryParams{
		HTTPClient: client,
	}
}

/*
PostV2HostsSummaryParams contains all the parameters to send to the API endpoint

	for the post v2 hosts summary operation.

	Typically these are written to a http.Request.
*/
type PostV2HostsSummaryParams struct {

	/* Query.

	   Summary
	*/
	Query *models.CybleComOdinAPIControllersV2IpservicesSummaryRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v2 hosts summary params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV2HostsSummaryParams) WithDefaults() *PostV2HostsSummaryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v2 hosts summary params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV2HostsSummaryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v2 hosts summary params
func (o *PostV2HostsSummaryParams) WithTimeout(timeout time.Duration) *PostV2HostsSummaryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v2 hosts summary params
func (o *PostV2HostsSummaryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v2 hosts summary params
func (o *PostV2HostsSummaryParams) WithContext(ctx context.Context) *PostV2HostsSummaryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v2 hosts summary params
func (o *PostV2HostsSummaryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v2 hosts summary params
func (o *PostV2HostsSummaryParams) WithHTTPClient(client *http.Client) *PostV2HostsSummaryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v2 hosts summary params
func (o *PostV2HostsSummaryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithQuery adds the query to the post v2 hosts summary params
func (o *PostV2HostsSummaryParams) WithQuery(query *models.CybleComOdinAPIControllersV2IpservicesSummaryRequest) *PostV2HostsSummaryParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the post v2 hosts summary params
func (o *PostV2HostsSummaryParams) SetQuery(query *models.CybleComOdinAPIControllersV2IpservicesSummaryRequest) {
	o.Query = query
}

// WriteToRequest writes these params to a swagger request
func (o *PostV2HostsSummaryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
