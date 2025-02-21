package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ChoiceColumn provides operations to manage the drive singleton.
type ChoiceColumn struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // If true, allows custom values that aren't in the configured choices.
    allowTextEntry *bool;
    // The list of values available for this column.
    choices []string;
    // How the choices are to be presented in the UX. Must be one of checkBoxes, dropDownMenu, or radioButtons
    displayAs *string;
}
// NewChoiceColumn instantiates a new choiceColumn and sets the default values.
func NewChoiceColumn()(*ChoiceColumn) {
    m := &ChoiceColumn{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateChoiceColumnFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateChoiceColumnFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewChoiceColumn(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ChoiceColumn) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAllowTextEntry gets the allowTextEntry property value. If true, allows custom values that aren't in the configured choices.
func (m *ChoiceColumn) GetAllowTextEntry()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowTextEntry
    }
}
// GetChoices gets the choices property value. The list of values available for this column.
func (m *ChoiceColumn) GetChoices()([]string) {
    if m == nil {
        return nil
    } else {
        return m.choices
    }
}
// GetDisplayAs gets the displayAs property value. How the choices are to be presented in the UX. Must be one of checkBoxes, dropDownMenu, or radioButtons
func (m *ChoiceColumn) GetDisplayAs()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayAs
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ChoiceColumn) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["allowTextEntry"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowTextEntry(val)
        }
        return nil
    }
    res["choices"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetChoices(res)
        }
        return nil
    }
    res["displayAs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayAs(val)
        }
        return nil
    }
    return res
}
func (m *ChoiceColumn) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ChoiceColumn) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowTextEntry", m.GetAllowTextEntry())
        if err != nil {
            return err
        }
    }
    if m.GetChoices() != nil {
        err := writer.WriteCollectionOfStringValues("choices", m.GetChoices())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayAs", m.GetDisplayAs())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ChoiceColumn) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAllowTextEntry sets the allowTextEntry property value. If true, allows custom values that aren't in the configured choices.
func (m *ChoiceColumn) SetAllowTextEntry(value *bool)() {
    if m != nil {
        m.allowTextEntry = value
    }
}
// SetChoices sets the choices property value. The list of values available for this column.
func (m *ChoiceColumn) SetChoices(value []string)() {
    if m != nil {
        m.choices = value
    }
}
// SetDisplayAs sets the displayAs property value. How the choices are to be presented in the UX. Must be one of checkBoxes, dropDownMenu, or radioButtons
func (m *ChoiceColumn) SetDisplayAs(value *string)() {
    if m != nil {
        m.displayAs = value
    }
}
