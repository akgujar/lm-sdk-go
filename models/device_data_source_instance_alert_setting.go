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

// DeviceDataSourceInstanceAlertSetting device data source instance alert setting
//
// swagger:model DeviceDataSourceInstanceAlertSetting
type DeviceDataSourceInstanceAlertSetting struct {

	// ad adv setting enabled
	AdAdvSettingEnabled bool `json:"adAdvSettingEnabled,omitempty"`

	// The polling interval of alert clear transition (0-60)
	// Example: 0
	AlertClearInterval int32 `json:"alertClearInterval,omitempty"`

	// The thresholds that should be associated with the datapoint. Note that you need to have a space between the operator and each threshold (e.g. > 1 2 3)
	// Example: \u003e= 70 90 95
	AlertExpr string `json:"alertExpr,omitempty"`

	// The note associated with the current alert threshold settings
	// Example: Router 1 Ping Check
	AlertExprNote string `json:"alertExprNote,omitempty"`

	// alert for no data (no alert-1, warning-2, error-3, critical-4)
	// Example: 1
	AlertForNoData int32 `json:"alertForNoData,omitempty"`

	// The polling interval of alert transition (0-60)
	// Example: 0
	AlertTransitionInterval int32 `json:"alertTransitionInterval,omitempty"`

	// The datapoint is effected alert disabled by which group
	// Read Only: true
	AlertingDisabledOn string `json:"alertingDisabledOn,omitempty"`

	// Collection Interval
	// Read Only: true
	CollectionInterval int64 `json:"collectionInterval,omitempty"`

	// critical ad adv setting
	CriticalAdAdvSetting string `json:"criticalAdAdvSetting,omitempty"`

	// The description of the datapoint the alert settings apply to
	// Read Only: true
	DataPointDescription string `json:"dataPointDescription,omitempty"`

	// The id of the Datapoint alert settings apply to
	// Read Only: true
	DataPointID int32 `json:"dataPointId,omitempty"`

	// The name of the datapoint the alert settings apply to
	// Read Only: true
	DataPointName string `json:"dataPointName,omitempty"`

	// The alias (name) of the DataSource instance the alert settings apply to
	// Read Only: true
	DataSourceInstanceAlias string `json:"dataSourceInstanceAlias,omitempty"`

	// The id of the DataSource instance alert settings apply to
	// Read Only: true
	DataSourceInstanceID int32 `json:"dataSourceInstanceId,omitempty"`

	// The full path of the device group
	// Read Only: true
	DeviceGroupFullPath string `json:"deviceGroupFullPath,omitempty"`

	// The ID of the device group
	// Read Only: true
	DeviceGroupID int32 `json:"deviceGroupId,omitempty"`

	// Whether or not alerting will be disabled for the datapoint
	// Example: true
	DisableAlerting bool `json:"disableAlerting,omitempty"`

	// The group full path lists who disable alert for this datapoint on devicegroup level
	// Read Only: true
	DisableDpAlertHostGroups string `json:"disableDpAlertHostGroups,omitempty"`

	// enable anomaly alert generation
	// Read Only: true
	EnableAnomalyAlertGeneration string `json:"enableAnomalyAlertGeneration,omitempty"`

	// enable anomaly alert suppression
	// Read Only: true
	EnableAnomalyAlertSuppression string `json:"enableAnomalyAlertSuppression,omitempty"`

	// error ad adv setting
	ErrorAdAdvSetting string `json:"errorAdAdvSetting,omitempty"`

	// The count that the alert must exist for this many poll cycles before the alert will be cleared
	// Example: 0
	// Read Only: true
	GlobalAlertClearTransitionInterval int32 `json:"globalAlertClearTransitionInterval,omitempty"`

	// The global alert expression for this datapoint
	// Read Only: true
	GlobalAlertExpr string `json:"globalAlertExpr,omitempty"`

	// The triggered alert level if we cannot collect data for this datapoint
	// Example: 1
	// Read Only: true
	GlobalAlertForNoData int32 `json:"globalAlertForNoData,omitempty"`

	// The count that the alert must exist for this many poll cycles before it will be triggered
	// Example: 0
	// Read Only: true
	GlobalAlertTransitionInterval int32 `json:"globalAlertTransitionInterval,omitempty"`

	// The global enable anomaly alert generation
	// Read Only: true
	GlobalEnableAnomalyAlertGeneration string `json:"globalEnableAnomalyAlertGeneration,omitempty"`

	// The global enable anomaly alert suppression
	// Read Only: true
	GlobalEnableAnomalyAlertSuppression string `json:"globalEnableAnomalyAlertSuppression,omitempty"`

	// The post processor parameters for complex DataPoints and global level configCheck threshold.
	GlobalPostProcessorParam string `json:"globalPostProcessorParam,omitempty"`

	// The id of this alert setting
	// Read Only: true
	ID int32 `json:"id,omitempty"`

	// Device group alert expression list base on the priority. The first is the highest priority and effected on this instance
	// Read Only: true
	ParentDeviceGroupAlertExprList []*DeviceGroupAlertThresholdInfo `json:"parentDeviceGroupAlertExprList,omitempty"`

	// Instance group alert expression list base on the priority. The first is the highest priority and effected on this instance
	// Read Only: true
	ParentInstanceGroupAlertExpr *InstanceGroupAlertThresholdInfo `json:"parentInstanceGroupAlertExpr,omitempty"`

	// Resource datasource alert expression list base on the priority. The first is the highest priority and effected on this instance
	// Read Only: true
	ParentResourceDataSourceAlertExpr *ResourceDataSourceAlertThresholdInfo `json:"parentResourceDataSourceAlertExpr,omitempty"`

	// The post processor parameter for complex DataPoint and instance level configCheck threshold.
	PostProcessorParam string `json:"postProcessorParam,omitempty"`

	// warn ad adv setting
	WarnAdAdvSetting string `json:"warnAdAdvSetting,omitempty"`
}

