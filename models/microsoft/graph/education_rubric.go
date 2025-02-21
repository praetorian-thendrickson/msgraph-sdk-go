package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// EducationRubric provides operations to manage the educationRoot singleton.
type EducationRubric struct {
    Entity
    // The user who created this resource.
    createdBy IdentitySetable;
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The description of this rubric.
    description EducationItemBodyable;
    // The name of this rubric.
    displayName *string;
    // The grading type of this rubric -- null for a no-points rubric, or educationAssignmentPointsGradeType for a points rubric.
    grading EducationAssignmentGradeTypeable;
    // The last user to modify the resource.
    lastModifiedBy IdentitySetable;
    // Moment in time when the resource was last modified.  The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The collection of levels making up this rubric.
    levels []RubricLevelable;
    // The collection of qualities making up this rubric.
    qualities []RubricQualityable;
}
// NewEducationRubric instantiates a new educationRubric and sets the default values.
func NewEducationRubric()(*EducationRubric) {
    m := &EducationRubric{
        Entity: *NewEntity(),
    }
    return m
}
// CreateEducationRubricFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationRubricFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewEducationRubric(), nil
}
// GetCreatedBy gets the createdBy property value. The user who created this resource.
func (m *EducationRubric) GetCreatedBy()(IdentitySetable) {
    if m == nil {
        return nil
    } else {
        return m.createdBy
    }
}
// GetCreatedDateTime gets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *EducationRubric) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDescription gets the description property value. The description of this rubric.
func (m *EducationRubric) GetDescription()(EducationItemBodyable) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. The name of this rubric.
func (m *EducationRubric) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationRubric) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateEducationItemBodyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val.(EducationItemBodyable))
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
    res["grading"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateEducationAssignmentGradeTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGrading(val.(EducationAssignmentGradeTypeable))
        }
        return nil
    }
    res["lastModifiedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedBy(val.(IdentitySetable))
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["levels"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRubricLevelFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RubricLevelable, len(val))
            for i, v := range val {
                res[i] = v.(RubricLevelable)
            }
            m.SetLevels(res)
        }
        return nil
    }
    res["qualities"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRubricQualityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RubricQualityable, len(val))
            for i, v := range val {
                res[i] = v.(RubricQualityable)
            }
            m.SetQualities(res)
        }
        return nil
    }
    return res
}
// GetGrading gets the grading property value. The grading type of this rubric -- null for a no-points rubric, or educationAssignmentPointsGradeType for a points rubric.
func (m *EducationRubric) GetGrading()(EducationAssignmentGradeTypeable) {
    if m == nil {
        return nil
    } else {
        return m.grading
    }
}
// GetLastModifiedBy gets the lastModifiedBy property value. The last user to modify the resource.
func (m *EducationRubric) GetLastModifiedBy()(IdentitySetable) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedBy
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Moment in time when the resource was last modified.  The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *EducationRubric) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetLevels gets the levels property value. The collection of levels making up this rubric.
func (m *EducationRubric) GetLevels()([]RubricLevelable) {
    if m == nil {
        return nil
    } else {
        return m.levels
    }
}
// GetQualities gets the qualities property value. The collection of qualities making up this rubric.
func (m *EducationRubric) GetQualities()([]RubricQualityable) {
    if m == nil {
        return nil
    } else {
        return m.qualities
    }
}
func (m *EducationRubric) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *EducationRubric) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("description", m.GetDescription())
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
        err = writer.WriteObjectValue("grading", m.GetGrading())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("lastModifiedBy", m.GetLastModifiedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetLevels() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetLevels()))
        for i, v := range m.GetLevels() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("levels", cast)
        if err != nil {
            return err
        }
    }
    if m.GetQualities() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetQualities()))
        for i, v := range m.GetQualities() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("qualities", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreatedBy sets the createdBy property value. The user who created this resource.
func (m *EducationRubric) SetCreatedBy(value IdentitySetable)() {
    if m != nil {
        m.createdBy = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *EducationRubric) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDescription sets the description property value. The description of this rubric.
func (m *EducationRubric) SetDescription(value EducationItemBodyable)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. The name of this rubric.
func (m *EducationRubric) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetGrading sets the grading property value. The grading type of this rubric -- null for a no-points rubric, or educationAssignmentPointsGradeType for a points rubric.
func (m *EducationRubric) SetGrading(value EducationAssignmentGradeTypeable)() {
    if m != nil {
        m.grading = value
    }
}
// SetLastModifiedBy sets the lastModifiedBy property value. The last user to modify the resource.
func (m *EducationRubric) SetLastModifiedBy(value IdentitySetable)() {
    if m != nil {
        m.lastModifiedBy = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Moment in time when the resource was last modified.  The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *EducationRubric) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetLevels sets the levels property value. The collection of levels making up this rubric.
func (m *EducationRubric) SetLevels(value []RubricLevelable)() {
    if m != nil {
        m.levels = value
    }
}
// SetQualities sets the qualities property value. The collection of qualities making up this rubric.
func (m *EducationRubric) SetQualities(value []RubricQualityable)() {
    if m != nil {
        m.qualities = value
    }
}
