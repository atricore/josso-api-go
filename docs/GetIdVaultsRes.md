# GetIdVaultsRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**IdVaults** | Pointer to [**[]EmbeddedIdentityVaultDTO**](EmbeddedIdentityVaultDTO.md) |  | [optional] 
**ValidationErrors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetIdVaultsRes

`func NewGetIdVaultsRes() *GetIdVaultsRes`

NewGetIdVaultsRes instantiates a new GetIdVaultsRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIdVaultsResWithDefaults

`func NewGetIdVaultsResWithDefaults() *GetIdVaultsRes`

NewGetIdVaultsResWithDefaults instantiates a new GetIdVaultsRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *GetIdVaultsRes) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetIdVaultsRes) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetIdVaultsRes) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetIdVaultsRes) HasError() bool`

HasError returns a boolean if a field has been set.

### GetIdVaults

`func (o *GetIdVaultsRes) GetIdVaults() []EmbeddedIdentityVaultDTO`

GetIdVaults returns the IdVaults field if non-nil, zero value otherwise.

### GetIdVaultsOk

`func (o *GetIdVaultsRes) GetIdVaultsOk() (*[]EmbeddedIdentityVaultDTO, bool)`

GetIdVaultsOk returns a tuple with the IdVaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdVaults

`func (o *GetIdVaultsRes) SetIdVaults(v []EmbeddedIdentityVaultDTO)`

SetIdVaults sets IdVaults field to given value.

### HasIdVaults

`func (o *GetIdVaultsRes) HasIdVaults() bool`

HasIdVaults returns a boolean if a field has been set.

### GetValidationErrors

`func (o *GetIdVaultsRes) GetValidationErrors() []string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *GetIdVaultsRes) GetValidationErrorsOk() (*[]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *GetIdVaultsRes) SetValidationErrors(v []string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *GetIdVaultsRes) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


