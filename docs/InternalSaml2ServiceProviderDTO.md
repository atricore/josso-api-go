# InternalSaml2ServiceProviderDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountLinkagePolicy** | Pointer to [**AccountLinkagePolicyDTO**](AccountLinkagePolicyDTO.md) |  | [optional] 
**ActiveBindings** | Pointer to **[]string** |  | [optional] 
**ActiveProfiles** | Pointer to **[]string** |  | [optional] 
**AuthenticationContract** | Pointer to [**AuthenticationContractDTO**](AuthenticationContractDTO.md) |  | [optional] 
**AuthenticationMechanism** | Pointer to [**AuthenticationMechanismDTO**](AuthenticationMechanismDTO.md) |  | [optional] 
**Config** | Pointer to [**ProviderConfigDTO**](ProviderConfigDTO.md) |  | [optional] 
**DashboardUrl** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**EnableMetadataEndpoint** | Pointer to **bool** |  | [optional] 
**ErrorBinding** | Pointer to **string** |  | [optional] 
**FederatedConnectionsA** | Pointer to [**[]FederatedConnectionDTO**](FederatedConnectionDTO.md) |  | [optional] 
**FederatedConnectionsB** | Pointer to [**[]FederatedConnectionDTO**](FederatedConnectionDTO.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**IdentityAppliance** | Pointer to [**IdentityApplianceDefinitionDTO**](IdentityApplianceDefinitionDTO.md) |  | [optional] 
**IdentityLookups** | Pointer to [**[]IdentityLookupDTO**](IdentityLookupDTO.md) |  | [optional] 
**IdentityMappingPolicy** | Pointer to [**IdentityMappingPolicyDTO**](IdentityMappingPolicyDTO.md) |  | [optional] 
**IsRemote** | Pointer to **bool** |  | [optional] 
**Location** | Pointer to [**LocationDTO**](LocationDTO.md) |  | [optional] 
**MessageTtl** | Pointer to **int32** |  | [optional] 
**MessageTtlTolerance** | Pointer to **int32** |  | [optional] 
**Metadata** | Pointer to [**ResourceDTO**](ResourceDTO.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Remote** | Pointer to **bool** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**ServiceConnection** | Pointer to [**ServiceConnectionDTO**](ServiceConnectionDTO.md) |  | [optional] 
**SignAuthenticationRequests** | Pointer to **bool** |  | [optional] 
**SignRequests** | Pointer to **bool** |  | [optional] 
**SignatureHash** | Pointer to **string** |  | [optional] 
**WantAssertionSigned** | Pointer to **bool** |  | [optional] 
**WantSLOResponseSigned** | Pointer to **bool** |  | [optional] 
**WantSignedRequests** | Pointer to **bool** |  | [optional] 
**X** | Pointer to **float64** |  | [optional] 
**Y** | Pointer to **float64** |  | [optional] 

## Methods

### NewInternalSaml2ServiceProviderDTO

`func NewInternalSaml2ServiceProviderDTO() *InternalSaml2ServiceProviderDTO`

NewInternalSaml2ServiceProviderDTO instantiates a new InternalSaml2ServiceProviderDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalSaml2ServiceProviderDTOWithDefaults

`func NewInternalSaml2ServiceProviderDTOWithDefaults() *InternalSaml2ServiceProviderDTO`

NewInternalSaml2ServiceProviderDTOWithDefaults instantiates a new InternalSaml2ServiceProviderDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountLinkagePolicy

`func (o *InternalSaml2ServiceProviderDTO) GetAccountLinkagePolicy() AccountLinkagePolicyDTO`

GetAccountLinkagePolicy returns the AccountLinkagePolicy field if non-nil, zero value otherwise.

### GetAccountLinkagePolicyOk

`func (o *InternalSaml2ServiceProviderDTO) GetAccountLinkagePolicyOk() (*AccountLinkagePolicyDTO, bool)`

GetAccountLinkagePolicyOk returns a tuple with the AccountLinkagePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountLinkagePolicy

`func (o *InternalSaml2ServiceProviderDTO) SetAccountLinkagePolicy(v AccountLinkagePolicyDTO)`

SetAccountLinkagePolicy sets AccountLinkagePolicy field to given value.

### HasAccountLinkagePolicy

`func (o *InternalSaml2ServiceProviderDTO) HasAccountLinkagePolicy() bool`

HasAccountLinkagePolicy returns a boolean if a field has been set.

### GetActiveBindings

`func (o *InternalSaml2ServiceProviderDTO) GetActiveBindings() []string`

GetActiveBindings returns the ActiveBindings field if non-nil, zero value otherwise.

### GetActiveBindingsOk

`func (o *InternalSaml2ServiceProviderDTO) GetActiveBindingsOk() (*[]string, bool)`

GetActiveBindingsOk returns a tuple with the ActiveBindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveBindings

`func (o *InternalSaml2ServiceProviderDTO) SetActiveBindings(v []string)`

SetActiveBindings sets ActiveBindings field to given value.

### HasActiveBindings

`func (o *InternalSaml2ServiceProviderDTO) HasActiveBindings() bool`

HasActiveBindings returns a boolean if a field has been set.

### GetActiveProfiles

`func (o *InternalSaml2ServiceProviderDTO) GetActiveProfiles() []string`

GetActiveProfiles returns the ActiveProfiles field if non-nil, zero value otherwise.

### GetActiveProfilesOk

`func (o *InternalSaml2ServiceProviderDTO) GetActiveProfilesOk() (*[]string, bool)`

GetActiveProfilesOk returns a tuple with the ActiveProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveProfiles

`func (o *InternalSaml2ServiceProviderDTO) SetActiveProfiles(v []string)`

SetActiveProfiles sets ActiveProfiles field to given value.

### HasActiveProfiles

`func (o *InternalSaml2ServiceProviderDTO) HasActiveProfiles() bool`

HasActiveProfiles returns a boolean if a field has been set.

### GetAuthenticationContract

`func (o *InternalSaml2ServiceProviderDTO) GetAuthenticationContract() AuthenticationContractDTO`

GetAuthenticationContract returns the AuthenticationContract field if non-nil, zero value otherwise.

### GetAuthenticationContractOk

`func (o *InternalSaml2ServiceProviderDTO) GetAuthenticationContractOk() (*AuthenticationContractDTO, bool)`

GetAuthenticationContractOk returns a tuple with the AuthenticationContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationContract

`func (o *InternalSaml2ServiceProviderDTO) SetAuthenticationContract(v AuthenticationContractDTO)`

SetAuthenticationContract sets AuthenticationContract field to given value.

### HasAuthenticationContract

`func (o *InternalSaml2ServiceProviderDTO) HasAuthenticationContract() bool`

HasAuthenticationContract returns a boolean if a field has been set.

### GetAuthenticationMechanism

`func (o *InternalSaml2ServiceProviderDTO) GetAuthenticationMechanism() AuthenticationMechanismDTO`

GetAuthenticationMechanism returns the AuthenticationMechanism field if non-nil, zero value otherwise.

### GetAuthenticationMechanismOk

`func (o *InternalSaml2ServiceProviderDTO) GetAuthenticationMechanismOk() (*AuthenticationMechanismDTO, bool)`

GetAuthenticationMechanismOk returns a tuple with the AuthenticationMechanism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMechanism

`func (o *InternalSaml2ServiceProviderDTO) SetAuthenticationMechanism(v AuthenticationMechanismDTO)`

SetAuthenticationMechanism sets AuthenticationMechanism field to given value.

### HasAuthenticationMechanism

`func (o *InternalSaml2ServiceProviderDTO) HasAuthenticationMechanism() bool`

HasAuthenticationMechanism returns a boolean if a field has been set.

### GetConfig

`func (o *InternalSaml2ServiceProviderDTO) GetConfig() ProviderConfigDTO`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *InternalSaml2ServiceProviderDTO) GetConfigOk() (*ProviderConfigDTO, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *InternalSaml2ServiceProviderDTO) SetConfig(v ProviderConfigDTO)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *InternalSaml2ServiceProviderDTO) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetDashboardUrl

`func (o *InternalSaml2ServiceProviderDTO) GetDashboardUrl() string`

GetDashboardUrl returns the DashboardUrl field if non-nil, zero value otherwise.

### GetDashboardUrlOk

`func (o *InternalSaml2ServiceProviderDTO) GetDashboardUrlOk() (*string, bool)`

GetDashboardUrlOk returns a tuple with the DashboardUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardUrl

`func (o *InternalSaml2ServiceProviderDTO) SetDashboardUrl(v string)`

SetDashboardUrl sets DashboardUrl field to given value.

### HasDashboardUrl

`func (o *InternalSaml2ServiceProviderDTO) HasDashboardUrl() bool`

HasDashboardUrl returns a boolean if a field has been set.

### GetDescription

`func (o *InternalSaml2ServiceProviderDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InternalSaml2ServiceProviderDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InternalSaml2ServiceProviderDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InternalSaml2ServiceProviderDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *InternalSaml2ServiceProviderDTO) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *InternalSaml2ServiceProviderDTO) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *InternalSaml2ServiceProviderDTO) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *InternalSaml2ServiceProviderDTO) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetElementId

`func (o *InternalSaml2ServiceProviderDTO) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *InternalSaml2ServiceProviderDTO) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *InternalSaml2ServiceProviderDTO) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *InternalSaml2ServiceProviderDTO) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetEnableMetadataEndpoint

`func (o *InternalSaml2ServiceProviderDTO) GetEnableMetadataEndpoint() bool`

GetEnableMetadataEndpoint returns the EnableMetadataEndpoint field if non-nil, zero value otherwise.

### GetEnableMetadataEndpointOk

`func (o *InternalSaml2ServiceProviderDTO) GetEnableMetadataEndpointOk() (*bool, bool)`

GetEnableMetadataEndpointOk returns a tuple with the EnableMetadataEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMetadataEndpoint

`func (o *InternalSaml2ServiceProviderDTO) SetEnableMetadataEndpoint(v bool)`

SetEnableMetadataEndpoint sets EnableMetadataEndpoint field to given value.

### HasEnableMetadataEndpoint

`func (o *InternalSaml2ServiceProviderDTO) HasEnableMetadataEndpoint() bool`

HasEnableMetadataEndpoint returns a boolean if a field has been set.

### GetErrorBinding

`func (o *InternalSaml2ServiceProviderDTO) GetErrorBinding() string`

GetErrorBinding returns the ErrorBinding field if non-nil, zero value otherwise.

### GetErrorBindingOk

`func (o *InternalSaml2ServiceProviderDTO) GetErrorBindingOk() (*string, bool)`

GetErrorBindingOk returns a tuple with the ErrorBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorBinding

`func (o *InternalSaml2ServiceProviderDTO) SetErrorBinding(v string)`

SetErrorBinding sets ErrorBinding field to given value.

### HasErrorBinding

`func (o *InternalSaml2ServiceProviderDTO) HasErrorBinding() bool`

HasErrorBinding returns a boolean if a field has been set.

### GetFederatedConnectionsA

`func (o *InternalSaml2ServiceProviderDTO) GetFederatedConnectionsA() []FederatedConnectionDTO`

GetFederatedConnectionsA returns the FederatedConnectionsA field if non-nil, zero value otherwise.

### GetFederatedConnectionsAOk

`func (o *InternalSaml2ServiceProviderDTO) GetFederatedConnectionsAOk() (*[]FederatedConnectionDTO, bool)`

GetFederatedConnectionsAOk returns a tuple with the FederatedConnectionsA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederatedConnectionsA

`func (o *InternalSaml2ServiceProviderDTO) SetFederatedConnectionsA(v []FederatedConnectionDTO)`

SetFederatedConnectionsA sets FederatedConnectionsA field to given value.

### HasFederatedConnectionsA

`func (o *InternalSaml2ServiceProviderDTO) HasFederatedConnectionsA() bool`

HasFederatedConnectionsA returns a boolean if a field has been set.

### GetFederatedConnectionsB

`func (o *InternalSaml2ServiceProviderDTO) GetFederatedConnectionsB() []FederatedConnectionDTO`

GetFederatedConnectionsB returns the FederatedConnectionsB field if non-nil, zero value otherwise.

### GetFederatedConnectionsBOk

`func (o *InternalSaml2ServiceProviderDTO) GetFederatedConnectionsBOk() (*[]FederatedConnectionDTO, bool)`

GetFederatedConnectionsBOk returns a tuple with the FederatedConnectionsB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederatedConnectionsB

`func (o *InternalSaml2ServiceProviderDTO) SetFederatedConnectionsB(v []FederatedConnectionDTO)`

SetFederatedConnectionsB sets FederatedConnectionsB field to given value.

### HasFederatedConnectionsB

`func (o *InternalSaml2ServiceProviderDTO) HasFederatedConnectionsB() bool`

HasFederatedConnectionsB returns a boolean if a field has been set.

