# DelegatedAuthenticationDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthnService** | Pointer to [**AuthenticationServiceDTO**](AuthenticationServiceDTO.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Idp** | Pointer to [**IdentityProviderDTO**](IdentityProviderDTO.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Waypoints** | Pointer to [**[]PointDTO**](PointDTO.md) |  | [optional] 

## Methods

### NewDelegatedAuthenticationDTO

`func NewDelegatedAuthenticationDTO() *DelegatedAuthenticationDTO`

NewDelegatedAuthenticationDTO instantiates a new DelegatedAuthenticationDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDelegatedAuthenticationDTOWithDefaults

`func NewDelegatedAuthenticationDTOWithDefaults() *DelegatedAuthenticationDTO`

NewDelegatedAuthenticationDTOWithDefaults instantiates a new DelegatedAuthenticationDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthnService

`func (o *DelegatedAuthenticationDTO) GetAuthnService() AuthenticationServiceDTO`

GetAuthnService returns the AuthnService field if non-nil, zero value otherwise.

### GetAuthnServiceOk

`func (o *DelegatedAuthenticationDTO) GetAuthnServiceOk() (*AuthenticationServiceDTO, bool)`

GetAuthnServiceOk returns a tuple with the AuthnService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthnService

`func (o *DelegatedAuthenticationDTO) SetAuthnService(v AuthenticationServiceDTO)`

SetAuthnService sets AuthnService field to given value.

### HasAuthnService

`func (o *DelegatedAuthenticationDTO) HasAuthnService() bool`

HasAuthnService returns a boolean if a field has been set.

### GetDescription

`func (o *DelegatedAuthenticationDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DelegatedAuthenticationDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DelegatedAuthenticationDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DelegatedAuthenticationDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetElementId

`func (o *DelegatedAuthenticationDTO) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *DelegatedAuthenticationDTO) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *DelegatedAuthenticationDTO) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *DelegatedAuthenticationDTO) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetId

`func (o *DelegatedAuthenticationDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DelegatedAuthenticationDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DelegatedAuthenticationDTO) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DelegatedAuthenticationDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdp

`func (o *DelegatedAuthenticationDTO) GetIdp() IdentityProviderDTO`

GetIdp returns the Idp field if non-nil, zero value otherwise.

### GetIdpOk

`func (o *DelegatedAuthenticationDTO) GetIdpOk() (*IdentityProviderDTO, bool)`

GetIdpOk returns a tuple with the Idp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdp

`func (o *DelegatedAuthenticationDTO) SetIdp(v IdentityProviderDTO)`

SetIdp sets Idp field to given value.

### HasIdp

`func (o *DelegatedAuthenticationDTO) HasIdp() bool`

HasIdp returns a boolean if a field has been set.

### GetName

`func (o *DelegatedAuthenticationDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DelegatedAuthenticationDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DelegatedAuthenticationDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DelegatedAuthenticationDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetWaypoints

`func (o *DelegatedAuthenticationDTO) GetWaypoints() []PointDTO`

GetWaypoints returns the Waypoints field if non-nil, zero value otherwise.

### GetWaypointsOk

`func (o *DelegatedAuthenticationDTO) GetWaypointsOk() (*[]PointDTO, bool)`

GetWaypointsOk returns a tuple with the Waypoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaypoints

`func (o *DelegatedAuthenticationDTO) SetWaypoints(v []PointDTO)`

SetWaypoints sets Waypoints field to given value.

### HasWaypoints

`func (o *DelegatedAuthenticationDTO) HasWaypoints() bool`

HasWaypoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


