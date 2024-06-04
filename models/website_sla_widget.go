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

// WebsiteSLAWidget website SLA widget
//
// swagger:model WebsiteSLAWidget
type WebsiteSLAWidget struct {
	dashboardIdField *int32

	descriptionField string

	idField int32

	intervalField int32

	lastUpdatedByField string

	lastUpdatedOnField int64

	nameField *string

	themeField string

	timescaleField string

	userPermissionField string

	// The threshold of color changes
	ColorThresholds []*ColorThreshold `json:"colorThresholds,omitempty"`

	// The days that SLA should be computed for, separated by commas. 1=Sunday, 2=Monday, 3=Tuesday, 4=Wednesday, 5=Thursday, 6=Friday, 7=Saturday
	DaysInWeek string `json:"daysInWeek,omitempty"`

	// The websites that should be used to compute the SLA
	// Required: true
	Items []*WebsiteItemConfig `json:"items"`

	// The period during the selected days that the SLA should be computed for. * = all day, or a time range can be specified in the format of "hh:mm TO hh:mm", e.g. "01:15 TO 17:15"
	PeriodInOneDay string `json:"periodInOneDay,omitempty"`

	// The specific timezone for the widget
	Timezone string `json:"timezone,omitempty"`
}

// DashboardID gets the dashboard Id of this subtype
func (m *WebsiteSLAWidget) DashboardID() *int32 {
	return m.dashboardIdField
}

// SetDashboardID sets the dashboard Id of this subtype
func (m *WebsiteSLAWidget) SetDashboardID(val *int32) {
	m.dashboardIdField = val
}

// Description gets the description of this subtype
func (m *WebsiteSLAWidget) Description() string {
	return m.descriptionField
}

// SetDescription sets the description of this subtype
func (m *WebsiteSLAWidget) SetDescription(val string) {
	m.descriptionField = val
}

// ID gets the id of this subtype
func (m *WebsiteSLAWidget) ID() int32 {
	return m.idField
}

// SetID sets the id of this subtype
func (m *WebsiteSLAWidget) SetID(val int32) {
	m.idField = val
}

// Interval gets the interval of this subtype
func (m *WebsiteSLAWidget) Interval() int32 {
	return m.intervalField
}

// SetInterval sets the interval of this subtype
func (m *WebsiteSLAWidget) SetInterval(val int32) {
	m.intervalField = val
}

// LastUpdatedBy gets the last updated by of this subtype
func (m *WebsiteSLAWidget) LastUpdatedBy() string {
	return m.lastUpdatedByField
}

// SetLastUpdatedBy sets the last updated by of this subtype
func (m *WebsiteSLAWidget) SetLastUpdatedBy(val string) {
	m.lastUpdatedByField = val
}

// LastUpdatedOn gets the last updated on of this subtype
func (m *WebsiteSLAWidget) LastUpdatedOn() int64 {
	return m.lastUpdatedOnField
}

// SetLastUpdatedOn sets the last updated on of this subtype
func (m *WebsiteSLAWidget) SetLastUpdatedOn(val int64) {
	m.lastUpdatedOnField = val
}

// Name gets the name of this subtype
func (m *WebsiteSLAWidget) Name() *string {
	return m.nameField
}

// SetName sets the name of this subtype
func (m *WebsiteSLAWidget) SetName(val *string) {
	m.nameField = val
}

// Theme gets the theme of this subtype
func (m *WebsiteSLAWidget) Theme() string {
	return m.themeField
}

// SetTheme sets the theme of this subtype
func (m *WebsiteSLAWidget) SetTheme(val string) {
	m.themeField = val
}

// Timescale gets the timescale of this subtype
func (m *WebsiteSLAWidget) Timescale() string {
	return m.timescaleField
}

// SetTimescale sets the timescale of this subtype
func (m *WebsiteSLAWidget) SetTimescale(val string) {
	m.timescaleField = val
}

// Type gets the type of this subtype
func (m *WebsiteSLAWidget) Type() string {
	return "websiteSLA"
}

// SetType sets the type of this subtype
func (m *WebsiteSLAWidget) SetType(val string) {
}

// UserPermission gets the user permission of this subtype
func (m *WebsiteSLAWidget) UserPermission() string {
	return m.userPermissionField
}

