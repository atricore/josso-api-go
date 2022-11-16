# StoreOidcRpReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdOrName** | Pointer to **string** |  | [optional] 
**OidcRp** | Pointer to [**ExternalOpenIDConnectRelayingPartyDTO**](ExternalOpenIDConnectRelayingPartyDTO.md) |  | [optional] 

## Methods

### NewStoreOidcRpReq

`func NewStoreOidcRpReq() *StoreOidcRpReq`

NewStoreOidcRpReq instantiates a new StoreOidcRpReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreOidcRpReqWithDefaults

`func NewStoreOidcRpReqWithDefaults() *StoreOidcRpReq`

NewStoreOidcRpReqWithDefaults instantiates a new StoreOidcRpReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdOrName

`func (o *StoreOidcRpReq) GetIdOrName() string`

GetIdOrName returns the IdOrName field if non-nil, zero value otherwise.

### GetIdOrNameOk

`func (o *StoreOidcRpReq) GetIdOrNameOk() (*string, bool)`

GetIdOrNameOk returns a tuple with the IdOrName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdOrName

`func (o *StoreOidcRpReq) SetIdOrName(v string)`

SetIdOrName sets IdOrName field to given value.

### HasIdOrName

`func (o *StoreOidcRpReq) HasIdOrName() bool`

HasIdOrName returns a boolean if a field has been set.

### GetOidcRp

`func (o *StoreOidcRpReq) GetOidcRp() ExternalOpenIDConnectRelayingPartyDTO`

GetOidcRp returns the OidcRp field if non-nil, zero value otherwise.

### GetOidcRpOk

`func (o *StoreOidcRpReq) GetOidcRpOk() (*ExternalOpenIDConnectRelayingPartyDTO, bool)`

GetOidcRpOk returns a tuple with the OidcRp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcRp

`func (o *StoreOidcRpReq) SetOidcRp(v ExternalOpenIDConnectRelayingPartyDTO)`

SetOidcRp sets OidcRp field to given value.

### HasOidcRp

`func (o *StoreOidcRpReq) HasOidcRp() bool`

HasOidcRp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


