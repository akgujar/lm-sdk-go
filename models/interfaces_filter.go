// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InterfacesFilter interfaces filter
//
// swagger:model InterfacesFilter
type InterfacesFilter struct {

	// device Id
	DeviceID int32 `json:"deviceId,omitempty"`

	// interface types
	InterfaceTypes []*InterfaceType `json:"interfaceTypes,omitempty"`
}

// Validate validates this interfaces filter
func (m *InterfacesFilter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInterfaceTypes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InterfacesFilter) validateInterfaceTypes(formats strfmt.Registry) error {
	if swag.IsZero(m.InterfaceTypes) { // not required
		return nil
	}

	for i := 0; i < len(m.InterfaceTypes); i++ {
		if swag.IsZero(m.InterfaceTypes[i]) { // not required
			continue
		}

		if m.InterfaceTypes[i] != nil {
			if err := m.InterfaceTypes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("interfaceTypes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("interfaceTypes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this interfaces filter based on the context it is used
func (m *InterfacesFilter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInterfaceTypes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InterfacesFilter) contextValidateInterfaceTypes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.InterfaceTypes); i++ {

		if m.InterfaceTypes[i] != nil {

			if swag.IsZero(m.InterfaceTypes[i]) { // not required
				return nil
			}

			if err := m.InterfaceTypes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("interfaceTypes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("interfaceTypes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *InterfacesFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InterfacesFilter) UnmarshalBinary(b []byte) error {
	var res InterfacesFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
