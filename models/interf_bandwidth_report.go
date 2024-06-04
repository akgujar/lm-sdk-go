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

// InterfBandwidthReport interf bandwidth report
//
// swagger:model InterfBandwidthReport
type InterfBandwidthReport struct {
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

	// The Time Range configured for the report. Last 2 hours | Last 24 hours | Last calendar day | Last 7 days | Last 14 days | Last 30 days | Last calendar month | Last 365 days | Any custom date range in this format: YYYY-MM-dd hh:mm TO YYYY-MM-dd hh:mm
	DateRange string `json:"dateRange,omitempty"`

	// The devices OR groups selected for the report, where multiple entities are separated by commas
	HostsVal string `json:"hostsVal,omitempty"`

	// host | group. The type of entities specified in the hostsVal field
	HostsValType string `json:"hostsValType,omitempty"`

	// true | false
	// false: Scale the number using 1000
	// true: Scale the number using 1024
	// Required: true
	IsBase1024 *bool `json:"isBase1024"`

	// The datapoint or calculation on a datapoint that will be included in the report, where each datapoint/calculation is specified by three fields: dataSourceId, instances (glob is okay)
	// Required: true
	Metrics []*Metric `json:"metrics"`

	// 0 | 1
	// 0: Text only - metrics will be displayed in a tabular format.
	// 1: One graph per instance - metrics will be displayed in a tabular format and one graph will be displayed per instance
	RowFormat int32 `json:"rowFormat,omitempty"`

	// true | false
	// false: Metrics will be displayed for all selected devices or groups
	// true: Metrics will only be displayed for the top ten device or groups
	// Required: true
	Top10Only *bool `json:"top10Only"`
}

// CustomReportTypeID gets the custom report type Id of this subtype
func (m *InterfBandwidthReport) CustomReportTypeID() int32 {
	return m.customReportTypeIdField
}

// SetCustomReportTypeID sets the custom report type Id of this subtype
func (m *InterfBandwidthReport) SetCustomReportTypeID(val int32) {
	m.customReportTypeIdField = val
}

// CustomReportTypeName gets the custom report type name of this subtype
func (m *InterfBandwidthReport) CustomReportTypeName() string {
	return m.customReportTypeNameField
}

// SetCustomReportTypeName sets the custom report type name of this subtype
func (m *InterfBandwidthReport) SetCustomReportTypeName(val string) {
	m.customReportTypeNameField = val
}

// Delivery gets the delivery of this subtype
func (m *InterfBandwidthReport) Delivery() string {
	return m.deliveryField
}

// SetDelivery sets the delivery of this subtype
func (m *InterfBandwidthReport) SetDelivery(val string) {
	m.deliveryField = val
}

// Description gets the description of this subtype
func (m *InterfBandwidthReport) Description() string {
	return m.descriptionField
}

// SetDescription sets the description of this subtype
func (m *InterfBandwidthReport) SetDescription(val string) {
	m.descriptionField = val
}

// EnableViewAsOtherUser gets the enable view as other user of this subtype
func (m *InterfBandwidthReport) EnableViewAsOtherUser() *bool {
	return m.enableViewAsOtherUserField
}

// SetEnableViewAsOtherUser sets the enable view as other user of this subtype
func (m *InterfBandwidthReport) SetEnableViewAsOtherUser(val *bool) {
	m.enableViewAsOtherUserField = val
}

// Format gets the format of this subtype
func (m *InterfBandwidthReport) Format() string {
	return m.formatField
}

// SetFormat sets the format of this subtype
func (m *InterfBandwidthReport) SetFormat(val string) {
	m.formatField = val
}

// GroupID gets the group Id of this subtype
func (m *InterfBandwidthReport) GroupID() int32 {
	return m.groupIdField
}

// SetGroupID sets the group Id of this subtype
func (m *InterfBandwidthReport) SetGroupID(val int32) {
	m.groupIdField = val
}

// ID gets the id of this subtype
func (m *InterfBandwidthReport) ID() int32 {
	return m.idField
}

