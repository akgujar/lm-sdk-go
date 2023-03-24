// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ConfigSource config source
//
// swagger:model ConfigSource
type ConfigSource struct {

	// The Applies To for the LMModule
	AppliesTo string `json:"appliesTo,omitempty"`

	// The ConfigSource audit version
	AuditVersion int64 `json:"auditVersion,omitempty"`

	// Auto discovery configuration
	AutoDiscoveryConfig *AutoDiscoveryConfiguration `json:"autoDiscoveryConfig,omitempty"`

	// The metadata checksum for the LMModule content
	// Read Only: true
	Checksum string `json:"checksum,omitempty"`

	// The ConfigSource data collect interval
	CollectInterval int32 `json:"collectInterval,omitempty"`

	// The method to collect data
	CollectMethod string `json:"collectMethod,omitempty"`

	collectorAttributeField CollectorAttribute

	// The List of ConfigChecks
	ConfigChecks []*ConfigCheck `json:"configChecks,omitempty"`

	// The description for the LMModule
	Description string `json:"description,omitempty"`

	// The ConfigSource display name
	DisplayName string `json:"displayName,omitempty"`

	// Enable active discovery if ConfigSource has multiple instances. true|false
	EnableAutoDiscovery bool `json:"enableAutoDiscovery,omitempty"`

	// Configuration file format. The values can be arbitrary|unix|java-properties|JSON|XML
	FileFormat string `json:"fileFormat,omitempty"`

	// The group the LMModule is in
	Group string `json:"group,omitempty"`

	// Whether the ConfigSource has multiple instances. true|false
	HasMultiInstances bool `json:"hasMultiInstances,omitempty"`

	// The ID of the LMModule
	// Read Only: true
	ID int32 `json:"id,omitempty"`

	// The local module's IntegrationMetadata, readable for troubleshooting purposes
	// Read Only: true
	InstallationMetadata *IntegrationMetadata `json:"installationMetadata,omitempty"`

	// The lineageId the LMModule belongs to
	// Read Only: true
	LineageID string `json:"lineageId,omitempty"`

	// The config source name
	Name string `json:"name,omitempty"`

	// The Tags for the LMModule
	Tags string `json:"tags,omitempty"`

	// The Technical Notes for the LMModule
	Technology string `json:"technology,omitempty"`

	// Timestamp format. ex. yyyy-MM-dd hh:mm:ss
	TimestampFormat string `json:"timestampFormat,omitempty"`

	// The ConfigSource version
	Version int64 `json:"version,omitempty"`
}

// CollectorAttribute gets the collector attribute of this base type
func (m *ConfigSource) CollectorAttribute() CollectorAttribute {
	return m.collectorAttributeField
}

