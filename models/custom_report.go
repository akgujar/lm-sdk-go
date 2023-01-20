// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CustomReport custom report
//
// swagger:model CustomReport
type CustomReport struct {
	customReportTypeIdField int32

	customReportTypeNameField string

	deliveryField string

	descriptionField string

	enableViewAsOtherUserField *bool

	formatField string

	groupIdField int32

	idField int32

	lastGenerateOnField int64

	lastGeneratePagesField int32

	lastGenerateSizeField int64

	lastmodifyUserIdField int32

	lastmodifyUserNameField string

	nameField *string

	recipientsField []*ReportRecipient

	reportLinkExpireField string

	reportLinkNumField int32

	scheduleField string

	scheduleTimezoneField string

	userPermissionField string

	// date range
	DateRange string `json:"dateRange,omitempty"`

	// macros
	Macros []*Macro `json:"macros,omitempty"`
}

// CustomReportTypeID gets the custom report type Id of this subtype
func (m *CustomReport) CustomReportTypeID() int32 {
	return m.customReportTypeIdField
}

// SetCustomReportTypeID sets the custom report type Id of this subtype
func (m *CustomReport) SetCustomReportTypeID(val int32) {
	m.customReportTypeIdField = val
}

// CustomReportTypeName gets the custom report type name of this subtype
func (m *CustomReport) CustomReportTypeName() string {
	return m.customReportTypeNameField
}

// SetCustomReportTypeName sets the custom report type name of this subtype
func (m *CustomReport) SetCustomReportTypeName(val string) {
	m.customReportTypeNameField = val
}

// Delivery gets the delivery of this subtype
func (m *CustomReport) Delivery() string {
	return m.deliveryField
}

// SetDelivery sets the delivery of this subtype
func (m *CustomReport) SetDelivery(val string) {
	m.deliveryField = val
}

// Description gets the description of this subtype
func (m *CustomReport) Description() string {
	return m.descriptionField
}

// SetDescription sets the description of this subtype
func (m *CustomReport) SetDescription(val string) {
	m.descriptionField = val
}

// EnableViewAsOtherUser gets the enable view as other user of this subtype
func (m *CustomReport) EnableViewAsOtherUser() *bool {
	return m.enableViewAsOtherUserField
}

// SetEnableViewAsOtherUser sets the enable view as other user of this subtype
func (m *CustomReport) SetEnableViewAsOtherUser(val *bool) {
	m.enableViewAsOtherUserField = val
}

// Format gets the format of this subtype
func (m *CustomReport) Format() string {
	return m.formatField
}

// SetFormat sets the format of this subtype
func (m *CustomReport) SetFormat(val string) {
	m.formatField = val
}

// GroupID gets the group Id of this subtype
func (m *CustomReport) GroupID() int32 {
	return m.groupIdField
}

// SetGroupID sets the group Id of this subtype
func (m *CustomReport) SetGroupID(val int32) {
	m.groupIdField = val
}

// ID gets the id of this subtype
func (m *CustomReport) ID() int32 {
	return m.idField
}

// SetID sets the id of this subtype
func (m *CustomReport) SetID(val int32) {
	m.idField = val
}

// LastGenerateOn gets the last generate on of this subtype
func (m *CustomReport) LastGenerateOn() int64 {
	return m.lastGenerateOnField
}

// SetLastGenerateOn sets the last generate on of this subtype
func (m *CustomReport) SetLastGenerateOn(val int64) {
	m.lastGenerateOnField = val
}

// LastGeneratePages gets the last generate pages of this subtype
func (m *CustomReport) LastGeneratePages() int32 {
	return m.lastGeneratePagesField
}

// SetLastGeneratePages sets the last generate pages of this subtype
func (m *CustomReport) SetLastGeneratePages(val int32) {
	m.lastGeneratePagesField = val
}

// LastGenerateSize gets the last generate size of this subtype
func (m *CustomReport) LastGenerateSize() int64 {
	return m.lastGenerateSizeField
}

