// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DNSAPIResponse dns API response
//
// swagger:model dns.APIResponse
type DNSAPIResponse struct {

	// data
	Data interface{} `json:"data,omitempty"`

	// pagination
	Pagination interface{} `json:"pagination,omitempty"`

	// success
	Success bool `json:"success,omitempty"`
}

// Validate validates this dns API response
func (m *DNSAPIResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this dns API response based on context it is used
func (m *DNSAPIResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DNSAPIResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DNSAPIResponse) UnmarshalBinary(b []byte) error {
	var res DNSAPIResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
