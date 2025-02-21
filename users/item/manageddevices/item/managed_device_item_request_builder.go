package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/odataerrors"
    i0fdd6e00c3a70ace39a5146d6c55b15af0de25c040237603b09168231c256f9b "github.com/microsoftgraph/msgraph-sdk-go/users/item/manageddevices/item/deviceconfigurationstates"
    i14175b6b023a6eb71ef6ba830a21cb6cd13dd609e7dc15bd405158e0b3df79ae "github.com/microsoftgraph/msgraph-sdk-go/users/item/manageddevices/item/requestremoteassistance"
    i1754afc016e1fdc501acb882736ad04e30bc5650f53d3c074ab8ac8a51c6fa8f "github.com/microsoftgraph/msgraph-sdk-go/users/item/manageddevices/item/devicecompliancepolicystates"
    i2b3c6b981da171994b1b46a7294676fb6bdc0c93bb24e519dd1306595ead2c80 "github.com/microsoftgraph/msgraph-sdk-go/users/item/manageddevices/item/updatewindowsdeviceaccount"
    i30838d0e3fd5b58dc055f8efb1004348c1be948971f57117d40f2bbf1b4535f0 "github.com/microsoftgraph/msgraph-sdk-go/users/item/manageddevices/item/cleanwindowsdevice"
    i322fafcb7ae8cab29e30bcc25707daf94d942483b0b2026f21e1a15ec8be0f92 "github.com/microsoftgraph/msgraph-sdk-go/users/item/manageddevices/item/shutdown"
    i4b83748504956b892ea1a7f149037fca181d069c7d1673044fada70771e33093 "github.com/microsoftgraph/msgraph-sdk-go/users/item/manageddevices/item/retire"
    i5889a0f91c17b2ff6be4d4d7f398164e86344dbddaf46d10d8dbf893bd1fdcae "github.com/microsoftgraph/msgraph-sdk-go/users/item/manageddevices/item/deleteuserfromsharedappledevice"
    i8cb1d9fa2d4d44aff3eee0f4bd2fa9c55841709e6d7067c757f8da63d9f42acb "github.com/microsoftgraph/msgraph-sdk-go/users/item/manageddevices/item/logoutsharedappledeviceactiveuser"
    i94ba78e1d9aea49972909e0476fb3f1e5a65ad70fb7b6b0bfa5b59fcc655e084 "github.com/microsoftgraph/msgraph-sdk-go/users/item/manageddevices/item/syncdevice"
    i992fc981e39a2a010a5298e4dd852c4b0b9d2285b09665f5884b4fd2dcca552e "github.com/microsoftgraph/msgraph-sdk-go/users/item/manageddevices/item/remotelock"
    i9daab2b0b5282b9d8617504ab740c1b4eeb654e8993c10684911865beb7f6513 "github.com/microsoftgraph/msgraph-sdk-go/users/item/manageddevices/item/devicecategory"
    ia76862a49870b0532e8c4b374ba7384c2818ea7114ca76a78ad21e98b0c92501 "github.com/microsoftgraph/msgraph-sdk-go/users/item/manageddevices/item/locatedevice"
    iad89fd53db2b505c2f511349b0beb9bdda197521ce058a20d54d9f5a00f41230 "github.com/microsoftgraph/msgraph-sdk-go/users/item/manageddevices/item/wipe"
    ic645e4837ba91a567171d12c28ba82eb15a0c5ba22dfc84e46d5174771406adc "github.com/microsoftgraph/msgraph-sdk-go/users/item/manageddevices/item/windowsdefenderscan"
    ic9510153a0140fff3d042d83894479343cdd52e080d68ae1293f3eb456c459e3 "github.com/microsoftgraph/msgraph-sdk-go/users/item/manageddevices/item/resetpasscode"
    icb15920be1a61530593e35306720a47cfee2b31808ee7b11ef7d482703361e75 "github.com/microsoftgraph/msgraph-sdk-go/users/item/manageddevices/item/rebootnow"
    id09c4b3a73bddb54267ba24e9836a171941a344b098f8a53e9044afc145d5939 "github.com/microsoftgraph/msgraph-sdk-go/users/item/manageddevices/item/disablelostmode"
    if07c243615b7cd8f004a1303c7440d0fb98ff8cb9456cb22dee43fbf230ad815 "github.com/microsoftgraph/msgraph-sdk-go/users/item/manageddevices/item/bypassactivationlock"
    if36ea65c1b212f583ceb77d3c50dc0ebe8e096afabd7f9662ae9ee0019217b7a "github.com/microsoftgraph/msgraph-sdk-go/users/item/manageddevices/item/recoverpasscode"
    if54975d74aac049350e7bc12c684f45645594faffab68df3d530a91cd5877663 "github.com/microsoftgraph/msgraph-sdk-go/users/item/manageddevices/item/windowsdefenderupdatesignatures"
    i5b5d439175020246e151a9558d6d568e6bc4bd9fdd376e082964b229d2b71629 "github.com/microsoftgraph/msgraph-sdk-go/users/item/manageddevices/item/devicecompliancepolicystates/item"
    i5e9268af6ea2183cbe612b8d18a9aeec3e11e6ad80e059611e3cafbb64907758 "github.com/microsoftgraph/msgraph-sdk-go/users/item/manageddevices/item/deviceconfigurationstates/item"
)

