# VirtualSaml2ServiceProviderDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountLinkagePolicy** | Pointer to [**AccountLinkagePolicyDTO**](AccountLinkagePolicyDTO.md) |  | [optional] 
**ActiveBindings** | Pointer to **[]string** |  | [optional] 
**ActiveProfiles** | Pointer to **[]string** |  | [optional] 
**AttributeProfile** | Pointer to [**AttributeProfileDTO**](AttributeProfileDTO.md) |  | [optional] 
**Config** | Pointer to [**ProviderConfigDTO**](ProviderConfigDTO.md) |  | [optional] 
**DashboardUrl** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**EnableMetadataEndpoint** | Pointer to **bool** |  | [optional] 
**EnableProxyExtension** | Pointer to **bool** |  | [optional] 
**EncryptAssertion** | Pointer to **bool** |  | [optional] 
**EncryptAssertionAlgorithm** | Pointer to **string** |  | [optional] 
**ErrorBinding** | Pointer to **string** |  | [optional] 
**FederatedConnectionsA** | Pointer to [**[]FederatedConnectionDTO**](FederatedConnectionDTO.md) |  | [optional] 
**FederatedConnectionsB** | Pointer to [**[]FederatedConnectionDTO**](FederatedConnectionDTO.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**IdentityAppliance** | Pointer to [**IdentityApplianceDefinitionDTO**](IdentityApplianceDefinitionDTO.md) |  | [optional] 
**IdentityLookups** | Pointer to [**[]IdentityLookupDTO**](IdentityLookupDTO.md) |  | [optional] 
**IdentityMappingPolicy** | Pointer to [**IdentityMappingPolicyDTO**](IdentityMappingPolicyDTO.md) |  | [optional] 
**IdpSignatureHash** | Pointer to **string** |  | [optional] 
**IgnoreRequestedNameIDPolicy** | Pointer to **bool** |  | [optional] 
**IsRemote** | Pointer to **bool** |  | [optional] 
**Location** | Pointer to [**LocationDTO**](LocationDTO.md) |  | [optional] 
**MessageTtl** | Pointer to **int32** |  | [optional] 
**MessageTtlTolerance** | Pointer to **int32** |  | [optional] 
**Metadata** | Pointer to [**ResourceDTO**](ResourceDTO.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Oauth2Enabled** | Pointer to **bool** |  | [optional] 
**Oauth2Key** | Pointer to **string** |  | [optional] 
**Oauth2RememberMeTokenValidity** | Pointer to **int64** |  | [optional] 
**Oauth2TokenValidity** | Pointer to **int64** |  | [optional] 
**OidcAccessTokenTimeToLive** | Pointer to **int32** |  | [optional] 
**OidcAuthzCodeTimeToLive** | Pointer to **int32** |  | [optional] 
**OidcIdTokenTimeToLive** | Pointer to **int32** |  | [optional] 
**OpenIdEnabled** | Pointer to **bool** |  | [optional] 
**Remote** | Pointer to **bool** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**SessionManagerFactory** | Pointer to [**SessionManagerFactoryDTO**](SessionManagerFactoryDTO.md) |  | [optional] 
**SignAuthenticationRequests** | Pointer to **bool** |  | [optional] 
**SignRequests** | Pointer to **bool** |  | [optional] 
**SpSignatureHash** | Pointer to **string** |  | [optional] 
**SsoSessionTimeout** | Pointer to **int32** |  | [optional] 
**SubjectAuthnPolicies** | Pointer to [**[]SubjectAuthenticationPolicyDTO**](SubjectAuthenticationPolicyDTO.md) |  | [optional] 
**SubjectNameIDPolicy** | Pointer to [**SubjectNameIdentifierPolicyDTO**](SubjectNameIdentifierPolicyDTO.md) |  | [optional] 
**WantAssertionSigned** | Pointer to **bool** |  | [optional] 
**WantAuthnRequestsSigned** | Pointer to **bool** |  | [optional] 
**WantSLOResponseSigned** | Pointer to **bool** |  | [optional] 
**WantSignedRequests** | Pointer to **bool** |  | [optional] 
**X** | Pointer to **float64** |  | [optional] 
**Y** | Pointer to **float64** |  | [optional] 

## Methods

### NewVirtualSaml2ServiceProviderDTO

`func NewVirtualSaml2ServiceProviderDTO() *VirtualSaml2ServiceProviderDTO`

NewVirtualSaml2ServiceProviderDTO instantiates a new VirtualSaml2ServiceProviderDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualSaml2ServiceProviderDTOWithDefaults

`func NewVirtualSaml2ServiceProviderDTOWithDefaults() *VirtualSaml2ServiceProviderDTO`

NewVirtualSaml2ServiceProviderDTOWithDefaults instantiates a new VirtualSaml2ServiceProviderDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountLinkagePolicy

`func (o *VirtualSaml2ServiceProviderDTO) GetAccountLinkagePolicy() AccountLinkagePolicyDTO`

GetAccountLinkagePolicy returns the AccountLinkagePolicy field if non-nil, zero value otherwise.

### GetAccountLinkagePolicyOk

`func (o *VirtualSaml2ServiceProviderDTO) GetAccountLinkagePolicyOk() (*AccountLinkagePolicyDTO, bool)`

GetAccountLinkagePolicyOk returns a tuple with the AccountLinkagePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountLinkagePolicy

`func (o *VirtualSaml2ServiceProviderDTO) SetAccountLinkagePolicy(v AccountLinkagePolicyDTO)`

SetAccountLinkagePolicy sets AccountLinkagePolicy field to given value.

### HasAccountLinkagePolicy

`func (o *VirtualSaml2ServiceProviderDTO) HasAccountLinkagePolicy() bool`

HasAccountLinkagePolicy returns a boolean if a field has been set.

### GetActiveBindings

`func (o *VirtualSaml2ServiceProviderDTO) GetActiveBindings() []string`

GetActiveBindings returns the ActiveBindings field if non-nil, zero value otherwise.

### GetActiveBindingsOk

`func (o *VirtualSaml2ServiceProviderDTO) GetActiveBindingsOk() (*[]string, bool)`

GetActiveBindingsOk returns a tuple with the ActiveBindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveBindings

`func (o *VirtualSaml2ServiceProviderDTO) SetActiveBindings(v []string)`

SetActiveBindings sets ActiveBindings field to given value.

### HasActiveBindings

`func (o *VirtualSaml2ServiceProviderDTO) HasActiveBindings() bool`

HasActiveBindings returns a boolean if a field has been set.

### GetActiveProfiles

`func (o *VirtualSaml2ServiceProviderDTO) GetActiveProfiles() []string`

GetActiveProfiles returns the ActiveProfiles field if non-nil, zero value otherwise.

### GetActiveProfilesOk

`func (o *VirtualSaml2ServiceProviderDTO) GetActiveProfilesOk() (*[]string, bool)`

GetActiveProfilesOk returns a tuple with the ActiveProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveProfiles

`func (o *VirtualSaml2ServiceProviderDTO) SetActiveProfiles(v []string)`

SetActiveProfiles sets ActiveProfiles field to given value.

### HasActiveProfiles

`func (o *VirtualSaml2ServiceProviderDTO) HasActiveProfiles() bool`

HasActiveProfiles returns a boolean if a field has been set.

### GetAttributeProfile

`func (o *VirtualSaml2ServiceProviderDTO) GetAttributeProfile() AttributeProfileDTO`

GetAttributeProfile returns the AttributeProfile field if non-nil, zero value otherwise.

### GetAttributeProfileOk

`func (o *VirtualSaml2ServiceProviderDTO) GetAttributeProfileOk() (*AttributeProfileDTO, bool)`

GetAttributeProfileOk returns a tuple with the AttributeProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeProfile

`func (o *VirtualSaml2ServiceProviderDTO) SetAttributeProfile(v AttributeProfileDTO)`

SetAttributeProfile sets AttributeProfile field to given value.

### HasAttributeProfile

`func (o *VirtualSaml2ServiceProviderDTO) HasAttributeProfile() bool`

HasAttributeProfile returns a boolean if a field has been set.

### GetConfig

`func (o *VirtualSaml2ServiceProviderDTO) GetConfig() ProviderConfigDTO`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *VirtualSaml2ServiceProviderDTO) GetConfigOk() (*ProviderConfigDTO, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *VirtualSaml2ServiceProviderDTO) SetConfig(v ProviderConfigDTO)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *VirtualSaml2ServiceProviderDTO) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetDashboardUrl

