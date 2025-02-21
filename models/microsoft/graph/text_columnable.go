package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TextColumnable 
type TextColumnable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAllowMultipleLines()(*bool)
    GetAppendChangesToExistingText()(*bool)
    GetLinesForEditing()(*int32)
    GetMaxLength()(*int32)
    GetTextType()(*string)
    SetAllowMultipleLines(value *bool)()
    SetAppendChangesToExistingText(value *bool)()
    SetLinesForEditing(value *int32)()
    SetMaxLength(value *int32)()
    SetTextType(value *string)()
}
