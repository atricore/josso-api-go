# GetSharepointRssRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**Resources** | Pointer to [**[]SharepointResourceDTO**](SharepointResourceDTO.md) |  | [optional] 
**ValidationErrors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetSharepointRssRes

`func NewGetSharepointRssRes() *GetSharepointRssRes`

NewGetSharepointRssRes instantiates a new GetSharepointRssRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSharepointRssResWithDefaults

`func NewGetSharepointRssResWithDefaults() *GetSharepointRssRes`

NewGetSharepointRssResWithDefaults instantiates a new GetSharepointRssRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *GetSharepointRssRes) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetSharepointRssRes) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetSharepointRssRes) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetSharepointRssRes) HasError() bool`

HasError returns a boolean if a field has been set.

### GetResources

`func (o *GetSharepointRssRes) GetResources() []SharepointResourceDTO`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *GetSharepointRssRes) GetResourcesOk() (*[]SharepointResourceDTO, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *GetSharepointRssRes) SetResources(v []SharepointResourceDTO)`

SetResources sets Resources field to given value.

### HasResources

`func (o *GetSharepointRssRes) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetValidationErrors

`func (o *GetSharepointRssRes) GetValidationErrors() []string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *GetSharepointRssRes) GetValidationErrorsOk() (*[]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *GetSharepointRssRes) SetValidationErrors(v []string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *GetSharepointRssRes) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


