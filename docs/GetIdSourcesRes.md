# GetIdSourcesRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**IdSources** | Pointer to [**[]IdSourceContainerDTO**](IdSourceContainerDTO.md) |  | [optional] 
**ValidationErrors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetIdSourcesRes

`func NewGetIdSourcesRes() *GetIdSourcesRes`

NewGetIdSourcesRes instantiates a new GetIdSourcesRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIdSourcesResWithDefaults

`func NewGetIdSourcesResWithDefaults() *GetIdSourcesRes`

NewGetIdSourcesResWithDefaults instantiates a new GetIdSourcesRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *GetIdSourcesRes) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetIdSourcesRes) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetIdSourcesRes) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetIdSourcesRes) HasError() bool`

HasError returns a boolean if a field has been set.

### GetIdSources

`func (o *GetIdSourcesRes) GetIdSources() []IdSourceContainerDTO`

GetIdSources returns the IdSources field if non-nil, zero value otherwise.

### GetIdSourcesOk

`func (o *GetIdSourcesRes) GetIdSourcesOk() (*[]IdSourceContainerDTO, bool)`

GetIdSourcesOk returns a tuple with the IdSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdSources

`func (o *GetIdSourcesRes) SetIdSources(v []IdSourceContainerDTO)`

SetIdSources sets IdSources field to given value.

### HasIdSources

`func (o *GetIdSourcesRes) HasIdSources() bool`

HasIdSources returns a boolean if a field has been set.

### GetValidationErrors

`func (o *GetIdSourcesRes) GetValidationErrors() []string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *GetIdSourcesRes) GetValidationErrorsOk() (*[]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *GetIdSourcesRes) SetValidationErrors(v []string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *GetIdSourcesRes) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


