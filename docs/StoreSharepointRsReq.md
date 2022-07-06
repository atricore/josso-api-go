# StoreSharepointRsReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdaName** | Pointer to **string** |  | [optional] 
**Resource** | Pointer to [**SharepointResourceDTO**](SharepointResourceDTO.md) |  | [optional] 

## Methods

### NewStoreSharepointRsReq

`func NewStoreSharepointRsReq() *StoreSharepointRsReq`

NewStoreSharepointRsReq instantiates a new StoreSharepointRsReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreSharepointRsReqWithDefaults

`func NewStoreSharepointRsReqWithDefaults() *StoreSharepointRsReq`

NewStoreSharepointRsReqWithDefaults instantiates a new StoreSharepointRsReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdaName

`func (o *StoreSharepointRsReq) GetIdaName() string`

GetIdaName returns the IdaName field if non-nil, zero value otherwise.

### GetIdaNameOk

`func (o *StoreSharepointRsReq) GetIdaNameOk() (*string, bool)`

GetIdaNameOk returns a tuple with the IdaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdaName

`func (o *StoreSharepointRsReq) SetIdaName(v string)`

SetIdaName sets IdaName field to given value.

### HasIdaName

`func (o *StoreSharepointRsReq) HasIdaName() bool`

HasIdaName returns a boolean if a field has been set.

### GetResource

`func (o *StoreSharepointRsReq) GetResource() SharepointResourceDTO`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *StoreSharepointRsReq) GetResourceOk() (*SharepointResourceDTO, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *StoreSharepointRsReq) SetResource(v SharepointResourceDTO)`

SetResource sets Resource field to given value.

### HasResource

`func (o *StoreSharepointRsReq) HasResource() bool`

HasResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


