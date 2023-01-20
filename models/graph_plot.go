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

// GraphPlot graph plot
//
// swagger:model GraphPlot
type GraphPlot struct {
	titleField string

	// The Base of the graph
	// Read Only: true
	Base int64 `json:"base,omitempty"`

	// true | false
	// Changes base scale from 1000 to 1024 if value is set to true
	// Read Only: true
	Base1024 *bool `json:"base1024,omitempty"`

	// The display priority of the graph in your LogicMonitor portal
	// Read Only: true
	DisplayPrio int32 `json:"displayPrio,omitempty"`

	// The name of the DataSource to be used to plot the graph
	// Read Only: true
	DsName string `json:"dsName,omitempty"`

	// Specifies the end TimeZone Offset of the graph
	// Read Only: true
	EndTZOffset int32 `json:"endTZOffset,omitempty"`

	// Specifies the end-time of the graph
	// Read Only: true
	EndTime int64 `json:"endTime,omitempty"`

	// The export file name
	// Read Only: true
	ExportFileName string `json:"exportFileName,omitempty"`

	// Specifies the height of graph
	// Read Only: true
	Height int32 `json:"height,omitempty"`

	// The Id of the graph
	// Read Only: true
	ID int32 `json:"id,omitempty"`

	// The matched instances of graph
	// Read Only: true
	// Unique: true
	Instances []int32 `json:"instances,omitempty"`

	// The properties of the graph and graph lines
	// Read Only: true
	Lines []*GraphPlotLine `json:"lines,omitempty"`

	// Specifies the maximum value of the graph
	// Read Only: true
	MaxValue interface{} `json:"maxValue,omitempty"`

	// Specifies the minimum value of the graph
	// Read Only: true
	MinValue interface{} `json:"minValue,omitempty"`

	// The Missing lines of the graph
	// Read Only: true
	Missinglines []string `json:"missinglines,omitempty"`

	// The Name of the Graph
	// Read Only: true
	Name string `json:"name,omitempty"`

	// true | false
	// Specifies if the graph is rigid or not
	// Read Only: true
	Rigid *bool `json:"rigid,omitempty"`

	// Scopes: use this field to find match opsnote
	// Read Only: true
	Scopes []*GraphOpsNoteScope `json:"scopes,omitempty"`

	// Specifies the start TimeZone Offset of the graph
	// Read Only: true
	StartTZOffset int32 `json:"startTZOffset,omitempty"`

	// Specifies the start-time of the graph
	// Read Only: true
	StartTime int64 `json:"startTime,omitempty"`

	// The Step of the graph
	// Read Only: true
	Step int64 `json:"step,omitempty"`

	// The specified timescale for the graph
	// Read Only: true
	TimeScale string `json:"timeScale,omitempty"`

	// The selected timezone for the graph
	// Read Only: true
	TimeZone string `json:"timeZone,omitempty"`

	// The Id of selected Time Zone
	// Read Only: true
	TimeZoneID string `json:"timeZoneId,omitempty"`

	// The timestamps of the graph
	// Read Only: true
	Timestamps []int64 `json:"timestamps,omitempty"`

	// The label that will be displayed along the y axis (Vertical Label)
	// Read Only: true
	VerticalLabel string `json:"verticalLabel,omitempty"`

	// Specifies the width of graph
	// Read Only: true
	Width int32 `json:"width,omitempty"`

	// The label that will be displayed along the X axis
	// Read Only: true
	XAxisName string `json:"xAxisName,omitempty"`
}

// Title gets the title of this subtype
func (m *GraphPlot) Title() string {
	return m.titleField
}

// SetTitle sets the title of this subtype
func (m *GraphPlot) SetTitle(val string) {
	m.titleField = val
}

// Type gets the type of this subtype
func (m *GraphPlot) Type() string {
	return "graph"
}

