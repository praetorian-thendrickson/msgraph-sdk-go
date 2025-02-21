package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TermsAndConditionsAcceptanceStatus provides operations to manage the deviceManagement singleton.
type TermsAndConditionsAcceptanceStatus struct {
    Entity
    // DateTime when the terms were last accepted by the user.
    acceptedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Most recent version number of the T&C accepted by the user.
    acceptedVersion *int32;
    // Navigation link to the terms and conditions that are assigned.
    termsAndConditions TermsAndConditionsable;
    // Display name of the user whose acceptance the entity represents.
    userDisplayName *string;
    // The userPrincipalName of the User that accepted the term.
    userPrincipalName *string;
}
// NewTermsAndConditionsAcceptanceStatus instantiates a new termsAndConditionsAcceptanceStatus and sets the default values.
func NewTermsAndConditionsAcceptanceStatus()(*TermsAndConditionsAcceptanceStatus) {
    m := &TermsAndConditionsAcceptanceStatus{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTermsAndConditionsAcceptanceStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTermsAndConditionsAcceptanceStatusFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewTermsAndConditionsAcceptanceStatus(), nil
}
// GetAcceptedDateTime gets the acceptedDateTime property value. DateTime when the terms were last accepted by the user.
func (m *TermsAndConditionsAcceptanceStatus) GetAcceptedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.acceptedDateTime
    }
}
// GetAcceptedVersion gets the acceptedVersion property value. Most recent version number of the T&C accepted by the user.
func (m *TermsAndConditionsAcceptanceStatus) GetAcceptedVersion()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.acceptedVersion
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TermsAndConditionsAcceptanceStatus) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["acceptedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAcceptedDateTime(val)
        }
        return nil
    }
    res["acceptedVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAcceptedVersion(val)
        }
        return nil
    }
    res["termsAndConditions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateTermsAndConditionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTermsAndConditions(val.(TermsAndConditionsable))
        }
        return nil
    }
    res["userDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserDisplayName(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    return res
}
// GetTermsAndConditions gets the termsAndConditions property value. Navigation link to the terms and conditions that are assigned.
func (m *TermsAndConditionsAcceptanceStatus) GetTermsAndConditions()(TermsAndConditionsable) {
    if m == nil {
        return nil
    } else {
        return m.termsAndConditions
    }
}
// GetUserDisplayName gets the userDisplayName property value. Display name of the user whose acceptance the entity represents.
func (m *TermsAndConditionsAcceptanceStatus) GetUserDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userDisplayName
    }
}
// GetUserPrincipalName gets the userPrincipalName property value. The userPrincipalName of the User that accepted the term.
func (m *TermsAndConditionsAcceptanceStatus) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
func (m *TermsAndConditionsAcceptanceStatus) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TermsAndConditionsAcceptanceStatus) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("acceptedDateTime", m.GetAcceptedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("acceptedVersion", m.GetAcceptedVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("termsAndConditions", m.GetTermsAndConditions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userDisplayName", m.GetUserDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAcceptedDateTime sets the acceptedDateTime property value. DateTime when the terms were last accepted by the user.
func (m *TermsAndConditionsAcceptanceStatus) SetAcceptedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.acceptedDateTime = value
    }
}
// SetAcceptedVersion sets the acceptedVersion property value. Most recent version number of the T&C accepted by the user.
func (m *TermsAndConditionsAcceptanceStatus) SetAcceptedVersion(value *int32)() {
    if m != nil {
        m.acceptedVersion = value
    }
}
// SetTermsAndConditions sets the termsAndConditions property value. Navigation link to the terms and conditions that are assigned.
func (m *TermsAndConditionsAcceptanceStatus) SetTermsAndConditions(value TermsAndConditionsable)() {
    if m != nil {
        m.termsAndConditions = value
    }
}
// SetUserDisplayName sets the userDisplayName property value. Display name of the user whose acceptance the entity represents.
func (m *TermsAndConditionsAcceptanceStatus) SetUserDisplayName(value *string)() {
    if m != nil {
        m.userDisplayName = value
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. The userPrincipalName of the User that accepted the term.
func (m *TermsAndConditionsAcceptanceStatus) SetUserPrincipalName(value *string)() {
    if m != nil {
        m.userPrincipalName = value
    }
}
