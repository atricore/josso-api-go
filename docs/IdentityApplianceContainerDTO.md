# IdentityApplianceContainerDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Appliance** | Pointer to [**IdentityApplianceDTO**](IdentityApplianceDTO.md) |  | [optional] 
**IdSources** | Pointer to **[]string** |  | [optional] 
**Providers** | Pointer to **[]string** |  | [optional] 

## Methods

### NewIdentityApplianceContainerDTO

`func NewIdentityApplianceContainerDTO() *IdentityApplianceContainerDTO`

NewIdentityApplianceContainerDTO instantiates a new IdentityApplianceContainerDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityApplianceContainerDTOWithDefaults

`func NewIdentityApplianceContainerDTOWithDefaults() *IdentityApplianceContainerDTO`

NewIdentityApplianceContainerDTOWithDefaults instantiates a new IdentityApplianceContainerDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppliance

`func (o *IdentityApplianceContainerDTO) GetAppliance() IdentityApplianceDTO`

GetAppliance returns the Appliance field if non-nil, zero value otherwise.

### GetApplianceOk

`func (o *IdentityApplianceContainerDTO) GetApplianceOk() (*IdentityApplianceDTO, bool)`

GetApplianceOk returns a tuple with the Appliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliance

`func (o *IdentityApplianceContainerDTO) SetAppliance(v IdentityApplianceDTO)`

SetAppliance sets Appliance field to given value.

### HasAppliance

`func (o *IdentityApplianceContainerDTO) HasAppliance() bool`

HasAppliance returns a boolean if a field has been set.

### GetIdSources

`func (o *IdentityApplianceContainerDTO) GetIdSources() []string`

GetIdSources returns the IdSources field if non-nil, zero value otherwise.

### GetIdSourcesOk

`func (o *IdentityApplianceContainerDTO) GetIdSourcesOk() (*[]string, bool)`

GetIdSourcesOk returns a tuple with the IdSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdSources

`func (o *IdentityApplianceContainerDTO) SetIdSources(v []string)`

SetIdSources sets IdSources field to given value.

### HasIdSources

`func (o *IdentityApplianceContainerDTO) HasIdSources() bool`

HasIdSources returns a boolean if a field has been set.

### GetProviders

`func (o *IdentityApplianceContainerDTO) GetProviders() []string`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *IdentityApplianceContainerDTO) GetProvidersOk() (*[]string, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *IdentityApplianceContainerDTO) SetProviders(v []string)`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *IdentityApplianceContainerDTO) HasProviders() bool`

HasProviders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


