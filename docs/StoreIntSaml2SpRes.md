# StoreIntSaml2SpRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**Sp** | Pointer to [**InternalSaml2ServiceProviderDTO**](InternalSaml2ServiceProviderDTO.md) |  | [optional] 
**ValidationErrors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewStoreIntSaml2SpRes

`func NewStoreIntSaml2SpRes() *StoreIntSaml2SpRes`

NewStoreIntSaml2SpRes instantiates a new StoreIntSaml2SpRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreIntSaml2SpResWithDefaults

`func NewStoreIntSaml2SpResWithDefaults() *StoreIntSaml2SpRes`

NewStoreIntSaml2SpResWithDefaults instantiates a new StoreIntSaml2SpRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *StoreIntSaml2SpRes) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *StoreIntSaml2SpRes) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *StoreIntSaml2SpRes) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *StoreIntSaml2SpRes) HasError() bool`

HasError returns a boolean if a field has been set.

### GetSp

`func (o *StoreIntSaml2SpRes) GetSp() InternalSaml2ServiceProviderDTO`

GetSp returns the Sp field if non-nil, zero value otherwise.

### GetSpOk

`func (o *StoreIntSaml2SpRes) GetSpOk() (*InternalSaml2ServiceProviderDTO, bool)`

GetSpOk returns a tuple with the Sp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSp

`func (o *StoreIntSaml2SpRes) SetSp(v InternalSaml2ServiceProviderDTO)`

SetSp sets Sp field to given value.

### HasSp

`func (o *StoreIntSaml2SpRes) HasSp() bool`

HasSp returns a boolean if a field has been set.

### GetValidationErrors

`func (o *StoreIntSaml2SpRes) GetValidationErrors() []string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *StoreIntSaml2SpRes) GetValidationErrorsOk() (*[]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *StoreIntSaml2SpRes) SetValidationErrors(v []string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *StoreIntSaml2SpRes) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


