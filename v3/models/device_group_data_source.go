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

// DeviceGroupDataSource device group data source
//
// swagger:model DeviceGroupDataSource
type DeviceGroupDataSource struct {

	// The Display Name of the DataSource
	// Read Only: true
	DataSourceDisplayName string `json:"dataSourceDisplayName,omitempty"`

	// The DataSource Group name
	// Read Only: true
	DataSourceGroupName string `json:"dataSourceGroupName,omitempty"`

	// The ID of the DataSource
	// Read Only: true
	DataSourceID int32 `json:"dataSourceId,omitempty"`

	// The Name of the DataSource
	// Read Only: true
	DataSourceName string `json:"dataSourceName,omitempty"`

	// The Type of the DataSource
	// Read Only: true
	DataSourceType string `json:"dataSourceType,omitempty"`

	// The ID of the Device Group for the DataSource
	// Read Only: true
	DeviceGroupID int32 `json:"deviceGroupId,omitempty"`

	// Boolean flag for enabling/disabling alerting for DataSource
	// Read Only: true
	DisableAlerting *bool `json:"disableAlerting,omitempty"`

	// Boolean flag for enabling/disabling monitoring of DataSource
	StopMonitoring bool `json:"stopMonitoring,omitempty"`
}

// Validate validates this device group data source
func (m *DeviceGroupDataSource) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this device group data source based on the context it is used
func (m *DeviceGroupDataSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDataSourceDisplayName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDataSourceGroupName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDataSourceID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDataSourceName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDataSourceType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeviceGroupID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisableAlerting(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceGroupDataSource) contextValidateDataSourceDisplayName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dataSourceDisplayName", "body", string(m.DataSourceDisplayName)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceGroupDataSource) contextValidateDataSourceGroupName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dataSourceGroupName", "body", string(m.DataSourceGroupName)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceGroupDataSource) contextValidateDataSourceID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dataSourceId", "body", int32(m.DataSourceID)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceGroupDataSource) contextValidateDataSourceName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dataSourceName", "body", string(m.DataSourceName)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceGroupDataSource) contextValidateDataSourceType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dataSourceType", "body", string(m.DataSourceType)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceGroupDataSource) contextValidateDeviceGroupID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "deviceGroupId", "body", int32(m.DeviceGroupID)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceGroupDataSource) contextValidateDisableAlerting(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "disableAlerting", "body", m.DisableAlerting); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceGroupDataSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceGroupDataSource) UnmarshalBinary(b []byte) error {
	var res DeviceGroupDataSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}