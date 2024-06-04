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

// AwsGlobalNetworkDeviceMethod aws global network device method
//
// swagger:model AwsGlobalNetworkDeviceMethod
type AwsGlobalNetworkDeviceMethod struct {
	AwsGlobalNetworkDeviceMethodAllOf1
}

// Name gets the name of this subtype
func (m *AwsGlobalNetworkDeviceMethod) Name() string {
	return "AwsGlobalNetworkDeviceMethod"
}

// SetName sets the name of this subtype
func (m *AwsGlobalNetworkDeviceMethod) SetName(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *AwsGlobalNetworkDeviceMethod) UnmarshalJSON(raw []byte) error {
	var data struct {
		AwsGlobalNetworkDeviceMethodAllOf1
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

	var result AwsGlobalNetworkDeviceMethod

	if base.Name != result.Name() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid name value: %q", base.Name)
	}
	result.AwsGlobalNetworkDeviceMethodAllOf1 = data.AwsGlobalNetworkDeviceMethodAllOf1

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m AwsGlobalNetworkDeviceMethod) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		AwsGlobalNetworkDeviceMethodAllOf1
	}{

		AwsGlobalNetworkDeviceMethodAllOf1: m.AwsGlobalNetworkDeviceMethodAllOf1,
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

// Validate validates this aws global network device method
func (m *AwsGlobalNetworkDeviceMethod) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with AwsGlobalNetworkDeviceMethodAllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this aws global network device method based on the context it is used
func (m *AwsGlobalNetworkDeviceMethod) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with AwsGlobalNetworkDeviceMethodAllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *AwsGlobalNetworkDeviceMethod) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AwsGlobalNetworkDeviceMethod) UnmarshalBinary(b []byte) error {
	var res AwsGlobalNetworkDeviceMethod
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AwsGlobalNetworkDeviceMethodAllOf1 aws global network device method all of1
//
// swagger:model AwsGlobalNetworkDeviceMethodAllOf1
type AwsGlobalNetworkDeviceMethodAllOf1 interface{}
