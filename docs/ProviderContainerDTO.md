# ProviderContainerDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FederatedProvider** | Pointer to [**FederatedProviderDTO**](FederatedProviderDTO.md) |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ServiceResource** | Pointer to [**ServiceResourceDTO**](ServiceResourceDTO.md) |  | [optional] 
**ServiceResourceType** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewProviderContainerDTO

`func NewProviderContainerDTO() *ProviderContainerDTO`

NewProviderContainerDTO instantiates a new ProviderContainerDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderContainerDTOWithDefaults

`func NewProviderContainerDTOWithDefaults() *ProviderContainerDTO`

NewProviderContainerDTOWithDefaults instantiates a new ProviderContainerDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFederatedProvider

`func (o *ProviderContainerDTO) GetFederatedProvider() FederatedProviderDTO`

GetFederatedProvider returns the FederatedProvider field if non-nil, zero value otherwise.

### GetFederatedProviderOk

`func (o *ProviderContainerDTO) GetFederatedProviderOk() (*FederatedProviderDTO, bool)`

GetFederatedProviderOk returns a tuple with the FederatedProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederatedProvider

`func (o *ProviderContainerDTO) SetFederatedProvider(v FederatedProviderDTO)`

SetFederatedProvider sets FederatedProvider field to given value.

### HasFederatedProvider

`func (o *ProviderContainerDTO) HasFederatedProvider() bool`

HasFederatedProvider returns a boolean if a field has been set.

### GetLocation

`func (o *ProviderContainerDTO) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ProviderContainerDTO) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ProviderContainerDTO) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ProviderContainerDTO) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetName

`func (o *ProviderContainerDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProviderContainerDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProviderContainerDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProviderContainerDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetServiceResource

`func (o *ProviderContainerDTO) GetServiceResource() ServiceResourceDTO`

GetServiceResource returns the ServiceResource field if non-nil, zero value otherwise.

### GetServiceResourceOk

`func (o *ProviderContainerDTO) GetServiceResourceOk() (*ServiceResourceDTO, bool)`

GetServiceResourceOk returns a tuple with the ServiceResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceResource

`func (o *ProviderContainerDTO) SetServiceResource(v ServiceResourceDTO)`

SetServiceResource sets ServiceResource field to given value.

### HasServiceResource

`func (o *ProviderContainerDTO) HasServiceResource() bool`

HasServiceResource returns a boolean if a field has been set.

### GetServiceResourceType

`func (o *ProviderContainerDTO) GetServiceResourceType() string`

GetServiceResourceType returns the ServiceResourceType field if non-nil, zero value otherwise.

### GetServiceResourceTypeOk

`func (o *ProviderContainerDTO) GetServiceResourceTypeOk() (*string, bool)`

GetServiceResourceTypeOk returns a tuple with the ServiceResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceResourceType

`func (o *ProviderContainerDTO) SetServiceResourceType(v string)`

SetServiceResourceType sets ServiceResourceType field to given value.

### HasServiceResourceType

`func (o *ProviderContainerDTO) HasServiceResourceType() bool`

HasServiceResourceType returns a boolean if a field has been set.

### GetType

`func (o *ProviderContainerDTO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProviderContainerDTO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProviderContainerDTO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ProviderContainerDTO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


