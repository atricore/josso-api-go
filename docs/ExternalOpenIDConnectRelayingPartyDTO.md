# ExternalOpenIDConnectRelayingPartyDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountLinkagePolicy** | Pointer to [**AccountLinkagePolicyDTO**](AccountLinkagePolicyDTO.md) |  | [optional] 
**ActiveBindings** | Pointer to **[]string** |  | [optional] 
**ActiveProfiles** | Pointer to **[]string** |  | [optional] 
**AuthorizedURIs** | Pointer to **[]string** |  | [optional] 
**ClientAuthnMethod** | Pointer to **string** |  | [optional] 
**ClientCert** | Pointer to **string** |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**ClientSecret** | Pointer to **string** |  | [optional] 
**ClientType** | Pointer to **string** |  | [optional] 
**Config** | Pointer to [**ProviderConfigDTO**](ProviderConfigDTO.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**EncryptionAlg** | Pointer to **string** |  | [optional] 
**EncryptionMethod** | Pointer to **string** |  | [optional] 
**FederatedConnectionsA** | Pointer to [**[]FederatedConnectionDTO**](FederatedConnectionDTO.md) |  | [optional] 
**FederatedConnectionsB** | Pointer to [**[]FederatedConnectionDTO**](FederatedConnectionDTO.md) |  | [optional] 
**Grants** | Pointer to **[]string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**IdTokenEncryptionAlg** | Pointer to **string** |  | [optional] 
**IdTokenEncryptionMethod** | Pointer to **string** |  | [optional] 
**IdTokenSigningAlg** | Pointer to **string** |  | [optional] 
**IdentityAppliance** | Pointer to [**IdentityApplianceDefinitionDTO**](IdentityApplianceDefinitionDTO.md) |  | [optional] 
**IdentityLookups** | Pointer to [**[]IdentityLookupDTO**](IdentityLookupDTO.md) |  | [optional] 
**IdentityMappingPolicy** | Pointer to [**IdentityMappingPolicyDTO**](IdentityMappingPolicyDTO.md) |  | [optional] 
**IsRemote** | Pointer to **bool** |  | [optional] 
**Location** | Pointer to [**LocationDTO**](LocationDTO.md) |  | [optional] 
**Metadata** | Pointer to [**ResourceDTO**](ResourceDTO.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PostLogoutRedirectionURIs** | Pointer to **[]string** |  | [optional] 
**Remote** | Pointer to **bool** |  | [optional] 
**ResponseTypes** | Pointer to **[]string** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**SigningAlg** | Pointer to **string** |  | [optional] 
**X** | Pointer to **float64** |  | [optional] 
**Y** | Pointer to **float64** |  | [optional] 

## Methods

### NewExternalOpenIDConnectRelayingPartyDTO

`func NewExternalOpenIDConnectRelayingPartyDTO() *ExternalOpenIDConnectRelayingPartyDTO`

NewExternalOpenIDConnectRelayingPartyDTO instantiates a new ExternalOpenIDConnectRelayingPartyDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalOpenIDConnectRelayingPartyDTOWithDefaults

`func NewExternalOpenIDConnectRelayingPartyDTOWithDefaults() *ExternalOpenIDConnectRelayingPartyDTO`

NewExternalOpenIDConnectRelayingPartyDTOWithDefaults instantiates a new ExternalOpenIDConnectRelayingPartyDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountLinkagePolicy

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetAccountLinkagePolicy() AccountLinkagePolicyDTO`

GetAccountLinkagePolicy returns the AccountLinkagePolicy field if non-nil, zero value otherwise.

### GetAccountLinkagePolicyOk

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetAccountLinkagePolicyOk() (*AccountLinkagePolicyDTO, bool)`

GetAccountLinkagePolicyOk returns a tuple with the AccountLinkagePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountLinkagePolicy

`func (o *ExternalOpenIDConnectRelayingPartyDTO) SetAccountLinkagePolicy(v AccountLinkagePolicyDTO)`

SetAccountLinkagePolicy sets AccountLinkagePolicy field to given value.

### HasAccountLinkagePolicy

`func (o *ExternalOpenIDConnectRelayingPartyDTO) HasAccountLinkagePolicy() bool`

HasAccountLinkagePolicy returns a boolean if a field has been set.

### GetActiveBindings

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetActiveBindings() []string`

GetActiveBindings returns the ActiveBindings field if non-nil, zero value otherwise.

### GetActiveBindingsOk

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetActiveBindingsOk() (*[]string, bool)`

GetActiveBindingsOk returns a tuple with the ActiveBindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveBindings

`func (o *ExternalOpenIDConnectRelayingPartyDTO) SetActiveBindings(v []string)`

SetActiveBindings sets ActiveBindings field to given value.

### HasActiveBindings

`func (o *ExternalOpenIDConnectRelayingPartyDTO) HasActiveBindings() bool`

HasActiveBindings returns a boolean if a field has been set.

### GetActiveProfiles

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetActiveProfiles() []string`

GetActiveProfiles returns the ActiveProfiles field if non-nil, zero value otherwise.

### GetActiveProfilesOk

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetActiveProfilesOk() (*[]string, bool)`

GetActiveProfilesOk returns a tuple with the ActiveProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveProfiles

`func (o *ExternalOpenIDConnectRelayingPartyDTO) SetActiveProfiles(v []string)`

SetActiveProfiles sets ActiveProfiles field to given value.

### HasActiveProfiles

`func (o *ExternalOpenIDConnectRelayingPartyDTO) HasActiveProfiles() bool`

HasActiveProfiles returns a boolean if a field has been set.

### GetAuthorizedURIs

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetAuthorizedURIs() []string`

GetAuthorizedURIs returns the AuthorizedURIs field if non-nil, zero value otherwise.

### GetAuthorizedURIsOk

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetAuthorizedURIsOk() (*[]string, bool)`

GetAuthorizedURIsOk returns a tuple with the AuthorizedURIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedURIs

`func (o *ExternalOpenIDConnectRelayingPartyDTO) SetAuthorizedURIs(v []string)`

SetAuthorizedURIs sets AuthorizedURIs field to given value.

### HasAuthorizedURIs

`func (o *ExternalOpenIDConnectRelayingPartyDTO) HasAuthorizedURIs() bool`

HasAuthorizedURIs returns a boolean if a field has been set.

### GetClientAuthnMethod

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetClientAuthnMethod() string`

GetClientAuthnMethod returns the ClientAuthnMethod field if non-nil, zero value otherwise.

### GetClientAuthnMethodOk

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetClientAuthnMethodOk() (*string, bool)`

GetClientAuthnMethodOk returns a tuple with the ClientAuthnMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAuthnMethod

`func (o *ExternalOpenIDConnectRelayingPartyDTO) SetClientAuthnMethod(v string)`

SetClientAuthnMethod sets ClientAuthnMethod field to given value.

### HasClientAuthnMethod

`func (o *ExternalOpenIDConnectRelayingPartyDTO) HasClientAuthnMethod() bool`

HasClientAuthnMethod returns a boolean if a field has been set.

### GetClientCert

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetClientCert() string`

GetClientCert returns the ClientCert field if non-nil, zero value otherwise.

### GetClientCertOk

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetClientCertOk() (*string, bool)`

GetClientCertOk returns a tuple with the ClientCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCert

`func (o *ExternalOpenIDConnectRelayingPartyDTO) SetClientCert(v string)`

SetClientCert sets ClientCert field to given value.

### HasClientCert

`func (o *ExternalOpenIDConnectRelayingPartyDTO) HasClientCert() bool`

HasClientCert returns a boolean if a field has been set.

### GetClientId

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ExternalOpenIDConnectRelayingPartyDTO) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *ExternalOpenIDConnectRelayingPartyDTO) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *ExternalOpenIDConnectRelayingPartyDTO) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *ExternalOpenIDConnectRelayingPartyDTO) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetClientType

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetClientType() string`

GetClientType returns the ClientType field if non-nil, zero value otherwise.

### GetClientTypeOk

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetClientTypeOk() (*string, bool)`

