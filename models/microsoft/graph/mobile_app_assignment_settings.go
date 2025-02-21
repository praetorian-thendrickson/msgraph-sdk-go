package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MobileAppAssignmentSettings provides operations to manage the deviceAppManagement singleton.
type MobileAppAssignmentSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
}
// NewMobileAppAssignmentSettings instantiates a new mobileAppAssignmentSettings and sets the default values.
func NewMobileAppAssignmentSettings()(*MobileAppAssignmentSettings) {
    m := &MobileAppAssignmentSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMobileAppAssignmentSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMobileAppAssignmentSettingsFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewMobileAppAssignmentSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MobileAppAssignmentSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MobileAppAssignmentSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    return res
}
func (m *MobileAppAssignmentSettings) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *MobileAppAssignmentSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MobileAppAssignmentSettings) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
