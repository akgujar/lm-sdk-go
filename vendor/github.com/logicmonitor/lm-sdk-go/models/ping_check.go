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

// PingCheck ping check
//
// swagger:model PingCheck
type PingCheck struct {
	checkpointsField []*WebsiteCheckPoint

	collectorsField []*WebsiteCollectorInfo

	descriptionField string

	disableAlertingField bool

	globalSmAlertCondField int32

	groupIdField int32

	idField int32

	individualAlertLevelField string

	individualSmAlertEnableField bool

	isInternalField bool

	lastUpdatedField int64

	nameField *string

	overallAlertLevelField string

	pollingIntervalField int32

	propertiesField []*NameAndValue

	statusField string

	stopMonitoringField bool

	stopMonitoringByFolderField *bool

	templateField interface{}

	testLocationField *WebsiteLocation

	transitionField int32

	useDefaultAlertSettingField bool

	useDefaultLocationSettingField bool

	userPermissionField string

	// The number of packets to send
	Count int32 `json:"count,omitempty"`

	// The URL to check, without the scheme or protocol (e.g http or https)
	// E.g. if the URL is "http://www.google.com, then the host="www.google.com"
	// Required: true
	Host *string `json:"host"`

	// The percentage of packets that should be returned in the time period specified by timeoutInMSPktsNotReceive for each ping check
	PercentPktsNotReceiveInTime int32 `json:"percentPktsNotReceiveInTime,omitempty"`

	// The time period that the percentage of packets specified by percentPktsNotReceiveInTime must be returned in for each ping check
	TimeoutInMSPktsNotReceive int64 `json:"timeoutInMSPktsNotReceive,omitempty"`
}

// Checkpoints gets the checkpoints of this subtype
func (m *PingCheck) Checkpoints() []*WebsiteCheckPoint {
	return m.checkpointsField
}

// SetCheckpoints sets the checkpoints of this subtype
func (m *PingCheck) SetCheckpoints(val []*WebsiteCheckPoint) {
	m.checkpointsField = val
}

// Collectors gets the collectors of this subtype
func (m *PingCheck) Collectors() []*WebsiteCollectorInfo {
	return m.collectorsField
}

// SetCollectors sets the collectors of this subtype
func (m *PingCheck) SetCollectors(val []*WebsiteCollectorInfo) {
	m.collectorsField = val
}

// Description gets the description of this subtype
func (m *PingCheck) Description() string {
	return m.descriptionField
}

// SetDescription sets the description of this subtype
func (m *PingCheck) SetDescription(val string) {
	m.descriptionField = val
}

// DisableAlerting gets the disable alerting of this subtype
func (m *PingCheck) DisableAlerting() bool {
	return m.disableAlertingField
}

// SetDisableAlerting sets the disable alerting of this subtype
func (m *PingCheck) SetDisableAlerting(val bool) {
	m.disableAlertingField = val
}

// GlobalSmAlertCond gets the global sm alert cond of this subtype
func (m *PingCheck) GlobalSmAlertCond() int32 {
	return m.globalSmAlertCondField
}

// SetGlobalSmAlertCond sets the global sm alert cond of this subtype
func (m *PingCheck) SetGlobalSmAlertCond(val int32) {
	m.globalSmAlertCondField = val
}

// GroupID gets the group Id of this subtype
func (m *PingCheck) GroupID() int32 {
	return m.groupIdField
}

// SetGroupID sets the group Id of this subtype
func (m *PingCheck) SetGroupID(val int32) {
	m.groupIdField = val
}

// ID gets the id of this subtype
func (m *PingCheck) ID() int32 {
	return m.idField
}

// SetID sets the id of this subtype
func (m *PingCheck) SetID(val int32) {
	m.idField = val
}

// IndividualAlertLevel gets the individual alert level of this subtype
func (m *PingCheck) IndividualAlertLevel() string {
	return m.individualAlertLevelField
}

// SetIndividualAlertLevel sets the individual alert level of this subtype
func (m *PingCheck) SetIndividualAlertLevel(val string) {
	m.individualAlertLevelField = val
}

