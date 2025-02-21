package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CertificateAuthorityable 
type CertificateAuthorityable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCertificate()([]byte)
    GetCertificateRevocationListUrl()(*string)
    GetDeltaCertificateRevocationListUrl()(*string)
    GetIsRootAuthority()(*bool)
    GetIssuer()(*string)
    GetIssuerSki()(*string)
    SetCertificate(value []byte)()
    SetCertificateRevocationListUrl(value *string)()
    SetDeltaCertificateRevocationListUrl(value *string)()
    SetIsRootAuthority(value *bool)()
    SetIssuer(value *string)()
    SetIssuerSki(value *string)()
}
