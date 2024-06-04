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

// Role role
//
// swagger:model Role
type Role struct {

	// Whether Two-Factor Authentication should be required for the entire account
	// Read Only: true
	AcctRequireTwoFA *bool `json:"acctRequireTwoFA,omitempty"`

	// The count of the users which are belongs to the role
	// Read Only: true
	AssociatedUserCount int32 `json:"associatedUserCount,omitempty"`

	// The label for the custom help URL as it will appear in the 'Help & Support' dropdown menu
	// Example: Internal Support Resources
	CustomHelpLabel string `json:"customHelpLabel,omitempty"`

	// The URL that should be added to the 'Help & Support' dropdown menu
	// Example: https://logicmonitor.com/support
	CustomHelpURL string `json:"customHelpURL,omitempty"`

	// The description of the role
	// Example: Administrator can do everything, including security-sensitive actions
	Description string `json:"description,omitempty"`

	// Whether Remote Session should be enabled at the account level
	// Read Only: true
	EnableRemoteSessionInCompanyLevel *bool `json:"enableRemoteSessionInCompanyLevel,omitempty"`

	// The Id of the role
	// Read Only: true
	ID int32 `json:"id,omitempty"`

	// The name of the role
	// Example: administrator
	// Required: true
	Name *string `json:"name"`

	// The account privileges associated with the role. Privileges can be added to a role for each area of your account
	// Required: true
	Privileges []*Privilege `json:"privileges"`

	// Whether or not users assigned this role should be required to acknowledge the EULA (end user license agreement)
	// Example: true
	RequireEULA bool `json:"requireEULA,omitempty"`

	// The group Id of the role
	// Example: 3
	RoleGroupID int32 `json:"roleGroupId,omitempty"`

	// Whether Two-Factor Authentication should be required for this role
	// Example: true
	TwoFARequired bool `json:"twoFARequired,omitempty"`

	// The permission of current role with the admin. The values can be write|read|none
	// Example: read
	// Read Only: true
	UserPermission string `json:"userPermission,omitempty"`
}

// Validate validates this role
func (m *Role) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivileges(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Role) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Role) validatePrivileges(formats strfmt.Registry) error {

	if err := validate.Required("privileges", "body", m.Privileges); err != nil {
		return err
	}

	for i := 0; i < len(m.Privileges); i++ {
		if swag.IsZero(m.Privileges[i]) { // not required
			continue
		}

		if m.Privileges[i] != nil {
			if err := m.Privileges[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("privileges" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("privileges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this role based on the context it is used
func (m *Role) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAcctRequireTwoFA(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAssociatedUserCount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEnableRemoteSessionInCompanyLevel(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrivileges(ctx, formats); err != nil {
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

func (m *Role) contextValidateAcctRequireTwoFA(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "acctRequireTwoFA", "body", m.AcctRequireTwoFA); err != nil {
		return err
	}

	return nil
}

func (m *Role) contextValidateAssociatedUserCount(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "associatedUserCount", "body", int32(m.AssociatedUserCount)); err != nil {
		return err
	}

	return nil
}

func (m *Role) contextValidateEnableRemoteSessionInCompanyLevel(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "enableRemoteSessionInCompanyLevel", "body", m.EnableRemoteSessionInCompanyLevel); err != nil {
		return err
	}

	return nil
}

func (m *Role) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int32(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *Role) contextValidatePrivileges(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Privileges); i++ {

		if m.Privileges[i] != nil {

			if swag.IsZero(m.Privileges[i]) { // not required
				return nil
			}

			if err := m.Privileges[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("privileges" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("privileges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Role) contextValidateUserPermission(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "userPermission", "body", string(m.UserPermission)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Role) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Role) UnmarshalBinary(b []byte) error {
	var res Role
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
