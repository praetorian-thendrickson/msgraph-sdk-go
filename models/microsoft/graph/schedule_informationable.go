package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ScheduleInformationable 
type ScheduleInformationable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAvailabilityView()(*string)
    GetError()(FreeBusyErrorable)
    GetScheduleId()(*string)
    GetScheduleItems()([]ScheduleItemable)
    GetWorkingHours()(WorkingHoursable)
    SetAvailabilityView(value *string)()
    SetError(value FreeBusyErrorable)()
    SetScheduleId(value *string)()
    SetScheduleItems(value []ScheduleItemable)()
    SetWorkingHours(value WorkingHoursable)()
}
