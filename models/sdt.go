// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// SDT SDT
//
// swagger:discriminator SDT type
type SDT interface {
	runtime.Validatable
	runtime.ContextValidatable

	// The name of the user that created the SDT
	// Read Only: true
	Admin() string
	SetAdmin(string)

	// The notes associated with the SDT
	// Example: Emergency prod deployment
	Comment() string
	SetComment(string)

	// The duration of the SDT in minutes
	// Example: 138
	Duration() int32
	SetDuration(int32)

	// The epoch time, in milliseconds, that the SDT will end
	// Example: 1534554000000
	EndDateTime() int64
	SetEndDateTime(int64)

	// The date, time and time zone that the SDT will end at
	// Read Only: true
	EndDateTimeOnLocal() string
	SetEndDateTimeOnLocal(string)

	// The values can be 1 | 2....| 24. Specifies the hour that the SDT ends for a repeating SDT
	// Example: 5
	EndHour() int32
	SetEndHour(int32)

	// The values can be 1 | 2....| 60. Specifies the minute of the hour that the SDT ends for a repeating SDT
	// Example: 18
	EndMinute() int32
	SetEndMinute(int32)

	// The values can be 1 | 2....| 24. Specifies the hour that the SDT will start for a repeating SDT (daily, weekly, or monthly)
	// Example: 3
	Hour() int32
	SetHour(int32)

	// The Id of the SDT. This value will be in the following format "XX_##" where XX will refer to the type of SDT and ## will refer to the number of SDTs of that type
	// Read Only: true
	ID() string
	SetID(string)

	// The values can be true|false, where true: the SDT is currently active
	// false: the SDT is currently inactive
	// Read Only: true
	IsEffective() *bool
	SetIsEffective(*bool)

	// The values can be 1 | 2....| 60. Specifies the minute of the hour that the SDT should begin for a repeating SDT
	// Example: 6
	Minute() int32
	SetMinute(int32)

	// The values can be 1 | 2....| 31. Specifies the day of the month that the SDT will be active for a monthly SDT
	// Example: 7
	MonthDay() int32
	SetMonthDay(int32)

	// The type of sdt. The values can be oneTime|weekly|monthly|daily|monthlyByWeek
	// Example: oneTime
	SDTType() string
	SetSDTType(string)

	// The epoch time, in milliseconds, that the SDT will start
	// Example: 1534460400000
	StartDateTime() int64
	SetStartDateTime(int64)

	// The date, time and time zone that the SDT will end at
	// Read Only: true
	StartDateTimeOnLocal() string
	SetStartDateTimeOnLocal(string)

	// The specific timezone for SDT
	// Example: America/Los_Angeles
	Timezone() string
	SetTimezone(string)

	// The type of resource that this SDT is for. The values can be CollectorSDT | DeviceDataSourceInstanceSDT | DeviceBatchJobSDT | DeviceClusterAlertDefSDT | DeviceDataSourceInstanceGroupSDT | DeviceDataSourceSDT | DeviceEventSourceSDT | ResourceGroupSDT | ResourceSDT | WebsiteCheckpointSDT | WebsiteGroupSDT | WebsiteSDT | DeviceLogPipeLineResourceSDT
	// Example: ResourceGroupSDT
	// Required: true
	Type() string
	SetType(string)

	// The week day of sdt. The values can be SUNDAY|MONDAY|TUESDAY|WEDNESDAY|THURSDAY|FRIDAY|SATURDAY
	// Example: Sunday
	WeekDay() string
	SetWeekDay(string)

	// The week of the month that the SDT will be active for a monthly SDT
	// Example: 1
	WeekOfMonth() string
	SetWeekOfMonth(string)

	// AdditionalProperties in base type shoud be handled just like regular properties
	// At this moment, the base type property is pushed down to the subtype
}

type sdt struct {
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

	typeField string

	weekDayField string

	weekOfMonthField string
}

// Admin gets the admin of this polymorphic type
func (m *sdt) Admin() string {
	return m.adminField
}

