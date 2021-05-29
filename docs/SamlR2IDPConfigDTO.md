# SamlR2IDPConfigDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**Encrypter** | Pointer to [**KeystoreDTO**](KeystoreDTO.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Signer** | Pointer to [**KeystoreDTO**](KeystoreDTO.md) |  | [optional] 
**UseSampleStore** | Pointer to **bool** |  | [optional] 
**UseSystemStore** | Pointer to **bool** |  | [optional] 

## Methods

### NewSamlR2IDPConfigDTO

`func NewSamlR2IDPConfigDTO() *SamlR2IDPConfigDTO`

NewSamlR2IDPConfigDTO instantiates a new SamlR2IDPConfigDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamlR2IDPConfigDTOWithDefaults

`func NewSamlR2IDPConfigDTOWithDefaults() *SamlR2IDPConfigDTO`

NewSamlR2IDPConfigDTOWithDefaults instantiates a new SamlR2IDPConfigDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *SamlR2IDPConfigDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SamlR2IDPConfigDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SamlR2IDPConfigDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SamlR2IDPConfigDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *SamlR2IDPConfigDTO) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SamlR2IDPConfigDTO) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SamlR2IDPConfigDTO) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *SamlR2IDPConfigDTO) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetElementId

`func (o *SamlR2IDPConfigDTO) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *SamlR2IDPConfigDTO) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *SamlR2IDPConfigDTO) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *SamlR2IDPConfigDTO) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetEncrypter

`func (o *SamlR2IDPConfigDTO) GetEncrypter() KeystoreDTO`

GetEncrypter returns the Encrypter field if non-nil, zero value otherwise.

### GetEncrypterOk

`func (o *SamlR2IDPConfigDTO) GetEncrypterOk() (*KeystoreDTO, bool)`

GetEncrypterOk returns a tuple with the Encrypter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypter

`func (o *SamlR2IDPConfigDTO) SetEncrypter(v KeystoreDTO)`

SetEncrypter sets Encrypter field to given value.

### HasEncrypter

`func (o *SamlR2IDPConfigDTO) HasEncrypter() bool`

HasEncrypter returns a boolean if a field has been set.

### GetId

`func (o *SamlR2IDPConfigDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SamlR2IDPConfigDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SamlR2IDPConfigDTO) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *SamlR2IDPConfigDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SamlR2IDPConfigDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SamlR2IDPConfigDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SamlR2IDPConfigDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SamlR2IDPConfigDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSigner

`func (o *SamlR2IDPConfigDTO) GetSigner() KeystoreDTO`

GetSigner returns the Signer field if non-nil, zero value otherwise.

### GetSignerOk

`func (o *SamlR2IDPConfigDTO) GetSignerOk() (*KeystoreDTO, bool)`

GetSignerOk returns a tuple with the Signer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigner

`func (o *SamlR2IDPConfigDTO) SetSigner(v KeystoreDTO)`

SetSigner sets Signer field to given value.

### HasSigner

`func (o *SamlR2IDPConfigDTO) HasSigner() bool`

HasSigner returns a boolean if a field has been set.

### GetUseSampleStore

`func (o *SamlR2IDPConfigDTO) GetUseSampleStore() bool`

GetUseSampleStore returns the UseSampleStore field if non-nil, zero value otherwise.

### GetUseSampleStoreOk

`func (o *SamlR2IDPConfigDTO) GetUseSampleStoreOk() (*bool, bool)`

GetUseSampleStoreOk returns a tuple with the UseSampleStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSampleStore

`func (o *SamlR2IDPConfigDTO) SetUseSampleStore(v bool)`

SetUseSampleStore sets UseSampleStore field to given value.

### HasUseSampleStore

`func (o *SamlR2IDPConfigDTO) HasUseSampleStore() bool`

HasUseSampleStore returns a boolean if a field has been set.

### GetUseSystemStore

`func (o *SamlR2IDPConfigDTO) GetUseSystemStore() bool`

GetUseSystemStore returns the UseSystemStore field if non-nil, zero value otherwise.

### GetUseSystemStoreOk

`func (o *SamlR2IDPConfigDTO) GetUseSystemStoreOk() (*bool, bool)`

GetUseSystemStoreOk returns a tuple with the UseSystemStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSystemStore

`func (o *SamlR2IDPConfigDTO) SetUseSystemStore(v bool)`

SetUseSystemStore sets UseSystemStore field to given value.

### HasUseSystemStore

`func (o *SamlR2IDPConfigDTO) HasUseSystemStore() bool`

HasUseSystemStore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


