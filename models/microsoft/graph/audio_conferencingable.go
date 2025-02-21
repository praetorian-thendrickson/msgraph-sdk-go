package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AudioConferencingable 
type AudioConferencingable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetConferenceId()(*string)
    GetDialinUrl()(*string)
    GetTollFreeNumber()(*string)
    GetTollFreeNumbers()([]string)
    GetTollNumber()(*string)
    GetTollNumbers()([]string)
    SetConferenceId(value *string)()
    SetDialinUrl(value *string)()
    SetTollFreeNumber(value *string)()
    SetTollFreeNumbers(value []string)()
    SetTollNumber(value *string)()
    SetTollNumbers(value []string)()
}
