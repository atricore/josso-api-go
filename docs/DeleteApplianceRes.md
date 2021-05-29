# DeleteApplianceRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**Removed** | Pointer to **bool** |  | [optional] 
**ValidationErrors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewDeleteApplianceRes

`func NewDeleteApplianceRes() *DeleteApplianceRes`

NewDeleteApplianceRes instantiates a new DeleteApplianceRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteApplianceResWithDefaults

`func NewDeleteApplianceResWithDefaults() *DeleteApplianceRes`

NewDeleteApplianceResWithDefaults instantiates a new DeleteApplianceRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *DeleteApplianceRes) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *DeleteApplianceRes) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *DeleteApplianceRes) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *DeleteApplianceRes) HasError() bool`

HasError returns a boolean if a field has been set.

### GetRemoved

`func (o *DeleteApplianceRes) GetRemoved() bool`

GetRemoved returns the Removed field if non-nil, zero value otherwise.

### GetRemovedOk

`func (o *DeleteApplianceRes) GetRemovedOk() (*bool, bool)`

GetRemovedOk returns a tuple with the Removed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoved

`func (o *DeleteApplianceRes) SetRemoved(v bool)`

SetRemoved sets Removed field to given value.

### HasRemoved

`func (o *DeleteApplianceRes) HasRemoved() bool`

HasRemoved returns a boolean if a field has been set.

### GetValidationErrors

`func (o *DeleteApplianceRes) GetValidationErrors() []string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *DeleteApplianceRes) GetValidationErrorsOk() (*[]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *DeleteApplianceRes) SetValidationErrors(v []string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *DeleteApplianceRes) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


