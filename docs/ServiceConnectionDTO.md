# ServiceConnectionDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Resource** | Pointer to [**ServiceResourceDTO**](ServiceResourceDTO.md) |  | [optional] 
**Sp** | Pointer to [**InternalSaml2ServiceProviderDTO**](InternalSaml2ServiceProviderDTO.md) |  | [optional] 
**Waypoints** | Pointer to [**[]PointDTO**](PointDTO.md) |  | [optional] 

## Methods

### NewServiceConnectionDTO

`func NewServiceConnectionDTO() *ServiceConnectionDTO`

NewServiceConnectionDTO instantiates a new ServiceConnectionDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceConnectionDTOWithDefaults

`func NewServiceConnectionDTOWithDefaults() *ServiceConnectionDTO`

NewServiceConnectionDTOWithDefaults instantiates a new ServiceConnectionDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ServiceConnectionDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceConnectionDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceConnectionDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServiceConnectionDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetElementId

`func (o *ServiceConnectionDTO) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *ServiceConnectionDTO) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *ServiceConnectionDTO) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *ServiceConnectionDTO) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetId

`func (o *ServiceConnectionDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceConnectionDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceConnectionDTO) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceConnectionDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ServiceConnectionDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceConnectionDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceConnectionDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceConnectionDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResource

`func (o *ServiceConnectionDTO) GetResource() ServiceResourceDTO`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ServiceConnectionDTO) GetResourceOk() (*ServiceResourceDTO, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ServiceConnectionDTO) SetResource(v ServiceResourceDTO)`

SetResource sets Resource field to given value.

### HasResource

`func (o *ServiceConnectionDTO) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetSp

`func (o *ServiceConnectionDTO) GetSp() InternalSaml2ServiceProviderDTO`

GetSp returns the Sp field if non-nil, zero value otherwise.

### GetSpOk

`func (o *ServiceConnectionDTO) GetSpOk() (*InternalSaml2ServiceProviderDTO, bool)`

GetSpOk returns a tuple with the Sp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSp

`func (o *ServiceConnectionDTO) SetSp(v InternalSaml2ServiceProviderDTO)`

SetSp sets Sp field to given value.

### HasSp

`func (o *ServiceConnectionDTO) HasSp() bool`

HasSp returns a boolean if a field has been set.

### GetWaypoints

`func (o *ServiceConnectionDTO) GetWaypoints() []PointDTO`

GetWaypoints returns the Waypoints field if non-nil, zero value otherwise.

### GetWaypointsOk

`func (o *ServiceConnectionDTO) GetWaypointsOk() (*[]PointDTO, bool)`

GetWaypointsOk returns a tuple with the Waypoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaypoints

`func (o *ServiceConnectionDTO) SetWaypoints(v []PointDTO)`

SetWaypoints sets Waypoints field to given value.

### HasWaypoints

`func (o *ServiceConnectionDTO) HasWaypoints() bool`

HasWaypoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


