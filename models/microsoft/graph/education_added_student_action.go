package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the educationRoot singleton.
type EducationAddedStudentAction int

const (
    NONE_EDUCATIONADDEDSTUDENTACTION EducationAddedStudentAction = iota
    ASSIGNIFOPEN_EDUCATIONADDEDSTUDENTACTION
    UNKNOWNFUTUREVALUE_EDUCATIONADDEDSTUDENTACTION
)

func (i EducationAddedStudentAction) String() string {
    return []string{"NONE", "ASSIGNIFOPEN", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseEducationAddedStudentAction(v string) (interface{}, error) {
    result := NONE_EDUCATIONADDEDSTUDENTACTION
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_EDUCATIONADDEDSTUDENTACTION
        case "ASSIGNIFOPEN":
            result = ASSIGNIFOPEN_EDUCATIONADDEDSTUDENTACTION
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_EDUCATIONADDEDSTUDENTACTION
        default:
            return 0, errors.New("Unknown EducationAddedStudentAction value: " + v)
    }
    return &result, nil
}
func SerializeEducationAddedStudentAction(values []EducationAddedStudentAction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
