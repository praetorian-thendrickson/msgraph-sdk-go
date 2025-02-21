package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementExportJob provides operations to manage the deviceManagement singleton.
type DeviceManagementExportJob struct {
    Entity
    // Time that the exported report expires
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Filters applied on the report
    filter *string;
    // Format of the exported report. Possible values are: csv, pdf.
    format *DeviceManagementReportFileFormat;
    // Configures how the requested export job is localized. Possible values are: localizedValuesAsAdditionalColumn, replaceLocalizableValues.
    localizationType *DeviceManagementExportJobLocalizationType;
    // Name of the report
    reportName *string;
    // Time that the exported report was requested
    requestDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Columns selected from the report
    select_escaped []string;
    // A snapshot is an identifiable subset of the dataset represented by the ReportName. A sessionId or CachedReportConfiguration id can be used here. If a sessionId is specified, Filter, Select, and OrderBy are applied to the data represented by the sessionId. Filter, Select, and OrderBy cannot be specified together with a CachedReportConfiguration id.
    snapshotId *string;
    // Status of the export job. Possible values are: unknown, notStarted, inProgress, completed, failed.
    status *DeviceManagementReportStatus;
    // Temporary location of the exported report
    url *string;
}
// NewDeviceManagementExportJob instantiates a new deviceManagementExportJob and sets the default values.
func NewDeviceManagementExportJob()(*DeviceManagementExportJob) {
    m := &DeviceManagementExportJob{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceManagementExportJobFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementExportJobFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewDeviceManagementExportJob(), nil
}
// GetExpirationDateTime gets the expirationDateTime property value. Time that the exported report expires
func (m *DeviceManagementExportJob) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.expirationDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementExportJob) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["expirationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDateTime(val)
        }
        return nil
    }
    res["filter"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFilter(val)
        }
        return nil
    }
    res["format"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementReportFileFormat)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormat(val.(*DeviceManagementReportFileFormat))
        }
        return nil
    }
    res["localizationType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementExportJobLocalizationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalizationType(val.(*DeviceManagementExportJobLocalizationType))
        }
        return nil
    }
    res["reportName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReportName(val)
        }
        return nil
    }
    res["requestDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestDateTime(val)
        }
        return nil
    }
    res["select"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSelect(res)
        }
        return nil
    }
    res["snapshotId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSnapshotId(val)
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementReportStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*DeviceManagementReportStatus))
        }
        return nil
    }
    res["url"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrl(val)
        }
        return nil
    }
    return res
}
// GetFilter gets the filter property value. Filters applied on the report
func (m *DeviceManagementExportJob) GetFilter()(*string) {
    if m == nil {
        return nil
    } else {
        return m.filter
    }
}
// GetFormat gets the format property value. Format of the exported report. Possible values are: csv, pdf.
func (m *DeviceManagementExportJob) GetFormat()(*DeviceManagementReportFileFormat) {
    if m == nil {
        return nil
    } else {
        return m.format
    }
}
// GetLocalizationType gets the localizationType property value. Configures how the requested export job is localized. Possible values are: localizedValuesAsAdditionalColumn, replaceLocalizableValues.
func (m *DeviceManagementExportJob) GetLocalizationType()(*DeviceManagementExportJobLocalizationType) {
    if m == nil {
        return nil
    } else {
        return m.localizationType
    }
}
// GetReportName gets the reportName property value. Name of the report
func (m *DeviceManagementExportJob) GetReportName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportName
    }
}
// GetRequestDateTime gets the requestDateTime property value. Time that the exported report was requested
func (m *DeviceManagementExportJob) GetRequestDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.requestDateTime
    }
}
// GetSelect gets the select property value. Columns selected from the report
func (m *DeviceManagementExportJob) GetSelect()([]string) {
    if m == nil {
        return nil
    } else {
        return m.select_escaped
    }
}
// GetSnapshotId gets the snapshotId property value. A snapshot is an identifiable subset of the dataset represented by the ReportName. A sessionId or CachedReportConfiguration id can be used here. If a sessionId is specified, Filter, Select, and OrderBy are applied to the data represented by the sessionId. Filter, Select, and OrderBy cannot be specified together with a CachedReportConfiguration id.
func (m *DeviceManagementExportJob) GetSnapshotId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.snapshotId
    }
}
// GetStatus gets the status property value. Status of the export job. Possible values are: unknown, notStarted, inProgress, completed, failed.
func (m *DeviceManagementExportJob) GetStatus()(*DeviceManagementReportStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetUrl gets the url property value. Temporary location of the exported report
func (m *DeviceManagementExportJob) GetUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.url
    }
}
func (m *DeviceManagementExportJob) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DeviceManagementExportJob) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("expirationDateTime", m.GetExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("filter", m.GetFilter())
        if err != nil {
            return err
        }
    }
    if m.GetFormat() != nil {
        cast := (*m.GetFormat()).String()
        err = writer.WriteStringValue("format", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetLocalizationType() != nil {
        cast := (*m.GetLocalizationType()).String()
        err = writer.WriteStringValue("localizationType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("reportName", m.GetReportName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("requestDateTime", m.GetRequestDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetSelect() != nil {
        err = writer.WriteCollectionOfStringValues("select", m.GetSelect())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("snapshotId", m.GetSnapshotId())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("url", m.GetUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetExpirationDateTime sets the expirationDateTime property value. Time that the exported report expires
func (m *DeviceManagementExportJob) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.expirationDateTime = value
    }
}
// SetFilter sets the filter property value. Filters applied on the report
func (m *DeviceManagementExportJob) SetFilter(value *string)() {
    if m != nil {
        m.filter = value
    }
}
// SetFormat sets the format property value. Format of the exported report. Possible values are: csv, pdf.
func (m *DeviceManagementExportJob) SetFormat(value *DeviceManagementReportFileFormat)() {
    if m != nil {
        m.format = value
    }
}
// SetLocalizationType sets the localizationType property value. Configures how the requested export job is localized. Possible values are: localizedValuesAsAdditionalColumn, replaceLocalizableValues.
func (m *DeviceManagementExportJob) SetLocalizationType(value *DeviceManagementExportJobLocalizationType)() {
    if m != nil {
        m.localizationType = value
    }
}
// SetReportName sets the reportName property value. Name of the report
func (m *DeviceManagementExportJob) SetReportName(value *string)() {
    if m != nil {
        m.reportName = value
    }
}
// SetRequestDateTime sets the requestDateTime property value. Time that the exported report was requested
func (m *DeviceManagementExportJob) SetRequestDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.requestDateTime = value
    }
}
// SetSelect sets the select property value. Columns selected from the report
func (m *DeviceManagementExportJob) SetSelect(value []string)() {
    if m != nil {
        m.select_escaped = value
    }
}
// SetSnapshotId sets the snapshotId property value. A snapshot is an identifiable subset of the dataset represented by the ReportName. A sessionId or CachedReportConfiguration id can be used here. If a sessionId is specified, Filter, Select, and OrderBy are applied to the data represented by the sessionId. Filter, Select, and OrderBy cannot be specified together with a CachedReportConfiguration id.
func (m *DeviceManagementExportJob) SetSnapshotId(value *string)() {
    if m != nil {
        m.snapshotId = value
    }
}
// SetStatus sets the status property value. Status of the export job. Possible values are: unknown, notStarted, inProgress, completed, failed.
func (m *DeviceManagementExportJob) SetStatus(value *DeviceManagementReportStatus)() {
    if m != nil {
        m.status = value
    }
}
// SetUrl sets the url property value. Temporary location of the exported report
func (m *DeviceManagementExportJob) SetUrl(value *string)() {
    if m != nil {
        m.url = value
    }
}
