package copytonotebook

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CopyToNotebookRequestBodyable 
type CopyToNotebookRequestBodyable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetGroupId()(*string)
    GetId()(*string)
    GetRenameAs()(*string)
    GetSiteCollectionId()(*string)
    GetSiteId()(*string)
    SetGroupId(value *string)()
    SetId(value *string)()
    SetRenameAs(value *string)()
    SetSiteCollectionId(value *string)()
    SetSiteId(value *string)()
}
