package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementable 
type DeviceManagementable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetApplePushNotificationCertificate()(ApplePushNotificationCertificateable)
    GetComplianceManagementPartners()([]ComplianceManagementPartnerable)
    GetConditionalAccessSettings()(OnPremisesConditionalAccessSettingsable)
    GetDetectedApps()([]DetectedAppable)
    GetDeviceCategories()([]DeviceCategoryable)
    GetDeviceCompliancePolicies()([]DeviceCompliancePolicyable)
    GetDeviceCompliancePolicyDeviceStateSummary()(DeviceCompliancePolicyDeviceStateSummaryable)
    GetDeviceCompliancePolicySettingStateSummaries()([]DeviceCompliancePolicySettingStateSummaryable)
    GetDeviceConfigurationDeviceStateSummaries()(DeviceConfigurationDeviceStateSummaryable)
    GetDeviceConfigurations()([]DeviceConfigurationable)
    GetDeviceEnrollmentConfigurations()([]DeviceEnrollmentConfigurationable)
    GetDeviceManagementPartners()([]DeviceManagementPartnerable)
    GetExchangeConnectors()([]DeviceManagementExchangeConnectorable)
    GetImportedWindowsAutopilotDeviceIdentities()([]ImportedWindowsAutopilotDeviceIdentityable)
    GetIntuneAccountId()(*string)
    GetIntuneBrand()(IntuneBrandable)
    GetIosUpdateStatuses()([]IosUpdateDeviceStatusable)
    GetManagedDeviceOverview()(ManagedDeviceOverviewable)
    GetManagedDevices()([]ManagedDeviceable)
    GetMobileThreatDefenseConnectors()([]MobileThreatDefenseConnectorable)
    GetNotificationMessageTemplates()([]NotificationMessageTemplateable)
    GetRemoteAssistancePartners()([]RemoteAssistancePartnerable)
    GetReports()(DeviceManagementReportsable)
    GetResourceOperations()([]ResourceOperationable)
    GetRoleAssignments()([]DeviceAndAppManagementRoleAssignmentable)
    GetRoleDefinitions()([]RoleDefinitionable)
    GetSettings()(DeviceManagementSettingsable)
    GetSoftwareUpdateStatusSummary()(SoftwareUpdateStatusSummaryable)
    GetSubscriptionState()(*DeviceManagementSubscriptionState)
    GetTelecomExpenseManagementPartners()([]TelecomExpenseManagementPartnerable)
    GetTermsAndConditions()([]TermsAndConditionsable)
    GetTroubleshootingEvents()([]DeviceManagementTroubleshootingEventable)
    GetWindowsAutopilotDeviceIdentities()([]WindowsAutopilotDeviceIdentityable)
    GetWindowsInformationProtectionAppLearningSummaries()([]WindowsInformationProtectionAppLearningSummaryable)
    GetWindowsInformationProtectionNetworkLearningSummaries()([]WindowsInformationProtectionNetworkLearningSummaryable)
    SetApplePushNotificationCertificate(value ApplePushNotificationCertificateable)()
    SetComplianceManagementPartners(value []ComplianceManagementPartnerable)()
    SetConditionalAccessSettings(value OnPremisesConditionalAccessSettingsable)()
    SetDetectedApps(value []DetectedAppable)()
    SetDeviceCategories(value []DeviceCategoryable)()
    SetDeviceCompliancePolicies(value []DeviceCompliancePolicyable)()
    SetDeviceCompliancePolicyDeviceStateSummary(value DeviceCompliancePolicyDeviceStateSummaryable)()
    SetDeviceCompliancePolicySettingStateSummaries(value []DeviceCompliancePolicySettingStateSummaryable)()
    SetDeviceConfigurationDeviceStateSummaries(value DeviceConfigurationDeviceStateSummaryable)()
    SetDeviceConfigurations(value []DeviceConfigurationable)()
    SetDeviceEnrollmentConfigurations(value []DeviceEnrollmentConfigurationable)()
    SetDeviceManagementPartners(value []DeviceManagementPartnerable)()
    SetExchangeConnectors(value []DeviceManagementExchangeConnectorable)()
    SetImportedWindowsAutopilotDeviceIdentities(value []ImportedWindowsAutopilotDeviceIdentityable)()
    SetIntuneAccountId(value *string)()
    SetIntuneBrand(value IntuneBrandable)()
    SetIosUpdateStatuses(value []IosUpdateDeviceStatusable)()
    SetManagedDeviceOverview(value ManagedDeviceOverviewable)()
    SetManagedDevices(value []ManagedDeviceable)()
    SetMobileThreatDefenseConnectors(value []MobileThreatDefenseConnectorable)()
    SetNotificationMessageTemplates(value []NotificationMessageTemplateable)()
    SetRemoteAssistancePartners(value []RemoteAssistancePartnerable)()
    SetReports(value DeviceManagementReportsable)()
    SetResourceOperations(value []ResourceOperationable)()
    SetRoleAssignments(value []DeviceAndAppManagementRoleAssignmentable)()
    SetRoleDefinitions(value []RoleDefinitionable)()
    SetSettings(value DeviceManagementSettingsable)()
    SetSoftwareUpdateStatusSummary(value SoftwareUpdateStatusSummaryable)()
    SetSubscriptionState(value *DeviceManagementSubscriptionState)()
    SetTelecomExpenseManagementPartners(value []TelecomExpenseManagementPartnerable)()
    SetTermsAndConditions(value []TermsAndConditionsable)()
    SetTroubleshootingEvents(value []DeviceManagementTroubleshootingEventable)()
    SetWindowsAutopilotDeviceIdentities(value []WindowsAutopilotDeviceIdentityable)()
    SetWindowsInformationProtectionAppLearningSummaries(value []WindowsInformationProtectionAppLearningSummaryable)()
    SetWindowsInformationProtectionNetworkLearningSummaries(value []WindowsInformationProtectionNetworkLearningSummaryable)()
}
