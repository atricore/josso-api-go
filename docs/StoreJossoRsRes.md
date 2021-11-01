# StoreJossoRsRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**Resource** | Pointer to [**JOSSO1ResourceDTO**](JOSSO1ResourceDTO.md) |  | [optional] 
**ValidationErrors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewStoreJossoRsRes

`func NewStoreJossoRsRes() *StoreJossoRsRes`

NewStoreJossoRsRes instantiates a new StoreJossoRsRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreJossoRsResWithDefaults

`func NewStoreJossoRsResWithDefaults() *StoreJossoRsRes`

NewStoreJossoRsResWithDefaults instantiates a new StoreJossoRsRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *StoreJossoRsRes) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *StoreJossoRsRes) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *StoreJossoRsRes) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *StoreJossoRsRes) HasError() bool`

HasError returns a boolean if a field has been set.

### GetResource

`func (o *StoreJossoRsRes) GetResource() JOSSO1ResourceDTO`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *StoreJossoRsRes) GetResourceOk() (*JOSSO1ResourceDTO, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *StoreJossoRsRes) SetResource(v JOSSO1ResourceDTO)`

SetResource sets Resource field to given value.

### HasResource

`func (o *StoreJossoRsRes) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetValidationErrors

`func (o *StoreJossoRsRes) GetValidationErrors() []string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *StoreJossoRsRes) GetValidationErrorsOk() (*[]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *StoreJossoRsRes) SetValidationErrors(v []string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *StoreJossoRsRes) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