### GetId

`func (o *InternalSaml2ServiceProviderDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InternalSaml2ServiceProviderDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InternalSaml2ServiceProviderDTO) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InternalSaml2ServiceProviderDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentityAppliance

`func (o *InternalSaml2ServiceProviderDTO) GetIdentityAppliance() IdentityApplianceDefinitionDTO`

GetIdentityAppliance returns the IdentityAppliance field if non-nil, zero value otherwise.

### GetIdentityApplianceOk

`func (o *InternalSaml2ServiceProviderDTO) GetIdentityApplianceOk() (*IdentityApplianceDefinitionDTO, bool)`

GetIdentityApplianceOk returns a tuple with the IdentityAppliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityAppliance

`func (o *InternalSaml2ServiceProviderDTO) SetIdentityAppliance(v IdentityApplianceDefinitionDTO)`

SetIdentityAppliance sets IdentityAppliance field to given value.

### HasIdentityAppliance

`func (o *InternalSaml2ServiceProviderDTO) HasIdentityAppliance() bool`

HasIdentityAppliance returns a boolean if a field has been set.

### GetIdentityLookups

`func (o *InternalSaml2ServiceProviderDTO) GetIdentityLookups() []IdentityLookupDTO`

GetIdentityLookups returns the IdentityLookups field if non-nil, zero value otherwise.

### GetIdentityLookupsOk

`func (o *InternalSaml2ServiceProviderDTO) GetIdentityLookupsOk() (*[]IdentityLookupDTO, bool)`

GetIdentityLookupsOk returns a tuple with the IdentityLookups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityLookups

`func (o *InternalSaml2ServiceProviderDTO) SetIdentityLookups(v []IdentityLookupDTO)`

SetIdentityLookups sets IdentityLookups field to given value.

### HasIdentityLookups

`func (o *InternalSaml2ServiceProviderDTO) HasIdentityLookups() bool`

HasIdentityLookups returns a boolean if a field has been set.

### GetIdentityMappingPolicy

`func (o *InternalSaml2ServiceProviderDTO) GetIdentityMappingPolicy() IdentityMappingPolicyDTO`

GetIdentityMappingPolicy returns the IdentityMappingPolicy field if non-nil, zero value otherwise.

### GetIdentityMappingPolicyOk

`func (o *InternalSaml2ServiceProviderDTO) GetIdentityMappingPolicyOk() (*IdentityMappingPolicyDTO, bool)`

GetIdentityMappingPolicyOk returns a tuple with the IdentityMappingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityMappingPolicy

`func (o *InternalSaml2ServiceProviderDTO) SetIdentityMappingPolicy(v IdentityMappingPolicyDTO)`

SetIdentityMappingPolicy sets IdentityMappingPolicy field to given value.

### HasIdentityMappingPolicy

`func (o *InternalSaml2ServiceProviderDTO) HasIdentityMappingPolicy() bool`

HasIdentityMappingPolicy returns a boolean if a field has been set.

### GetIsRemote

`func (o *InternalSaml2ServiceProviderDTO) GetIsRemote() bool`

GetIsRemote returns the IsRemote field if non-nil, zero value otherwise.

### GetIsRemoteOk

`func (o *InternalSaml2ServiceProviderDTO) GetIsRemoteOk() (*bool, bool)`

GetIsRemoteOk returns a tuple with the IsRemote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemote

`func (o *InternalSaml2ServiceProviderDTO) SetIsRemote(v bool)`

SetIsRemote sets IsRemote field to given value.

### HasIsRemote

`func (o *InternalSaml2ServiceProviderDTO) HasIsRemote() bool`

HasIsRemote returns a boolean if a field has been set.

### GetLocation

`func (o *InternalSaml2ServiceProviderDTO) GetLocation() LocationDTO`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *InternalSaml2ServiceProviderDTO) GetLocationOk() (*LocationDTO, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *InternalSaml2ServiceProviderDTO) SetLocation(v LocationDTO)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *InternalSaml2ServiceProviderDTO) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetMessageTtl

`func (o *InternalSaml2ServiceProviderDTO) GetMessageTtl() int32`

GetMessageTtl returns the MessageTtl field if non-nil, zero value otherwise.

### GetMessageTtlOk

`func (o *InternalSaml2ServiceProviderDTO) GetMessageTtlOk() (*int32, bool)`

GetMessageTtlOk returns a tuple with the MessageTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageTtl

`func (o *InternalSaml2ServiceProviderDTO) SetMessageTtl(v int32)`

SetMessageTtl sets MessageTtl field to given value.

### HasMessageTtl

`func (o *InternalSaml2ServiceProviderDTO) HasMessageTtl() bool`

HasMessageTtl returns a boolean if a field has been set.

### GetMessageTtlTolerance

`func (o *InternalSaml2ServiceProviderDTO) GetMessageTtlTolerance() int32`

GetMessageTtlTolerance returns the MessageTtlTolerance field if non-nil, zero value otherwise.

### GetMessageTtlToleranceOk

`func (o *InternalSaml2ServiceProviderDTO) GetMessageTtlToleranceOk() (*int32, bool)`

GetMessageTtlToleranceOk returns a tuple with the MessageTtlTolerance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageTtlTolerance

`func (o *InternalSaml2ServiceProviderDTO) SetMessageTtlTolerance(v int32)`

SetMessageTtlTolerance sets MessageTtlTolerance field to given value.

### HasMessageTtlTolerance

`func (o *InternalSaml2ServiceProviderDTO) HasMessageTtlTolerance() bool`

HasMessageTtlTolerance returns a boolean if a field has been set.

### GetMetadata

`func (o *InternalSaml2ServiceProviderDTO) GetMetadata() ResourceDTO`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *InternalSaml2ServiceProviderDTO) GetMetadataOk() (*ResourceDTO, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *InternalSaml2ServiceProviderDTO) SetMetadata(v ResourceDTO)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *InternalSaml2ServiceProviderDTO) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *InternalSaml2ServiceProviderDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InternalSaml2ServiceProviderDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InternalSaml2ServiceProviderDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InternalSaml2ServiceProviderDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRemote

`func (o *InternalSaml2ServiceProviderDTO) GetRemote() bool`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *InternalSaml2ServiceProviderDTO) GetRemoteOk() (*bool, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *InternalSaml2ServiceProviderDTO) SetRemote(v bool)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *InternalSaml2ServiceProviderDTO) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### GetRole

`func (o *InternalSaml2ServiceProviderDTO) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *InternalSaml2ServiceProviderDTO) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *InternalSaml2ServiceProviderDTO) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *InternalSaml2ServiceProviderDTO) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetServiceConnection

`func (o *InternalSaml2ServiceProviderDTO) GetServiceConnection() ServiceConnectionDTO`

GetServiceConnection returns the ServiceConnection field if non-nil, zero value otherwise.

### GetServiceConnectionOk

`func (o *InternalSaml2ServiceProviderDTO) GetServiceConnectionOk() (*ServiceConnectionDTO, bool)`

GetServiceConnectionOk returns a tuple with the ServiceConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceConnection

`func (o *InternalSaml2ServiceProviderDTO) SetServiceConnection(v ServiceConnectionDTO)`

SetServiceConnection sets ServiceConnection field to given value.

### HasServiceConnection

`func (o *InternalSaml2ServiceProviderDTO) HasServiceConnection() bool`

HasServiceConnection returns a boolean if a field has been set.

### GetSignAuthenticationRequests

`func (o *InternalSaml2ServiceProviderDTO) GetSignAuthenticationRequests() bool`

GetSignAuthenticationRequests returns the SignAuthenticationRequests field if non-nil, zero value otherwise.

### GetSignAuthenticationRequestsOk

`func (o *InternalSaml2ServiceProviderDTO) GetSignAuthenticationRequestsOk() (*bool, bool)`

GetSignAuthenticationRequestsOk returns a tuple with the SignAuthenticationRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignAuthenticationRequests

`func (o *InternalSaml2ServiceProviderDTO) SetSignAuthenticationRequests(v bool)`

SetSignAuthenticationRequests sets SignAuthenticationRequests field to given value.

### HasSignAuthenticationRequests

`func (o *InternalSaml2ServiceProviderDTO) HasSignAuthenticationRequests() bool`

HasSignAuthenticationRequests returns a boolean if a field has been set.

### GetSignRequests

`func (o *InternalSaml2ServiceProviderDTO) GetSignRequests() bool`

GetSignRequests returns the SignRequests field if non-nil, zero value otherwise.

### GetSignRequestsOk

`func (o *InternalSaml2ServiceProviderDTO) GetSignRequestsOk() (*bool, bool)`

