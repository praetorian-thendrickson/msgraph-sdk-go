package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// EducationAssignmentDefaults provides operations to manage the educationRoot singleton.
type EducationAssignmentDefaults struct {
    Entity
    // Class-level default behavior for handling students who are added after the assignment is published. Possible values are: none, assignIfOpen.
    addedStudentAction *EducationAddedStudentAction;
    // Optional field to control adding assignments to students' and teachers' calendars when the assignment is published. The possible values are: none, studentsAndPublisher, studentsAndTeamOwners, unknownFutureValue, and studentsOnly. Note that you must use the Prefer: include-unknown-enum-members request header to get the following value(s) in this evolvable enum: studentsOnly. The default value is none.
    addToCalendarAction *EducationAddToCalendarOptions;
    // Class-level default value for due time field. Default value is 23:59:00.
    dueTime *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.TimeOnly;
    // Default Teams channel to which notifications will be sent. Default value is null.
    notificationChannelUrl *string;
}
// NewEducationAssignmentDefaults instantiates a new educationAssignmentDefaults and sets the default values.
func NewEducationAssignmentDefaults()(*EducationAssignmentDefaults) {
    m := &EducationAssignmentDefaults{
        Entity: *NewEntity(),
    }
    return m
}
// CreateEducationAssignmentDefaultsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationAssignmentDefaultsFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewEducationAssignmentDefaults(), nil
}
// GetAddedStudentAction gets the addedStudentAction property value. Class-level default behavior for handling students who are added after the assignment is published. Possible values are: none, assignIfOpen.
func (m *EducationAssignmentDefaults) GetAddedStudentAction()(*EducationAddedStudentAction) {
    if m == nil {
        return nil
    } else {
        return m.addedStudentAction
    }
}
// GetAddToCalendarAction gets the addToCalendarAction property value. Optional field to control adding assignments to students' and teachers' calendars when the assignment is published. The possible values are: none, studentsAndPublisher, studentsAndTeamOwners, unknownFutureValue, and studentsOnly. Note that you must use the Prefer: include-unknown-enum-members request header to get the following value(s) in this evolvable enum: studentsOnly. The default value is none.
func (m *EducationAssignmentDefaults) GetAddToCalendarAction()(*EducationAddToCalendarOptions) {
    if m == nil {
        return nil
    } else {
        return m.addToCalendarAction
    }
}
// GetDueTime gets the dueTime property value. Class-level default value for due time field. Default value is 23:59:00.
func (m *EducationAssignmentDefaults) GetDueTime()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.TimeOnly) {
    if m == nil {
        return nil
    } else {
        return m.dueTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationAssignmentDefaults) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["addedStudentAction"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseEducationAddedStudentAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddedStudentAction(val.(*EducationAddedStudentAction))
        }
        return nil
    }
    res["addToCalendarAction"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseEducationAddToCalendarOptions)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddToCalendarAction(val.(*EducationAddToCalendarOptions))
        }
        return nil
    }
    res["dueTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDueTime(val)
        }
        return nil
    }
    res["notificationChannelUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationChannelUrl(val)
        }
        return nil
    }
    return res
}
// GetNotificationChannelUrl gets the notificationChannelUrl property value. Default Teams channel to which notifications will be sent. Default value is null.
func (m *EducationAssignmentDefaults) GetNotificationChannelUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.notificationChannelUrl
    }
}
func (m *EducationAssignmentDefaults) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *EducationAssignmentDefaults) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAddedStudentAction() != nil {
        cast := (*m.GetAddedStudentAction()).String()
        err = writer.WriteStringValue("addedStudentAction", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAddToCalendarAction() != nil {
        cast := (*m.GetAddToCalendarAction()).String()
        err = writer.WriteStringValue("addToCalendarAction", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeOnlyValue("dueTime", m.GetDueTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("notificationChannelUrl", m.GetNotificationChannelUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAddedStudentAction sets the addedStudentAction property value. Class-level default behavior for handling students who are added after the assignment is published. Possible values are: none, assignIfOpen.
func (m *EducationAssignmentDefaults) SetAddedStudentAction(value *EducationAddedStudentAction)() {
    if m != nil {
        m.addedStudentAction = value
    }
}
// SetAddToCalendarAction sets the addToCalendarAction property value. Optional field to control adding assignments to students' and teachers' calendars when the assignment is published. The possible values are: none, studentsAndPublisher, studentsAndTeamOwners, unknownFutureValue, and studentsOnly. Note that you must use the Prefer: include-unknown-enum-members request header to get the following value(s) in this evolvable enum: studentsOnly. The default value is none.
func (m *EducationAssignmentDefaults) SetAddToCalendarAction(value *EducationAddToCalendarOptions)() {
    if m != nil {
        m.addToCalendarAction = value
    }
}
// SetDueTime sets the dueTime property value. Class-level default value for due time field. Default value is 23:59:00.
func (m *EducationAssignmentDefaults) SetDueTime(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.TimeOnly)() {
    if m != nil {
        m.dueTime = value
    }
}
// SetNotificationChannelUrl sets the notificationChannelUrl property value. Default Teams channel to which notifications will be sent. Default value is null.
func (m *EducationAssignmentDefaults) SetNotificationChannelUrl(value *string)() {
    if m != nil {
        m.notificationChannelUrl = value
    }
}
