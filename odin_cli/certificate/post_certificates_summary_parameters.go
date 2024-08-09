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

// NewPostCertificatesSummaryParams creates a new PostCertificatesSummaryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostCertificatesSummaryParams() *PostCertificatesSummaryParams {
	return &PostCertificatesSummaryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostCertificatesSummaryParamsWithTimeout creates a new PostCertificatesSummaryParams object
// with the ability to set a timeout on a request.
func NewPostCertificatesSummaryParamsWithTimeout(timeout time.Duration) *PostCertificatesSummaryParams {
	return &PostCertificatesSummaryParams{
		timeout: timeout,
	}
}

// NewPostCertificatesSummaryParamsWithContext creates a new PostCertificatesSummaryParams object
// with the ability to set a context for a request.
func NewPostCertificatesSummaryParamsWithContext(ctx context.Context) *PostCertificatesSummaryParams {
	return &PostCertificatesSummaryParams{
		Context: ctx,
	}
}

// NewPostCertificatesSummaryParamsWithHTTPClient creates a new PostCertificatesSummaryParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostCertificatesSummaryParamsWithHTTPClient(client *http.Client) *PostCertificatesSummaryParams {
	return &PostCertificatesSummaryParams{
		HTTPClient: client,
	}
}

/*
PostCertificatesSummaryParams contains all the parameters to send to the API endpoint

	for the post certificates summary operation.

	Typically these are written to a http.Request.
*/
type PostCertificatesSummaryParams struct {

	/* Query.

	   Summary
	*/
	Query *models.CertificateCertSummaryRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post certificates summary params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostCertificatesSummaryParams) WithDefaults() *PostCertificatesSummaryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post certificates summary params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostCertificatesSummaryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post certificates summary params
func (o *PostCertificatesSummaryParams) WithTimeout(timeout time.Duration) *PostCertificatesSummaryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post certificates summary params
func (o *PostCertificatesSummaryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post certificates summary params
func (o *PostCertificatesSummaryParams) WithContext(ctx context.Context) *PostCertificatesSummaryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post certificates summary params
func (o *PostCertificatesSummaryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post certificates summary params
func (o *PostCertificatesSummaryParams) WithHTTPClient(client *http.Client) *PostCertificatesSummaryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post certificates summary params
func (o *PostCertificatesSummaryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithQuery adds the query to the post certificates summary params
func (o *PostCertificatesSummaryParams) WithQuery(query *models.CertificateCertSummaryRequest) *PostCertificatesSummaryParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the post certificates summary params
func (o *PostCertificatesSummaryParams) SetQuery(query *models.CertificateCertSummaryRequest) {
	o.Query = query
}

// WriteToRequest writes these params to a swagger request
func (o *PostCertificatesSummaryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
