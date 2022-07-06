# SharepointResourceDTO

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
**RelayingPartyLocation** | Pointer to [**LocationDTO**](LocationDTO.md) |  | [optional] 
**ServiceConnection** | Pointer to [**ServiceConnectionDTO**](ServiceConnectionDTO.md) |  | [optional] 
**SloLocation** | Pointer to [**LocationDTO**](LocationDTO.md) |  | [optional] 
**SloLocationEnabled** | Pointer to **bool** |  | [optional] 
**StsEncryptingCertSubject** | Pointer to **string** |  | [optional] 
**StsMetadata** | Pointer to [**ResourceDTO**](ResourceDTO.md) |  | [optional] 
**StsSigningCertSubject** | Pointer to **string** |  | [optional] 
**X** | Pointer to **float64** |  | [optional] 
**Y** | Pointer to **float64** |  | [optional] 

## Methods

### NewSharepointResourceDTO

`func NewSharepointResourceDTO() *SharepointResourceDTO`

NewSharepointResourceDTO instantiates a new SharepointResourceDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharepointResourceDTOWithDefaults

`func NewSharepointResourceDTOWithDefaults() *SharepointResourceDTO`

NewSharepointResourceDTOWithDefaults instantiates a new SharepointResourceDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivation

`func (o *SharepointResourceDTO) GetActivation() ActivationDTO`

GetActivation returns the Activation field if non-nil, zero value otherwise.

### GetActivationOk

`func (o *SharepointResourceDTO) GetActivationOk() (*ActivationDTO, bool)`

GetActivationOk returns a tuple with the Activation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivation

`func (o *SharepointResourceDTO) SetActivation(v ActivationDTO)`

SetActivation sets Activation field to given value.

### HasActivation

`func (o *SharepointResourceDTO) HasActivation() bool`

HasActivation returns a boolean if a field has been set.

### GetDefaultResource

`func (o *SharepointResourceDTO) GetDefaultResource() string`

GetDefaultResource returns the DefaultResource field if non-nil, zero value otherwise.

### GetDefaultResourceOk

`func (o *SharepointResourceDTO) GetDefaultResourceOk() (*string, bool)`

GetDefaultResourceOk returns a tuple with the DefaultResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultResource

`func (o *SharepointResourceDTO) SetDefaultResource(v string)`

SetDefaultResource sets DefaultResource field to given value.

### HasDefaultResource

`func (o *SharepointResourceDTO) HasDefaultResource() bool`

HasDefaultResource returns a boolean if a field has been set.

### GetDescription

`func (o *SharepointResourceDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SharepointResourceDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SharepointResourceDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SharepointResourceDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetElementId

`func (o *SharepointResourceDTO) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *SharepointResourceDTO) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *SharepointResourceDTO) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *SharepointResourceDTO) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetId

`func (o *SharepointResourceDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SharepointResourceDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SharepointResourceDTO) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *SharepointResourceDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIgnoredWebResources

`func (o *SharepointResourceDTO) GetIgnoredWebResources() []string`

GetIgnoredWebResources returns the IgnoredWebResources field if non-nil, zero value otherwise.

### GetIgnoredWebResourcesOk

`func (o *SharepointResourceDTO) GetIgnoredWebResourcesOk() (*[]string, bool)`

GetIgnoredWebResourcesOk returns a tuple with the IgnoredWebResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoredWebResources

`func (o *SharepointResourceDTO) SetIgnoredWebResources(v []string)`

SetIgnoredWebResources sets IgnoredWebResources field to given value.

### HasIgnoredWebResources

`func (o *SharepointResourceDTO) HasIgnoredWebResources() bool`

HasIgnoredWebResources returns a boolean if a field has been set.

### GetName

`func (o *SharepointResourceDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SharepointResourceDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SharepointResourceDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SharepointResourceDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPartnerAppLocation

`func (o *SharepointResourceDTO) GetPartnerAppLocation() LocationDTO`

GetPartnerAppLocation returns the PartnerAppLocation field if non-nil, zero value otherwise.

### GetPartnerAppLocationOk

`func (o *SharepointResourceDTO) GetPartnerAppLocationOk() (*LocationDTO, bool)`

GetPartnerAppLocationOk returns a tuple with the PartnerAppLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerAppLocation

`func (o *SharepointResourceDTO) SetPartnerAppLocation(v LocationDTO)`

SetPartnerAppLocation sets PartnerAppLocation field to given value.

### HasPartnerAppLocation

`func (o *SharepointResourceDTO) HasPartnerAppLocation() bool`

HasPartnerAppLocation returns a boolean if a field has been set.

### GetRelayingPartyLocation

`func (o *SharepointResourceDTO) GetRelayingPartyLocation() LocationDTO`

GetRelayingPartyLocation returns the RelayingPartyLocation field if non-nil, zero value otherwise.

### GetRelayingPartyLocationOk

`func (o *SharepointResourceDTO) GetRelayingPartyLocationOk() (*LocationDTO, bool)`

GetRelayingPartyLocationOk returns a tuple with the RelayingPartyLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayingPartyLocation

`func (o *SharepointResourceDTO) SetRelayingPartyLocation(v LocationDTO)`

SetRelayingPartyLocation sets RelayingPartyLocation field to given value.

### HasRelayingPartyLocation

`func (o *SharepointResourceDTO) HasRelayingPartyLocation() bool`

HasRelayingPartyLocation returns a boolean if a field has been set.

### GetServiceConnection

`func (o *SharepointResourceDTO) GetServiceConnection() ServiceConnectionDTO`

GetServiceConnection returns the ServiceConnection field if non-nil, zero value otherwise.

### GetServiceConnectionOk

`func (o *SharepointResourceDTO) GetServiceConnectionOk() (*ServiceConnectionDTO, bool)`

GetServiceConnectionOk returns a tuple with the ServiceConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceConnection

`func (o *SharepointResourceDTO) SetServiceConnection(v ServiceConnectionDTO)`

SetServiceConnection sets ServiceConnection field to given value.

### HasServiceConnection

`func (o *SharepointResourceDTO) HasServiceConnection() bool`

HasServiceConnection returns a boolean if a field has been set.

### GetSloLocation

`func (o *SharepointResourceDTO) GetSloLocation() LocationDTO`

GetSloLocation returns the SloLocation field if non-nil, zero value otherwise.

### GetSloLocationOk

`func (o *SharepointResourceDTO) GetSloLocationOk() (*LocationDTO, bool)`

GetSloLocationOk returns a tuple with the SloLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSloLocation

`func (o *SharepointResourceDTO) SetSloLocation(v LocationDTO)`

SetSloLocation sets SloLocation field to given value.

### HasSloLocation

`func (o *SharepointResourceDTO) HasSloLocation() bool`

HasSloLocation returns a boolean if a field has been set.

### GetSloLocationEnabled

`func (o *SharepointResourceDTO) GetSloLocationEnabled() bool`

GetSloLocationEnabled returns the SloLocationEnabled field if non-nil, zero value otherwise.

### GetSloLocationEnabledOk

`func (o *SharepointResourceDTO) GetSloLocationEnabledOk() (*bool, bool)`

GetSloLocationEnabledOk returns a tuple with the SloLocationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSloLocationEnabled

`func (o *SharepointResourceDTO) SetSloLocationEnabled(v bool)`

SetSloLocationEnabled sets SloLocationEnabled field to given value.

### HasSloLocationEnabled

`func (o *SharepointResourceDTO) HasSloLocationEnabled() bool`

HasSloLocationEnabled returns a boolean if a field has been set.

### GetStsEncryptingCertSubject

`func (o *SharepointResourceDTO) GetStsEncryptingCertSubject() string`

GetStsEncryptingCertSubject returns the StsEncryptingCertSubject field if non-nil, zero value otherwise.

### GetStsEncryptingCertSubjectOk

`func (o *SharepointResourceDTO) GetStsEncryptingCertSubjectOk() (*string, bool)`

GetStsEncryptingCertSubjectOk returns a tuple with the StsEncryptingCertSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStsEncryptingCertSubject

`func (o *SharepointResourceDTO) SetStsEncryptingCertSubject(v string)`

SetStsEncryptingCertSubject sets StsEncryptingCertSubject field to given value.

### HasStsEncryptingCertSubject

`func (o *SharepointResourceDTO) HasStsEncryptingCertSubject() bool`

HasStsEncryptingCertSubject returns a boolean if a field has been set.

### GetStsMetadata

`func (o *SharepointResourceDTO) GetStsMetadata() ResourceDTO`

GetStsMetadata returns the StsMetadata field if non-nil, zero value otherwise.

### GetStsMetadataOk

`func (o *SharepointResourceDTO) GetStsMetadataOk() (*ResourceDTO, bool)`

GetStsMetadataOk returns a tuple with the StsMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStsMetadata

`func (o *SharepointResourceDTO) SetStsMetadata(v ResourceDTO)`

SetStsMetadata sets StsMetadata field to given value.

### HasStsMetadata

`func (o *SharepointResourceDTO) HasStsMetadata() bool`

HasStsMetadata returns a boolean if a field has been set.

### GetStsSigningCertSubject

`func (o *SharepointResourceDTO) GetStsSigningCertSubject() string`

GetStsSigningCertSubject returns the StsSigningCertSubject field if non-nil, zero value otherwise.

### GetStsSigningCertSubjectOk

`func (o *SharepointResourceDTO) GetStsSigningCertSubjectOk() (*string, bool)`

GetStsSigningCertSubjectOk returns a tuple with the StsSigningCertSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStsSigningCertSubject

`func (o *SharepointResourceDTO) SetStsSigningCertSubject(v string)`

SetStsSigningCertSubject sets StsSigningCertSubject field to given value.

### HasStsSigningCertSubject

`func (o *SharepointResourceDTO) HasStsSigningCertSubject() bool`

HasStsSigningCertSubject returns a boolean if a field has been set.

### GetX

`func (o *SharepointResourceDTO) GetX() float64`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *SharepointResourceDTO) GetXOk() (*float64, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *SharepointResourceDTO) SetX(v float64)`

SetX sets X field to given value.

### HasX

`func (o *SharepointResourceDTO) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *SharepointResourceDTO) GetY() float64`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *SharepointResourceDTO) GetYOk() (*float64, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *SharepointResourceDTO) SetY(v float64)`

SetY sets Y field to given value.

### HasY

`func (o *SharepointResourceDTO) HasY() bool`

HasY returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


