package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MailFolder provides operations to manage the drive singleton.
type MailFolder struct {
    Entity
    // The number of immediate child mailFolders in the current mailFolder.
    childFolderCount *int32;
    // The collection of child folders in the mailFolder.
    childFolders []MailFolderable;
    // The mailFolder's display name.
    displayName *string;
    // Indicates whether the mailFolder is hidden. This property can be set only when creating the folder. Find more information in Hidden mail folders.
    isHidden *bool;
    // The collection of rules that apply to the user's Inbox folder.
    messageRules []MessageRuleable;
    // The collection of messages in the mailFolder.
    messages []Messageable;
    // The collection of multi-value extended properties defined for the mailFolder. Read-only. Nullable.
    multiValueExtendedProperties []MultiValueLegacyExtendedPropertyable;
    // The unique identifier for the mailFolder's parent mailFolder.
    parentFolderId *string;
    // The collection of single-value extended properties defined for the mailFolder. Read-only. Nullable.
    singleValueExtendedProperties []SingleValueLegacyExtendedPropertyable;
    // The number of items in the mailFolder.
    totalItemCount *int32;
    // The number of items in the mailFolder marked as unread.
    unreadItemCount *int32;
}
// NewMailFolder instantiates a new mailFolder and sets the default values.
func NewMailFolder()(*MailFolder) {
    m := &MailFolder{
        Entity: *NewEntity(),
    }
    return m
}
// CreateMailFolderFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMailFolderFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewMailFolder(), nil
}
// GetChildFolderCount gets the childFolderCount property value. The number of immediate child mailFolders in the current mailFolder.
func (m *MailFolder) GetChildFolderCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.childFolderCount
    }
}
// GetChildFolders gets the childFolders property value. The collection of child folders in the mailFolder.
func (m *MailFolder) GetChildFolders()([]MailFolderable) {
    if m == nil {
        return nil
    } else {
        return m.childFolders
    }
}
// GetDisplayName gets the displayName property value. The mailFolder's display name.
func (m *MailFolder) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MailFolder) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["childFolderCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChildFolderCount(val)
        }
        return nil
    }
    res["childFolders"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMailFolderFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MailFolderable, len(val))
            for i, v := range val {
                res[i] = v.(MailFolderable)
            }
            m.SetChildFolders(res)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["isHidden"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsHidden(val)
        }
        return nil
    }
    res["messageRules"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMessageRuleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MessageRuleable, len(val))
            for i, v := range val {
                res[i] = v.(MessageRuleable)
            }
            m.SetMessageRules(res)
        }
        return nil
    }
    res["messages"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMessageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Messageable, len(val))
            for i, v := range val {
                res[i] = v.(Messageable)
            }
            m.SetMessages(res)
        }
        return nil
    }
    res["multiValueExtendedProperties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMultiValueLegacyExtendedPropertyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MultiValueLegacyExtendedPropertyable, len(val))
            for i, v := range val {
                res[i] = v.(MultiValueLegacyExtendedPropertyable)
            }
            m.SetMultiValueExtendedProperties(res)
        }
        return nil
    }
    res["parentFolderId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentFolderId(val)
        }
        return nil
    }
    res["singleValueExtendedProperties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSingleValueLegacyExtendedPropertyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SingleValueLegacyExtendedPropertyable, len(val))
            for i, v := range val {
                res[i] = v.(SingleValueLegacyExtendedPropertyable)
            }
            m.SetSingleValueExtendedProperties(res)
        }
        return nil
    }
    res["totalItemCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalItemCount(val)
        }
        return nil
    }
    res["unreadItemCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnreadItemCount(val)
        }
        return nil
    }
    return res
}
// GetIsHidden gets the isHidden property value. Indicates whether the mailFolder is hidden. This property can be set only when creating the folder. Find more information in Hidden mail folders.
func (m *MailFolder) GetIsHidden()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isHidden
    }
}
// GetMessageRules gets the messageRules property value. The collection of rules that apply to the user's Inbox folder.
func (m *MailFolder) GetMessageRules()([]MessageRuleable) {
    if m == nil {
        return nil
    } else {
        return m.messageRules
    }
}
// GetMessages gets the messages property value. The collection of messages in the mailFolder.
func (m *MailFolder) GetMessages()([]Messageable) {
    if m == nil {
        return nil
    } else {
        return m.messages
    }
}
// GetMultiValueExtendedProperties gets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the mailFolder. Read-only. Nullable.
func (m *MailFolder) GetMultiValueExtendedProperties()([]MultiValueLegacyExtendedPropertyable) {
    if m == nil {
        return nil
    } else {
        return m.multiValueExtendedProperties
    }
}
// GetParentFolderId gets the parentFolderId property value. The unique identifier for the mailFolder's parent mailFolder.
func (m *MailFolder) GetParentFolderId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.parentFolderId
    }
}
// GetSingleValueExtendedProperties gets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the mailFolder. Read-only. Nullable.
func (m *MailFolder) GetSingleValueExtendedProperties()([]SingleValueLegacyExtendedPropertyable) {
    if m == nil {
        return nil
    } else {
        return m.singleValueExtendedProperties
    }
}
// GetTotalItemCount gets the totalItemCount property value. The number of items in the mailFolder.
func (m *MailFolder) GetTotalItemCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.totalItemCount
    }
}
// GetUnreadItemCount gets the unreadItemCount property value. The number of items in the mailFolder marked as unread.
func (m *MailFolder) GetUnreadItemCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.unreadItemCount
    }
}
func (m *MailFolder) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *MailFolder) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("childFolderCount", m.GetChildFolderCount())
        if err != nil {
            return err
        }
    }
    if m.GetChildFolders() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetChildFolders()))
        for i, v := range m.GetChildFolders() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("childFolders", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isHidden", m.GetIsHidden())
        if err != nil {
            return err
        }
    }
    if m.GetMessageRules() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMessageRules()))
        for i, v := range m.GetMessageRules() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("messageRules", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMessages() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMessages()))
        for i, v := range m.GetMessages() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("messages", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMultiValueExtendedProperties() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMultiValueExtendedProperties()))
        for i, v := range m.GetMultiValueExtendedProperties() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("multiValueExtendedProperties", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("parentFolderId", m.GetParentFolderId())
        if err != nil {
            return err
        }
    }
    if m.GetSingleValueExtendedProperties() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSingleValueExtendedProperties()))
        for i, v := range m.GetSingleValueExtendedProperties() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("singleValueExtendedProperties", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("totalItemCount", m.GetTotalItemCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("unreadItemCount", m.GetUnreadItemCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetChildFolderCount sets the childFolderCount property value. The number of immediate child mailFolders in the current mailFolder.
func (m *MailFolder) SetChildFolderCount(value *int32)() {
    if m != nil {
        m.childFolderCount = value
    }
}
// SetChildFolders sets the childFolders property value. The collection of child folders in the mailFolder.
func (m *MailFolder) SetChildFolders(value []MailFolderable)() {
    if m != nil {
        m.childFolders = value
    }
}
// SetDisplayName sets the displayName property value. The mailFolder's display name.
func (m *MailFolder) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetIsHidden sets the isHidden property value. Indicates whether the mailFolder is hidden. This property can be set only when creating the folder. Find more information in Hidden mail folders.
func (m *MailFolder) SetIsHidden(value *bool)() {
    if m != nil {
        m.isHidden = value
    }
}
// SetMessageRules sets the messageRules property value. The collection of rules that apply to the user's Inbox folder.
func (m *MailFolder) SetMessageRules(value []MessageRuleable)() {
    if m != nil {
        m.messageRules = value
    }
}
// SetMessages sets the messages property value. The collection of messages in the mailFolder.
func (m *MailFolder) SetMessages(value []Messageable)() {
    if m != nil {
        m.messages = value
    }
}
// SetMultiValueExtendedProperties sets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the mailFolder. Read-only. Nullable.
func (m *MailFolder) SetMultiValueExtendedProperties(value []MultiValueLegacyExtendedPropertyable)() {
    if m != nil {
        m.multiValueExtendedProperties = value
    }
}
// SetParentFolderId sets the parentFolderId property value. The unique identifier for the mailFolder's parent mailFolder.
func (m *MailFolder) SetParentFolderId(value *string)() {
    if m != nil {
        m.parentFolderId = value
    }
}
// SetSingleValueExtendedProperties sets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the mailFolder. Read-only. Nullable.
func (m *MailFolder) SetSingleValueExtendedProperties(value []SingleValueLegacyExtendedPropertyable)() {
    if m != nil {
        m.singleValueExtendedProperties = value
    }
}
// SetTotalItemCount sets the totalItemCount property value. The number of items in the mailFolder.
func (m *MailFolder) SetTotalItemCount(value *int32)() {
    if m != nil {
        m.totalItemCount = value
    }
}
// SetUnreadItemCount sets the unreadItemCount property value. The number of items in the mailFolder marked as unread.
func (m *MailFolder) SetUnreadItemCount(value *int32)() {
    if m != nil {
        m.unreadItemCount = value
    }
}
