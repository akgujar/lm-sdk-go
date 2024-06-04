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
	"github.com/go-openapi/validate"
)

// EntityProperty entity property
//
// swagger:model EntityProperty
type EntityProperty struct {

	// The inherit list of the property
	// Read Only: true
	InheritList []*InheritanceProp `json:"inheritList,omitempty"`

	// The name of the property
	// Read Only: true
	Name string `json:"name,omitempty"`

	// The type of property. The values can be Inherit|System|Custom
	// Read Only: true
	Type string `json:"type,omitempty"`

	// The value of the property
	// Read Only: true
	Value string `json:"value,omitempty"`
}

// Validate validates this entity property
func (m *EntityProperty) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInheritList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EntityProperty) validateInheritList(formats strfmt.Registry) error {
	if swag.IsZero(m.InheritList) { // not required
		return nil
	}

	for i := 0; i < len(m.InheritList); i++ {
		if swag.IsZero(m.InheritList[i]) { // not required
			continue
		}

		if m.InheritList[i] != nil {
			if err := m.InheritList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("inheritList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("inheritList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this entity property based on the context it is used
func (m *EntityProperty) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInheritList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateValue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EntityProperty) contextValidateInheritList(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "inheritList", "body", []*InheritanceProp(m.InheritList)); err != nil {
		return err
	}

	for i := 0; i < len(m.InheritList); i++ {

		if m.InheritList[i] != nil {

			if swag.IsZero(m.InheritList[i]) { // not required
				return nil
			}

			if err := m.InheritList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("inheritList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("inheritList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EntityProperty) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

func (m *EntityProperty) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "type", "body", string(m.Type)); err != nil {
		return err
	}

	return nil
}

func (m *EntityProperty) contextValidateValue(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "value", "body", string(m.Value)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EntityProperty) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EntityProperty) UnmarshalBinary(b []byte) error {
	var res EntityProperty
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
