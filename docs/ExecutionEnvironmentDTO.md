# ExecutionEnvironmentDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activations** | Pointer to [**[]ActivationDTO**](ActivationDTO.md) |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**BindingLocation** | Pointer to [**LocationDTO**](LocationDTO.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
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

### NewExecutionEnvironmentDTO

`func NewExecutionEnvironmentDTO() *ExecutionEnvironmentDTO`

NewExecutionEnvironmentDTO instantiates a new ExecutionEnvironmentDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionEnvironmentDTOWithDefaults

`func NewExecutionEnvironmentDTOWithDefaults() *ExecutionEnvironmentDTO`

NewExecutionEnvironmentDTOWithDefaults instantiates a new ExecutionEnvironmentDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivations

`func (o *ExecutionEnvironmentDTO) GetActivations() []ActivationDTO`

GetActivations returns the Activations field if non-nil, zero value otherwise.

### GetActivationsOk

`func (o *ExecutionEnvironmentDTO) GetActivationsOk() (*[]ActivationDTO, bool)`

GetActivationsOk returns a tuple with the Activations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivations

`func (o *ExecutionEnvironmentDTO) SetActivations(v []ActivationDTO)`

SetActivations sets Activations field to given value.

### HasActivations

`func (o *ExecutionEnvironmentDTO) HasActivations() bool`

HasActivations returns a boolean if a field has been set.

### GetActive

`func (o *ExecutionEnvironmentDTO) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ExecutionEnvironmentDTO) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ExecutionEnvironmentDTO) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ExecutionEnvironmentDTO) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetBindingLocation

`func (o *ExecutionEnvironmentDTO) GetBindingLocation() LocationDTO`

GetBindingLocation returns the BindingLocation field if non-nil, zero value otherwise.

### GetBindingLocationOk

`func (o *ExecutionEnvironmentDTO) GetBindingLocationOk() (*LocationDTO, bool)`

GetBindingLocationOk returns a tuple with the BindingLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingLocation

`func (o *ExecutionEnvironmentDTO) SetBindingLocation(v LocationDTO)`

SetBindingLocation sets BindingLocation field to given value.

### HasBindingLocation

`func (o *ExecutionEnvironmentDTO) HasBindingLocation() bool`

HasBindingLocation returns a boolean if a field has been set.

### GetDescription

`func (o *ExecutionEnvironmentDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExecutionEnvironmentDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExecutionEnvironmentDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExecutionEnvironmentDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *ExecutionEnvironmentDTO) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ExecutionEnvironmentDTO) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ExecutionEnvironmentDTO) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ExecutionEnvironmentDTO) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetElementId

`func (o *ExecutionEnvironmentDTO) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *ExecutionEnvironmentDTO) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *ExecutionEnvironmentDTO) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *ExecutionEnvironmentDTO) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetId

`func (o *ExecutionEnvironmentDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExecutionEnvironmentDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExecutionEnvironmentDTO) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ExecutionEnvironmentDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstallDemoApps

`func (o *ExecutionEnvironmentDTO) GetInstallDemoApps() bool`

GetInstallDemoApps returns the InstallDemoApps field if non-nil, zero value otherwise.

### GetInstallDemoAppsOk

`func (o *ExecutionEnvironmentDTO) GetInstallDemoAppsOk() (*bool, bool)`

GetInstallDemoAppsOk returns a tuple with the InstallDemoApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallDemoApps

`func (o *ExecutionEnvironmentDTO) SetInstallDemoApps(v bool)`

SetInstallDemoApps sets InstallDemoApps field to given value.

### HasInstallDemoApps

`func (o *ExecutionEnvironmentDTO) HasInstallDemoApps() bool`

HasInstallDemoApps returns a boolean if a field has been set.

### GetInstallUri

`func (o *ExecutionEnvironmentDTO) GetInstallUri() string`

GetInstallUri returns the InstallUri field if non-nil, zero value otherwise.

