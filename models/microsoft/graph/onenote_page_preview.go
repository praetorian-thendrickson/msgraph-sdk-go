package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OnenotePagePreview provides operations to call the preview method.
type OnenotePagePreview struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    links OnenotePagePreviewLinksable;
    // 
    previewText *string;
}
// NewOnenotePagePreview instantiates a new onenotePagePreview and sets the default values.
func NewOnenotePagePreview()(*OnenotePagePreview) {
    m := &OnenotePagePreview{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateOnenotePagePreviewFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnenotePagePreviewFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewOnenotePagePreview(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OnenotePagePreview) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnenotePagePreview) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["links"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateOnenotePagePreviewLinksFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLinks(val.(OnenotePagePreviewLinksable))
        }
        return nil
    }
    res["previewText"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreviewText(val)
        }
        return nil
    }
    return res
}
// GetLinks gets the links property value. 
func (m *OnenotePagePreview) GetLinks()(OnenotePagePreviewLinksable) {
    if m == nil {
        return nil
    } else {
        return m.links
    }
}
// GetPreviewText gets the previewText property value. 
func (m *OnenotePagePreview) GetPreviewText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.previewText
    }
}
func (m *OnenotePagePreview) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *OnenotePagePreview) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("links", m.GetLinks())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("previewText", m.GetPreviewText())
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
func (m *OnenotePagePreview) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetLinks sets the links property value. 
func (m *OnenotePagePreview) SetLinks(value OnenotePagePreviewLinksable)() {
    if m != nil {
        m.links = value
    }
}
// SetPreviewText sets the previewText property value. 
func (m *OnenotePagePreview) SetPreviewText(value *string)() {
    if m != nil {
        m.previewText = value
    }
}
