// Code generated by go-swagger; DO NOT EDIT.

package exposed_buckets

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

// NewPostExposedBucketsSummaryParams creates a new PostExposedBucketsSummaryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostExposedBucketsSummaryParams() *PostExposedBucketsSummaryParams {
	return &PostExposedBucketsSummaryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostExposedBucketsSummaryParamsWithTimeout creates a new PostExposedBucketsSummaryParams object
// with the ability to set a timeout on a request.
func NewPostExposedBucketsSummaryParamsWithTimeout(timeout time.Duration) *PostExposedBucketsSummaryParams {
	return &PostExposedBucketsSummaryParams{
		timeout: timeout,
	}
}

// NewPostExposedBucketsSummaryParamsWithContext creates a new PostExposedBucketsSummaryParams object
// with the ability to set a context for a request.
func NewPostExposedBucketsSummaryParamsWithContext(ctx context.Context) *PostExposedBucketsSummaryParams {
	return &PostExposedBucketsSummaryParams{
		Context: ctx,
	}
}

// NewPostExposedBucketsSummaryParamsWithHTTPClient creates a new PostExposedBucketsSummaryParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostExposedBucketsSummaryParamsWithHTTPClient(client *http.Client) *PostExposedBucketsSummaryParams {
	return &PostExposedBucketsSummaryParams{
		HTTPClient: client,
	}
}

/*
PostExposedBucketsSummaryParams contains all the parameters to send to the API endpoint

	for the post exposed buckets summary operation.

	Typically these are written to a http.Request.
*/
type PostExposedBucketsSummaryParams struct {

	/* Query.

	   Summary Request
	*/
	Query *models.ExposedSummaryRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post exposed buckets summary params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostExposedBucketsSummaryParams) WithDefaults() *PostExposedBucketsSummaryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post exposed buckets summary params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostExposedBucketsSummaryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post exposed buckets summary params
func (o *PostExposedBucketsSummaryParams) WithTimeout(timeout time.Duration) *PostExposedBucketsSummaryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post exposed buckets summary params
func (o *PostExposedBucketsSummaryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post exposed buckets summary params
func (o *PostExposedBucketsSummaryParams) WithContext(ctx context.Context) *PostExposedBucketsSummaryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post exposed buckets summary params
func (o *PostExposedBucketsSummaryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post exposed buckets summary params
func (o *PostExposedBucketsSummaryParams) WithHTTPClient(client *http.Client) *PostExposedBucketsSummaryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post exposed buckets summary params
func (o *PostExposedBucketsSummaryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithQuery adds the query to the post exposed buckets summary params
func (o *PostExposedBucketsSummaryParams) WithQuery(query *models.ExposedSummaryRequest) *PostExposedBucketsSummaryParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the post exposed buckets summary params
func (o *PostExposedBucketsSummaryParams) SetQuery(query *models.ExposedSummaryRequest) {
	o.Query = query
}

// WriteToRequest writes these params to a swagger request
func (o *PostExposedBucketsSummaryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
