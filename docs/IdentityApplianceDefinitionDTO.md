# IdentityApplianceDefinitionDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveFeatures** | Pointer to **[]string** |  | [optional] 
**AuthenticationServices** | Pointer to [**[]AuthenticationServiceDTO**](AuthenticationServiceDTO.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**ExecutionEnvironments** | Pointer to [**[]ExecutionEnvironmentDTO**](ExecutionEnvironmentDTO.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**IdentitySources** | Pointer to [**[]IdentitySourceDTO**](IdentitySourceDTO.md) |  | [optional] 
**IdpSelector** | Pointer to [**EntitySelectionStrategyDTO**](EntitySelectionStrategyDTO.md) |  | [optional] 
**Keystore** | Pointer to [**KeystoreDTO**](KeystoreDTO.md) |  | [optional] 
**LastModification** | Pointer to **time.Time** |  | [optional] 
**Location** | Pointer to [**LocationDTO**](LocationDTO.md) |  | [optional] 
**ModelVersion** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Providers** | Pointer to [**[]ProviderDTO**](ProviderDTO.md) |  | [optional] 
**RequiredBundles** | Pointer to **[]string** |  | [optional] 
**Revision** | Pointer to **int32** |  | [optional] 
**SecurityConfig** | Pointer to [**IdentityApplianceSecurityConfigDTO**](IdentityApplianceSecurityConfigDTO.md) |  | [optional] 
**ServiceResources** | Pointer to [**[]ServiceResourceDTO**](ServiceResourceDTO.md) |  | [optional] 
**SupportedRoles** | Pointer to **[]string** |  | [optional] 
**UserDashboardBranding** | Pointer to [**UserDashboardBrandingDTO**](UserDashboardBrandingDTO.md) |  | [optional] 

## Methods

### NewIdentityApplianceDefinitionDTO

`func NewIdentityApplianceDefinitionDTO() *IdentityApplianceDefinitionDTO`

NewIdentityApplianceDefinitionDTO instantiates a new IdentityApplianceDefinitionDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityApplianceDefinitionDTOWithDefaults

`func NewIdentityApplianceDefinitionDTOWithDefaults() *IdentityApplianceDefinitionDTO`

NewIdentityApplianceDefinitionDTOWithDefaults instantiates a new IdentityApplianceDefinitionDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveFeatures

`func (o *IdentityApplianceDefinitionDTO) GetActiveFeatures() []string`

GetActiveFeatures returns the ActiveFeatures field if non-nil, zero value otherwise.

### GetActiveFeaturesOk

`func (o *IdentityApplianceDefinitionDTO) GetActiveFeaturesOk() (*[]string, bool)`

GetActiveFeaturesOk returns a tuple with the ActiveFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveFeatures

`func (o *IdentityApplianceDefinitionDTO) SetActiveFeatures(v []string)`

SetActiveFeatures sets ActiveFeatures field to given value.

### HasActiveFeatures

`func (o *IdentityApplianceDefinitionDTO) HasActiveFeatures() bool`

HasActiveFeatures returns a boolean if a field has been set.

### GetAuthenticationServices

`func (o *IdentityApplianceDefinitionDTO) GetAuthenticationServices() []AuthenticationServiceDTO`

GetAuthenticationServices returns the AuthenticationServices field if non-nil, zero value otherwise.

### GetAuthenticationServicesOk

`func (o *IdentityApplianceDefinitionDTO) GetAuthenticationServicesOk() (*[]AuthenticationServiceDTO, bool)`

GetAuthenticationServicesOk returns a tuple with the AuthenticationServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationServices

`func (o *IdentityApplianceDefinitionDTO) SetAuthenticationServices(v []AuthenticationServiceDTO)`

SetAuthenticationServices sets AuthenticationServices field to given value.

### HasAuthenticationServices

`func (o *IdentityApplianceDefinitionDTO) HasAuthenticationServices() bool`

HasAuthenticationServices returns a boolean if a field has been set.

### GetDescription

`func (o *IdentityApplianceDefinitionDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IdentityApplianceDefinitionDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IdentityApplianceDefinitionDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IdentityApplianceDefinitionDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *IdentityApplianceDefinitionDTO) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IdentityApplianceDefinitionDTO) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IdentityApplianceDefinitionDTO) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IdentityApplianceDefinitionDTO) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetElementId

`func (o *IdentityApplianceDefinitionDTO) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *IdentityApplianceDefinitionDTO) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *IdentityApplianceDefinitionDTO) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *IdentityApplianceDefinitionDTO) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetExecutionEnvironments

`func (o *IdentityApplianceDefinitionDTO) GetExecutionEnvironments() []ExecutionEnvironmentDTO`

GetExecutionEnvironments returns the ExecutionEnvironments field if non-nil, zero value otherwise.