// SetID sets the id of this subtype
func (m *InterfBandwidthReport) SetID(val int32) {
	m.idField = val
}

// LastGenerateOn gets the last generate on of this subtype
func (m *InterfBandwidthReport) LastGenerateOn() int64 {
	return m.lastGenerateOnField
}

// SetLastGenerateOn sets the last generate on of this subtype
func (m *InterfBandwidthReport) SetLastGenerateOn(val int64) {
	m.lastGenerateOnField = val
}

// LastGeneratePages gets the last generate pages of this subtype
func (m *InterfBandwidthReport) LastGeneratePages() int32 {
	return m.lastGeneratePagesField
}

// SetLastGeneratePages sets the last generate pages of this subtype
func (m *InterfBandwidthReport) SetLastGeneratePages(val int32) {
	m.lastGeneratePagesField = val
}

// LastGenerateSize gets the last generate size of this subtype
func (m *InterfBandwidthReport) LastGenerateSize() int64 {
	return m.lastGenerateSizeField
}

// SetLastGenerateSize sets the last generate size of this subtype
func (m *InterfBandwidthReport) SetLastGenerateSize(val int64) {
	m.lastGenerateSizeField = val
}

// LastmodifyUserID gets the lastmodify user Id of this subtype
func (m *InterfBandwidthReport) LastmodifyUserID() int32 {
	return m.lastmodifyUserIdField
}

// SetLastmodifyUserID sets the lastmodify user Id of this subtype
func (m *InterfBandwidthReport) SetLastmodifyUserID(val int32) {
	m.lastmodifyUserIdField = val
}

// LastmodifyUserName gets the lastmodify user name of this subtype
func (m *InterfBandwidthReport) LastmodifyUserName() string {
	return m.lastmodifyUserNameField
}

// SetLastmodifyUserName sets the lastmodify user name of this subtype
func (m *InterfBandwidthReport) SetLastmodifyUserName(val string) {
	m.lastmodifyUserNameField = val
}

// Name gets the name of this subtype
func (m *InterfBandwidthReport) Name() *string {
	return m.nameField
}

// SetName sets the name of this subtype
func (m *InterfBandwidthReport) SetName(val *string) {
	m.nameField = val
}

// Recipients gets the recipients of this subtype
func (m *InterfBandwidthReport) Recipients() []*ReportRecipient {
	return m.recipientsField
}

// SetRecipients sets the recipients of this subtype
func (m *InterfBandwidthReport) SetRecipients(val []*ReportRecipient) {
	m.recipientsField = val
}

// ReportLinkExpire gets the report link expire of this subtype
func (m *InterfBandwidthReport) ReportLinkExpire() string {
	return m.reportLinkExpireField
}

// SetReportLinkExpire sets the report link expire of this subtype
func (m *InterfBandwidthReport) SetReportLinkExpire(val string) {
	m.reportLinkExpireField = val
}

// ReportLinkNum gets the report link num of this subtype
func (m *InterfBandwidthReport) ReportLinkNum() int32 {
	return m.reportLinkNumField
}

// SetReportLinkNum sets the report link num of this subtype
func (m *InterfBandwidthReport) SetReportLinkNum(val int32) {
	m.reportLinkNumField = val
}

// Schedule gets the schedule of this subtype
func (m *InterfBandwidthReport) Schedule() string {
	return m.scheduleField
}

// SetSchedule sets the schedule of this subtype
func (m *InterfBandwidthReport) SetSchedule(val string) {
	m.scheduleField = val
}

// ScheduleTimezone gets the schedule timezone of this subtype
func (m *InterfBandwidthReport) ScheduleTimezone() string {
	return m.scheduleTimezoneField
}

// SetScheduleTimezone sets the schedule timezone of this subtype
func (m *InterfBandwidthReport) SetScheduleTimezone(val string) {
	m.scheduleTimezoneField = val
}

// Type gets the type of this subtype
func (m *InterfBandwidthReport) Type() string {
	return "Interfaces Bandwidth"
}