// SetCollectorAttribute sets the collector attribute of this base type
func (m *ConfigSource) SetCollectorAttribute(val CollectorAttribute) {
	m.collectorAttributeField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *ConfigSource) UnmarshalJSON(raw []byte) error {
	var data struct {
		AppliesTo string `json:"appliesTo,omitempty"`

		AuditVersion int64 `json:"auditVersion,omitempty"`

		AutoDiscoveryConfig *AutoDiscoveryConfiguration `json:"autoDiscoveryConfig,omitempty"`

		Checksum string `json:"checksum,omitempty"`

		CollectInterval int32 `json:"collectInterval,omitempty"`

		CollectMethod string `json:"collectMethod,omitempty"`

		CollectorAttribute json.RawMessage `json:"collectorAttribute,omitempty"`

		ConfigChecks []*ConfigCheck `json:"configChecks,omitempty"`

		Description string `json:"description,omitempty"`

		DisplayName string `json:"displayName,omitempty"`

		EnableAutoDiscovery bool `json:"enableAutoDiscovery,omitempty"`

		FileFormat string `json:"fileFormat,omitempty"`

		Group string `json:"group,omitempty"`

		HasMultiInstances bool `json:"hasMultiInstances,omitempty"`

		ID int32 `json:"id,omitempty"`

		InstallationMetadata *IntegrationMetadata `json:"installationMetadata,omitempty"`

		LineageID string `json:"lineageId,omitempty"`

		Name string `json:"name,omitempty"`

		Tags string `json:"tags,omitempty"`

		Technology string `json:"technology,omitempty"`

		TimestampFormat string `json:"timestampFormat,omitempty"`

		Version int64 `json:"version,omitempty"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var propCollectorAttribute CollectorAttribute
	if string(data.CollectorAttribute) != "null" {
		collectorAttribute, err := UnmarshalCollectorAttribute(bytes.NewBuffer(data.CollectorAttribute), runtime.JSONConsumer())
		if err != nil && err != io.EOF {
			return err
		}
		propCollectorAttribute = collectorAttribute
	}

	var result ConfigSource

	// appliesTo
	result.AppliesTo = data.AppliesTo

	// auditVersion
	result.AuditVersion = data.AuditVersion

	// autoDiscoveryConfig
	result.AutoDiscoveryConfig = data.AutoDiscoveryConfig

	// checksum
	result.Checksum = data.Checksum

	// collectInterval
	result.CollectInterval = data.CollectInterval

	// collectMethod
	result.CollectMethod = data.CollectMethod

	// collectorAttribute
	result.collectorAttributeField = propCollectorAttribute

	// configChecks
	result.ConfigChecks = data.ConfigChecks

	// description
	result.Description = data.Description

	// displayName
	result.DisplayName = data.DisplayName

	// enableAutoDiscovery
	result.EnableAutoDiscovery = data.EnableAutoDiscovery

	// fileFormat
	result.FileFormat = data.FileFormat

	// group
	result.Group = data.Group

	// hasMultiInstances
	result.HasMultiInstances = data.HasMultiInstances

	// id
	result.ID = data.ID

	// installationMetadata
	result.InstallationMetadata = data.InstallationMetadata

	// lineageId
	result.LineageID = data.LineageID

	// name
	result.Name = data.Name

	// tags
	result.Tags = data.Tags

	// technology
	result.Technology = data.Technology

	// timestampFormat
	result.TimestampFormat = data.TimestampFormat

	// version
	result.Version = data.Version

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m ConfigSource) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		AppliesTo string `json:"appliesTo,omitempty"`

		AuditVersion int64 `json:"auditVersion,omitempty"`

		AutoDiscoveryConfig *AutoDiscoveryConfiguration `json:"autoDiscoveryConfig,omitempty"`

		Checksum string `json:"checksum,omitempty"`

		CollectInterval int32 `json:"collectInterval,omitempty"`

		CollectMethod string `json:"collectMethod,omitempty"`

		ConfigChecks []*ConfigCheck `json:"configChecks,omitempty"`

		Description string `json:"description,omitempty"`

		DisplayName string `json:"displayName,omitempty"`

		EnableAutoDiscovery bool `json:"enableAutoDiscovery,omitempty"`

		FileFormat string `json:"fileFormat,omitempty"`

		Group string `json:"group,omitempty"`

		HasMultiInstances bool `json:"hasMultiInstances,omitempty"`

		ID int32 `json:"id,omitempty"`

		InstallationMetadata *IntegrationMetadata `json:"installationMetadata,omitempty"`

		LineageID string `json:"lineageId,omitempty"`

		Name string `json:"name,omitempty"`

		Tags string `json:"tags,omitempty"`

		Technology string `json:"technology,omitempty"`

		TimestampFormat string `json:"timestampFormat,omitempty"`

		Version int64 `json:"version,omitempty"`
	}{

		AppliesTo: m.AppliesTo,

		AuditVersion: m.AuditVersion,

		AutoDiscoveryConfig: m.AutoDiscoveryConfig,

		Checksum: m.Checksum,

		CollectInterval: m.CollectInterval,

		CollectMethod: m.CollectMethod,

		ConfigChecks: m.ConfigChecks,

		Description: m.Description,

		DisplayName: m.DisplayName,

		EnableAutoDiscovery: m.EnableAutoDiscovery,

		FileFormat: m.FileFormat,

		Group: m.Group,

		HasMultiInstances: m.HasMultiInstances,

		ID: m.ID,

		InstallationMetadata: m.InstallationMetadata,

		LineageID: m.LineageID,

		Name: m.Name,

		Tags: m.Tags,

		Technology: m.Technology,

		TimestampFormat: m.TimestampFormat,

		Version: m.Version,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		CollectorAttribute CollectorAttribute `json:"collectorAttribute,omitempty"`
	}{

		CollectorAttribute: m.collectorAttributeField,
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this config source
func (m *ConfigSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAutoDiscoveryConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCollectorAttribute(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConfigChecks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstallationMetadata(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigSource) validateAutoDiscoveryConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.AutoDiscoveryConfig) { // not required
		return nil
	}

	if m.AutoDiscoveryConfig != nil {
		if err := m.AutoDiscoveryConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("autoDiscoveryConfig")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigSource) validateCollectorAttribute(formats strfmt.Registry) error {
	if swag.IsZero(m.CollectorAttribute()) { // not required
		return nil
	}

	if err := m.CollectorAttribute().Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("collectorAttribute")
		}
		return err
	}

	return nil
}

func (m *ConfigSource) validateConfigChecks(formats strfmt.Registry) error {
	if swag.IsZero(m.ConfigChecks) { // not required
		return nil
	}

	for i := 0; i < len(m.ConfigChecks); i++ {
		if swag.IsZero(m.ConfigChecks[i]) { // not required
			continue
		}

		if m.ConfigChecks[i] != nil {
			if err := m.ConfigChecks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("configChecks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConfigSource) validateInstallationMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.InstallationMetadata) { // not required
		return nil
	}

	if m.InstallationMetadata != nil {
		if err := m.InstallationMetadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("installationMetadata")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this config source based on the context it is used
func (m *ConfigSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAutoDiscoveryConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateChecksum(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCollectorAttribute(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConfigChecks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInstallationMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLineageID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigSource) contextValidateAutoDiscoveryConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.AutoDiscoveryConfig != nil {
		if err := m.AutoDiscoveryConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("autoDiscoveryConfig")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigSource) contextValidateChecksum(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "checksum", "body", string(m.Checksum)); err != nil {
		return err
	}

	return nil
}

func (m *ConfigSource) contextValidateCollectorAttribute(ctx context.Context, formats strfmt.Registry) error {

	if err := m.CollectorAttribute().ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("collectorAttribute")
		}
		return err
	}

	return nil
}

func (m *ConfigSource) contextValidateConfigChecks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ConfigChecks); i++ {

		if m.ConfigChecks[i] != nil {
			if err := m.ConfigChecks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("configChecks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConfigSource) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int32(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *ConfigSource) contextValidateInstallationMetadata(ctx context.Context, formats strfmt.Registry) error {

	if m.InstallationMetadata != nil {
		if err := m.InstallationMetadata.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("installationMetadata")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigSource) contextValidateLineageID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "lineageId", "body", string(m.LineageID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigSource) UnmarshalBinary(b []byte) error {
	var res ConfigSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}