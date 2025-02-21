package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Quota provides operations to manage the drive singleton.
type Quota struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Total space consumed by files in the recycle bin, in bytes. Read-only.
    deleted *int64;
    // Total space remaining before reaching the quota limit, in bytes. Read-only.
    remaining *int64;
    // Enumeration value that indicates the state of the storage space. Read-only.
    state *string;
    // Information about the drive's storage quota plans. Only in Personal OneDrive.
    storagePlanInformation StoragePlanInformationable;
    // Total allowed storage space, in bytes. Read-only.
    total *int64;
    // Total space used, in bytes. Read-only.
    used *int64;
}
// NewQuota instantiates a new quota and sets the default values.
func NewQuota()(*Quota) {
    m := &Quota{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateQuotaFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateQuotaFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewQuota(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Quota) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDeleted gets the deleted property value. Total space consumed by files in the recycle bin, in bytes. Read-only.
func (m *Quota) GetDeleted()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.deleted
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Quota) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["deleted"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeleted(val)
        }
        return nil
    }
    res["remaining"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemaining(val)
        }
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val)
        }
        return nil
    }
    res["storagePlanInformation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateStoragePlanInformationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStoragePlanInformation(val.(StoragePlanInformationable))
        }
        return nil
    }
    res["total"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotal(val)
        }
        return nil
    }
    res["used"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsed(val)
        }
        return nil
    }
    return res
}
// GetRemaining gets the remaining property value. Total space remaining before reaching the quota limit, in bytes. Read-only.
func (m *Quota) GetRemaining()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.remaining
    }
}
// GetState gets the state property value. Enumeration value that indicates the state of the storage space. Read-only.
func (m *Quota) GetState()(*string) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// GetStoragePlanInformation gets the storagePlanInformation property value. Information about the drive's storage quota plans. Only in Personal OneDrive.
func (m *Quota) GetStoragePlanInformation()(StoragePlanInformationable) {
    if m == nil {
        return nil
    } else {
        return m.storagePlanInformation
    }
}
// GetTotal gets the total property value. Total allowed storage space, in bytes. Read-only.
func (m *Quota) GetTotal()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.total
    }
}
// GetUsed gets the used property value. Total space used, in bytes. Read-only.
func (m *Quota) GetUsed()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.used
    }
}
func (m *Quota) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Quota) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("deleted", m.GetDeleted())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("remaining", m.GetRemaining())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("state", m.GetState())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("storagePlanInformation", m.GetStoragePlanInformation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("total", m.GetTotal())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("used", m.GetUsed())
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
func (m *Quota) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDeleted sets the deleted property value. Total space consumed by files in the recycle bin, in bytes. Read-only.
func (m *Quota) SetDeleted(value *int64)() {
    if m != nil {
        m.deleted = value
    }
}
// SetRemaining sets the remaining property value. Total space remaining before reaching the quota limit, in bytes. Read-only.
func (m *Quota) SetRemaining(value *int64)() {
    if m != nil {
        m.remaining = value
    }
}
// SetState sets the state property value. Enumeration value that indicates the state of the storage space. Read-only.
func (m *Quota) SetState(value *string)() {
    if m != nil {
        m.state = value
    }
}
// SetStoragePlanInformation sets the storagePlanInformation property value. Information about the drive's storage quota plans. Only in Personal OneDrive.
func (m *Quota) SetStoragePlanInformation(value StoragePlanInformationable)() {
    if m != nil {
        m.storagePlanInformation = value
    }
}
// SetTotal sets the total property value. Total allowed storage space, in bytes. Read-only.
func (m *Quota) SetTotal(value *int64)() {
    if m != nil {
        m.total = value
    }
}
// SetUsed sets the used property value. Total space used, in bytes. Read-only.
func (m *Quota) SetUsed(value *int64)() {
    if m != nil {
        m.used = value
    }
}
