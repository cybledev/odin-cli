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

// NewPostExposedFilesSearchParams creates a new PostExposedFilesSearchParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostExposedFilesSearchParams() *PostExposedFilesSearchParams {
	return &PostExposedFilesSearchParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostExposedFilesSearchParamsWithTimeout creates a new PostExposedFilesSearchParams object
// with the ability to set a timeout on a request.
func NewPostExposedFilesSearchParamsWithTimeout(timeout time.Duration) *PostExposedFilesSearchParams {
	return &PostExposedFilesSearchParams{
		timeout: timeout,
	}
}

// NewPostExposedFilesSearchParamsWithContext creates a new PostExposedFilesSearchParams object
// with the ability to set a context for a request.
func NewPostExposedFilesSearchParamsWithContext(ctx context.Context) *PostExposedFilesSearchParams {
	return &PostExposedFilesSearchParams{
		Context: ctx,
	}
}

// NewPostExposedFilesSearchParamsWithHTTPClient creates a new PostExposedFilesSearchParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostExposedFilesSearchParamsWithHTTPClient(client *http.Client) *PostExposedFilesSearchParams {
	return &PostExposedFilesSearchParams{
		HTTPClient: client,
	}
}

/*
PostExposedFilesSearchParams contains all the parameters to send to the API endpoint

	for the post exposed files search operation.

	Typically these are written to a http.Request.
*/
type PostExposedFilesSearchParams struct {

	/* Query.

	   Search Query
	*/
	Query *models.ExposedSearchRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post exposed files search params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostExposedFilesSearchParams) WithDefaults() *PostExposedFilesSearchParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post exposed files search params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostExposedFilesSearchParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post exposed files search params
func (o *PostExposedFilesSearchParams) WithTimeout(timeout time.Duration) *PostExposedFilesSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post exposed files search params
func (o *PostExposedFilesSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post exposed files search params
func (o *PostExposedFilesSearchParams) WithContext(ctx context.Context) *PostExposedFilesSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post exposed files search params
func (o *PostExposedFilesSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post exposed files search params
func (o *PostExposedFilesSearchParams) WithHTTPClient(client *http.Client) *PostExposedFilesSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post exposed files search params
func (o *PostExposedFilesSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithQuery adds the query to the post exposed files search params
func (o *PostExposedFilesSearchParams) WithQuery(query *models.ExposedSearchRequest) *PostExposedFilesSearchParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the post exposed files search params
func (o *PostExposedFilesSearchParams) SetQuery(query *models.ExposedSearchRequest) {
	o.Query = query
}

// WriteToRequest writes these params to a swagger request
func (o *PostExposedFilesSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