// SetType sets the type of this subtype
func (m *InterfBandwidthReport) SetType(val string) {
}

// UserPermission gets the user permission of this subtype
func (m *InterfBandwidthReport) UserPermission() string {
	return m.userPermissionField
}

// SetUserPermission sets the user permission of this subtype
func (m *InterfBandwidthReport) SetUserPermission(val string) {
	m.userPermissionField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *InterfBandwidthReport) UnmarshalJSON(raw []byte) error {
	var data struct {

		// The Time Range configured for the report. Last 2 hours | Last 24 hours | Last calendar day | Last 7 days | Last 14 days | Last 30 days | Last calendar month | Last 365 days | Any custom date range in this format: YYYY-MM-dd hh:mm TO YYYY-MM-dd hh:mm
		DateRange string `json:"dateRange,omitempty"`

		// The devices OR groups selected for the report, where multiple entities are separated by commas
		HostsVal string `json:"hostsVal,omitempty"`

		// host | group. The type of entities specified in the hostsVal field
		HostsValType string `json:"hostsValType,omitempty"`

		// true | false
		// false: Scale the number using 1000
		// true: Scale the number using 1024
		// Required: true
		IsBase1024 *bool `json:"isBase1024"`

		// The datapoint or calculation on a datapoint that will be included in the report, where each datapoint/calculation is specified by three fields: dataSourceId, instances (glob is okay)
		// Required: true
		Metrics []*Metric `json:"metrics"`

		// 0 | 1
		// 0: Text only - metrics will be displayed in a tabular format.
		// 1: One graph per instance - metrics will be displayed in a tabular format and one graph will be displayed per instance
		RowFormat int32 `json:"rowFormat,omitempty"`

		// true | false
		// false: Metrics will be displayed for all selected devices or groups
		// true: Metrics will only be displayed for the top ten device or groups
		// Required: true
		Top10Only *bool `json:"top10Only"`
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

	var result InterfBandwidthReport

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
	result.HostsVal = data.HostsVal
	result.HostsValType = data.HostsValType
	result.IsBase1024 = data.IsBase1024
	result.Metrics = data.Metrics
	result.RowFormat = data.RowFormat
	result.Top10Only = data.Top10Only

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m InterfBandwidthReport) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// The Time Range configured for the report. Last 2 hours | Last 24 hours | Last calendar day | Last 7 days | Last 14 days | Last 30 days | Last calendar month | Last 365 days | Any custom date range in this format: YYYY-MM-dd hh:mm TO YYYY-MM-dd hh:mm
		DateRange string `json:"dateRange,omitempty"`

		// The devices OR groups selected for the report, where multiple entities are separated by commas
		HostsVal string `json:"hostsVal,omitempty"`

		// host | group. The type of entities specified in the hostsVal field
		HostsValType string `json:"hostsValType,omitempty"`

		// true | false
		// false: Scale the number using 1000
		// true: Scale the number using 1024
		// Required: true
		IsBase1024 *bool `json:"isBase1024"`

		// The datapoint or calculation on a datapoint that will be included in the report, where each datapoint/calculation is specified by three fields: dataSourceId, instances (glob is okay)
		// Required: true
		Metrics []*Metric `json:"metrics"`

		// 0 | 1
		// 0: Text only - metrics will be displayed in a tabular format.
		// 1: One graph per instance - metrics will be displayed in a tabular format and one graph will be displayed per instance
		RowFormat int32 `json:"rowFormat,omitempty"`

		// true | false
		// false: Metrics will be displayed for all selected devices or groups
		// true: Metrics will only be displayed for the top ten device or groups
		// Required: true
		Top10Only *bool `json:"top10Only"`
	}{

		DateRange: m.DateRange,

		HostsVal: m.HostsVal,

		HostsValType: m.HostsValType,

		IsBase1024: m.IsBase1024,

		Metrics: m.Metrics,

		RowFormat: m.RowFormat,

		Top10Only: m.Top10Only,
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

// Validate validates this interf bandwidth report
func (m *InterfBandwidthReport) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecipients(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsBase1024(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetrics(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTop10Only(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InterfBandwidthReport) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name()); err != nil {
		return err
	}

	return nil
}

func (m *InterfBandwidthReport) validateRecipients(formats strfmt.Registry) error {

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
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("recipients" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *InterfBandwidthReport) validateIsBase1024(formats strfmt.Registry) error {

	if err := validate.Required("isBase1024", "body", m.IsBase1024); err != nil {
		return err
	}

	return nil
}

func (m *InterfBandwidthReport) validateMetrics(formats strfmt.Registry) error {

	if err := validate.Required("metrics", "body", m.Metrics); err != nil {
		return err
	}

	for i := 0; i < len(m.Metrics); i++ {
		if swag.IsZero(m.Metrics[i]) { // not required
			continue
		}

		if m.Metrics[i] != nil {
			if err := m.Metrics[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("metrics" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("metrics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *InterfBandwidthReport) validateTop10Only(formats strfmt.Registry) error {

	if err := validate.Required("top10Only", "body", m.Top10Only); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this interf bandwidth report based on the context it is used
func (m *InterfBandwidthReport) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if err := m.contextValidateMetrics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InterfBandwidthReport) contextValidateCustomReportTypeID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "customReportTypeId", "body", int32(m.CustomReportTypeID())); err != nil {
		return err
	}

	return nil
}

func (m *InterfBandwidthReport) contextValidateCustomReportTypeName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "customReportTypeName", "body", string(m.CustomReportTypeName())); err != nil {
		return err
	}

	return nil
}

func (m *InterfBandwidthReport) contextValidateEnableViewAsOtherUser(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "enableViewAsOtherUser", "body", m.EnableViewAsOtherUser()); err != nil {
		return err
	}

	return nil
}

func (m *InterfBandwidthReport) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int32(m.ID())); err != nil {
		return err
	}

	return nil
}

func (m *InterfBandwidthReport) contextValidateLastGenerateOn(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "lastGenerateOn", "body", int64(m.LastGenerateOn())); err != nil {
		return err
	}

	return nil
}

