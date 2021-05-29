# FederatedConnectionDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelA** | Pointer to [**FederatedChannelDTO**](FederatedChannelDTO.md) |  | [optional] 
**ChannelB** | Pointer to [**FederatedChannelDTO**](FederatedChannelDTO.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**RoleA** | Pointer to [**FederatedProviderDTO**](FederatedProviderDTO.md) |  | [optional] 
**RoleB** | Pointer to [**FederatedProviderDTO**](FederatedProviderDTO.md) |  | [optional] 
**Waypoints** | Pointer to [**[]PointDTO**](PointDTO.md) |  | [optional] 

## Methods

### NewFederatedConnectionDTO

`func NewFederatedConnectionDTO() *FederatedConnectionDTO`

NewFederatedConnectionDTO instantiates a new FederatedConnectionDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFederatedConnectionDTOWithDefaults

`func NewFederatedConnectionDTOWithDefaults() *FederatedConnectionDTO`

NewFederatedConnectionDTOWithDefaults instantiates a new FederatedConnectionDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelA

`func (o *FederatedConnectionDTO) GetChannelA() FederatedChannelDTO`

GetChannelA returns the ChannelA field if non-nil, zero value otherwise.

### GetChannelAOk

`func (o *FederatedConnectionDTO) GetChannelAOk() (*FederatedChannelDTO, bool)`

GetChannelAOk returns a tuple with the ChannelA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelA

`func (o *FederatedConnectionDTO) SetChannelA(v FederatedChannelDTO)`

SetChannelA sets ChannelA field to given value.

### HasChannelA

`func (o *FederatedConnectionDTO) HasChannelA() bool`

HasChannelA returns a boolean if a field has been set.

### GetChannelB

`func (o *FederatedConnectionDTO) GetChannelB() FederatedChannelDTO`

GetChannelB returns the ChannelB field if non-nil, zero value otherwise.

### GetChannelBOk

`func (o *FederatedConnectionDTO) GetChannelBOk() (*FederatedChannelDTO, bool)`

GetChannelBOk returns a tuple with the ChannelB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelB

`func (o *FederatedConnectionDTO) SetChannelB(v FederatedChannelDTO)`

SetChannelB sets ChannelB field to given value.

### HasChannelB

`func (o *FederatedConnectionDTO) HasChannelB() bool`

HasChannelB returns a boolean if a field has been set.

### GetDescription

`func (o *FederatedConnectionDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FederatedConnectionDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FederatedConnectionDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FederatedConnectionDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetElementId

`func (o *FederatedConnectionDTO) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *FederatedConnectionDTO) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *FederatedConnectionDTO) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *FederatedConnectionDTO) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetId

`func (o *FederatedConnectionDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FederatedConnectionDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FederatedConnectionDTO) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *FederatedConnectionDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *FederatedConnectionDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FederatedConnectionDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FederatedConnectionDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FederatedConnectionDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRoleA

`func (o *FederatedConnectionDTO) GetRoleA() FederatedProviderDTO`

GetRoleA returns the RoleA field if non-nil, zero value otherwise.

### GetRoleAOk

`func (o *FederatedConnectionDTO) GetRoleAOk() (*FederatedProviderDTO, bool)`

GetRoleAOk returns a tuple with the RoleA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleA

`func (o *FederatedConnectionDTO) SetRoleA(v FederatedProviderDTO)`

SetRoleA sets RoleA field to given value.

### HasRoleA

`func (o *FederatedConnectionDTO) HasRoleA() bool`

HasRoleA returns a boolean if a field has been set.

### GetRoleB

`func (o *FederatedConnectionDTO) GetRoleB() FederatedProviderDTO`

GetRoleB returns the RoleB field if non-nil, zero value otherwise.

### GetRoleBOk

`func (o *FederatedConnectionDTO) GetRoleBOk() (*FederatedProviderDTO, bool)`

GetRoleBOk returns a tuple with the RoleB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleB

`func (o *FederatedConnectionDTO) SetRoleB(v FederatedProviderDTO)`

SetRoleB sets RoleB field to given value.

### HasRoleB

`func (o *FederatedConnectionDTO) HasRoleB() bool`

HasRoleB returns a boolean if a field has been set.

### GetWaypoints

`func (o *FederatedConnectionDTO) GetWaypoints() []PointDTO`

GetWaypoints returns the Waypoints field if non-nil, zero value otherwise.

### GetWaypointsOk

`func (o *FederatedConnectionDTO) GetWaypointsOk() (*[]PointDTO, bool)`

GetWaypointsOk returns a tuple with the Waypoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaypoints

`func (o *FederatedConnectionDTO) SetWaypoints(v []PointDTO)`

SetWaypoints sets Waypoints field to given value.

### HasWaypoints

`func (o *FederatedConnectionDTO) HasWaypoints() bool`

HasWaypoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


