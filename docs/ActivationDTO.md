# ActivationDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**ExecutionEnv** | Pointer to [**ExecutionEnvironmentDTO**](ExecutionEnvironmentDTO.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Resource** | Pointer to [**ServiceResourceDTO**](ServiceResourceDTO.md) |  | [optional] 
**Sp** | Pointer to [**InternalSaml2ServiceProviderDTO**](InternalSaml2ServiceProviderDTO.md) |  | [optional] 
**Waypoints** | Pointer to [**[]PointDTO**](PointDTO.md) |  | [optional] 

## Methods

### NewActivationDTO

`func NewActivationDTO() *ActivationDTO`

NewActivationDTO instantiates a new ActivationDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivationDTOWithDefaults

`func NewActivationDTOWithDefaults() *ActivationDTO`

NewActivationDTOWithDefaults instantiates a new ActivationDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ActivationDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ActivationDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ActivationDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ActivationDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetElementId

`func (o *ActivationDTO) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *ActivationDTO) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *ActivationDTO) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *ActivationDTO) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetExecutionEnv

`func (o *ActivationDTO) GetExecutionEnv() ExecutionEnvironmentDTO`

GetExecutionEnv returns the ExecutionEnv field if non-nil, zero value otherwise.

### GetExecutionEnvOk

`func (o *ActivationDTO) GetExecutionEnvOk() (*ExecutionEnvironmentDTO, bool)`

GetExecutionEnvOk returns a tuple with the ExecutionEnv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionEnv

`func (o *ActivationDTO) SetExecutionEnv(v ExecutionEnvironmentDTO)`

SetExecutionEnv sets ExecutionEnv field to given value.

### HasExecutionEnv

`func (o *ActivationDTO) HasExecutionEnv() bool`

HasExecutionEnv returns a boolean if a field has been set.

### GetId

`func (o *ActivationDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActivationDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActivationDTO) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ActivationDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ActivationDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ActivationDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ActivationDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ActivationDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResource

`func (o *ActivationDTO) GetResource() ServiceResourceDTO`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ActivationDTO) GetResourceOk() (*ServiceResourceDTO, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ActivationDTO) SetResource(v ServiceResourceDTO)`

SetResource sets Resource field to given value.

### HasResource

`func (o *ActivationDTO) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetSp

`func (o *ActivationDTO) GetSp() InternalSaml2ServiceProviderDTO`

GetSp returns the Sp field if non-nil, zero value otherwise.

### GetSpOk

`func (o *ActivationDTO) GetSpOk() (*InternalSaml2ServiceProviderDTO, bool)`

GetSpOk returns a tuple with the Sp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSp

`func (o *ActivationDTO) SetSp(v InternalSaml2ServiceProviderDTO)`

SetSp sets Sp field to given value.

### HasSp

`func (o *ActivationDTO) HasSp() bool`

HasSp returns a boolean if a field has been set.

### GetWaypoints

`func (o *ActivationDTO) GetWaypoints() []PointDTO`

GetWaypoints returns the Waypoints field if non-nil, zero value otherwise.

### GetWaypointsOk

`func (o *ActivationDTO) GetWaypointsOk() (*[]PointDTO, bool)`

GetWaypointsOk returns a tuple with the Waypoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaypoints

`func (o *ActivationDTO) SetWaypoints(v []PointDTO)`

SetWaypoints sets Waypoints field to given value.

### HasWaypoints

`func (o *ActivationDTO) HasWaypoints() bool`

HasWaypoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


