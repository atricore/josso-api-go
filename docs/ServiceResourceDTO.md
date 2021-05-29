# ServiceResourceDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activation** | Pointer to [**ActivationDTO**](ActivationDTO.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ServiceConnection** | Pointer to [**ServiceConnectionDTO**](ServiceConnectionDTO.md) |  | [optional] 
**X** | Pointer to **float64** |  | [optional] 
**Y** | Pointer to **float64** |  | [optional] 

## Methods

### NewServiceResourceDTO

`func NewServiceResourceDTO() *ServiceResourceDTO`

NewServiceResourceDTO instantiates a new ServiceResourceDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceResourceDTOWithDefaults

`func NewServiceResourceDTOWithDefaults() *ServiceResourceDTO`

NewServiceResourceDTOWithDefaults instantiates a new ServiceResourceDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivation

`func (o *ServiceResourceDTO) GetActivation() ActivationDTO`

GetActivation returns the Activation field if non-nil, zero value otherwise.

### GetActivationOk

`func (o *ServiceResourceDTO) GetActivationOk() (*ActivationDTO, bool)`

GetActivationOk returns a tuple with the Activation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivation

`func (o *ServiceResourceDTO) SetActivation(v ActivationDTO)`

SetActivation sets Activation field to given value.

### HasActivation

`func (o *ServiceResourceDTO) HasActivation() bool`

HasActivation returns a boolean if a field has been set.

### GetDescription

`func (o *ServiceResourceDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceResourceDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceResourceDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServiceResourceDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetElementId

`func (o *ServiceResourceDTO) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *ServiceResourceDTO) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *ServiceResourceDTO) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *ServiceResourceDTO) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetId

`func (o *ServiceResourceDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceResourceDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceResourceDTO) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceResourceDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ServiceResourceDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceResourceDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceResourceDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceResourceDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetServiceConnection

`func (o *ServiceResourceDTO) GetServiceConnection() ServiceConnectionDTO`

GetServiceConnection returns the ServiceConnection field if non-nil, zero value otherwise.

### GetServiceConnectionOk

`func (o *ServiceResourceDTO) GetServiceConnectionOk() (*ServiceConnectionDTO, bool)`

GetServiceConnectionOk returns a tuple with the ServiceConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceConnection

`func (o *ServiceResourceDTO) SetServiceConnection(v ServiceConnectionDTO)`

SetServiceConnection sets ServiceConnection field to given value.

### HasServiceConnection

`func (o *ServiceResourceDTO) HasServiceConnection() bool`

HasServiceConnection returns a boolean if a field has been set.

### GetX

`func (o *ServiceResourceDTO) GetX() float64`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *ServiceResourceDTO) GetXOk() (*float64, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *ServiceResourceDTO) SetX(v float64)`

SetX sets X field to given value.

### HasX

`func (o *ServiceResourceDTO) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *ServiceResourceDTO) GetY() float64`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *ServiceResourceDTO) GetYOk() (*float64, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *ServiceResourceDTO) SetY(v float64)`

SetY sets Y field to given value.

### HasY

`func (o *ServiceResourceDTO) HasY() bool`

HasY returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