GetClientTypeOk returns a tuple with the ClientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientType

`func (o *ExternalOpenIDConnectRelayingPartyDTO) SetClientType(v string)`

SetClientType sets ClientType field to given value.

### HasClientType

`func (o *ExternalOpenIDConnectRelayingPartyDTO) HasClientType() bool`

HasClientType returns a boolean if a field has been set.

### GetConfig

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetConfig() ProviderConfigDTO`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetConfigOk() (*ProviderConfigDTO, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ExternalOpenIDConnectRelayingPartyDTO) SetConfig(v ProviderConfigDTO)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ExternalOpenIDConnectRelayingPartyDTO) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetDescription

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExternalOpenIDConnectRelayingPartyDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExternalOpenIDConnectRelayingPartyDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ExternalOpenIDConnectRelayingPartyDTO) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ExternalOpenIDConnectRelayingPartyDTO) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetElementId

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *ExternalOpenIDConnectRelayingPartyDTO) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *ExternalOpenIDConnectRelayingPartyDTO) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetEncryptionAlg

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetEncryptionAlg() string`

GetEncryptionAlg returns the EncryptionAlg field if non-nil, zero value otherwise.

### GetEncryptionAlgOk

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetEncryptionAlgOk() (*string, bool)`

GetEncryptionAlgOk returns a tuple with the EncryptionAlg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionAlg

`func (o *ExternalOpenIDConnectRelayingPartyDTO) SetEncryptionAlg(v string)`

SetEncryptionAlg sets EncryptionAlg field to given value.

### HasEncryptionAlg

`func (o *ExternalOpenIDConnectRelayingPartyDTO) HasEncryptionAlg() bool`

HasEncryptionAlg returns a boolean if a field has been set.

### GetEncryptionMethod

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetEncryptionMethod() string`