`func (o *VirtualSaml2ServiceProviderDTO) GetDashboardUrl() string`

GetDashboardUrl returns the DashboardUrl field if non-nil, zero value otherwise.

### GetDashboardUrlOk

`func (o *VirtualSaml2ServiceProviderDTO) GetDashboardUrlOk() (*string, bool)`

GetDashboardUrlOk returns a tuple with the DashboardUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardUrl

`func (o *VirtualSaml2ServiceProviderDTO) SetDashboardUrl(v string)`

SetDashboardUrl sets DashboardUrl field to given value.

### HasDashboardUrl

`func (o *VirtualSaml2ServiceProviderDTO) HasDashboardUrl() bool`

HasDashboardUrl returns a boolean if a field has been set.

### GetDescription

`func (o *VirtualSaml2ServiceProviderDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VirtualSaml2ServiceProviderDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VirtualSaml2ServiceProviderDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VirtualSaml2ServiceProviderDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *VirtualSaml2ServiceProviderDTO) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *VirtualSaml2ServiceProviderDTO) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *VirtualSaml2ServiceProviderDTO) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *VirtualSaml2ServiceProviderDTO) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetElementId

`func (o *VirtualSaml2ServiceProviderDTO) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *VirtualSaml2ServiceProviderDTO) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *VirtualSaml2ServiceProviderDTO) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *VirtualSaml2ServiceProviderDTO) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetEnableMetadataEndpoint

`func (o *VirtualSaml2ServiceProviderDTO) GetEnableMetadataEndpoint() bool`

GetEnableMetadataEndpoint returns the EnableMetadataEndpoint field if non-nil, zero value otherwise.

### GetEnableMetadataEndpointOk

`func (o *VirtualSaml2ServiceProviderDTO) GetEnableMetadataEndpointOk() (*bool, bool)`

GetEnableMetadataEndpointOk returns a tuple with the EnableMetadataEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMetadataEndpoint

`func (o *VirtualSaml2ServiceProviderDTO) SetEnableMetadataEndpoint(v bool)`

SetEnableMetadataEndpoint sets EnableMetadataEndpoint field to given value.

### HasEnableMetadataEndpoint

`func (o *VirtualSaml2ServiceProviderDTO) HasEnableMetadataEndpoint() bool`

HasEnableMetadataEndpoint returns a boolean if a field has been set.

