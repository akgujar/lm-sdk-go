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

// WebsiteLocation website location
//
// swagger:model WebsiteLocation
type WebsiteLocation struct {

	// This field only for the SiteMonitor Groups, does not include Internal Service Groups
	// Example: true
	All interface{} `json:"all,omitempty"`

	// The Internal Service Groups Ids
	CollectorIds []int32 `json:"collectorIds,omitempty"`

	// The collector info of the services
	Collectors []*WebsiteCollectorInfo `json:"collectors,omitempty"`

	// The SiteMonitor Groups Ids
	// Example: [1, 2, 4, 3, 5, 6]
	SmgIds []int32 `json:"smgIds,omitempty"`
}

// Validate validates this website location
func (m *WebsiteLocation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCollectors(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebsiteLocation) validateCollectors(formats strfmt.Registry) error {
	if swag.IsZero(m.Collectors) { // not required
		return nil
	}

	for i := 0; i < len(m.Collectors); i++ {
		if swag.IsZero(m.Collectors[i]) { // not required
			continue
		}

		if m.Collectors[i] != nil {
			if err := m.Collectors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("collectors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("collectors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this website location based on the context it is used
func (m *WebsiteLocation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCollectors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebsiteLocation) contextValidateCollectors(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Collectors); i++ {

		if m.Collectors[i] != nil {

			if swag.IsZero(m.Collectors[i]) { // not required
				return nil
			}

			if err := m.Collectors[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("collectors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("collectors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *WebsiteLocation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebsiteLocation) UnmarshalBinary(b []byte) error {
	var res WebsiteLocation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
