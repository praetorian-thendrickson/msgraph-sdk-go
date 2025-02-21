package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the identityContainer singleton.
type PersistentBrowserSessionMode int

const (
    ALWAYS_PERSISTENTBROWSERSESSIONMODE PersistentBrowserSessionMode = iota
    NEVER_PERSISTENTBROWSERSESSIONMODE
)

func (i PersistentBrowserSessionMode) String() string {
    return []string{"ALWAYS", "NEVER"}[i]
}
func ParsePersistentBrowserSessionMode(v string) (interface{}, error) {
    result := ALWAYS_PERSISTENTBROWSERSESSIONMODE
    switch strings.ToUpper(v) {
        case "ALWAYS":
            result = ALWAYS_PERSISTENTBROWSERSESSIONMODE
        case "NEVER":
            result = NEVER_PERSISTENTBROWSERSESSIONMODE
        default:
            return 0, errors.New("Unknown PersistentBrowserSessionMode value: " + v)
    }
    return &result, nil
}
func SerializePersistentBrowserSessionMode(values []PersistentBrowserSessionMode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
