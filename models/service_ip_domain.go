// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ServiceIPDomain service IP domain
//
// swagger:model service.IPDomain
type ServiceIPDomain struct {

	// country
	Country string `json:"country,omitempty"`

	// dns servers
	DNSServers []string `json:"dns_servers"`

	// emails
	Emails []string `json:"emails"`

	// last updated at
	LastUpdatedAt string `json:"last_updated_at,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// web server
	WebServer string `json:"web_server,omitempty"`
}

// Validate validates this service IP domain
func (m *ServiceIPDomain) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this service IP domain based on context it is used
func (m *ServiceIPDomain) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceIPDomain) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceIPDomain) UnmarshalBinary(b []byte) error {
	var res ServiceIPDomain
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
