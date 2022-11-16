# StoreIdVaultReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdOrName** | Pointer to **string** |  | [optional] 
**IdVault** | Pointer to [**EmbeddedIdentityVaultDTO**](EmbeddedIdentityVaultDTO.md) |  | [optional] 

## Methods

### NewStoreIdVaultReq

`func NewStoreIdVaultReq() *StoreIdVaultReq`

NewStoreIdVaultReq instantiates a new StoreIdVaultReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreIdVaultReqWithDefaults

`func NewStoreIdVaultReqWithDefaults() *StoreIdVaultReq`

NewStoreIdVaultReqWithDefaults instantiates a new StoreIdVaultReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdOrName

`func (o *StoreIdVaultReq) GetIdOrName() string`

GetIdOrName returns the IdOrName field if non-nil, zero value otherwise.

### GetIdOrNameOk

`func (o *StoreIdVaultReq) GetIdOrNameOk() (*string, bool)`

GetIdOrNameOk returns a tuple with the IdOrName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdOrName

`func (o *StoreIdVaultReq) SetIdOrName(v string)`

SetIdOrName sets IdOrName field to given value.

### HasIdOrName

`func (o *StoreIdVaultReq) HasIdOrName() bool`

HasIdOrName returns a boolean if a field has been set.

### GetIdVault

`func (o *StoreIdVaultReq) GetIdVault() EmbeddedIdentityVaultDTO`

GetIdVault returns the IdVault field if non-nil, zero value otherwise.

### GetIdVaultOk

`func (o *StoreIdVaultReq) GetIdVaultOk() (*EmbeddedIdentityVaultDTO, bool)`

GetIdVaultOk returns a tuple with the IdVault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdVault

`func (o *StoreIdVaultReq) SetIdVault(v EmbeddedIdentityVaultDTO)`

SetIdVault sets IdVault field to given value.

### HasIdVault

`func (o *StoreIdVaultReq) HasIdVault() bool`

HasIdVault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


