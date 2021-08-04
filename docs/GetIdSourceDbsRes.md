# GetIdSourceDbsRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**IdSourceDbs** | Pointer to [**[]DbIdentitySourceDTO**](DbIdentitySourceDTO.md) |  | [optional] 
**ValidationErrors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetIdSourceDbsRes

`func NewGetIdSourceDbsRes() *GetIdSourceDbsRes`

NewGetIdSourceDbsRes instantiates a new GetIdSourceDbsRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIdSourceDbsResWithDefaults

`func NewGetIdSourceDbsResWithDefaults() *GetIdSourceDbsRes`

NewGetIdSourceDbsResWithDefaults instantiates a new GetIdSourceDbsRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *GetIdSourceDbsRes) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetIdSourceDbsRes) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetIdSourceDbsRes) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetIdSourceDbsRes) HasError() bool`

HasError returns a boolean if a field has been set.

### GetIdSourceDbs

`func (o *GetIdSourceDbsRes) GetIdSourceDbs() []DbIdentitySourceDTO`

GetIdSourceDbs returns the IdSourceDbs field if non-nil, zero value otherwise.

### GetIdSourceDbsOk

`func (o *GetIdSourceDbsRes) GetIdSourceDbsOk() (*[]DbIdentitySourceDTO, bool)`

GetIdSourceDbsOk returns a tuple with the IdSourceDbs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdSourceDbs

`func (o *GetIdSourceDbsRes) SetIdSourceDbs(v []DbIdentitySourceDTO)`

SetIdSourceDbs sets IdSourceDbs field to given value.

### HasIdSourceDbs

`func (o *GetIdSourceDbsRes) HasIdSourceDbs() bool`

HasIdSourceDbs returns a boolean if a field has been set.

### GetValidationErrors

`func (o *GetIdSourceDbsRes) GetValidationErrors() []string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *GetIdSourceDbsRes) GetValidationErrorsOk() (*[]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *GetIdSourceDbsRes) SetValidationErrors(v []string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *GetIdSourceDbsRes) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


