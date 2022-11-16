# StoreApplianceReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Appliance** | Pointer to [**IdentityApplianceDefinitionDTO**](IdentityApplianceDefinitionDTO.md) |  | [optional] 
**IdOrName** | Pointer to **string** |  | [optional] 

## Methods

### NewStoreApplianceReq

`func NewStoreApplianceReq() *StoreApplianceReq`

NewStoreApplianceReq instantiates a new StoreApplianceReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreApplianceReqWithDefaults

`func NewStoreApplianceReqWithDefaults() *StoreApplianceReq`

NewStoreApplianceReqWithDefaults instantiates a new StoreApplianceReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppliance

`func (o *StoreApplianceReq) GetAppliance() IdentityApplianceDefinitionDTO`

GetAppliance returns the Appliance field if non-nil, zero value otherwise.

### GetApplianceOk

`func (o *StoreApplianceReq) GetApplianceOk() (*IdentityApplianceDefinitionDTO, bool)`

GetApplianceOk returns a tuple with the Appliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliance

`func (o *StoreApplianceReq) SetAppliance(v IdentityApplianceDefinitionDTO)`

SetAppliance sets Appliance field to given value.

### HasAppliance

`func (o *StoreApplianceReq) HasAppliance() bool`

HasAppliance returns a boolean if a field has been set.

### GetIdOrName

`func (o *StoreApplianceReq) GetIdOrName() string`

GetIdOrName returns the IdOrName field if non-nil, zero value otherwise.

### GetIdOrNameOk

`func (o *StoreApplianceReq) GetIdOrNameOk() (*string, bool)`

GetIdOrNameOk returns a tuple with the IdOrName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdOrName

`func (o *StoreApplianceReq) SetIdOrName(v string)`

SetIdOrName sets IdOrName field to given value.

### HasIdOrName

`func (o *StoreApplianceReq) HasIdOrName() bool`

HasIdOrName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