### GetEnableProxyExtension

`func (o *VirtualSaml2ServiceProviderDTO) GetEnableProxyExtension() bool`

GetEnableProxyExtension returns the EnableProxyExtension field if non-nil, zero value otherwise.

### GetEnableProxyExtensionOk

`func (o *VirtualSaml2ServiceProviderDTO) GetEnableProxyExtensionOk() (*bool, bool)`

GetEnableProxyExtensionOk returns a tuple with the EnableProxyExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableProxyExtension

`func (o *VirtualSaml2ServiceProviderDTO) SetEnableProxyExtension(v bool)`

SetEnableProxyExtension sets EnableProxyExtension field to given value.

### HasEnableProxyExtension

`func (o *VirtualSaml2ServiceProviderDTO) HasEnableProxyExtension() bool`

HasEnableProxyExtension returns a boolean if a field has been set.

### GetEncryptAssertion

`func (o *VirtualSaml2ServiceProviderDTO) GetEncryptAssertion() bool`

GetEncryptAssertion returns the EncryptAssertion field if non-nil, zero value otherwise.

### GetEncryptAssertionOk

`func (o *VirtualSaml2ServiceProviderDTO) GetEncryptAssertionOk() (*bool, bool)`

GetEncryptAssertionOk returns a tuple with the EncryptAssertion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptAssertion

`func (o *VirtualSaml2ServiceProviderDTO) SetEncryptAssertion(v bool)`

SetEncryptAssertion sets EncryptAssertion field to given value.

### HasEncryptAssertion

`func (o *VirtualSaml2ServiceProviderDTO) HasEncryptAssertion() bool`

HasEncryptAssertion returns a boolean if a field has been set.

### GetEncryptAssertionAlgorithm

`func (o *VirtualSaml2ServiceProviderDTO) GetEncryptAssertionAlgorithm() string`

GetEncryptAssertionAlgorithm returns the EncryptAssertionAlgorithm field if non-nil, zero value otherwise.

### GetEncryptAssertionAlgorithmOk

`func (o *VirtualSaml2ServiceProviderDTO) GetEncryptAssertionAlgorithmOk() (*string, bool)`

GetEncryptAssertionAlgorithmOk returns a tuple with the EncryptAssertionAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptAssertionAlgorithm

`func (o *VirtualSaml2ServiceProviderDTO) SetEncryptAssertionAlgorithm(v string)`

SetEncryptAssertionAlgorithm sets EncryptAssertionAlgorithm field to given value.

### HasEncryptAssertionAlgorithm

`func (o *VirtualSaml2ServiceProviderDTO) HasEncryptAssertionAlgorithm() bool`

HasEncryptAssertionAlgorithm returns a boolean if a field has been set.

### GetErrorBinding

`func (o *VirtualSaml2ServiceProviderDTO) GetErrorBinding() string`

GetErrorBinding returns the ErrorBinding field if non-nil, zero value otherwise.

### GetErrorBindingOk

`func (o *VirtualSaml2ServiceProviderDTO) GetErrorBindingOk() (*string, bool)`

GetErrorBindingOk returns a tuple with the ErrorBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorBinding

`func (o *VirtualSaml2ServiceProviderDTO) SetErrorBinding(v string)`

SetErrorBinding sets ErrorBinding field to given value.

### HasErrorBinding

`func (o *VirtualSaml2ServiceProviderDTO) HasErrorBinding() bool`

HasErrorBinding returns a boolean if a field has been set.

### GetFederatedConnectionsA

`func (o *VirtualSaml2ServiceProviderDTO) GetFederatedConnectionsA() []FederatedConnectionDTO`

GetFederatedConnectionsA returns the FederatedConnectionsA field if non-nil, zero value otherwise.

### GetFederatedConnectionsAOk

`func (o *VirtualSaml2ServiceProviderDTO) GetFederatedConnectionsAOk() (*[]FederatedConnectionDTO, bool)`

GetFederatedConnectionsAOk returns a tuple with the FederatedConnectionsA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederatedConnectionsA

`func (o *VirtualSaml2ServiceProviderDTO) SetFederatedConnectionsA(v []FederatedConnectionDTO)`

SetFederatedConnectionsA sets FederatedConnectionsA field to given value.

### HasFederatedConnectionsA

`func (o *VirtualSaml2ServiceProviderDTO) HasFederatedConnectionsA() bool`

HasFederatedConnectionsA returns a boolean if a field has been set.

### GetFederatedConnectionsB

`func (o *VirtualSaml2ServiceProviderDTO) GetFederatedConnectionsB() []FederatedConnectionDTO`

GetFederatedConnectionsB returns the FederatedConnectionsB field if non-nil, zero value otherwise.

### GetFederatedConnectionsBOk

`func (o *VirtualSaml2ServiceProviderDTO) GetFederatedConnectionsBOk() (*[]FederatedConnectionDTO, bool)`

GetFederatedConnectionsBOk returns a tuple with the FederatedConnectionsB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederatedConnectionsB

`func (o *VirtualSaml2ServiceProviderDTO) SetFederatedConnectionsB(v []FederatedConnectionDTO)`

SetFederatedConnectionsB sets FederatedConnectionsB field to given value.

### HasFederatedConnectionsB

`func (o *VirtualSaml2ServiceProviderDTO) HasFederatedConnectionsB() bool`

HasFederatedConnectionsB returns a boolean if a field has been set.

### GetId

