package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the drive singleton.
type ManagementAgentType int

const (
    EAS_MANAGEMENTAGENTTYPE ManagementAgentType = iota
    MDM_MANAGEMENTAGENTTYPE
    EASMDM_MANAGEMENTAGENTTYPE
    INTUNECLIENT_MANAGEMENTAGENTTYPE
    EASINTUNECLIENT_MANAGEMENTAGENTTYPE
    CONFIGURATIONMANAGERCLIENT_MANAGEMENTAGENTTYPE
    CONFIGURATIONMANAGERCLIENTMDM_MANAGEMENTAGENTTYPE
    CONFIGURATIONMANAGERCLIENTMDMEAS_MANAGEMENTAGENTTYPE
    UNKNOWN_MANAGEMENTAGENTTYPE
    JAMF_MANAGEMENTAGENTTYPE
    GOOGLECLOUDDEVICEPOLICYCONTROLLER_MANAGEMENTAGENTTYPE
    MICROSOFT365MANAGEDMDM_MANAGEMENTAGENTTYPE
    MSSENSE_MANAGEMENTAGENTTYPE
)

func (i ManagementAgentType) String() string {
    return []string{"EAS", "MDM", "EASMDM", "INTUNECLIENT", "EASINTUNECLIENT", "CONFIGURATIONMANAGERCLIENT", "CONFIGURATIONMANAGERCLIENTMDM", "CONFIGURATIONMANAGERCLIENTMDMEAS", "UNKNOWN", "JAMF", "GOOGLECLOUDDEVICEPOLICYCONTROLLER", "MICROSOFT365MANAGEDMDM", "MSSENSE"}[i]
}
func ParseManagementAgentType(v string) (interface{}, error) {
    result := EAS_MANAGEMENTAGENTTYPE
    switch strings.ToUpper(v) {
        case "EAS":
            result = EAS_MANAGEMENTAGENTTYPE
        case "MDM":
            result = MDM_MANAGEMENTAGENTTYPE
        case "EASMDM":
            result = EASMDM_MANAGEMENTAGENTTYPE
        case "INTUNECLIENT":
            result = INTUNECLIENT_MANAGEMENTAGENTTYPE
        case "EASINTUNECLIENT":
            result = EASINTUNECLIENT_MANAGEMENTAGENTTYPE
        case "CONFIGURATIONMANAGERCLIENT":
            result = CONFIGURATIONMANAGERCLIENT_MANAGEMENTAGENTTYPE
        case "CONFIGURATIONMANAGERCLIENTMDM":
            result = CONFIGURATIONMANAGERCLIENTMDM_MANAGEMENTAGENTTYPE
        case "CONFIGURATIONMANAGERCLIENTMDMEAS":
            result = CONFIGURATIONMANAGERCLIENTMDMEAS_MANAGEMENTAGENTTYPE
        case "UNKNOWN":
            result = UNKNOWN_MANAGEMENTAGENTTYPE
        case "JAMF":
            result = JAMF_MANAGEMENTAGENTTYPE
        case "GOOGLECLOUDDEVICEPOLICYCONTROLLER":
            result = GOOGLECLOUDDEVICEPOLICYCONTROLLER_MANAGEMENTAGENTTYPE
        case "MICROSOFT365MANAGEDMDM":
            result = MICROSOFT365MANAGEDMDM_MANAGEMENTAGENTTYPE
        case "MSSENSE":
            result = MSSENSE_MANAGEMENTAGENTTYPE
        default:
            return 0, errors.New("Unknown ManagementAgentType value: " + v)
    }
    return &result, nil
}
func SerializeManagementAgentType(values []ManagementAgentType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
