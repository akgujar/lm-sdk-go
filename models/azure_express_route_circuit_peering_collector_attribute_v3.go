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
	"github.com/go-openapi/validate"
)

// AzureExpressRouteCircuitPeeringCollectorAttributeV3 azure express route circuit peering collector attribute v3
//
// swagger:model AzureExpressRouteCircuitPeeringCollectorAttributeV3
type AzureExpressRouteCircuitPeeringCollectorAttributeV3 struct {

	// period
	// Required: true
	Period *int32 `json:"period"`
}

// Name gets the name of this subtype
func (m *AzureExpressRouteCircuitPeeringCollectorAttributeV3) Name() string {
	return "AzureExpressRouteCircuitPeeringCollectorAttributeV3"
}

// SetName sets the name of this subtype
func (m *AzureExpressRouteCircuitPeeringCollectorAttributeV3) SetName(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *AzureExpressRouteCircuitPeeringCollectorAttributeV3) UnmarshalJSON(raw []byte) error {
	var data struct {

		// period
		// Required: true
		Period *int32 `json:"period"`
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

	var result AzureExpressRouteCircuitPeeringCollectorAttributeV3

	if base.Name != result.Name() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid name value: %q", base.Name)
	}

	result.Period = data.Period

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m AzureExpressRouteCircuitPeeringCollectorAttributeV3) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// period
		// Required: true
		Period *int32 `json:"period"`
	}{

		Period: m.Period,
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

// Validate validates this azure express route circuit peering collector attribute v3
func (m *AzureExpressRouteCircuitPeeringCollectorAttributeV3) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePeriod(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AzureExpressRouteCircuitPeeringCollectorAttributeV3) validatePeriod(formats strfmt.Registry) error {

	if err := validate.Required("period", "body", m.Period); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this azure express route circuit peering collector attribute v3 based on the context it is used
func (m *AzureExpressRouteCircuitPeeringCollectorAttributeV3) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *AzureExpressRouteCircuitPeeringCollectorAttributeV3) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureExpressRouteCircuitPeeringCollectorAttributeV3) UnmarshalBinary(b []byte) error {
	var res AzureExpressRouteCircuitPeeringCollectorAttributeV3
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
