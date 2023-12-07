// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WebsiteGroupSDT website group SDT
//
// swagger:model WebsiteGroupSDT
type WebsiteGroupSDT struct {
	adminField string

	commentField string

	durationField int32

	endDateTimeField int64

	endDateTimeOnLocalField string

	endHourField int32

	endMinuteField int32

	hourField int32

	idField string

	isEffectiveField *bool

	minuteField int32

	monthDayField int32

	sdtTypeField string

	startDateTimeField int64

	startDateTimeOnLocalField string

	timezoneField string

	weekDayField string

	weekOfMonthField string

	// The Id of the service group that the SDT applies to
	// Required: true
	WebsiteGroupID *int32 `json:"websiteGroupId"`

	// The name of the service group that the SDT applies to
	// Read Only: true
	WebsiteGroupName string `json:"websiteGroupName,omitempty"`
}

// Admin gets the admin of this subtype
func (m *WebsiteGroupSDT) Admin() string {
	return m.adminField
}

// SetAdmin sets the admin of this subtype
func (m *WebsiteGroupSDT) SetAdmin(val string) {
	m.adminField = val
}

// Comment gets the comment of this subtype
func (m *WebsiteGroupSDT) Comment() string {
	return m.commentField
}

// SetComment sets the comment of this subtype
func (m *WebsiteGroupSDT) SetComment(val string) {
	m.commentField = val
}

// Duration gets the duration of this subtype
func (m *WebsiteGroupSDT) Duration() int32 {
	return m.durationField
}

// SetDuration sets the duration of this subtype
func (m *WebsiteGroupSDT) SetDuration(val int32) {
	m.durationField = val
}

// EndDateTime gets the end date time of this subtype
func (m *WebsiteGroupSDT) EndDateTime() int64 {
	return m.endDateTimeField
}

// SetEndDateTime sets the end date time of this subtype
func (m *WebsiteGroupSDT) SetEndDateTime(val int64) {
	m.endDateTimeField = val
}

// EndDateTimeOnLocal gets the end date time on local of this subtype
func (m *WebsiteGroupSDT) EndDateTimeOnLocal() string {
	return m.endDateTimeOnLocalField
}

// SetEndDateTimeOnLocal sets the end date time on local of this subtype
func (m *WebsiteGroupSDT) SetEndDateTimeOnLocal(val string) {
	m.endDateTimeOnLocalField = val
}

// EndHour gets the end hour of this subtype
func (m *WebsiteGroupSDT) EndHour() int32 {
	return m.endHourField
}

// SetEndHour sets the end hour of this subtype
func (m *WebsiteGroupSDT) SetEndHour(val int32) {
	m.endHourField = val
}

// EndMinute gets the end minute of this subtype
func (m *WebsiteGroupSDT) EndMinute() int32 {
	return m.endMinuteField
}

// SetEndMinute sets the end minute of this subtype
func (m *WebsiteGroupSDT) SetEndMinute(val int32) {
	m.endMinuteField = val
}

// Hour gets the hour of this subtype
func (m *WebsiteGroupSDT) Hour() int32 {
	return m.hourField
}

// SetHour sets the hour of this subtype
func (m *WebsiteGroupSDT) SetHour(val int32) {
	m.hourField = val
}

// ID gets the id of this subtype
func (m *WebsiteGroupSDT) ID() string {
	return m.idField
}

// SetID sets the id of this subtype
func (m *WebsiteGroupSDT) SetID(val string) {
	m.idField = val
}

// IsEffective gets the is effective of this subtype
func (m *WebsiteGroupSDT) IsEffective() *bool {
	return m.isEffectiveField
}

// SetIsEffective sets the is effective of this subtype
func (m *WebsiteGroupSDT) SetIsEffective(val *bool) {
	m.isEffectiveField = val
}

// Minute gets the minute of this subtype
func (m *WebsiteGroupSDT) Minute() int32 {
	return m.minuteField
}

// SetMinute sets the minute of this subtype
func (m *WebsiteGroupSDT) SetMinute(val int32) {
	m.minuteField = val
}

// MonthDay gets the month day of this subtype
func (m *WebsiteGroupSDT) MonthDay() int32 {
	return m.monthDayField
}

// SetMonthDay sets the month day of this subtype
func (m *WebsiteGroupSDT) SetMonthDay(val int32) {
	m.monthDayField = val
}

