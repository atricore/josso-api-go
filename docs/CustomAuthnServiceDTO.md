# CustomAuthnServiceDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthnCtxClass** | Pointer to **string** |  | [optional] 
**AuthnMechanismType** | Pointer to **string** |  | [optional] 
**CustomClass** | Pointer to [**CustomClassDTO**](CustomClassDTO.md) |  | [optional] 
**DelegatedAuthentications** | Pointer to [**[]DelegatedAuthenticationDTO**](DelegatedAuthenticationDTO.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PreAuthnClaimCollectorConfig** | Pointer to **string** |  | [optional] 
**PreAuthnClaimCollectorType** | Pointer to **string** |  | [optional] 
**PreAuthnServiceURL** | Pointer to **string** |  | [optional] 
**UseCredentialStore** | Pointer to **bool** |  | [optional] 
**X** | Pointer to **float64** |  | [optional] 
**Y** | Pointer to **float64** |  | [optional] 

## Methods

### NewCustomAuthnServiceDTO

`func NewCustomAuthnServiceDTO() *CustomAuthnServiceDTO`

NewCustomAuthnServiceDTO instantiates a new CustomAuthnServiceDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomAuthnServiceDTOWithDefaults

`func NewCustomAuthnServiceDTOWithDefaults() *CustomAuthnServiceDTO`

NewCustomAuthnServiceDTOWithDefaults instantiates a new CustomAuthnServiceDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthnCtxClass

`func (o *CustomAuthnServiceDTO) GetAuthnCtxClass() string`

GetAuthnCtxClass returns the AuthnCtxClass field if non-nil, zero value otherwise.

### GetAuthnCtxClassOk

`func (o *CustomAuthnServiceDTO) GetAuthnCtxClassOk() (*string, bool)`

GetAuthnCtxClassOk returns a tuple with the AuthnCtxClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthnCtxClass

`func (o *CustomAuthnServiceDTO) SetAuthnCtxClass(v string)`

SetAuthnCtxClass sets AuthnCtxClass field to given value.

### HasAuthnCtxClass

`func (o *CustomAuthnServiceDTO) HasAuthnCtxClass() bool`

HasAuthnCtxClass returns a boolean if a field has been set.

### GetAuthnMechanismType

`func (o *CustomAuthnServiceDTO) GetAuthnMechanismType() string`

GetAuthnMechanismType returns the AuthnMechanismType field if non-nil, zero value otherwise.

### GetAuthnMechanismTypeOk

`func (o *CustomAuthnServiceDTO) GetAuthnMechanismTypeOk() (*string, bool)`

GetAuthnMechanismTypeOk returns a tuple with the AuthnMechanismType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthnMechanismType

`func (o *CustomAuthnServiceDTO) SetAuthnMechanismType(v string)`

SetAuthnMechanismType sets AuthnMechanismType field to given value.

### HasAuthnMechanismType

`func (o *CustomAuthnServiceDTO) HasAuthnMechanismType() bool`

HasAuthnMechanismType returns a boolean if a field has been set.

### GetCustomClass

`func (o *CustomAuthnServiceDTO) GetCustomClass() CustomClassDTO`

GetCustomClass returns the CustomClass field if non-nil, zero value otherwise.

### GetCustomClassOk

`func (o *CustomAuthnServiceDTO) GetCustomClassOk() (*CustomClassDTO, bool)`

GetCustomClassOk returns a tuple with the CustomClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomClass

`func (o *CustomAuthnServiceDTO) SetCustomClass(v CustomClassDTO)`

SetCustomClass sets CustomClass field to given value.

### HasCustomClass

`func (o *CustomAuthnServiceDTO) HasCustomClass() bool`

HasCustomClass returns a boolean if a field has been set.

### GetDelegatedAuthentications

`func (o *CustomAuthnServiceDTO) GetDelegatedAuthentications() []DelegatedAuthenticationDTO`

GetDelegatedAuthentications returns the DelegatedAuthentications field if non-nil, zero value otherwise.

### GetDelegatedAuthenticationsOk

`func (o *CustomAuthnServiceDTO) GetDelegatedAuthenticationsOk() (*[]DelegatedAuthenticationDTO, bool)`

GetDelegatedAuthenticationsOk returns a tuple with the DelegatedAuthentications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatedAuthentications

`func (o *CustomAuthnServiceDTO) SetDelegatedAuthentications(v []DelegatedAuthenticationDTO)`

SetDelegatedAuthentications sets DelegatedAuthentications field to given value.

### HasDelegatedAuthentications

`func (o *CustomAuthnServiceDTO) HasDelegatedAuthentications() bool`

HasDelegatedAuthentications returns a boolean if a field has been set.

### GetDescription

`func (o *CustomAuthnServiceDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomAuthnServiceDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomAuthnServiceDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CustomAuthnServiceDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *CustomAuthnServiceDTO) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CustomAuthnServiceDTO) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CustomAuthnServiceDTO) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CustomAuthnServiceDTO) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetElementId

`func (o *CustomAuthnServiceDTO) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *CustomAuthnServiceDTO) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *CustomAuthnServiceDTO) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *CustomAuthnServiceDTO) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetId

`func (o *CustomAuthnServiceDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomAuthnServiceDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomAuthnServiceDTO) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CustomAuthnServiceDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CustomAuthnServiceDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomAuthnServiceDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomAuthnServiceDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CustomAuthnServiceDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPreAuthnClaimCollectorConfig

`func (o *CustomAuthnServiceDTO) GetPreAuthnClaimCollectorConfig() string`

GetPreAuthnClaimCollectorConfig returns the PreAuthnClaimCollectorConfig field if non-nil, zero value otherwise.

### GetPreAuthnClaimCollectorConfigOk

`func (o *CustomAuthnServiceDTO) GetPreAuthnClaimCollectorConfigOk() (*string, bool)`

GetPreAuthnClaimCollectorConfigOk returns a tuple with the PreAuthnClaimCollectorConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreAuthnClaimCollectorConfig

`func (o *CustomAuthnServiceDTO) SetPreAuthnClaimCollectorConfig(v string)`

SetPreAuthnClaimCollectorConfig sets PreAuthnClaimCollectorConfig field to given value.

### HasPreAuthnClaimCollectorConfig

`func (o *CustomAuthnServiceDTO) HasPreAuthnClaimCollectorConfig() bool`

HasPreAuthnClaimCollectorConfig returns a boolean if a field has been set.

### GetPreAuthnClaimCollectorType

`func (o *CustomAuthnServiceDTO) GetPreAuthnClaimCollectorType() string`

GetPreAuthnClaimCollectorType returns the PreAuthnClaimCollectorType field if non-nil, zero value otherwise.

### GetPreAuthnClaimCollectorTypeOk

`func (o *CustomAuthnServiceDTO) GetPreAuthnClaimCollectorTypeOk() (*string, bool)`

GetPreAuthnClaimCollectorTypeOk returns a tuple with the PreAuthnClaimCollectorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreAuthnClaimCollectorType

`func (o *CustomAuthnServiceDTO) SetPreAuthnClaimCollectorType(v string)`

SetPreAuthnClaimCollectorType sets PreAuthnClaimCollectorType field to given value.

### HasPreAuthnClaimCollectorType

`func (o *CustomAuthnServiceDTO) HasPreAuthnClaimCollectorType() bool`

HasPreAuthnClaimCollectorType returns a boolean if a field has been set.

### GetPreAuthnServiceURL

`func (o *CustomAuthnServiceDTO) GetPreAuthnServiceURL() string`

GetPreAuthnServiceURL returns the PreAuthnServiceURL field if non-nil, zero value otherwise.

### GetPreAuthnServiceURLOk

`func (o *CustomAuthnServiceDTO) GetPreAuthnServiceURLOk() (*string, bool)`

GetPreAuthnServiceURLOk returns a tuple with the PreAuthnServiceURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreAuthnServiceURL

`func (o *CustomAuthnServiceDTO) SetPreAuthnServiceURL(v string)`

SetPreAuthnServiceURL sets PreAuthnServiceURL field to given value.

### HasPreAuthnServiceURL

`func (o *CustomAuthnServiceDTO) HasPreAuthnServiceURL() bool`

HasPreAuthnServiceURL returns a boolean if a field has been set.

### GetUseCredentialStore

`func (o *CustomAuthnServiceDTO) GetUseCredentialStore() bool`

GetUseCredentialStore returns the UseCredentialStore field if non-nil, zero value otherwise.

### GetUseCredentialStoreOk

`func (o *CustomAuthnServiceDTO) GetUseCredentialStoreOk() (*bool, bool)`

GetUseCredentialStoreOk returns a tuple with the UseCredentialStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCredentialStore

`func (o *CustomAuthnServiceDTO) SetUseCredentialStore(v bool)`

SetUseCredentialStore sets UseCredentialStore field to given value.

### HasUseCredentialStore

`func (o *CustomAuthnServiceDTO) HasUseCredentialStore() bool`

HasUseCredentialStore returns a boolean if a field has been set.

### GetX

`func (o *CustomAuthnServiceDTO) GetX() float64`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *CustomAuthnServiceDTO) GetXOk() (*float64, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *CustomAuthnServiceDTO) SetX(v float64)`

SetX sets X field to given value.

### HasX

`func (o *CustomAuthnServiceDTO) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *CustomAuthnServiceDTO) GetY() float64`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *CustomAuthnServiceDTO) GetYOk() (*float64, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *CustomAuthnServiceDTO) SetY(v float64)`

SetY sets Y field to given value.

### HasY

`func (o *CustomAuthnServiceDTO) HasY() bool`

HasY returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


