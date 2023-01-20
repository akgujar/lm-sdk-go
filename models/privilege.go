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

// Privilege privilege
//
// swagger:model Privilege
type Privilege struct {

	// The privilege object identifier
	// Example: 123
	// Required: true
	ObjectID *string `json:"objectId"`

	// The privilege object name
	// Read Only: true
	ObjectName string `json:"objectName,omitempty"`

	// The privilege object type. The values can be dashboard_group|dashboard|host_group|service_group|website_group|report_group|remoteSession|chat|setting|device_dashboard|help|logs|configNeedDeviceManagePermission|map|resourceMapTab|tracesManageTab
	// Example: dashboard group
	// Required: true
	ObjectType *string `json:"objectType"`

	// The privilege operation
	// Example: write
	// Required: true
	Operation *string `json:"operation"`

	// The highest privilege operation on its children operations
	// Read Only: true
	SubOperation string `json:"subOperation,omitempty"`
}

// Validate validates this privilege
func (m *Privilege) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateObjectID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObjectType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Privilege) validateObjectID(formats strfmt.Registry) error {

	if err := validate.Required("objectId", "body", m.ObjectID); err != nil {
		return err
	}

	return nil
}

func (m *Privilege) validateObjectType(formats strfmt.Registry) error {

	if err := validate.Required("objectType", "body", m.ObjectType); err != nil {
		return err
	}

	return nil
}

func (m *Privilege) validateOperation(formats strfmt.Registry) error {

	if err := validate.Required("operation", "body", m.Operation); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this privilege based on the context it is used
func (m *Privilege) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateObjectName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSubOperation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Privilege) contextValidateObjectName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "objectName", "body", string(m.ObjectName)); err != nil {
		return err
	}

	return nil
}

func (m *Privilege) contextValidateSubOperation(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "subOperation", "body", string(m.SubOperation)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Privilege) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Privilege) UnmarshalBinary(b []byte) error {
	var res Privilege
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
