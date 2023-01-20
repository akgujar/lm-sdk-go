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

// CollectorAttribute collector attribute
//
// swagger:discriminator CollectorAttribute name
type CollectorAttribute interface {
	runtime.Validatable
	runtime.ContextValidatable

	// The data collector's name
	// Required: true
	Name() string
	SetName(string)

	// AdditionalProperties in base type shoud be handled just like regular properties
	// At this moment, the base type property is pushed down to the subtype
}

type collectorAttribute struct {
	nameField string
}

// Name gets the name of this polymorphic type
func (m *collectorAttribute) Name() string {
	return "CollectorAttribute"
}

// SetName sets the name of this polymorphic type
func (m *collectorAttribute) SetName(val string) {
}

// UnmarshalCollectorAttributeSlice unmarshals polymorphic slices of CollectorAttribute
func UnmarshalCollectorAttributeSlice(reader io.Reader, consumer runtime.Consumer) ([]CollectorAttribute, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []CollectorAttribute
	for _, element := range elements {
		obj, err := unmarshalCollectorAttribute(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalCollectorAttribute unmarshals polymorphic CollectorAttribute
func UnmarshalCollectorAttribute(reader io.Reader, consumer runtime.Consumer) (CollectorAttribute, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalCollectorAttribute(data, consumer)
}

func unmarshalCollectorAttribute(data []byte, consumer runtime.Consumer) (CollectorAttribute, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the name property.
	var getType struct {
		Name string `json:"name"`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("name", "body", getType.Name); err != nil {
		return nil, err
	}

	// The value of name is used to determine which type to create and unmarshal the data into
	switch getType.Name {
	case "AggregateCollectorAttribute":
		var result AggregateCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "AwsEbsVolumeSnapshotCollectorAttributeV3":
		var result AwsEbsVolumeSnapshotCollectorAttributeV3
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "AwsRdsPerformanceInsightsCollectorAttribute":
		var result AwsRdsPerformanceInsightsCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "AwsRdsServiceLimitsCollectorAttributeV3":
		var result AwsRdsServiceLimitsCollectorAttributeV3
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "AzureActiveDirectoryAppSecretCollectorAttribute":
		var result AzureActiveDirectoryAppSecretCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "AzureActiveDirectorySyncCollectorAttribute":
		var result AzureActiveDirectorySyncCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "AzureActiveDirectoryUsersCollectorAttribute":
		var result AzureActiveDirectoryUsersCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "AzureAppServiceEnvironmentMultiRolePoolCollectorAttributeV3":
		var result AzureAppServiceEnvironmentMultiRolePoolCollectorAttributeV3
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "AzureAutomationAccountCertificateCollectorAttributeV3":
		var result AzureAutomationAccountCertificateCollectorAttributeV3
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "AzureBackupJobCollectorAttributeV3":
		var result AzureBackupJobCollectorAttributeV3
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "AzureBackupProtectedItemBackupJobCollectorAttribute":
		var result AzureBackupProtectedItemBackupJobCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "AzureBackupProtectedItemHealthCollectorAttribute":
		var result AzureBackupProtectedItemHealthCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "AzureCostManagementCollectorAttribute":
		var result AzureCostManagementCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "AzureEABillingCollectorAttribute":
		var result AzureEABillingCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "AzureExpressRouteCircuitPeeringCollectorAttributeV3":
		var result AzureExpressRouteCircuitPeeringCollectorAttributeV3
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "AzureLogAnalyticsReplicationJobCollectorAttribute":
		var result AzureLogAnalyticsReplicationJobCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "AzureLogAnalyticsWorkspacesCollectorAttribute":
		var result AzureLogAnalyticsWorkspacesCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "AzureRecoveryServiceRTOCollectorAttribute":
		var result AzureRecoveryServiceRTOCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "AzureRecoveryServiceVaultSRCollectorAttribute":
		var result AzureRecoveryServiceVaultSRCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "AzureRecoveryServicesVaultAgentsCollectorAttribute":
		var result AzureRecoveryServicesVaultAgentsCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "AzureReplicationDisasterRecoveryCollectorAttribute":
		var result AzureReplicationDisasterRecoveryCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "AzureReplicationJobCollectorAttributeV3":
		var result AzureReplicationJobCollectorAttributeV3
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "AzureVMBackupStatusCollectorAttributeV3":
		var result AzureVMBackupStatusCollectorAttributeV3
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "AzureVMBackupStatusLogAnalyticsCollectorAttribute":
		var result AzureVMBackupStatusLogAnalyticsCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "AzureVNGConnectionCollectorAttributeV3":
		var result AzureVNGConnectionCollectorAttributeV3
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "AzureVirtualDesktopHostPoolsCollectorAttributeV3":
		var result AzureVirtualDesktopHostPoolsCollectorAttributeV3
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "AzureVirtualDesktopSessionHostsCollectorAttributeV3":
		var result AzureVirtualDesktopSessionHostsCollectorAttributeV3
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "AzureWebJobCollectorAttributeV3":
		var result AzureWebJobCollectorAttributeV3
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "CollectorAttribute":
		var result collectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "GcpBackendServiceHealthCollectorAttributeV3":
		var result GcpBackendServiceHealthCollectorAttributeV3
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "GcpBillingBigQuerySourceCollectorAttribute":
		var result GcpBillingBigQuerySourceCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "GcpBillingCollectorAttribute":
		var result GcpBillingCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "GcpComputeServiceLimitsCollectorAttributeV3":
		var result GcpComputeServiceLimitsCollectorAttributeV3
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "GcpStackDriverCollectorAttributeV3":
		var result GcpStackDriverCollectorAttributeV3
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "NetflowAlertModulesCollectorAttributeV3":
		var result NetflowAlertModulesCollectorAttributeV3
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "OpenMetricCollectorAttributeV3":
		var result OpenMetricCollectorAttributeV3
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "PaaSMongoDbCollectorAttribute":
		var result PaaSMongoDbCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "PushModulesCollectorAttribute":
		var result PushModulesCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "SaasAirbrakeCollectorAttribute":
		var result SaasAirbrakeCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "SaasOffice365CsvReportCollectorAttributeV3":
		var result SaasOffice365CsvReportCollectorAttributeV3
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "SaasOffice365HealthCollectorAttributeV3":
		var result SaasOffice365HealthCollectorAttributeV3
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "SaasOffice365SharepointReportCollectorAttributeV3":
		var result SaasOffice365SharepointReportCollectorAttributeV3
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "SaasOffice365SkusCollectorAttribute":
		var result SaasOffice365SkusCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "SaasOffice365TeamsCallsQosCollectorAttributeV3":
		var result SaasOffice365TeamsCallsQosCollectorAttributeV3
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "SaasSalesforceInstanceStatusCollectorAttributeV3":
		var result SaasSalesforceInstanceStatusCollectorAttributeV3
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "SaasSalesforceJsonCollectorAttributeV3":
		var result SaasSalesforceJSONCollectorAttributeV3
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "SaasSalesforceSOQLQueryCollectorAttributeV3":
		var result SaasSalesforceSOQLQueryCollectorAttributeV3
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "SaasSlackHealthCollectorAttribute":
		var result SaasSlackHealthCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "SaasStatusCollectorAttribute":
		var result SaasStatusCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "SaasWebexCollectorAttribute":
		var result SaasWebexCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "SaasZoomJsonCollectorAttributeV3":
		var result SaasZoomJSONCollectorAttributeV3
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "SaasZoomPlanUsageCollectorAttributeV3":
		var result SaasZoomPlanUsageCollectorAttributeV3
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "SyntheticsSeleniumCollectorAttribute":
		var result SyntheticsSeleniumCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "awsautoscalingservicelimits":
		var result AwsAutoScalingServiceLimitsCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "awsbilling":
		var result AwsBillingCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "awsbillingreport":
		var result AwsBillingReportCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "awsclassicelbservicelimits":
		var result AwsClassicElbServiceLimitsCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "awscloudwatch":
		var result AwsCloudWatchCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "awsdynamodb":
		var result AwsDynamodbCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "awsec2reservedinstance":
		var result AwsEC2ReservedInstanceCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "awsec2reservedinstancecoverage":
		var result AwsEC2ReservedInstanceCoverageCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "awsec2scheduledevents":
		var result AwsEC2ScheduledEventsCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "awsec2servicelimits":
		var result AwsEc2ServiceLimitsCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "awsecsservicedetails":
		var result AwsEcsServiceDetailsCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "awss3":
		var result AwsS3CollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "awsservicelimitsfromtrustedadvisor":
		var result AwsServiceLimitsFromTrustedAdvisorCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "awssesservicelimits":
		var result AwsSesServiceLimitsCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "awssqs":
		var result AwsSqsCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "azurebilling":
		var result AzureBillingCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "azureinsights":
		var result AzureInsightsCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "azurenetworkservicelimits":
		var result AzureNetworkServiceLimitsCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "azureresourcehealth":
		var result AzureResourceHealthCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "azurestorageservicelimits":
		var result AzureStorageServiceLimitsCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "azurevmservicelimits":
		var result AzureVMServiceLimitsCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "batchscript":
		var result BatchScriptCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "cim":
		var result CIMCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "datapump":
		var result DataPumpCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "dns":
		var result DNSCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "esx":
		var result ESXCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "internal":
		var result InternalCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "ipmi":
		var result IPMICollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "jdbc":
		var result JDBCCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "jmx":
		var result JMXCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "memcached":
		var result MemcachedCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "netapp":
		var result NetAppCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "perfmon":
		var result PerfmonCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "ping":
		var result PingCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "script":
		var result ScriptCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "sdkscript":
		var result SDKScriptCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "snmp":
		var result SNMPCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "tcp":
		var result TCPCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "udp":
		var result UDPCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "webpage":
		var result WebPageCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "wmi":
		var result WMICollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "xen":
		var result XENCollectorAttribute
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	}
	return nil, errors.New(422, "invalid name value: %q", getType.Name)
}

// Validate validates this collector attribute
func (m *collectorAttribute) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this collector attribute based on context it is used
func (m *collectorAttribute) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
