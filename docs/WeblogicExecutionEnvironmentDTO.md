# WeblogicExecutionEnvironmentDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activations** | Pointer to [**[]ActivationDTO**](ActivationDTO.md) |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**BindingLocation** | Pointer to [**LocationDTO**](LocationDTO.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**InstallDemoApps** | Pointer to **bool** |  | [optional] 
**InstallUri** | Pointer to **string** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OverwriteOriginalSetup** | Pointer to **bool** |  | [optional] 
**PlatformId** | Pointer to **string** |  | [optional] 
**TargetJDK** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**X** | Pointer to **float64** |  | [optional] 
**Y** | Pointer to **float64** |  | [optional] 

## Methods

### NewWeblogicExecutionEnvironmentDTO

`func NewWeblogicExecutionEnvironmentDTO() *WeblogicExecutionEnvironmentDTO`

NewWeblogicExecutionEnvironmentDTO instantiates a new WeblogicExecutionEnvironmentDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWeblogicExecutionEnvironmentDTOWithDefaults

`func NewWeblogicExecutionEnvironmentDTOWithDefaults() *WeblogicExecutionEnvironmentDTO`

NewWeblogicExecutionEnvironmentDTOWithDefaults instantiates a new WeblogicExecutionEnvironmentDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivations

`func (o *WeblogicExecutionEnvironmentDTO) GetActivations() []ActivationDTO`

GetActivations returns the Activations field if non-nil, zero value otherwise.

### GetActivationsOk

`func (o *WeblogicExecutionEnvironmentDTO) GetActivationsOk() (*[]ActivationDTO, bool)`

GetActivationsOk returns a tuple with the Activations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivations

`func (o *WeblogicExecutionEnvironmentDTO) SetActivations(v []ActivationDTO)`

SetActivations sets Activations field to given value.

### HasActivations

`func (o *WeblogicExecutionEnvironmentDTO) HasActivations() bool`

HasActivations returns a boolean if a field has been set.

### GetActive

`func (o *WeblogicExecutionEnvironmentDTO) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *WeblogicExecutionEnvironmentDTO) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *WeblogicExecutionEnvironmentDTO) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *WeblogicExecutionEnvironmentDTO) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetBindingLocation

`func (o *WeblogicExecutionEnvironmentDTO) GetBindingLocation() LocationDTO`

GetBindingLocation returns the BindingLocation field if non-nil, zero value otherwise.

### GetBindingLocationOk

`func (o *WeblogicExecutionEnvironmentDTO) GetBindingLocationOk() (*LocationDTO, bool)`

GetBindingLocationOk returns a tuple with the BindingLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingLocation

`func (o *WeblogicExecutionEnvironmentDTO) SetBindingLocation(v LocationDTO)`

SetBindingLocation sets BindingLocation field to given value.

### HasBindingLocation

`func (o *WeblogicExecutionEnvironmentDTO) HasBindingLocation() bool`

HasBindingLocation returns a boolean if a field has been set.

### GetDescription

`func (o *WeblogicExecutionEnvironmentDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WeblogicExecutionEnvironmentDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WeblogicExecutionEnvironmentDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WeblogicExecutionEnvironmentDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *WeblogicExecutionEnvironmentDTO) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *WeblogicExecutionEnvironmentDTO) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *WeblogicExecutionEnvironmentDTO) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *WeblogicExecutionEnvironmentDTO) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDomain

`func (o *WeblogicExecutionEnvironmentDTO) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *WeblogicExecutionEnvironmentDTO) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *WeblogicExecutionEnvironmentDTO) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *WeblogicExecutionEnvironmentDTO) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetElementId

`func (o *WeblogicExecutionEnvironmentDTO) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *WeblogicExecutionEnvironmentDTO) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *WeblogicExecutionEnvironmentDTO) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *WeblogicExecutionEnvironmentDTO) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetId

`func (o *WeblogicExecutionEnvironmentDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WeblogicExecutionEnvironmentDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WeblogicExecutionEnvironmentDTO) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *WeblogicExecutionEnvironmentDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstallDemoApps

`func (o *WeblogicExecutionEnvironmentDTO) GetInstallDemoApps() bool`

GetInstallDemoApps returns the InstallDemoApps field if non-nil, zero value otherwise.

### GetInstallDemoAppsOk

`func (o *WeblogicExecutionEnvironmentDTO) GetInstallDemoAppsOk() (*bool, bool)`

GetInstallDemoAppsOk returns a tuple with the InstallDemoApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallDemoApps

`func (o *WeblogicExecutionEnvironmentDTO) SetInstallDemoApps(v bool)`

SetInstallDemoApps sets InstallDemoApps field to given value.

### HasInstallDemoApps

`func (o *WeblogicExecutionEnvironmentDTO) HasInstallDemoApps() bool`

HasInstallDemoApps returns a boolean if a field has been set.

### GetInstallUri

`func (o *WeblogicExecutionEnvironmentDTO) GetInstallUri() string`

GetInstallUri returns the InstallUri field if non-nil, zero value otherwise.

### GetInstallUriOk

`func (o *WeblogicExecutionEnvironmentDTO) GetInstallUriOk() (*string, bool)`

GetInstallUriOk returns a tuple with the InstallUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallUri

`func (o *WeblogicExecutionEnvironmentDTO) SetInstallUri(v string)`

SetInstallUri sets InstallUri field to given value.

### HasInstallUri

`func (o *WeblogicExecutionEnvironmentDTO) HasInstallUri() bool`

HasInstallUri returns a boolean if a field has been set.

### GetLocation

`func (o *WeblogicExecutionEnvironmentDTO) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *WeblogicExecutionEnvironmentDTO) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *WeblogicExecutionEnvironmentDTO) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *WeblogicExecutionEnvironmentDTO) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetName

`func (o *WeblogicExecutionEnvironmentDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WeblogicExecutionEnvironmentDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WeblogicExecutionEnvironmentDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WeblogicExecutionEnvironmentDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOverwriteOriginalSetup

`func (o *WeblogicExecutionEnvironmentDTO) GetOverwriteOriginalSetup() bool`

GetOverwriteOriginalSetup returns the OverwriteOriginalSetup field if non-nil, zero value otherwise.

### GetOverwriteOriginalSetupOk

`func (o *WeblogicExecutionEnvironmentDTO) GetOverwriteOriginalSetupOk() (*bool, bool)`

GetOverwriteOriginalSetupOk returns a tuple with the OverwriteOriginalSetup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteOriginalSetup

`func (o *WeblogicExecutionEnvironmentDTO) SetOverwriteOriginalSetup(v bool)`

SetOverwriteOriginalSetup sets OverwriteOriginalSetup field to given value.

### HasOverwriteOriginalSetup

`func (o *WeblogicExecutionEnvironmentDTO) HasOverwriteOriginalSetup() bool`

HasOverwriteOriginalSetup returns a boolean if a field has been set.

### GetPlatformId

`func (o *WeblogicExecutionEnvironmentDTO) GetPlatformId() string`

GetPlatformId returns the PlatformId field if non-nil, zero value otherwise.

### GetPlatformIdOk

`func (o *WeblogicExecutionEnvironmentDTO) GetPlatformIdOk() (*string, bool)`

GetPlatformIdOk returns a tuple with the PlatformId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformId

`func (o *WeblogicExecutionEnvironmentDTO) SetPlatformId(v string)`

SetPlatformId sets PlatformId field to given value.

### HasPlatformId

`func (o *WeblogicExecutionEnvironmentDTO) HasPlatformId() bool`

HasPlatformId returns a boolean if a field has been set.

### GetTargetJDK

`func (o *WeblogicExecutionEnvironmentDTO) GetTargetJDK() string`

GetTargetJDK returns the TargetJDK field if non-nil, zero value otherwise.

### GetTargetJDKOk

`func (o *WeblogicExecutionEnvironmentDTO) GetTargetJDKOk() (*string, bool)`

GetTargetJDKOk returns a tuple with the TargetJDK field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetJDK

`func (o *WeblogicExecutionEnvironmentDTO) SetTargetJDK(v string)`

SetTargetJDK sets TargetJDK field to given value.

### HasTargetJDK

`func (o *WeblogicExecutionEnvironmentDTO) HasTargetJDK() bool`

HasTargetJDK returns a boolean if a field has been set.

### GetType

`func (o *WeblogicExecutionEnvironmentDTO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WeblogicExecutionEnvironmentDTO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WeblogicExecutionEnvironmentDTO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WeblogicExecutionEnvironmentDTO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetX

`func (o *WeblogicExecutionEnvironmentDTO) GetX() float64`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *WeblogicExecutionEnvironmentDTO) GetXOk() (*float64, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *WeblogicExecutionEnvironmentDTO) SetX(v float64)`

SetX sets X field to given value.

### HasX

`func (o *WeblogicExecutionEnvironmentDTO) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *WeblogicExecutionEnvironmentDTO) GetY() float64`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *WeblogicExecutionEnvironmentDTO) GetYOk() (*float64, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *WeblogicExecutionEnvironmentDTO) SetY(v float64)`

SetY sets Y field to given value.

### HasY

`func (o *WeblogicExecutionEnvironmentDTO) HasY() bool`

HasY returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