// SDTType gets the sdt type of this subtype
func (m *WebsiteGroupSDT) SDTType() string {
	return m.sdtTypeField
}

// SetSDTType sets the sdt type of this subtype
func (m *WebsiteGroupSDT) SetSDTType(val string) {
	m.sdtTypeField = val
}

// StartDateTime gets the start date time of this subtype
func (m *WebsiteGroupSDT) StartDateTime() int64 {
	return m.startDateTimeField
}

// SetStartDateTime sets the start date time of this subtype
func (m *WebsiteGroupSDT) SetStartDateTime(val int64) {
	m.startDateTimeField = val
}

// StartDateTimeOnLocal gets the start date time on local of this subtype
func (m *WebsiteGroupSDT) StartDateTimeOnLocal() string {
	return m.startDateTimeOnLocalField
}

// SetStartDateTimeOnLocal sets the start date time on local of this subtype
func (m *WebsiteGroupSDT) SetStartDateTimeOnLocal(val string) {
	m.startDateTimeOnLocalField = val
}

// Timezone gets the timezone of this subtype
func (m *WebsiteGroupSDT) Timezone() string {
	return m.timezoneField
}

// SetTimezone sets the timezone of this subtype
func (m *WebsiteGroupSDT) SetTimezone(val string) {
	m.timezoneField = val
}

// Type gets the type of this subtype
func (m *WebsiteGroupSDT) Type() string {
	return "WebsiteGroupSDT"
}

// SetType sets the type of this subtype
func (m *WebsiteGroupSDT) SetType(val string) {
}

// WeekDay gets the week day of this subtype
func (m *WebsiteGroupSDT) WeekDay() string {
	return m.weekDayField
}

// SetWeekDay sets the week day of this subtype
func (m *WebsiteGroupSDT) SetWeekDay(val string) {
	m.weekDayField = val
}

// WeekOfMonth gets the week of month of this subtype
func (m *WebsiteGroupSDT) WeekOfMonth() string {
	return m.weekOfMonthField
}

