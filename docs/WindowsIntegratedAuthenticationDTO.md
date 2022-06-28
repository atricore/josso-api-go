# WindowsIntegratedAuthenticationDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DelegatedAuthentications** | Pointer to [**[]DelegatedAuthenticationDTO**](DelegatedAuthenticationDTO.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**DomainController** | Pointer to **string** |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**KeyTab** | Pointer to [**ResourceDTO**](ResourceDTO.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OverwriteKerberosSetup** | Pointer to **bool** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**ServiceClass** | Pointer to **string** |  | [optional] 
**ServiceName** | Pointer to **string** |  | [optional] 
**X** | Pointer to **float64** |  | [optional] 
**Y** | Pointer to **float64** |  | [optional] 

## Methods

### NewWindowsIntegratedAuthenticationDTO

`func NewWindowsIntegratedAuthenticationDTO() *WindowsIntegratedAuthenticationDTO`

NewWindowsIntegratedAuthenticationDTO instantiates a new WindowsIntegratedAuthenticationDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWindowsIntegratedAuthenticationDTOWithDefaults

`func NewWindowsIntegratedAuthenticationDTOWithDefaults() *WindowsIntegratedAuthenticationDTO`

NewWindowsIntegratedAuthenticationDTOWithDefaults instantiates a new WindowsIntegratedAuthenticationDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelegatedAuthentications

`func (o *WindowsIntegratedAuthenticationDTO) GetDelegatedAuthentications() []DelegatedAuthenticationDTO`

GetDelegatedAuthentications returns the DelegatedAuthentications field if non-nil, zero value otherwise.

### GetDelegatedAuthenticationsOk

`func (o *WindowsIntegratedAuthenticationDTO) GetDelegatedAuthenticationsOk() (*[]DelegatedAuthenticationDTO, bool)`

GetDelegatedAuthenticationsOk returns a tuple with the DelegatedAuthentications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatedAuthentications

`func (o *WindowsIntegratedAuthenticationDTO) SetDelegatedAuthentications(v []DelegatedAuthenticationDTO)`

SetDelegatedAuthentications sets DelegatedAuthentications field to given value.

### HasDelegatedAuthentications

`func (o *WindowsIntegratedAuthenticationDTO) HasDelegatedAuthentications() bool`

HasDelegatedAuthentications returns a boolean if a field has been set.

### GetDescription

`func (o *WindowsIntegratedAuthenticationDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WindowsIntegratedAuthenticationDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WindowsIntegratedAuthenticationDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WindowsIntegratedAuthenticationDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *WindowsIntegratedAuthenticationDTO) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *WindowsIntegratedAuthenticationDTO) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *WindowsIntegratedAuthenticationDTO) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *WindowsIntegratedAuthenticationDTO) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDomain

`func (o *WindowsIntegratedAuthenticationDTO) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *WindowsIntegratedAuthenticationDTO) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *WindowsIntegratedAuthenticationDTO) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *WindowsIntegratedAuthenticationDTO) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetDomainController

`func (o *WindowsIntegratedAuthenticationDTO) GetDomainController() string`

GetDomainController returns the DomainController field if non-nil, zero value otherwise.

### GetDomainControllerOk

`func (o *WindowsIntegratedAuthenticationDTO) GetDomainControllerOk() (*string, bool)`

GetDomainControllerOk returns a tuple with the DomainController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainController

`func (o *WindowsIntegratedAuthenticationDTO) SetDomainController(v string)`

SetDomainController sets DomainController field to given value.

### HasDomainController

`func (o *WindowsIntegratedAuthenticationDTO) HasDomainController() bool`

HasDomainController returns a boolean if a field has been set.

### GetElementId

`func (o *WindowsIntegratedAuthenticationDTO) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *WindowsIntegratedAuthenticationDTO) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *WindowsIntegratedAuthenticationDTO) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *WindowsIntegratedAuthenticationDTO) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetHost

`func (o *WindowsIntegratedAuthenticationDTO) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *WindowsIntegratedAuthenticationDTO) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *WindowsIntegratedAuthenticationDTO) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *WindowsIntegratedAuthenticationDTO) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *WindowsIntegratedAuthenticationDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WindowsIntegratedAuthenticationDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WindowsIntegratedAuthenticationDTO) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *WindowsIntegratedAuthenticationDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKeyTab

