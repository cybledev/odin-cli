// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CertificateAPIResponse certificate API response
//
// swagger:model certificate.APIResponse
type CertificateAPIResponse struct {

	// data
	Data interface{} `json:"data,omitempty"`

	// pagination
	Pagination interface{} `json:"pagination,omitempty"`

	// success
	Success bool `json:"success,omitempty"`
}

// Validate validates this certificate API response
func (m *CertificateAPIResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this certificate API response based on context it is used
func (m *CertificateAPIResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CertificateAPIResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CertificateAPIResponse) UnmarshalBinary(b []byte) error {
	var res CertificateAPIResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
