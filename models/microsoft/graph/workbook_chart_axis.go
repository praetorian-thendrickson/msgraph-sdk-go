package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WorkbookChartAxis provides operations to manage the drive singleton.
type WorkbookChartAxis struct {
    Entity
    // Represents the formatting of a chart object, which includes line and font formatting. Read-only.
    format WorkbookChartAxisFormatable;
    // Returns a gridlines object that represents the major gridlines for the specified axis. Read-only.
    majorGridlines WorkbookChartGridlinesable;
    // Represents the interval between two major tick marks. Can be set to a numeric value or an empty string.  The returned value is always a number.
    majorUnit Jsonable;
    // Represents the maximum value on the value axis.  Can be set to a numeric value or an empty string (for automatic axis values).  The returned value is always a number.
    maximum Jsonable;
    // Represents the minimum value on the value axis. Can be set to a numeric value or an empty string (for automatic axis values).  The returned value is always a number.
    minimum Jsonable;
    // Returns a Gridlines object that represents the minor gridlines for the specified axis. Read-only.
    minorGridlines WorkbookChartGridlinesable;
    // Represents the interval between two minor tick marks. 'Can be set to a numeric value or an empty string (for automatic axis values). The returned value is always a number.
    minorUnit Jsonable;
    // Represents the axis title. Read-only.
    title WorkbookChartAxisTitleable;
}
// NewWorkbookChartAxis instantiates a new workbookChartAxis and sets the default values.
func NewWorkbookChartAxis()(*WorkbookChartAxis) {
    m := &WorkbookChartAxis{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookChartAxisFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookChartAxisFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewWorkbookChartAxis(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookChartAxis) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["format"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookChartAxisFormatFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormat(val.(WorkbookChartAxisFormatable))
        }
        return nil
    }
    res["majorGridlines"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookChartGridlinesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMajorGridlines(val.(WorkbookChartGridlinesable))
        }
        return nil
    }
    res["majorUnit"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMajorUnit(val.(Jsonable))
        }
        return nil
    }
    res["maximum"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximum(val.(Jsonable))
        }
        return nil
    }
    res["minimum"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimum(val.(Jsonable))
        }
        return nil
    }
    res["minorGridlines"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookChartGridlinesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinorGridlines(val.(WorkbookChartGridlinesable))
        }
        return nil
    }
    res["minorUnit"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinorUnit(val.(Jsonable))
        }
        return nil
    }
    res["title"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookChartAxisTitleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val.(WorkbookChartAxisTitleable))
        }
        return nil
    }
    return res
}
// GetFormat gets the format property value. Represents the formatting of a chart object, which includes line and font formatting. Read-only.
func (m *WorkbookChartAxis) GetFormat()(WorkbookChartAxisFormatable) {
    if m == nil {
        return nil
    } else {
        return m.format
    }
}
// GetMajorGridlines gets the majorGridlines property value. Returns a gridlines object that represents the major gridlines for the specified axis. Read-only.
func (m *WorkbookChartAxis) GetMajorGridlines()(WorkbookChartGridlinesable) {
    if m == nil {
        return nil
    } else {
        return m.majorGridlines
    }
}
// GetMajorUnit gets the majorUnit property value. Represents the interval between two major tick marks. Can be set to a numeric value or an empty string.  The returned value is always a number.
func (m *WorkbookChartAxis) GetMajorUnit()(Jsonable) {
    if m == nil {
        return nil
    } else {
        return m.majorUnit
    }
}
// GetMaximum gets the maximum property value. Represents the maximum value on the value axis.  Can be set to a numeric value or an empty string (for automatic axis values).  The returned value is always a number.
func (m *WorkbookChartAxis) GetMaximum()(Jsonable) {
    if m == nil {
        return nil
    } else {
        return m.maximum
    }
}
// GetMinimum gets the minimum property value. Represents the minimum value on the value axis. Can be set to a numeric value or an empty string (for automatic axis values).  The returned value is always a number.
func (m *WorkbookChartAxis) GetMinimum()(Jsonable) {
    if m == nil {
        return nil
    } else {
        return m.minimum
    }
}
// GetMinorGridlines gets the minorGridlines property value. Returns a Gridlines object that represents the minor gridlines for the specified axis. Read-only.
func (m *WorkbookChartAxis) GetMinorGridlines()(WorkbookChartGridlinesable) {
    if m == nil {
        return nil
    } else {
        return m.minorGridlines
    }
}
// GetMinorUnit gets the minorUnit property value. Represents the interval between two minor tick marks. 'Can be set to a numeric value or an empty string (for automatic axis values). The returned value is always a number.
func (m *WorkbookChartAxis) GetMinorUnit()(Jsonable) {
    if m == nil {
        return nil
    } else {
        return m.minorUnit
    }
}
// GetTitle gets the title property value. Represents the axis title. Read-only.
func (m *WorkbookChartAxis) GetTitle()(WorkbookChartAxisTitleable) {
    if m == nil {
        return nil
    } else {
        return m.title
    }
}
func (m *WorkbookChartAxis) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *WorkbookChartAxis) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("format", m.GetFormat())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("majorGridlines", m.GetMajorGridlines())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("majorUnit", m.GetMajorUnit())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("maximum", m.GetMaximum())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("minimum", m.GetMinimum())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("minorGridlines", m.GetMinorGridlines())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("minorUnit", m.GetMinorUnit())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("title", m.GetTitle())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFormat sets the format property value. Represents the formatting of a chart object, which includes line and font formatting. Read-only.
func (m *WorkbookChartAxis) SetFormat(value WorkbookChartAxisFormatable)() {
    if m != nil {
        m.format = value
    }
}
// SetMajorGridlines sets the majorGridlines property value. Returns a gridlines object that represents the major gridlines for the specified axis. Read-only.
func (m *WorkbookChartAxis) SetMajorGridlines(value WorkbookChartGridlinesable)() {
    if m != nil {
        m.majorGridlines = value
    }
}
// SetMajorUnit sets the majorUnit property value. Represents the interval between two major tick marks. Can be set to a numeric value or an empty string.  The returned value is always a number.
func (m *WorkbookChartAxis) SetMajorUnit(value Jsonable)() {
    if m != nil {
        m.majorUnit = value
    }
}
// SetMaximum sets the maximum property value. Represents the maximum value on the value axis.  Can be set to a numeric value or an empty string (for automatic axis values).  The returned value is always a number.
func (m *WorkbookChartAxis) SetMaximum(value Jsonable)() {
    if m != nil {
        m.maximum = value
    }
}
// SetMinimum sets the minimum property value. Represents the minimum value on the value axis. Can be set to a numeric value or an empty string (for automatic axis values).  The returned value is always a number.
func (m *WorkbookChartAxis) SetMinimum(value Jsonable)() {
    if m != nil {
        m.minimum = value
    }
}
// SetMinorGridlines sets the minorGridlines property value. Returns a Gridlines object that represents the minor gridlines for the specified axis. Read-only.
func (m *WorkbookChartAxis) SetMinorGridlines(value WorkbookChartGridlinesable)() {
    if m != nil {
        m.minorGridlines = value
    }
}
// SetMinorUnit sets the minorUnit property value. Represents the interval between two minor tick marks. 'Can be set to a numeric value or an empty string (for automatic axis values). The returned value is always a number.
func (m *WorkbookChartAxis) SetMinorUnit(value Jsonable)() {
    if m != nil {
        m.minorUnit = value
    }
}
// SetTitle sets the title property value. Represents the axis title. Read-only.
func (m *WorkbookChartAxis) SetTitle(value WorkbookChartAxisTitleable)() {
    if m != nil {
        m.title = value
    }
}