GetSignRequestsOk returns a tuple with the SignRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignRequests

`func (o *InternalSaml2ServiceProviderDTO) SetSignRequests(v bool)`

SetSignRequests sets SignRequests field to given value.

### HasSignRequests

`func (o *InternalSaml2ServiceProviderDTO) HasSignRequests() bool`

HasSignRequests returns a boolean if a field has been set.

### GetSignatureHash

`func (o *InternalSaml2ServiceProviderDTO) GetSignatureHash() string`

GetSignatureHash returns the SignatureHash field if non-nil, zero value otherwise.

### GetSignatureHashOk

`func (o *InternalSaml2ServiceProviderDTO) GetSignatureHashOk() (*string, bool)`

GetSignatureHashOk returns a tuple with the SignatureHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureHash

`func (o *InternalSaml2ServiceProviderDTO) SetSignatureHash(v string)`

SetSignatureHash sets SignatureHash field to given value.

### HasSignatureHash

`func (o *InternalSaml2ServiceProviderDTO) HasSignatureHash() bool`

HasSignatureHash returns a boolean if a field has been set.

### GetWantAssertionSigned

`func (o *InternalSaml2ServiceProviderDTO) GetWantAssertionSigned() bool`

GetWantAssertionSigned returns the WantAssertionSigned field if non-nil, zero value otherwise.

### GetWantAssertionSignedOk

`func (o *InternalSaml2ServiceProviderDTO) GetWantAssertionSignedOk() (*bool, bool)`

GetWantAssertionSignedOk returns a tuple with the WantAssertionSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWantAssertionSigned

`func (o *InternalSaml2ServiceProviderDTO) SetWantAssertionSigned(v bool)`

SetWantAssertionSigned sets WantAssertionSigned field to given value.

### HasWantAssertionSigned

`func (o *InternalSaml2ServiceProviderDTO) HasWantAssertionSigned() bool`

HasWantAssertionSigned returns a boolean if a field has been set.

### GetWantSLOResponseSigned

`func (o *InternalSaml2ServiceProviderDTO) GetWantSLOResponseSigned() bool`

GetWantSLOResponseSigned returns the WantSLOResponseSigned field if non-nil, zero value otherwise.

### GetWantSLOResponseSignedOk

`func (o *InternalSaml2ServiceProviderDTO) GetWantSLOResponseSignedOk() (*bool, bool)`

GetWantSLOResponseSignedOk returns a tuple with the WantSLOResponseSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWantSLOResponseSigned

`func (o *InternalSaml2ServiceProviderDTO) SetWantSLOResponseSigned(v bool)`

SetWantSLOResponseSigned sets WantSLOResponseSigned field to given value.

### HasWantSLOResponseSigned

`func (o *InternalSaml2ServiceProviderDTO) HasWantSLOResponseSigned() bool`

HasWantSLOResponseSigned returns a boolean if a field has been set.

### GetWantSignedRequests

`func (o *InternalSaml2ServiceProviderDTO) GetWantSignedRequests() bool`

GetWantSignedRequests returns the WantSignedRequests field if non-nil, zero value otherwise.

### GetWantSignedRequestsOk

`func (o *InternalSaml2ServiceProviderDTO) GetWantSignedRequestsOk() (*bool, bool)`

GetWantSignedRequestsOk returns a tuple with the WantSignedRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWantSignedRequests

`func (o *InternalSaml2ServiceProviderDTO) SetWantSignedRequests(v bool)`

SetWantSignedRequests sets WantSignedRequests field to given value.

### HasWantSignedRequests

`func (o *InternalSaml2ServiceProviderDTO) HasWantSignedRequests() bool`

HasWantSignedRequests returns a boolean if a field has been set.

### GetX

`func (o *InternalSaml2ServiceProviderDTO) GetX() float64`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *InternalSaml2ServiceProviderDTO) GetXOk() (*float64, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *InternalSaml2ServiceProviderDTO) SetX(v float64)`

SetX sets X field to given value.

### HasX

`func (o *InternalSaml2ServiceProviderDTO) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *InternalSaml2ServiceProviderDTO) GetY() float64`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *InternalSaml2ServiceProviderDTO) GetYOk() (*float64, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *InternalSaml2ServiceProviderDTO) SetY(v float64)`

SetY sets Y field to given value.

### HasY

`func (o *InternalSaml2ServiceProviderDTO) HasY() bool`

HasY returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