### GetInstallUriOk

`func (o *ExecutionEnvironmentDTO) GetInstallUriOk() (*string, bool)`

GetInstallUriOk returns a tuple with the InstallUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallUri

`func (o *ExecutionEnvironmentDTO) SetInstallUri(v string)`

SetInstallUri sets InstallUri field to given value.

### HasInstallUri

`func (o *ExecutionEnvironmentDTO) HasInstallUri() bool`

HasInstallUri returns a boolean if a field has been set.

### GetLocation

`func (o *ExecutionEnvironmentDTO) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ExecutionEnvironmentDTO) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ExecutionEnvironmentDTO) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ExecutionEnvironmentDTO) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetName

`func (o *ExecutionEnvironmentDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExecutionEnvironmentDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExecutionEnvironmentDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExecutionEnvironmentDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOverwriteOriginalSetup

`func (o *ExecutionEnvironmentDTO) GetOverwriteOriginalSetup() bool`

GetOverwriteOriginalSetup returns the OverwriteOriginalSetup field if non-nil, zero value otherwise.

### GetOverwriteOriginalSetupOk

`func (o *ExecutionEnvironmentDTO) GetOverwriteOriginalSetupOk() (*bool, bool)`

GetOverwriteOriginalSetupOk returns a tuple with the OverwriteOriginalSetup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteOriginalSetup

`func (o *ExecutionEnvironmentDTO) SetOverwriteOriginalSetup(v bool)`

SetOverwriteOriginalSetup sets OverwriteOriginalSetup field to given value.

### HasOverwriteOriginalSetup

`func (o *ExecutionEnvironmentDTO) HasOverwriteOriginalSetup() bool`

HasOverwriteOriginalSetup returns a boolean if a field has been set.

### GetPlatformId

`func (o *ExecutionEnvironmentDTO) GetPlatformId() string`

GetPlatformId returns the PlatformId field if non-nil, zero value otherwise.

### GetPlatformIdOk

`func (o *ExecutionEnvironmentDTO) GetPlatformIdOk() (*string, bool)`

GetPlatformIdOk returns a tuple with the PlatformId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformId

`func (o *ExecutionEnvironmentDTO) SetPlatformId(v string)`

SetPlatformId sets PlatformId field to given value.

### HasPlatformId

`func (o *ExecutionEnvironmentDTO) HasPlatformId() bool`

HasPlatformId returns a boolean if a field has been set.

### GetTargetJDK

`func (o *ExecutionEnvironmentDTO) GetTargetJDK() string`

GetTargetJDK returns the TargetJDK field if non-nil, zero value otherwise.

### GetTargetJDKOk

`func (o *ExecutionEnvironmentDTO) GetTargetJDKOk() (*string, bool)`

GetTargetJDKOk returns a tuple with the TargetJDK field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetJDK

`func (o *ExecutionEnvironmentDTO) SetTargetJDK(v string)`

SetTargetJDK sets TargetJDK field to given value.

### HasTargetJDK

`func (o *ExecutionEnvironmentDTO) HasTargetJDK() bool`

HasTargetJDK returns a boolean if a field has been set.

### GetType

`func (o *ExecutionEnvironmentDTO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExecutionEnvironmentDTO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExecutionEnvironmentDTO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ExecutionEnvironmentDTO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetX

`func (o *ExecutionEnvironmentDTO) GetX() float64`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *ExecutionEnvironmentDTO) GetXOk() (*float64, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *ExecutionEnvironmentDTO) SetX(v float64)`

SetX sets X field to given value.

### HasX

`func (o *ExecutionEnvironmentDTO) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *ExecutionEnvironmentDTO) GetY() float64`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *ExecutionEnvironmentDTO) GetYOk() (*float64, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *ExecutionEnvironmentDTO) SetY(v float64)`

SetY sets Y field to given value.

### HasY

`func (o *ExecutionEnvironmentDTO) HasY() bool`

HasY returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


