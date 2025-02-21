package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AudioConferencing provides operations to manage the cloudCommunications singleton.
type AudioConferencing struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The conference id of the online meeting.
    conferenceId *string;
    // A URL to the externally-accessible web page that contains dial-in information.
    dialinUrl *string;
    // 
    tollFreeNumber *string;
    // List of toll-free numbers that are displayed in the meeting invite.
    tollFreeNumbers []string;
    // 
    tollNumber *string;
    // List of toll numbers that are displayed in the meeting invite.
    tollNumbers []string;
}
// NewAudioConferencing instantiates a new audioConferencing and sets the default values.
func NewAudioConferencing()(*AudioConferencing) {
    m := &AudioConferencing{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAudioConferencingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAudioConferencingFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewAudioConferencing(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AudioConferencing) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetConferenceId gets the conferenceId property value. The conference id of the online meeting.
func (m *AudioConferencing) GetConferenceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.conferenceId
    }
}
// GetDialinUrl gets the dialinUrl property value. A URL to the externally-accessible web page that contains dial-in information.
func (m *AudioConferencing) GetDialinUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.dialinUrl
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AudioConferencing) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["conferenceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConferenceId(val)
        }
        return nil
    }
    res["dialinUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDialinUrl(val)
        }
        return nil
    }
    res["tollFreeNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTollFreeNumber(val)
        }
        return nil
    }
    res["tollFreeNumbers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetTollFreeNumbers(res)
        }
        return nil
    }
    res["tollNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTollNumber(val)
        }
        return nil
    }
    res["tollNumbers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetTollNumbers(res)
        }
        return nil
    }
    return res
}
// GetTollFreeNumber gets the tollFreeNumber property value. 
func (m *AudioConferencing) GetTollFreeNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tollFreeNumber
    }
}
// GetTollFreeNumbers gets the tollFreeNumbers property value. List of toll-free numbers that are displayed in the meeting invite.
func (m *AudioConferencing) GetTollFreeNumbers()([]string) {
    if m == nil {
        return nil
    } else {
        return m.tollFreeNumbers
    }
}
// GetTollNumber gets the tollNumber property value. 
func (m *AudioConferencing) GetTollNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tollNumber
    }
}
// GetTollNumbers gets the tollNumbers property value. List of toll numbers that are displayed in the meeting invite.
func (m *AudioConferencing) GetTollNumbers()([]string) {
    if m == nil {
        return nil
    } else {
        return m.tollNumbers
    }
}
func (m *AudioConferencing) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AudioConferencing) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("conferenceId", m.GetConferenceId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("dialinUrl", m.GetDialinUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("tollFreeNumber", m.GetTollFreeNumber())
        if err != nil {
            return err
        }
    }
    if m.GetTollFreeNumbers() != nil {
        err := writer.WriteCollectionOfStringValues("tollFreeNumbers", m.GetTollFreeNumbers())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("tollNumber", m.GetTollNumber())
        if err != nil {
            return err
        }
    }
    if m.GetTollNumbers() != nil {
        err := writer.WriteCollectionOfStringValues("tollNumbers", m.GetTollNumbers())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AudioConferencing) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetConferenceId sets the conferenceId property value. The conference id of the online meeting.
func (m *AudioConferencing) SetConferenceId(value *string)() {
    if m != nil {
        m.conferenceId = value
    }
}
// SetDialinUrl sets the dialinUrl property value. A URL to the externally-accessible web page that contains dial-in information.
func (m *AudioConferencing) SetDialinUrl(value *string)() {
    if m != nil {
        m.dialinUrl = value
    }
}
// SetTollFreeNumber sets the tollFreeNumber property value. 
func (m *AudioConferencing) SetTollFreeNumber(value *string)() {
    if m != nil {
        m.tollFreeNumber = value
    }
}
// SetTollFreeNumbers sets the tollFreeNumbers property value. List of toll-free numbers that are displayed in the meeting invite.
func (m *AudioConferencing) SetTollFreeNumbers(value []string)() {
    if m != nil {
        m.tollFreeNumbers = value
    }
}
// SetTollNumber sets the tollNumber property value. 
func (m *AudioConferencing) SetTollNumber(value *string)() {
    if m != nil {
        m.tollNumber = value
    }
}
// SetTollNumbers sets the tollNumbers property value. List of toll numbers that are displayed in the meeting invite.
func (m *AudioConferencing) SetTollNumbers(value []string)() {
    if m != nil {
        m.tollNumbers = value
    }
}
