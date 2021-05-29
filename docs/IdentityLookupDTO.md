# IdentityLookupDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**IdentitySource** | Pointer to [**IdentitySourceDTO**](IdentitySourceDTO.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to [**ProviderDTO**](ProviderDTO.md) |  | [optional] 
**Waypoints** | Pointer to [**[]PointDTO**](PointDTO.md) |  | [optional] 

## Methods

### NewIdentityLookupDTO

`func NewIdentityLookupDTO() *IdentityLookupDTO`

NewIdentityLookupDTO instantiates a new IdentityLookupDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityLookupDTOWithDefaults

`func NewIdentityLookupDTOWithDefaults() *IdentityLookupDTO`

NewIdentityLookupDTOWithDefaults instantiates a new IdentityLookupDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *IdentityLookupDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IdentityLookupDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IdentityLookupDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IdentityLookupDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetElementId

`func (o *IdentityLookupDTO) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *IdentityLookupDTO) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *IdentityLookupDTO) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *IdentityLookupDTO) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetId

`func (o *IdentityLookupDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentityLookupDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentityLookupDTO) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *IdentityLookupDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentitySource

`func (o *IdentityLookupDTO) GetIdentitySource() IdentitySourceDTO`

GetIdentitySource returns the IdentitySource field if non-nil, zero value otherwise.

### GetIdentitySourceOk

`func (o *IdentityLookupDTO) GetIdentitySourceOk() (*IdentitySourceDTO, bool)`

GetIdentitySourceOk returns a tuple with the IdentitySource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentitySource

`func (o *IdentityLookupDTO) SetIdentitySource(v IdentitySourceDTO)`

SetIdentitySource sets IdentitySource field to given value.

### HasIdentitySource

`func (o *IdentityLookupDTO) HasIdentitySource() bool`

HasIdentitySource returns a boolean if a field has been set.

### GetName

`func (o *IdentityLookupDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdentityLookupDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdentityLookupDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IdentityLookupDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProvider

`func (o *IdentityLookupDTO) GetProvider() ProviderDTO`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *IdentityLookupDTO) GetProviderOk() (*ProviderDTO, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *IdentityLookupDTO) SetProvider(v ProviderDTO)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *IdentityLookupDTO) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetWaypoints

`func (o *IdentityLookupDTO) GetWaypoints() []PointDTO`

GetWaypoints returns the Waypoints field if non-nil, zero value otherwise.

### GetWaypointsOk

`func (o *IdentityLookupDTO) GetWaypointsOk() (*[]PointDTO, bool)`

GetWaypointsOk returns a tuple with the Waypoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaypoints

`func (o *IdentityLookupDTO) SetWaypoints(v []PointDTO)`

SetWaypoints sets Waypoints field to given value.

### HasWaypoints

`func (o *IdentityLookupDTO) HasWaypoints() bool`

HasWaypoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


