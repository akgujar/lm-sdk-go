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

// AlertThresholdReport alert threshold report
//
// swagger:model AlertThresholdReport
type AlertThresholdReport struct {
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

	// The columns that will be displayed in the report. You should specify the columns in the order in which you'd like them to be displayed
	Columns []*DynamicColumn `json:"columns,omitempty"`

	// The datapoint to be included in the report. Glob expressions supported
	DataPoint string `json:"dataPoint,omitempty"`

	// The name of the datasource instance to be included in the report, where the syntax should be dataSourceDisplayName-InstanceName. If you're referencing a single instance datasource, you can just specify dataSourceDisplayName. Glob expressions supported
	DataSourceInstanceName string `json:"dataSourceInstanceName,omitempty"`

	// The display name of the device(s) to be included in the report. Glob expressions supported
	DeviceDisplayName string `json:"deviceDisplayName,omitempty"`

	// true: only variations from the global thresholds will be displayed
	// false: all thresholds will be displayed, including global thresholds an custom group and instance level thresholds
	// the default value is true
	ExcludeGlobal interface{} `json:"excludeGlobal,omitempty"`

	// The full path of the group whose member devices you are going to include in the report. Glob expressions supported
	GroupFullPath string `json:"groupFullPath,omitempty"`

	// host | datasource | datapoint
	// host: displayed thresholds will be sorted by device
	// datasource: displayed thresholds will be sorted by datasource instance
	// datapoint: displayed thresholds will be sorted by datapoint (metric)
	SortedBy string `json:"sortedBy,omitempty"`
}

// CustomReportTypeID gets the custom report type Id of this subtype
func (m *AlertThresholdReport) CustomReportTypeID() int32 {
	return m.customReportTypeIdField
}

// SetCustomReportTypeID sets the custom report type Id of this subtype
func (m *AlertThresholdReport) SetCustomReportTypeID(val int32) {
	m.customReportTypeIdField = val
}

// CustomReportTypeName gets the custom report type name of this subtype
func (m *AlertThresholdReport) CustomReportTypeName() string {
	return m.customReportTypeNameField
}

// SetCustomReportTypeName sets the custom report type name of this subtype
func (m *AlertThresholdReport) SetCustomReportTypeName(val string) {
	m.customReportTypeNameField = val
}

// Delivery gets the delivery of this subtype
func (m *AlertThresholdReport) Delivery() string {
	return m.deliveryField
}

// SetDelivery sets the delivery of this subtype
func (m *AlertThresholdReport) SetDelivery(val string) {
	m.deliveryField = val
}

// Description gets the description of this subtype
func (m *AlertThresholdReport) Description() string {
	return m.descriptionField
}

// SetDescription sets the description of this subtype
func (m *AlertThresholdReport) SetDescription(val string) {
	m.descriptionField = val
}

// EnableViewAsOtherUser gets the enable view as other user of this subtype
func (m *AlertThresholdReport) EnableViewAsOtherUser() *bool {
	return m.enableViewAsOtherUserField
}

// SetEnableViewAsOtherUser sets the enable view as other user of this subtype
func (m *AlertThresholdReport) SetEnableViewAsOtherUser(val *bool) {
	m.enableViewAsOtherUserField = val
}

// Format gets the format of this subtype
func (m *AlertThresholdReport) Format() string {
	return m.formatField
}

// SetFormat sets the format of this subtype
func (m *AlertThresholdReport) SetFormat(val string) {
	m.formatField = val
}

// GroupID gets the group Id of this subtype
func (m *AlertThresholdReport) GroupID() int32 {
	return m.groupIdField
}

// SetGroupID sets the group Id of this subtype
func (m *AlertThresholdReport) SetGroupID(val int32) {
	m.groupIdField = val
}

// ID gets the id of this subtype
func (m *AlertThresholdReport) ID() int32 {
	return m.idField
}

// SetID sets the id of this subtype
func (m *AlertThresholdReport) SetID(val int32) {
	m.idField = val
}

// LastGenerateOn gets the last generate on of this subtype
func (m *AlertThresholdReport) LastGenerateOn() int64 {
	return m.lastGenerateOnField
}

// SetLastGenerateOn sets the last generate on of this subtype
func (m *AlertThresholdReport) SetLastGenerateOn(val int64) {
	m.lastGenerateOnField = val
}

// LastGeneratePages gets the last generate pages of this subtype
func (m *AlertThresholdReport) LastGeneratePages() int32 {
	return m.lastGeneratePagesField
}

// SetLastGeneratePages sets the last generate pages of this subtype
func (m *AlertThresholdReport) SetLastGeneratePages(val int32) {
	m.lastGeneratePagesField = val
}

// LastGenerateSize gets the last generate size of this subtype
func (m *AlertThresholdReport) LastGenerateSize() int64 {
	return m.lastGenerateSizeField
}

// SetLastGenerateSize sets the last generate size of this subtype
func (m *AlertThresholdReport) SetLastGenerateSize(val int64) {
	m.lastGenerateSizeField = val
}

