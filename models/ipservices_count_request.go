// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IpservicesCountRequest ipservices count request
//
// swagger:model ipservices.CountRequest
type IpservicesCountRequest struct {

	// query
	Query string `json:"query,omitempty"`
}

// Validate validates this ipservices count request
func (m *IpservicesCountRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ipservices count request based on context it is used
func (m *IpservicesCountRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IpservicesCountRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IpservicesCountRequest) UnmarshalBinary(b []byte) error {
	var res IpservicesCountRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
