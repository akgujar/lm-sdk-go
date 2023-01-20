// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InterfaceType interface type
//
// swagger:model InterfaceType
type InterfaceType struct {

	// if Id
	IfID int64 `json:"ifId,omitempty"`

	// if position
	IfPosition string `json:"ifPosition,omitempty"`
}

// Validate validates this interface type
func (m *InterfaceType) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this interface type based on context it is used
func (m *InterfaceType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InterfaceType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InterfaceType) UnmarshalBinary(b []byte) error {
	var res InterfaceType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
