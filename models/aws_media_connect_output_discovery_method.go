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

// AwsMediaConnectOutputDiscoveryMethod aws media connect output discovery method
//
// swagger:model AwsMediaConnectOutputDiscoveryMethod
type AwsMediaConnectOutputDiscoveryMethod struct {
	AwsMediaConnectOutputDiscoveryMethodAllOf1
}

// Name gets the name of this subtype
func (m *AwsMediaConnectOutputDiscoveryMethod) Name() string {
	return "ad_awsmediaconnectoutput"
}

// SetName sets the name of this subtype
func (m *AwsMediaConnectOutputDiscoveryMethod) SetName(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *AwsMediaConnectOutputDiscoveryMethod) UnmarshalJSON(raw []byte) error {
	var data struct {
		AwsMediaConnectOutputDiscoveryMethodAllOf1
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

	var result AwsMediaConnectOutputDiscoveryMethod

	if base.Name != result.Name() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid name value: %q", base.Name)
	}
	result.AwsMediaConnectOutputDiscoveryMethodAllOf1 = data.AwsMediaConnectOutputDiscoveryMethodAllOf1

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m AwsMediaConnectOutputDiscoveryMethod) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		AwsMediaConnectOutputDiscoveryMethodAllOf1
	}{

		AwsMediaConnectOutputDiscoveryMethodAllOf1: m.AwsMediaConnectOutputDiscoveryMethodAllOf1,
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

// Validate validates this aws media connect output discovery method
func (m *AwsMediaConnectOutputDiscoveryMethod) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with AwsMediaConnectOutputDiscoveryMethodAllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this aws media connect output discovery method based on the context it is used
func (m *AwsMediaConnectOutputDiscoveryMethod) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with AwsMediaConnectOutputDiscoveryMethodAllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *AwsMediaConnectOutputDiscoveryMethod) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AwsMediaConnectOutputDiscoveryMethod) UnmarshalBinary(b []byte) error {
	var res AwsMediaConnectOutputDiscoveryMethod
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AwsMediaConnectOutputDiscoveryMethodAllOf1 aws media connect output discovery method all of1
//
// swagger:model AwsMediaConnectOutputDiscoveryMethodAllOf1
type AwsMediaConnectOutputDiscoveryMethodAllOf1 interface{}
