# ExternalSaml2IdentityProviderDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveBindings** | Pointer to **[]string** |  | [optional] 
**ActiveProfiles** | Pointer to **[]string** |  | [optional] 
**Config** | Pointer to [**ProviderConfigDTO**](ProviderConfigDTO.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**FederatedConnectionsA** | Pointer to [**[]FederatedConnectionDTO**](FederatedConnectionDTO.md) |  | [optional] 
**FederatedConnectionsB** | Pointer to [**[]FederatedConnectionDTO**](FederatedConnectionDTO.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**IdentityAppliance** | Pointer to [**IdentityApplianceDefinitionDTO**](IdentityApplianceDefinitionDTO.md) |  | [optional] 
**IdentityLookups** | Pointer to [**[]IdentityLookupDTO**](IdentityLookupDTO.md) |  | [optional] 
**IsRemote** | Pointer to **bool** |  | [optional] 
**Location** | Pointer to [**LocationDTO**](LocationDTO.md) |  | [optional] 
**Metadata** | Pointer to [**ResourceDTO**](ResourceDTO.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Remote** | Pointer to **bool** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**X** | Pointer to **float64** |  | [optional] 
**Y** | Pointer to **float64** |  | [optional] 

## Methods

### NewExternalSaml2IdentityProviderDTO

`func NewExternalSaml2IdentityProviderDTO() *ExternalSaml2IdentityProviderDTO`

NewExternalSaml2IdentityProviderDTO instantiates a new ExternalSaml2IdentityProviderDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalSaml2IdentityProviderDTOWithDefaults

`func NewExternalSaml2IdentityProviderDTOWithDefaults() *ExternalSaml2IdentityProviderDTO`

NewExternalSaml2IdentityProviderDTOWithDefaults instantiates a new ExternalSaml2IdentityProviderDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveBindings

`func (o *ExternalSaml2IdentityProviderDTO) GetActiveBindings() []string`

GetActiveBindings returns the ActiveBindings field if non-nil, zero value otherwise.

### GetActiveBindingsOk

`func (o *ExternalSaml2IdentityProviderDTO) GetActiveBindingsOk() (*[]string, bool)`

GetActiveBindingsOk returns a tuple with the ActiveBindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveBindings

`func (o *ExternalSaml2IdentityProviderDTO) SetActiveBindings(v []string)`

SetActiveBindings sets ActiveBindings field to given value.

### HasActiveBindings

`func (o *ExternalSaml2IdentityProviderDTO) HasActiveBindings() bool`

HasActiveBindings returns a boolean if a field has been set.

### GetActiveProfiles

`func (o *ExternalSaml2IdentityProviderDTO) GetActiveProfiles() []string`

GetActiveProfiles returns the ActiveProfiles field if non-nil, zero value otherwise.

### GetActiveProfilesOk

`func (o *ExternalSaml2IdentityProviderDTO) GetActiveProfilesOk() (*[]string, bool)`

GetActiveProfilesOk returns a tuple with the ActiveProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveProfiles

`func (o *ExternalSaml2IdentityProviderDTO) SetActiveProfiles(v []string)`

SetActiveProfiles sets ActiveProfiles field to given value.

### HasActiveProfiles

`func (o *ExternalSaml2IdentityProviderDTO) HasActiveProfiles() bool`

HasActiveProfiles returns a boolean if a field has been set.

### GetConfig

`func (o *ExternalSaml2IdentityProviderDTO) GetConfig() ProviderConfigDTO`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ExternalSaml2IdentityProviderDTO) GetConfigOk() (*ProviderConfigDTO, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ExternalSaml2IdentityProviderDTO) SetConfig(v ProviderConfigDTO)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ExternalSaml2IdentityProviderDTO) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetDescription

`func (o *ExternalSaml2IdentityProviderDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExternalSaml2IdentityProviderDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExternalSaml2IdentityProviderDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExternalSaml2IdentityProviderDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *ExternalSaml2IdentityProviderDTO) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ExternalSaml2IdentityProviderDTO) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ExternalSaml2IdentityProviderDTO) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ExternalSaml2IdentityProviderDTO) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetElementId

`func (o *ExternalSaml2IdentityProviderDTO) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *ExternalSaml2IdentityProviderDTO) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *ExternalSaml2IdentityProviderDTO) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *ExternalSaml2IdentityProviderDTO) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetFederatedConnectionsA

`func (o *ExternalSaml2IdentityProviderDTO) GetFederatedConnectionsA() []FederatedConnectionDTO`

GetFederatedConnectionsA returns the FederatedConnectionsA field if non-nil, zero value otherwise.

### GetFederatedConnectionsAOk

`func (o *ExternalSaml2IdentityProviderDTO) GetFederatedConnectionsAOk() (*[]FederatedConnectionDTO, bool)`

GetFederatedConnectionsAOk returns a tuple with the FederatedConnectionsA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederatedConnectionsA

`func (o *ExternalSaml2IdentityProviderDTO) SetFederatedConnectionsA(v []FederatedConnectionDTO)`

SetFederatedConnectionsA sets FederatedConnectionsA field to given value.

### HasFederatedConnectionsA

`func (o *ExternalSaml2IdentityProviderDTO) HasFederatedConnectionsA() bool`

HasFederatedConnectionsA returns a boolean if a field has been set.

### GetFederatedConnectionsB

`func (o *ExternalSaml2IdentityProviderDTO) GetFederatedConnectionsB() []FederatedConnectionDTO`

GetFederatedConnectionsB returns the FederatedConnectionsB field if non-nil, zero value otherwise.

