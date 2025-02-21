package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// LookupColumnable 
type LookupColumnable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAllowMultipleValues()(*bool)
    GetAllowUnlimitedLength()(*bool)
    GetColumnName()(*string)
    GetListId()(*string)
    GetPrimaryLookupColumnId()(*string)
    SetAllowMultipleValues(value *bool)()
    SetAllowUnlimitedLength(value *bool)()
    SetColumnName(value *string)()
    SetListId(value *string)()
    SetPrimaryLookupColumnId(value *string)()
}
