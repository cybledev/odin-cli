// Code generated by go-swagger; DO NOT EDIT.

package certificate

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

// NewPostCertificatesScrollNextParams creates a new PostCertificatesScrollNextParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostCertificatesScrollNextParams() *PostCertificatesScrollNextParams {
	return &PostCertificatesScrollNextParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostCertificatesScrollNextParamsWithTimeout creates a new PostCertificatesScrollNextParams object
// with the ability to set a timeout on a request.
func NewPostCertificatesScrollNextParamsWithTimeout(timeout time.Duration) *PostCertificatesScrollNextParams {
	return &PostCertificatesScrollNextParams{
		timeout: timeout,
	}
}

// NewPostCertificatesScrollNextParamsWithContext creates a new PostCertificatesScrollNextParams object
// with the ability to set a context for a request.
func NewPostCertificatesScrollNextParamsWithContext(ctx context.Context) *PostCertificatesScrollNextParams {
	return &PostCertificatesScrollNextParams{
		Context: ctx,
	}
}

// NewPostCertificatesScrollNextParamsWithHTTPClient creates a new PostCertificatesScrollNextParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostCertificatesScrollNextParamsWithHTTPClient(client *http.Client) *PostCertificatesScrollNextParams {
	return &PostCertificatesScrollNextParams{
		HTTPClient: client,
	}
}

/*
PostCertificatesScrollNextParams contains all the parameters to send to the API endpoint

	for the post certificates scroll next operation.

	Typically these are written to a http.Request.
*/
type PostCertificatesScrollNextParams struct {

	/* Query.

	   Search Query
	*/
	Query *models.CertificateNextBatchRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post certificates scroll next params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostCertificatesScrollNextParams) WithDefaults() *PostCertificatesScrollNextParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post certificates scroll next params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostCertificatesScrollNextParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post certificates scroll next params
func (o *PostCertificatesScrollNextParams) WithTimeout(timeout time.Duration) *PostCertificatesScrollNextParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post certificates scroll next params
func (o *PostCertificatesScrollNextParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post certificates scroll next params
func (o *PostCertificatesScrollNextParams) WithContext(ctx context.Context) *PostCertificatesScrollNextParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post certificates scroll next params
func (o *PostCertificatesScrollNextParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post certificates scroll next params
func (o *PostCertificatesScrollNextParams) WithHTTPClient(client *http.Client) *PostCertificatesScrollNextParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post certificates scroll next params
func (o *PostCertificatesScrollNextParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithQuery adds the query to the post certificates scroll next params
func (o *PostCertificatesScrollNextParams) WithQuery(query *models.CertificateNextBatchRequest) *PostCertificatesScrollNextParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the post certificates scroll next params
func (o *PostCertificatesScrollNextParams) SetQuery(query *models.CertificateNextBatchRequest) {
	o.Query = query
}

// WriteToRequest writes these params to a swagger request
func (o *PostCertificatesScrollNextParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
