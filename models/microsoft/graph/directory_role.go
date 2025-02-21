package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DirectoryRole provides operations to manage the collection of directoryRole entities.
type DirectoryRole struct {
    DirectoryObject
    // The description for the directory role. Read-only. Supports $filter (eq), $search, $select.
    description *string;
    // The display name for the directory role. Read-only. Supports $filter (eq), $search, $select.
    displayName *string;
    // Users that are members of this directory role. HTTP Methods: GET, POST, DELETE. Read-only. Nullable. Supports $expand.
    members []DirectoryObjectable;
    // The id of the directoryRoleTemplate that this role is based on. The property must be specified when activating a directory role in a tenant with a POST operation. After the directory role has been activated, the property is read only. Supports $filter (eq), $select.
    roleTemplateId *string;
    // Members of this directory role that are scoped to administrative units. Read-only. Nullable.
    scopedMembers []ScopedRoleMembershipable;
}
// NewDirectoryRole instantiates a new directoryRole and sets the default values.
func NewDirectoryRole()(*DirectoryRole) {
    m := &DirectoryRole{
        DirectoryObject: *NewDirectoryObject(),
    }
    return m
}
// CreateDirectoryRoleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDirectoryRoleFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewDirectoryRole(), nil
}
// GetDescription gets the description property value. The description for the directory role. Read-only. Supports $filter (eq), $search, $select.
func (m *DirectoryRole) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. The display name for the directory role. Read-only. Supports $filter (eq), $search, $select.
func (m *DirectoryRole) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DirectoryRole) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
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
    res["members"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDirectoryObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObjectable, len(val))
            for i, v := range val {
                res[i] = v.(DirectoryObjectable)
            }
            m.SetMembers(res)
        }
        return nil
    }
    res["roleTemplateId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleTemplateId(val)
        }
        return nil
    }
    res["scopedMembers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateScopedRoleMembershipFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ScopedRoleMembershipable, len(val))
            for i, v := range val {
                res[i] = v.(ScopedRoleMembershipable)
            }
            m.SetScopedMembers(res)
        }
        return nil
    }
    return res
}
// GetMembers gets the members property value. Users that are members of this directory role. HTTP Methods: GET, POST, DELETE. Read-only. Nullable. Supports $expand.
func (m *DirectoryRole) GetMembers()([]DirectoryObjectable) {
    if m == nil {
        return nil
    } else {
        return m.members
    }
}
// GetRoleTemplateId gets the roleTemplateId property value. The id of the directoryRoleTemplate that this role is based on. The property must be specified when activating a directory role in a tenant with a POST operation. After the directory role has been activated, the property is read only. Supports $filter (eq), $select.
func (m *DirectoryRole) GetRoleTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.roleTemplateId
    }
}
// GetScopedMembers gets the scopedMembers property value. Members of this directory role that are scoped to administrative units. Read-only. Nullable.
func (m *DirectoryRole) GetScopedMembers()([]ScopedRoleMembershipable) {
    if m == nil {
        return nil
    } else {
        return m.scopedMembers
    }
}
func (m *DirectoryRole) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DirectoryRole) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.DirectoryObject.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
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
    if m.GetMembers() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMembers()))
        for i, v := range m.GetMembers() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("members", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("roleTemplateId", m.GetRoleTemplateId())
        if err != nil {
            return err
        }
    }
    if m.GetScopedMembers() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetScopedMembers()))
        for i, v := range m.GetScopedMembers() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("scopedMembers", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDescription sets the description property value. The description for the directory role. Read-only. Supports $filter (eq), $search, $select.
func (m *DirectoryRole) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. The display name for the directory role. Read-only. Supports $filter (eq), $search, $select.
func (m *DirectoryRole) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetMembers sets the members property value. Users that are members of this directory role. HTTP Methods: GET, POST, DELETE. Read-only. Nullable. Supports $expand.
func (m *DirectoryRole) SetMembers(value []DirectoryObjectable)() {
    if m != nil {
        m.members = value
    }
}
// SetRoleTemplateId sets the roleTemplateId property value. The id of the directoryRoleTemplate that this role is based on. The property must be specified when activating a directory role in a tenant with a POST operation. After the directory role has been activated, the property is read only. Supports $filter (eq), $select.
func (m *DirectoryRole) SetRoleTemplateId(value *string)() {
    if m != nil {
        m.roleTemplateId = value
    }
}
// SetScopedMembers sets the scopedMembers property value. Members of this directory role that are scoped to administrative units. Read-only. Nullable.
func (m *DirectoryRole) SetScopedMembers(value []ScopedRoleMembershipable)() {
    if m != nil {
        m.scopedMembers = value
    }
}