// SetLastGenerateSize sets the last generate size of this subtype
func (m *CustomReport) SetLastGenerateSize(val int64) {
	m.lastGenerateSizeField = val
}

// LastmodifyUserID gets the lastmodify user Id of this subtype
func (m *CustomReport) LastmodifyUserID() int32 {
	return m.lastmodifyUserIdField
}

// SetLastmodifyUserID sets the lastmodify user Id of this subtype
func (m *CustomReport) SetLastmodifyUserID(val int32) {
	m.lastmodifyUserIdField = val
}

// LastmodifyUserName gets the lastmodify user name of this subtype
func (m *CustomReport) LastmodifyUserName() string {
	return m.lastmodifyUserNameField
}

// SetLastmodifyUserName sets the lastmodify user name of this subtype
func (m *CustomReport) SetLastmodifyUserName(val string) {
	m.lastmodifyUserNameField = val
}

// Name gets the name of this subtype
func (m *CustomReport) Name() *string {
	return m.nameField
}

// SetName sets the name of this subtype
func (m *CustomReport) SetName(val *string) {
	m.nameField = val
}

// Recipients gets the recipients of this subtype
func (m *CustomReport) Recipients() []*ReportRecipient {
	return m.recipientsField
}

// SetRecipients sets the recipients of this subtype
func (m *CustomReport) SetRecipients(val []*ReportRecipient) {
	m.recipientsField = val
}

// ReportLinkExpire gets the report link expire of this subtype
func (m *CustomReport) ReportLinkExpire() string {
	return m.reportLinkExpireField
}

// SetReportLinkExpire sets the report link expire of this subtype
func (m *CustomReport) SetReportLinkExpire(val string) {
	m.reportLinkExpireField = val
}

// ReportLinkNum gets the report link num of this subtype
func (m *CustomReport) ReportLinkNum() int32 {
	return m.reportLinkNumField
}

// SetReportLinkNum sets the report link num of this subtype
func (m *CustomReport) SetReportLinkNum(val int32) {
	m.reportLinkNumField = val
}

// Schedule gets the schedule of this subtype
func (m *CustomReport) Schedule() string {
	return m.scheduleField
}

// SetSchedule sets the schedule of this subtype
func (m *CustomReport) SetSchedule(val string) {
	m.scheduleField = val
}

// ScheduleTimezone gets the schedule timezone of this subtype
func (m *CustomReport) ScheduleTimezone() string {
	return m.scheduleTimezoneField
}

// SetScheduleTimezone sets the schedule timezone of this subtype
func (m *CustomReport) SetScheduleTimezone(val string) {
	m.scheduleTimezoneField = val
}

// Type gets the type of this subtype
func (m *CustomReport) Type() string {
	return "Word template"
}

// SetType sets the type of this subtype
func (m *CustomReport) SetType(val string) {
}

// UserPermission gets the user permission of this subtype
func (m *CustomReport) UserPermission() string {
	return m.userPermissionField
}

