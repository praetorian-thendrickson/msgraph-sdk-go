package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ManagedDeviceMobileAppConfiguration provides operations to manage the deviceAppManagement singleton.
type ManagedDeviceMobileAppConfiguration struct {
    Entity
    // The list of group assignemenets for app configration.
    assignments []ManagedDeviceMobileAppConfigurationAssignmentable;
    // DateTime the object was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Admin provided description of the Device Configuration.
    description *string;
    // List of ManagedDeviceMobileAppConfigurationDeviceStatus.
    deviceStatuses []ManagedDeviceMobileAppConfigurationDeviceStatusable;
    // App configuration device status summary.
    deviceStatusSummary ManagedDeviceMobileAppConfigurationDeviceSummaryable;
    // Admin provided name of the device configuration.
    displayName *string;
    // DateTime the object was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // the associated app.
    targetedMobileApps []string;
    // List of ManagedDeviceMobileAppConfigurationUserStatus.
    userStatuses []ManagedDeviceMobileAppConfigurationUserStatusable;
    // App configuration user status summary.
    userStatusSummary ManagedDeviceMobileAppConfigurationUserSummaryable;
    // Version of the device configuration.
    version *int32;
}
// NewManagedDeviceMobileAppConfiguration instantiates a new managedDeviceMobileAppConfiguration and sets the default values.
func NewManagedDeviceMobileAppConfiguration()(*ManagedDeviceMobileAppConfiguration) {
    m := &ManagedDeviceMobileAppConfiguration{
        Entity: *NewEntity(),
    }
    return m
}
// CreateManagedDeviceMobileAppConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedDeviceMobileAppConfigurationFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewManagedDeviceMobileAppConfiguration(), nil
}
// GetAssignments gets the assignments property value. The list of group assignemenets for app configration.
func (m *ManagedDeviceMobileAppConfiguration) GetAssignments()([]ManagedDeviceMobileAppConfigurationAssignmentable) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
// GetCreatedDateTime gets the createdDateTime property value. DateTime the object was created.
func (m *ManagedDeviceMobileAppConfiguration) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDescription gets the description property value. Admin provided description of the Device Configuration.
func (m *ManagedDeviceMobileAppConfiguration) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDeviceStatuses gets the deviceStatuses property value. List of ManagedDeviceMobileAppConfigurationDeviceStatus.
func (m *ManagedDeviceMobileAppConfiguration) GetDeviceStatuses()([]ManagedDeviceMobileAppConfigurationDeviceStatusable) {
    if m == nil {
        return nil
    } else {
        return m.deviceStatuses
    }
}
// GetDeviceStatusSummary gets the deviceStatusSummary property value. App configuration device status summary.
func (m *ManagedDeviceMobileAppConfiguration) GetDeviceStatusSummary()(ManagedDeviceMobileAppConfigurationDeviceSummaryable) {
    if m == nil {
        return nil
    } else {
        return m.deviceStatusSummary
    }
}
// GetDisplayName gets the displayName property value. Admin provided name of the device configuration.
func (m *ManagedDeviceMobileAppConfiguration) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedDeviceMobileAppConfiguration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedDeviceMobileAppConfigurationAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedDeviceMobileAppConfigurationAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedDeviceMobileAppConfigurationAssignmentable)
            }
            m.SetAssignments(res)
        }
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["deviceStatuses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedDeviceMobileAppConfigurationDeviceStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedDeviceMobileAppConfigurationDeviceStatusable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedDeviceMobileAppConfigurationDeviceStatusable)
            }
            m.SetDeviceStatuses(res)
        }
        return nil
    }
    res["deviceStatusSummary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateManagedDeviceMobileAppConfigurationDeviceSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceStatusSummary(val.(ManagedDeviceMobileAppConfigurationDeviceSummaryable))
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["targetedMobileApps"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetTargetedMobileApps(res)
        }
        return nil
    }
    res["userStatuses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedDeviceMobileAppConfigurationUserStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedDeviceMobileAppConfigurationUserStatusable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedDeviceMobileAppConfigurationUserStatusable)
            }
            m.SetUserStatuses(res)
        }
        return nil
    }
    res["userStatusSummary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateManagedDeviceMobileAppConfigurationUserSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserStatusSummary(val.(ManagedDeviceMobileAppConfigurationUserSummaryable))
        }
        return nil
    }
    res["version"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. DateTime the object was last modified.
func (m *ManagedDeviceMobileAppConfiguration) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetTargetedMobileApps gets the targetedMobileApps property value. the associated app.
func (m *ManagedDeviceMobileAppConfiguration) GetTargetedMobileApps()([]string) {
    if m == nil {
        return nil
    } else {
        return m.targetedMobileApps
    }
}
// GetUserStatuses gets the userStatuses property value. List of ManagedDeviceMobileAppConfigurationUserStatus.
func (m *ManagedDeviceMobileAppConfiguration) GetUserStatuses()([]ManagedDeviceMobileAppConfigurationUserStatusable) {
    if m == nil {
        return nil
    } else {
        return m.userStatuses
    }
}
// GetUserStatusSummary gets the userStatusSummary property value. App configuration user status summary.
func (m *ManagedDeviceMobileAppConfiguration) GetUserStatusSummary()(ManagedDeviceMobileAppConfigurationUserSummaryable) {
    if m == nil {
        return nil
    } else {
        return m.userStatusSummary
    }
}
// GetVersion gets the version property value. Version of the device configuration.
func (m *ManagedDeviceMobileAppConfiguration) GetVersion()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
func (m *ManagedDeviceMobileAppConfiguration) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ManagedDeviceMobileAppConfiguration) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignments() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("assignments", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceStatuses() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeviceStatuses()))
        for i, v := range m.GetDeviceStatuses() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("deviceStatuses", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deviceStatusSummary", m.GetDeviceStatusSummary())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetTargetedMobileApps() != nil {
        err = writer.WriteCollectionOfStringValues("targetedMobileApps", m.GetTargetedMobileApps())
        if err != nil {
            return err
        }
    }
    if m.GetUserStatuses() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserStatuses()))
        for i, v := range m.GetUserStatuses() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userStatuses", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userStatusSummary", m.GetUserStatusSummary())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignments sets the assignments property value. The list of group assignemenets for app configration.
