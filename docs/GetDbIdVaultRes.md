# GetDbIdVaultRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**SamlR2SPConfigDTO**](SamlR2SPConfigDTO.md) |  | [optional] 
**DbIdVault** | Pointer to [**DbIdentityVaultDTO**](DbIdentityVaultDTO.md) |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**ValidationErrors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetDbIdVaultRes

`func NewGetDbIdVaultRes() *GetDbIdVaultRes`

NewGetDbIdVaultRes instantiates a new GetDbIdVaultRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDbIdVaultResWithDefaults

`func NewGetDbIdVaultResWithDefaults() *GetDbIdVaultRes`

NewGetDbIdVaultResWithDefaults instantiates a new GetDbIdVaultRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *GetDbIdVaultRes) GetConfig() SamlR2SPConfigDTO`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetDbIdVaultRes) GetConfigOk() (*SamlR2SPConfigDTO, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetDbIdVaultRes) SetConfig(v SamlR2SPConfigDTO)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetDbIdVaultRes) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetDbIdVault

`func (o *GetDbIdVaultRes) GetDbIdVault() DbIdentityVaultDTO`

GetDbIdVault returns the DbIdVault field if non-nil, zero value otherwise.

### GetDbIdVaultOk

`func (o *GetDbIdVaultRes) GetDbIdVaultOk() (*DbIdentityVaultDTO, bool)`

GetDbIdVaultOk returns a tuple with the DbIdVault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbIdVault

`func (o *GetDbIdVaultRes) SetDbIdVault(v DbIdentityVaultDTO)`

SetDbIdVault sets DbIdVault field to given value.

### HasDbIdVault

`func (o *GetDbIdVaultRes) HasDbIdVault() bool`

HasDbIdVault returns a boolean if a field has been set.

### GetError

`func (o *GetDbIdVaultRes) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetDbIdVaultRes) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetDbIdVaultRes) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetDbIdVaultRes) HasError() bool`

HasError returns a boolean if a field has been set.

### GetValidationErrors

`func (o *GetDbIdVaultRes) GetValidationErrors() []string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *GetDbIdVaultRes) GetValidationErrorsOk() (*[]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *GetDbIdVaultRes) SetValidationErrors(v []string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *GetDbIdVaultRes) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


