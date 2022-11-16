# StoreIdSourceDbReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdOrName** | Pointer to **string** |  | [optional] 
**IdSourceDb** | Pointer to [**DbIdentitySourceDTO**](DbIdentitySourceDTO.md) |  | [optional] 
**IdSourceLdap** | Pointer to [**DbIdentitySourceDTO**](DbIdentitySourceDTO.md) |  | [optional] 

## Methods

### NewStoreIdSourceDbReq

`func NewStoreIdSourceDbReq() *StoreIdSourceDbReq`

NewStoreIdSourceDbReq instantiates a new StoreIdSourceDbReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreIdSourceDbReqWithDefaults

`func NewStoreIdSourceDbReqWithDefaults() *StoreIdSourceDbReq`

NewStoreIdSourceDbReqWithDefaults instantiates a new StoreIdSourceDbReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdOrName

`func (o *StoreIdSourceDbReq) GetIdOrName() string`

GetIdOrName returns the IdOrName field if non-nil, zero value otherwise.

### GetIdOrNameOk

`func (o *StoreIdSourceDbReq) GetIdOrNameOk() (*string, bool)`

GetIdOrNameOk returns a tuple with the IdOrName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdOrName

`func (o *StoreIdSourceDbReq) SetIdOrName(v string)`

SetIdOrName sets IdOrName field to given value.

### HasIdOrName

`func (o *StoreIdSourceDbReq) HasIdOrName() bool`

HasIdOrName returns a boolean if a field has been set.

### GetIdSourceDb

`func (o *StoreIdSourceDbReq) GetIdSourceDb() DbIdentitySourceDTO`

GetIdSourceDb returns the IdSourceDb field if non-nil, zero value otherwise.

### GetIdSourceDbOk

`func (o *StoreIdSourceDbReq) GetIdSourceDbOk() (*DbIdentitySourceDTO, bool)`

GetIdSourceDbOk returns a tuple with the IdSourceDb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdSourceDb

`func (o *StoreIdSourceDbReq) SetIdSourceDb(v DbIdentitySourceDTO)`

SetIdSourceDb sets IdSourceDb field to given value.

### HasIdSourceDb

`func (o *StoreIdSourceDbReq) HasIdSourceDb() bool`

HasIdSourceDb returns a boolean if a field has been set.

### GetIdSourceLdap

`func (o *StoreIdSourceDbReq) GetIdSourceLdap() DbIdentitySourceDTO`

GetIdSourceLdap returns the IdSourceLdap field if non-nil, zero value otherwise.

### GetIdSourceLdapOk

`func (o *StoreIdSourceDbReq) GetIdSourceLdapOk() (*DbIdentitySourceDTO, bool)`

GetIdSourceLdapOk returns a tuple with the IdSourceLdap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdSourceLdap

`func (o *StoreIdSourceDbReq) SetIdSourceLdap(v DbIdentitySourceDTO)`

SetIdSourceLdap sets IdSourceLdap field to given value.

### HasIdSourceLdap

`func (o *StoreIdSourceDbReq) HasIdSourceLdap() bool`

HasIdSourceLdap returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


