package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WorkbookTableColumn provides operations to manage the drive singleton.
type WorkbookTableColumn struct {
    Entity
    // Retrieve the filter applied to the column. Read-only.
    filter WorkbookFilterable;
    // Returns the index number of the column within the columns collection of the table. Zero-indexed. Read-only.
    index *int32;
    // Returns the name of the table column.
    name *string;
    // Represents the raw values of the specified range. The data returned could be of type string, number, or a boolean. Cell that contain an error will return the error string.
    values Jsonable;
}
// NewWorkbookTableColumn instantiates a new workbookTableColumn and sets the default values.
func NewWorkbookTableColumn()(*WorkbookTableColumn) {
    m := &WorkbookTableColumn{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookTableColumnFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookTableColumnFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewWorkbookTableColumn(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookTableColumn) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["filter"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookFilterFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFilter(val.(WorkbookFilterable))
        }
        return nil
    }
    res["index"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIndex(val)
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
    res["values"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValues(val.(Jsonable))
        }
        return nil
    }
    return res
}
// GetFilter gets the filter property value. Retrieve the filter applied to the column. Read-only.
func (m *WorkbookTableColumn) GetFilter()(WorkbookFilterable) {
    if m == nil {
        return nil
    } else {
        return m.filter
    }
}
// GetIndex gets the index property value. Returns the index number of the column within the columns collection of the table. Zero-indexed. Read-only.
func (m *WorkbookTableColumn) GetIndex()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.index
    }
}
// GetName gets the name property value. Returns the name of the table column.
func (m *WorkbookTableColumn) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetValues gets the values property value. Represents the raw values of the specified range. The data returned could be of type string, number, or a boolean. Cell that contain an error will return the error string.
func (m *WorkbookTableColumn) GetValues()(Jsonable) {
    if m == nil {
        return nil
    } else {
        return m.values
    }
}
func (m *WorkbookTableColumn) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *WorkbookTableColumn) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("filter", m.GetFilter())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("index", m.GetIndex())
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
        err = writer.WriteObjectValue("values", m.GetValues())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFilter sets the filter property value. Retrieve the filter applied to the column. Read-only.
func (m *WorkbookTableColumn) SetFilter(value WorkbookFilterable)() {
    if m != nil {
        m.filter = value
    }
}
// SetIndex sets the index property value. Returns the index number of the column within the columns collection of the table. Zero-indexed. Read-only.
func (m *WorkbookTableColumn) SetIndex(value *int32)() {
    if m != nil {
        m.index = value
    }
}
// SetName sets the name property value. Returns the name of the table column.
func (m *WorkbookTableColumn) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetValues sets the values property value. Represents the raw values of the specified range. The data returned could be of type string, number, or a boolean. Cell that contain an error will return the error string.
func (m *WorkbookTableColumn) SetValues(value Jsonable)() {
    if m != nil {
        m.values = value
    }
}
