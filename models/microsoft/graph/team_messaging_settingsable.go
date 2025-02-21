package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamMessagingSettingsable 
type TeamMessagingSettingsable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAllowChannelMentions()(*bool)
    GetAllowOwnerDeleteMessages()(*bool)
    GetAllowTeamMentions()(*bool)
    GetAllowUserDeleteMessages()(*bool)
    GetAllowUserEditMessages()(*bool)
    SetAllowChannelMentions(value *bool)()
    SetAllowOwnerDeleteMessages(value *bool)()
    SetAllowTeamMentions(value *bool)()
    SetAllowUserDeleteMessages(value *bool)()
    SetAllowUserEditMessages(value *bool)()
}
