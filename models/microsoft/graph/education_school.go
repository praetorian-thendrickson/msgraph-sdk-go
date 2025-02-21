package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// EducationSchool provides operations to manage the educationRoot singleton.
type EducationSchool struct {
    EducationOrganization
    // Address of the school.
    address PhysicalAddressable;
    // The underlying administrativeUnit for this school.
    administrativeUnit AdministrativeUnitable;
    // Classes taught at the school. Nullable.
    classes []EducationClassable;
    // Entity who created the school.
    createdBy IdentitySetable;
    // ID of school in syncing system.
    externalId *string;
    // ID of principal in syncing system.
    externalPrincipalId *string;
    // 
    fax *string;
    // Highest grade taught.
    highestGrade *string;
    // Lowest grade taught.
    lowestGrade *string;
    // Phone number of school.
    phone *string;
    // Email address of the principal.
    principalEmail *string;
    // Name of the principal.
    principalName *string;
    // School Number.
    schoolNumber *string;
    // Users in the school. Nullable.
    users []EducationUserable;
}
// NewEducationSchool instantiates a new educationSchool and sets the default values.
func NewEducationSchool()(*EducationSchool) {
    m := &EducationSchool{
        EducationOrganization: *NewEducationOrganization(),
    }
    return m
}
// CreateEducationSchoolFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationSchoolFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewEducationSchool(), nil
}
// GetAddress gets the address property value. Address of the school.
func (m *EducationSchool) GetAddress()(PhysicalAddressable) {
    if m == nil {
        return nil
    } else {
        return m.address
    }
}
// GetAdministrativeUnit gets the administrativeUnit property value. The underlying administrativeUnit for this school.
func (m *EducationSchool) GetAdministrativeUnit()(AdministrativeUnitable) {
    if m == nil {
        return nil
    } else {
        return m.administrativeUnit
    }
}
// GetClasses gets the classes property value. Classes taught at the school. Nullable.
func (m *EducationSchool) GetClasses()([]EducationClassable) {
    if m == nil {
        return nil
    } else {
        return m.classes
    }
}
// GetCreatedBy gets the createdBy property value. Entity who created the school.
func (m *EducationSchool) GetCreatedBy()(IdentitySetable) {
    if m == nil {
        return nil
    } else {
        return m.createdBy
    }
}
// GetExternalId gets the externalId property value. ID of school in syncing system.
func (m *EducationSchool) GetExternalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalId
    }
}
// GetExternalPrincipalId gets the externalPrincipalId property value. ID of principal in syncing system.
func (m *EducationSchool) GetExternalPrincipalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalPrincipalId
    }
}
// GetFax gets the fax property value. 
func (m *EducationSchool) GetFax()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fax
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationSchool) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.EducationOrganization.GetFieldDeserializers()
    res["address"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreatePhysicalAddressFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddress(val.(PhysicalAddressable))
        }
        return nil
    }
    res["administrativeUnit"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateAdministrativeUnitFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdministrativeUnit(val.(AdministrativeUnitable))
        }
        return nil
    }
    res["classes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEducationClassFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EducationClassable, len(val))
            for i, v := range val {
                res[i] = v.(EducationClassable)
            }
            m.SetClasses(res)
        }
        return nil
    }
    res["createdBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(IdentitySetable))
        }
        return nil
    }
    res["externalId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalId(val)
        }
        return nil
    }
    res["externalPrincipalId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalPrincipalId(val)
        }
        return nil
    }
    res["fax"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFax(val)
        }
        return nil
    }
    res["highestGrade"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHighestGrade(val)
        }
        return nil
    }
    res["lowestGrade"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLowestGrade(val)
        }
        return nil
    }
    res["phone"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhone(val)
        }
        return nil
    }
    res["principalEmail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrincipalEmail(val)
        }
        return nil
    }
    res["principalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrincipalName(val)
        }
        return nil
    }
    res["schoolNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchoolNumber(val)
        }
        return nil
    }
    res["users"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEducationUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EducationUserable, len(val))
            for i, v := range val {
                res[i] = v.(EducationUserable)
            }
            m.SetUsers(res)
        }
        return nil
    }
    return res
}
// GetHighestGrade gets the highestGrade property value. Highest grade taught.
func (m *EducationSchool) GetHighestGrade()(*string) {
    if m == nil {
        return nil
    } else {
        return m.highestGrade
    }
}
// GetLowestGrade gets the lowestGrade property value. Lowest grade taught.
func (m *EducationSchool) GetLowestGrade()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lowestGrade
    }
}
// GetPhone gets the phone property value. Phone number of school.
func (m *EducationSchool) GetPhone()(*string) {
    if m == nil {
        return nil
    } else {
        return m.phone
    }
}
// GetPrincipalEmail gets the principalEmail property value. Email address of the principal.
func (m *EducationSchool) GetPrincipalEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.principalEmail
    }
}
// GetPrincipalName gets the principalName property value. Name of the principal.
func (m *EducationSchool) GetPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.principalName
    }
}
// GetSchoolNumber gets the schoolNumber property value. School Number.
func (m *EducationSchool) GetSchoolNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.schoolNumber
    }
}
// GetUsers gets the users property value. Users in the school. Nullable.
func (m *EducationSchool) GetUsers()([]EducationUserable) {
    if m == nil {
        return nil
    } else {
        return m.users
    }
}
func (m *EducationSchool) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *EducationSchool) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.EducationOrganization.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("address", m.GetAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("administrativeUnit", m.GetAdministrativeUnit())
        if err != nil {
            return err
        }
    }
    if m.GetClasses() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetClasses()))
        for i, v := range m.GetClasses() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("classes", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("externalId", m.GetExternalId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("externalPrincipalId", m.GetExternalPrincipalId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("fax", m.GetFax())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("highestGrade", m.GetHighestGrade())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("lowestGrade", m.GetLowestGrade())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("phone", m.GetPhone())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("principalEmail", m.GetPrincipalEmail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("principalName", m.GetPrincipalName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("schoolNumber", m.GetSchoolNumber())
        if err != nil {
            return err
        }
    }
    if m.GetUsers() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUsers()))
        for i, v := range m.GetUsers() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("users", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAddress sets the address property value. Address of the school.
func (m *EducationSchool) SetAddress(value PhysicalAddressable)() {
    if m != nil {
        m.address = value
    }
}
// SetAdministrativeUnit sets the administrativeUnit property value. The underlying administrativeUnit for this school.
func (m *EducationSchool) SetAdministrativeUnit(value AdministrativeUnitable)() {
    if m != nil {
        m.administrativeUnit = value
    }
}
// SetClasses sets the classes property value. Classes taught at the school. Nullable.
func (m *EducationSchool) SetClasses(value []EducationClassable)() {
    if m != nil {
        m.classes = value
    }
}
// SetCreatedBy sets the createdBy property value. Entity who created the school.
func (m *EducationSchool) SetCreatedBy(value IdentitySetable)() {
    if m != nil {
        m.createdBy = value
    }
}
// SetExternalId sets the externalId property value. ID of school in syncing system.
func (m *EducationSchool) SetExternalId(value *string)() {
    if m != nil {
        m.externalId = value
    }
}
// SetExternalPrincipalId sets the externalPrincipalId property value. ID of principal in syncing system.
func (m *EducationSchool) SetExternalPrincipalId(value *string)() {
    if m != nil {
        m.externalPrincipalId = value
    }
}
// SetFax sets the fax property value. 
func (m *EducationSchool) SetFax(value *string)() {
    if m != nil {
        m.fax = value
    }
}
// SetHighestGrade sets the highestGrade property value. Highest grade taught.
func (m *EducationSchool) SetHighestGrade(value *string)() {
    if m != nil {
        m.highestGrade = value
    }
}
// SetLowestGrade sets the lowestGrade property value. Lowest grade taught.
func (m *EducationSchool) SetLowestGrade(value *string)() {
    if m != nil {
        m.lowestGrade = value
    }
}
// SetPhone sets the phone property value. Phone number of school.
func (m *EducationSchool) SetPhone(value *string)() {
    if m != nil {
        m.phone = value
    }
}
// SetPrincipalEmail sets the principalEmail property value. Email address of the principal.
func (m *EducationSchool) SetPrincipalEmail(value *string)() {
    if m != nil {
        m.principalEmail = value
    }
}
// SetPrincipalName sets the principalName property value. Name of the principal.
func (m *EducationSchool) SetPrincipalName(value *string)() {
    if m != nil {
        m.principalName = value
    }
}
// SetSchoolNumber sets the schoolNumber property value. School Number.
func (m *EducationSchool) SetSchoolNumber(value *string)() {
    if m != nil {
        m.schoolNumber = value
    }
}
// SetUsers sets the users property value. Users in the school. Nullable.
func (m *EducationSchool) SetUsers(value []EducationUserable)() {
    if m != nil {
        m.users = value
    }
}
