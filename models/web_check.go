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

// WebCheck web check
//
// swagger:model WebCheck
type WebCheck struct {
	checkpointsField []*WebsiteCheckPoint

	collectorsField []*WebsiteCollectorInfo

	descriptionField string

	disableAlertingField bool

	domainField string

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

	rolePrivilegesField []string

	statusField string

	stepsField []*WebCheckStep

	stopMonitoringField bool

	stopMonitoringByFolderField *bool

	templateField interface{}

	testLocationField *WebsiteLocation

	transitionField int32

	useDefaultAlertSettingField bool

	useDefaultLocationSettingField bool

	userPermissionField string

	// The threshold (in days) for triggering SSL certification alerts
	// Example: \u003c 200 100 50
	AlertExpr string `json:"alertExpr,omitempty"`

	// Whether or not SSL should be ignored, the default value is true
	// Example: true
	IgnoreSSL interface{} `json:"ignoreSSL,omitempty"`

	// The time in milliseconds that the page must load within for each step to avoid triggering an alert
	// Example: 30000
	PageLoadAlertTimeInMS int64 `json:"pageLoadAlertTimeInMS,omitempty"`

	// The scheme or protocol associated with the URL to check. Acceptable values are: http, https
	// Example: https
	Schema string `json:"schema,omitempty"`

	// Whether or not SSL expiration alerts should be triggered
	// Example: false
	TriggerSSLExpirationAlert bool `json:"triggerSSLExpirationAlert,omitempty"`

	// Whether or not SSL status alerts should be triggered
	// Example: false
	TriggerSSLStatusAlert bool `json:"triggerSSLStatusAlert,omitempty"`
}

// Checkpoints gets the checkpoints of this subtype
func (m *WebCheck) Checkpoints() []*WebsiteCheckPoint {
	return m.checkpointsField
}

// SetCheckpoints sets the checkpoints of this subtype
func (m *WebCheck) SetCheckpoints(val []*WebsiteCheckPoint) {
	m.checkpointsField = val
}

// Collectors gets the collectors of this subtype
func (m *WebCheck) Collectors() []*WebsiteCollectorInfo {
	return m.collectorsField
}

// SetCollectors sets the collectors of this subtype
func (m *WebCheck) SetCollectors(val []*WebsiteCollectorInfo) {
	m.collectorsField = val
}

// Description gets the description of this subtype
func (m *WebCheck) Description() string {
	return m.descriptionField
}

// SetDescription sets the description of this subtype
func (m *WebCheck) SetDescription(val string) {
	m.descriptionField = val
}

// DisableAlerting gets the disable alerting of this subtype
func (m *WebCheck) DisableAlerting() bool {
	return m.disableAlertingField
}

// SetDisableAlerting sets the disable alerting of this subtype
func (m *WebCheck) SetDisableAlerting(val bool) {
	m.disableAlertingField = val
}

// Domain gets the domain of this subtype
func (m *WebCheck) Domain() string {
	return m.domainField
}

// SetDomain sets the domain of this subtype
func (m *WebCheck) SetDomain(val string) {
	m.domainField = val
}

// GlobalSmAlertCond gets the global sm alert cond of this subtype
func (m *WebCheck) GlobalSmAlertCond() int32 {
	return m.globalSmAlertCondField
}

// SetGlobalSmAlertCond sets the global sm alert cond of this subtype
func (m *WebCheck) SetGlobalSmAlertCond(val int32) {
	m.globalSmAlertCondField = val
}

// GroupID gets the group Id of this subtype
func (m *WebCheck) GroupID() int32 {
	return m.groupIdField
}

// SetGroupID sets the group Id of this subtype
func (m *WebCheck) SetGroupID(val int32) {
	m.groupIdField = val
}

// ID gets the id of this subtype
func (m *WebCheck) ID() int32 {
	return m.idField
}

// SetID sets the id of this subtype
func (m *WebCheck) SetID(val int32) {
	m.idField = val
}

// IndividualAlertLevel gets the individual alert level of this subtype
func (m *WebCheck) IndividualAlertLevel() string {
	return m.individualAlertLevelField
}

// SetIndividualAlertLevel sets the individual alert level of this subtype
func (m *WebCheck) SetIndividualAlertLevel(val string) {
	m.individualAlertLevelField = val
}