`func (o *VirtualSaml2ServiceProviderDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VirtualSaml2ServiceProviderDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VirtualSaml2ServiceProviderDTO) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *VirtualSaml2ServiceProviderDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentityAppliance

`func (o *VirtualSaml2ServiceProviderDTO) GetIdentityAppliance() IdentityApplianceDefinitionDTO`

GetIdentityAppliance returns the IdentityAppliance field if non-nil, zero value otherwise.

### GetIdentityApplianceOk

`func (o *VirtualSaml2ServiceProviderDTO) GetIdentityApplianceOk() (*IdentityApplianceDefinitionDTO, bool)`

GetIdentityApplianceOk returns a tuple with the IdentityAppliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityAppliance

`func (o *VirtualSaml2ServiceProviderDTO) SetIdentityAppliance(v IdentityApplianceDefinitionDTO)`

SetIdentityAppliance sets IdentityAppliance field to given value.

### HasIdentityAppliance

`func (o *VirtualSaml2ServiceProviderDTO) HasIdentityAppliance() bool`

HasIdentityAppliance returns a boolean if a field has been set.

### GetIdentityLookups

`func (o *VirtualSaml2ServiceProviderDTO) GetIdentityLookups() []IdentityLookupDTO`

GetIdentityLookups returns the IdentityLookups field if non-nil, zero value otherwise.

### GetIdentityLookupsOk

`func (o *VirtualSaml2ServiceProviderDTO) GetIdentityLookupsOk() (*[]IdentityLookupDTO, bool)`

GetIdentityLookupsOk returns a tuple with the IdentityLookups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityLookups

`func (o *VirtualSaml2ServiceProviderDTO) SetIdentityLookups(v []IdentityLookupDTO)`

SetIdentityLookups sets IdentityLookups field to given value.

### HasIdentityLookups

`func (o *VirtualSaml2ServiceProviderDTO) HasIdentityLookups() bool`

HasIdentityLookups returns a boolean if a field has been set.

### GetIdentityMappingPolicy

`func (o *VirtualSaml2ServiceProviderDTO) GetIdentityMappingPolicy() IdentityMappingPolicyDTO`

GetIdentityMappingPolicy returns the IdentityMappingPolicy field if non-nil, zero value otherwise.

### GetIdentityMappingPolicyOk

`func (o *VirtualSaml2ServiceProviderDTO) GetIdentityMappingPolicyOk() (*IdentityMappingPolicyDTO, bool)`

GetIdentityMappingPolicyOk returns a tuple with the IdentityMappingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityMappingPolicy

`func (o *VirtualSaml2ServiceProviderDTO) SetIdentityMappingPolicy(v IdentityMappingPolicyDTO)`

SetIdentityMappingPolicy sets IdentityMappingPolicy field to given value.

### HasIdentityMappingPolicy

`func (o *VirtualSaml2ServiceProviderDTO) HasIdentityMappingPolicy() bool`

HasIdentityMappingPolicy returns a boolean if a field has been set.

### GetIdpSignatureHash

`func (o *VirtualSaml2ServiceProviderDTO) GetIdpSignatureHash() string`

GetIdpSignatureHash returns the IdpSignatureHash field if non-nil, zero value otherwise.

### GetIdpSignatureHashOk

`func (o *VirtualSaml2ServiceProviderDTO) GetIdpSignatureHashOk() (*string, bool)`

GetIdpSignatureHashOk returns a tuple with the IdpSignatureHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpSignatureHash

`func (o *VirtualSaml2ServiceProviderDTO) SetIdpSignatureHash(v string)`

SetIdpSignatureHash sets IdpSignatureHash field to given value.

### HasIdpSignatureHash

`func (o *VirtualSaml2ServiceProviderDTO) HasIdpSignatureHash() bool`

HasIdpSignatureHash returns a boolean if a field has been set.

### GetIgnoreRequestedNameIDPolicy

`func (o *VirtualSaml2ServiceProviderDTO) GetIgnoreRequestedNameIDPolicy() bool`

GetIgnoreRequestedNameIDPolicy returns the IgnoreRequestedNameIDPolicy field if non-nil, zero value otherwise.

### GetIgnoreRequestedNameIDPolicyOk

`func (o *VirtualSaml2ServiceProviderDTO) GetIgnoreRequestedNameIDPolicyOk() (*bool, bool)`

GetIgnoreRequestedNameIDPolicyOk returns a tuple with the IgnoreRequestedNameIDPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreRequestedNameIDPolicy

`func (o *VirtualSaml2ServiceProviderDTO) SetIgnoreRequestedNameIDPolicy(v bool)`

SetIgnoreRequestedNameIDPolicy sets IgnoreRequestedNameIDPolicy field to given value.

### HasIgnoreRequestedNameIDPolicy

`func (o *VirtualSaml2ServiceProviderDTO) HasIgnoreRequestedNameIDPolicy() bool`

HasIgnoreRequestedNameIDPolicy returns a boolean if a field has been set.

### GetIsRemote

`func (o *VirtualSaml2ServiceProviderDTO) GetIsRemote() bool`

GetIsRemote returns the IsRemote field if non-nil, zero value otherwise.

### GetIsRemoteOk

`func (o *VirtualSaml2ServiceProviderDTO) GetIsRemoteOk() (*bool, bool)`

GetIsRemoteOk returns a tuple with the IsRemote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemote

`func (o *VirtualSaml2ServiceProviderDTO) SetIsRemote(v bool)`

SetIsRemote sets IsRemote field to given value.

### HasIsRemote

`func (o *VirtualSaml2ServiceProviderDTO) HasIsRemote() bool`

HasIsRemote returns a boolean if a field has been set.

### GetLocation

`func (o *VirtualSaml2ServiceProviderDTO) GetLocation() LocationDTO`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *VirtualSaml2ServiceProviderDTO) GetLocationOk() (*LocationDTO, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *VirtualSaml2ServiceProviderDTO) SetLocation(v LocationDTO)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *VirtualSaml2ServiceProviderDTO) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetMessageTtl

`func (o *VirtualSaml2ServiceProviderDTO) GetMessageTtl() int32`

GetMessageTtl returns the MessageTtl field if non-nil, zero value otherwise.

### GetMessageTtlOk

`func (o *VirtualSaml2ServiceProviderDTO) GetMessageTtlOk() (*int32, bool)`

GetMessageTtlOk returns a tuple with the MessageTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageTtl

`func (o *VirtualSaml2ServiceProviderDTO) SetMessageTtl(v int32)`

SetMessageTtl sets MessageTtl field to given value.

### HasMessageTtl

`func (o *VirtualSaml2ServiceProviderDTO) HasMessageTtl() bool`

HasMessageTtl returns a boolean if a field has been set.

### GetMessageTtlTolerance

`func (o *VirtualSaml2ServiceProviderDTO) GetMessageTtlTolerance() int32`

GetMessageTtlTolerance returns the MessageTtlTolerance field if non-nil, zero value otherwise.

### GetMessageTtlToleranceOk

`func (o *VirtualSaml2ServiceProviderDTO) GetMessageTtlToleranceOk() (*int32, bool)`

GetMessageTtlToleranceOk returns a tuple with the MessageTtlTolerance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageTtlTolerance

`func (o *VirtualSaml2ServiceProviderDTO) SetMessageTtlTolerance(v int32)`

SetMessageTtlTolerance sets MessageTtlTolerance field to given value.

### HasMessageTtlTolerance

`func (o *VirtualSaml2ServiceProviderDTO) HasMessageTtlTolerance() bool`

HasMessageTtlTolerance returns a boolean if a field has been set.

### GetMetadata

`func (o *VirtualSaml2ServiceProviderDTO) GetMetadata() ResourceDTO`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *VirtualSaml2ServiceProviderDTO) GetMetadataOk() (*ResourceDTO, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *VirtualSaml2ServiceProviderDTO) SetMetadata(v ResourceDTO)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *VirtualSaml2ServiceProviderDTO) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *VirtualSaml2ServiceProviderDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualSaml2ServiceProviderDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualSaml2ServiceProviderDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualSaml2ServiceProviderDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOauth2Enabled

`func (o *VirtualSaml2ServiceProviderDTO) GetOauth2Enabled() bool`

GetOauth2Enabled returns the Oauth2Enabled field if non-nil, zero value otherwise.

### GetOauth2EnabledOk

`func (o *VirtualSaml2ServiceProviderDTO) GetOauth2EnabledOk() (*bool, bool)`

GetOauth2EnabledOk returns a tuple with the Oauth2Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2Enabled

`func (o *VirtualSaml2ServiceProviderDTO) SetOauth2Enabled(v bool)`

SetOauth2Enabled sets Oauth2Enabled field to given value.

### HasOauth2Enabled

`func (o *VirtualSaml2ServiceProviderDTO) HasOauth2Enabled() bool`

HasOauth2Enabled returns a boolean if a field has been set.

### GetOauth2Key

`func (o *VirtualSaml2ServiceProviderDTO) GetOauth2Key() string`

GetOauth2Key returns the Oauth2Key field if non-nil, zero value otherwise.

### GetOauth2KeyOk

`func (o *VirtualSaml2ServiceProviderDTO) GetOauth2KeyOk() (*string, bool)`

GetOauth2KeyOk returns a tuple with the Oauth2Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2Key

`func (o *VirtualSaml2ServiceProviderDTO) SetOauth2Key(v string)`

SetOauth2Key sets Oauth2Key field to given value.

### HasOauth2Key

`func (o *VirtualSaml2ServiceProviderDTO) HasOauth2Key() bool`

HasOauth2Key returns a boolean if a field has been set.

### GetOauth2RememberMeTokenValidity

`func (o *VirtualSaml2ServiceProviderDTO) GetOauth2RememberMeTokenValidity() int64`

GetOauth2RememberMeTokenValidity returns the Oauth2RememberMeTokenValidity field if non-nil, zero value otherwise.

### GetOauth2RememberMeTokenValidityOk

`func (o *VirtualSaml2ServiceProviderDTO) GetOauth2RememberMeTokenValidityOk() (*int64, bool)`

GetOauth2RememberMeTokenValidityOk returns a tuple with the Oauth2RememberMeTokenValidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2RememberMeTokenValidity

`func (o *VirtualSaml2ServiceProviderDTO) SetOauth2RememberMeTokenValidity(v int64)`

SetOauth2RememberMeTokenValidity sets Oauth2RememberMeTokenValidity field to given value.

### HasOauth2RememberMeTokenValidity

`func (o *VirtualSaml2ServiceProviderDTO) HasOauth2RememberMeTokenValidity() bool`

HasOauth2RememberMeTokenValidity returns a boolean if a field has been set.

### GetOauth2TokenValidity

`func (o *VirtualSaml2ServiceProviderDTO) GetOauth2TokenValidity() int64`

GetOauth2TokenValidity returns the Oauth2TokenValidity field if non-nil, zero value otherwise.

### GetOauth2TokenValidityOk

`func (o *VirtualSaml2ServiceProviderDTO) GetOauth2TokenValidityOk() (*int64, bool)`

GetOauth2TokenValidityOk returns a tuple with the Oauth2TokenValidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2TokenValidity

`func (o *VirtualSaml2ServiceProviderDTO) SetOauth2TokenValidity(v int64)`

SetOauth2TokenValidity sets Oauth2TokenValidity field to given value.

### HasOauth2TokenValidity

`func (o *VirtualSaml2ServiceProviderDTO) HasOauth2TokenValidity() bool`

HasOauth2TokenValidity returns a boolean if a field has been set.

### GetOidcAccessTokenTimeToLive

`func (o *VirtualSaml2ServiceProviderDTO) GetOidcAccessTokenTimeToLive() int32`

GetOidcAccessTokenTimeToLive returns the OidcAccessTokenTimeToLive field if non-nil, zero value otherwise.

### GetOidcAccessTokenTimeToLiveOk

`func (o *VirtualSaml2ServiceProviderDTO) GetOidcAccessTokenTimeToLiveOk() (*int32, bool)`

GetOidcAccessTokenTimeToLiveOk returns a tuple with the OidcAccessTokenTimeToLive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcAccessTokenTimeToLive

`func (o *VirtualSaml2ServiceProviderDTO) SetOidcAccessTokenTimeToLive(v int32)`

SetOidcAccessTokenTimeToLive sets OidcAccessTokenTimeToLive field to given value.

### HasOidcAccessTokenTimeToLive

`func (o *VirtualSaml2ServiceProviderDTO) HasOidcAccessTokenTimeToLive() bool`

HasOidcAccessTokenTimeToLive returns a boolean if a field has been set.

### GetOidcAuthzCodeTimeToLive

`func (o *VirtualSaml2ServiceProviderDTO) GetOidcAuthzCodeTimeToLive() int32`

GetOidcAuthzCodeTimeToLive returns the OidcAuthzCodeTimeToLive field if non-nil, zero value otherwise.

### GetOidcAuthzCodeTimeToLiveOk

`func (o *VirtualSaml2ServiceProviderDTO) GetOidcAuthzCodeTimeToLiveOk() (*int32, bool)`

GetOidcAuthzCodeTimeToLiveOk returns a tuple with the OidcAuthzCodeTimeToLive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcAuthzCodeTimeToLive

`func (o *VirtualSaml2ServiceProviderDTO) SetOidcAuthzCodeTimeToLive(v int32)`

SetOidcAuthzCodeTimeToLive sets OidcAuthzCodeTimeToLive field to given value.

### HasOidcAuthzCodeTimeToLive

`func (o *VirtualSaml2ServiceProviderDTO) HasOidcAuthzCodeTimeToLive() bool`

HasOidcAuthzCodeTimeToLive returns a boolean if a field has been set.

### GetOidcIdTokenTimeToLive

`func (o *VirtualSaml2ServiceProviderDTO) GetOidcIdTokenTimeToLive() int32`

GetOidcIdTokenTimeToLive returns the OidcIdTokenTimeToLive field if non-nil, zero value otherwise.

### GetOidcIdTokenTimeToLiveOk

`func (o *VirtualSaml2ServiceProviderDTO) GetOidcIdTokenTimeToLiveOk() (*int32, bool)`

GetOidcIdTokenTimeToLiveOk returns a tuple with the OidcIdTokenTimeToLive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcIdTokenTimeToLive

`func (o *VirtualSaml2ServiceProviderDTO) SetOidcIdTokenTimeToLive(v int32)`

SetOidcIdTokenTimeToLive sets OidcIdTokenTimeToLive field to given value.

### HasOidcIdTokenTimeToLive

`func (o *VirtualSaml2ServiceProviderDTO) HasOidcIdTokenTimeToLive() bool`

HasOidcIdTokenTimeToLive returns a boolean if a field has been set.

### GetOpenIdEnabled

`func (o *VirtualSaml2ServiceProviderDTO) GetOpenIdEnabled() bool`

GetOpenIdEnabled returns the OpenIdEnabled field if non-nil, zero value otherwise.

### GetOpenIdEnabledOk

`func (o *VirtualSaml2ServiceProviderDTO) GetOpenIdEnabledOk() (*bool, bool)`

GetOpenIdEnabledOk returns a tuple with the OpenIdEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenIdEnabled

`func (o *VirtualSaml2ServiceProviderDTO) SetOpenIdEnabled(v bool)`

SetOpenIdEnabled sets OpenIdEnabled field to given value.

### HasOpenIdEnabled

`func (o *VirtualSaml2ServiceProviderDTO) HasOpenIdEnabled() bool`

HasOpenIdEnabled returns a boolean if a field has been set.

### GetRemote

`func (o *VirtualSaml2ServiceProviderDTO) GetRemote() bool`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *VirtualSaml2ServiceProviderDTO) GetRemoteOk() (*bool, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *VirtualSaml2ServiceProviderDTO) SetRemote(v bool)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *VirtualSaml2ServiceProviderDTO) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### GetRole

`func (o *VirtualSaml2ServiceProviderDTO) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *VirtualSaml2ServiceProviderDTO) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *VirtualSaml2ServiceProviderDTO) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *VirtualSaml2ServiceProviderDTO) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSessionManagerFactory

`func (o *VirtualSaml2ServiceProviderDTO) GetSessionManagerFactory() SessionManagerFactoryDTO`

GetSessionManagerFactory returns the SessionManagerFactory field if non-nil, zero value otherwise.

### GetSessionManagerFactoryOk

`func (o *VirtualSaml2ServiceProviderDTO) GetSessionManagerFactoryOk() (*SessionManagerFactoryDTO, bool)`

GetSessionManagerFactoryOk returns a tuple with the SessionManagerFactory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionManagerFactory

`func (o *VirtualSaml2ServiceProviderDTO) SetSessionManagerFactory(v SessionManagerFactoryDTO)`

SetSessionManagerFactory sets SessionManagerFactory field to given value.

### HasSessionManagerFactory

`func (o *VirtualSaml2ServiceProviderDTO) HasSessionManagerFactory() bool`

HasSessionManagerFactory returns a boolean if a field has been set.

### GetSignAuthenticationRequests

`func (o *VirtualSaml2ServiceProviderDTO) GetSignAuthenticationRequests() bool`

GetSignAuthenticationRequests returns the SignAuthenticationRequests field if non-nil, zero value otherwise.

### GetSignAuthenticationRequestsOk

`func (o *VirtualSaml2ServiceProviderDTO) GetSignAuthenticationRequestsOk() (*bool, bool)`

