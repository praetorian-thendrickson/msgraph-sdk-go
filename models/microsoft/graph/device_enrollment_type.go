package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the drive singleton.
type DeviceEnrollmentType int

const (
    UNKNOWN_DEVICEENROLLMENTTYPE DeviceEnrollmentType = iota
    USERENROLLMENT_DEVICEENROLLMENTTYPE
    DEVICEENROLLMENTMANAGER_DEVICEENROLLMENTTYPE
    APPLEBULKWITHUSER_DEVICEENROLLMENTTYPE
    APPLEBULKWITHOUTUSER_DEVICEENROLLMENTTYPE
    WINDOWSAZUREADJOIN_DEVICEENROLLMENTTYPE
    WINDOWSBULKUSERLESS_DEVICEENROLLMENTTYPE
    WINDOWSAUTOENROLLMENT_DEVICEENROLLMENTTYPE
    WINDOWSBULKAZUREDOMAINJOIN_DEVICEENROLLMENTTYPE
    WINDOWSCOMANAGEMENT_DEVICEENROLLMENTTYPE
    WINDOWSAZUREADJOINUSINGDEVICEAUTH_DEVICEENROLLMENTTYPE
    APPLEUSERENROLLMENT_DEVICEENROLLMENTTYPE
    APPLEUSERENROLLMENTWITHSERVICEACCOUNT_DEVICEENROLLMENTTYPE
)

func (i DeviceEnrollmentType) String() string {
    return []string{"UNKNOWN", "USERENROLLMENT", "DEVICEENROLLMENTMANAGER", "APPLEBULKWITHUSER", "APPLEBULKWITHOUTUSER", "WINDOWSAZUREADJOIN", "WINDOWSBULKUSERLESS", "WINDOWSAUTOENROLLMENT", "WINDOWSBULKAZUREDOMAINJOIN", "WINDOWSCOMANAGEMENT", "WINDOWSAZUREADJOINUSINGDEVICEAUTH", "APPLEUSERENROLLMENT", "APPLEUSERENROLLMENTWITHSERVICEACCOUNT"}[i]
}
func ParseDeviceEnrollmentType(v string) (interface{}, error) {
    result := UNKNOWN_DEVICEENROLLMENTTYPE
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_DEVICEENROLLMENTTYPE
        case "USERENROLLMENT":
            result = USERENROLLMENT_DEVICEENROLLMENTTYPE
        case "DEVICEENROLLMENTMANAGER":
            result = DEVICEENROLLMENTMANAGER_DEVICEENROLLMENTTYPE
        case "APPLEBULKWITHUSER":
            result = APPLEBULKWITHUSER_DEVICEENROLLMENTTYPE
        case "APPLEBULKWITHOUTUSER":
            result = APPLEBULKWITHOUTUSER_DEVICEENROLLMENTTYPE
        case "WINDOWSAZUREADJOIN":
            result = WINDOWSAZUREADJOIN_DEVICEENROLLMENTTYPE
        case "WINDOWSBULKUSERLESS":
            result = WINDOWSBULKUSERLESS_DEVICEENROLLMENTTYPE
        case "WINDOWSAUTOENROLLMENT":
            result = WINDOWSAUTOENROLLMENT_DEVICEENROLLMENTTYPE
        case "WINDOWSBULKAZUREDOMAINJOIN":
            result = WINDOWSBULKAZUREDOMAINJOIN_DEVICEENROLLMENTTYPE
        case "WINDOWSCOMANAGEMENT":
            result = WINDOWSCOMANAGEMENT_DEVICEENROLLMENTTYPE
        case "WINDOWSAZUREADJOINUSINGDEVICEAUTH":
            result = WINDOWSAZUREADJOINUSINGDEVICEAUTH_DEVICEENROLLMENTTYPE
        case "APPLEUSERENROLLMENT":
            result = APPLEUSERENROLLMENT_DEVICEENROLLMENTTYPE
        case "APPLEUSERENROLLMENTWITHSERVICEACCOUNT":
            result = APPLEUSERENROLLMENTWITHSERVICEACCOUNT_DEVICEENROLLMENTTYPE
        default:
            return 0, errors.New("Unknown DeviceEnrollmentType value: " + v)
    }
    return &result, nil
}
func SerializeDeviceEnrollmentType(values []DeviceEnrollmentType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
