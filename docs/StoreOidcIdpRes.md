# StoreOidcIdpRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**Idp** | Pointer to [**GenericOpenIDConnectIdentityProviderDTO**](GenericOpenIDConnectIdentityProviderDTO.md) |  | [optional] 
**ValidationErrors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewStoreOidcIdpRes

`func NewStoreOidcIdpRes() *StoreOidcIdpRes`

NewStoreOidcIdpRes instantiates a new StoreOidcIdpRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreOidcIdpResWithDefaults

`func NewStoreOidcIdpResWithDefaults() *StoreOidcIdpRes`

NewStoreOidcIdpResWithDefaults instantiates a new StoreOidcIdpRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *StoreOidcIdpRes) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *StoreOidcIdpRes) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *StoreOidcIdpRes) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *StoreOidcIdpRes) HasError() bool`

HasError returns a boolean if a field has been set.

### GetIdp

`func (o *StoreOidcIdpRes) GetIdp() GenericOpenIDConnectIdentityProviderDTO`

GetIdp returns the Idp field if non-nil, zero value otherwise.

### GetIdpOk

`func (o *StoreOidcIdpRes) GetIdpOk() (*GenericOpenIDConnectIdentityProviderDTO, bool)`

GetIdpOk returns a tuple with the Idp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdp

`func (o *StoreOidcIdpRes) SetIdp(v GenericOpenIDConnectIdentityProviderDTO)`

SetIdp sets Idp field to given value.

### HasIdp

`func (o *StoreOidcIdpRes) HasIdp() bool`

HasIdp returns a boolean if a field has been set.

### GetValidationErrors

`func (o *StoreOidcIdpRes) GetValidationErrors() []string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *StoreOidcIdpRes) GetValidationErrorsOk() (*[]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *StoreOidcIdpRes) SetValidationErrors(v []string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *StoreOidcIdpRes) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


