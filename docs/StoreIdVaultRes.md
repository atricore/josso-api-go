# StoreIdVaultRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**IdVault** | Pointer to [**EmbeddedIdentityVaultDTO**](EmbeddedIdentityVaultDTO.md) |  | [optional] 
**ValidationErrors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewStoreIdVaultRes

`func NewStoreIdVaultRes() *StoreIdVaultRes`

NewStoreIdVaultRes instantiates a new StoreIdVaultRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreIdVaultResWithDefaults

`func NewStoreIdVaultResWithDefaults() *StoreIdVaultRes`

NewStoreIdVaultResWithDefaults instantiates a new StoreIdVaultRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *StoreIdVaultRes) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *StoreIdVaultRes) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *StoreIdVaultRes) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *StoreIdVaultRes) HasError() bool`

HasError returns a boolean if a field has been set.

### GetIdVault

`func (o *StoreIdVaultRes) GetIdVault() EmbeddedIdentityVaultDTO`

GetIdVault returns the IdVault field if non-nil, zero value otherwise.

### GetIdVaultOk

`func (o *StoreIdVaultRes) GetIdVaultOk() (*EmbeddedIdentityVaultDTO, bool)`

GetIdVaultOk returns a tuple with the IdVault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdVault

`func (o *StoreIdVaultRes) SetIdVault(v EmbeddedIdentityVaultDTO)`

SetIdVault sets IdVault field to given value.

### HasIdVault

`func (o *StoreIdVaultRes) HasIdVault() bool`

HasIdVault returns a boolean if a field has been set.

### GetValidationErrors

`func (o *StoreIdVaultRes) GetValidationErrors() []string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *StoreIdVaultRes) GetValidationErrorsOk() (*[]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *StoreIdVaultRes) SetValidationErrors(v []string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *StoreIdVaultRes) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


