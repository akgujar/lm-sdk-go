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

// DeviceGroupDataSourceDataPointConfig device group data source data point config
//
// swagger:model DeviceGroupDataSourceDataPointConfig
type DeviceGroupDataSourceDataPointConfig struct {

	// ad adv setting enabled
	AdAdvSettingEnabled bool `json:"adAdvSettingEnabled,omitempty"`

	// The count that the alert must exist for this many poll cycles before the alert will be cleared (0-60)
	// Example: 0
	AlertClearTransitionInterval int32 `json:"alertClearTransitionInterval,omitempty"`

	// alert expr
	// Required: true
	AlertExpr *string `json:"alertExpr"`

	// alert expr note
	AlertExprNote string `json:"alertExprNote,omitempty"`

	// The triggered alert level if we cannot collect data for this datapoint. The values can be 1-4 (1:no alert, 2:warn alert, 3:error alert, 4:critical alert)
	// Example: 1
	AlertForNoData int32 `json:"alertForNoData,omitempty"`

	// The count that the alert must exist for this many poll cycles before it will be triggered (0-60)
	// Example: 0
	AlertTransitionInterval int32 `json:"alertTransitionInterval,omitempty"`

	// Collection Interval
	// Read Only: true
	CollectionInterval int64 `json:"collectionInterval,omitempty"`

	// critical ad adv setting
	CriticalAdAdvSetting string `json:"criticalAdAdvSetting,omitempty"`

	// data point description
	// Read Only: true
	DataPointDescription string `json:"dataPointDescription,omitempty"`

	// data point Id
	// Required: true
	DataPointID *int32 `json:"dataPointId"`

	// data point name
	// Required: true
	DataPointName *string `json:"dataPointName"`

	// disable alerting
	DisableAlerting bool `json:"disableAlerting,omitempty"`

	// enable anomaly alert generation
	EnableAnomalyAlertGeneration string `json:"enableAnomalyAlertGeneration,omitempty"`

	// enable anomaly alert suppression
	EnableAnomalyAlertSuppression string `json:"enableAnomalyAlertSuppression,omitempty"`

	// error ad adv setting
	ErrorAdAdvSetting string `json:"errorAdAdvSetting,omitempty"`

	// The count that the alert must exist for this many poll cycles before the alert will be cleared
	// Example: 0
	// Read Only: true
	GlobalAlertClearTransitionInterval int32 `json:"globalAlertClearTransitionInterval,omitempty"`

	// global alert expr
	// Read Only: true
	GlobalAlertExpr string `json:"globalAlertExpr,omitempty"`

	// The triggered alert level if we cannot collect data for this datapoint. The values can be 1-4 (1:no alert, 2:warn alert, 3:error alert, 4:critical alert)
	// Example: 1
	// Read Only: true
	GlobalAlertForNoData int32 `json:"globalAlertForNoData,omitempty"`

	// The count that the alert must exist for this many poll cycles before it will be triggered
	// Example: 0
	// Read Only: true
	GlobalAlertTransitionInterval int32 `json:"globalAlertTransitionInterval,omitempty"`

	// global enable anomaly alert generation
	GlobalEnableAnomalyAlertGeneration string `json:"globalEnableAnomalyAlertGeneration,omitempty"`

	// global enable anomaly alert suppression
	GlobalEnableAnomalyAlertSuppression string `json:"globalEnableAnomalyAlertSuppression,omitempty"`

	// warn ad adv setting
	WarnAdAdvSetting string `json:"warnAdAdvSetting,omitempty"`
}

// Validate validates this device group data source data point config
func (m *DeviceGroupDataSourceDataPointConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlertExpr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDataPointID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDataPointName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceGroupDataSourceDataPointConfig) validateAlertExpr(formats strfmt.Registry) error {

	if err := validate.Required("alertExpr", "body", m.AlertExpr); err != nil {
		return err
	}

	return nil
}

func (m *DeviceGroupDataSourceDataPointConfig) validateDataPointID(formats strfmt.Registry) error {

	if err := validate.Required("dataPointId", "body", m.DataPointID); err != nil {
		return err
	}

	return nil
}

func (m *DeviceGroupDataSourceDataPointConfig) validateDataPointName(formats strfmt.Registry) error {

	if err := validate.Required("dataPointName", "body", m.DataPointName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this device group data source data point config based on the context it is used
func (m *DeviceGroupDataSourceDataPointConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCollectionInterval(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDataPointDescription(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGlobalAlertClearTransitionInterval(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGlobalAlertExpr(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGlobalAlertForNoData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGlobalAlertTransitionInterval(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceGroupDataSourceDataPointConfig) contextValidateCollectionInterval(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "collectionInterval", "body", int64(m.CollectionInterval)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceGroupDataSourceDataPointConfig) contextValidateDataPointDescription(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dataPointDescription", "body", string(m.DataPointDescription)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceGroupDataSourceDataPointConfig) contextValidateGlobalAlertClearTransitionInterval(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "globalAlertClearTransitionInterval", "body", int32(m.GlobalAlertClearTransitionInterval)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceGroupDataSourceDataPointConfig) contextValidateGlobalAlertExpr(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "globalAlertExpr", "body", string(m.GlobalAlertExpr)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceGroupDataSourceDataPointConfig) contextValidateGlobalAlertForNoData(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "globalAlertForNoData", "body", int32(m.GlobalAlertForNoData)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceGroupDataSourceDataPointConfig) contextValidateGlobalAlertTransitionInterval(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "globalAlertTransitionInterval", "body", int32(m.GlobalAlertTransitionInterval)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceGroupDataSourceDataPointConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceGroupDataSourceDataPointConfig) UnmarshalBinary(b []byte) error {
	var res DeviceGroupDataSourceDataPointConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
