// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CertificateCertificateSummaryResponse certificate certificate summary response
//
// swagger:model certificate.CertificateSummaryResponse
type CertificateCertificateSummaryResponse struct {

	// data
	Data *CertificateCertificateSummaryResponseData `json:"data,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// success
	Success bool `json:"success,omitempty"`
}

// Validate validates this certificate certificate summary response
func (m *CertificateCertificateSummaryResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CertificateCertificateSummaryResponse) validateData(formats strfmt.Registry) error {
	if swag.IsZero(m.Data) { // not required
		return nil
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this certificate certificate summary response based on the context it is used
func (m *CertificateCertificateSummaryResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CertificateCertificateSummaryResponse) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if m.Data != nil {

		if swag.IsZero(m.Data) { // not required
			return nil
		}

		if err := m.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CertificateCertificateSummaryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CertificateCertificateSummaryResponse) UnmarshalBinary(b []byte) error {
	var res CertificateCertificateSummaryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CertificateCertificateSummaryResponseData certificate certificate summary response data
//
// swagger:model CertificateCertificateSummaryResponseData
type CertificateCertificateSummaryResponseData struct {

	// buckets
	Buckets []*CertificateCertificateSummaryResponseDataBucketsItems0 `json:"buckets"`

	// doc count error upper bound
	DocCountErrorUpperBound int64 `json:"doc_count_error_upper_bound,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// sum other doc count
	SumOtherDocCount int64 `json:"sum_other_doc_count,omitempty"`
}

// Validate validates this certificate certificate summary response data
func (m *CertificateCertificateSummaryResponseData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuckets(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CertificateCertificateSummaryResponseData) validateBuckets(formats strfmt.Registry) error {
	if swag.IsZero(m.Buckets) { // not required
		return nil
	}

	for i := 0; i < len(m.Buckets); i++ {
		if swag.IsZero(m.Buckets[i]) { // not required
			continue
		}

		if m.Buckets[i] != nil {
			if err := m.Buckets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + "buckets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("data" + "." + "buckets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this certificate certificate summary response data based on the context it is used
func (m *CertificateCertificateSummaryResponseData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBuckets(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CertificateCertificateSummaryResponseData) contextValidateBuckets(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Buckets); i++ {

		if m.Buckets[i] != nil {

			if swag.IsZero(m.Buckets[i]) { // not required
				return nil
			}

			if err := m.Buckets[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + "buckets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("data" + "." + "buckets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CertificateCertificateSummaryResponseData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CertificateCertificateSummaryResponseData) UnmarshalBinary(b []byte) error {
	var res CertificateCertificateSummaryResponseData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CertificateCertificateSummaryResponseDataBucketsItems0 certificate certificate summary response data buckets items0
//
// swagger:model CertificateCertificateSummaryResponseDataBucketsItems0
type CertificateCertificateSummaryResponseDataBucketsItems0 struct {

	// doc count
	DocCount int64 `json:"doc_count,omitempty"`

	// key
	Key string `json:"key,omitempty"`
}

// Validate validates this certificate certificate summary response data buckets items0
func (m *CertificateCertificateSummaryResponseDataBucketsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this certificate certificate summary response data buckets items0 based on context it is used
func (m *CertificateCertificateSummaryResponseDataBucketsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CertificateCertificateSummaryResponseDataBucketsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CertificateCertificateSummaryResponseDataBucketsItems0) UnmarshalBinary(b []byte) error {
	var res CertificateCertificateSummaryResponseDataBucketsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
