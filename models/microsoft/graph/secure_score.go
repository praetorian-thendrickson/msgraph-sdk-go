package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SecureScore provides operations to manage the security singleton.
type SecureScore struct {
    Entity
    // Active user count of the given tenant.
    activeUserCount *int32;
    // Average score by different scopes (for example, average by industry, average by seating) and control category (Identity, Data, Device, Apps, Infrastructure) within the scope.
    averageComparativeScores []AverageComparativeScoreable;
    // GUID string for tenant ID.
    azureTenantId *string;
    // Contains tenant scores for a set of controls.
    controlScores []ControlScoreable;
    // The date when the entity is created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Tenant current attained score on specified date.
    currentScore *float64;
    // Microsoft-provided services for the tenant (for example, Exchange online, Skype, Sharepoint).
    enabledServices []string;
    // Licensed user count of the given tenant.
    licensedUserCount *int32;
    // Tenant maximum possible score on specified date.
    maxScore *float64;
    // Complex type containing details about the security product/service vendor, provider, and subprovider (for example, vendor=Microsoft; provider=SecureScore). Required.
    vendorInformation SecurityVendorInformationable;
}
// NewSecureScore instantiates a new secureScore and sets the default values.
func NewSecureScore()(*SecureScore) {
    m := &SecureScore{
        Entity: *NewEntity(),
    }
    return m
}
// CreateSecureScoreFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSecureScoreFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewSecureScore(), nil
}
// GetActiveUserCount gets the activeUserCount property value. Active user count of the given tenant.
func (m *SecureScore) GetActiveUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.activeUserCount
    }
}
// GetAverageComparativeScores gets the averageComparativeScores property value. Average score by different scopes (for example, average by industry, average by seating) and control category (Identity, Data, Device, Apps, Infrastructure) within the scope.
func (m *SecureScore) GetAverageComparativeScores()([]AverageComparativeScoreable) {
    if m == nil {
        return nil
    } else {
        return m.averageComparativeScores
    }
}
// GetAzureTenantId gets the azureTenantId property value. GUID string for tenant ID.
func (m *SecureScore) GetAzureTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.azureTenantId
    }
}
// GetControlScores gets the controlScores property value. Contains tenant scores for a set of controls.
func (m *SecureScore) GetControlScores()([]ControlScoreable) {
    if m == nil {
        return nil
    } else {
        return m.controlScores
    }
}
// GetCreatedDateTime gets the createdDateTime property value. The date when the entity is created.
func (m *SecureScore) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetCurrentScore gets the currentScore property value. Tenant current attained score on specified date.
func (m *SecureScore) GetCurrentScore()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.currentScore
    }
}
// GetEnabledServices gets the enabledServices property value. Microsoft-provided services for the tenant (for example, Exchange online, Skype, Sharepoint).
func (m *SecureScore) GetEnabledServices()([]string) {
    if m == nil {
        return nil
    } else {
        return m.enabledServices
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SecureScore) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activeUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActiveUserCount(val)
        }
        return nil
    }
    res["averageComparativeScores"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAverageComparativeScoreFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AverageComparativeScoreable, len(val))
            for i, v := range val {
                res[i] = v.(AverageComparativeScoreable)
            }
            m.SetAverageComparativeScores(res)
        }
        return nil
    }
    res["azureTenantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureTenantId(val)
        }
        return nil
    }
    res["controlScores"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateControlScoreFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ControlScoreable, len(val))
            for i, v := range val {
                res[i] = v.(ControlScoreable)
            }
            m.SetControlScores(res)
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
    res["currentScore"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentScore(val)
        }
        return nil
    }
    res["enabledServices"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetEnabledServices(res)
        }
        return nil
    }
    res["licensedUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLicensedUserCount(val)
        }
        return nil
    }
    res["maxScore"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxScore(val)
        }
        return nil
    }
    res["vendorInformation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateSecurityVendorInformationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVendorInformation(val.(SecurityVendorInformationable))
        }
        return nil
    }
    return res
}
// GetLicensedUserCount gets the licensedUserCount property value. Licensed user count of the given tenant.
func (m *SecureScore) GetLicensedUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.licensedUserCount
    }
}
// GetMaxScore gets the maxScore property value. Tenant maximum possible score on specified date.
func (m *SecureScore) GetMaxScore()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.maxScore
    }
}
// GetVendorInformation gets the vendorInformation property value. Complex type containing details about the security product/service vendor, provider, and subprovider (for example, vendor=Microsoft; provider=SecureScore). Required.
func (m *SecureScore) GetVendorInformation()(SecurityVendorInformationable) {
    if m == nil {
        return nil
    } else {
        return m.vendorInformation
    }
}
func (m *SecureScore) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *SecureScore) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("activeUserCount", m.GetActiveUserCount())
        if err != nil {
            return err
        }
    }
    if m.GetAverageComparativeScores() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAverageComparativeScores()))
        for i, v := range m.GetAverageComparativeScores() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("averageComparativeScores", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("azureTenantId", m.GetAzureTenantId())
        if err != nil {
            return err
        }
    }
    if m.GetControlScores() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetControlScores()))
        for i, v := range m.GetControlScores() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("controlScores", cast)
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
        err = writer.WriteFloat64Value("currentScore", m.GetCurrentScore())
        if err != nil {
            return err
        }
    }
    if m.GetEnabledServices() != nil {
        err = writer.WriteCollectionOfStringValues("enabledServices", m.GetEnabledServices())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("licensedUserCount", m.GetLicensedUserCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("maxScore", m.GetMaxScore())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("vendorInformation", m.GetVendorInformation())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActiveUserCount sets the activeUserCount property value. Active user count of the given tenant.
func (m *SecureScore) SetActiveUserCount(value *int32)() {
    if m != nil {
        m.activeUserCount = value
    }
}
// SetAverageComparativeScores sets the averageComparativeScores property value. Average score by different scopes (for example, average by industry, average by seating) and control category (Identity, Data, Device, Apps, Infrastructure) within the scope.
func (m *SecureScore) SetAverageComparativeScores(value []AverageComparativeScoreable)() {
    if m != nil {
        m.averageComparativeScores = value
    }
}
// SetAzureTenantId sets the azureTenantId property value. GUID string for tenant ID.
func (m *SecureScore) SetAzureTenantId(value *string)() {
    if m != nil {
        m.azureTenantId = value
    }
}
// SetControlScores sets the controlScores property value. Contains tenant scores for a set of controls.
func (m *SecureScore) SetControlScores(value []ControlScoreable)() {
    if m != nil {
        m.controlScores = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The date when the entity is created.
func (m *SecureScore) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetCurrentScore sets the currentScore property value. Tenant current attained score on specified date.
func (m *SecureScore) SetCurrentScore(value *float64)() {
    if m != nil {
        m.currentScore = value
    }
}
// SetEnabledServices sets the enabledServices property value. Microsoft-provided services for the tenant (for example, Exchange online, Skype, Sharepoint).
func (m *SecureScore) SetEnabledServices(value []string)() {
    if m != nil {
        m.enabledServices = value
    }
}
// SetLicensedUserCount sets the licensedUserCount property value. Licensed user count of the given tenant.
func (m *SecureScore) SetLicensedUserCount(value *int32)() {
    if m != nil {
        m.licensedUserCount = value
    }
}
// SetMaxScore sets the maxScore property value. Tenant maximum possible score on specified date.
func (m *SecureScore) SetMaxScore(value *float64)() {
    if m != nil {
        m.maxScore = value
    }
}
// SetVendorInformation sets the vendorInformation property value. Complex type containing details about the security product/service vendor, provider, and subprovider (for example, vendor=Microsoft; provider=SecureScore). Required.
func (m *SecureScore) SetVendorInformation(value SecurityVendorInformationable)() {
    if m != nil {
        m.vendorInformation = value
    }
}