// IndividualSmAlertEnable gets the individual sm alert enable of this subtype
func (m *PingCheck) IndividualSmAlertEnable() bool {
	return m.individualSmAlertEnableField
}

// SetIndividualSmAlertEnable sets the individual sm alert enable of this subtype
func (m *PingCheck) SetIndividualSmAlertEnable(val bool) {
	m.individualSmAlertEnableField = val
}

// IsInternal gets the is internal of this subtype
func (m *PingCheck) IsInternal() bool {
	return m.isInternalField
}

// SetIsInternal sets the is internal of this subtype
func (m *PingCheck) SetIsInternal(val bool) {
	m.isInternalField = val
}

// LastUpdated gets the last updated of this subtype
func (m *PingCheck) LastUpdated() int64 {
	return m.lastUpdatedField
}

// SetLastUpdated sets the last updated of this subtype
func (m *PingCheck) SetLastUpdated(val int64) {
	m.lastUpdatedField = val
}

// Name gets the name of this subtype
func (m *PingCheck) Name() *string {
	return m.nameField
}

// SetName sets the name of this subtype
func (m *PingCheck) SetName(val *string) {
	m.nameField = val
}

// OverallAlertLevel gets the overall alert level of this subtype
func (m *PingCheck) OverallAlertLevel() string {
	return m.overallAlertLevelField
}

// SetOverallAlertLevel sets the overall alert level of this subtype
func (m *PingCheck) SetOverallAlertLevel(val string) {
	m.overallAlertLevelField = val
}

// PollingInterval gets the polling interval of this subtype
func (m *PingCheck) PollingInterval() int32 {
	return m.pollingIntervalField
}

// SetPollingInterval sets the polling interval of this subtype
func (m *PingCheck) SetPollingInterval(val int32) {
	m.pollingIntervalField = val
}

// Properties gets the properties of this subtype
func (m *PingCheck) Properties() []*NameAndValue {
	return m.propertiesField
}

// SetProperties sets the properties of this subtype
func (m *PingCheck) SetProperties(val []*NameAndValue) {
	m.propertiesField = val
}

// Status gets the status of this subtype
func (m *PingCheck) Status() string {
	return m.statusField
}

// SetStatus sets the status of this subtype
func (m *PingCheck) SetStatus(val string) {
	m.statusField = val
}

// StopMonitoring gets the stop monitoring of this subtype
func (m *PingCheck) StopMonitoring() bool {
	return m.stopMonitoringField
}

// SetStopMonitoring sets the stop monitoring of this subtype
func (m *PingCheck) SetStopMonitoring(val bool) {
	m.stopMonitoringField = val
}

// StopMonitoringByFolder gets the stop monitoring by folder of this subtype
func (m *PingCheck) StopMonitoringByFolder() *bool {
	return m.stopMonitoringByFolderField
}

// SetStopMonitoringByFolder sets the stop monitoring by folder of this subtype
func (m *PingCheck) SetStopMonitoringByFolder(val *bool) {
	m.stopMonitoringByFolderField = val
}

// Template gets the template of this subtype
func (m *PingCheck) Template() interface{} {
	return m.templateField
}

// SetTemplate sets the template of this subtype
func (m *PingCheck) SetTemplate(val interface{}) {
	m.templateField = val
}

// TestLocation gets the test location of this subtype
func (m *PingCheck) TestLocation() *WebsiteLocation {
	return m.testLocationField
}

// SetTestLocation sets the test location of this subtype
func (m *PingCheck) SetTestLocation(val *WebsiteLocation) {
	m.testLocationField = val
}

// Transition gets the transition of this subtype
func (m *PingCheck) Transition() int32 {
	return m.transitionField
}

// SetTransition sets the transition of this subtype
func (m *PingCheck) SetTransition(val int32) {
	m.transitionField = val
}

// Type gets the type of this subtype
func (m *PingCheck) Type() string {
	return "pingcheck"
}

// SetType sets the type of this subtype
func (m *PingCheck) SetType(val string) {
}

// UseDefaultAlertSetting gets the use default alert setting of this subtype
func (m *PingCheck) UseDefaultAlertSetting() bool {
	return m.useDefaultAlertSettingField
}

