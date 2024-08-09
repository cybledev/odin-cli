// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ExposedFile exposed file
//
// swagger:model exposed.File
type ExposedFile struct {

	// accessible
	Accessible bool `json:"accessible,omitempty"`

	// bucket
	Bucket string `json:"bucket,omitempty"`

	// category
	Category string `json:"category,omitempty"`

	// etag
	Etag string `json:"etag,omitempty"`

	// ext
	Ext string `json:"ext,omitempty"`

	// ext desc
	ExtDesc string `json:"ext_desc,omitempty"`

	// ins at
	InsAt string `json:"ins_at,omitempty"`

	// label
	Label string `json:"label,omitempty"`

	// mod at
	ModAt string `json:"mod_at,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// path
	Path string `json:"path,omitempty"`

	// provider
	Provider string `json:"provider,omitempty"`

	// region
	Region string `json:"region,omitempty"`

	// scan at
	ScanAt string `json:"scan_at,omitempty"`

	// size
	Size int64 `json:"size,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this exposed file
func (m *ExposedFile) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this exposed file based on context it is used
func (m *ExposedFile) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ExposedFile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExposedFile) UnmarshalBinary(b []byte) error {
	var res ExposedFile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
