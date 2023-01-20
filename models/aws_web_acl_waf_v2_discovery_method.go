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

// AwsWebACLWafV2DiscoveryMethod aws web ACL waf v2 discovery method
//
// swagger:model AwsWebACLWafV2DiscoveryMethod
type AwsWebACLWafV2DiscoveryMethod struct {
	AwsWebACLWafV2DiscoveryMethodAllOf1
}

// Name gets the name of this subtype
func (m *AwsWebACLWafV2DiscoveryMethod) Name() string {
	return "ad_awswebaclwafv2"
}

// SetName sets the name of this subtype
func (m *AwsWebACLWafV2DiscoveryMethod) SetName(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *AwsWebACLWafV2DiscoveryMethod) UnmarshalJSON(raw []byte) error {
	var data struct {
		AwsWebACLWafV2DiscoveryMethodAllOf1
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

	var result AwsWebACLWafV2DiscoveryMethod

	if base.Name != result.Name() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid name value: %q", base.Name)
	}
	result.AwsWebACLWafV2DiscoveryMethodAllOf1 = data.AwsWebACLWafV2DiscoveryMethodAllOf1

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m AwsWebACLWafV2DiscoveryMethod) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		AwsWebACLWafV2DiscoveryMethodAllOf1
	}{

		AwsWebACLWafV2DiscoveryMethodAllOf1: m.AwsWebACLWafV2DiscoveryMethodAllOf1,
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

// Validate validates this aws web ACL waf v2 discovery method
func (m *AwsWebACLWafV2DiscoveryMethod) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with AwsWebACLWafV2DiscoveryMethodAllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this aws web ACL waf v2 discovery method based on the context it is used
func (m *AwsWebACLWafV2DiscoveryMethod) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with AwsWebACLWafV2DiscoveryMethodAllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *AwsWebACLWafV2DiscoveryMethod) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AwsWebACLWafV2DiscoveryMethod) UnmarshalBinary(b []byte) error {
	var res AwsWebACLWafV2DiscoveryMethod
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AwsWebACLWafV2DiscoveryMethodAllOf1 aws web ACL waf v2 discovery method all of1
//
// swagger:model AwsWebACLWafV2DiscoveryMethodAllOf1
type AwsWebACLWafV2DiscoveryMethodAllOf1 interface{}