// Validate validates this device data source instance alert setting
func (m *DeviceDataSourceInstanceAlertSetting) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateParentDeviceGroupAlertExprList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParentInstanceGroupAlertExpr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParentResourceDataSourceAlertExpr(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceDataSourceInstanceAlertSetting) validateParentDeviceGroupAlertExprList(formats strfmt.Registry) error {
	if swag.IsZero(m.ParentDeviceGroupAlertExprList) { // not required
		return nil
	}

	for i := 0; i < len(m.ParentDeviceGroupAlertExprList); i++ {
		if swag.IsZero(m.ParentDeviceGroupAlertExprList[i]) { // not required
			continue
		}

		if m.ParentDeviceGroupAlertExprList[i] != nil {
			if err := m.ParentDeviceGroupAlertExprList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parentDeviceGroupAlertExprList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeviceDataSourceInstanceAlertSetting) validateParentInstanceGroupAlertExpr(formats strfmt.Registry) error {
	if swag.IsZero(m.ParentInstanceGroupAlertExpr) { // not required
		return nil
	}

	if m.ParentInstanceGroupAlertExpr != nil {
		if err := m.ParentInstanceGroupAlertExpr.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parentInstanceGroupAlertExpr")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceDataSourceInstanceAlertSetting) validateParentResourceDataSourceAlertExpr(formats strfmt.Registry) error {
	if swag.IsZero(m.ParentResourceDataSourceAlertExpr) { // not required
		return nil
	}

	if m.ParentResourceDataSourceAlertExpr != nil {
		if err := m.ParentResourceDataSourceAlertExpr.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parentResourceDataSourceAlertExpr")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this device data source instance alert setting based on the context it is used
func (m *DeviceDataSourceInstanceAlertSetting) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAlertingDisabledOn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCollectionInterval(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDataPointDescription(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDataPointID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDataPointName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDataSourceInstanceAlias(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDataSourceInstanceID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeviceGroupFullPath(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeviceGroupID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisableDpAlertHostGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEnableAnomalyAlertGeneration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEnableAnomalyAlertSuppression(ctx, formats); err != nil {
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

	if err := m.contextValidateGlobalEnableAnomalyAlertGeneration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGlobalEnableAnomalyAlertSuppression(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateParentDeviceGroupAlertExprList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateParentInstanceGroupAlertExpr(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateParentResourceDataSourceAlertExpr(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceDataSourceInstanceAlertSetting) contextValidateAlertingDisabledOn(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "alertingDisabledOn", "body", string(m.AlertingDisabledOn)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceDataSourceInstanceAlertSetting) contextValidateCollectionInterval(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "collectionInterval", "body", int64(m.CollectionInterval)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceDataSourceInstanceAlertSetting) contextValidateDataPointDescription(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dataPointDescription", "body", string(m.DataPointDescription)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceDataSourceInstanceAlertSetting) contextValidateDataPointID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dataPointId", "body", int32(m.DataPointID)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceDataSourceInstanceAlertSetting) contextValidateDataPointName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dataPointName", "body", string(m.DataPointName)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceDataSourceInstanceAlertSetting) contextValidateDataSourceInstanceAlias(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dataSourceInstanceAlias", "body", string(m.DataSourceInstanceAlias)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceDataSourceInstanceAlertSetting) contextValidateDataSourceInstanceID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dataSourceInstanceId", "body", int32(m.DataSourceInstanceID)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceDataSourceInstanceAlertSetting) contextValidateDeviceGroupFullPath(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "deviceGroupFullPath", "body", string(m.DeviceGroupFullPath)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceDataSourceInstanceAlertSetting) contextValidateDeviceGroupID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "deviceGroupId", "body", int32(m.DeviceGroupID)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceDataSourceInstanceAlertSetting) contextValidateDisableDpAlertHostGroups(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "disableDpAlertHostGroups", "body", string(m.DisableDpAlertHostGroups)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceDataSourceInstanceAlertSetting) contextValidateEnableAnomalyAlertGeneration(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "enableAnomalyAlertGeneration", "body", string(m.EnableAnomalyAlertGeneration)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceDataSourceInstanceAlertSetting) contextValidateEnableAnomalyAlertSuppression(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "enableAnomalyAlertSuppression", "body", string(m.EnableAnomalyAlertSuppression)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceDataSourceInstanceAlertSetting) contextValidateGlobalAlertClearTransitionInterval(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "globalAlertClearTransitionInterval", "body", int32(m.GlobalAlertClearTransitionInterval)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceDataSourceInstanceAlertSetting) contextValidateGlobalAlertExpr(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "globalAlertExpr", "body", string(m.GlobalAlertExpr)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceDataSourceInstanceAlertSetting) contextValidateGlobalAlertForNoData(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "globalAlertForNoData", "body", int32(m.GlobalAlertForNoData)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceDataSourceInstanceAlertSetting) contextValidateGlobalAlertTransitionInterval(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "globalAlertTransitionInterval", "body", int32(m.GlobalAlertTransitionInterval)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceDataSourceInstanceAlertSetting) contextValidateGlobalEnableAnomalyAlertGeneration(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "globalEnableAnomalyAlertGeneration", "body", string(m.GlobalEnableAnomalyAlertGeneration)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceDataSourceInstanceAlertSetting) contextValidateGlobalEnableAnomalyAlertSuppression(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "globalEnableAnomalyAlertSuppression", "body", string(m.GlobalEnableAnomalyAlertSuppression)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceDataSourceInstanceAlertSetting) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int32(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceDataSourceInstanceAlertSetting) contextValidateParentDeviceGroupAlertExprList(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "parentDeviceGroupAlertExprList", "body", []*DeviceGroupAlertThresholdInfo(m.ParentDeviceGroupAlertExprList)); err != nil {
		return err
	}

	for i := 0; i < len(m.ParentDeviceGroupAlertExprList); i++ {

		if m.ParentDeviceGroupAlertExprList[i] != nil {
			if err := m.ParentDeviceGroupAlertExprList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parentDeviceGroupAlertExprList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeviceDataSourceInstanceAlertSetting) contextValidateParentInstanceGroupAlertExpr(ctx context.Context, formats strfmt.Registry) error {

	if m.ParentInstanceGroupAlertExpr != nil {
		if err := m.ParentInstanceGroupAlertExpr.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parentInstanceGroupAlertExpr")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceDataSourceInstanceAlertSetting) contextValidateParentResourceDataSourceAlertExpr(ctx context.Context, formats strfmt.Registry) error {

	if m.ParentResourceDataSourceAlertExpr != nil {
		if err := m.ParentResourceDataSourceAlertExpr.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parentResourceDataSourceAlertExpr")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceDataSourceInstanceAlertSetting) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceDataSourceInstanceAlertSetting) UnmarshalBinary(b []byte) error {
	var res DeviceDataSourceInstanceAlertSetting
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
