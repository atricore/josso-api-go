# GetIdVaultRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**IdVault** | Pointer to [**EmbeddedIdentityVaultDTO**](EmbeddedIdentityVaultDTO.md) |  | [optional] 
**ValidationErrors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetIdVaultRes

`func NewGetIdVaultRes() *GetIdVaultRes`

NewGetIdVaultRes instantiates a new GetIdVaultRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIdVaultResWithDefaults

`func NewGetIdVaultResWithDefaults() *GetIdVaultRes`

NewGetIdVaultResWithDefaults instantiates a new GetIdVaultRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *GetIdVaultRes) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetIdVaultRes) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetIdVaultRes) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetIdVaultRes) HasError() bool`

HasError returns a boolean if a field has been set.

### GetIdVault

`func (o *GetIdVaultRes) GetIdVault() EmbeddedIdentityVaultDTO`

GetIdVault returns the IdVault field if non-nil, zero value otherwise.

### GetIdVaultOk

`func (o *GetIdVaultRes) GetIdVaultOk() (*EmbeddedIdentityVaultDTO, bool)`

GetIdVaultOk returns a tuple with the IdVault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdVault

`func (o *GetIdVaultRes) SetIdVault(v EmbeddedIdentityVaultDTO)`

SetIdVault sets IdVault field to given value.

### HasIdVault

`func (o *GetIdVaultRes) HasIdVault() bool`

HasIdVault returns a boolean if a field has been set.

### GetValidationErrors

`func (o *GetIdVaultRes) GetValidationErrors() []string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *GetIdVaultRes) GetValidationErrorsOk() (*[]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *GetIdVaultRes) SetValidationErrors(v []string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *GetIdVaultRes) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