// LastmodifyUserID gets the lastmodify user Id of this subtype
func (m *AlertThresholdReport) LastmodifyUserID() int32 {
	return m.lastmodifyUserIdField
}

// SetLastmodifyUserID sets the lastmodify user Id of this subtype
func (m *AlertThresholdReport) SetLastmodifyUserID(val int32) {
	m.lastmodifyUserIdField = val
}

// LastmodifyUserName gets the lastmodify user name of this subtype
func (m *AlertThresholdReport) LastmodifyUserName() string {
	return m.lastmodifyUserNameField
}

// SetLastmodifyUserName sets the lastmodify user name of this subtype
func (m *AlertThresholdReport) SetLastmodifyUserName(val string) {
	m.lastmodifyUserNameField = val
}

// Name gets the name of this subtype
func (m *AlertThresholdReport) Name() *string {
	return m.nameField
}

// SetName sets the name of this subtype
func (m *AlertThresholdReport) SetName(val *string) {
	m.nameField = val
}

// Recipients gets the recipients of this subtype
func (m *AlertThresholdReport) Recipients() []*ReportRecipient {
	return m.recipientsField
}

// SetRecipients sets the recipients of this subtype
func (m *AlertThresholdReport) SetRecipients(val []*ReportRecipient) {
	m.recipientsField = val
}

// ReportLinkExpire gets the report link expire of this subtype
func (m *AlertThresholdReport) ReportLinkExpire() string {
	return m.reportLinkExpireField
}

// SetReportLinkExpire sets the report link expire of this subtype
func (m *AlertThresholdReport) SetReportLinkExpire(val string) {
	m.reportLinkExpireField = val
}

// ReportLinkNum gets the report link num of this subtype
func (m *AlertThresholdReport) ReportLinkNum() int32 {
	return m.reportLinkNumField
}

// SetReportLinkNum sets the report link num of this subtype
func (m *AlertThresholdReport) SetReportLinkNum(val int32) {
	m.reportLinkNumField = val
}

// Schedule gets the schedule of this subtype
func (m *AlertThresholdReport) Schedule() string {
	return m.scheduleField
}

// SetSchedule sets the schedule of this subtype
func (m *AlertThresholdReport) SetSchedule(val string) {
	m.scheduleField = val
}

// ScheduleTimezone gets the schedule timezone of this subtype
func (m *AlertThresholdReport) ScheduleTimezone() string {
	return m.scheduleTimezoneField
}

// SetScheduleTimezone sets the schedule timezone of this subtype
func (m *AlertThresholdReport) SetScheduleTimezone(val string) {
	m.scheduleTimezoneField = val
}

// Type gets the type of this subtype
func (m *AlertThresholdReport) Type() string {
	return "Alert threshold"
}

// SetType sets the type of this subtype
func (m *AlertThresholdReport) SetType(val string) {
}

// UserPermission gets the user permission of this subtype
func (m *AlertThresholdReport) UserPermission() string {
	return m.userPermissionField
}

// SetUserPermission sets the user permission of this subtype
func (m *AlertThresholdReport) SetUserPermission(val string) {
	m.userPermissionField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *AlertThresholdReport) UnmarshalJSON(raw []byte) error {
	var data struct {

		// The columns that will be displayed in the report. You should specify the columns in the order in which you'd like them to be displayed
		Columns []*DynamicColumn `json:"columns,omitempty"`

		// The datapoint to be included in the report. Glob expressions supported
		DataPoint string `json:"dataPoint,omitempty"`

		// The name of the datasource instance to be included in the report, where the syntax should be dataSourceDisplayName-InstanceName. If you're referencing a single instance datasource, you can just specify dataSourceDisplayName. Glob expressions supported
		DataSourceInstanceName string `json:"dataSourceInstanceName,omitempty"`

		// The display name of the device(s) to be included in the report. Glob expressions supported
		DeviceDisplayName string `json:"deviceDisplayName,omitempty"`

		// true: only variations from the global thresholds will be displayed
		// false: all thresholds will be displayed, including global thresholds an custom group and instance level thresholds
		// the default value is true
		ExcludeGlobal interface{} `json:"excludeGlobal,omitempty"`

		// The full path of the group whose member devices you are going to include in the report. Glob expressions supported
		GroupFullPath string `json:"groupFullPath,omitempty"`

		// host | datasource | datapoint
		// host: displayed thresholds will be sorted by device
		// datasource: displayed thresholds will be sorted by datasource instance
		// datapoint: displayed thresholds will be sorted by datapoint (metric)
		SortedBy string `json:"sortedBy,omitempty"`
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

	var result AlertThresholdReport

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

	result.Columns = data.Columns
	result.DataPoint = data.DataPoint
	result.DataSourceInstanceName = data.DataSourceInstanceName
	result.DeviceDisplayName = data.DeviceDisplayName
	result.ExcludeGlobal = data.ExcludeGlobal
	result.GroupFullPath = data.GroupFullPath
	result.SortedBy = data.SortedBy

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m AlertThresholdReport) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// The columns that will be displayed in the report. You should specify the columns in the order in which you'd like them to be displayed
		Columns []*DynamicColumn `json:"columns,omitempty"`

		// The datapoint to be included in the report. Glob expressions supported
		DataPoint string `json:"dataPoint,omitempty"`

		// The name of the datasource instance to be included in the report, where the syntax should be dataSourceDisplayName-InstanceName. If you're referencing a single instance datasource, you can just specify dataSourceDisplayName. Glob expressions supported
		DataSourceInstanceName string `json:"dataSourceInstanceName,omitempty"`

		// The display name of the device(s) to be included in the report. Glob expressions supported
		DeviceDisplayName string `json:"deviceDisplayName,omitempty"`

		// true: only variations from the global thresholds will be displayed
		// false: all thresholds will be displayed, including global thresholds an custom group and instance level thresholds
		// the default value is true
		ExcludeGlobal interface{} `json:"excludeGlobal,omitempty"`

		// The full path of the group whose member devices you are going to include in the report. Glob expressions supported
		GroupFullPath string `json:"groupFullPath,omitempty"`

		// host | datasource | datapoint
		// host: displayed thresholds will be sorted by device
		// datasource: displayed thresholds will be sorted by datasource instance
		// datapoint: displayed thresholds will be sorted by datapoint (metric)
		SortedBy string `json:"sortedBy,omitempty"`
	}{

		Columns: m.Columns,

		DataPoint: m.DataPoint,

		DataSourceInstanceName: m.DataSourceInstanceName,

		DeviceDisplayName: m.DeviceDisplayName,

		ExcludeGlobal: m.ExcludeGlobal,

		GroupFullPath: m.GroupFullPath,

		SortedBy: m.SortedBy,
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

// Validate validates this alert threshold report
func (m *AlertThresholdReport) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecipients(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateColumns(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AlertThresholdReport) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name()); err != nil {
		return err
	}

	return nil
}