// IndividualSmAlertEnable gets the individual sm alert enable of this subtype
func (m *WebCheck) IndividualSmAlertEnable() bool {
	return m.individualSmAlertEnableField
}

// SetIndividualSmAlertEnable sets the individual sm alert enable of this subtype
func (m *WebCheck) SetIndividualSmAlertEnable(val bool) {
	m.individualSmAlertEnableField = val
}

// IsInternal gets the is internal of this subtype
func (m *WebCheck) IsInternal() bool {
	return m.isInternalField
}

// SetIsInternal sets the is internal of this subtype
func (m *WebCheck) SetIsInternal(val bool) {
	m.isInternalField = val
}

// LastUpdated gets the last updated of this subtype
func (m *WebCheck) LastUpdated() int64 {
	return m.lastUpdatedField
}

// SetLastUpdated sets the last updated of this subtype
func (m *WebCheck) SetLastUpdated(val int64) {
	m.lastUpdatedField = val
}

// Name gets the name of this subtype
func (m *WebCheck) Name() *string {
	return m.nameField
}

// SetName sets the name of this subtype
func (m *WebCheck) SetName(val *string) {
	m.nameField = val
}

// OverallAlertLevel gets the overall alert level of this subtype
func (m *WebCheck) OverallAlertLevel() string {
	return m.overallAlertLevelField
}

// SetOverallAlertLevel sets the overall alert level of this subtype
func (m *WebCheck) SetOverallAlertLevel(val string) {
	m.overallAlertLevelField = val
}

// PollingInterval gets the polling interval of this subtype
func (m *WebCheck) PollingInterval() int32 {
	return m.pollingIntervalField
}

// SetPollingInterval sets the polling interval of this subtype
func (m *WebCheck) SetPollingInterval(val int32) {
	m.pollingIntervalField = val
}

// Properties gets the properties of this subtype
func (m *WebCheck) Properties() []*NameAndValue {
	return m.propertiesField
}

// SetProperties sets the properties of this subtype
func (m *WebCheck) SetProperties(val []*NameAndValue) {
	m.propertiesField = val
}

// RolePrivileges gets the role privileges of this subtype
func (m *WebCheck) RolePrivileges() []string {
	return m.rolePrivilegesField
}

// SetRolePrivileges sets the role privileges of this subtype
func (m *WebCheck) SetRolePrivileges(val []string) {
	m.rolePrivilegesField = val
}

// Status gets the status of this subtype
func (m *WebCheck) Status() string {
	return m.statusField
}

// SetStatus sets the status of this subtype
func (m *WebCheck) SetStatus(val string) {
	m.statusField = val
}

// Steps gets the steps of this subtype
func (m *WebCheck) Steps() []*WebCheckStep {
	return m.stepsField
}

// SetSteps sets the steps of this subtype
func (m *WebCheck) SetSteps(val []*WebCheckStep) {
	m.stepsField = val
}

// StopMonitoring gets the stop monitoring of this subtype
func (m *WebCheck) StopMonitoring() bool {
	return m.stopMonitoringField
}

// SetStopMonitoring sets the stop monitoring of this subtype
func (m *WebCheck) SetStopMonitoring(val bool) {
	m.stopMonitoringField = val
}

// StopMonitoringByFolder gets the stop monitoring by folder of this subtype
func (m *WebCheck) StopMonitoringByFolder() *bool {
	return m.stopMonitoringByFolderField
}

// SetStopMonitoringByFolder sets the stop monitoring by folder of this subtype
func (m *WebCheck) SetStopMonitoringByFolder(val *bool) {
	m.stopMonitoringByFolderField = val
}

// Template gets the template of this subtype
func (m *WebCheck) Template() interface{} {
	return m.templateField
}

// SetTemplate sets the template of this subtype
func (m *WebCheck) SetTemplate(val interface{}) {
	m.templateField = val
}

// TestLocation gets the test location of this subtype
func (m *WebCheck) TestLocation() *WebsiteLocation {
	return m.testLocationField
}

// SetTestLocation sets the test location of this subtype
func (m *WebCheck) SetTestLocation(val *WebsiteLocation) {
	m.testLocationField = val
}

