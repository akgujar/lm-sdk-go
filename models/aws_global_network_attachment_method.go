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

// AwsGlobalNetworkAttachmentMethod aws global network attachment method
//
// swagger:model AwsGlobalNetworkAttachmentMethod
type AwsGlobalNetworkAttachmentMethod struct {
	AwsGlobalNetworkAttachmentMethodAllOf1
}

// Name gets the name of this subtype
func (m *AwsGlobalNetworkAttachmentMethod) Name() string {
	return "AwsGlobalNetworkAttachmentMethod"
}

// SetName sets the name of this subtype
func (m *AwsGlobalNetworkAttachmentMethod) SetName(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *AwsGlobalNetworkAttachmentMethod) UnmarshalJSON(raw []byte) error {
	var data struct {
		AwsGlobalNetworkAttachmentMethodAllOf1
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

	var result AwsGlobalNetworkAttachmentMethod

	if base.Name != result.Name() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid name value: %q", base.Name)
	}
	result.AwsGlobalNetworkAttachmentMethodAllOf1 = data.AwsGlobalNetworkAttachmentMethodAllOf1

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m AwsGlobalNetworkAttachmentMethod) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		AwsGlobalNetworkAttachmentMethodAllOf1
	}{

		AwsGlobalNetworkAttachmentMethodAllOf1: m.AwsGlobalNetworkAttachmentMethodAllOf1,
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

// Validate validates this aws global network attachment method
func (m *AwsGlobalNetworkAttachmentMethod) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with AwsGlobalNetworkAttachmentMethodAllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this aws global network attachment method based on the context it is used
func (m *AwsGlobalNetworkAttachmentMethod) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with AwsGlobalNetworkAttachmentMethodAllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *AwsGlobalNetworkAttachmentMethod) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AwsGlobalNetworkAttachmentMethod) UnmarshalBinary(b []byte) error {
	var res AwsGlobalNetworkAttachmentMethod
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AwsGlobalNetworkAttachmentMethodAllOf1 aws global network attachment method all of1
//
// swagger:model AwsGlobalNetworkAttachmentMethodAllOf1
type AwsGlobalNetworkAttachmentMethodAllOf1 interface{}
