package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ThumbnailSet provides operations to manage the drive singleton.
type ThumbnailSet struct {
    Entity
    // A 1920x1920 scaled thumbnail.
    large Thumbnailable;
    // A 176x176 scaled thumbnail.
    medium Thumbnailable;
    // A 48x48 cropped thumbnail.
    small Thumbnailable;
    // A custom thumbnail image or the original image used to generate other thumbnails.
    source Thumbnailable;
}
// NewThumbnailSet instantiates a new thumbnailSet and sets the default values.
func NewThumbnailSet()(*ThumbnailSet) {
    m := &ThumbnailSet{
        Entity: *NewEntity(),
    }
    return m
}
// CreateThumbnailSetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateThumbnailSetFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewThumbnailSet(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ThumbnailSet) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["large"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateThumbnailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLarge(val.(Thumbnailable))
        }
        return nil
    }
    res["medium"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateThumbnailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMedium(val.(Thumbnailable))
        }
        return nil
    }
    res["small"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateThumbnailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSmall(val.(Thumbnailable))
        }
        return nil
    }
    res["source"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateThumbnailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSource(val.(Thumbnailable))
        }
        return nil
    }
    return res
}
// GetLarge gets the large property value. A 1920x1920 scaled thumbnail.
func (m *ThumbnailSet) GetLarge()(Thumbnailable) {
    if m == nil {
        return nil
    } else {
        return m.large
    }
}
// GetMedium gets the medium property value. A 176x176 scaled thumbnail.
func (m *ThumbnailSet) GetMedium()(Thumbnailable) {
    if m == nil {
        return nil
    } else {
        return m.medium
    }
}
// GetSmall gets the small property value. A 48x48 cropped thumbnail.
func (m *ThumbnailSet) GetSmall()(Thumbnailable) {
    if m == nil {
        return nil
    } else {
        return m.small
    }
}
// GetSource gets the source property value. A custom thumbnail image or the original image used to generate other thumbnails.
func (m *ThumbnailSet) GetSource()(Thumbnailable) {
    if m == nil {
        return nil
    } else {
        return m.source
    }
}
func (m *ThumbnailSet) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ThumbnailSet) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("large", m.GetLarge())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("medium", m.GetMedium())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("small", m.GetSmall())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("source", m.GetSource())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLarge sets the large property value. A 1920x1920 scaled thumbnail.
func (m *ThumbnailSet) SetLarge(value Thumbnailable)() {
    if m != nil {
        m.large = value
    }
}
// SetMedium sets the medium property value. A 176x176 scaled thumbnail.
func (m *ThumbnailSet) SetMedium(value Thumbnailable)() {
    if m != nil {
        m.medium = value
    }
}
// SetSmall sets the small property value. A 48x48 cropped thumbnail.
func (m *ThumbnailSet) SetSmall(value Thumbnailable)() {
    if m != nil {
        m.small = value
    }
}
// SetSource sets the source property value. A custom thumbnail image or the original image used to generate other thumbnails.
func (m *ThumbnailSet) SetSource(value Thumbnailable)() {
    if m != nil {
        m.source = value
    }
}
