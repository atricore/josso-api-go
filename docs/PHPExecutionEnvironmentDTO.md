# PHPExecutionEnvironmentDTO

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
**PhpEnvironmentType** | Pointer to **string** |  | [optional] 
**PlatformId** | Pointer to **string** |  | [optional] 
**TargetJDK** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**X** | Pointer to **float64** |  | [optional] 
**Y** | Pointer to **float64** |  | [optional] 

## Methods

### NewPHPExecutionEnvironmentDTO

`func NewPHPExecutionEnvironmentDTO() *PHPExecutionEnvironmentDTO`

NewPHPExecutionEnvironmentDTO instantiates a new PHPExecutionEnvironmentDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPHPExecutionEnvironmentDTOWithDefaults

`func NewPHPExecutionEnvironmentDTOWithDefaults() *PHPExecutionEnvironmentDTO`

NewPHPExecutionEnvironmentDTOWithDefaults instantiates a new PHPExecutionEnvironmentDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivations

`func (o *PHPExecutionEnvironmentDTO) GetActivations() []ActivationDTO`

GetActivations returns the Activations field if non-nil, zero value otherwise.

### GetActivationsOk

`func (o *PHPExecutionEnvironmentDTO) GetActivationsOk() (*[]ActivationDTO, bool)`

GetActivationsOk returns a tuple with the Activations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivations

`func (o *PHPExecutionEnvironmentDTO) SetActivations(v []ActivationDTO)`

SetActivations sets Activations field to given value.

### HasActivations

`func (o *PHPExecutionEnvironmentDTO) HasActivations() bool`

HasActivations returns a boolean if a field has been set.

### GetActive

`func (o *PHPExecutionEnvironmentDTO) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PHPExecutionEnvironmentDTO) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PHPExecutionEnvironmentDTO) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PHPExecutionEnvironmentDTO) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetBindingLocation

`func (o *PHPExecutionEnvironmentDTO) GetBindingLocation() LocationDTO`

GetBindingLocation returns the BindingLocation field if non-nil, zero value otherwise.

### GetBindingLocationOk

`func (o *PHPExecutionEnvironmentDTO) GetBindingLocationOk() (*LocationDTO, bool)`

GetBindingLocationOk returns a tuple with the BindingLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingLocation

`func (o *PHPExecutionEnvironmentDTO) SetBindingLocation(v LocationDTO)`

SetBindingLocation sets BindingLocation field to given value.

### HasBindingLocation

`func (o *PHPExecutionEnvironmentDTO) HasBindingLocation() bool`

HasBindingLocation returns a boolean if a field has been set.

### GetDescription

`func (o *PHPExecutionEnvironmentDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PHPExecutionEnvironmentDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PHPExecutionEnvironmentDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PHPExecutionEnvironmentDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *PHPExecutionEnvironmentDTO) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PHPExecutionEnvironmentDTO) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PHPExecutionEnvironmentDTO) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *PHPExecutionEnvironmentDTO) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetElementId

`func (o *PHPExecutionEnvironmentDTO) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *PHPExecutionEnvironmentDTO) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *PHPExecutionEnvironmentDTO) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *PHPExecutionEnvironmentDTO) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetId

`func (o *PHPExecutionEnvironmentDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PHPExecutionEnvironmentDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PHPExecutionEnvironmentDTO) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PHPExecutionEnvironmentDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstallDemoApps

`func (o *PHPExecutionEnvironmentDTO) GetInstallDemoApps() bool`

GetInstallDemoApps returns the InstallDemoApps field if non-nil, zero value otherwise.

### GetInstallDemoAppsOk

`func (o *PHPExecutionEnvironmentDTO) GetInstallDemoAppsOk() (*bool, bool)`

GetInstallDemoAppsOk returns a tuple with the InstallDemoApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallDemoApps

`func (o *PHPExecutionEnvironmentDTO) SetInstallDemoApps(v bool)`

SetInstallDemoApps sets InstallDemoApps field to given value.

### HasInstallDemoApps

`func (o *PHPExecutionEnvironmentDTO) HasInstallDemoApps() bool`

HasInstallDemoApps returns a boolean if a field has been set.

### GetInstallUri

`func (o *PHPExecutionEnvironmentDTO) GetInstallUri() string`

GetInstallUri returns the InstallUri field if non-nil, zero value otherwise.

### GetInstallUriOk

`func (o *PHPExecutionEnvironmentDTO) GetInstallUriOk() (*string, bool)`

GetInstallUriOk returns a tuple with the InstallUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallUri

`func (o *PHPExecutionEnvironmentDTO) SetInstallUri(v string)`

SetInstallUri sets InstallUri field to given value.