// SetUseDefaultAlertSetting sets the use default alert setting of this subtype
func (m *PingCheck) SetUseDefaultAlertSetting(val bool) {
	m.useDefaultAlertSettingField = val
}

// UseDefaultLocationSetting gets the use default location setting of this subtype
func (m *PingCheck) UseDefaultLocationSetting() bool {
	return m.useDefaultLocationSettingField
}

// SetUseDefaultLocationSetting sets the use default location setting of this subtype
func (m *PingCheck) SetUseDefaultLocationSetting(val bool) {
	m.useDefaultLocationSettingField = val
}

// UserPermission gets the user permission of this subtype
func (m *PingCheck) UserPermission() string {
	return m.userPermissionField
}

// SetUserPermission sets the user permission of this subtype
func (m *PingCheck) SetUserPermission(val string) {
	m.userPermissionField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *PingCheck) UnmarshalJSON(raw []byte) error {
	var data struct {

		// The number of packets to send
		Count int32 `json:"count,omitempty"`

		// The URL to check, without the scheme or protocol (e.g http or https)
		// E.g. if the URL is "http://www.google.com, then the host="www.google.com"
		// Required: true
		Host *string `json:"host"`

		// The percentage of packets that should be returned in the time period specified by timeoutInMSPktsNotReceive for each ping check
		PercentPktsNotReceiveInTime int32 `json:"percentPktsNotReceiveInTime,omitempty"`

		// The time period that the percentage of packets specified by percentPktsNotReceiveInTime must be returned in for each ping check
		TimeoutInMSPktsNotReceive int64 `json:"timeoutInMSPktsNotReceive,omitempty"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		Checkpoints []*WebsiteCheckPoint `json:"checkpoints,omitempty"`

		Collectors []*WebsiteCollectorInfo `json:"collectors,omitempty"`

		Description string `json:"description,omitempty"`

		DisableAlerting bool `json:"disableAlerting,omitempty"`

		GlobalSmAlertCond int32 `json:"globalSmAlertCond,omitempty"`

		GroupID int32 `json:"groupId,omitempty"`

		ID int32 `json:"id,omitempty"`

		IndividualAlertLevel string `json:"individualAlertLevel,omitempty"`

		IndividualSmAlertEnable bool `json:"individualSmAlertEnable,omitempty"`

		IsInternal bool `json:"isInternal,omitempty"`

		LastUpdated int64 `json:"lastUpdated,omitempty"`

		Name *string `json:"name"`

		OverallAlertLevel string `json:"overallAlertLevel,omitempty"`

		PollingInterval int32 `json:"pollingInterval,omitempty"`

		Properties []*NameAndValue `json:"properties,omitempty"`

		Status string `json:"status,omitempty"`

		StopMonitoring bool `json:"stopMonitoring,omitempty"`

		StopMonitoringByFolder *bool `json:"stopMonitoringByFolder,omitempty"`

		Template interface{} `json:"template,omitempty"`

		TestLocation *WebsiteLocation `json:"testLocation"`

		Transition int32 `json:"transition,omitempty"`

		Type string `json:"type"`

		UseDefaultAlertSetting bool `json:"useDefaultAlertSetting,omitempty"`

		UseDefaultLocationSetting bool `json:"useDefaultLocationSetting,omitempty"`

		UserPermission string `json:"userPermission,omitempty"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result PingCheck

	result.checkpointsField = base.Checkpoints

	result.collectorsField = base.Collectors

	result.descriptionField = base.Description

	result.disableAlertingField = base.DisableAlerting

	result.globalSmAlertCondField = base.GlobalSmAlertCond

	result.groupIdField = base.GroupID

	result.idField = base.ID

	result.individualAlertLevelField = base.IndividualAlertLevel

	result.individualSmAlertEnableField = base.IndividualSmAlertEnable

	result.isInternalField = base.IsInternal

	result.lastUpdatedField = base.LastUpdated

	result.nameField = base.Name

	result.overallAlertLevelField = base.OverallAlertLevel

	result.pollingIntervalField = base.PollingInterval

	result.propertiesField = base.Properties

	result.statusField = base.Status

	result.stopMonitoringField = base.StopMonitoring

	result.stopMonitoringByFolderField = base.StopMonitoringByFolder

	result.templateField = base.Template

	result.testLocationField = base.TestLocation

	result.transitionField = base.Transition

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}
	result.useDefaultAlertSettingField = base.UseDefaultAlertSetting

	result.useDefaultLocationSettingField = base.UseDefaultLocationSetting

	result.userPermissionField = base.UserPermission

	result.Count = data.Count
	result.Host = data.Host
	result.PercentPktsNotReceiveInTime = data.PercentPktsNotReceiveInTime
	result.TimeoutInMSPktsNotReceive = data.TimeoutInMSPktsNotReceive

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m PingCheck) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// The number of packets to send
		Count int32 `json:"count,omitempty"`

		// The URL to check, without the scheme or protocol (e.g http or https)
		// E.g. if the URL is "http://www.google.com, then the host="www.google.com"
		// Required: true
		Host *string `json:"host"`

		// The percentage of packets that should be returned in the time period specified by timeoutInMSPktsNotReceive for each ping check
		PercentPktsNotReceiveInTime int32 `json:"percentPktsNotReceiveInTime,omitempty"`

		// The time period that the percentage of packets specified by percentPktsNotReceiveInTime must be returned in for each ping check
		TimeoutInMSPktsNotReceive int64 `json:"timeoutInMSPktsNotReceive,omitempty"`
	}{

		Count: m.Count,

		Host: m.Host,

		PercentPktsNotReceiveInTime: m.PercentPktsNotReceiveInTime,

		TimeoutInMSPktsNotReceive: m.TimeoutInMSPktsNotReceive,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Checkpoints []*WebsiteCheckPoint `json:"checkpoints,omitempty"`

		Collectors []*WebsiteCollectorInfo `json:"collectors,omitempty"`

		Description string `json:"description,omitempty"`

		DisableAlerting bool `json:"disableAlerting,omitempty"`

		GlobalSmAlertCond int32 `json:"globalSmAlertCond,omitempty"`

		GroupID int32 `json:"groupId,omitempty"`

		ID int32 `json:"id,omitempty"`

		IndividualAlertLevel string `json:"individualAlertLevel,omitempty"`

		IndividualSmAlertEnable bool `json:"individualSmAlertEnable,omitempty"`

		IsInternal bool `json:"isInternal,omitempty"`

		LastUpdated int64 `json:"lastUpdated,omitempty"`

		Name *string `json:"name"`

		OverallAlertLevel string `json:"overallAlertLevel,omitempty"`

		PollingInterval int32 `json:"pollingInterval,omitempty"`

		Properties []*NameAndValue `json:"properties,omitempty"`

		Status string `json:"status,omitempty"`

		StopMonitoring bool `json:"stopMonitoring,omitempty"`

		StopMonitoringByFolder *bool `json:"stopMonitoringByFolder,omitempty"`

		Template interface{} `json:"template,omitempty"`

		TestLocation *WebsiteLocation `json:"testLocation"`

		Transition int32 `json:"transition,omitempty"`

		Type string `json:"type"`

		UseDefaultAlertSetting bool `json:"useDefaultAlertSetting,omitempty"`

		UseDefaultLocationSetting bool `json:"useDefaultLocationSetting,omitempty"`

		UserPermission string `json:"userPermission,omitempty"`
	}{

		Checkpoints: m.Checkpoints(),

		Collectors: m.Collectors(),

		Description: m.Description(),

		DisableAlerting: m.DisableAlerting(),

		GlobalSmAlertCond: m.GlobalSmAlertCond(),

		GroupID: m.GroupID(),

		ID: m.ID(),

		IndividualAlertLevel: m.IndividualAlertLevel(),

		IndividualSmAlertEnable: m.IndividualSmAlertEnable(),

		IsInternal: m.IsInternal(),

		LastUpdated: m.LastUpdated(),

		Name: m.Name(),

		OverallAlertLevel: m.OverallAlertLevel(),

		PollingInterval: m.PollingInterval(),

		Properties: m.Properties(),

		Status: m.Status(),

		StopMonitoring: m.StopMonitoring(),

		StopMonitoringByFolder: m.StopMonitoringByFolder(),

		Template: m.Template(),

		TestLocation: m.TestLocation(),

		Transition: m.Transition(),

		Type: m.Type(),

		UseDefaultAlertSetting: m.UseDefaultAlertSetting(),

		UseDefaultLocationSetting: m.UseDefaultLocationSetting(),

		UserPermission: m.UserPermission(),
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this ping check
func (m *PingCheck) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCheckpoints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCollectors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProperties(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTestLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHost(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PingCheck) validateCheckpoints(formats strfmt.Registry) error {

	if swag.IsZero(m.Checkpoints()) { // not required
		return nil
	}

	for i := 0; i < len(m.Checkpoints()); i++ {
		if swag.IsZero(m.checkpointsField[i]) { // not required
			continue
		}

		if m.checkpointsField[i] != nil {
			if err := m.checkpointsField[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("checkpoints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PingCheck) validateCollectors(formats strfmt.Registry) error {

	if swag.IsZero(m.Collectors()) { // not required
		return nil
	}

	for i := 0; i < len(m.Collectors()); i++ {
		if swag.IsZero(m.collectorsField[i]) { // not required
			continue
		}

		if m.collectorsField[i] != nil {
			if err := m.collectorsField[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("collectors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PingCheck) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name()); err != nil {
		return err
	}

	return nil
}

func (m *PingCheck) validateProperties(formats strfmt.Registry) error {

	if swag.IsZero(m.Properties()) { // not required
		return nil
	}

	for i := 0; i < len(m.Properties()); i++ {
		if swag.IsZero(m.propertiesField[i]) { // not required
			continue
		}

		if m.propertiesField[i] != nil {
			if err := m.propertiesField[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("properties" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PingCheck) validateTestLocation(formats strfmt.Registry) error {

	if err := validate.Required("testLocation", "body", m.TestLocation()); err != nil {
		return err
	}

	if m.TestLocation() != nil {
		if err := m.TestLocation().Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("testLocation")
			}
			return err
		}
	}

	return nil
}

func (m *PingCheck) validateHost(formats strfmt.Registry) error {

	if err := validate.Required("host", "body", m.Host); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this ping check based on the context it is used
func (m *PingCheck) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCheckpoints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCollectors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGroupID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastUpdated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProperties(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStopMonitoringByFolder(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTestLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PingCheck) contextValidateCheckpoints(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Checkpoints()); i++ {

		if m.checkpointsField[i] != nil {
			if err := m.checkpointsField[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("checkpoints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PingCheck) contextValidateCollectors(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "collectors", "body", []*WebsiteCollectorInfo(m.Collectors())); err != nil {
		return err
	}

	for i := 0; i < len(m.Collectors()); i++ {

		if m.collectorsField[i] != nil {
			if err := m.collectorsField[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("collectors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PingCheck) contextValidateGroupID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "groupId", "body", int32(m.GroupID())); err != nil {
		return err
	}

	return nil
}

func (m *PingCheck) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int32(m.ID())); err != nil {
		return err
	}

	return nil
}

func (m *PingCheck) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "lastUpdated", "body", int64(m.LastUpdated())); err != nil {
		return err
	}

	return nil
}

func (m *PingCheck) contextValidateProperties(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "properties", "body", []*NameAndValue(m.Properties())); err != nil {
		return err
	}

	for i := 0; i < len(m.Properties()); i++ {

		if m.propertiesField[i] != nil {
			if err := m.propertiesField[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("properties" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PingCheck) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "status", "body", string(m.Status())); err != nil {
		return err
	}

	return nil
}

func (m *PingCheck) contextValidateStopMonitoringByFolder(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "stopMonitoringByFolder", "body", m.StopMonitoringByFolder()); err != nil {
		return err
	}

	return nil
}

func (m *PingCheck) contextValidateTestLocation(ctx context.Context, formats strfmt.Registry) error {

	if m.TestLocation() != nil {
		if err := m.TestLocation().ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("testLocation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PingCheck) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PingCheck) UnmarshalBinary(b []byte) error {
	var res PingCheck
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
