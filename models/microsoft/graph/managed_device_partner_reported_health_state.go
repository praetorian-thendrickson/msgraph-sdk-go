package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the drive singleton.
type ManagedDevicePartnerReportedHealthState int

const (
    UNKNOWN_MANAGEDDEVICEPARTNERREPORTEDHEALTHSTATE ManagedDevicePartnerReportedHealthState = iota
    ACTIVATED_MANAGEDDEVICEPARTNERREPORTEDHEALTHSTATE
    DEACTIVATED_MANAGEDDEVICEPARTNERREPORTEDHEALTHSTATE
    SECURED_MANAGEDDEVICEPARTNERREPORTEDHEALTHSTATE
    LOWSEVERITY_MANAGEDDEVICEPARTNERREPORTEDHEALTHSTATE
    MEDIUMSEVERITY_MANAGEDDEVICEPARTNERREPORTEDHEALTHSTATE
    HIGHSEVERITY_MANAGEDDEVICEPARTNERREPORTEDHEALTHSTATE
    UNRESPONSIVE_MANAGEDDEVICEPARTNERREPORTEDHEALTHSTATE
    COMPROMISED_MANAGEDDEVICEPARTNERREPORTEDHEALTHSTATE
    MISCONFIGURED_MANAGEDDEVICEPARTNERREPORTEDHEALTHSTATE
)

func (i ManagedDevicePartnerReportedHealthState) String() string {
    return []string{"UNKNOWN", "ACTIVATED", "DEACTIVATED", "SECURED", "LOWSEVERITY", "MEDIUMSEVERITY", "HIGHSEVERITY", "UNRESPONSIVE", "COMPROMISED", "MISCONFIGURED"}[i]
}
func ParseManagedDevicePartnerReportedHealthState(v string) (interface{}, error) {
    result := UNKNOWN_MANAGEDDEVICEPARTNERREPORTEDHEALTHSTATE
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_MANAGEDDEVICEPARTNERREPORTEDHEALTHSTATE
        case "ACTIVATED":
            result = ACTIVATED_MANAGEDDEVICEPARTNERREPORTEDHEALTHSTATE
        case "DEACTIVATED":
            result = DEACTIVATED_MANAGEDDEVICEPARTNERREPORTEDHEALTHSTATE
        case "SECURED":
            result = SECURED_MANAGEDDEVICEPARTNERREPORTEDHEALTHSTATE
        case "LOWSEVERITY":
            result = LOWSEVERITY_MANAGEDDEVICEPARTNERREPORTEDHEALTHSTATE
        case "MEDIUMSEVERITY":
            result = MEDIUMSEVERITY_MANAGEDDEVICEPARTNERREPORTEDHEALTHSTATE
        case "HIGHSEVERITY":
            result = HIGHSEVERITY_MANAGEDDEVICEPARTNERREPORTEDHEALTHSTATE
        case "UNRESPONSIVE":
            result = UNRESPONSIVE_MANAGEDDEVICEPARTNERREPORTEDHEALTHSTATE
        case "COMPROMISED":
            result = COMPROMISED_MANAGEDDEVICEPARTNERREPORTEDHEALTHSTATE
        case "MISCONFIGURED":
            result = MISCONFIGURED_MANAGEDDEVICEPARTNERREPORTEDHEALTHSTATE
        default:
            return 0, errors.New("Unknown ManagedDevicePartnerReportedHealthState value: " + v)
    }
    return &result, nil
}
func SerializeManagedDevicePartnerReportedHealthState(values []ManagedDevicePartnerReportedHealthState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