func (m *AlertThresholdReport) validateRecipients(formats strfmt.Registry) error {

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

func (m *AlertThresholdReport) validateColumns(formats strfmt.Registry) error {

	if swag.IsZero(m.Columns) { // not required
		return nil
	}

	for i := 0; i < len(m.Columns); i++ {
		if swag.IsZero(m.Columns[i]) { // not required
			continue
		}

		if m.Columns[i] != nil {
			if err := m.Columns[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("columns" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this alert threshold report based on the context it is used
func (m *AlertThresholdReport) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if err := m.contextValidateColumns(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AlertThresholdReport) contextValidateCustomReportTypeID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "customReportTypeId", "body", int32(m.CustomReportTypeID())); err != nil {
		return err
	}

	return nil
}

func (m *AlertThresholdReport) contextValidateCustomReportTypeName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "customReportTypeName", "body", string(m.CustomReportTypeName())); err != nil {
		return err
	}

	return nil
}

func (m *AlertThresholdReport) contextValidateEnableViewAsOtherUser(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "enableViewAsOtherUser", "body", m.EnableViewAsOtherUser()); err != nil {
		return err
	}

	return nil
}

func (m *AlertThresholdReport) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int32(m.ID())); err != nil {
		return err
	}

	return nil
}

func (m *AlertThresholdReport) contextValidateLastGenerateOn(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "lastGenerateOn", "body", int64(m.LastGenerateOn())); err != nil {
		return err
	}

	return nil
}

func (m *AlertThresholdReport) contextValidateLastGeneratePages(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "lastGeneratePages", "body", int32(m.LastGeneratePages())); err != nil {
		return err
	}

	return nil
}

func (m *AlertThresholdReport) contextValidateLastGenerateSize(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "lastGenerateSize", "body", int64(m.LastGenerateSize())); err != nil {
		return err
	}

	return nil
}

func (m *AlertThresholdReport) contextValidateLastmodifyUserID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "lastmodifyUserId", "body", int32(m.LastmodifyUserID())); err != nil {
		return err
	}

	return nil
}

func (m *AlertThresholdReport) contextValidateLastmodifyUserName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "lastmodifyUserName", "body", string(m.LastmodifyUserName())); err != nil {
		return err
	}

	return nil
}

func (m *AlertThresholdReport) contextValidateRecipients(ctx context.Context, formats strfmt.Registry) error {

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

func (m *AlertThresholdReport) contextValidateReportLinkNum(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "reportLinkNum", "body", int32(m.ReportLinkNum())); err != nil {
		return err
	}

	return nil
}

func (m *AlertThresholdReport) contextValidateUserPermission(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "userPermission", "body", string(m.UserPermission())); err != nil {
		return err
	}

	return nil
}

func (m *AlertThresholdReport) contextValidateColumns(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Columns); i++ {

		if m.Columns[i] != nil {
			if err := m.Columns[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("columns" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AlertThresholdReport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AlertThresholdReport) UnmarshalBinary(b []byte) error {
	var res AlertThresholdReport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
