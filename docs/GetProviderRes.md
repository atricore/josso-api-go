# GetProviderRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to [**ProviderContainerDTO**](ProviderContainerDTO.md) |  | [optional] 
**ValidationErrors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetProviderRes

`func NewGetProviderRes() *GetProviderRes`

NewGetProviderRes instantiates a new GetProviderRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProviderResWithDefaults

`func NewGetProviderResWithDefaults() *GetProviderRes`

NewGetProviderResWithDefaults instantiates a new GetProviderRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *GetProviderRes) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetProviderRes) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetProviderRes) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetProviderRes) HasError() bool`

HasError returns a boolean if a field has been set.

### GetProvider

`func (o *GetProviderRes) GetProvider() ProviderContainerDTO`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *GetProviderRes) GetProviderOk() (*ProviderContainerDTO, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *GetProviderRes) SetProvider(v ProviderContainerDTO)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *GetProviderRes) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetValidationErrors

`func (o *GetProviderRes) GetValidationErrors() []string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *GetProviderRes) GetValidationErrorsOk() (*[]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *GetProviderRes) SetValidationErrors(v []string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *GetProviderRes) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


