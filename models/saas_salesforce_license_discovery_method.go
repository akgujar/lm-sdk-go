// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SaasSalesforceLicenseDiscoveryMethod saas salesforce license discovery method
//
// swagger:model SaasSalesforceLicenseDiscoveryMethod
type SaasSalesforceLicenseDiscoveryMethod struct {
	SaasSalesforceLicenseDiscoveryMethodAllOf1
}

// Name gets the name of this subtype
func (m *SaasSalesforceLicenseDiscoveryMethod) Name() string {
	return "SaasSalesforceLicenseDiscoveryMethod"
}

// SetName sets the name of this subtype
func (m *SaasSalesforceLicenseDiscoveryMethod) SetName(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *SaasSalesforceLicenseDiscoveryMethod) UnmarshalJSON(raw []byte) error {
	var data struct {
		SaasSalesforceLicenseDiscoveryMethodAllOf1
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		Name string `json:"name"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result SaasSalesforceLicenseDiscoveryMethod

	if base.Name != result.Name() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid name value: %q", base.Name)
	}
	result.SaasSalesforceLicenseDiscoveryMethodAllOf1 = data.SaasSalesforceLicenseDiscoveryMethodAllOf1

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m SaasSalesforceLicenseDiscoveryMethod) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		SaasSalesforceLicenseDiscoveryMethodAllOf1
	}{

		SaasSalesforceLicenseDiscoveryMethodAllOf1: m.SaasSalesforceLicenseDiscoveryMethodAllOf1,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Name string `json:"name"`
	}{

		Name: m.Name(),
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this saas salesforce license discovery method
func (m *SaasSalesforceLicenseDiscoveryMethod) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with SaasSalesforceLicenseDiscoveryMethodAllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this saas salesforce license discovery method based on the context it is used
func (m *SaasSalesforceLicenseDiscoveryMethod) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with SaasSalesforceLicenseDiscoveryMethodAllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *SaasSalesforceLicenseDiscoveryMethod) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SaasSalesforceLicenseDiscoveryMethod) UnmarshalBinary(b []byte) error {
	var res SaasSalesforceLicenseDiscoveryMethod
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SaasSalesforceLicenseDiscoveryMethodAllOf1 saas salesforce license discovery method all of1
//
// swagger:model SaasSalesforceLicenseDiscoveryMethodAllOf1
type SaasSalesforceLicenseDiscoveryMethodAllOf1 interface{}