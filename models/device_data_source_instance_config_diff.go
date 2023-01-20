// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DeviceDataSourceInstanceConfigDiff device data source instance config diff
//
// swagger:model DeviceDataSourceInstanceConfigDiff
type DeviceDataSourceInstanceConfigDiff struct {

	// Configuration content
	// Read Only: true
	Content string `json:"content,omitempty"`

	// Diff row number
	// Read Only: true
	RowNo int32 `json:"rowNo,omitempty"`

	// Diff type. The values can be : add|remove
	// Read Only: true
	Type string `json:"type,omitempty"`
}

// Validate validates this device data source instance config diff
func (m *DeviceDataSourceInstanceConfigDiff) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this device data source instance config diff based on the context it is used
func (m *DeviceDataSourceInstanceConfigDiff) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateContent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRowNo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceDataSourceInstanceConfigDiff) contextValidateContent(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "content", "body", string(m.Content)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceDataSourceInstanceConfigDiff) contextValidateRowNo(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "rowNo", "body", int32(m.RowNo)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceDataSourceInstanceConfigDiff) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "type", "body", string(m.Type)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceDataSourceInstanceConfigDiff) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceDataSourceInstanceConfigDiff) UnmarshalBinary(b []byte) error {
	var res DeviceDataSourceInstanceConfigDiff
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
