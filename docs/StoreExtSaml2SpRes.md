# StoreExtSaml2SpRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**Sp** | Pointer to [**ExternalSaml2ServiceProviderDTO**](ExternalSaml2ServiceProviderDTO.md) |  | [optional] 
**ValidationErrors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewStoreExtSaml2SpRes

`func NewStoreExtSaml2SpRes() *StoreExtSaml2SpRes`

NewStoreExtSaml2SpRes instantiates a new StoreExtSaml2SpRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreExtSaml2SpResWithDefaults

`func NewStoreExtSaml2SpResWithDefaults() *StoreExtSaml2SpRes`

NewStoreExtSaml2SpResWithDefaults instantiates a new StoreExtSaml2SpRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *StoreExtSaml2SpRes) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *StoreExtSaml2SpRes) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *StoreExtSaml2SpRes) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *StoreExtSaml2SpRes) HasError() bool`

HasError returns a boolean if a field has been set.

### GetSp

`func (o *StoreExtSaml2SpRes) GetSp() ExternalSaml2ServiceProviderDTO`

GetSp returns the Sp field if non-nil, zero value otherwise.

### GetSpOk

`func (o *StoreExtSaml2SpRes) GetSpOk() (*ExternalSaml2ServiceProviderDTO, bool)`

GetSpOk returns a tuple with the Sp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSp

`func (o *StoreExtSaml2SpRes) SetSp(v ExternalSaml2ServiceProviderDTO)`

SetSp sets Sp field to given value.

### HasSp

`func (o *StoreExtSaml2SpRes) HasSp() bool`

HasSp returns a boolean if a field has been set.

### GetValidationErrors

`func (o *StoreExtSaml2SpRes) GetValidationErrors() []string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *StoreExtSaml2SpRes) GetValidationErrorsOk() (*[]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *StoreExtSaml2SpRes) SetValidationErrors(v []string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *StoreExtSaml2SpRes) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


