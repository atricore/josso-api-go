# StoreIdpGoogleReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdaName** | Pointer to **string** |  | [optional] 
**Idp** | Pointer to [**GoogleOpenIDConnectIdentityProviderDTO**](GoogleOpenIDConnectIdentityProviderDTO.md) |  | [optional] 

## Methods

### NewStoreIdpGoogleReq

`func NewStoreIdpGoogleReq() *StoreIdpGoogleReq`

NewStoreIdpGoogleReq instantiates a new StoreIdpGoogleReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreIdpGoogleReqWithDefaults

`func NewStoreIdpGoogleReqWithDefaults() *StoreIdpGoogleReq`

NewStoreIdpGoogleReqWithDefaults instantiates a new StoreIdpGoogleReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdaName

`func (o *StoreIdpGoogleReq) GetIdaName() string`

GetIdaName returns the IdaName field if non-nil, zero value otherwise.

### GetIdaNameOk

`func (o *StoreIdpGoogleReq) GetIdaNameOk() (*string, bool)`

GetIdaNameOk returns a tuple with the IdaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdaName

`func (o *StoreIdpGoogleReq) SetIdaName(v string)`

SetIdaName sets IdaName field to given value.

### HasIdaName

`func (o *StoreIdpGoogleReq) HasIdaName() bool`

HasIdaName returns a boolean if a field has been set.

### GetIdp

`func (o *StoreIdpGoogleReq) GetIdp() GoogleOpenIDConnectIdentityProviderDTO`

GetIdp returns the Idp field if non-nil, zero value otherwise.

### GetIdpOk

`func (o *StoreIdpGoogleReq) GetIdpOk() (*GoogleOpenIDConnectIdentityProviderDTO, bool)`

GetIdpOk returns a tuple with the Idp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdp

`func (o *StoreIdpGoogleReq) SetIdp(v GoogleOpenIDConnectIdentityProviderDTO)`

SetIdp sets Idp field to given value.

### HasIdp

`func (o *StoreIdpGoogleReq) HasIdp() bool`

HasIdp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