GetEncryptionMethod returns the EncryptionMethod field if non-nil, zero value otherwise.

### GetEncryptionMethodOk

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetEncryptionMethodOk() (*string, bool)`

GetEncryptionMethodOk returns a tuple with the EncryptionMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionMethod

`func (o *ExternalOpenIDConnectRelayingPartyDTO) SetEncryptionMethod(v string)`

SetEncryptionMethod sets EncryptionMethod field to given value.

### HasEncryptionMethod

`func (o *ExternalOpenIDConnectRelayingPartyDTO) HasEncryptionMethod() bool`

HasEncryptionMethod returns a boolean if a field has been set.

### GetFederatedConnectionsA

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetFederatedConnectionsA() []FederatedConnectionDTO`

GetFederatedConnectionsA returns the FederatedConnectionsA field if non-nil, zero value otherwise.

### GetFederatedConnectionsAOk

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetFederatedConnectionsAOk() (*[]FederatedConnectionDTO, bool)`

GetFederatedConnectionsAOk returns a tuple with the FederatedConnectionsA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederatedConnectionsA

`func (o *ExternalOpenIDConnectRelayingPartyDTO) SetFederatedConnectionsA(v []FederatedConnectionDTO)`

SetFederatedConnectionsA sets FederatedConnectionsA field to given value.

### HasFederatedConnectionsA

`func (o *ExternalOpenIDConnectRelayingPartyDTO) HasFederatedConnectionsA() bool`

HasFederatedConnectionsA returns a boolean if a field has been set.

### GetFederatedConnectionsB

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetFederatedConnectionsB() []FederatedConnectionDTO`

GetFederatedConnectionsB returns the FederatedConnectionsB field if non-nil, zero value otherwise.

### GetFederatedConnectionsBOk

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetFederatedConnectionsBOk() (*[]FederatedConnectionDTO, bool)`

GetFederatedConnectionsBOk returns a tuple with the FederatedConnectionsB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederatedConnectionsB

`func (o *ExternalOpenIDConnectRelayingPartyDTO) SetFederatedConnectionsB(v []FederatedConnectionDTO)`

SetFederatedConnectionsB sets FederatedConnectionsB field to given value.

### HasFederatedConnectionsB

`func (o *ExternalOpenIDConnectRelayingPartyDTO) HasFederatedConnectionsB() bool`

HasFederatedConnectionsB returns a boolean if a field has been set.

### GetGrants

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetGrants() []string`

GetGrants returns the Grants field if non-nil, zero value otherwise.

### GetGrantsOk

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetGrantsOk() (*[]string, bool)`

GetGrantsOk returns a tuple with the Grants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrants

`func (o *ExternalOpenIDConnectRelayingPartyDTO) SetGrants(v []string)`

SetGrants sets Grants field to given value.

### HasGrants

`func (o *ExternalOpenIDConnectRelayingPartyDTO) HasGrants() bool`

HasGrants returns a boolean if a field has been set.

### GetId

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalOpenIDConnectRelayingPartyDTO) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ExternalOpenIDConnectRelayingPartyDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdTokenEncryptionAlg

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetIdTokenEncryptionAlg() string`

GetIdTokenEncryptionAlg returns the IdTokenEncryptionAlg field if non-nil, zero value otherwise.