// SetUserPermission sets the user permission of this subtype
func (m *CustomReport) SetUserPermission(val string) {
	m.userPermissionField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *CustomReport) UnmarshalJSON(raw []byte) error {
	var data struct {

		// date range
		DateRange string `json:"dateRange,omitempty"`

		// macros
		Macros []*Macro `json:"macros,omitempty"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		CustomReportTypeID int32 `json:"customReportTypeId,omitempty"`

		CustomReportTypeName string `json:"customReportTypeName,omitempty"`

		Delivery string `json:"delivery,omitempty"`

		Description string `json:"description,omitempty"`

		EnableViewAsOtherUser *bool `json:"enableViewAsOtherUser,omitempty"`

		Format string `json:"format,omitempty"`

		GroupID int32 `json:"groupId,omitempty"`

		ID int32 `json:"id,omitempty"`

		LastGenerateOn int64 `json:"lastGenerateOn,omitempty"`

		LastGeneratePages int32 `json:"lastGeneratePages,omitempty"`

		LastGenerateSize int64 `json:"lastGenerateSize,omitempty"`

		LastmodifyUserID int32 `json:"lastmodifyUserId,omitempty"`

		LastmodifyUserName string `json:"lastmodifyUserName,omitempty"`

		Name *string `json:"name"`

		Recipients []*ReportRecipient `json:"recipients,omitempty"`

		ReportLinkExpire string `json:"reportLinkExpire,omitempty"`

		ReportLinkNum int32 `json:"reportLinkNum,omitempty"`

		Schedule string `json:"schedule,omitempty"`

		ScheduleTimezone string `json:"scheduleTimezone,omitempty"`

		Type string `json:"type"`

		UserPermission string `json:"userPermission,omitempty"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result CustomReport

	result.customReportTypeIdField = base.CustomReportTypeID

	result.customReportTypeNameField = base.CustomReportTypeName

	result.deliveryField = base.Delivery

	result.descriptionField = base.Description

	result.enableViewAsOtherUserField = base.EnableViewAsOtherUser

	result.formatField = base.Format

	result.groupIdField = base.GroupID

	result.idField = base.ID

	result.lastGenerateOnField = base.LastGenerateOn

	result.lastGeneratePagesField = base.LastGeneratePages

	result.lastGenerateSizeField = base.LastGenerateSize

	result.lastmodifyUserIdField = base.LastmodifyUserID

	result.lastmodifyUserNameField = base.LastmodifyUserName

	result.nameField = base.Name

	result.recipientsField = base.Recipients

	result.reportLinkExpireField = base.ReportLinkExpire

	result.reportLinkNumField = base.ReportLinkNum

	result.scheduleField = base.Schedule

	result.scheduleTimezoneField = base.ScheduleTimezone

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}
	result.userPermissionField = base.UserPermission

	result.DateRange = data.DateRange
	result.Macros = data.Macros

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m CustomReport) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// date range
		DateRange string `json:"dateRange,omitempty"`

		// macros
		Macros []*Macro `json:"macros,omitempty"`
	}{

		DateRange: m.DateRange,

		Macros: m.Macros,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		CustomReportTypeID int32 `json:"customReportTypeId,omitempty"`

		CustomReportTypeName string `json:"customReportTypeName,omitempty"`

		Delivery string `json:"delivery,omitempty"`

		Description string `json:"description,omitempty"`

		EnableViewAsOtherUser *bool `json:"enableViewAsOtherUser,omitempty"`

		Format string `json:"format,omitempty"`

		GroupID int32 `json:"groupId,omitempty"`

		ID int32 `json:"id,omitempty"`

		LastGenerateOn int64 `json:"lastGenerateOn,omitempty"`

		LastGeneratePages int32 `json:"lastGeneratePages,omitempty"`

		LastGenerateSize int64 `json:"lastGenerateSize,omitempty"`

		LastmodifyUserID int32 `json:"lastmodifyUserId,omitempty"`

		LastmodifyUserName string `json:"lastmodifyUserName,omitempty"`

		Name *string `json:"name"`

		Recipients []*ReportRecipient `json:"recipients,omitempty"`

		ReportLinkExpire string `json:"reportLinkExpire,omitempty"`

		ReportLinkNum int32 `json:"reportLinkNum,omitempty"`

		Schedule string `json:"schedule,omitempty"`

		ScheduleTimezone string `json:"scheduleTimezone,omitempty"`

		Type string `json:"type"`

		UserPermission string `json:"userPermission,omitempty"`
	}{

		CustomReportTypeID: m.CustomReportTypeID(),

		CustomReportTypeName: m.CustomReportTypeName(),

		Delivery: m.Delivery(),

		Description: m.Description(),

		EnableViewAsOtherUser: m.EnableViewAsOtherUser(),

		Format: m.Format(),

		GroupID: m.GroupID(),

		ID: m.ID(),

		LastGenerateOn: m.LastGenerateOn(),

		LastGeneratePages: m.LastGeneratePages(),

		LastGenerateSize: m.LastGenerateSize(),

		LastmodifyUserID: m.LastmodifyUserID(),

		LastmodifyUserName: m.LastmodifyUserName(),

		Name: m.Name(),

		Recipients: m.Recipients(),

		ReportLinkExpire: m.ReportLinkExpire(),

		ReportLinkNum: m.ReportLinkNum(),

		Schedule: m.Schedule(),

		ScheduleTimezone: m.ScheduleTimezone(),

		Type: m.Type(),

		UserPermission: m.UserPermission(),
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this custom report
func (m *CustomReport) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecipients(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMacros(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustomReport) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name()); err != nil {
		return err
	}

	return nil
}

