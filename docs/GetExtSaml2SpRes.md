# GetExtSaml2SpRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**SamlR2SPConfigDTO**](SamlR2SPConfigDTO.md) |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**Isp** | Pointer to [**ExternalSaml2ServiceProviderDTO**](ExternalSaml2ServiceProviderDTO.md) |  | [optional] 
**ValidationErrors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetExtSaml2SpRes

`func NewGetExtSaml2SpRes() *GetExtSaml2SpRes`

NewGetExtSaml2SpRes instantiates a new GetExtSaml2SpRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetExtSaml2SpResWithDefaults

`func NewGetExtSaml2SpResWithDefaults() *GetExtSaml2SpRes`

NewGetExtSaml2SpResWithDefaults instantiates a new GetExtSaml2SpRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *GetExtSaml2SpRes) GetConfig() SamlR2SPConfigDTO`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetExtSaml2SpRes) GetConfigOk() (*SamlR2SPConfigDTO, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetExtSaml2SpRes) SetConfig(v SamlR2SPConfigDTO)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetExtSaml2SpRes) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetError

`func (o *GetExtSaml2SpRes) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetExtSaml2SpRes) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetExtSaml2SpRes) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetExtSaml2SpRes) HasError() bool`

HasError returns a boolean if a field has been set.

### GetIsp

`func (o *GetExtSaml2SpRes) GetIsp() ExternalSaml2ServiceProviderDTO`

GetIsp returns the Isp field if non-nil, zero value otherwise.

### GetIspOk

`func (o *GetExtSaml2SpRes) GetIspOk() (*ExternalSaml2ServiceProviderDTO, bool)`

GetIspOk returns a tuple with the Isp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsp

`func (o *GetExtSaml2SpRes) SetIsp(v ExternalSaml2ServiceProviderDTO)`

SetIsp sets Isp field to given value.

### HasIsp

`func (o *GetExtSaml2SpRes) HasIsp() bool`

HasIsp returns a boolean if a field has been set.

### GetValidationErrors

`func (o *GetExtSaml2SpRes) GetValidationErrors() []string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *GetExtSaml2SpRes) GetValidationErrorsOk() (*[]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *GetExtSaml2SpRes) SetValidationErrors(v []string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *GetExtSaml2SpRes) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