func (m *ManagedDeviceMobileAppConfiguration) SetAssignments(value []ManagedDeviceMobileAppConfigurationAssignmentable)() {
    if m != nil {
        m.assignments = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. DateTime the object was created.
func (m *ManagedDeviceMobileAppConfiguration) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDescription sets the description property value. Admin provided description of the Device Configuration.
func (m *ManagedDeviceMobileAppConfiguration) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDeviceStatuses sets the deviceStatuses property value. List of ManagedDeviceMobileAppConfigurationDeviceStatus.
func (m *ManagedDeviceMobileAppConfiguration) SetDeviceStatuses(value []ManagedDeviceMobileAppConfigurationDeviceStatusable)() {
    if m != nil {
        m.deviceStatuses = value
    }
}
// SetDeviceStatusSummary sets the deviceStatusSummary property value. App configuration device status summary.
func (m *ManagedDeviceMobileAppConfiguration) SetDeviceStatusSummary(value ManagedDeviceMobileAppConfigurationDeviceSummaryable)() {
    if m != nil {
        m.deviceStatusSummary = value
    }
}
// SetDisplayName sets the displayName property value. Admin provided name of the device configuration.
func (m *ManagedDeviceMobileAppConfiguration) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. DateTime the object was last modified.
func (m *ManagedDeviceMobileAppConfiguration) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetTargetedMobileApps sets the targetedMobileApps property value. the associated app.
func (m *ManagedDeviceMobileAppConfiguration) SetTargetedMobileApps(value []string)() {
    if m != nil {
        m.targetedMobileApps = value
    }
}
// SetUserStatuses sets the userStatuses property value. List of ManagedDeviceMobileAppConfigurationUserStatus.
func (m *ManagedDeviceMobileAppConfiguration) SetUserStatuses(value []ManagedDeviceMobileAppConfigurationUserStatusable)() {
    if m != nil {
        m.userStatuses = value
    }
}
// SetUserStatusSummary sets the userStatusSummary property value. App configuration user status summary.
func (m *ManagedDeviceMobileAppConfiguration) SetUserStatusSummary(value ManagedDeviceMobileAppConfigurationUserSummaryable)() {
    if m != nil {
        m.userStatusSummary = value
    }
}
// SetVersion sets the version property value. Version of the device configuration.
func (m *ManagedDeviceMobileAppConfiguration) SetVersion(value *int32)() {
    if m != nil {
        m.version = value
    }
}
