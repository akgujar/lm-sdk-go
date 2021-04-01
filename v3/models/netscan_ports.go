// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NetscanPorts netscan ports
//
// swagger:model NetscanPorts
type NetscanPorts struct {

	// Whether or not default ports should be used
	// Example: true
	IsGlobalDefault bool `json:"isGlobalDefault,omitempty"`

	// The ports that should be used in the Netscan
	// Example: 21,22,23
	Value string `json:"value,omitempty"`
}

// Validate validates this netscan ports
func (m *NetscanPorts) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this netscan ports based on context it is used
func (m *NetscanPorts) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NetscanPorts) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetscanPorts) UnmarshalBinary(b []byte) error {
	var res NetscanPorts
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}