GetSignAuthenticationRequestsOk returns a tuple with the SignAuthenticationRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignAuthenticationRequests

`func (o *VirtualSaml2ServiceProviderDTO) SetSignAuthenticationRequests(v bool)`

SetSignAuthenticationRequests sets SignAuthenticationRequests field to given value.

### HasSignAuthenticationRequests

`func (o *VirtualSaml2ServiceProviderDTO) HasSignAuthenticationRequests() bool`

HasSignAuthenticationRequests returns a boolean if a field has been set.

### GetSignRequests

`func (o *VirtualSaml2ServiceProviderDTO) GetSignRequests() bool`

GetSignRequests returns the SignRequests field if non-nil, zero value otherwise.

### GetSignRequestsOk

`func (o *VirtualSaml2ServiceProviderDTO) GetSignRequestsOk() (*bool, bool)`

GetSignRequestsOk returns a tuple with the SignRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignRequests

`func (o *VirtualSaml2ServiceProviderDTO) SetSignRequests(v bool)`

SetSignRequests sets SignRequests field to given value.

### HasSignRequests

`func (o *VirtualSaml2ServiceProviderDTO) HasSignRequests() bool`

HasSignRequests returns a boolean if a field has been set.

### GetSpSignatureHash

`func (o *VirtualSaml2ServiceProviderDTO) GetSpSignatureHash() string`