// SetType sets the type of this subtype
func (m *GraphPlot) SetType(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *GraphPlot) UnmarshalJSON(raw []byte) error {
	var data struct {

		// The Base of the graph
		// Read Only: true
		Base int64 `json:"base,omitempty"`

		// true | false
		// Changes base scale from 1000 to 1024 if value is set to true
		// Read Only: true
		Base1024 *bool `json:"base1024,omitempty"`

		// The display priority of the graph in your LogicMonitor portal
		// Read Only: true
		DisplayPrio int32 `json:"displayPrio,omitempty"`

		// The name of the DataSource to be used to plot the graph
		// Read Only: true
		DsName string `json:"dsName,omitempty"`

		// Specifies the end TimeZone Offset of the graph
		// Read Only: true
		EndTZOffset int32 `json:"endTZOffset,omitempty"`

		// Specifies the end-time of the graph
		// Read Only: true
		EndTime int64 `json:"endTime,omitempty"`

		// The export file name
		// Read Only: true
		ExportFileName string `json:"exportFileName,omitempty"`

		// Specifies the height of graph
		// Read Only: true
		Height int32 `json:"height,omitempty"`

		// The Id of the graph
		// Read Only: true
		ID int32 `json:"id,omitempty"`

		// The matched instances of graph
		// Read Only: true
		// Unique: true
		Instances []int32 `json:"instances,omitempty"`

		// The properties of the graph and graph lines
		// Read Only: true
		Lines []*GraphPlotLine `json:"lines,omitempty"`

		// Specifies the maximum value of the graph
		// Read Only: true
		MaxValue interface{} `json:"maxValue,omitempty"`

		// Specifies the minimum value of the graph
		// Read Only: true
		MinValue interface{} `json:"minValue,omitempty"`

		// The Missing lines of the graph
		// Read Only: true
		Missinglines []string `json:"missinglines,omitempty"`

		// The Name of the Graph
		// Read Only: true
		Name string `json:"name,omitempty"`

		// true | false
		// Specifies if the graph is rigid or not
		// Read Only: true
		Rigid *bool `json:"rigid,omitempty"`

		// Scopes: use this field to find match opsnote
		// Read Only: true
		Scopes []*GraphOpsNoteScope `json:"scopes,omitempty"`

		// Specifies the start TimeZone Offset of the graph
		// Read Only: true
		StartTZOffset int32 `json:"startTZOffset,omitempty"`

		// Specifies the start-time of the graph
		// Read Only: true
		StartTime int64 `json:"startTime,omitempty"`

		// The Step of the graph
		// Read Only: true
		Step int64 `json:"step,omitempty"`

		// The specified timescale for the graph
		// Read Only: true
		TimeScale string `json:"timeScale,omitempty"`

		// The selected timezone for the graph
		// Read Only: true
		TimeZone string `json:"timeZone,omitempty"`

		// The Id of selected Time Zone
		// Read Only: true
		TimeZoneID string `json:"timeZoneId,omitempty"`

		// The timestamps of the graph
		// Read Only: true
		Timestamps []int64 `json:"timestamps,omitempty"`

		// The label that will be displayed along the y axis (Vertical Label)
		// Read Only: true
		VerticalLabel string `json:"verticalLabel,omitempty"`

		// Specifies the width of graph
		// Read Only: true
		Width int32 `json:"width,omitempty"`

		// The label that will be displayed along the X axis
		// Read Only: true
		XAxisName string `json:"xAxisName,omitempty"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		Title string `json:"title,omitempty"`

		Type string `json:"type,omitempty"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result GraphPlot

	result.titleField = base.Title

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}

	result.Base = data.Base
	result.Base1024 = data.Base1024
	result.DisplayPrio = data.DisplayPrio
	result.DsName = data.DsName
	result.EndTZOffset = data.EndTZOffset
	result.EndTime = data.EndTime
	result.ExportFileName = data.ExportFileName
	result.Height = data.Height
	result.ID = data.ID
	result.Instances = data.Instances
	result.Lines = data.Lines
	result.MaxValue = data.MaxValue
	result.MinValue = data.MinValue
	result.Missinglines = data.Missinglines
	result.Name = data.Name
	result.Rigid = data.Rigid
	result.Scopes = data.Scopes
	result.StartTZOffset = data.StartTZOffset
	result.StartTime = data.StartTime
	result.Step = data.Step
	result.TimeScale = data.TimeScale
	result.TimeZone = data.TimeZone
	result.TimeZoneID = data.TimeZoneID
	result.Timestamps = data.Timestamps
	result.VerticalLabel = data.VerticalLabel
	result.Width = data.Width
	result.XAxisName = data.XAxisName

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m GraphPlot) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// The Base of the graph
		// Read Only: true
		Base int64 `json:"base,omitempty"`

		// true | false
		// Changes base scale from 1000 to 1024 if value is set to true
		// Read Only: true
		Base1024 *bool `json:"base1024,omitempty"`

		// The display priority of the graph in your LogicMonitor portal
		// Read Only: true
		DisplayPrio int32 `json:"displayPrio,omitempty"`

		// The name of the DataSource to be used to plot the graph
		// Read Only: true
		DsName string `json:"dsName,omitempty"`

		// Specifies the end TimeZone Offset of the graph
		// Read Only: true
		EndTZOffset int32 `json:"endTZOffset,omitempty"`

		// Specifies the end-time of the graph
		// Read Only: true
		EndTime int64 `json:"endTime,omitempty"`

		// The export file name
		// Read Only: true
		ExportFileName string `json:"exportFileName,omitempty"`

		// Specifies the height of graph
		// Read Only: true
		Height int32 `json:"height,omitempty"`

		// The Id of the graph
		// Read Only: true
		ID int32 `json:"id,omitempty"`

		// The matched instances of graph
		// Read Only: true
		// Unique: true
		Instances []int32 `json:"instances,omitempty"`

		// The properties of the graph and graph lines
		// Read Only: true
		Lines []*GraphPlotLine `json:"lines,omitempty"`

		// Specifies the maximum value of the graph
		// Read Only: true
		MaxValue interface{} `json:"maxValue,omitempty"`

		// Specifies the minimum value of the graph
		// Read Only: true
		MinValue interface{} `json:"minValue,omitempty"`

		// The Missing lines of the graph
		// Read Only: true
		Missinglines []string `json:"missinglines,omitempty"`

		// The Name of the Graph
		// Read Only: true
		Name string `json:"name,omitempty"`

		// true | false
		// Specifies if the graph is rigid or not
		// Read Only: true
		Rigid *bool `json:"rigid,omitempty"`

		// Scopes: use this field to find match opsnote
		// Read Only: true
		Scopes []*GraphOpsNoteScope `json:"scopes,omitempty"`

		// Specifies the start TimeZone Offset of the graph
		// Read Only: true
		StartTZOffset int32 `json:"startTZOffset,omitempty"`

		// Specifies the start-time of the graph
		// Read Only: true
		StartTime int64 `json:"startTime,omitempty"`

		// The Step of the graph
		// Read Only: true
		Step int64 `json:"step,omitempty"`

		// The specified timescale for the graph
		// Read Only: true
		TimeScale string `json:"timeScale,omitempty"`

		// The selected timezone for the graph
		// Read Only: true
		TimeZone string `json:"timeZone,omitempty"`

		// The Id of selected Time Zone
		// Read Only: true
		TimeZoneID string `json:"timeZoneId,omitempty"`

		// The timestamps of the graph
		// Read Only: true
		Timestamps []int64 `json:"timestamps,omitempty"`

		// The label that will be displayed along the y axis (Vertical Label)
		// Read Only: true
		VerticalLabel string `json:"verticalLabel,omitempty"`

		// Specifies the width of graph
		// Read Only: true
		Width int32 `json:"width,omitempty"`

		// The label that will be displayed along the X axis
		// Read Only: true
		XAxisName string `json:"xAxisName,omitempty"`
	}{

		Base: m.Base,

		Base1024: m.Base1024,

		DisplayPrio: m.DisplayPrio,

		DsName: m.DsName,

		EndTZOffset: m.EndTZOffset,

		EndTime: m.EndTime,

		ExportFileName: m.ExportFileName,

		Height: m.Height,

		ID: m.ID,

		Instances: m.Instances,

		Lines: m.Lines,

		MaxValue: m.MaxValue,

		MinValue: m.MinValue,

		Missinglines: m.Missinglines,

		Name: m.Name,

		Rigid: m.Rigid,

		Scopes: m.Scopes,

		StartTZOffset: m.StartTZOffset,

		StartTime: m.StartTime,

		Step: m.Step,

		TimeScale: m.TimeScale,

		TimeZone: m.TimeZone,

		TimeZoneID: m.TimeZoneID,

		Timestamps: m.Timestamps,

		VerticalLabel: m.VerticalLabel,

		Width: m.Width,

		XAxisName: m.XAxisName,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Title string `json:"title,omitempty"`

		Type string `json:"type,omitempty"`
	}{

		Title: m.Title(),

		Type: m.Type(),
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this graph plot
func (m *GraphPlot) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInstances(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLines(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScopes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GraphPlot) validateInstances(formats strfmt.Registry) error {

	if swag.IsZero(m.Instances) { // not required
		return nil
	}

	if err := validate.UniqueItems("instances", "body", m.Instances); err != nil {
		return err
	}

	return nil
}

func (m *GraphPlot) validateLines(formats strfmt.Registry) error {

	if swag.IsZero(m.Lines) { // not required
		return nil
	}

	for i := 0; i < len(m.Lines); i++ {
		if swag.IsZero(m.Lines[i]) { // not required
			continue
		}

		if m.Lines[i] != nil {
			if err := m.Lines[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("lines" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GraphPlot) validateScopes(formats strfmt.Registry) error {

	if swag.IsZero(m.Scopes) { // not required
		return nil
	}

	for i := 0; i < len(m.Scopes); i++ {
		if swag.IsZero(m.Scopes[i]) { // not required
			continue
		}

		if m.Scopes[i] != nil {
			if err := m.Scopes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("scopes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this graph plot based on the context it is used
func (m *GraphPlot) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBase(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBase1024(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisplayPrio(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDsName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEndTZOffset(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEndTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExportFileName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHeight(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInstances(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLines(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMissinglines(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRigid(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateScopes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStartTZOffset(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStartTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStep(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTimeScale(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTimeZone(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTimeZoneID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTimestamps(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVerticalLabel(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWidth(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateXAxisName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GraphPlot) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "type", "body", string(m.Type())); err != nil {
		return err
	}

	return nil
}

func (m *GraphPlot) contextValidateBase(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "base", "body", int64(m.Base)); err != nil {
		return err
	}

	return nil
}

func (m *GraphPlot) contextValidateBase1024(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "base1024", "body", m.Base1024); err != nil {
		return err
	}

	return nil
}

func (m *GraphPlot) contextValidateDisplayPrio(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "displayPrio", "body", int32(m.DisplayPrio)); err != nil {
		return err
	}

	return nil
}

func (m *GraphPlot) contextValidateDsName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dsName", "body", string(m.DsName)); err != nil {
		return err
	}

	return nil
}

func (m *GraphPlot) contextValidateEndTZOffset(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "endTZOffset", "body", int32(m.EndTZOffset)); err != nil {
		return err
	}

	return nil
}

func (m *GraphPlot) contextValidateEndTime(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "endTime", "body", int64(m.EndTime)); err != nil {
		return err
	}

	return nil
}

func (m *GraphPlot) contextValidateExportFileName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "exportFileName", "body", string(m.ExportFileName)); err != nil {
		return err
	}

	return nil
}

func (m *GraphPlot) contextValidateHeight(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "height", "body", int32(m.Height)); err != nil {
		return err
	}

	return nil
}

func (m *GraphPlot) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int32(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *GraphPlot) contextValidateInstances(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "instances", "body", []int32(m.Instances)); err != nil {
		return err
	}

	return nil
}

func (m *GraphPlot) contextValidateLines(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "lines", "body", []*GraphPlotLine(m.Lines)); err != nil {
		return err
	}

	for i := 0; i < len(m.Lines); i++ {

		if m.Lines[i] != nil {
			if err := m.Lines[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("lines" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GraphPlot) contextValidateMissinglines(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "missinglines", "body", []string(m.Missinglines)); err != nil {
		return err
	}

	return nil
}

func (m *GraphPlot) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

func (m *GraphPlot) contextValidateRigid(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "rigid", "body", m.Rigid); err != nil {
		return err
	}

	return nil
}

func (m *GraphPlot) contextValidateScopes(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "scopes", "body", []*GraphOpsNoteScope(m.Scopes)); err != nil {
		return err
	}

	for i := 0; i < len(m.Scopes); i++ {

		if m.Scopes[i] != nil {
			if err := m.Scopes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("scopes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GraphPlot) contextValidateStartTZOffset(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "startTZOffset", "body", int32(m.StartTZOffset)); err != nil {
		return err
	}

	return nil
}

func (m *GraphPlot) contextValidateStartTime(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "startTime", "body", int64(m.StartTime)); err != nil {
		return err
	}

	return nil
}

func (m *GraphPlot) contextValidateStep(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "step", "body", int64(m.Step)); err != nil {
		return err
	}

	return nil
}

func (m *GraphPlot) contextValidateTimeScale(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "timeScale", "body", string(m.TimeScale)); err != nil {
		return err
	}

	return nil
}

func (m *GraphPlot) contextValidateTimeZone(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "timeZone", "body", string(m.TimeZone)); err != nil {
		return err
	}

	return nil
}

func (m *GraphPlot) contextValidateTimeZoneID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "timeZoneId", "body", string(m.TimeZoneID)); err != nil {
		return err
	}

	return nil
}

func (m *GraphPlot) contextValidateTimestamps(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "timestamps", "body", []int64(m.Timestamps)); err != nil {
		return err
	}

	return nil
}

func (m *GraphPlot) contextValidateVerticalLabel(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "verticalLabel", "body", string(m.VerticalLabel)); err != nil {
		return err
	}

	return nil
}

func (m *GraphPlot) contextValidateWidth(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "width", "body", int32(m.Width)); err != nil {
		return err
	}

	return nil
}

func (m *GraphPlot) contextValidateXAxisName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "xAxisName", "body", string(m.XAxisName)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GraphPlot) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GraphPlot) UnmarshalBinary(b []byte) error {
	var res GraphPlot
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