// ManagedDeviceItemRequestBuilder provides operations to manage the managedDevices property of the microsoft.graph.user entity.
type ManagedDeviceItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ManagedDeviceItemRequestBuilderDeleteOptions options for Delete
type ManagedDeviceItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ManagedDeviceItemRequestBuilderGetOptions options for Get
type ManagedDeviceItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ManagedDeviceItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ManagedDeviceItemRequestBuilderGetQueryParameters the managed devices associated with the user.
type ManagedDeviceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ManagedDeviceItemRequestBuilderPatchOptions options for Patch
type ManagedDeviceItemRequestBuilderPatchOptions struct {
    // 
    Body i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ManagedDeviceable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ManagedDeviceItemRequestBuilder) BypassActivationLock()(*if07c243615b7cd8f004a1303c7440d0fb98ff8cb9456cb22dee43fbf230ad815.BypassActivationLockRequestBuilder) {
    return if07c243615b7cd8f004a1303c7440d0fb98ff8cb9456cb22dee43fbf230ad815.NewBypassActivationLockRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceItemRequestBuilder) CleanWindowsDevice()(*i30838d0e3fd5b58dc055f8efb1004348c1be948971f57117d40f2bbf1b4535f0.CleanWindowsDeviceRequestBuilder) {
    return i30838d0e3fd5b58dc055f8efb1004348c1be948971f57117d40f2bbf1b4535f0.NewCleanWindowsDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewManagedDeviceItemRequestBuilderInternal instantiates a new ManagedDeviceItemRequestBuilder and sets the default values.
func NewManagedDeviceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedDeviceItemRequestBuilder) {
    m := &ManagedDeviceItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/managedDevices/{managedDevice_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewManagedDeviceItemRequestBuilder instantiates a new ManagedDeviceItemRequestBuilder and sets the default values.
func NewManagedDeviceItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedDeviceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedDeviceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property managedDevices for users
func (m *ManagedDeviceItemRequestBuilder) CreateDeleteRequestInformation(options *ManagedDeviceItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
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
// CreateGetRequestInformation the managed devices associated with the user.
func (m *ManagedDeviceItemRequestBuilder) CreateGetRequestInformation(options *ManagedDeviceItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property managedDevices in users
func (m *ManagedDeviceItemRequestBuilder) CreatePatchRequestInformation(options *ManagedDeviceItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property managedDevices for users
func (m *ManagedDeviceItemRequestBuilder) Delete(options *ManagedDeviceItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
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
func (m *ManagedDeviceItemRequestBuilder) DeleteUserFromSharedAppleDevice()(*i5889a0f91c17b2ff6be4d4d7f398164e86344dbddaf46d10d8dbf893bd1fdcae.DeleteUserFromSharedAppleDeviceRequestBuilder) {
    return i5889a0f91c17b2ff6be4d4d7f398164e86344dbddaf46d10d8dbf893bd1fdcae.NewDeleteUserFromSharedAppleDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceItemRequestBuilder) DeviceCategory()(*i9daab2b0b5282b9d8617504ab740c1b4eeb654e8993c10684911865beb7f6513.DeviceCategoryRequestBuilder) {
    return i9daab2b0b5282b9d8617504ab740c1b4eeb654e8993c10684911865beb7f6513.NewDeviceCategoryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceItemRequestBuilder) DeviceCompliancePolicyStates()(*i1754afc016e1fdc501acb882736ad04e30bc5650f53d3c074ab8ac8a51c6fa8f.DeviceCompliancePolicyStatesRequestBuilder) {
    return i1754afc016e1fdc501acb882736ad04e30bc5650f53d3c074ab8ac8a51c6fa8f.NewDeviceCompliancePolicyStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceCompliancePolicyStatesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.managedDevices.item.deviceCompliancePolicyStates.item collection
func (m *ManagedDeviceItemRequestBuilder) DeviceCompliancePolicyStatesById(id string)(*i5b5d439175020246e151a9558d6d568e6bc4bd9fdd376e082964b229d2b71629.DeviceCompliancePolicyStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceCompliancePolicyState_id"] = id
    }
    return i5b5d439175020246e151a9558d6d568e6bc4bd9fdd376e082964b229d2b71629.NewDeviceCompliancePolicyStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedDeviceItemRequestBuilder) DeviceConfigurationStates()(*i0fdd6e00c3a70ace39a5146d6c55b15af0de25c040237603b09168231c256f9b.DeviceConfigurationStatesRequestBuilder) {
    return i0fdd6e00c3a70ace39a5146d6c55b15af0de25c040237603b09168231c256f9b.NewDeviceConfigurationStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceConfigurationStatesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.managedDevices.item.deviceConfigurationStates.item collection
func (m *ManagedDeviceItemRequestBuilder) DeviceConfigurationStatesById(id string)(*i5e9268af6ea2183cbe612b8d18a9aeec3e11e6ad80e059611e3cafbb64907758.DeviceConfigurationStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceConfigurationState_id"] = id
    }
    return i5e9268af6ea2183cbe612b8d18a9aeec3e11e6ad80e059611e3cafbb64907758.NewDeviceConfigurationStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedDeviceItemRequestBuilder) DisableLostMode()(*id09c4b3a73bddb54267ba24e9836a171941a344b098f8a53e9044afc145d5939.DisableLostModeRequestBuilder) {
    return id09c4b3a73bddb54267ba24e9836a171941a344b098f8a53e9044afc145d5939.NewDisableLostModeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the managed devices associated with the user.
func (m *ManagedDeviceItemRequestBuilder) Get(options *ManagedDeviceItemRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ManagedDeviceable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateManagedDeviceFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ManagedDeviceable), nil
}
func (m *ManagedDeviceItemRequestBuilder) LocateDevice()(*ia76862a49870b0532e8c4b374ba7384c2818ea7114ca76a78ad21e98b0c92501.LocateDeviceRequestBuilder) {
    return ia76862a49870b0532e8c4b374ba7384c2818ea7114ca76a78ad21e98b0c92501.NewLocateDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceItemRequestBuilder) LogoutSharedAppleDeviceActiveUser()(*i8cb1d9fa2d4d44aff3eee0f4bd2fa9c55841709e6d7067c757f8da63d9f42acb.LogoutSharedAppleDeviceActiveUserRequestBuilder) {
    return i8cb1d9fa2d4d44aff3eee0f4bd2fa9c55841709e6d7067c757f8da63d9f42acb.NewLogoutSharedAppleDeviceActiveUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property managedDevices in users
func (m *ManagedDeviceItemRequestBuilder) Patch(options *ManagedDeviceItemRequestBuilderPatchOptions)(error) {
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
func (m *ManagedDeviceItemRequestBuilder) RebootNow()(*icb15920be1a61530593e35306720a47cfee2b31808ee7b11ef7d482703361e75.RebootNowRequestBuilder) {
    return icb15920be1a61530593e35306720a47cfee2b31808ee7b11ef7d482703361e75.NewRebootNowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceItemRequestBuilder) RecoverPasscode()(*if36ea65c1b212f583ceb77d3c50dc0ebe8e096afabd7f9662ae9ee0019217b7a.RecoverPasscodeRequestBuilder) {
    return if36ea65c1b212f583ceb77d3c50dc0ebe8e096afabd7f9662ae9ee0019217b7a.NewRecoverPasscodeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceItemRequestBuilder) RemoteLock()(*i992fc981e39a2a010a5298e4dd852c4b0b9d2285b09665f5884b4fd2dcca552e.RemoteLockRequestBuilder) {
    return i992fc981e39a2a010a5298e4dd852c4b0b9d2285b09665f5884b4fd2dcca552e.NewRemoteLockRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceItemRequestBuilder) RequestRemoteAssistance()(*i14175b6b023a6eb71ef6ba830a21cb6cd13dd609e7dc15bd405158e0b3df79ae.RequestRemoteAssistanceRequestBuilder) {
    return i14175b6b023a6eb71ef6ba830a21cb6cd13dd609e7dc15bd405158e0b3df79ae.NewRequestRemoteAssistanceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceItemRequestBuilder) ResetPasscode()(*ic9510153a0140fff3d042d83894479343cdd52e080d68ae1293f3eb456c459e3.ResetPasscodeRequestBuilder) {
    return ic9510153a0140fff3d042d83894479343cdd52e080d68ae1293f3eb456c459e3.NewResetPasscodeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceItemRequestBuilder) Retire()(*i4b83748504956b892ea1a7f149037fca181d069c7d1673044fada70771e33093.RetireRequestBuilder) {
    return i4b83748504956b892ea1a7f149037fca181d069c7d1673044fada70771e33093.NewRetireRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceItemRequestBuilder) ShutDown()(*i322fafcb7ae8cab29e30bcc25707daf94d942483b0b2026f21e1a15ec8be0f92.ShutDownRequestBuilder) {
    return i322fafcb7ae8cab29e30bcc25707daf94d942483b0b2026f21e1a15ec8be0f92.NewShutDownRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceItemRequestBuilder) SyncDevice()(*i94ba78e1d9aea49972909e0476fb3f1e5a65ad70fb7b6b0bfa5b59fcc655e084.SyncDeviceRequestBuilder) {
    return i94ba78e1d9aea49972909e0476fb3f1e5a65ad70fb7b6b0bfa5b59fcc655e084.NewSyncDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceItemRequestBuilder) UpdateWindowsDeviceAccount()(*i2b3c6b981da171994b1b46a7294676fb6bdc0c93bb24e519dd1306595ead2c80.UpdateWindowsDeviceAccountRequestBuilder) {
    return i2b3c6b981da171994b1b46a7294676fb6bdc0c93bb24e519dd1306595ead2c80.NewUpdateWindowsDeviceAccountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceItemRequestBuilder) WindowsDefenderScan()(*ic645e4837ba91a567171d12c28ba82eb15a0c5ba22dfc84e46d5174771406adc.WindowsDefenderScanRequestBuilder) {
    return ic645e4837ba91a567171d12c28ba82eb15a0c5ba22dfc84e46d5174771406adc.NewWindowsDefenderScanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceItemRequestBuilder) WindowsDefenderUpdateSignatures()(*if54975d74aac049350e7bc12c684f45645594faffab68df3d530a91cd5877663.WindowsDefenderUpdateSignaturesRequestBuilder) {
    return if54975d74aac049350e7bc12c684f45645594faffab68df3d530a91cd5877663.NewWindowsDefenderUpdateSignaturesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceItemRequestBuilder) Wipe()(*iad89fd53db2b505c2f511349b0beb9bdda197521ce058a20d54d9f5a00f41230.WipeRequestBuilder) {
    return iad89fd53db2b505c2f511349b0beb9bdda197521ce058a20d54d9f5a00f41230.NewWipeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
