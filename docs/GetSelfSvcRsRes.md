# GetSelfSvcRsRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**Resource** | Pointer to [**SelfServicesResourceDTO**](SelfServicesResourceDTO.md) |  | [optional] 
**ValidationErrors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetSelfSvcRsRes

`func NewGetSelfSvcRsRes() *GetSelfSvcRsRes`

NewGetSelfSvcRsRes instantiates a new GetSelfSvcRsRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSelfSvcRsResWithDefaults

`func NewGetSelfSvcRsResWithDefaults() *GetSelfSvcRsRes`

NewGetSelfSvcRsResWithDefaults instantiates a new GetSelfSvcRsRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *GetSelfSvcRsRes) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetSelfSvcRsRes) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetSelfSvcRsRes) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetSelfSvcRsRes) HasError() bool`

HasError returns a boolean if a field has been set.

### GetResource

`func (o *GetSelfSvcRsRes) GetResource() SelfServicesResourceDTO`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *GetSelfSvcRsRes) GetResourceOk() (*SelfServicesResourceDTO, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *GetSelfSvcRsRes) SetResource(v SelfServicesResourceDTO)`

SetResource sets Resource field to given value.

### HasResource

`func (o *GetSelfSvcRsRes) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetValidationErrors

`func (o *GetSelfSvcRsRes) GetValidationErrors() []string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *GetSelfSvcRsRes) GetValidationErrorsOk() (*[]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *GetSelfSvcRsRes) SetValidationErrors(v []string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *GetSelfSvcRsRes) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


