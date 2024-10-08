// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ServiceFullCveData service full cve data
//
// swagger:model service.FullCveData
type ServiceFullCveData struct {

	// id
	ID string `json:"id,omitempty"`

	// score
	Score float64 `json:"score,omitempty"`

	// severity
	Severity string `json:"severity,omitempty"`
}

// Validate validates this service full cve data
func (m *ServiceFullCveData) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this service full cve data based on context it is used
func (m *ServiceFullCveData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceFullCveData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceFullCveData) UnmarshalBinary(b []byte) error {
	var res ServiceFullCveData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