func (m *CustomReport) validateRecipients(formats strfmt.Registry) error {

	if swag.IsZero(m.Recipients()) { // not required
		return nil
	}

	for i := 0; i < len(m.Recipients()); i++ {
		if swag.IsZero(m.recipientsField[i]) { // not required
			continue
		}

		if m.recipientsField[i] != nil {
			if err := m.recipientsField[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("recipients" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CustomReport) validateMacros(formats strfmt.Registry) error {

	if swag.IsZero(m.Macros) { // not required
		return nil
	}

	for i := 0; i < len(m.Macros); i++ {
		if swag.IsZero(m.Macros[i]) { // not required
			continue
		}

		if m.Macros[i] != nil {
			if err := m.Macros[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("macros" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this custom report based on the context it is used
func (m *CustomReport) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCustomReportTypeID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCustomReportTypeName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEnableViewAsOtherUser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastGenerateOn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastGeneratePages(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastGenerateSize(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastmodifyUserID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastmodifyUserName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRecipients(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReportLinkNum(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUserPermission(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMacros(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustomReport) contextValidateCustomReportTypeID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "customReportTypeId", "body", int32(m.CustomReportTypeID())); err != nil {
		return err
	}

	return nil
}

func (m *CustomReport) contextValidateCustomReportTypeName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "customReportTypeName", "body", string(m.CustomReportTypeName())); err != nil {
		return err
	}

	return nil
}

func (m *CustomReport) contextValidateEnableViewAsOtherUser(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "enableViewAsOtherUser", "body", m.EnableViewAsOtherUser()); err != nil {
		return err
	}

	return nil
}

func (m *CustomReport) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int32(m.ID())); err != nil {
		return err
	}

	return nil
}

func (m *CustomReport) contextValidateLastGenerateOn(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "lastGenerateOn", "body", int64(m.LastGenerateOn())); err != nil {
		return err
	}

	return nil
}

func (m *CustomReport) contextValidateLastGeneratePages(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "lastGeneratePages", "body", int32(m.LastGeneratePages())); err != nil {
		return err
	}

	return nil
}

func (m *CustomReport) contextValidateLastGenerateSize(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "lastGenerateSize", "body", int64(m.LastGenerateSize())); err != nil {
		return err
	}

	return nil
}

func (m *CustomReport) contextValidateLastmodifyUserID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "lastmodifyUserId", "body", int32(m.LastmodifyUserID())); err != nil {
		return err
	}

	return nil
}

func (m *CustomReport) contextValidateLastmodifyUserName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "lastmodifyUserName", "body", string(m.LastmodifyUserName())); err != nil {
		return err
	}

	return nil
}

func (m *CustomReport) contextValidateRecipients(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Recipients()); i++ {

		if m.recipientsField[i] != nil {
			if err := m.recipientsField[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("recipients" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CustomReport) contextValidateReportLinkNum(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "reportLinkNum", "body", int32(m.ReportLinkNum())); err != nil {
		return err
	}

	return nil
}

func (m *CustomReport) contextValidateUserPermission(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "userPermission", "body", string(m.UserPermission())); err != nil {
		return err
	}

	return nil
}

func (m *CustomReport) contextValidateMacros(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Macros); i++ {

		if m.Macros[i] != nil {
			if err := m.Macros[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("macros" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CustomReport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomReport) UnmarshalBinary(b []byte) error {
	var res CustomReport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
