# GetDbIdVaultsRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DbIdVaults** | Pointer to [**[]DbIdentityVaultDTO**](DbIdentityVaultDTO.md) |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**ValidationErrors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetDbIdVaultsRes

`func NewGetDbIdVaultsRes() *GetDbIdVaultsRes`

NewGetDbIdVaultsRes instantiates a new GetDbIdVaultsRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDbIdVaultsResWithDefaults

`func NewGetDbIdVaultsResWithDefaults() *GetDbIdVaultsRes`

NewGetDbIdVaultsResWithDefaults instantiates a new GetDbIdVaultsRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDbIdVaults

`func (o *GetDbIdVaultsRes) GetDbIdVaults() []DbIdentityVaultDTO`

GetDbIdVaults returns the DbIdVaults field if non-nil, zero value otherwise.

### GetDbIdVaultsOk

`func (o *GetDbIdVaultsRes) GetDbIdVaultsOk() (*[]DbIdentityVaultDTO, bool)`

GetDbIdVaultsOk returns a tuple with the DbIdVaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbIdVaults

`func (o *GetDbIdVaultsRes) SetDbIdVaults(v []DbIdentityVaultDTO)`

SetDbIdVaults sets DbIdVaults field to given value.

### HasDbIdVaults

`func (o *GetDbIdVaultsRes) HasDbIdVaults() bool`

HasDbIdVaults returns a boolean if a field has been set.

### GetError

`func (o *GetDbIdVaultsRes) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetDbIdVaultsRes) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetDbIdVaultsRes) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetDbIdVaultsRes) HasError() bool`

HasError returns a boolean if a field has been set.

### GetValidationErrors

`func (o *GetDbIdVaultsRes) GetValidationErrors() []string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *GetDbIdVaultsRes) GetValidationErrorsOk() (*[]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *GetDbIdVaultsRes) SetValidationErrors(v []string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *GetDbIdVaultsRes) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


