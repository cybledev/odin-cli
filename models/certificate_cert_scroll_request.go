// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CertificateCertScrollRequest certificate cert scroll request
//
// swagger:model certificate.CertScrollRequest
type CertificateCertScrollRequest struct {

	// from
	From int64 `json:"from,omitempty"`

	// limit
	// Required: true
	Limit *int64 `json:"limit"`

	// query
	// Required: true
	Query *string `json:"query"`
}

// Validate validates this certificate cert scroll request
func (m *CertificateCertScrollRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLimit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuery(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CertificateCertScrollRequest) validateLimit(formats strfmt.Registry) error {

	if err := validate.Required("limit", "body", m.Limit); err != nil {
		return err
	}

	return nil
}

func (m *CertificateCertScrollRequest) validateQuery(formats strfmt.Registry) error {

	if err := validate.Required("query", "body", m.Query); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this certificate cert scroll request based on context it is used
func (m *CertificateCertScrollRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CertificateCertScrollRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CertificateCertScrollRequest) UnmarshalBinary(b []byte) error {
	var res CertificateCertScrollRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
