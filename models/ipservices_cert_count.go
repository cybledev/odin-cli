// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IpservicesCertCount ipservices cert count
//
// swagger:model ipservices.CertCount
type IpservicesCertCount struct {

	// count
	Count int64 `json:"count,omitempty"`
}

// Validate validates this ipservices cert count
func (m *IpservicesCertCount) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ipservices cert count based on context it is used
func (m *IpservicesCertCount) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IpservicesCertCount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IpservicesCertCount) UnmarshalBinary(b []byte) error {
	var res IpservicesCertCount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
