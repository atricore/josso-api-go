# StoreOidcIdpReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdOrName** | Pointer to **string** |  | [optional] 
**Idp** | Pointer to [**GenericOpenIDConnectIdentityProviderDTO**](GenericOpenIDConnectIdentityProviderDTO.md) |  | [optional] 

## Methods

### NewStoreOidcIdpReq

`func NewStoreOidcIdpReq() *StoreOidcIdpReq`

NewStoreOidcIdpReq instantiates a new StoreOidcIdpReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreOidcIdpReqWithDefaults

`func NewStoreOidcIdpReqWithDefaults() *StoreOidcIdpReq`

NewStoreOidcIdpReqWithDefaults instantiates a new StoreOidcIdpReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdOrName

`func (o *StoreOidcIdpReq) GetIdOrName() string`

GetIdOrName returns the IdOrName field if non-nil, zero value otherwise.

### GetIdOrNameOk

`func (o *StoreOidcIdpReq) GetIdOrNameOk() (*string, bool)`

GetIdOrNameOk returns a tuple with the IdOrName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdOrName

`func (o *StoreOidcIdpReq) SetIdOrName(v string)`

SetIdOrName sets IdOrName field to given value.

### HasIdOrName

`func (o *StoreOidcIdpReq) HasIdOrName() bool`

HasIdOrName returns a boolean if a field has been set.

### GetIdp

`func (o *StoreOidcIdpReq) GetIdp() GenericOpenIDConnectIdentityProviderDTO`

GetIdp returns the Idp field if non-nil, zero value otherwise.

### GetIdpOk

`func (o *StoreOidcIdpReq) GetIdpOk() (*GenericOpenIDConnectIdentityProviderDTO, bool)`

GetIdpOk returns a tuple with the Idp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdp

`func (o *StoreOidcIdpReq) SetIdp(v GenericOpenIDConnectIdentityProviderDTO)`

SetIdp sets Idp field to given value.

### HasIdp

`func (o *StoreOidcIdpReq) HasIdp() bool`

HasIdp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


