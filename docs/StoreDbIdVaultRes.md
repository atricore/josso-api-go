# StoreDbIdVaultRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DbIdVault** | Pointer to [**DbIdentityVaultDTO**](DbIdentityVaultDTO.md) |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**ValidationErrors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewStoreDbIdVaultRes

`func NewStoreDbIdVaultRes() *StoreDbIdVaultRes`

NewStoreDbIdVaultRes instantiates a new StoreDbIdVaultRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreDbIdVaultResWithDefaults

`func NewStoreDbIdVaultResWithDefaults() *StoreDbIdVaultRes`

NewStoreDbIdVaultResWithDefaults instantiates a new StoreDbIdVaultRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDbIdVault

`func (o *StoreDbIdVaultRes) GetDbIdVault() DbIdentityVaultDTO`

GetDbIdVault returns the DbIdVault field if non-nil, zero value otherwise.

### GetDbIdVaultOk

`func (o *StoreDbIdVaultRes) GetDbIdVaultOk() (*DbIdentityVaultDTO, bool)`

GetDbIdVaultOk returns a tuple with the DbIdVault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbIdVault

`func (o *StoreDbIdVaultRes) SetDbIdVault(v DbIdentityVaultDTO)`

SetDbIdVault sets DbIdVault field to given value.

### HasDbIdVault

`func (o *StoreDbIdVaultRes) HasDbIdVault() bool`

HasDbIdVault returns a boolean if a field has been set.

### GetError

`func (o *StoreDbIdVaultRes) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *StoreDbIdVaultRes) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *StoreDbIdVaultRes) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *StoreDbIdVaultRes) HasError() bool`

HasError returns a boolean if a field has been set.

### GetValidationErrors

`func (o *StoreDbIdVaultRes) GetValidationErrors() []string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *StoreDbIdVaultRes) GetValidationErrorsOk() (*[]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *StoreDbIdVaultRes) SetValidationErrors(v []string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *StoreDbIdVaultRes) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