// Transition gets the transition of this subtype
func (m *WebCheck) Transition() int32 {
	return m.transitionField
}

// SetTransition sets the transition of this subtype
func (m *WebCheck) SetTransition(val int32) {
	m.transitionField = val
}

// Type gets the type of this subtype
func (m *WebCheck) Type() string {
	return "webcheck"
}

// SetType sets the type of this subtype
func (m *WebCheck) SetType(val string) {
}

// UseDefaultAlertSetting gets the use default alert setting of this subtype
func (m *WebCheck) UseDefaultAlertSetting() bool {
	return m.useDefaultAlertSettingField
}

// SetUseDefaultAlertSetting sets the use default alert setting of this subtype
func (m *WebCheck) SetUseDefaultAlertSetting(val bool) {
	m.useDefaultAlertSettingField = val
}

// UseDefaultLocationSetting gets the use default location setting of this subtype
func (m *WebCheck) UseDefaultLocationSetting() bool {
	return m.useDefaultLocationSettingField
}

// SetUseDefaultLocationSetting sets the use default location setting of this subtype
func (m *WebCheck) SetUseDefaultLocationSetting(val bool) {
	m.useDefaultLocationSettingField = val
}

// UserPermission gets the user permission of this subtype
func (m *WebCheck) UserPermission() string {
	return m.userPermissionField
}

