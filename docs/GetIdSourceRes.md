# GetIdSourceRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**IdSource** | Pointer to [**IdSourceContainerDTO**](IdSourceContainerDTO.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**ValidationErrors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetIdSourceRes

`func NewGetIdSourceRes() *GetIdSourceRes`

NewGetIdSourceRes instantiates a new GetIdSourceRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIdSourceResWithDefaults

`func NewGetIdSourceResWithDefaults() *GetIdSourceRes`

NewGetIdSourceResWithDefaults instantiates a new GetIdSourceRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *GetIdSourceRes) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetIdSourceRes) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetIdSourceRes) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetIdSourceRes) HasError() bool`

HasError returns a boolean if a field has been set.

### GetIdSource

`func (o *GetIdSourceRes) GetIdSource() IdSourceContainerDTO`

GetIdSource returns the IdSource field if non-nil, zero value otherwise.

### GetIdSourceOk

`func (o *GetIdSourceRes) GetIdSourceOk() (*IdSourceContainerDTO, bool)`

GetIdSourceOk returns a tuple with the IdSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdSource

`func (o *GetIdSourceRes) SetIdSource(v IdSourceContainerDTO)`

SetIdSource sets IdSource field to given value.

### HasIdSource

`func (o *GetIdSourceRes) HasIdSource() bool`

HasIdSource returns a boolean if a field has been set.

### GetType

`func (o *GetIdSourceRes) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetIdSourceRes) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetIdSourceRes) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetIdSourceRes) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValidationErrors

`func (o *GetIdSourceRes) GetValidationErrors() []string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *GetIdSourceRes) GetValidationErrorsOk() (*[]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *GetIdSourceRes) SetValidationErrors(v []string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *GetIdSourceRes) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


