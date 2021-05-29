# DeleteIdPRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**Removed** | Pointer to **bool** |  | [optional] 
**ValidationErrors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewDeleteIdPRes

`func NewDeleteIdPRes() *DeleteIdPRes`

NewDeleteIdPRes instantiates a new DeleteIdPRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteIdPResWithDefaults

`func NewDeleteIdPResWithDefaults() *DeleteIdPRes`

NewDeleteIdPResWithDefaults instantiates a new DeleteIdPRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *DeleteIdPRes) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *DeleteIdPRes) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *DeleteIdPRes) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *DeleteIdPRes) HasError() bool`

HasError returns a boolean if a field has been set.

### GetRemoved

`func (o *DeleteIdPRes) GetRemoved() bool`

GetRemoved returns the Removed field if non-nil, zero value otherwise.

### GetRemovedOk

`func (o *DeleteIdPRes) GetRemovedOk() (*bool, bool)`

GetRemovedOk returns a tuple with the Removed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoved

`func (o *DeleteIdPRes) SetRemoved(v bool)`

SetRemoved sets Removed field to given value.

### HasRemoved

`func (o *DeleteIdPRes) HasRemoved() bool`

HasRemoved returns a boolean if a field has been set.

### GetValidationErrors

`func (o *DeleteIdPRes) GetValidationErrors() []string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *DeleteIdPRes) GetValidationErrorsOk() (*[]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *DeleteIdPRes) SetValidationErrors(v []string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *DeleteIdPRes) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


