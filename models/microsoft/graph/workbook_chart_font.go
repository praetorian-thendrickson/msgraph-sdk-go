package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WorkbookChartFont provides operations to manage the drive singleton.
type WorkbookChartFont struct {
    Entity
    // Represents the bold status of font.
    bold *bool;
    // HTML color code representation of the text color. E.g. #FF0000 represents Red.
    color *string;
    // Represents the italic status of the font.
    italic *bool;
    // Font name (e.g. 'Calibri')
    name *string;
    // Size of the font (e.g. 11)
    size *float64;
    // Type of underline applied to the font. The possible values are: None, Single.
    underline *string;
}
// NewWorkbookChartFont instantiates a new workbookChartFont and sets the default values.
func NewWorkbookChartFont()(*WorkbookChartFont) {
    m := &WorkbookChartFont{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookChartFontFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookChartFontFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewWorkbookChartFont(), nil
}
// GetBold gets the bold property value. Represents the bold status of font.
func (m *WorkbookChartFont) GetBold()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.bold
    }
}
// GetColor gets the color property value. HTML color code representation of the text color. E.g. #FF0000 represents Red.
func (m *WorkbookChartFont) GetColor()(*string) {
    if m == nil {
        return nil
    } else {
        return m.color
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookChartFont) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["bold"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBold(val)
        }
        return nil
    }
    res["color"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetColor(val)
        }
        return nil
    }
    res["italic"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetItalic(val)
        }
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["size"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSize(val)
        }
        return nil
    }
    res["underline"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnderline(val)
        }
        return nil
    }
    return res
}
// GetItalic gets the italic property value. Represents the italic status of the font.
func (m *WorkbookChartFont) GetItalic()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.italic
    }
}
// GetName gets the name property value. Font name (e.g. 'Calibri')
func (m *WorkbookChartFont) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetSize gets the size property value. Size of the font (e.g. 11)
func (m *WorkbookChartFont) GetSize()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.size
    }
}
// GetUnderline gets the underline property value. Type of underline applied to the font. The possible values are: None, Single.
func (m *WorkbookChartFont) GetUnderline()(*string) {
    if m == nil {
        return nil
    } else {
        return m.underline
    }
}
func (m *WorkbookChartFont) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *WorkbookChartFont) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("bold", m.GetBold())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("color", m.GetColor())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("italic", m.GetItalic())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("size", m.GetSize())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("underline", m.GetUnderline())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBold sets the bold property value. Represents the bold status of font.
func (m *WorkbookChartFont) SetBold(value *bool)() {
    if m != nil {
        m.bold = value
    }
}
// SetColor sets the color property value. HTML color code representation of the text color. E.g. #FF0000 represents Red.
func (m *WorkbookChartFont) SetColor(value *string)() {
    if m != nil {
        m.color = value
    }
}
// SetItalic sets the italic property value. Represents the italic status of the font.
func (m *WorkbookChartFont) SetItalic(value *bool)() {
    if m != nil {
        m.italic = value
    }
}
// SetName sets the name property value. Font name (e.g. 'Calibri')
func (m *WorkbookChartFont) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetSize sets the size property value. Size of the font (e.g. 11)
func (m *WorkbookChartFont) SetSize(value *float64)() {
    if m != nil {
        m.size = value
    }
}
// SetUnderline sets the underline property value. Type of underline applied to the font. The possible values are: None, Single.
func (m *WorkbookChartFont) SetUnderline(value *string)() {
    if m != nil {
        m.underline = value
    }
}