### GetFederatedConnectionsBOk

`func (o *ExternalSaml2IdentityProviderDTO) GetFederatedConnectionsBOk() (*[]FederatedConnectionDTO, bool)`

GetFederatedConnectionsBOk returns a tuple with the FederatedConnectionsB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederatedConnectionsB

`func (o *ExternalSaml2IdentityProviderDTO) SetFederatedConnectionsB(v []FederatedConnectionDTO)`

SetFederatedConnectionsB sets FederatedConnectionsB field to given value.

### HasFederatedConnectionsB

`func (o *ExternalSaml2IdentityProviderDTO) HasFederatedConnectionsB() bool`

HasFederatedConnectionsB returns a boolean if a field has been set.

### GetId

`func (o *ExternalSaml2IdentityProviderDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalSaml2IdentityProviderDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalSaml2IdentityProviderDTO) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ExternalSaml2IdentityProviderDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentityAppliance

`func (o *ExternalSaml2IdentityProviderDTO) GetIdentityAppliance() IdentityApplianceDefinitionDTO`

GetIdentityAppliance returns the IdentityAppliance field if non-nil, zero value otherwise.

### GetIdentityApplianceOk

`func (o *ExternalSaml2IdentityProviderDTO) GetIdentityApplianceOk() (*IdentityApplianceDefinitionDTO, bool)`

GetIdentityApplianceOk returns a tuple with the IdentityAppliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityAppliance

`func (o *ExternalSaml2IdentityProviderDTO) SetIdentityAppliance(v IdentityApplianceDefinitionDTO)`

SetIdentityAppliance sets IdentityAppliance field to given value.

### HasIdentityAppliance

`func (o *ExternalSaml2IdentityProviderDTO) HasIdentityAppliance() bool`

HasIdentityAppliance returns a boolean if a field has been set.

### GetIdentityLookups

`func (o *ExternalSaml2IdentityProviderDTO) GetIdentityLookups() []IdentityLookupDTO`

GetIdentityLookups returns the IdentityLookups field if non-nil, zero value otherwise.

### GetIdentityLookupsOk

`func (o *ExternalSaml2IdentityProviderDTO) GetIdentityLookupsOk() (*[]IdentityLookupDTO, bool)`

GetIdentityLookupsOk returns a tuple with the IdentityLookups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityLookups

`func (o *ExternalSaml2IdentityProviderDTO) SetIdentityLookups(v []IdentityLookupDTO)`

SetIdentityLookups sets IdentityLookups field to given value.

### HasIdentityLookups

`func (o *ExternalSaml2IdentityProviderDTO) HasIdentityLookups() bool`

HasIdentityLookups returns a boolean if a field has been set.

### GetIsRemote

`func (o *ExternalSaml2IdentityProviderDTO) GetIsRemote() bool`

GetIsRemote returns the IsRemote field if non-nil, zero value otherwise.

### GetIsRemoteOk

`func (o *ExternalSaml2IdentityProviderDTO) GetIsRemoteOk() (*bool, bool)`

GetIsRemoteOk returns a tuple with the IsRemote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemote

`func (o *ExternalSaml2IdentityProviderDTO) SetIsRemote(v bool)`

SetIsRemote sets IsRemote field to given value.

### HasIsRemote

`func (o *ExternalSaml2IdentityProviderDTO) HasIsRemote() bool`

HasIsRemote returns a boolean if a field has been set.

### GetLocation

`func (o *ExternalSaml2IdentityProviderDTO) GetLocation() LocationDTO`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ExternalSaml2IdentityProviderDTO) GetLocationOk() (*LocationDTO, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ExternalSaml2IdentityProviderDTO) SetLocation(v LocationDTO)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ExternalSaml2IdentityProviderDTO) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetMetadata

`func (o *ExternalSaml2IdentityProviderDTO) GetMetadata() ResourceDTO`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ExternalSaml2IdentityProviderDTO) GetMetadataOk() (*ResourceDTO, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ExternalSaml2IdentityProviderDTO) SetMetadata(v ResourceDTO)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ExternalSaml2IdentityProviderDTO) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *ExternalSaml2IdentityProviderDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExternalSaml2IdentityProviderDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExternalSaml2IdentityProviderDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExternalSaml2IdentityProviderDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRemote

`func (o *ExternalSaml2IdentityProviderDTO) GetRemote() bool`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *ExternalSaml2IdentityProviderDTO) GetRemoteOk() (*bool, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *ExternalSaml2IdentityProviderDTO) SetRemote(v bool)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *ExternalSaml2IdentityProviderDTO) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### GetRole

`func (o *ExternalSaml2IdentityProviderDTO) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ExternalSaml2IdentityProviderDTO) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ExternalSaml2IdentityProviderDTO) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *ExternalSaml2IdentityProviderDTO) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetX

`func (o *ExternalSaml2IdentityProviderDTO) GetX() float64`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *ExternalSaml2IdentityProviderDTO) GetXOk() (*float64, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *ExternalSaml2IdentityProviderDTO) SetX(v float64)`

SetX sets X field to given value.

### HasX

`func (o *ExternalSaml2IdentityProviderDTO) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *ExternalSaml2IdentityProviderDTO) GetY() float64`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *ExternalSaml2IdentityProviderDTO) GetYOk() (*float64, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *ExternalSaml2IdentityProviderDTO) SetY(v float64)`

SetY sets Y field to given value.

### HasY

`func (o *ExternalSaml2IdentityProviderDTO) HasY() bool`

HasY returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


