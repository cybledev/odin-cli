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

// DNSDNSRequest dns DNS request
//
// swagger:model dns.DNSRequest
type DNSDNSRequest struct {

	// domain
	// Required: true
	Domain *string `json:"domain"`

	// limit
	// Required: true
	Limit *int64 `json:"limit"`

	// published after
	PublishedAfter string `json:"publishedAfter,omitempty"`

	// published before
	PublishedBefore string `json:"publishedBefore,omitempty"`

	// start
	Start []float64 `json:"start"`
}

// Validate validates this dns DNS request
func (m *DNSDNSRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDomain(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLimit(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DNSDNSRequest) validateDomain(formats strfmt.Registry) error {

	if err := validate.Required("domain", "body", m.Domain); err != nil {
		return err
	}

	return nil
}

func (m *DNSDNSRequest) validateLimit(formats strfmt.Registry) error {

	if err := validate.Required("limit", "body", m.Limit); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this dns DNS request based on context it is used
func (m *DNSDNSRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DNSDNSRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DNSDNSRequest) UnmarshalBinary(b []byte) error {
	var res DNSDNSRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
