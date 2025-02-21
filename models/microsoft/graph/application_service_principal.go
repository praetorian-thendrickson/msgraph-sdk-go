package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ApplicationServicePrincipal provides operations to call the instantiate method.
type ApplicationServicePrincipal struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    application Applicationable;
    // 
    servicePrincipal ServicePrincipalable;
}
// NewApplicationServicePrincipal instantiates a new applicationServicePrincipal and sets the default values.
func NewApplicationServicePrincipal()(*ApplicationServicePrincipal) {
    m := &ApplicationServicePrincipal{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateApplicationServicePrincipalFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateApplicationServicePrincipalFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewApplicationServicePrincipal(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ApplicationServicePrincipal) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetApplication gets the application property value. 
func (m *ApplicationServicePrincipal) GetApplication()(Applicationable) {
    if m == nil {
        return nil
    } else {
        return m.application
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ApplicationServicePrincipal) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["application"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateApplicationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplication(val.(Applicationable))
        }
        return nil
    }
    res["servicePrincipal"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateServicePrincipalFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePrincipal(val.(ServicePrincipalable))
        }
        return nil
    }
    return res
}
// GetServicePrincipal gets the servicePrincipal property value. 
func (m *ApplicationServicePrincipal) GetServicePrincipal()(ServicePrincipalable) {
    if m == nil {
        return nil
    } else {
        return m.servicePrincipal
    }
}
func (m *ApplicationServicePrincipal) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ApplicationServicePrincipal) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("application", m.GetApplication())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("servicePrincipal", m.GetServicePrincipal())
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
func (m *ApplicationServicePrincipal) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetApplication sets the application property value. 
func (m *ApplicationServicePrincipal) SetApplication(value Applicationable)() {
    if m != nil {
        m.application = value
    }
}
// SetServicePrincipal sets the servicePrincipal property value. 
func (m *ApplicationServicePrincipal) SetServicePrincipal(value ServicePrincipalable)() {
    if m != nil {
        m.servicePrincipal = value
    }
}