// SetAdmin sets the admin of this polymorphic type
func (m *sdt) SetAdmin(val string) {
	m.adminField = val
}

// Comment gets the comment of this polymorphic type
func (m *sdt) Comment() string {
	return m.commentField
}

// SetComment sets the comment of this polymorphic type
func (m *sdt) SetComment(val string) {
	m.commentField = val
}

// Duration gets the duration of this polymorphic type
func (m *sdt) Duration() int32 {
	return m.durationField
}

// SetDuration sets the duration of this polymorphic type
func (m *sdt) SetDuration(val int32) {
	m.durationField = val
}

// EndDateTime gets the end date time of this polymorphic type
func (m *sdt) EndDateTime() int64 {
	return m.endDateTimeField
}

// SetEndDateTime sets the end date time of this polymorphic type
func (m *sdt) SetEndDateTime(val int64) {
	m.endDateTimeField = val
}

// EndDateTimeOnLocal gets the end date time on local of this polymorphic type
func (m *sdt) EndDateTimeOnLocal() string {
	return m.endDateTimeOnLocalField
}

// SetEndDateTimeOnLocal sets the end date time on local of this polymorphic type
func (m *sdt) SetEndDateTimeOnLocal(val string) {
	m.endDateTimeOnLocalField = val
}

// EndHour gets the end hour of this polymorphic type
func (m *sdt) EndHour() int32 {
	return m.endHourField
}

// SetEndHour sets the end hour of this polymorphic type
func (m *sdt) SetEndHour(val int32) {
	m.endHourField = val
}

// EndMinute gets the end minute of this polymorphic type
func (m *sdt) EndMinute() int32 {
	return m.endMinuteField
}

// SetEndMinute sets the end minute of this polymorphic type
func (m *sdt) SetEndMinute(val int32) {
	m.endMinuteField = val
}

// Hour gets the hour of this polymorphic type
func (m *sdt) Hour() int32 {
	return m.hourField
}

// SetHour sets the hour of this polymorphic type
func (m *sdt) SetHour(val int32) {
	m.hourField = val
}

// ID gets the id of this polymorphic type
func (m *sdt) ID() string {
	return m.idField
}

// SetID sets the id of this polymorphic type
func (m *sdt) SetID(val string) {
	m.idField = val
}

// IsEffective gets the is effective of this polymorphic type
func (m *sdt) IsEffective() *bool {
	return m.isEffectiveField
}

// SetIsEffective sets the is effective of this polymorphic type
func (m *sdt) SetIsEffective(val *bool) {
	m.isEffectiveField = val
}

// Minute gets the minute of this polymorphic type
func (m *sdt) Minute() int32 {
	return m.minuteField
}

// SetMinute sets the minute of this polymorphic type
func (m *sdt) SetMinute(val int32) {
	m.minuteField = val
}

// MonthDay gets the month day of this polymorphic type
func (m *sdt) MonthDay() int32 {
	return m.monthDayField
}

// SetMonthDay sets the month day of this polymorphic type
func (m *sdt) SetMonthDay(val int32) {
	m.monthDayField = val
}

// SDTType gets the sdt type of this polymorphic type
func (m *sdt) SDTType() string {
	return m.sdtTypeField
}

// SetSDTType sets the sdt type of this polymorphic type
func (m *sdt) SetSDTType(val string) {
	m.sdtTypeField = val
}

// StartDateTime gets the start date time of this polymorphic type
func (m *sdt) StartDateTime() int64 {
	return m.startDateTimeField
}

// SetStartDateTime sets the start date time of this polymorphic type
func (m *sdt) SetStartDateTime(val int64) {
	m.startDateTimeField = val
}

// StartDateTimeOnLocal gets the start date time on local of this polymorphic type
func (m *sdt) StartDateTimeOnLocal() string {
	return m.startDateTimeOnLocalField
}