// SetUserPermission sets the user permission of this subtype
func (m *WebsiteSLAWidget) SetUserPermission(val string) {
	m.userPermissionField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *WebsiteSLAWidget) UnmarshalJSON(raw []byte) error {
	var data struct {

		// The threshold of color changes
		ColorThresholds []*ColorThreshold `json:"colorThresholds,omitempty"`

		// The days that SLA should be computed for, separated by commas. 1=Sunday, 2=Monday, 3=Tuesday, 4=Wednesday, 5=Thursday, 6=Friday, 7=Saturday
		DaysInWeek string `json:"daysInWeek,omitempty"`

		// The websites that should be used to compute the SLA
		// Required: true
		Items []*WebsiteItemConfig `json:"items"`

		// The period during the selected days that the SLA should be computed for. * = all day, or a time range can be specified in the format of "hh:mm TO hh:mm", e.g. "01:15 TO 17:15"
		PeriodInOneDay string `json:"periodInOneDay,omitempty"`

		// The specific timezone for the widget
		Timezone string `json:"timezone,omitempty"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		DashboardID *int32 `json:"dashboardId"`

		Description string `json:"description,omitempty"`

		ID int32 `json:"id,omitempty"`

		Interval int32 `json:"interval,omitempty"`

		LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`

		LastUpdatedOn int64 `json:"lastUpdatedOn,omitempty"`

		Name *string `json:"name"`

		Theme string `json:"theme,omitempty"`

		Timescale string `json:"timescale,omitempty"`

		Type string `json:"type"`

		UserPermission string `json:"userPermission,omitempty"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result WebsiteSLAWidget

	result.dashboardIdField = base.DashboardID

	result.descriptionField = base.Description

	result.idField = base.ID

	result.intervalField = base.Interval

	result.lastUpdatedByField = base.LastUpdatedBy

	result.lastUpdatedOnField = base.LastUpdatedOn

	result.nameField = base.Name

	result.themeField = base.Theme

	result.timescaleField = base.Timescale

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}
	result.userPermissionField = base.UserPermission

	result.ColorThresholds = data.ColorThresholds
	result.DaysInWeek = data.DaysInWeek
	result.Items = data.Items
	result.PeriodInOneDay = data.PeriodInOneDay
	result.Timezone = data.Timezone

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m WebsiteSLAWidget) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// The threshold of color changes
		ColorThresholds []*ColorThreshold `json:"colorThresholds,omitempty"`

		// The days that SLA should be computed for, separated by commas. 1=Sunday, 2=Monday, 3=Tuesday, 4=Wednesday, 5=Thursday, 6=Friday, 7=Saturday
		DaysInWeek string `json:"daysInWeek,omitempty"`

		// The websites that should be used to compute the SLA
		// Required: true
		Items []*WebsiteItemConfig `json:"items"`

		// The period during the selected days that the SLA should be computed for. * = all day, or a time range can be specified in the format of "hh:mm TO hh:mm", e.g. "01:15 TO 17:15"
		PeriodInOneDay string `json:"periodInOneDay,omitempty"`

		// The specific timezone for the widget
		Timezone string `json:"timezone,omitempty"`
	}{

		ColorThresholds: m.ColorThresholds,

		DaysInWeek: m.DaysInWeek,

		Items: m.Items,

		PeriodInOneDay: m.PeriodInOneDay,

		Timezone: m.Timezone,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		DashboardID *int32 `json:"dashboardId"`

		Description string `json:"description,omitempty"`

		ID int32 `json:"id,omitempty"`

		Interval int32 `json:"interval,omitempty"`

		LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`

		LastUpdatedOn int64 `json:"lastUpdatedOn,omitempty"`

		Name *string `json:"name"`

		Theme string `json:"theme,omitempty"`

		Timescale string `json:"timescale,omitempty"`

		Type string `json:"type"`

		UserPermission string `json:"userPermission,omitempty"`
	}{

		DashboardID: m.DashboardID(),

		Description: m.Description(),

		ID: m.ID(),

		Interval: m.Interval(),

		LastUpdatedBy: m.LastUpdatedBy(),

		LastUpdatedOn: m.LastUpdatedOn(),

		Name: m.Name(),

		Theme: m.Theme(),

		Timescale: m.Timescale(),

		Type: m.Type(),

		UserPermission: m.UserPermission(),
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this website SLA widget
func (m *WebsiteSLAWidget) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDashboardID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateColorThresholds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebsiteSLAWidget) validateDashboardID(formats strfmt.Registry) error {

	if err := validate.Required("dashboardId", "body", m.DashboardID()); err != nil {
		return err
	}

	return nil
}

func (m *WebsiteSLAWidget) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name()); err != nil {
		return err
	}

	return nil
}

func (m *WebsiteSLAWidget) validateColorThresholds(formats strfmt.Registry) error {

	if swag.IsZero(m.ColorThresholds) { // not required
		return nil
	}

	for i := 0; i < len(m.ColorThresholds); i++ {
		if swag.IsZero(m.ColorThresholds[i]) { // not required
			continue
		}

		if m.ColorThresholds[i] != nil {
			if err := m.ColorThresholds[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("colorThresholds" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("colorThresholds" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WebsiteSLAWidget) validateItems(formats strfmt.Registry) error {

	if err := validate.Required("items", "body", m.Items); err != nil {
		return err
	}

	for i := 0; i < len(m.Items); i++ {
		if swag.IsZero(m.Items[i]) { // not required
			continue
		}

		if m.Items[i] != nil {
			if err := m.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this website SLA widget based on the context it is used
func (m *WebsiteSLAWidget) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLastUpdatedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastUpdatedOn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUserPermission(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateColorThresholds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebsiteSLAWidget) contextValidateLastUpdatedBy(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "lastUpdatedBy", "body", string(m.LastUpdatedBy())); err != nil {
		return err
	}

	return nil
}

func (m *WebsiteSLAWidget) contextValidateLastUpdatedOn(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "lastUpdatedOn", "body", int64(m.LastUpdatedOn())); err != nil {
		return err
	}

	return nil
}

func (m *WebsiteSLAWidget) contextValidateUserPermission(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "userPermission", "body", string(m.UserPermission())); err != nil {
		return err
	}

	return nil
}

func (m *WebsiteSLAWidget) contextValidateColorThresholds(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ColorThresholds); i++ {

		if m.ColorThresholds[i] != nil {

			if swag.IsZero(m.ColorThresholds[i]) { // not required
				return nil
			}

			if err := m.ColorThresholds[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("colorThresholds" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("colorThresholds" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WebsiteSLAWidget) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Items); i++ {

		if m.Items[i] != nil {

			if swag.IsZero(m.Items[i]) { // not required
				return nil
			}

			if err := m.Items[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *WebsiteSLAWidget) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebsiteSLAWidget) UnmarshalBinary(b []byte) error {
	var res WebsiteSLAWidget
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
