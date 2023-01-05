# SelfServicesResourceDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activation** | Pointer to [**ActivationDTO**](ActivationDTO.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Location** | Pointer to [**LocationDTO**](LocationDTO.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Secret** | Pointer to **string** |  | [optional] 
**ServiceConnection** | Pointer to [**ServiceConnectionDTO**](ServiceConnectionDTO.md) |  | [optional] 
**X** | Pointer to **float64** |  | [optional] 
**Y** | Pointer to **float64** |  | [optional] 

## Methods

### NewSelfServicesResourceDTO

`func NewSelfServicesResourceDTO() *SelfServicesResourceDTO`

NewSelfServicesResourceDTO instantiates a new SelfServicesResourceDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelfServicesResourceDTOWithDefaults

`func NewSelfServicesResourceDTOWithDefaults() *SelfServicesResourceDTO`

NewSelfServicesResourceDTOWithDefaults instantiates a new SelfServicesResourceDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivation

`func (o *SelfServicesResourceDTO) GetActivation() ActivationDTO`

GetActivation returns the Activation field if non-nil, zero value otherwise.

### GetActivationOk

`func (o *SelfServicesResourceDTO) GetActivationOk() (*ActivationDTO, bool)`

GetActivationOk returns a tuple with the Activation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivation

`func (o *SelfServicesResourceDTO) SetActivation(v ActivationDTO)`

SetActivation sets Activation field to given value.

### HasActivation

`func (o *SelfServicesResourceDTO) HasActivation() bool`

HasActivation returns a boolean if a field has been set.

### GetDescription

`func (o *SelfServicesResourceDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SelfServicesResourceDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SelfServicesResourceDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SelfServicesResourceDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetElementId

`func (o *SelfServicesResourceDTO) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *SelfServicesResourceDTO) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *SelfServicesResourceDTO) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *SelfServicesResourceDTO) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetId

`func (o *SelfServicesResourceDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SelfServicesResourceDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SelfServicesResourceDTO) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *SelfServicesResourceDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocation

`func (o *SelfServicesResourceDTO) GetLocation() LocationDTO`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SelfServicesResourceDTO) GetLocationOk() (*LocationDTO, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SelfServicesResourceDTO) SetLocation(v LocationDTO)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *SelfServicesResourceDTO) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetName

`func (o *SelfServicesResourceDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SelfServicesResourceDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SelfServicesResourceDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SelfServicesResourceDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSecret

`func (o *SelfServicesResourceDTO) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *SelfServicesResourceDTO) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *SelfServicesResourceDTO) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *SelfServicesResourceDTO) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetServiceConnection

`func (o *SelfServicesResourceDTO) GetServiceConnection() ServiceConnectionDTO`

GetServiceConnection returns the ServiceConnection field if non-nil, zero value otherwise.

### GetServiceConnectionOk

`func (o *SelfServicesResourceDTO) GetServiceConnectionOk() (*ServiceConnectionDTO, bool)`

GetServiceConnectionOk returns a tuple with the ServiceConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceConnection

`func (o *SelfServicesResourceDTO) SetServiceConnection(v ServiceConnectionDTO)`

SetServiceConnection sets ServiceConnection field to given value.

### HasServiceConnection

`func (o *SelfServicesResourceDTO) HasServiceConnection() bool`

HasServiceConnection returns a boolean if a field has been set.

### GetX

`func (o *SelfServicesResourceDTO) GetX() float64`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *SelfServicesResourceDTO) GetXOk() (*float64, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *SelfServicesResourceDTO) SetX(v float64)`

SetX sets X field to given value.

### HasX

`func (o *SelfServicesResourceDTO) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *SelfServicesResourceDTO) GetY() float64`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *SelfServicesResourceDTO) GetYOk() (*float64, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *SelfServicesResourceDTO) SetY(v float64)`

SetY sets Y field to given value.

### HasY

`func (o *SelfServicesResourceDTO) HasY() bool`

HasY returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


