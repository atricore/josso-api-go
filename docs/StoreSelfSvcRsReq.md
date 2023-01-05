# StoreSelfSvcRsReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdOrName** | Pointer to **string** |  | [optional] 
**Resource** | Pointer to [**SelfServicesResourceDTO**](SelfServicesResourceDTO.md) |  | [optional] 

## Methods

### NewStoreSelfSvcRsReq

`func NewStoreSelfSvcRsReq() *StoreSelfSvcRsReq`

NewStoreSelfSvcRsReq instantiates a new StoreSelfSvcRsReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreSelfSvcRsReqWithDefaults

`func NewStoreSelfSvcRsReqWithDefaults() *StoreSelfSvcRsReq`

NewStoreSelfSvcRsReqWithDefaults instantiates a new StoreSelfSvcRsReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdOrName

`func (o *StoreSelfSvcRsReq) GetIdOrName() string`

GetIdOrName returns the IdOrName field if non-nil, zero value otherwise.

### GetIdOrNameOk

`func (o *StoreSelfSvcRsReq) GetIdOrNameOk() (*string, bool)`

GetIdOrNameOk returns a tuple with the IdOrName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdOrName

`func (o *StoreSelfSvcRsReq) SetIdOrName(v string)`

SetIdOrName sets IdOrName field to given value.

### HasIdOrName

`func (o *StoreSelfSvcRsReq) HasIdOrName() bool`

HasIdOrName returns a boolean if a field has been set.

### GetResource

`func (o *StoreSelfSvcRsReq) GetResource() SelfServicesResourceDTO`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *StoreSelfSvcRsReq) GetResourceOk() (*SelfServicesResourceDTO, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *StoreSelfSvcRsReq) SetResource(v SelfServicesResourceDTO)`

SetResource sets Resource field to given value.

### HasResource

`func (o *StoreSelfSvcRsReq) HasResource() bool`

HasResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