// SetWeekOfMonth sets the week of month of this subtype
func (m *WebsiteGroupSDT) SetWeekOfMonth(val string) {
	m.weekOfMonthField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *WebsiteGroupSDT) UnmarshalJSON(raw []byte) error {
	var data struct {

		// The Id of the service group that the SDT applies to
		// Required: true
		WebsiteGroupID *int32 `json:"websiteGroupId"`

		// The name of the service group that the SDT applies to
		// Read Only: true
		WebsiteGroupName string `json:"websiteGroupName,omitempty"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		Admin string `json:"admin,omitempty"`

		Comment string `json:"comment,omitempty"`

		Duration int32 `json:"duration,omitempty"`

		EndDateTime int64 `json:"endDateTime,omitempty"`

		EndDateTimeOnLocal string `json:"endDateTimeOnLocal,omitempty"`

		EndHour int32 `json:"endHour,omitempty"`

		EndMinute int32 `json:"endMinute,omitempty"`

		Hour int32 `json:"hour,omitempty"`

		ID string `json:"id,omitempty"`

		IsEffective *bool `json:"isEffective,omitempty"`

		Minute int32 `json:"minute,omitempty"`

		MonthDay int32 `json:"monthDay,omitempty"`

		SDTType string `json:"sdtType,omitempty"`

		StartDateTime int64 `json:"startDateTime,omitempty"`

		StartDateTimeOnLocal string `json:"startDateTimeOnLocal,omitempty"`

		Timezone string `json:"timezone,omitempty"`

		Type string `json:"type"`

		WeekDay string `json:"weekDay,omitempty"`

		WeekOfMonth string `json:"weekOfMonth,omitempty"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result WebsiteGroupSDT

	result.adminField = base.Admin

	result.commentField = base.Comment

	result.durationField = base.Duration

	result.endDateTimeField = base.EndDateTime

	result.endDateTimeOnLocalField = base.EndDateTimeOnLocal

	result.endHourField = base.EndHour

	result.endMinuteField = base.EndMinute

	result.hourField = base.Hour

	result.idField = base.ID

	result.isEffectiveField = base.IsEffective

	result.minuteField = base.Minute

	result.monthDayField = base.MonthDay

	result.sdtTypeField = base.SDTType

	result.startDateTimeField = base.StartDateTime

	result.startDateTimeOnLocalField = base.StartDateTimeOnLocal

	result.timezoneField = base.Timezone

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}
	result.weekDayField = base.WeekDay

	result.weekOfMonthField = base.WeekOfMonth

	result.WebsiteGroupID = data.WebsiteGroupID
	result.WebsiteGroupName = data.WebsiteGroupName

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m WebsiteGroupSDT) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// The Id of the service group that the SDT applies to
		// Required: true
		WebsiteGroupID *int32 `json:"websiteGroupId"`

		// The name of the service group that the SDT applies to
		// Read Only: true
		WebsiteGroupName string `json:"websiteGroupName,omitempty"`
	}{

		WebsiteGroupID: m.WebsiteGroupID,

		WebsiteGroupName: m.WebsiteGroupName,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Admin string `json:"admin,omitempty"`

		Comment string `json:"comment,omitempty"`

		Duration int32 `json:"duration,omitempty"`

		EndDateTime int64 `json:"endDateTime,omitempty"`

		EndDateTimeOnLocal string `json:"endDateTimeOnLocal,omitempty"`

		EndHour int32 `json:"endHour,omitempty"`

		EndMinute int32 `json:"endMinute,omitempty"`

		Hour int32 `json:"hour,omitempty"`

		ID string `json:"id,omitempty"`

		IsEffective *bool `json:"isEffective,omitempty"`

		Minute int32 `json:"minute,omitempty"`

		MonthDay int32 `json:"monthDay,omitempty"`

		SDTType string `json:"sdtType,omitempty"`

		StartDateTime int64 `json:"startDateTime,omitempty"`

		StartDateTimeOnLocal string `json:"startDateTimeOnLocal,omitempty"`

		Timezone string `json:"timezone,omitempty"`

		Type string `json:"type"`

		WeekDay string `json:"weekDay,omitempty"`

		WeekOfMonth string `json:"weekOfMonth,omitempty"`
	}{

		Admin: m.Admin(),

		Comment: m.Comment(),

		Duration: m.Duration(),

		EndDateTime: m.EndDateTime(),

		EndDateTimeOnLocal: m.EndDateTimeOnLocal(),

		EndHour: m.EndHour(),

		EndMinute: m.EndMinute(),

		Hour: m.Hour(),

		ID: m.ID(),

		IsEffective: m.IsEffective(),

		Minute: m.Minute(),

		MonthDay: m.MonthDay(),

		SDTType: m.SDTType(),

		StartDateTime: m.StartDateTime(),

		StartDateTimeOnLocal: m.StartDateTimeOnLocal(),

		Timezone: m.Timezone(),

		Type: m.Type(),

		WeekDay: m.WeekDay(),

		WeekOfMonth: m.WeekOfMonth(),
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this website group SDT
func (m *WebsiteGroupSDT) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateWebsiteGroupID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebsiteGroupSDT) validateWebsiteGroupID(formats strfmt.Registry) error {

	if err := validate.Required("websiteGroupId", "body", m.WebsiteGroupID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this website group SDT based on the context it is used
func (m *WebsiteGroupSDT) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAdmin(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEndDateTimeOnLocal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIsEffective(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStartDateTimeOnLocal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWebsiteGroupName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebsiteGroupSDT) contextValidateAdmin(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "admin", "body", string(m.Admin())); err != nil {
		return err
	}

	return nil
}

func (m *WebsiteGroupSDT) contextValidateEndDateTimeOnLocal(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "endDateTimeOnLocal", "body", string(m.EndDateTimeOnLocal())); err != nil {
		return err
	}

	return nil
}

func (m *WebsiteGroupSDT) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID())); err != nil {
		return err
	}

	return nil
}

func (m *WebsiteGroupSDT) contextValidateIsEffective(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "isEffective", "body", m.IsEffective()); err != nil {
		return err
	}

	return nil
}

func (m *WebsiteGroupSDT) contextValidateStartDateTimeOnLocal(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "startDateTimeOnLocal", "body", string(m.StartDateTimeOnLocal())); err != nil {
		return err
	}

	return nil
}

func (m *WebsiteGroupSDT) contextValidateWebsiteGroupName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "websiteGroupName", "body", string(m.WebsiteGroupName)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WebsiteGroupSDT) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebsiteGroupSDT) UnmarshalBinary(b []byte) error {
	var res WebsiteGroupSDT
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
