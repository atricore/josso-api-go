# GetIntSaml2SpRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**SamlR2SPConfigDTO**](SamlR2SPConfigDTO.md) |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**Sp** | Pointer to [**InternalSaml2ServiceProviderDTO**](InternalSaml2ServiceProviderDTO.md) |  | [optional] 
**ValidationErrors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetIntSaml2SpRes

`func NewGetIntSaml2SpRes() *GetIntSaml2SpRes`

NewGetIntSaml2SpRes instantiates a new GetIntSaml2SpRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIntSaml2SpResWithDefaults

`func NewGetIntSaml2SpResWithDefaults() *GetIntSaml2SpRes`

NewGetIntSaml2SpResWithDefaults instantiates a new GetIntSaml2SpRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *GetIntSaml2SpRes) GetConfig() SamlR2SPConfigDTO`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetIntSaml2SpRes) GetConfigOk() (*SamlR2SPConfigDTO, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetIntSaml2SpRes) SetConfig(v SamlR2SPConfigDTO)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetIntSaml2SpRes) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetError

`func (o *GetIntSaml2SpRes) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetIntSaml2SpRes) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetIntSaml2SpRes) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetIntSaml2SpRes) HasError() bool`

HasError returns a boolean if a field has been set.

### GetSp

`func (o *GetIntSaml2SpRes) GetSp() InternalSaml2ServiceProviderDTO`

GetSp returns the Sp field if non-nil, zero value otherwise.

### GetSpOk

`func (o *GetIntSaml2SpRes) GetSpOk() (*InternalSaml2ServiceProviderDTO, bool)`

GetSpOk returns a tuple with the Sp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSp

`func (o *GetIntSaml2SpRes) SetSp(v InternalSaml2ServiceProviderDTO)`

SetSp sets Sp field to given value.

### HasSp

`func (o *GetIntSaml2SpRes) HasSp() bool`

HasSp returns a boolean if a field has been set.

### GetValidationErrors

`func (o *GetIntSaml2SpRes) GetValidationErrors() []string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *GetIntSaml2SpRes) GetValidationErrorsOk() (*[]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *GetIntSaml2SpRes) SetValidationErrors(v []string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *GetIntSaml2SpRes) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