// SetUserPermission sets the user permission of this subtype
func (m *WebCheck) SetUserPermission(val string) {
	m.userPermissionField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *WebCheck) UnmarshalJSON(raw []byte) error {
	var data struct {

		// The threshold (in days) for triggering SSL certification alerts
		// Example: \u003c 200 100 50
		AlertExpr string `json:"alertExpr,omitempty"`

		// Whether or not SSL should be ignored, the default value is true
		// Example: true
		IgnoreSSL interface{} `json:"ignoreSSL,omitempty"`

		// The time in milliseconds that the page must load within for each step to avoid triggering an alert
		// Example: 30000
		PageLoadAlertTimeInMS int64 `json:"pageLoadAlertTimeInMS,omitempty"`

		// The scheme or protocol associated with the URL to check. Acceptable values are: http, https
		// Example: https
		Schema string `json:"schema,omitempty"`

		// Whether or not SSL expiration alerts should be triggered
		// Example: false
		TriggerSSLExpirationAlert bool `json:"triggerSSLExpirationAlert,omitempty"`

		// Whether or not SSL status alerts should be triggered
		// Example: false
		TriggerSSLStatusAlert bool `json:"triggerSSLStatusAlert,omitempty"`
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

		Domain string `json:"domain,omitempty"`

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

		RolePrivileges []string `json:"rolePrivileges,omitempty"`

		Status string `json:"status,omitempty"`

		Steps []*WebCheckStep `json:"steps,omitempty"`

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

	var result WebCheck

	result.checkpointsField = base.Checkpoints

	result.collectorsField = base.Collectors

	result.descriptionField = base.Description

	result.disableAlertingField = base.DisableAlerting

	result.domainField = base.Domain

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

	result.rolePrivilegesField = base.RolePrivileges

	result.statusField = base.Status

	result.stepsField = base.Steps

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

	result.AlertExpr = data.AlertExpr
	result.IgnoreSSL = data.IgnoreSSL
	result.PageLoadAlertTimeInMS = data.PageLoadAlertTimeInMS
	result.Schema = data.Schema
	result.TriggerSSLExpirationAlert = data.TriggerSSLExpirationAlert
	result.TriggerSSLStatusAlert = data.TriggerSSLStatusAlert

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m WebCheck) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// The threshold (in days) for triggering SSL certification alerts
		// Example: \u003c 200 100 50
		AlertExpr string `json:"alertExpr,omitempty"`

		// Whether or not SSL should be ignored, the default value is true
		// Example: true
		IgnoreSSL interface{} `json:"ignoreSSL,omitempty"`

		// The time in milliseconds that the page must load within for each step to avoid triggering an alert
		// Example: 30000
		PageLoadAlertTimeInMS int64 `json:"pageLoadAlertTimeInMS,omitempty"`

		// The scheme or protocol associated with the URL to check. Acceptable values are: http, https
		// Example: https
		Schema string `json:"schema,omitempty"`

		// Whether or not SSL expiration alerts should be triggered
		// Example: false
		TriggerSSLExpirationAlert bool `json:"triggerSSLExpirationAlert,omitempty"`

		// Whether or not SSL status alerts should be triggered
		// Example: false
		TriggerSSLStatusAlert bool `json:"triggerSSLStatusAlert,omitempty"`
	}{

		AlertExpr: m.AlertExpr,

		IgnoreSSL: m.IgnoreSSL,

		PageLoadAlertTimeInMS: m.PageLoadAlertTimeInMS,

		Schema: m.Schema,

		TriggerSSLExpirationAlert: m.TriggerSSLExpirationAlert,

		TriggerSSLStatusAlert: m.TriggerSSLStatusAlert,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Checkpoints []*WebsiteCheckPoint `json:"checkpoints,omitempty"`

		Collectors []*WebsiteCollectorInfo `json:"collectors,omitempty"`

		Description string `json:"description,omitempty"`

		DisableAlerting bool `json:"disableAlerting,omitempty"`

		Domain string `json:"domain,omitempty"`

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

		RolePrivileges []string `json:"rolePrivileges,omitempty"`

		Status string `json:"status,omitempty"`

		Steps []*WebCheckStep `json:"steps,omitempty"`

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

		Domain: m.Domain(),

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

		RolePrivileges: m.RolePrivileges(),

		Status: m.Status(),

		Steps: m.Steps(),

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

// Validate validates this web check
func (m *WebCheck) Validate(formats strfmt.Registry) error {
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

	if err := m.validateSteps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTestLocation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebCheck) validateCheckpoints(formats strfmt.Registry) error {

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

func (m *WebCheck) validateCollectors(formats strfmt.Registry) error {

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

func (m *WebCheck) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name()); err != nil {
		return err
	}

	return nil
}

func (m *WebCheck) validateProperties(formats strfmt.Registry) error {

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

func (m *WebCheck) validateSteps(formats strfmt.Registry) error {

	if swag.IsZero(m.Steps()) { // not required
		return nil
	}

	for i := 0; i < len(m.Steps()); i++ {
		if swag.IsZero(m.stepsField[i]) { // not required
			continue
		}

		if m.stepsField[i] != nil {
			if err := m.stepsField[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("steps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WebCheck) validateTestLocation(formats strfmt.Registry) error {

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

// ContextValidate validate this web check based on the context it is used
func (m *WebCheck) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if err := m.contextValidateRolePrivileges(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSteps(ctx, formats); err != nil {
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

func (m *WebCheck) contextValidateCheckpoints(ctx context.Context, formats strfmt.Registry) error {

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

func (m *WebCheck) contextValidateCollectors(ctx context.Context, formats strfmt.Registry) error {

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

func (m *WebCheck) contextValidateGroupID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "groupId", "body", int32(m.GroupID())); err != nil {
		return err
	}

	return nil
}

func (m *WebCheck) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int32(m.ID())); err != nil {
		return err
	}

	return nil
}

func (m *WebCheck) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "lastUpdated", "body", int64(m.LastUpdated())); err != nil {
		return err
	}

	return nil
}

func (m *WebCheck) contextValidateProperties(ctx context.Context, formats strfmt.Registry) error {

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

func (m *WebCheck) contextValidateRolePrivileges(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "rolePrivileges", "body", []string(m.RolePrivileges())); err != nil {
		return err
	}

	return nil
}

func (m *WebCheck) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "status", "body", string(m.Status())); err != nil {
		return err
	}

	return nil
}

func (m *WebCheck) contextValidateSteps(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Steps()); i++ {

		if m.stepsField[i] != nil {
			if err := m.stepsField[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("steps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WebCheck) contextValidateStopMonitoringByFolder(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "stopMonitoringByFolder", "body", m.StopMonitoringByFolder()); err != nil {
		return err
	}

	return nil
}

func (m *WebCheck) contextValidateTestLocation(ctx context.Context, formats strfmt.Registry) error {

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
func (m *WebCheck) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebCheck) UnmarshalBinary(b []byte) error {
	var res WebCheck
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
