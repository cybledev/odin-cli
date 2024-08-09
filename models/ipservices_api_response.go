// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IpservicesAPIResponse ipservices API response
//
// swagger:model ipservices.APIResponse
type IpservicesAPIResponse struct {

	// data
	Data interface{} `json:"data,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// pagination
	Pagination interface{} `json:"pagination,omitempty"`

	// success
	Success bool `json:"success,omitempty"`
}

// Validate validates this ipservices API response
func (m *IpservicesAPIResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ipservices API response based on context it is used
func (m *IpservicesAPIResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IpservicesAPIResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IpservicesAPIResponse) UnmarshalBinary(b []byte) error {
	var res IpservicesAPIResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
