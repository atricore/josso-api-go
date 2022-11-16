# StoreDbIdVaultReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DbIdVault** | Pointer to [**DbIdentityVaultDTO**](DbIdentityVaultDTO.md) |  | [optional] 
**IdOrName** | Pointer to **string** |  | [optional] 

## Methods

### NewStoreDbIdVaultReq

`func NewStoreDbIdVaultReq() *StoreDbIdVaultReq`

NewStoreDbIdVaultReq instantiates a new StoreDbIdVaultReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreDbIdVaultReqWithDefaults

`func NewStoreDbIdVaultReqWithDefaults() *StoreDbIdVaultReq`

NewStoreDbIdVaultReqWithDefaults instantiates a new StoreDbIdVaultReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDbIdVault

`func (o *StoreDbIdVaultReq) GetDbIdVault() DbIdentityVaultDTO`

GetDbIdVault returns the DbIdVault field if non-nil, zero value otherwise.

### GetDbIdVaultOk

`func (o *StoreDbIdVaultReq) GetDbIdVaultOk() (*DbIdentityVaultDTO, bool)`

GetDbIdVaultOk returns a tuple with the DbIdVault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbIdVault

`func (o *StoreDbIdVaultReq) SetDbIdVault(v DbIdentityVaultDTO)`

SetDbIdVault sets DbIdVault field to given value.

### HasDbIdVault

`func (o *StoreDbIdVaultReq) HasDbIdVault() bool`

HasDbIdVault returns a boolean if a field has been set.

### GetIdOrName

`func (o *StoreDbIdVaultReq) GetIdOrName() string`

GetIdOrName returns the IdOrName field if non-nil, zero value otherwise.

### GetIdOrNameOk

`func (o *StoreDbIdVaultReq) GetIdOrNameOk() (*string, bool)`

GetIdOrNameOk returns a tuple with the IdOrName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdOrName

`func (o *StoreDbIdVaultReq) SetIdOrName(v string)`

SetIdOrName sets IdOrName field to given value.

### HasIdOrName

`func (o *StoreDbIdVaultReq) HasIdOrName() bool`

HasIdOrName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


