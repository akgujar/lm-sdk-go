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

// DeviceGroupData device group data
//
// swagger:model DeviceGroupData
type DeviceGroupData struct {

	// The Applies to custom query for this group (only for dynamic groups)
	// Read Only: true
	AppliesTo string `json:"appliesTo,omitempty"`

	// The number of instances in each AWS region (only applies to AWS groups)
	// Read Only: true
	AwsRegionsInfo string `json:"awsRegionsInfo,omitempty"`

	// The number of instances in each Azure region (only applies to Azure groups)
	// Read Only: true
	AzureRegionsInfo string `json:"azureRegionsInfo,omitempty"`

	// The description of the device group
	// Read Only: true
	Description string `json:"description,omitempty"`

	// The full path of the device group (i.e. if the group 'Dev' is under a parent group named 'Production', the fullPath would be 'Production/Dev'
	// Read Only: true
	FullPath string `json:"fullPath,omitempty"`

	// gcp regions info
	// Read Only: true
	GcpRegionsInfo string `json:"gcpRegionsInfo,omitempty"`

	// The type of device group: normal and dynamic device groups will have groupType=Normal, and AWS groups will have a groupType value of AWS/SERVICE (e.g. AWS/S3)
	// Read Only: true
	GroupType string `json:"groupType,omitempty"`

	// The id of the device group
	// Read Only: true
	ID int32 `json:"id,omitempty"`

	// The name of the device group
	// Read Only: true
	Name string `json:"name,omitempty"`

	// The number of AWS and normal devices that belong only to this device group (doesn't include devices in sub-groups)
	// Read Only: true
	NumOfDirectDevices int64 `json:"numOfDirectDevices,omitempty"`

	// The number of sub-groups that belong only to this device group (doesn't include groups under sub-groups)
	// Read Only: true
	NumOfDirectSubGroups int64 `json:"numOfDirectSubGroups,omitempty"`

	// The number of total devices, including both AWS and normal devices, that belong to this device group (includes normal devices in sub groups)
	// Read Only: true
	NumOfHosts int32 `json:"numOfHosts,omitempty"`

	// The role privilege operations for the device group that are granted to the user that made this API request
	// Read Only: true
	RolePrivileges []string `json:"rolePrivileges,omitempty"`

	// The permissions for the device group that are granted to the user that made this API request
	// Read Only: true
	UserPermission string `json:"userPermission,omitempty"`
}

// Validate validates this device group data
func (m *DeviceGroupData) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this device group data based on the context it is used
func (m *DeviceGroupData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAppliesTo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAwsRegionsInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAzureRegionsInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDescription(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFullPath(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGcpRegionsInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGroupType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNumOfDirectDevices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNumOfDirectSubGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNumOfHosts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRolePrivileges(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUserPermission(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceGroupData) contextValidateAppliesTo(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "appliesTo", "body", string(m.AppliesTo)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceGroupData) contextValidateAwsRegionsInfo(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "awsRegionsInfo", "body", string(m.AwsRegionsInfo)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceGroupData) contextValidateAzureRegionsInfo(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "azureRegionsInfo", "body", string(m.AzureRegionsInfo)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceGroupData) contextValidateDescription(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "description", "body", string(m.Description)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceGroupData) contextValidateFullPath(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "fullPath", "body", string(m.FullPath)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceGroupData) contextValidateGcpRegionsInfo(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "gcpRegionsInfo", "body", string(m.GcpRegionsInfo)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceGroupData) contextValidateGroupType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "groupType", "body", string(m.GroupType)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceGroupData) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int32(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceGroupData) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceGroupData) contextValidateNumOfDirectDevices(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "numOfDirectDevices", "body", int64(m.NumOfDirectDevices)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceGroupData) contextValidateNumOfDirectSubGroups(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "numOfDirectSubGroups", "body", int64(m.NumOfDirectSubGroups)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceGroupData) contextValidateNumOfHosts(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "numOfHosts", "body", int32(m.NumOfHosts)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceGroupData) contextValidateRolePrivileges(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "rolePrivileges", "body", []string(m.RolePrivileges)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceGroupData) contextValidateUserPermission(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "userPermission", "body", string(m.UserPermission)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceGroupData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceGroupData) UnmarshalBinary(b []byte) error {
	var res DeviceGroupData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
