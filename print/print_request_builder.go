package print

import (
    i0bda270a8db970823b8449367f0a1b6167aeb67bb1a0b116061e10c72e9e8d5d "github.com/microsoftgraph/msgraph-sdk-go/print/printers"
    i0f704d64ae3c276c8ff23cafb44315ef4067a668629360d124bdabfea38afdd3 "github.com/microsoftgraph/msgraph-sdk-go/print/connectors"
    i54feb3a1bee5e2d3891bf4b38d7ad46df95f464ec5502fc27ce12645532dcf04 "github.com/microsoftgraph/msgraph-sdk-go/print/shares"
    i802e345f3915c26a53a0e538ad96acb01b0ad515a9510581f00bcedb2e2d1266 "github.com/microsoftgraph/msgraph-sdk-go/print/services"
    ib24f76d25d0e8f20d062be56f8b94808153dc20498d344b76af7867075ea8ad9 "github.com/microsoftgraph/msgraph-sdk-go/print/operations"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    ied8956e290d0c898cb2abe2b85d9373567ead7ae2e05a776597718dd3f1709e8 "github.com/microsoftgraph/msgraph-sdk-go/print/taskdefinitions"
    i0b3afc35ae08ae48e886b1ea1ae2fe47da46f82bd161cb48f9ff550676f78301 "github.com/microsoftgraph/msgraph-sdk-go/print/services/item"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i7b3d94bc34459d3d2fa55e5dc4096ab7096e743ffc216c2d05b8f451f231fde1 "github.com/microsoftgraph/msgraph-sdk-go/print/printers/item"
    i827522a32e9dfc8181a7f7793aec64fdf0814658430464c0b64af19c8e4e9a96 "github.com/microsoftgraph/msgraph-sdk-go/print/taskdefinitions/item"
    i8e18f5c596955d2371ff489e959f3488755280f5b7014ba5a0f5b95b7dc44bf2 "github.com/microsoftgraph/msgraph-sdk-go/print/shares/item"
    id0e61828b4320e56c027152ab07b04414264a89aa169cd832394397ada28defd "github.com/microsoftgraph/msgraph-sdk-go/print/operations/item"
    iebd42d165d4ff6aca1eec85e8351f67b97a4208ea00a5bc8ac67f4f1f191b677 "github.com/microsoftgraph/msgraph-sdk-go/print/connectors/item"
    i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/odataerrors"
)

