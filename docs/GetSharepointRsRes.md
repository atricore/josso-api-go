# GetSharepointRsRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**Resource** | Pointer to [**SharepointResourceDTO**](SharepointResourceDTO.md) |  | [optional] 
**ValidationErrors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetSharepointRsRes

`func NewGetSharepointRsRes() *GetSharepointRsRes`

NewGetSharepointRsRes instantiates a new GetSharepointRsRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSharepointRsResWithDefaults

`func NewGetSharepointRsResWithDefaults() *GetSharepointRsRes`

NewGetSharepointRsResWithDefaults instantiates a new GetSharepointRsRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *GetSharepointRsRes) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetSharepointRsRes) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetSharepointRsRes) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetSharepointRsRes) HasError() bool`

HasError returns a boolean if a field has been set.

### GetResource

`func (o *GetSharepointRsRes) GetResource() SharepointResourceDTO`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *GetSharepointRsRes) GetResourceOk() (*SharepointResourceDTO, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *GetSharepointRsRes) SetResource(v SharepointResourceDTO)`

SetResource sets Resource field to given value.

### HasResource

`func (o *GetSharepointRsRes) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetValidationErrors

`func (o *GetSharepointRsRes) GetValidationErrors() []string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *GetSharepointRsRes) GetValidationErrorsOk() (*[]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *GetSharepointRsRes) SetValidationErrors(v []string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *GetSharepointRsRes) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