### GetIdTokenEncryptionAlgOk

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetIdTokenEncryptionAlgOk() (*string, bool)`

GetIdTokenEncryptionAlgOk returns a tuple with the IdTokenEncryptionAlg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTokenEncryptionAlg

`func (o *ExternalOpenIDConnectRelayingPartyDTO) SetIdTokenEncryptionAlg(v string)`

SetIdTokenEncryptionAlg sets IdTokenEncryptionAlg field to given value.

### HasIdTokenEncryptionAlg

`func (o *ExternalOpenIDConnectRelayingPartyDTO) HasIdTokenEncryptionAlg() bool`

HasIdTokenEncryptionAlg returns a boolean if a field has been set.

### GetIdTokenEncryptionMethod

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetIdTokenEncryptionMethod() string`

GetIdTokenEncryptionMethod returns the IdTokenEncryptionMethod field if non-nil, zero value otherwise.

### GetIdTokenEncryptionMethodOk

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetIdTokenEncryptionMethodOk() (*string, bool)`

GetIdTokenEncryptionMethodOk returns a tuple with the IdTokenEncryptionMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTokenEncryptionMethod

`func (o *ExternalOpenIDConnectRelayingPartyDTO) SetIdTokenEncryptionMethod(v string)`

SetIdTokenEncryptionMethod sets IdTokenEncryptionMethod field to given value.

### HasIdTokenEncryptionMethod

`func (o *ExternalOpenIDConnectRelayingPartyDTO) HasIdTokenEncryptionMethod() bool`

HasIdTokenEncryptionMethod returns a boolean if a field has been set.

### GetIdTokenSigningAlg

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetIdTokenSigningAlg() string`

GetIdTokenSigningAlg returns the IdTokenSigningAlg field if non-nil, zero value otherwise.

### GetIdTokenSigningAlgOk

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetIdTokenSigningAlgOk() (*string, bool)`

GetIdTokenSigningAlgOk returns a tuple with the IdTokenSigningAlg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTokenSigningAlg

`func (o *ExternalOpenIDConnectRelayingPartyDTO) SetIdTokenSigningAlg(v string)`

SetIdTokenSigningAlg sets IdTokenSigningAlg field to given value.

### HasIdTokenSigningAlg

`func (o *ExternalOpenIDConnectRelayingPartyDTO) HasIdTokenSigningAlg() bool`

HasIdTokenSigningAlg returns a boolean if a field has been set.

### GetIdentityAppliance

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetIdentityAppliance() IdentityApplianceDefinitionDTO`

GetIdentityAppliance returns the IdentityAppliance field if non-nil, zero value otherwise.

### GetIdentityApplianceOk

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetIdentityApplianceOk() (*IdentityApplianceDefinitionDTO, bool)`

GetIdentityApplianceOk returns a tuple with the IdentityAppliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityAppliance

`func (o *ExternalOpenIDConnectRelayingPartyDTO) SetIdentityAppliance(v IdentityApplianceDefinitionDTO)`

SetIdentityAppliance sets IdentityAppliance field to given value.

### HasIdentityAppliance

`func (o *ExternalOpenIDConnectRelayingPartyDTO) HasIdentityAppliance() bool`

HasIdentityAppliance returns a boolean if a field has been set.

### GetIdentityLookups

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetIdentityLookups() []IdentityLookupDTO`

GetIdentityLookups returns the IdentityLookups field if non-nil, zero value otherwise.

### GetIdentityLookupsOk

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetIdentityLookupsOk() (*[]IdentityLookupDTO, bool)`

GetIdentityLookupsOk returns a tuple with the IdentityLookups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityLookups

`func (o *ExternalOpenIDConnectRelayingPartyDTO) SetIdentityLookups(v []IdentityLookupDTO)`

SetIdentityLookups sets IdentityLookups field to given value.

### HasIdentityLookups

`func (o *ExternalOpenIDConnectRelayingPartyDTO) HasIdentityLookups() bool`

HasIdentityLookups returns a boolean if a field has been set.

### GetIdentityMappingPolicy

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetIdentityMappingPolicy() IdentityMappingPolicyDTO`

GetIdentityMappingPolicy returns the IdentityMappingPolicy field if non-nil, zero value otherwise.

### GetIdentityMappingPolicyOk

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetIdentityMappingPolicyOk() (*IdentityMappingPolicyDTO, bool)`

GetIdentityMappingPolicyOk returns a tuple with the IdentityMappingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityMappingPolicy

`func (o *ExternalOpenIDConnectRelayingPartyDTO) SetIdentityMappingPolicy(v IdentityMappingPolicyDTO)`

SetIdentityMappingPolicy sets IdentityMappingPolicy field to given value.

### HasIdentityMappingPolicy

`func (o *ExternalOpenIDConnectRelayingPartyDTO) HasIdentityMappingPolicy() bool`

HasIdentityMappingPolicy returns a boolean if a field has been set.

### GetIsRemote

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetIsRemote() bool`