// SetStartDateTimeOnLocal sets the start date time on local of this polymorphic type
func (m *sdt) SetStartDateTimeOnLocal(val string) {
	m.startDateTimeOnLocalField = val
}

// Timezone gets the timezone of this polymorphic type
func (m *sdt) Timezone() string {
	return m.timezoneField
}

// SetTimezone sets the timezone of this polymorphic type
func (m *sdt) SetTimezone(val string) {
	m.timezoneField = val
}

// Type gets the type of this polymorphic type
func (m *sdt) Type() string {
	return "SDT"
}

// SetType sets the type of this polymorphic type
func (m *sdt) SetType(val string) {
}

// WeekDay gets the week day of this polymorphic type
func (m *sdt) WeekDay() string {
	return m.weekDayField
}

// SetWeekDay sets the week day of this polymorphic type
func (m *sdt) SetWeekDay(val string) {
	m.weekDayField = val
}

// WeekOfMonth gets the week of month of this polymorphic type
func (m *sdt) WeekOfMonth() string {
	return m.weekOfMonthField
}

// SetWeekOfMonth sets the week of month of this polymorphic type
func (m *sdt) SetWeekOfMonth(val string) {
	m.weekOfMonthField = val
}

// UnmarshalSDTSlice unmarshals polymorphic slices of SDT
func UnmarshalSDTSlice(reader io.Reader, consumer runtime.Consumer) ([]SDT, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []SDT
	for _, element := range elements {
		obj, err := unmarshalSDT(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalSDT unmarshals polymorphic SDT
func UnmarshalSDT(reader io.Reader, consumer runtime.Consumer) (SDT, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalSDT(data, consumer)
}

func unmarshalSDT(data []byte, consumer runtime.Consumer) (SDT, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the type property.
	var getType struct {
		Type string `json:"type"`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("type", "body", getType.Type); err != nil {
		return nil, err
	}

	// The value of type is used to determine which type to create and unmarshal the data into
	switch getType.Type {
	case "CollectorSDT":
		var result CollectorSDT
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "DeviceBatchJobSDT":
		var result DeviceBatchJobSDT
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "DeviceClusterAlertDefSDT":
		var result DeviceClusterAlertDefSDT
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "DeviceDataSourceInstanceGroupSDT":
		var result DeviceDataSourceInstanceGroupSDT
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "DeviceDataSourceInstanceSDT":
		var result DeviceDataSourceInstanceSDT
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "DeviceDataSourceSDT":
		var result DeviceDataSourceSDT
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "DeviceEventSourceSDT":
		var result DeviceEventSourceSDT
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "DeviceLogPipeLineResourceSDT":
		var result DeviceLogPipeLineResourceSDT
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "ResourceGroupSDT":
		var result ResourceGroupSDT
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "ResourceSDT":
		var result ResourceSDT
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "SDT":
		var result sdt
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "WebsiteCheckpointSDT":
		var result WebsiteCheckpointSDT
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "WebsiteGroupSDT":
		var result WebsiteGroupSDT
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "WebsiteSDT":
		var result WebsiteSDT
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	}
	return nil, errors.New(422, "invalid type value: %q", getType.Type)
}

// Validate validates this SDT
func (m *sdt) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this SDT based on the context it is used
func (m *sdt) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *sdt) contextValidateAdmin(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "admin", "body", string(m.Admin())); err != nil {
		return err
	}

	return nil
}

func (m *sdt) contextValidateEndDateTimeOnLocal(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "endDateTimeOnLocal", "body", string(m.EndDateTimeOnLocal())); err != nil {
		return err
	}

	return nil
}

func (m *sdt) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID())); err != nil {
		return err
	}

	return nil
}

func (m *sdt) contextValidateIsEffective(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "isEffective", "body", m.IsEffective()); err != nil {
		return err
	}

	return nil
}

func (m *sdt) contextValidateStartDateTimeOnLocal(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "startDateTimeOnLocal", "body", string(m.StartDateTimeOnLocal())); err != nil {
		return err
	}

	return nil
}