// PrintRequestBuilder provides operations to manage the print singleton.
type PrintRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// PrintRequestBuilderGetOptions options for Get
type PrintRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *PrintRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// PrintRequestBuilderGetQueryParameters get print
type PrintRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// PrintRequestBuilderPatchOptions options for Patch
type PrintRequestBuilderPatchOptions struct {
    // 
    Body i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Printable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *PrintRequestBuilder) Connectors()(*i0f704d64ae3c276c8ff23cafb44315ef4067a668629360d124bdabfea38afdd3.ConnectorsRequestBuilder) {
    return i0f704d64ae3c276c8ff23cafb44315ef4067a668629360d124bdabfea38afdd3.NewConnectorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ConnectorsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.print.connectors.item collection
func (m *PrintRequestBuilder) ConnectorsById(id string)(*iebd42d165d4ff6aca1eec85e8351f67b97a4208ea00a5bc8ac67f4f1f191b677.PrintConnectorItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printConnector_id"] = id
    }
    return iebd42d165d4ff6aca1eec85e8351f67b97a4208ea00a5bc8ac67f4f1f191b677.NewPrintConnectorItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewPrintRequestBuilderInternal instantiates a new PrintRequestBuilder and sets the default values.
func NewPrintRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrintRequestBuilder) {
    m := &PrintRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/print{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPrintRequestBuilder instantiates a new PrintRequestBuilder and sets the default values.
func NewPrintRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrintRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrintRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get print
func (m *PrintRequestBuilder) CreateGetRequestInformation(options *PrintRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update print
func (m *PrintRequestBuilder) CreatePatchRequestInformation(options *PrintRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Get get print
func (m *PrintRequestBuilder) Get(options *PrintRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Printable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreatePrintFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Printable), nil
}
func (m *PrintRequestBuilder) Operations()(*ib24f76d25d0e8f20d062be56f8b94808153dc20498d344b76af7867075ea8ad9.OperationsRequestBuilder) {
    return ib24f76d25d0e8f20d062be56f8b94808153dc20498d344b76af7867075ea8ad9.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.print.operations.item collection
func (m *PrintRequestBuilder) OperationsById(id string)(*id0e61828b4320e56c027152ab07b04414264a89aa169cd832394397ada28defd.PrintOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printOperation_id"] = id
    }
    return id0e61828b4320e56c027152ab07b04414264a89aa169cd832394397ada28defd.NewPrintOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update print
func (m *PrintRequestBuilder) Patch(options *PrintRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *PrintRequestBuilder) Printers()(*i0bda270a8db970823b8449367f0a1b6167aeb67bb1a0b116061e10c72e9e8d5d.PrintersRequestBuilder) {
    return i0bda270a8db970823b8449367f0a1b6167aeb67bb1a0b116061e10c72e9e8d5d.NewPrintersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PrintersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.print.printers.item collection
func (m *PrintRequestBuilder) PrintersById(id string)(*i7b3d94bc34459d3d2fa55e5dc4096ab7096e743ffc216c2d05b8f451f231fde1.PrinterItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printer_id"] = id
    }
    return i7b3d94bc34459d3d2fa55e5dc4096ab7096e743ffc216c2d05b8f451f231fde1.NewPrinterItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PrintRequestBuilder) Services()(*i802e345f3915c26a53a0e538ad96acb01b0ad515a9510581f00bcedb2e2d1266.ServicesRequestBuilder) {
    return i802e345f3915c26a53a0e538ad96acb01b0ad515a9510581f00bcedb2e2d1266.NewServicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.print.services.item collection
func (m *PrintRequestBuilder) ServicesById(id string)(*i0b3afc35ae08ae48e886b1ea1ae2fe47da46f82bd161cb48f9ff550676f78301.PrintServiceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printService_id"] = id
    }
    return i0b3afc35ae08ae48e886b1ea1ae2fe47da46f82bd161cb48f9ff550676f78301.NewPrintServiceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PrintRequestBuilder) Shares()(*i54feb3a1bee5e2d3891bf4b38d7ad46df95f464ec5502fc27ce12645532dcf04.SharesRequestBuilder) {
    return i54feb3a1bee5e2d3891bf4b38d7ad46df95f464ec5502fc27ce12645532dcf04.NewSharesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SharesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.print.shares.item collection
func (m *PrintRequestBuilder) SharesById(id string)(*i8e18f5c596955d2371ff489e959f3488755280f5b7014ba5a0f5b95b7dc44bf2.PrinterShareItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printerShare_id"] = id
    }
    return i8e18f5c596955d2371ff489e959f3488755280f5b7014ba5a0f5b95b7dc44bf2.NewPrinterShareItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PrintRequestBuilder) TaskDefinitions()(*ied8956e290d0c898cb2abe2b85d9373567ead7ae2e05a776597718dd3f1709e8.TaskDefinitionsRequestBuilder) {
    return ied8956e290d0c898cb2abe2b85d9373567ead7ae2e05a776597718dd3f1709e8.NewTaskDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TaskDefinitionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.print.taskDefinitions.item collection
func (m *PrintRequestBuilder) TaskDefinitionsById(id string)(*i827522a32e9dfc8181a7f7793aec64fdf0814658430464c0b64af19c8e4e9a96.PrintTaskDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printTaskDefinition_id"] = id
    }
    return i827522a32e9dfc8181a7f7793aec64fdf0814658430464c0b64af19c8e4e9a96.NewPrintTaskDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
