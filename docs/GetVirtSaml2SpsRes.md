# GetVirtSaml2SpsRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**Sps** | Pointer to [**[]VirtualSaml2ServiceProviderDTO**](VirtualSaml2ServiceProviderDTO.md) |  | [optional] 
**ValidationErrors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetVirtSaml2SpsRes

`func NewGetVirtSaml2SpsRes() *GetVirtSaml2SpsRes`

NewGetVirtSaml2SpsRes instantiates a new GetVirtSaml2SpsRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVirtSaml2SpsResWithDefaults

`func NewGetVirtSaml2SpsResWithDefaults() *GetVirtSaml2SpsRes`

NewGetVirtSaml2SpsResWithDefaults instantiates a new GetVirtSaml2SpsRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *GetVirtSaml2SpsRes) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetVirtSaml2SpsRes) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetVirtSaml2SpsRes) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetVirtSaml2SpsRes) HasError() bool`

HasError returns a boolean if a field has been set.

### GetSps

`func (o *GetVirtSaml2SpsRes) GetSps() []VirtualSaml2ServiceProviderDTO`

GetSps returns the Sps field if non-nil, zero value otherwise.

### GetSpsOk

`func (o *GetVirtSaml2SpsRes) GetSpsOk() (*[]VirtualSaml2ServiceProviderDTO, bool)`

GetSpsOk returns a tuple with the Sps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSps

`func (o *GetVirtSaml2SpsRes) SetSps(v []VirtualSaml2ServiceProviderDTO)`

SetSps sets Sps field to given value.

### HasSps

`func (o *GetVirtSaml2SpsRes) HasSps() bool`

HasSps returns a boolean if a field has been set.

### GetValidationErrors

`func (o *GetVirtSaml2SpsRes) GetValidationErrors() []string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *GetVirtSaml2SpsRes) GetValidationErrorsOk() (*[]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *GetVirtSaml2SpsRes) SetValidationErrors(v []string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *GetVirtSaml2SpsRes) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