GetSpSignatureHash returns the SpSignatureHash field if non-nil, zero value otherwise.

### GetSpSignatureHashOk

`func (o *VirtualSaml2ServiceProviderDTO) GetSpSignatureHashOk() (*string, bool)`

GetSpSignatureHashOk returns a tuple with the SpSignatureHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpSignatureHash

`func (o *VirtualSaml2ServiceProviderDTO) SetSpSignatureHash(v string)`

SetSpSignatureHash sets SpSignatureHash field to given value.

### HasSpSignatureHash

`func (o *VirtualSaml2ServiceProviderDTO) HasSpSignatureHash() bool`

HasSpSignatureHash returns a boolean if a field has been set.

### GetSsoSessionTimeout

`func (o *VirtualSaml2ServiceProviderDTO) GetSsoSessionTimeout() int32`

GetSsoSessionTimeout returns the SsoSessionTimeout field if non-nil, zero value otherwise.

### GetSsoSessionTimeoutOk

`func (o *VirtualSaml2ServiceProviderDTO) GetSsoSessionTimeoutOk() (*int32, bool)`

GetSsoSessionTimeoutOk returns a tuple with the SsoSessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoSessionTimeout

`func (o *VirtualSaml2ServiceProviderDTO) SetSsoSessionTimeout(v int32)`

SetSsoSessionTimeout sets SsoSessionTimeout field to given value.

### HasSsoSessionTimeout

`func (o *VirtualSaml2ServiceProviderDTO) HasSsoSessionTimeout() bool`

HasSsoSessionTimeout returns a boolean if a field has been set.

### GetSubjectAuthnPolicies

`func (o *VirtualSaml2ServiceProviderDTO) GetSubjectAuthnPolicies() []SubjectAuthenticationPolicyDTO`

GetSubjectAuthnPolicies returns the SubjectAuthnPolicies field if non-nil, zero value otherwise.

### GetSubjectAuthnPoliciesOk

`func (o *VirtualSaml2ServiceProviderDTO) GetSubjectAuthnPoliciesOk() (*[]SubjectAuthenticationPolicyDTO, bool)`

GetSubjectAuthnPoliciesOk returns a tuple with the SubjectAuthnPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectAuthnPolicies

`func (o *VirtualSaml2ServiceProviderDTO) SetSubjectAuthnPolicies(v []SubjectAuthenticationPolicyDTO)`

SetSubjectAuthnPolicies sets SubjectAuthnPolicies field to given value.

### HasSubjectAuthnPolicies

`func (o *VirtualSaml2ServiceProviderDTO) HasSubjectAuthnPolicies() bool`

HasSubjectAuthnPolicies returns a boolean if a field has been set.

