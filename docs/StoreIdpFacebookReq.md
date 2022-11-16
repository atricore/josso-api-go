# StoreIdpFacebookReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdOrName** | Pointer to **string** |  | [optional] 
**Idp** | Pointer to [**FacebookOpenIDConnectIdentityProviderDTO**](FacebookOpenIDConnectIdentityProviderDTO.md) |  | [optional] 

## Methods

### NewStoreIdpFacebookReq

`func NewStoreIdpFacebookReq() *StoreIdpFacebookReq`

NewStoreIdpFacebookReq instantiates a new StoreIdpFacebookReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreIdpFacebookReqWithDefaults

`func NewStoreIdpFacebookReqWithDefaults() *StoreIdpFacebookReq`

NewStoreIdpFacebookReqWithDefaults instantiates a new StoreIdpFacebookReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdOrName

`func (o *StoreIdpFacebookReq) GetIdOrName() string`

GetIdOrName returns the IdOrName field if non-nil, zero value otherwise.

### GetIdOrNameOk

`func (o *StoreIdpFacebookReq) GetIdOrNameOk() (*string, bool)`

GetIdOrNameOk returns a tuple with the IdOrName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdOrName

`func (o *StoreIdpFacebookReq) SetIdOrName(v string)`

SetIdOrName sets IdOrName field to given value.

### HasIdOrName

`func (o *StoreIdpFacebookReq) HasIdOrName() bool`

HasIdOrName returns a boolean if a field has been set.

### GetIdp

`func (o *StoreIdpFacebookReq) GetIdp() FacebookOpenIDConnectIdentityProviderDTO`

GetIdp returns the Idp field if non-nil, zero value otherwise.

### GetIdpOk

`func (o *StoreIdpFacebookReq) GetIdpOk() (*FacebookOpenIDConnectIdentityProviderDTO, bool)`

GetIdpOk returns a tuple with the Idp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdp

`func (o *StoreIdpFacebookReq) SetIdp(v FacebookOpenIDConnectIdentityProviderDTO)`

SetIdp sets Idp field to given value.

### HasIdp

`func (o *StoreIdpFacebookReq) HasIdp() bool`

HasIdp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


