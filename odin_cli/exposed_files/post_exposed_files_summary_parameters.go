// Code generated by go-swagger; DO NOT EDIT.

package exposed_files

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

// NewPostExposedFilesSummaryParams creates a new PostExposedFilesSummaryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostExposedFilesSummaryParams() *PostExposedFilesSummaryParams {
	return &PostExposedFilesSummaryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostExposedFilesSummaryParamsWithTimeout creates a new PostExposedFilesSummaryParams object
// with the ability to set a timeout on a request.
func NewPostExposedFilesSummaryParamsWithTimeout(timeout time.Duration) *PostExposedFilesSummaryParams {
	return &PostExposedFilesSummaryParams{
		timeout: timeout,
	}
}

// NewPostExposedFilesSummaryParamsWithContext creates a new PostExposedFilesSummaryParams object
// with the ability to set a context for a request.
func NewPostExposedFilesSummaryParamsWithContext(ctx context.Context) *PostExposedFilesSummaryParams {
	return &PostExposedFilesSummaryParams{
		Context: ctx,
	}
}

// NewPostExposedFilesSummaryParamsWithHTTPClient creates a new PostExposedFilesSummaryParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostExposedFilesSummaryParamsWithHTTPClient(client *http.Client) *PostExposedFilesSummaryParams {
	return &PostExposedFilesSummaryParams{
		HTTPClient: client,
	}
}

/*
PostExposedFilesSummaryParams contains all the parameters to send to the API endpoint

	for the post exposed files summary operation.

	Typically these are written to a http.Request.
*/
type PostExposedFilesSummaryParams struct {

	/* Query.

	   Summary Request
	*/
	Query *models.ExposedSummaryRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post exposed files summary params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostExposedFilesSummaryParams) WithDefaults() *PostExposedFilesSummaryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post exposed files summary params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostExposedFilesSummaryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post exposed files summary params
func (o *PostExposedFilesSummaryParams) WithTimeout(timeout time.Duration) *PostExposedFilesSummaryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post exposed files summary params
func (o *PostExposedFilesSummaryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post exposed files summary params
func (o *PostExposedFilesSummaryParams) WithContext(ctx context.Context) *PostExposedFilesSummaryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post exposed files summary params
func (o *PostExposedFilesSummaryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post exposed files summary params
func (o *PostExposedFilesSummaryParams) WithHTTPClient(client *http.Client) *PostExposedFilesSummaryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post exposed files summary params
func (o *PostExposedFilesSummaryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithQuery adds the query to the post exposed files summary params
func (o *PostExposedFilesSummaryParams) WithQuery(query *models.ExposedSummaryRequest) *PostExposedFilesSummaryParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the post exposed files summary params
func (o *PostExposedFilesSummaryParams) SetQuery(query *models.ExposedSummaryRequest) {
	o.Query = query
}

// WriteToRequest writes these params to a swagger request
func (o *PostExposedFilesSummaryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
