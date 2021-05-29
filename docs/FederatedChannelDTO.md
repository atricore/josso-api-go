# FederatedChannelDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveBindings** | Pointer to **[]string** |  | [optional] 
**ActiveProfiles** | Pointer to **[]string** |  | [optional] 
**ConnectionA** | Pointer to [**FederatedConnectionDTO**](FederatedConnectionDTO.md) |  | [optional] 
**ConnectionB** | Pointer to [**FederatedConnectionDTO**](FederatedConnectionDTO.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Location** | Pointer to [**LocationDTO**](LocationDTO.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OverrideProviderSetup** | Pointer to **bool** |  | [optional] 

## Methods

### NewFederatedChannelDTO

`func NewFederatedChannelDTO() *FederatedChannelDTO`

NewFederatedChannelDTO instantiates a new FederatedChannelDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFederatedChannelDTOWithDefaults

`func NewFederatedChannelDTOWithDefaults() *FederatedChannelDTO`

NewFederatedChannelDTOWithDefaults instantiates a new FederatedChannelDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveBindings

`func (o *FederatedChannelDTO) GetActiveBindings() []string`

GetActiveBindings returns the ActiveBindings field if non-nil, zero value otherwise.

### GetActiveBindingsOk

`func (o *FederatedChannelDTO) GetActiveBindingsOk() (*[]string, bool)`

GetActiveBindingsOk returns a tuple with the ActiveBindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveBindings

`func (o *FederatedChannelDTO) SetActiveBindings(v []string)`

SetActiveBindings sets ActiveBindings field to given value.

### HasActiveBindings

`func (o *FederatedChannelDTO) HasActiveBindings() bool`

HasActiveBindings returns a boolean if a field has been set.

### GetActiveProfiles

`func (o *FederatedChannelDTO) GetActiveProfiles() []string`

GetActiveProfiles returns the ActiveProfiles field if non-nil, zero value otherwise.

### GetActiveProfilesOk

`func (o *FederatedChannelDTO) GetActiveProfilesOk() (*[]string, bool)`

GetActiveProfilesOk returns a tuple with the ActiveProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveProfiles

`func (o *FederatedChannelDTO) SetActiveProfiles(v []string)`

SetActiveProfiles sets ActiveProfiles field to given value.

### HasActiveProfiles

`func (o *FederatedChannelDTO) HasActiveProfiles() bool`

HasActiveProfiles returns a boolean if a field has been set.

### GetConnectionA

`func (o *FederatedChannelDTO) GetConnectionA() FederatedConnectionDTO`

GetConnectionA returns the ConnectionA field if non-nil, zero value otherwise.

### GetConnectionAOk

`func (o *FederatedChannelDTO) GetConnectionAOk() (*FederatedConnectionDTO, bool)`

GetConnectionAOk returns a tuple with the ConnectionA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionA

`func (o *FederatedChannelDTO) SetConnectionA(v FederatedConnectionDTO)`

SetConnectionA sets ConnectionA field to given value.

### HasConnectionA

`func (o *FederatedChannelDTO) HasConnectionA() bool`

HasConnectionA returns a boolean if a field has been set.

### GetConnectionB

`func (o *FederatedChannelDTO) GetConnectionB() FederatedConnectionDTO`

GetConnectionB returns the ConnectionB field if non-nil, zero value otherwise.

### GetConnectionBOk

`func (o *FederatedChannelDTO) GetConnectionBOk() (*FederatedConnectionDTO, bool)`

GetConnectionBOk returns a tuple with the ConnectionB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionB

`func (o *FederatedChannelDTO) SetConnectionB(v FederatedConnectionDTO)`

SetConnectionB sets ConnectionB field to given value.

### HasConnectionB

`func (o *FederatedChannelDTO) HasConnectionB() bool`

HasConnectionB returns a boolean if a field has been set.

### GetDescription

`func (o *FederatedChannelDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FederatedChannelDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FederatedChannelDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FederatedChannelDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *FederatedChannelDTO) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *FederatedChannelDTO) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *FederatedChannelDTO) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *FederatedChannelDTO) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetElementId

`func (o *FederatedChannelDTO) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *FederatedChannelDTO) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *FederatedChannelDTO) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *FederatedChannelDTO) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetId

`func (o *FederatedChannelDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FederatedChannelDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FederatedChannelDTO) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *FederatedChannelDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocation

`func (o *FederatedChannelDTO) GetLocation() LocationDTO`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *FederatedChannelDTO) GetLocationOk() (*LocationDTO, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *FederatedChannelDTO) SetLocation(v LocationDTO)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *FederatedChannelDTO) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetName

`func (o *FederatedChannelDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FederatedChannelDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FederatedChannelDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FederatedChannelDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOverrideProviderSetup

`func (o *FederatedChannelDTO) GetOverrideProviderSetup() bool`

GetOverrideProviderSetup returns the OverrideProviderSetup field if non-nil, zero value otherwise.

### GetOverrideProviderSetupOk

`func (o *FederatedChannelDTO) GetOverrideProviderSetupOk() (*bool, bool)`

GetOverrideProviderSetupOk returns a tuple with the OverrideProviderSetup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideProviderSetup

`func (o *FederatedChannelDTO) SetOverrideProviderSetup(v bool)`

SetOverrideProviderSetup sets OverrideProviderSetup field to given value.

### HasOverrideProviderSetup

`func (o *FederatedChannelDTO) HasOverrideProviderSetup() bool`

HasOverrideProviderSetup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