`func (o *WindowsIntegratedAuthenticationDTO) GetKeyTab() ResourceDTO`

GetKeyTab returns the KeyTab field if non-nil, zero value otherwise.

### GetKeyTabOk

`func (o *WindowsIntegratedAuthenticationDTO) GetKeyTabOk() (*ResourceDTO, bool)`

GetKeyTabOk returns a tuple with the KeyTab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyTab

`func (o *WindowsIntegratedAuthenticationDTO) SetKeyTab(v ResourceDTO)`

SetKeyTab sets KeyTab field to given value.

### HasKeyTab

`func (o *WindowsIntegratedAuthenticationDTO) HasKeyTab() bool`

HasKeyTab returns a boolean if a field has been set.

### GetName

`func (o *WindowsIntegratedAuthenticationDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WindowsIntegratedAuthenticationDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WindowsIntegratedAuthenticationDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WindowsIntegratedAuthenticationDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOverwriteKerberosSetup

`func (o *WindowsIntegratedAuthenticationDTO) GetOverwriteKerberosSetup() bool`

GetOverwriteKerberosSetup returns the OverwriteKerberosSetup field if non-nil, zero value otherwise.

### GetOverwriteKerberosSetupOk

`func (o *WindowsIntegratedAuthenticationDTO) GetOverwriteKerberosSetupOk() (*bool, bool)`

GetOverwriteKerberosSetupOk returns a tuple with the OverwriteKerberosSetup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteKerberosSetup

`func (o *WindowsIntegratedAuthenticationDTO) SetOverwriteKerberosSetup(v bool)`

SetOverwriteKerberosSetup sets OverwriteKerberosSetup field to given value.

### HasOverwriteKerberosSetup

`func (o *WindowsIntegratedAuthenticationDTO) HasOverwriteKerberosSetup() bool`

HasOverwriteKerberosSetup returns a boolean if a field has been set.

### GetPort

`func (o *WindowsIntegratedAuthenticationDTO) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *WindowsIntegratedAuthenticationDTO) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *WindowsIntegratedAuthenticationDTO) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *WindowsIntegratedAuthenticationDTO) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProtocol

`func (o *WindowsIntegratedAuthenticationDTO) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *WindowsIntegratedAuthenticationDTO) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *WindowsIntegratedAuthenticationDTO) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *WindowsIntegratedAuthenticationDTO) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetServiceClass

`func (o *WindowsIntegratedAuthenticationDTO) GetServiceClass() string`

GetServiceClass returns the ServiceClass field if non-nil, zero value otherwise.

### GetServiceClassOk

`func (o *WindowsIntegratedAuthenticationDTO) GetServiceClassOk() (*string, bool)`

GetServiceClassOk returns a tuple with the ServiceClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceClass

`func (o *WindowsIntegratedAuthenticationDTO) SetServiceClass(v string)`

SetServiceClass sets ServiceClass field to given value.

### HasServiceClass

`func (o *WindowsIntegratedAuthenticationDTO) HasServiceClass() bool`

HasServiceClass returns a boolean if a field has been set.

### GetServiceName

`func (o *WindowsIntegratedAuthenticationDTO) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *WindowsIntegratedAuthenticationDTO) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *WindowsIntegratedAuthenticationDTO) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *WindowsIntegratedAuthenticationDTO) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetX

`func (o *WindowsIntegratedAuthenticationDTO) GetX() float64`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *WindowsIntegratedAuthenticationDTO) GetXOk() (*float64, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *WindowsIntegratedAuthenticationDTO) SetX(v float64)`

SetX sets X field to given value.

### HasX

`func (o *WindowsIntegratedAuthenticationDTO) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *WindowsIntegratedAuthenticationDTO) GetY() float64`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *WindowsIntegratedAuthenticationDTO) GetYOk() (*float64, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *WindowsIntegratedAuthenticationDTO) SetY(v float64)`

SetY sets Y field to given value.

### HasY

`func (o *WindowsIntegratedAuthenticationDTO) HasY() bool`

HasY returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


