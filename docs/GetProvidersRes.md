# GetProvidersRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**Providers** | Pointer to [**[]ProviderContainerDTO**](ProviderContainerDTO.md) |  | [optional] 
**ValidationErrors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetProvidersRes

`func NewGetProvidersRes() *GetProvidersRes`

NewGetProvidersRes instantiates a new GetProvidersRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProvidersResWithDefaults

`func NewGetProvidersResWithDefaults() *GetProvidersRes`

NewGetProvidersResWithDefaults instantiates a new GetProvidersRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *GetProvidersRes) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetProvidersRes) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetProvidersRes) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetProvidersRes) HasError() bool`

HasError returns a boolean if a field has been set.

### GetProviders

`func (o *GetProvidersRes) GetProviders() []ProviderContainerDTO`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *GetProvidersRes) GetProvidersOk() (*[]ProviderContainerDTO, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *GetProvidersRes) SetProviders(v []ProviderContainerDTO)`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *GetProvidersRes) HasProviders() bool`

HasProviders returns a boolean if a field has been set.

### GetValidationErrors

`func (o *GetProvidersRes) GetValidationErrors() []string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *GetProvidersRes) GetValidationErrorsOk() (*[]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *GetProvidersRes) SetValidationErrors(v []string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *GetProvidersRes) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