GetIsRemote returns the IsRemote field if non-nil, zero value otherwise.

### GetIsRemoteOk

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetIsRemoteOk() (*bool, bool)`

GetIsRemoteOk returns a tuple with the IsRemote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemote

`func (o *ExternalOpenIDConnectRelayingPartyDTO) SetIsRemote(v bool)`

SetIsRemote sets IsRemote field to given value.

### HasIsRemote

`func (o *ExternalOpenIDConnectRelayingPartyDTO) HasIsRemote() bool`

HasIsRemote returns a boolean if a field has been set.

### GetLocation

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetLocation() LocationDTO`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetLocationOk() (*LocationDTO, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ExternalOpenIDConnectRelayingPartyDTO) SetLocation(v LocationDTO)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ExternalOpenIDConnectRelayingPartyDTO) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetMetadata

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetMetadata() ResourceDTO`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetMetadataOk() (*ResourceDTO, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ExternalOpenIDConnectRelayingPartyDTO) SetMetadata(v ResourceDTO)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ExternalOpenIDConnectRelayingPartyDTO) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExternalOpenIDConnectRelayingPartyDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExternalOpenIDConnectRelayingPartyDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPostLogoutRedirectionURIs

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetPostLogoutRedirectionURIs() []string`

GetPostLogoutRedirectionURIs returns the PostLogoutRedirectionURIs field if non-nil, zero value otherwise.

### GetPostLogoutRedirectionURIsOk

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetPostLogoutRedirectionURIsOk() (*[]string, bool)`

GetPostLogoutRedirectionURIsOk returns a tuple with the PostLogoutRedirectionURIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostLogoutRedirectionURIs

`func (o *ExternalOpenIDConnectRelayingPartyDTO) SetPostLogoutRedirectionURIs(v []string)`

SetPostLogoutRedirectionURIs sets PostLogoutRedirectionURIs field to given value.

### HasPostLogoutRedirectionURIs

`func (o *ExternalOpenIDConnectRelayingPartyDTO) HasPostLogoutRedirectionURIs() bool`

HasPostLogoutRedirectionURIs returns a boolean if a field has been set.

### GetRemote

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetRemote() bool`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetRemoteOk() (*bool, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *ExternalOpenIDConnectRelayingPartyDTO) SetRemote(v bool)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *ExternalOpenIDConnectRelayingPartyDTO) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### GetResponseTypes

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetResponseTypes() []string`

GetResponseTypes returns the ResponseTypes field if non-nil, zero value otherwise.

### GetResponseTypesOk

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetResponseTypesOk() (*[]string, bool)`

GetResponseTypesOk returns a tuple with the ResponseTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTypes

`func (o *ExternalOpenIDConnectRelayingPartyDTO) SetResponseTypes(v []string)`

SetResponseTypes sets ResponseTypes field to given value.

### HasResponseTypes

`func (o *ExternalOpenIDConnectRelayingPartyDTO) HasResponseTypes() bool`

HasResponseTypes returns a boolean if a field has been set.

### GetRole

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ExternalOpenIDConnectRelayingPartyDTO) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *ExternalOpenIDConnectRelayingPartyDTO) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSigningAlg

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetSigningAlg() string`

GetSigningAlg returns the SigningAlg field if non-nil, zero value otherwise.

### GetSigningAlgOk

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetSigningAlgOk() (*string, bool)`

GetSigningAlgOk returns a tuple with the SigningAlg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningAlg

`func (o *ExternalOpenIDConnectRelayingPartyDTO) SetSigningAlg(v string)`

SetSigningAlg sets SigningAlg field to given value.

### HasSigningAlg

`func (o *ExternalOpenIDConnectRelayingPartyDTO) HasSigningAlg() bool`

HasSigningAlg returns a boolean if a field has been set.

### GetX

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetX() float64`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetXOk() (*float64, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *ExternalOpenIDConnectRelayingPartyDTO) SetX(v float64)`

SetX sets X field to given value.

### HasX

`func (o *ExternalOpenIDConnectRelayingPartyDTO) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetY() float64`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *ExternalOpenIDConnectRelayingPartyDTO) GetYOk() (*float64, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *ExternalOpenIDConnectRelayingPartyDTO) SetY(v float64)`

SetY sets Y field to given value.

### HasY

`func (o *ExternalOpenIDConnectRelayingPartyDTO) HasY() bool`

HasY returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