### HasInstallUri

`func (o *PHPExecutionEnvironmentDTO) HasInstallUri() bool`

HasInstallUri returns a boolean if a field has been set.

### GetLocation

`func (o *PHPExecutionEnvironmentDTO) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PHPExecutionEnvironmentDTO) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PHPExecutionEnvironmentDTO) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PHPExecutionEnvironmentDTO) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetName

`func (o *PHPExecutionEnvironmentDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PHPExecutionEnvironmentDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PHPExecutionEnvironmentDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PHPExecutionEnvironmentDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOverwriteOriginalSetup

`func (o *PHPExecutionEnvironmentDTO) GetOverwriteOriginalSetup() bool`

GetOverwriteOriginalSetup returns the OverwriteOriginalSetup field if non-nil, zero value otherwise.

### GetOverwriteOriginalSetupOk

`func (o *PHPExecutionEnvironmentDTO) GetOverwriteOriginalSetupOk() (*bool, bool)`

GetOverwriteOriginalSetupOk returns a tuple with the OverwriteOriginalSetup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteOriginalSetup

`func (o *PHPExecutionEnvironmentDTO) SetOverwriteOriginalSetup(v bool)`

SetOverwriteOriginalSetup sets OverwriteOriginalSetup field to given value.

### HasOverwriteOriginalSetup

`func (o *PHPExecutionEnvironmentDTO) HasOverwriteOriginalSetup() bool`

HasOverwriteOriginalSetup returns a boolean if a field has been set.

### GetPhpEnvironmentType

`func (o *PHPExecutionEnvironmentDTO) GetPhpEnvironmentType() string`

GetPhpEnvironmentType returns the PhpEnvironmentType field if non-nil, zero value otherwise.

### GetPhpEnvironmentTypeOk

`func (o *PHPExecutionEnvironmentDTO) GetPhpEnvironmentTypeOk() (*string, bool)`

GetPhpEnvironmentTypeOk returns a tuple with the PhpEnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhpEnvironmentType

`func (o *PHPExecutionEnvironmentDTO) SetPhpEnvironmentType(v string)`

SetPhpEnvironmentType sets PhpEnvironmentType field to given value.

### HasPhpEnvironmentType

`func (o *PHPExecutionEnvironmentDTO) HasPhpEnvironmentType() bool`

HasPhpEnvironmentType returns a boolean if a field has been set.

### GetPlatformId

`func (o *PHPExecutionEnvironmentDTO) GetPlatformId() string`

GetPlatformId returns the PlatformId field if non-nil, zero value otherwise.

### GetPlatformIdOk

`func (o *PHPExecutionEnvironmentDTO) GetPlatformIdOk() (*string, bool)`

GetPlatformIdOk returns a tuple with the PlatformId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformId

`func (o *PHPExecutionEnvironmentDTO) SetPlatformId(v string)`

SetPlatformId sets PlatformId field to given value.

### HasPlatformId

`func (o *PHPExecutionEnvironmentDTO) HasPlatformId() bool`

HasPlatformId returns a boolean if a field has been set.

### GetTargetJDK

`func (o *PHPExecutionEnvironmentDTO) GetTargetJDK() string`

GetTargetJDK returns the TargetJDK field if non-nil, zero value otherwise.

### GetTargetJDKOk

`func (o *PHPExecutionEnvironmentDTO) GetTargetJDKOk() (*string, bool)`

GetTargetJDKOk returns a tuple with the TargetJDK field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetJDK

`func (o *PHPExecutionEnvironmentDTO) SetTargetJDK(v string)`

SetTargetJDK sets TargetJDK field to given value.

### HasTargetJDK

`func (o *PHPExecutionEnvironmentDTO) HasTargetJDK() bool`

HasTargetJDK returns a boolean if a field has been set.

### GetType

`func (o *PHPExecutionEnvironmentDTO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PHPExecutionEnvironmentDTO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PHPExecutionEnvironmentDTO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PHPExecutionEnvironmentDTO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetX

`func (o *PHPExecutionEnvironmentDTO) GetX() float64`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *PHPExecutionEnvironmentDTO) GetXOk() (*float64, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *PHPExecutionEnvironmentDTO) SetX(v float64)`

SetX sets X field to given value.

### HasX

`func (o *PHPExecutionEnvironmentDTO) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *PHPExecutionEnvironmentDTO) GetY() float64`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *PHPExecutionEnvironmentDTO) GetYOk() (*float64, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *PHPExecutionEnvironmentDTO) SetY(v float64)`

SetY sets Y field to given value.

### HasY

`func (o *PHPExecutionEnvironmentDTO) HasY() bool`

HasY returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


