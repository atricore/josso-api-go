# StoreJossoRsReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdaName** | Pointer to **string** |  | [optional] 
**Resource** | Pointer to [**JOSSO1ResourceDTO**](JOSSO1ResourceDTO.md) |  | [optional] 

## Methods

### NewStoreJossoRsReq

`func NewStoreJossoRsReq() *StoreJossoRsReq`

NewStoreJossoRsReq instantiates a new StoreJossoRsReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreJossoRsReqWithDefaults

`func NewStoreJossoRsReqWithDefaults() *StoreJossoRsReq`

NewStoreJossoRsReqWithDefaults instantiates a new StoreJossoRsReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdaName

`func (o *StoreJossoRsReq) GetIdaName() string`

GetIdaName returns the IdaName field if non-nil, zero value otherwise.

### GetIdaNameOk

`func (o *StoreJossoRsReq) GetIdaNameOk() (*string, bool)`

GetIdaNameOk returns a tuple with the IdaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdaName

`func (o *StoreJossoRsReq) SetIdaName(v string)`

SetIdaName sets IdaName field to given value.

### HasIdaName

`func (o *StoreJossoRsReq) HasIdaName() bool`

HasIdaName returns a boolean if a field has been set.

### GetResource

`func (o *StoreJossoRsReq) GetResource() JOSSO1ResourceDTO`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *StoreJossoRsReq) GetResourceOk() (*JOSSO1ResourceDTO, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *StoreJossoRsReq) SetResource(v JOSSO1ResourceDTO)`

SetResource sets Resource field to given value.

### HasResource

`func (o *StoreJossoRsReq) HasResource() bool`

HasResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


