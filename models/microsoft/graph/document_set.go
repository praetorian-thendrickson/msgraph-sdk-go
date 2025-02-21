package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DocumentSet provides operations to manage the drive singleton.
type DocumentSet struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Content types allowed in document set.
    allowedContentTypes []ContentTypeInfoable;
    // Default contents of document set.
    defaultContents []DocumentSetContentable;
    // Specifies whether to push welcome page changes to inherited content types.
    propagateWelcomePageChanges *bool;
    // 
    sharedColumns []ColumnDefinitionable;
    // Add the name of the document set to each file name.
    shouldPrefixNameToFile *bool;
    // 
    welcomePageColumns []ColumnDefinitionable;
    // Welcome page absolute URL.
    welcomePageUrl *string;
}
// NewDocumentSet instantiates a new documentSet and sets the default values.
func NewDocumentSet()(*DocumentSet) {
    m := &DocumentSet{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDocumentSetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDocumentSetFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewDocumentSet(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DocumentSet) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAllowedContentTypes gets the allowedContentTypes property value. Content types allowed in document set.
func (m *DocumentSet) GetAllowedContentTypes()([]ContentTypeInfoable) {
    if m == nil {
        return nil
    } else {
        return m.allowedContentTypes
    }
}
// GetDefaultContents gets the defaultContents property value. Default contents of document set.
func (m *DocumentSet) GetDefaultContents()([]DocumentSetContentable) {
    if m == nil {
        return nil
    } else {
        return m.defaultContents
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DocumentSet) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["allowedContentTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateContentTypeInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ContentTypeInfoable, len(val))
            for i, v := range val {
                res[i] = v.(ContentTypeInfoable)
            }
            m.SetAllowedContentTypes(res)
        }
        return nil
    }
    res["defaultContents"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDocumentSetContentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DocumentSetContentable, len(val))
            for i, v := range val {
                res[i] = v.(DocumentSetContentable)
            }
            m.SetDefaultContents(res)
        }
        return nil
    }
    res["propagateWelcomePageChanges"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPropagateWelcomePageChanges(val)
        }
        return nil
    }
    res["sharedColumns"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateColumnDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ColumnDefinitionable, len(val))
            for i, v := range val {
                res[i] = v.(ColumnDefinitionable)
            }
            m.SetSharedColumns(res)
        }
        return nil
    }
    res["shouldPrefixNameToFile"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShouldPrefixNameToFile(val)
        }
        return nil
    }
    res["welcomePageColumns"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateColumnDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ColumnDefinitionable, len(val))
            for i, v := range val {
                res[i] = v.(ColumnDefinitionable)
            }
            m.SetWelcomePageColumns(res)
        }
        return nil
    }
    res["welcomePageUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWelcomePageUrl(val)
        }
        return nil
    }
    return res
}
// GetPropagateWelcomePageChanges gets the propagateWelcomePageChanges property value. Specifies whether to push welcome page changes to inherited content types.
func (m *DocumentSet) GetPropagateWelcomePageChanges()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.propagateWelcomePageChanges
    }
}
// GetSharedColumns gets the sharedColumns property value. 
func (m *DocumentSet) GetSharedColumns()([]ColumnDefinitionable) {
    if m == nil {
        return nil
    } else {
        return m.sharedColumns
    }
}
// GetShouldPrefixNameToFile gets the shouldPrefixNameToFile property value. Add the name of the document set to each file name.
func (m *DocumentSet) GetShouldPrefixNameToFile()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.shouldPrefixNameToFile
    }
}
// GetWelcomePageColumns gets the welcomePageColumns property value. 
func (m *DocumentSet) GetWelcomePageColumns()([]ColumnDefinitionable) {
    if m == nil {
        return nil
    } else {
        return m.welcomePageColumns
    }
}
// GetWelcomePageUrl gets the welcomePageUrl property value. Welcome page absolute URL.
func (m *DocumentSet) GetWelcomePageUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.welcomePageUrl
    }
}
func (m *DocumentSet) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DocumentSet) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetAllowedContentTypes() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAllowedContentTypes()))
        for i, v := range m.GetAllowedContentTypes() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("allowedContentTypes", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefaultContents() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDefaultContents()))
        for i, v := range m.GetDefaultContents() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("defaultContents", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("propagateWelcomePageChanges", m.GetPropagateWelcomePageChanges())
        if err != nil {
            return err
        }
    }
    if m.GetSharedColumns() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSharedColumns()))
        for i, v := range m.GetSharedColumns() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("sharedColumns", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("shouldPrefixNameToFile", m.GetShouldPrefixNameToFile())
        if err != nil {
            return err
        }
    }
    if m.GetWelcomePageColumns() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetWelcomePageColumns()))
        for i, v := range m.GetWelcomePageColumns() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("welcomePageColumns", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("welcomePageUrl", m.GetWelcomePageUrl())
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
func (m *DocumentSet) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAllowedContentTypes sets the allowedContentTypes property value. Content types allowed in document set.
func (m *DocumentSet) SetAllowedContentTypes(value []ContentTypeInfoable)() {
    if m != nil {
        m.allowedContentTypes = value
    }
}
// SetDefaultContents sets the defaultContents property value. Default contents of document set.
func (m *DocumentSet) SetDefaultContents(value []DocumentSetContentable)() {
    if m != nil {
        m.defaultContents = value
    }
}
// SetPropagateWelcomePageChanges sets the propagateWelcomePageChanges property value. Specifies whether to push welcome page changes to inherited content types.
func (m *DocumentSet) SetPropagateWelcomePageChanges(value *bool)() {
    if m != nil {
        m.propagateWelcomePageChanges = value
    }
}
// SetSharedColumns sets the sharedColumns property value. 
func (m *DocumentSet) SetSharedColumns(value []ColumnDefinitionable)() {
    if m != nil {
        m.sharedColumns = value
    }
}
// SetShouldPrefixNameToFile sets the shouldPrefixNameToFile property value. Add the name of the document set to each file name.
func (m *DocumentSet) SetShouldPrefixNameToFile(value *bool)() {
    if m != nil {
        m.shouldPrefixNameToFile = value
    }
}
// SetWelcomePageColumns sets the welcomePageColumns property value. 
func (m *DocumentSet) SetWelcomePageColumns(value []ColumnDefinitionable)() {
    if m != nil {
        m.welcomePageColumns = value
    }
}
// SetWelcomePageUrl sets the welcomePageUrl property value. Welcome page absolute URL.
func (m *DocumentSet) SetWelcomePageUrl(value *string)() {
    if m != nil {
        m.welcomePageUrl = value
    }
}
