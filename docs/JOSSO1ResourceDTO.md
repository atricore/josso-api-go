# JOSSO1ResourceDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activation** | Pointer to [**ActivationDTO**](ActivationDTO.md) |  | [optional] 
**DefaultResource** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**IgnoredWebResources** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PartnerAppLocation** | Pointer to [**LocationDTO**](LocationDTO.md) |  | [optional] 
**ServiceConnection** | Pointer to [**ServiceConnectionDTO**](ServiceConnectionDTO.md) |  | [optional] 
**SloLocation** | Pointer to [**LocationDTO**](LocationDTO.md) |  | [optional] 
**SloLocationEnabled** | Pointer to **bool** |  | [optional] 
**X** | Pointer to **float64** |  | [optional] 
**Y** | Pointer to **float64** |  | [optional] 

## Methods

### NewJOSSO1ResourceDTO

`func NewJOSSO1ResourceDTO() *JOSSO1ResourceDTO`

NewJOSSO1ResourceDTO instantiates a new JOSSO1ResourceDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJOSSO1ResourceDTOWithDefaults

`func NewJOSSO1ResourceDTOWithDefaults() *JOSSO1ResourceDTO`

NewJOSSO1ResourceDTOWithDefaults instantiates a new JOSSO1ResourceDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivation

`func (o *JOSSO1ResourceDTO) GetActivation() ActivationDTO`

GetActivation returns the Activation field if non-nil, zero value otherwise.

### GetActivationOk

`func (o *JOSSO1ResourceDTO) GetActivationOk() (*ActivationDTO, bool)`

GetActivationOk returns a tuple with the Activation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivation

`func (o *JOSSO1ResourceDTO) SetActivation(v ActivationDTO)`

SetActivation sets Activation field to given value.

### HasActivation

`func (o *JOSSO1ResourceDTO) HasActivation() bool`

HasActivation returns a boolean if a field has been set.

### GetDefaultResource

`func (o *JOSSO1ResourceDTO) GetDefaultResource() string`

GetDefaultResource returns the DefaultResource field if non-nil, zero value otherwise.

### GetDefaultResourceOk

`func (o *JOSSO1ResourceDTO) GetDefaultResourceOk() (*string, bool)`

GetDefaultResourceOk returns a tuple with the DefaultResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultResource

`func (o *JOSSO1ResourceDTO) SetDefaultResource(v string)`

SetDefaultResource sets DefaultResource field to given value.

### HasDefaultResource

`func (o *JOSSO1ResourceDTO) HasDefaultResource() bool`

HasDefaultResource returns a boolean if a field has been set.

### GetDescription

`func (o *JOSSO1ResourceDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JOSSO1ResourceDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JOSSO1ResourceDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *JOSSO1ResourceDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetElementId

`func (o *JOSSO1ResourceDTO) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *JOSSO1ResourceDTO) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *JOSSO1ResourceDTO) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *JOSSO1ResourceDTO) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetId

`func (o *JOSSO1ResourceDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JOSSO1ResourceDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JOSSO1ResourceDTO) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *JOSSO1ResourceDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIgnoredWebResources

`func (o *JOSSO1ResourceDTO) GetIgnoredWebResources() []string`

GetIgnoredWebResources returns the IgnoredWebResources field if non-nil, zero value otherwise.

### GetIgnoredWebResourcesOk

`func (o *JOSSO1ResourceDTO) GetIgnoredWebResourcesOk() (*[]string, bool)`

GetIgnoredWebResourcesOk returns a tuple with the IgnoredWebResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoredWebResources

`func (o *JOSSO1ResourceDTO) SetIgnoredWebResources(v []string)`

SetIgnoredWebResources sets IgnoredWebResources field to given value.

### HasIgnoredWebResources

`func (o *JOSSO1ResourceDTO) HasIgnoredWebResources() bool`

HasIgnoredWebResources returns a boolean if a field has been set.

### GetName

`func (o *JOSSO1ResourceDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JOSSO1ResourceDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JOSSO1ResourceDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *JOSSO1ResourceDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPartnerAppLocation

`func (o *JOSSO1ResourceDTO) GetPartnerAppLocation() LocationDTO`

GetPartnerAppLocation returns the PartnerAppLocation field if non-nil, zero value otherwise.

### GetPartnerAppLocationOk

`func (o *JOSSO1ResourceDTO) GetPartnerAppLocationOk() (*LocationDTO, bool)`

GetPartnerAppLocationOk returns a tuple with the PartnerAppLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerAppLocation

`func (o *JOSSO1ResourceDTO) SetPartnerAppLocation(v LocationDTO)`

SetPartnerAppLocation sets PartnerAppLocation field to given value.

### HasPartnerAppLocation

`func (o *JOSSO1ResourceDTO) HasPartnerAppLocation() bool`

HasPartnerAppLocation returns a boolean if a field has been set.

### GetServiceConnection

`func (o *JOSSO1ResourceDTO) GetServiceConnection() ServiceConnectionDTO`

GetServiceConnection returns the ServiceConnection field if non-nil, zero value otherwise.

### GetServiceConnectionOk

`func (o *JOSSO1ResourceDTO) GetServiceConnectionOk() (*ServiceConnectionDTO, bool)`

GetServiceConnectionOk returns a tuple with the ServiceConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceConnection

`func (o *JOSSO1ResourceDTO) SetServiceConnection(v ServiceConnectionDTO)`

SetServiceConnection sets ServiceConnection field to given value.

### HasServiceConnection

`func (o *JOSSO1ResourceDTO) HasServiceConnection() bool`

HasServiceConnection returns a boolean if a field has been set.

### GetSloLocation

`func (o *JOSSO1ResourceDTO) GetSloLocation() LocationDTO`

GetSloLocation returns the SloLocation field if non-nil, zero value otherwise.

### GetSloLocationOk

`func (o *JOSSO1ResourceDTO) GetSloLocationOk() (*LocationDTO, bool)`

GetSloLocationOk returns a tuple with the SloLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSloLocation

`func (o *JOSSO1ResourceDTO) SetSloLocation(v LocationDTO)`

SetSloLocation sets SloLocation field to given value.

### HasSloLocation

`func (o *JOSSO1ResourceDTO) HasSloLocation() bool`

HasSloLocation returns a boolean if a field has been set.

### GetSloLocationEnabled

`func (o *JOSSO1ResourceDTO) GetSloLocationEnabled() bool`

GetSloLocationEnabled returns the SloLocationEnabled field if non-nil, zero value otherwise.

### GetSloLocationEnabledOk

`func (o *JOSSO1ResourceDTO) GetSloLocationEnabledOk() (*bool, bool)`

GetSloLocationEnabledOk returns a tuple with the SloLocationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSloLocationEnabled

`func (o *JOSSO1ResourceDTO) SetSloLocationEnabled(v bool)`

SetSloLocationEnabled sets SloLocationEnabled field to given value.

### HasSloLocationEnabled

`func (o *JOSSO1ResourceDTO) HasSloLocationEnabled() bool`

HasSloLocationEnabled returns a boolean if a field has been set.

### GetX

`func (o *JOSSO1ResourceDTO) GetX() float64`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *JOSSO1ResourceDTO) GetXOk() (*float64, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *JOSSO1ResourceDTO) SetX(v float64)`

SetX sets X field to given value.

### HasX

`func (o *JOSSO1ResourceDTO) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *JOSSO1ResourceDTO) GetY() float64`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *JOSSO1ResourceDTO) GetYOk() (*float64, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *JOSSO1ResourceDTO) SetY(v float64)`

SetY sets Y field to given value.

### HasY

`func (o *JOSSO1ResourceDTO) HasY() bool`

HasY returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