### GetSubjectNameIDPolicy

`func (o *VirtualSaml2ServiceProviderDTO) GetSubjectNameIDPolicy() SubjectNameIdentifierPolicyDTO`

GetSubjectNameIDPolicy returns the SubjectNameIDPolicy field if non-nil, zero value otherwise.

### GetSubjectNameIDPolicyOk

`func (o *VirtualSaml2ServiceProviderDTO) GetSubjectNameIDPolicyOk() (*SubjectNameIdentifierPolicyDTO, bool)`

GetSubjectNameIDPolicyOk returns a tuple with the SubjectNameIDPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectNameIDPolicy

`func (o *VirtualSaml2ServiceProviderDTO) SetSubjectNameIDPolicy(v SubjectNameIdentifierPolicyDTO)`

SetSubjectNameIDPolicy sets SubjectNameIDPolicy field to given value.

### HasSubjectNameIDPolicy

`func (o *VirtualSaml2ServiceProviderDTO) HasSubjectNameIDPolicy() bool`

HasSubjectNameIDPolicy returns a boolean if a field has been set.

### GetWantAssertionSigned

`func (o *VirtualSaml2ServiceProviderDTO) GetWantAssertionSigned() bool`

GetWantAssertionSigned returns the WantAssertionSigned field if non-nil, zero value otherwise.

### GetWantAssertionSignedOk

`func (o *VirtualSaml2ServiceProviderDTO) GetWantAssertionSignedOk() (*bool, bool)`

GetWantAssertionSignedOk returns a tuple with the WantAssertionSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWantAssertionSigned

`func (o *VirtualSaml2ServiceProviderDTO) SetWantAssertionSigned(v bool)`

SetWantAssertionSigned sets WantAssertionSigned field to given value.

### HasWantAssertionSigned

`func (o *VirtualSaml2ServiceProviderDTO) HasWantAssertionSigned() bool`

HasWantAssertionSigned returns a boolean if a field has been set.

### GetWantAuthnRequestsSigned

`func (o *VirtualSaml2ServiceProviderDTO) GetWantAuthnRequestsSigned() bool`

GetWantAuthnRequestsSigned returns the WantAuthnRequestsSigned field if non-nil, zero value otherwise.

### GetWantAuthnRequestsSignedOk

`func (o *VirtualSaml2ServiceProviderDTO) GetWantAuthnRequestsSignedOk() (*bool, bool)`

GetWantAuthnRequestsSignedOk returns a tuple with the WantAuthnRequestsSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWantAuthnRequestsSigned

`func (o *VirtualSaml2ServiceProviderDTO) SetWantAuthnRequestsSigned(v bool)`

SetWantAuthnRequestsSigned sets WantAuthnRequestsSigned field to given value.

### HasWantAuthnRequestsSigned

`func (o *VirtualSaml2ServiceProviderDTO) HasWantAuthnRequestsSigned() bool`

HasWantAuthnRequestsSigned returns a boolean if a field has been set.

### GetWantSLOResponseSigned

`func (o *VirtualSaml2ServiceProviderDTO) GetWantSLOResponseSigned() bool`

GetWantSLOResponseSigned returns the WantSLOResponseSigned field if non-nil, zero value otherwise.

### GetWantSLOResponseSignedOk

`func (o *VirtualSaml2ServiceProviderDTO) GetWantSLOResponseSignedOk() (*bool, bool)`

GetWantSLOResponseSignedOk returns a tuple with the WantSLOResponseSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWantSLOResponseSigned

`func (o *VirtualSaml2ServiceProviderDTO) SetWantSLOResponseSigned(v bool)`

SetWantSLOResponseSigned sets WantSLOResponseSigned field to given value.

### HasWantSLOResponseSigned

`func (o *VirtualSaml2ServiceProviderDTO) HasWantSLOResponseSigned() bool`

HasWantSLOResponseSigned returns a boolean if a field has been set.

### GetWantSignedRequests

`func (o *VirtualSaml2ServiceProviderDTO) GetWantSignedRequests() bool`

GetWantSignedRequests returns the WantSignedRequests field if non-nil, zero value otherwise.

### GetWantSignedRequestsOk

`func (o *VirtualSaml2ServiceProviderDTO) GetWantSignedRequestsOk() (*bool, bool)`

GetWantSignedRequestsOk returns a tuple with the WantSignedRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWantSignedRequests

`func (o *VirtualSaml2ServiceProviderDTO) SetWantSignedRequests(v bool)`

SetWantSignedRequests sets WantSignedRequests field to given value.

### HasWantSignedRequests

`func (o *VirtualSaml2ServiceProviderDTO) HasWantSignedRequests() bool`

HasWantSignedRequests returns a boolean if a field has been set.

### GetX

`func (o *VirtualSaml2ServiceProviderDTO) GetX() float64`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *VirtualSaml2ServiceProviderDTO) GetXOk() (*float64, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *VirtualSaml2ServiceProviderDTO) SetX(v float64)`

SetX sets X field to given value.

### HasX

`func (o *VirtualSaml2ServiceProviderDTO) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *VirtualSaml2ServiceProviderDTO) GetY() float64`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *VirtualSaml2ServiceProviderDTO) GetYOk() (*float64, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *VirtualSaml2ServiceProviderDTO) SetY(v float64)`

SetY sets Y field to given value.

### HasY

`func (o *VirtualSaml2ServiceProviderDTO) HasY() bool`

HasY returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


