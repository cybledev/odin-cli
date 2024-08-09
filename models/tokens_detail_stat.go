// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TokensDetailStat tokens detail stat
//
// swagger:model tokens.DetailStat
type TokensDetailStat struct {

	// data
	Data interface{} `json:"data,omitempty"`

	// left
	Left int64 `json:"left,omitempty"`

	// plan
	Plan int64 `json:"plan,omitempty"`
}

// Validate validates this tokens detail stat
func (m *TokensDetailStat) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this tokens detail stat based on context it is used
func (m *TokensDetailStat) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TokensDetailStat) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TokensDetailStat) UnmarshalBinary(b []byte) error {
	var res TokensDetailStat
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
