# IdentityApplianceUnitDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BundleName** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**Group** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Providers** | Pointer to [**[]ProviderDTO**](ProviderDTO.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewIdentityApplianceUnitDTO

`func NewIdentityApplianceUnitDTO() *IdentityApplianceUnitDTO`

NewIdentityApplianceUnitDTO instantiates a new IdentityApplianceUnitDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityApplianceUnitDTOWithDefaults

`func NewIdentityApplianceUnitDTOWithDefaults() *IdentityApplianceUnitDTO`

NewIdentityApplianceUnitDTOWithDefaults instantiates a new IdentityApplianceUnitDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBundleName

`func (o *IdentityApplianceUnitDTO) GetBundleName() string`

GetBundleName returns the BundleName field if non-nil, zero value otherwise.

### GetBundleNameOk

`func (o *IdentityApplianceUnitDTO) GetBundleNameOk() (*string, bool)`

GetBundleNameOk returns a tuple with the BundleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleName

`func (o *IdentityApplianceUnitDTO) SetBundleName(v string)`

SetBundleName sets BundleName field to given value.

### HasBundleName

`func (o *IdentityApplianceUnitDTO) HasBundleName() bool`

HasBundleName returns a boolean if a field has been set.

### GetDescription

`func (o *IdentityApplianceUnitDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IdentityApplianceUnitDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IdentityApplianceUnitDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IdentityApplianceUnitDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetElementId

`func (o *IdentityApplianceUnitDTO) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *IdentityApplianceUnitDTO) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *IdentityApplianceUnitDTO) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *IdentityApplianceUnitDTO) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetGroup

`func (o *IdentityApplianceUnitDTO) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *IdentityApplianceUnitDTO) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *IdentityApplianceUnitDTO) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *IdentityApplianceUnitDTO) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetId

`func (o *IdentityApplianceUnitDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentityApplianceUnitDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentityApplianceUnitDTO) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *IdentityApplianceUnitDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IdentityApplianceUnitDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdentityApplianceUnitDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdentityApplianceUnitDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IdentityApplianceUnitDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProviders

`func (o *IdentityApplianceUnitDTO) GetProviders() []ProviderDTO`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *IdentityApplianceUnitDTO) GetProvidersOk() (*[]ProviderDTO, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *IdentityApplianceUnitDTO) SetProviders(v []ProviderDTO)`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *IdentityApplianceUnitDTO) HasProviders() bool`

HasProviders returns a boolean if a field has been set.

### GetType

`func (o *IdentityApplianceUnitDTO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IdentityApplianceUnitDTO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IdentityApplianceUnitDTO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IdentityApplianceUnitDTO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *IdentityApplianceUnitDTO) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *IdentityApplianceUnitDTO) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *IdentityApplianceUnitDTO) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *IdentityApplianceUnitDTO) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


