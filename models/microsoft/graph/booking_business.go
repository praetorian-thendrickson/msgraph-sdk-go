package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// BookingBusiness provides operations to manage the solutionsRoot singleton.
type BookingBusiness struct {
    Entity
    // The street address of the business. The address property, together with phone and webSiteUrl, appear in the footer of a business scheduling page. The attribute type of physicalAddress is not supported in v1.0. Internally we map the addresses to the type others.
    address PhysicalAddressable;
    // All the appointments of this business. Read-only. Nullable.
    appointments []BookingAppointmentable;
    // The hours of operation for the business.
    businessHours []BookingWorkHoursable;
    // The type of business.
    businessType *string;
    // The set of appointments of this business in a specified date range. Read-only. Nullable.
    calendarView []BookingAppointmentable;
    // All the customers of this business. Read-only. Nullable.
    customers []BookingCustomerBaseable;
    // All the custom questions of this business. Read-only. Nullable.
    customQuestions []BookingCustomQuestionable;
    // The code for the currency that the business operates in on Microsoft Bookings.
    defaultCurrencyIso *string;
    // The name of the business, which interfaces with customers. This name appears at the top of the business scheduling page.
    displayName *string;
    // The email address for the business.
    email *string;
    // The scheduling page has been made available to external customers. Use the publish and unpublish actions to set this property. Read-only.
    isPublished *bool;
    // The telephone number for the business. The phone property, together with address and webSiteUrl, appear in the footer of a business scheduling page.
    phone *string;
    // The URL for the scheduling page, which is set after you publish or unpublish the page. Read-only.
    publicUrl *string;
    // Specifies how bookings can be created for this business.
    schedulingPolicy BookingSchedulingPolicyable;
    // All the services offered by this business. Read-only. Nullable.
    services []BookingServiceable;
    // All the staff members that provide services in this business. Read-only. Nullable.
    staffMembers []BookingStaffMemberBaseable;
    // The URL of the business web site. The webSiteUrl property, together with address, phone, appear in the footer of a business scheduling page.
    webSiteUrl *string;
}
// NewBookingBusiness instantiates a new bookingBusiness and sets the default values.
func NewBookingBusiness()(*BookingBusiness) {
    m := &BookingBusiness{
        Entity: *NewEntity(),
    }
    return m
}
// CreateBookingBusinessFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBookingBusinessFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewBookingBusiness(), nil
}
// GetAddress gets the address property value. The street address of the business. The address property, together with phone and webSiteUrl, appear in the footer of a business scheduling page. The attribute type of physicalAddress is not supported in v1.0. Internally we map the addresses to the type others.
func (m *BookingBusiness) GetAddress()(PhysicalAddressable) {
    if m == nil {
        return nil
    } else {
        return m.address
    }
}
// GetAppointments gets the appointments property value. All the appointments of this business. Read-only. Nullable.
func (m *BookingBusiness) GetAppointments()([]BookingAppointmentable) {
    if m == nil {
        return nil
    } else {
        return m.appointments
    }
}
// GetBusinessHours gets the businessHours property value. The hours of operation for the business.
func (m *BookingBusiness) GetBusinessHours()([]BookingWorkHoursable) {
    if m == nil {
        return nil
    } else {
        return m.businessHours
    }
}
// GetBusinessType gets the businessType property value. The type of business.
func (m *BookingBusiness) GetBusinessType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.businessType
    }
}
// GetCalendarView gets the calendarView property value. The set of appointments of this business in a specified date range. Read-only. Nullable.
func (m *BookingBusiness) GetCalendarView()([]BookingAppointmentable) {
    if m == nil {
        return nil
    } else {
        return m.calendarView
    }
}
// GetCustomers gets the customers property value. All the customers of this business. Read-only. Nullable.
func (m *BookingBusiness) GetCustomers()([]BookingCustomerBaseable) {
    if m == nil {
        return nil
    } else {
        return m.customers
    }
}
// GetCustomQuestions gets the customQuestions property value. All the custom questions of this business. Read-only. Nullable.
func (m *BookingBusiness) GetCustomQuestions()([]BookingCustomQuestionable) {
    if m == nil {
        return nil
    } else {
        return m.customQuestions
    }
}
// GetDefaultCurrencyIso gets the defaultCurrencyIso property value. The code for the currency that the business operates in on Microsoft Bookings.
func (m *BookingBusiness) GetDefaultCurrencyIso()(*string) {
    if m == nil {
        return nil
    } else {
        return m.defaultCurrencyIso
    }
}
// GetDisplayName gets the displayName property value. The name of the business, which interfaces with customers. This name appears at the top of the business scheduling page.
func (m *BookingBusiness) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetEmail gets the email property value. The email address for the business.
func (m *BookingBusiness) GetEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.email
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BookingBusiness) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["appointments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateBookingAppointmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BookingAppointmentable, len(val))
            for i, v := range val {
                res[i] = v.(BookingAppointmentable)
            }
            m.SetAppointments(res)
        }
        return nil
    }
    res["businessHours"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateBookingWorkHoursFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BookingWorkHoursable, len(val))
            for i, v := range val {
                res[i] = v.(BookingWorkHoursable)
            }
            m.SetBusinessHours(res)
        }
        return nil
    }
    res["businessType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBusinessType(val)
        }
        return nil
    }
    res["calendarView"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateBookingAppointmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BookingAppointmentable, len(val))
            for i, v := range val {
                res[i] = v.(BookingAppointmentable)
            }
            m.SetCalendarView(res)
        }
        return nil
    }
    res["customers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateBookingCustomerBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BookingCustomerBaseable, len(val))
            for i, v := range val {
                res[i] = v.(BookingCustomerBaseable)
            }
            m.SetCustomers(res)
        }
        return nil
    }
    res["customQuestions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateBookingCustomQuestionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BookingCustomQuestionable, len(val))
            for i, v := range val {
                res[i] = v.(BookingCustomQuestionable)
            }
            m.SetCustomQuestions(res)
        }
        return nil
    }
    res["defaultCurrencyIso"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultCurrencyIso(val)
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
    res["email"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmail(val)
        }
        return nil
    }
    res["isPublished"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPublished(val)
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
    res["publicUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublicUrl(val)
        }
        return nil
    }
    res["schedulingPolicy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateBookingSchedulingPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchedulingPolicy(val.(BookingSchedulingPolicyable))
        }
        return nil
    }
    res["services"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateBookingServiceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BookingServiceable, len(val))
            for i, v := range val {
                res[i] = v.(BookingServiceable)
            }
            m.SetServices(res)
        }
        return nil
    }
    res["staffMembers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateBookingStaffMemberBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BookingStaffMemberBaseable, len(val))
            for i, v := range val {
                res[i] = v.(BookingStaffMemberBaseable)
            }
            m.SetStaffMembers(res)
        }
        return nil
    }
    res["webSiteUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebSiteUrl(val)
        }
        return nil
    }
    return res
}
// GetIsPublished gets the isPublished property value. The scheduling page has been made available to external customers. Use the publish and unpublish actions to set this property. Read-only.
func (m *BookingBusiness) GetIsPublished()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isPublished
    }
}
// GetPhone gets the phone property value. The telephone number for the business. The phone property, together with address and webSiteUrl, appear in the footer of a business scheduling page.
func (m *BookingBusiness) GetPhone()(*string) {
    if m == nil {
        return nil
    } else {
        return m.phone
    }
}
// GetPublicUrl gets the publicUrl property value. The URL for the scheduling page, which is set after you publish or unpublish the page. Read-only.
func (m *BookingBusiness) GetPublicUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.publicUrl
    }
}
// GetSchedulingPolicy gets the schedulingPolicy property value. Specifies how bookings can be created for this business.
func (m *BookingBusiness) GetSchedulingPolicy()(BookingSchedulingPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.schedulingPolicy
    }
}
// GetServices gets the services property value. All the services offered by this business. Read-only. Nullable.
func (m *BookingBusiness) GetServices()([]BookingServiceable) {
    if m == nil {
        return nil
    } else {
        return m.services
    }
}
// GetStaffMembers gets the staffMembers property value. All the staff members that provide services in this business. Read-only. Nullable.
func (m *BookingBusiness) GetStaffMembers()([]BookingStaffMemberBaseable) {
    if m == nil {
        return nil
    } else {
        return m.staffMembers
    }
}
// GetWebSiteUrl gets the webSiteUrl property value. The URL of the business web site. The webSiteUrl property, together with address, phone, appear in the footer of a business scheduling page.
func (m *BookingBusiness) GetWebSiteUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webSiteUrl
    }
}
func (m *BookingBusiness) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *BookingBusiness) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("address", m.GetAddress())
        if err != nil {
            return err
        }
    }
    if m.GetAppointments() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAppointments()))
        for i, v := range m.GetAppointments() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("appointments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetBusinessHours() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetBusinessHours()))
        for i, v := range m.GetBusinessHours() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("businessHours", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("businessType", m.GetBusinessType())
        if err != nil {
            return err
        }
    }
    if m.GetCalendarView() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCalendarView()))
        for i, v := range m.GetCalendarView() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("calendarView", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCustomers() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCustomers()))
        for i, v := range m.GetCustomers() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("customers", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCustomQuestions() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCustomQuestions()))
        for i, v := range m.GetCustomQuestions() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("customQuestions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("defaultCurrencyIso", m.GetDefaultCurrencyIso())
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
        err = writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isPublished", m.GetIsPublished())
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
        err = writer.WriteStringValue("publicUrl", m.GetPublicUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("schedulingPolicy", m.GetSchedulingPolicy())
        if err != nil {
            return err
        }
    }
    if m.GetServices() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetServices()))
        for i, v := range m.GetServices() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("services", cast)
        if err != nil {
            return err
        }
    }
    if m.GetStaffMembers() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetStaffMembers()))
        for i, v := range m.GetStaffMembers() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("staffMembers", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("webSiteUrl", m.GetWebSiteUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAddress sets the address property value. The street address of the business. The address property, together with phone and webSiteUrl, appear in the footer of a business scheduling page. The attribute type of physicalAddress is not supported in v1.0. Internally we map the addresses to the type others.
func (m *BookingBusiness) SetAddress(value PhysicalAddressable)() {
    if m != nil {
        m.address = value
    }
}
// SetAppointments sets the appointments property value. All the appointments of this business. Read-only. Nullable.
func (m *BookingBusiness) SetAppointments(value []BookingAppointmentable)() {
    if m != nil {
        m.appointments = value
    }
}
// SetBusinessHours sets the businessHours property value. The hours of operation for the business.
func (m *BookingBusiness) SetBusinessHours(value []BookingWorkHoursable)() {
    if m != nil {
        m.businessHours = value
    }
}
// SetBusinessType sets the businessType property value. The type of business.
func (m *BookingBusiness) SetBusinessType(value *string)() {
    if m != nil {
        m.businessType = value
    }
}
// SetCalendarView sets the calendarView property value. The set of appointments of this business in a specified date range. Read-only. Nullable.
func (m *BookingBusiness) SetCalendarView(value []BookingAppointmentable)() {
    if m != nil {
        m.calendarView = value
    }
}
// SetCustomers sets the customers property value. All the customers of this business. Read-only. Nullable.
func (m *BookingBusiness) SetCustomers(value []BookingCustomerBaseable)() {
    if m != nil {
        m.customers = value
    }
}
// SetCustomQuestions sets the customQuestions property value. All the custom questions of this business. Read-only. Nullable.
func (m *BookingBusiness) SetCustomQuestions(value []BookingCustomQuestionable)() {
    if m != nil {
        m.customQuestions = value
    }
}
// SetDefaultCurrencyIso sets the defaultCurrencyIso property value. The code for the currency that the business operates in on Microsoft Bookings.
func (m *BookingBusiness) SetDefaultCurrencyIso(value *string)() {
    if m != nil {
        m.defaultCurrencyIso = value
    }
}
// SetDisplayName sets the displayName property value. The name of the business, which interfaces with customers. This name appears at the top of the business scheduling page.
func (m *BookingBusiness) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetEmail sets the email property value. The email address for the business.
func (m *BookingBusiness) SetEmail(value *string)() {
    if m != nil {
        m.email = value
    }
}
// SetIsPublished sets the isPublished property value. The scheduling page has been made available to external customers. Use the publish and unpublish actions to set this property. Read-only.
func (m *BookingBusiness) SetIsPublished(value *bool)() {
    if m != nil {
        m.isPublished = value
    }
}
// SetPhone sets the phone property value. The telephone number for the business. The phone property, together with address and webSiteUrl, appear in the footer of a business scheduling page.
func (m *BookingBusiness) SetPhone(value *string)() {
    if m != nil {
        m.phone = value
    }
}
// SetPublicUrl sets the publicUrl property value. The URL for the scheduling page, which is set after you publish or unpublish the page. Read-only.
func (m *BookingBusiness) SetPublicUrl(value *string)() {
    if m != nil {
        m.publicUrl = value
    }
}
// SetSchedulingPolicy sets the schedulingPolicy property value. Specifies how bookings can be created for this business.
func (m *BookingBusiness) SetSchedulingPolicy(value BookingSchedulingPolicyable)() {
    if m != nil {
        m.schedulingPolicy = value
    }
}
// SetServices sets the services property value. All the services offered by this business. Read-only. Nullable.
func (m *BookingBusiness) SetServices(value []BookingServiceable)() {
    if m != nil {
        m.services = value
    }
}
// SetStaffMembers sets the staffMembers property value. All the staff members that provide services in this business. Read-only. Nullable.
func (m *BookingBusiness) SetStaffMembers(value []BookingStaffMemberBaseable)() {
    if m != nil {
        m.staffMembers = value
    }
}
// SetWebSiteUrl sets the webSiteUrl property value. The URL of the business web site. The webSiteUrl property, together with address, phone, appear in the footer of a business scheduling page.
func (m *BookingBusiness) SetWebSiteUrl(value *string)() {
    if m != nil {
        m.webSiteUrl = value
    }
}
