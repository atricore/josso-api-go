# StoreIntSaml2SpReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdOrName** | Pointer to **string** |  | [optional] 
**Sp** | Pointer to [**InternalSaml2ServiceProviderDTO**](InternalSaml2ServiceProviderDTO.md) |  | [optional] 

## Methods

### NewStoreIntSaml2SpReq

`func NewStoreIntSaml2SpReq() *StoreIntSaml2SpReq`

NewStoreIntSaml2SpReq instantiates a new StoreIntSaml2SpReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreIntSaml2SpReqWithDefaults

`func NewStoreIntSaml2SpReqWithDefaults() *StoreIntSaml2SpReq`

NewStoreIntSaml2SpReqWithDefaults instantiates a new StoreIntSaml2SpReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdOrName

`func (o *StoreIntSaml2SpReq) GetIdOrName() string`

GetIdOrName returns the IdOrName field if non-nil, zero value otherwise.

### GetIdOrNameOk

`func (o *StoreIntSaml2SpReq) GetIdOrNameOk() (*string, bool)`

GetIdOrNameOk returns a tuple with the IdOrName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdOrName

`func (o *StoreIntSaml2SpReq) SetIdOrName(v string)`

SetIdOrName sets IdOrName field to given value.

### HasIdOrName

`func (o *StoreIntSaml2SpReq) HasIdOrName() bool`

HasIdOrName returns a boolean if a field has been set.

### GetSp

`func (o *StoreIntSaml2SpReq) GetSp() InternalSaml2ServiceProviderDTO`

GetSp returns the Sp field if non-nil, zero value otherwise.

### GetSpOk

`func (o *StoreIntSaml2SpReq) GetSpOk() (*InternalSaml2ServiceProviderDTO, bool)`

GetSpOk returns a tuple with the Sp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSp

`func (o *StoreIntSaml2SpReq) SetSp(v InternalSaml2ServiceProviderDTO)`

SetSp sets Sp field to given value.

### HasSp

`func (o *StoreIntSaml2SpReq) HasSp() bool`

HasSp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


