# StoreIdPSaml2Res

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**Idp** | Pointer to [**ExternalSaml2IdentityProviderDTO**](ExternalSaml2IdentityProviderDTO.md) |  | [optional] 
**ValidationErrors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewStoreIdPSaml2Res

`func NewStoreIdPSaml2Res() *StoreIdPSaml2Res`

NewStoreIdPSaml2Res instantiates a new StoreIdPSaml2Res object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreIdPSaml2ResWithDefaults

`func NewStoreIdPSaml2ResWithDefaults() *StoreIdPSaml2Res`

NewStoreIdPSaml2ResWithDefaults instantiates a new StoreIdPSaml2Res object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *StoreIdPSaml2Res) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *StoreIdPSaml2Res) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *StoreIdPSaml2Res) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *StoreIdPSaml2Res) HasError() bool`

HasError returns a boolean if a field has been set.

### GetIdp

`func (o *StoreIdPSaml2Res) GetIdp() ExternalSaml2IdentityProviderDTO`

GetIdp returns the Idp field if non-nil, zero value otherwise.

### GetIdpOk

`func (o *StoreIdPSaml2Res) GetIdpOk() (*ExternalSaml2IdentityProviderDTO, bool)`

GetIdpOk returns a tuple with the Idp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdp

`func (o *StoreIdPSaml2Res) SetIdp(v ExternalSaml2IdentityProviderDTO)`

SetIdp sets Idp field to given value.

### HasIdp

`func (o *StoreIdPSaml2Res) HasIdp() bool`

HasIdp returns a boolean if a field has been set.

### GetValidationErrors

`func (o *StoreIdPSaml2Res) GetValidationErrors() []string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *StoreIdPSaml2Res) GetValidationErrorsOk() (*[]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *StoreIdPSaml2Res) SetValidationErrors(v []string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *StoreIdPSaml2Res) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


