# StoreIdSourceLdapReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdOrName** | Pointer to **string** |  | [optional] 
**IdSourceLdap** | Pointer to [**LdapIdentitySourceDTO**](LdapIdentitySourceDTO.md) |  | [optional] 

## Methods

### NewStoreIdSourceLdapReq

`func NewStoreIdSourceLdapReq() *StoreIdSourceLdapReq`

NewStoreIdSourceLdapReq instantiates a new StoreIdSourceLdapReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreIdSourceLdapReqWithDefaults

`func NewStoreIdSourceLdapReqWithDefaults() *StoreIdSourceLdapReq`

NewStoreIdSourceLdapReqWithDefaults instantiates a new StoreIdSourceLdapReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdOrName

`func (o *StoreIdSourceLdapReq) GetIdOrName() string`

GetIdOrName returns the IdOrName field if non-nil, zero value otherwise.

### GetIdOrNameOk

`func (o *StoreIdSourceLdapReq) GetIdOrNameOk() (*string, bool)`

GetIdOrNameOk returns a tuple with the IdOrName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdOrName

`func (o *StoreIdSourceLdapReq) SetIdOrName(v string)`

SetIdOrName sets IdOrName field to given value.

### HasIdOrName

`func (o *StoreIdSourceLdapReq) HasIdOrName() bool`

HasIdOrName returns a boolean if a field has been set.

### GetIdSourceLdap

`func (o *StoreIdSourceLdapReq) GetIdSourceLdap() LdapIdentitySourceDTO`

GetIdSourceLdap returns the IdSourceLdap field if non-nil, zero value otherwise.

### GetIdSourceLdapOk

`func (o *StoreIdSourceLdapReq) GetIdSourceLdapOk() (*LdapIdentitySourceDTO, bool)`

GetIdSourceLdapOk returns a tuple with the IdSourceLdap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdSourceLdap

`func (o *StoreIdSourceLdapReq) SetIdSourceLdap(v LdapIdentitySourceDTO)`

SetIdSourceLdap sets IdSourceLdap field to given value.

### HasIdSourceLdap

`func (o *StoreIdSourceLdapReq) HasIdSourceLdap() bool`

HasIdSourceLdap returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


