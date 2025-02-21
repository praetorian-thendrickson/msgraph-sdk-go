package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SingleValueLegacyExtendedProperty provides operations to manage the drive singleton.
type SingleValueLegacyExtendedProperty struct {
    Entity
    // A property value.
    value *string;
}
// NewSingleValueLegacyExtendedProperty instantiates a new singleValueLegacyExtendedProperty and sets the default values.
func NewSingleValueLegacyExtendedProperty()(*SingleValueLegacyExtendedProperty) {
    m := &SingleValueLegacyExtendedProperty{
        Entity: *NewEntity(),
    }
    return m
}
// CreateSingleValueLegacyExtendedPropertyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSingleValueLegacyExtendedPropertyFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewSingleValueLegacyExtendedProperty(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SingleValueLegacyExtendedProperty) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["value"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. A property value.
func (m *SingleValueLegacyExtendedProperty) GetValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
func (m *SingleValueLegacyExtendedProperty) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *SingleValueLegacyExtendedProperty) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. A property value.
func (m *SingleValueLegacyExtendedProperty) SetValue(value *string)() {
    if m != nil {
        m.value = value
    }
}