### GetExecutionEnvironmentsOk

`func (o *IdentityApplianceDefinitionDTO) GetExecutionEnvironmentsOk() (*[]ExecutionEnvironmentDTO, bool)`

GetExecutionEnvironmentsOk returns a tuple with the ExecutionEnvironments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionEnvironments

`func (o *IdentityApplianceDefinitionDTO) SetExecutionEnvironments(v []ExecutionEnvironmentDTO)`

SetExecutionEnvironments sets ExecutionEnvironments field to given value.

### HasExecutionEnvironments

`func (o *IdentityApplianceDefinitionDTO) HasExecutionEnvironments() bool`

HasExecutionEnvironments returns a boolean if a field has been set.

### GetId

`func (o *IdentityApplianceDefinitionDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentityApplianceDefinitionDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentityApplianceDefinitionDTO) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *IdentityApplianceDefinitionDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentitySources

`func (o *IdentityApplianceDefinitionDTO) GetIdentitySources() []IdentitySourceDTO`

GetIdentitySources returns the IdentitySources field if non-nil, zero value otherwise.

### GetIdentitySourcesOk

`func (o *IdentityApplianceDefinitionDTO) GetIdentitySourcesOk() (*[]IdentitySourceDTO, bool)`

GetIdentitySourcesOk returns a tuple with the IdentitySources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentitySources

`func (o *IdentityApplianceDefinitionDTO) SetIdentitySources(v []IdentitySourceDTO)`

SetIdentitySources sets IdentitySources field to given value.

### HasIdentitySources

`func (o *IdentityApplianceDefinitionDTO) HasIdentitySources() bool`

HasIdentitySources returns a boolean if a field has been set.

### GetIdpSelector

`func (o *IdentityApplianceDefinitionDTO) GetIdpSelector() EntitySelectionStrategyDTO`

GetIdpSelector returns the IdpSelector field if non-nil, zero value otherwise.

### GetIdpSelectorOk

`func (o *IdentityApplianceDefinitionDTO) GetIdpSelectorOk() (*EntitySelectionStrategyDTO, bool)`

GetIdpSelectorOk returns a tuple with the IdpSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpSelector

`func (o *IdentityApplianceDefinitionDTO) SetIdpSelector(v EntitySelectionStrategyDTO)`

SetIdpSelector sets IdpSelector field to given value.

### HasIdpSelector

`func (o *IdentityApplianceDefinitionDTO) HasIdpSelector() bool`

HasIdpSelector returns a boolean if a field has been set.

### GetKeystore

`func (o *IdentityApplianceDefinitionDTO) GetKeystore() KeystoreDTO`

GetKeystore returns the Keystore field if non-nil, zero value otherwise.

### GetKeystoreOk

`func (o *IdentityApplianceDefinitionDTO) GetKeystoreOk() (*KeystoreDTO, bool)`

GetKeystoreOk returns a tuple with the Keystore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeystore

`func (o *IdentityApplianceDefinitionDTO) SetKeystore(v KeystoreDTO)`

SetKeystore sets Keystore field to given value.

### HasKeystore

`func (o *IdentityApplianceDefinitionDTO) HasKeystore() bool`

HasKeystore returns a boolean if a field has been set.

### GetLastModification

`func (o *IdentityApplianceDefinitionDTO) GetLastModification() time.Time`

GetLastModification returns the LastModification field if non-nil, zero value otherwise.

### GetLastModificationOk

`func (o *IdentityApplianceDefinitionDTO) GetLastModificationOk() (*time.Time, bool)`

GetLastModificationOk returns a tuple with the LastModification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModification

`func (o *IdentityApplianceDefinitionDTO) SetLastModification(v time.Time)`

SetLastModification sets LastModification field to given value.

### HasLastModification

`func (o *IdentityApplianceDefinitionDTO) HasLastModification() bool`

HasLastModification returns a boolean if a field has been set.

### GetLocation

`func (o *IdentityApplianceDefinitionDTO) GetLocation() LocationDTO`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *IdentityApplianceDefinitionDTO) GetLocationOk() (*LocationDTO, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *IdentityApplianceDefinitionDTO) SetLocation(v LocationDTO)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *IdentityApplianceDefinitionDTO) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetModelVersion

`func (o *IdentityApplianceDefinitionDTO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *IdentityApplianceDefinitionDTO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *IdentityApplianceDefinitionDTO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *IdentityApplianceDefinitionDTO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *IdentityApplianceDefinitionDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdentityApplianceDefinitionDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdentityApplianceDefinitionDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IdentityApplianceDefinitionDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *IdentityApplianceDefinitionDTO) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *IdentityApplianceDefinitionDTO) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *IdentityApplianceDefinitionDTO) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *IdentityApplianceDefinitionDTO) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetProviders

`func (o *IdentityApplianceDefinitionDTO) GetProviders() []ProviderDTO`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *IdentityApplianceDefinitionDTO) GetProvidersOk() (*[]ProviderDTO, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *IdentityApplianceDefinitionDTO) SetProviders(v []ProviderDTO)`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *IdentityApplianceDefinitionDTO) HasProviders() bool`

HasProviders returns a boolean if a field has been set.

### GetRequiredBundles

`func (o *IdentityApplianceDefinitionDTO) GetRequiredBundles() []string`

GetRequiredBundles returns the RequiredBundles field if non-nil, zero value otherwise.

### GetRequiredBundlesOk

`func (o *IdentityApplianceDefinitionDTO) GetRequiredBundlesOk() (*[]string, bool)`

GetRequiredBundlesOk returns a tuple with the RequiredBundles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredBundles

`func (o *IdentityApplianceDefinitionDTO) SetRequiredBundles(v []string)`

SetRequiredBundles sets RequiredBundles field to given value.

### HasRequiredBundles

`func (o *IdentityApplianceDefinitionDTO) HasRequiredBundles() bool`

HasRequiredBundles returns a boolean if a field has been set.

### GetRevision

`func (o *IdentityApplianceDefinitionDTO) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *IdentityApplianceDefinitionDTO) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *IdentityApplianceDefinitionDTO) SetRevision(v int32)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *IdentityApplianceDefinitionDTO) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetSecurityConfig

`func (o *IdentityApplianceDefinitionDTO) GetSecurityConfig() IdentityApplianceSecurityConfigDTO`

GetSecurityConfig returns the SecurityConfig field if non-nil, zero value otherwise.

### GetSecurityConfigOk

`func (o *IdentityApplianceDefinitionDTO) GetSecurityConfigOk() (*IdentityApplianceSecurityConfigDTO, bool)`

GetSecurityConfigOk returns a tuple with the SecurityConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityConfig

`func (o *IdentityApplianceDefinitionDTO) SetSecurityConfig(v IdentityApplianceSecurityConfigDTO)`

SetSecurityConfig sets SecurityConfig field to given value.

### HasSecurityConfig

`func (o *IdentityApplianceDefinitionDTO) HasSecurityConfig() bool`

HasSecurityConfig returns a boolean if a field has been set.

### GetServiceResources

`func (o *IdentityApplianceDefinitionDTO) GetServiceResources() []ServiceResourceDTO`

GetServiceResources returns the ServiceResources field if non-nil, zero value otherwise.

### GetServiceResourcesOk

`func (o *IdentityApplianceDefinitionDTO) GetServiceResourcesOk() (*[]ServiceResourceDTO, bool)`

GetServiceResourcesOk returns a tuple with the ServiceResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceResources

`func (o *IdentityApplianceDefinitionDTO) SetServiceResources(v []ServiceResourceDTO)`

SetServiceResources sets ServiceResources field to given value.

### HasServiceResources

`func (o *IdentityApplianceDefinitionDTO) HasServiceResources() bool`

HasServiceResources returns a boolean if a field has been set.

### GetSupportedRoles

`func (o *IdentityApplianceDefinitionDTO) GetSupportedRoles() []string`

GetSupportedRoles returns the SupportedRoles field if non-nil, zero value otherwise.

### GetSupportedRolesOk

`func (o *IdentityApplianceDefinitionDTO) GetSupportedRolesOk() (*[]string, bool)`

GetSupportedRolesOk returns a tuple with the SupportedRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedRoles

`func (o *IdentityApplianceDefinitionDTO) SetSupportedRoles(v []string)`

SetSupportedRoles sets SupportedRoles field to given value.

### HasSupportedRoles

`func (o *IdentityApplianceDefinitionDTO) HasSupportedRoles() bool`

HasSupportedRoles returns a boolean if a field has been set.

### GetUserDashboardBranding

`func (o *IdentityApplianceDefinitionDTO) GetUserDashboardBranding() UserDashboardBrandingDTO`

GetUserDashboardBranding returns the UserDashboardBranding field if non-nil, zero value otherwise.

### GetUserDashboardBrandingOk

`func (o *IdentityApplianceDefinitionDTO) GetUserDashboardBrandingOk() (*UserDashboardBrandingDTO, bool)`

GetUserDashboardBrandingOk returns a tuple with the UserDashboardBranding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDashboardBranding

`func (o *IdentityApplianceDefinitionDTO) SetUserDashboardBranding(v UserDashboardBrandingDTO)`

SetUserDashboardBranding sets UserDashboardBranding field to given value.

### HasUserDashboardBranding

`func (o *IdentityApplianceDefinitionDTO) HasUserDashboardBranding() bool`

HasUserDashboardBranding returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


