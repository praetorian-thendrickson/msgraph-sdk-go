package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AndroidManagedAppProtection provides operations to manage the deviceAppManagement singleton.
type AndroidManagedAppProtection struct {
    TargetedManagedAppProtection
    // List of apps to which the policy is deployed.
    apps []ManagedMobileAppable;
    // Friendly name of the preferred custom browser to open weblink on Android. When this property is configured, ManagedBrowserToOpenLinksRequired should be true.
    customBrowserDisplayName *string;
    // Unique identifier of the preferred custom browser to open weblink on Android. When this property is configured, ManagedBrowserToOpenLinksRequired should be true.
    customBrowserPackageId *string;
    // Count of apps to which the current policy is deployed.
    deployedAppCount *int32;
    // Navigation property to deployment summary of the configuration.
    deploymentSummary ManagedAppPolicyDeploymentSummaryable;
    // When this setting is enabled, app level encryption is disabled if device level encryption is enabled
    disableAppEncryptionIfDeviceEncryptionIsEnabled *bool;
    // Indicates whether application data for managed apps should be encrypted
    encryptAppData *bool;
    // Define the oldest required Android security patch level a user can have to gain secure access to the app.
    minimumRequiredPatchVersion *string;
    // Define the oldest recommended Android security patch level a user can have for secure access to the app.
    minimumWarningPatchVersion *string;
    // Indicates whether a managed user can take screen captures of managed apps
    screenCaptureBlocked *bool;
}
// NewAndroidManagedAppProtection instantiates a new androidManagedAppProtection and sets the default values.
func NewAndroidManagedAppProtection()(*AndroidManagedAppProtection) {
    m := &AndroidManagedAppProtection{
        TargetedManagedAppProtection: *NewTargetedManagedAppProtection(),
    }
    return m
}
// CreateAndroidManagedAppProtectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidManagedAppProtectionFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewAndroidManagedAppProtection(), nil
}
// GetApps gets the apps property value. List of apps to which the policy is deployed.
func (m *AndroidManagedAppProtection) GetApps()([]ManagedMobileAppable) {
    if m == nil {
        return nil
    } else {
        return m.apps
    }
}
// GetCustomBrowserDisplayName gets the customBrowserDisplayName property value. Friendly name of the preferred custom browser to open weblink on Android. When this property is configured, ManagedBrowserToOpenLinksRequired should be true.
func (m *AndroidManagedAppProtection) GetCustomBrowserDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customBrowserDisplayName
    }
}
// GetCustomBrowserPackageId gets the customBrowserPackageId property value. Unique identifier of the preferred custom browser to open weblink on Android. When this property is configured, ManagedBrowserToOpenLinksRequired should be true.
func (m *AndroidManagedAppProtection) GetCustomBrowserPackageId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customBrowserPackageId
    }
}
// GetDeployedAppCount gets the deployedAppCount property value. Count of apps to which the current policy is deployed.
func (m *AndroidManagedAppProtection) GetDeployedAppCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.deployedAppCount
    }
}
// GetDeploymentSummary gets the deploymentSummary property value. Navigation property to deployment summary of the configuration.
func (m *AndroidManagedAppProtection) GetDeploymentSummary()(ManagedAppPolicyDeploymentSummaryable) {
    if m == nil {
        return nil
    } else {
        return m.deploymentSummary
    }
}
// GetDisableAppEncryptionIfDeviceEncryptionIsEnabled gets the disableAppEncryptionIfDeviceEncryptionIsEnabled property value. When this setting is enabled, app level encryption is disabled if device level encryption is enabled
func (m *AndroidManagedAppProtection) GetDisableAppEncryptionIfDeviceEncryptionIsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.disableAppEncryptionIfDeviceEncryptionIsEnabled
    }
}
// GetEncryptAppData gets the encryptAppData property value. Indicates whether application data for managed apps should be encrypted
func (m *AndroidManagedAppProtection) GetEncryptAppData()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.encryptAppData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidManagedAppProtection) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.TargetedManagedAppProtection.GetFieldDeserializers()
    res["apps"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedMobileAppFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedMobileAppable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedMobileAppable)
            }
            m.SetApps(res)
        }
        return nil
    }
    res["customBrowserDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomBrowserDisplayName(val)
        }
        return nil
    }
    res["customBrowserPackageId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomBrowserPackageId(val)
        }
        return nil
    }
    res["deployedAppCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeployedAppCount(val)
        }
        return nil
    }
    res["deploymentSummary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateManagedAppPolicyDeploymentSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeploymentSummary(val.(ManagedAppPolicyDeploymentSummaryable))
        }
        return nil
    }
    res["disableAppEncryptionIfDeviceEncryptionIsEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisableAppEncryptionIfDeviceEncryptionIsEnabled(val)
        }
        return nil
    }
    res["encryptAppData"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEncryptAppData(val)
        }
        return nil
    }
    res["minimumRequiredPatchVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumRequiredPatchVersion(val)
        }
        return nil
    }
    res["minimumWarningPatchVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumWarningPatchVersion(val)
        }
        return nil
    }
    res["screenCaptureBlocked"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScreenCaptureBlocked(val)
        }
        return nil
    }
    return res
}
// GetMinimumRequiredPatchVersion gets the minimumRequiredPatchVersion property value. Define the oldest required Android security patch level a user can have to gain secure access to the app.
func (m *AndroidManagedAppProtection) GetMinimumRequiredPatchVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.minimumRequiredPatchVersion
    }
}
// GetMinimumWarningPatchVersion gets the minimumWarningPatchVersion property value. Define the oldest recommended Android security patch level a user can have for secure access to the app.
func (m *AndroidManagedAppProtection) GetMinimumWarningPatchVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.minimumWarningPatchVersion
    }
}
// GetScreenCaptureBlocked gets the screenCaptureBlocked property value. Indicates whether a managed user can take screen captures of managed apps
func (m *AndroidManagedAppProtection) GetScreenCaptureBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.screenCaptureBlocked
    }
}
func (m *AndroidManagedAppProtection) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AndroidManagedAppProtection) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.TargetedManagedAppProtection.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetApps() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetApps()))
        for i, v := range m.GetApps() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("apps", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customBrowserDisplayName", m.GetCustomBrowserDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customBrowserPackageId", m.GetCustomBrowserPackageId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("deployedAppCount", m.GetDeployedAppCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deploymentSummary", m.GetDeploymentSummary())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("disableAppEncryptionIfDeviceEncryptionIsEnabled", m.GetDisableAppEncryptionIfDeviceEncryptionIsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("encryptAppData", m.GetEncryptAppData())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("minimumRequiredPatchVersion", m.GetMinimumRequiredPatchVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("minimumWarningPatchVersion", m.GetMinimumWarningPatchVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("screenCaptureBlocked", m.GetScreenCaptureBlocked())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApps sets the apps property value. List of apps to which the policy is deployed.
func (m *AndroidManagedAppProtection) SetApps(value []ManagedMobileAppable)() {
    if m != nil {
        m.apps = value
    }
}
// SetCustomBrowserDisplayName sets the customBrowserDisplayName property value. Friendly name of the preferred custom browser to open weblink on Android. When this property is configured, ManagedBrowserToOpenLinksRequired should be true.
func (m *AndroidManagedAppProtection) SetCustomBrowserDisplayName(value *string)() {
    if m != nil {
        m.customBrowserDisplayName = value
    }
}
// SetCustomBrowserPackageId sets the customBrowserPackageId property value. Unique identifier of the preferred custom browser to open weblink on Android. When this property is configured, ManagedBrowserToOpenLinksRequired should be true.
func (m *AndroidManagedAppProtection) SetCustomBrowserPackageId(value *string)() {
    if m != nil {
        m.customBrowserPackageId = value
    }
}
// SetDeployedAppCount sets the deployedAppCount property value. Count of apps to which the current policy is deployed.
func (m *AndroidManagedAppProtection) SetDeployedAppCount(value *int32)() {
    if m != nil {
        m.deployedAppCount = value
    }
}
// SetDeploymentSummary sets the deploymentSummary property value. Navigation property to deployment summary of the configuration.
func (m *AndroidManagedAppProtection) SetDeploymentSummary(value ManagedAppPolicyDeploymentSummaryable)() {
    if m != nil {
        m.deploymentSummary = value
    }
}
// SetDisableAppEncryptionIfDeviceEncryptionIsEnabled sets the disableAppEncryptionIfDeviceEncryptionIsEnabled property value. When this setting is enabled, app level encryption is disabled if device level encryption is enabled
func (m *AndroidManagedAppProtection) SetDisableAppEncryptionIfDeviceEncryptionIsEnabled(value *bool)() {
    if m != nil {
        m.disableAppEncryptionIfDeviceEncryptionIsEnabled = value
    }
}
// SetEncryptAppData sets the encryptAppData property value. Indicates whether application data for managed apps should be encrypted
func (m *AndroidManagedAppProtection) SetEncryptAppData(value *bool)() {
    if m != nil {
        m.encryptAppData = value
    }
}
// SetMinimumRequiredPatchVersion sets the minimumRequiredPatchVersion property value. Define the oldest required Android security patch level a user can have to gain secure access to the app.
func (m *AndroidManagedAppProtection) SetMinimumRequiredPatchVersion(value *string)() {
    if m != nil {
        m.minimumRequiredPatchVersion = value
    }
}
// SetMinimumWarningPatchVersion sets the minimumWarningPatchVersion property value. Define the oldest recommended Android security patch level a user can have for secure access to the app.
func (m *AndroidManagedAppProtection) SetMinimumWarningPatchVersion(value *string)() {
    if m != nil {
        m.minimumWarningPatchVersion = value
    }
}
// SetScreenCaptureBlocked sets the screenCaptureBlocked property value. Indicates whether a managed user can take screen captures of managed apps
func (m *AndroidManagedAppProtection) SetScreenCaptureBlocked(value *bool)() {
    if m != nil {
        m.screenCaptureBlocked = value
    }
}
