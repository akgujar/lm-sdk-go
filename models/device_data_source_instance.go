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

// DeviceDataSourceInstance device data source instance
//
// swagger:model DeviceDataSourceInstance
type DeviceDataSourceInstance struct {

	// Any instance level auto properties assigned to the instance
	AutoProperties []*NameAndValue `json:"autoProperties,omitempty"`

	// The id of the collector the datasource instance is associated with
	// Read Only: true
	CollectorID int32 `json:"collectorId,omitempty"`

	// Any instance level properties assigned to the instance
	CustomProperties []*NameAndValue `json:"customProperties,omitempty"`

	// The id of the datasource definition that the instance represents
	// Read Only: true
	DataSourceID int32 `json:"dataSourceId,omitempty"`

	// The type of LogicModule, e.g. DS (datasource)
	// Read Only: true
	DataSourceType string `json:"dataSourceType,omitempty"`

	// The description of the datasource instance
	// Example: Ping Test
	Description string `json:"description,omitempty"`

	// The id of the unique device-datasource the instance is associated with
	// Read Only: true
	DeviceDataSourceID int32 `json:"deviceDataSourceId,omitempty"`

	// The display name of the device the datasource instance is associated with
	// Read Only: true
	DeviceDisplayName string `json:"deviceDisplayName,omitempty"`

	// The id of the device the datasource instance is associated with
	// Read Only: true
	DeviceID int32 `json:"deviceId,omitempty"`

	// Whether or not alerting is disabled for the instance
	// Example: true
	DisableAlerting bool `json:"disableAlerting,omitempty"`

	// The instance alias. This is the descriptive name of the instance, and should be unique for the device/datasource combination
	// Example: Ping
	// Required: true
	DisplayName *string `json:"displayName"`

	// The id of the instance group associated with the datasource instance
	// Example: 211
	GroupID int32 `json:"groupId,omitempty"`

	// The name of the instance group associated with the datasource instance
	// Read Only: true
	GroupName string `json:"groupName,omitempty"`

	// The Id of the datasource instance
	// Read Only: true
	ID int32 `json:"id,omitempty"`

	// Whether or not UNC Monitoring enabled for device
	// Example: true
	IsUNCInstance bool `json:"isUNCInstance,omitempty"`

	// Whether or not Active Discovery is enabled, and thus whether or not the instance description is editable
	// Example: true
	LockDescription bool `json:"lockDescription,omitempty"`

	// The name of the datasource instance, in the format of: datasourceName-instanceAlias
	// Read Only: true
	Name string `json:"name,omitempty"`

	// Whether or not monitoring is disabled for the instance
	// Example: true
	StopMonitoring bool `json:"stopMonitoring,omitempty"`

	// Any instance level system properties assigned to the instance
	SystemProperties []*NameAndValue `json:"systemProperties,omitempty"`

	// The variable part of the instance, used to query data from a device. For example, variable part of the SNMP OID tree. This value must be unique for the device/datasource combination, unless two-dimensional active discovery is used
	// Example: 1
	// Required: true
	WildValue *string `json:"wildValue"`

	// Only used for two dimensional active discovery. When used, during Active Discovery runs, the token ##WILDVALUE## is replaces with the value of ALIAS and the token ##WILDVALUE2## is replaced with the value of the second part alias. This value must be unique for the device/datasource/WILDVALUE combination
	// Example: 1
	WildValue2 string `json:"wildValue2,omitempty"`
}

// Validate validates this device data source instance
func (m *DeviceDataSourceInstance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAutoProperties(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomProperties(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisplayName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSystemProperties(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWildValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceDataSourceInstance) validateAutoProperties(formats strfmt.Registry) error {
	if swag.IsZero(m.AutoProperties) { // not required
		return nil
	}

	for i := 0; i < len(m.AutoProperties); i++ {
		if swag.IsZero(m.AutoProperties[i]) { // not required
			continue
		}

		if m.AutoProperties[i] != nil {
			if err := m.AutoProperties[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("autoProperties" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("autoProperties" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeviceDataSourceInstance) validateCustomProperties(formats strfmt.Registry) error {
	if swag.IsZero(m.CustomProperties) { // not required
		return nil
	}

	for i := 0; i < len(m.CustomProperties); i++ {
		if swag.IsZero(m.CustomProperties[i]) { // not required
			continue
		}

		if m.CustomProperties[i] != nil {
			if err := m.CustomProperties[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("customProperties" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("customProperties" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeviceDataSourceInstance) validateDisplayName(formats strfmt.Registry) error {

	if err := validate.Required("displayName", "body", m.DisplayName); err != nil {
		return err
	}

	return nil
}

func (m *DeviceDataSourceInstance) validateSystemProperties(formats strfmt.Registry) error {
	if swag.IsZero(m.SystemProperties) { // not required
		return nil
	}

	for i := 0; i < len(m.SystemProperties); i++ {
		if swag.IsZero(m.SystemProperties[i]) { // not required
			continue
		}

		if m.SystemProperties[i] != nil {
			if err := m.SystemProperties[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("systemProperties" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("systemProperties" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeviceDataSourceInstance) validateWildValue(formats strfmt.Registry) error {

	if err := validate.Required("wildValue", "body", m.WildValue); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this device data source instance based on the context it is used
func (m *DeviceDataSourceInstance) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAutoProperties(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCollectorID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCustomProperties(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDataSourceID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDataSourceType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeviceDataSourceID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeviceDisplayName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeviceID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGroupName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSystemProperties(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceDataSourceInstance) contextValidateAutoProperties(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AutoProperties); i++ {

		if m.AutoProperties[i] != nil {

			if swag.IsZero(m.AutoProperties[i]) { // not required
				return nil
			}

			if err := m.AutoProperties[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("autoProperties" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("autoProperties" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeviceDataSourceInstance) contextValidateCollectorID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "collectorId", "body", int32(m.CollectorID)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceDataSourceInstance) contextValidateCustomProperties(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CustomProperties); i++ {

		if m.CustomProperties[i] != nil {

			if swag.IsZero(m.CustomProperties[i]) { // not required
				return nil
			}

			if err := m.CustomProperties[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("customProperties" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("customProperties" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeviceDataSourceInstance) contextValidateDataSourceID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dataSourceId", "body", int32(m.DataSourceID)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceDataSourceInstance) contextValidateDataSourceType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dataSourceType", "body", string(m.DataSourceType)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceDataSourceInstance) contextValidateDeviceDataSourceID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "deviceDataSourceId", "body", int32(m.DeviceDataSourceID)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceDataSourceInstance) contextValidateDeviceDisplayName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "deviceDisplayName", "body", string(m.DeviceDisplayName)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceDataSourceInstance) contextValidateDeviceID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "deviceId", "body", int32(m.DeviceID)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceDataSourceInstance) contextValidateGroupName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "groupName", "body", string(m.GroupName)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceDataSourceInstance) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int32(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceDataSourceInstance) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceDataSourceInstance) contextValidateSystemProperties(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SystemProperties); i++ {

		if m.SystemProperties[i] != nil {

			if swag.IsZero(m.SystemProperties[i]) { // not required
				return nil
			}

			if err := m.SystemProperties[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("systemProperties" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("systemProperties" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceDataSourceInstance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceDataSourceInstance) UnmarshalBinary(b []byte) error {
	var res DeviceDataSourceInstance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
