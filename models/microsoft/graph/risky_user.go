package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RiskyUser provides operations to manage the identityProtectionRoot singleton.
type RiskyUser struct {
    Entity
    // The activity related to user risk level change
    history []RiskyUserHistoryItemable;
    // Indicates whether the user is deleted. Possible values are: true, false.
    isDeleted *bool;
    // Indicates whether a user's risky state is being processed by the backend.
    isProcessing *bool;
    // Details of the detected risk. Possible values are: none, adminGeneratedTemporaryPassword, userPerformedSecuredPasswordChange, userPerformedSecuredPasswordReset, adminConfirmedSigninSafe, aiConfirmedSigninSafe, userPassedMFADrivenByRiskBasedPolicy, adminDismissedAllRiskForUser, adminConfirmedSigninCompromised, hidden, adminConfirmedUserCompromised, unknownFutureValue.
    riskDetail *RiskDetail;
    // The date and time that the risky user was last updated.  The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    riskLastUpdatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Level of the detected risky user. Possible values are: low, medium, high, hidden, none, unknownFutureValue.
    riskLevel *RiskLevel;
    // State of the user's risk. Possible values are: none, confirmedSafe, remediated, dismissed, atRisk, confirmedCompromised, unknownFutureValue.
    riskState *RiskState;
    // Risky user display name.
    userDisplayName *string;
    // Risky user principal name.
    userPrincipalName *string;
}
// NewRiskyUser instantiates a new riskyUser and sets the default values.
func NewRiskyUser()(*RiskyUser) {
    m := &RiskyUser{
        Entity: *NewEntity(),
    }
    return m
}
// CreateRiskyUserFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRiskyUserFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewRiskyUser(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RiskyUser) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["history"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRiskyUserHistoryItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RiskyUserHistoryItemable, len(val))
            for i, v := range val {
                res[i] = v.(RiskyUserHistoryItemable)
            }
            m.SetHistory(res)
        }
        return nil
    }
    res["isDeleted"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDeleted(val)
        }
        return nil
    }
    res["isProcessing"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsProcessing(val)
        }
        return nil
    }
    res["riskDetail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRiskDetail)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRiskDetail(val.(*RiskDetail))
        }
        return nil
    }
    res["riskLastUpdatedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRiskLastUpdatedDateTime(val)
        }
        return nil
    }
    res["riskLevel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRiskLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRiskLevel(val.(*RiskLevel))
        }
        return nil
    }
    res["riskState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRiskState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRiskState(val.(*RiskState))
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
// GetHistory gets the history property value. The activity related to user risk level change
func (m *RiskyUser) GetHistory()([]RiskyUserHistoryItemable) {
    if m == nil {
        return nil
    } else {
        return m.history
    }
}
// GetIsDeleted gets the isDeleted property value. Indicates whether the user is deleted. Possible values are: true, false.
func (m *RiskyUser) GetIsDeleted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDeleted
    }
}
// GetIsProcessing gets the isProcessing property value. Indicates whether a user's risky state is being processed by the backend.
func (m *RiskyUser) GetIsProcessing()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isProcessing
    }
}
// GetRiskDetail gets the riskDetail property value. Details of the detected risk. Possible values are: none, adminGeneratedTemporaryPassword, userPerformedSecuredPasswordChange, userPerformedSecuredPasswordReset, adminConfirmedSigninSafe, aiConfirmedSigninSafe, userPassedMFADrivenByRiskBasedPolicy, adminDismissedAllRiskForUser, adminConfirmedSigninCompromised, hidden, adminConfirmedUserCompromised, unknownFutureValue.
func (m *RiskyUser) GetRiskDetail()(*RiskDetail) {
    if m == nil {
        return nil
    } else {
        return m.riskDetail
    }
}
// GetRiskLastUpdatedDateTime gets the riskLastUpdatedDateTime property value. The date and time that the risky user was last updated.  The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *RiskyUser) GetRiskLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.riskLastUpdatedDateTime
    }
}
// GetRiskLevel gets the riskLevel property value. Level of the detected risky user. Possible values are: low, medium, high, hidden, none, unknownFutureValue.
func (m *RiskyUser) GetRiskLevel()(*RiskLevel) {
    if m == nil {
        return nil
    } else {
        return m.riskLevel
    }
}
// GetRiskState gets the riskState property value. State of the user's risk. Possible values are: none, confirmedSafe, remediated, dismissed, atRisk, confirmedCompromised, unknownFutureValue.
func (m *RiskyUser) GetRiskState()(*RiskState) {
    if m == nil {
        return nil
    } else {
        return m.riskState
    }
}
// GetUserDisplayName gets the userDisplayName property value. Risky user display name.
func (m *RiskyUser) GetUserDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userDisplayName
    }
}
// GetUserPrincipalName gets the userPrincipalName property value. Risky user principal name.
func (m *RiskyUser) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
func (m *RiskyUser) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *RiskyUser) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetHistory() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetHistory()))
        for i, v := range m.GetHistory() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("history", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDeleted", m.GetIsDeleted())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isProcessing", m.GetIsProcessing())
        if err != nil {
            return err
        }
    }
    if m.GetRiskDetail() != nil {
        cast := (*m.GetRiskDetail()).String()
        err = writer.WriteStringValue("riskDetail", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("riskLastUpdatedDateTime", m.GetRiskLastUpdatedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetRiskLevel() != nil {
        cast := (*m.GetRiskLevel()).String()
        err = writer.WriteStringValue("riskLevel", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetRiskState() != nil {
        cast := (*m.GetRiskState()).String()
        err = writer.WriteStringValue("riskState", &cast)
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
// SetHistory sets the history property value. The activity related to user risk level change
func (m *RiskyUser) SetHistory(value []RiskyUserHistoryItemable)() {
    if m != nil {
        m.history = value
    }
}
// SetIsDeleted sets the isDeleted property value. Indicates whether the user is deleted. Possible values are: true, false.
func (m *RiskyUser) SetIsDeleted(value *bool)() {
    if m != nil {
        m.isDeleted = value
    }
}
// SetIsProcessing sets the isProcessing property value. Indicates whether a user's risky state is being processed by the backend.
func (m *RiskyUser) SetIsProcessing(value *bool)() {
    if m != nil {
        m.isProcessing = value
    }
}
// SetRiskDetail sets the riskDetail property value. Details of the detected risk. Possible values are: none, adminGeneratedTemporaryPassword, userPerformedSecuredPasswordChange, userPerformedSecuredPasswordReset, adminConfirmedSigninSafe, aiConfirmedSigninSafe, userPassedMFADrivenByRiskBasedPolicy, adminDismissedAllRiskForUser, adminConfirmedSigninCompromised, hidden, adminConfirmedUserCompromised, unknownFutureValue.
func (m *RiskyUser) SetRiskDetail(value *RiskDetail)() {
    if m != nil {
        m.riskDetail = value
    }
}
// SetRiskLastUpdatedDateTime sets the riskLastUpdatedDateTime property value. The date and time that the risky user was last updated.  The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *RiskyUser) SetRiskLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.riskLastUpdatedDateTime = value
    }
}
// SetRiskLevel sets the riskLevel property value. Level of the detected risky user. Possible values are: low, medium, high, hidden, none, unknownFutureValue.
func (m *RiskyUser) SetRiskLevel(value *RiskLevel)() {
    if m != nil {
        m.riskLevel = value
    }
}
// SetRiskState sets the riskState property value. State of the user's risk. Possible values are: none, confirmedSafe, remediated, dismissed, atRisk, confirmedCompromised, unknownFutureValue.
func (m *RiskyUser) SetRiskState(value *RiskState)() {
    if m != nil {
        m.riskState = value
    }
}
// SetUserDisplayName sets the userDisplayName property value. Risky user display name.
func (m *RiskyUser) SetUserDisplayName(value *string)() {
    if m != nil {
        m.userDisplayName = value
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. Risky user principal name.
func (m *RiskyUser) SetUserPrincipalName(value *string)() {
    if m != nil {
        m.userPrincipalName = value
    }
}