func (m *InterfBandwidthReport) contextValidateLastGeneratePages(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "lastGeneratePages", "body", int32(m.LastGeneratePages())); err != nil {
		return err
	}

	return nil
}

func (m *InterfBandwidthReport) contextValidateLastGenerateSize(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "lastGenerateSize", "body", int64(m.LastGenerateSize())); err != nil {
		return err
	}

	return nil
}

func (m *InterfBandwidthReport) contextValidateLastmodifyUserID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "lastmodifyUserId", "body", int32(m.LastmodifyUserID())); err != nil {
		return err
	}

	return nil
}

func (m *InterfBandwidthReport) contextValidateLastmodifyUserName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "lastmodifyUserName", "body", string(m.LastmodifyUserName())); err != nil {
		return err
	}

	return nil
}

func (m *InterfBandwidthReport) contextValidateRecipients(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Recipients()); i++ {

		if m.recipientsField[i] != nil {

			if swag.IsZero(m.recipientsField[i]) { // not required
				return nil
			}

			if err := m.recipientsField[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("recipients" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("recipients" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *InterfBandwidthReport) contextValidateReportLinkNum(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "reportLinkNum", "body", int32(m.ReportLinkNum())); err != nil {
		return err
	}

	return nil
}

func (m *InterfBandwidthReport) contextValidateUserPermission(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "userPermission", "body", string(m.UserPermission())); err != nil {
		return err
	}

	return nil
}

func (m *InterfBandwidthReport) contextValidateMetrics(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Metrics); i++ {

		if m.Metrics[i] != nil {

			if swag.IsZero(m.Metrics[i]) { // not required
				return nil
			}

			if err := m.Metrics[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("metrics" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("metrics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *InterfBandwidthReport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InterfBandwidthReport) UnmarshalBinary(b []byte) error {
	var res InterfBandwidthReport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
