# StoreIdPSaml2Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdOrName** | Pointer to **string** |  | [optional] 
**Idp** | Pointer to [**ExternalSaml2IdentityProviderDTO**](ExternalSaml2IdentityProviderDTO.md) |  | [optional] 

## Methods

### NewStoreIdPSaml2Req

`func NewStoreIdPSaml2Req() *StoreIdPSaml2Req`

NewStoreIdPSaml2Req instantiates a new StoreIdPSaml2Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreIdPSaml2ReqWithDefaults

`func NewStoreIdPSaml2ReqWithDefaults() *StoreIdPSaml2Req`

NewStoreIdPSaml2ReqWithDefaults instantiates a new StoreIdPSaml2Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdOrName

`func (o *StoreIdPSaml2Req) GetIdOrName() string`

GetIdOrName returns the IdOrName field if non-nil, zero value otherwise.

### GetIdOrNameOk

`func (o *StoreIdPSaml2Req) GetIdOrNameOk() (*string, bool)`

GetIdOrNameOk returns a tuple with the IdOrName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdOrName

`func (o *StoreIdPSaml2Req) SetIdOrName(v string)`

SetIdOrName sets IdOrName field to given value.

### HasIdOrName

`func (o *StoreIdPSaml2Req) HasIdOrName() bool`

HasIdOrName returns a boolean if a field has been set.

### GetIdp

`func (o *StoreIdPSaml2Req) GetIdp() ExternalSaml2IdentityProviderDTO`

GetIdp returns the Idp field if non-nil, zero value otherwise.

### GetIdpOk

`func (o *StoreIdPSaml2Req) GetIdpOk() (*ExternalSaml2IdentityProviderDTO, bool)`

GetIdpOk returns a tuple with the Idp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdp

`func (o *StoreIdPSaml2Req) SetIdp(v ExternalSaml2IdentityProviderDTO)`

SetIdp sets Idp field to given value.

### HasIdp

`func (o *StoreIdPSaml2Req) HasIdp() bool`

HasIdp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


