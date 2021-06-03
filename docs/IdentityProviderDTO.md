# IdentityProviderDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveBindings** | Pointer to **[]string** |  | [optional] 
**ActiveProfiles** | Pointer to **[]string** |  | [optional] 
**AttributeProfile** | Pointer to [**AttributeProfileDTO**](AttributeProfileDTO.md) |  | [optional] 
**AuthenticationContract** | Pointer to [**AuthenticationContractDTO**](AuthenticationContractDTO.md) |  | [optional] 
**AuthenticationMechanisms** | Pointer to [**[]AuthenticationMechanismDTO**](AuthenticationMechanismDTO.md) |  | [optional] 
**Config** | Pointer to [**ProviderConfigDTO**](ProviderConfigDTO.md) |  | [optional] 
**DashboardUrl** | Pointer to **string** |  | [optional] 
**DelegatedAuthentications** | Pointer to [**[]DelegatedAuthenticationDTO**](DelegatedAuthenticationDTO.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DestroyPreviousSession** | Pointer to **bool** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**EmissionPolicy** | Pointer to [**AuthenticationAssertionEmissionPolicyDTO**](AuthenticationAssertionEmissionPolicyDTO.md) |  | [optional] 
**EnableMetadataEndpoint** | Pointer to **bool** |  | [optional] 
**EncryptAssertion** | Pointer to **bool** |  | [optional] 
**EncryptAssertionAlgorithm** | Pointer to **string** |  | [optional] 
**ErrorBinding** | Pointer to **string** |  | [optional] 
**ExternallyHostedIdentityConfirmationTokenService** | Pointer to **bool** |  | [optional] 
**FederatedConnectionsA** | Pointer to [**[]FederatedConnectionDTO**](FederatedConnectionDTO.md) |  | [optional] 
**FederatedConnectionsB** | Pointer to [**[]FederatedConnectionDTO**](FederatedConnectionDTO.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**IdentityLookups** | Pointer to [**[]IdentityLookupDTO**](IdentityLookupDTO.md) |  | [optional] 
**IdentityAppliance** | Pointer to [**IdentityApplianceDefinitionDTO**](IdentityApplianceDefinitionDTO.md) |  | [optional] 
**IdentityConfirmationEnabled** | Pointer to **bool** |  | [optional] 
**IdentityConfirmationOAuth2AuthorizationServerEndpoint** | Pointer to **string** |  | [optional] 
**IdentityConfirmationOAuth2ClientId** | Pointer to **string** |  | [optional] 
**IdentityConfirmationOAuth2ClientSecret** | Pointer to **string** |  | [optional] 
**IdentityConfirmationPolicy** | Pointer to [**ExtensionDTO**](ExtensionDTO.md) |  | [optional] 
**IgnoreRequestedNameIDPolicy** | Pointer to **bool** |  | [optional] 
**IsRemote** | Pointer to **bool** |  | [optional] 
**Location** | Pointer to [**LocationDTO**](LocationDTO.md) |  | [optional] 
**MaxSessionsPerUser** | Pointer to **int32** |  | [optional] 
**MessageTtl** | Pointer to **int32** |  | [optional] 
**MessageTtlTolerance** | Pointer to **int32** |  | [optional] 
**Metadata** | Pointer to [**ResourceDTO**](ResourceDTO.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Oauth2Clients** | Pointer to [**[]OAuth2ClientDTO**](OAuth2ClientDTO.md) |  | [optional] 
**Oauth2ClientsConfig** | Pointer to **string** |  | [optional] 
**Oauth2Enabled** | Pointer to **bool** |  | [optional] 
**Oauth2Key** | Pointer to **string** |  | [optional] 
**Oauth2RememberMeTokenValidity** | Pointer to **int64** |  | [optional] 
**Oauth2TokenValidity** | Pointer to **int64** |  | [optional] 
**OidcAccessTokenTimeToLive** | Pointer to **int32** |  | [optional] 
**OidcAuthzCodeTimeToLive** | Pointer to **int32** |  | [optional] 
**OidcIdTokenTimeToLive** | Pointer to **int32** |  | [optional] 
**OpenIdEnabled** | Pointer to **bool** |  | [optional] 
**PwdlessAuthnEnabled** | Pointer to **bool** |  | [optional] 
**PwdlessAuthnFrom** | Pointer to **string** |  | [optional] 
**PwdlessAuthnSubject** | Pointer to **string** |  | [optional] 
**PwdlessAuthnTemplate** | Pointer to **string** |  | [optional] 
**PwdlessAuthnTo** | Pointer to **string** |  | [optional] 
**Remote** | Pointer to **bool** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**SessionManagerFactory** | Pointer to [**SessionManagerFactoryDTO**](SessionManagerFactoryDTO.md) |  | [optional] 
**SignRequests** | Pointer to **bool** |  | [optional] 
**SignatureHash** | Pointer to **string** |  | [optional] 
**SsoSessionTimeout** | Pointer to **int32** |  | [optional] 
**SubjectAuthnPolicies** | Pointer to [**[]SubjectAuthenticationPolicyDTO**](SubjectAuthenticationPolicyDTO.md) |  | [optional] 
**SubjectNameIDPolicy** | Pointer to [**SubjectNameIdentifierPolicyDTO**](SubjectNameIdentifierPolicyDTO.md) |  | [optional] 
**UserDashboardBranding** | Pointer to **string** |  | [optional] 
**WantAuthnRequestsSigned** | Pointer to **bool** |  | [optional] 
**WantSignedRequests** | Pointer to **bool** |  | [optional] 
**X** | Pointer to **float64** |  | [optional] 
**Y** | Pointer to **float64** |  | [optional] 

## Methods

### NewIdentityProviderDTO

`func NewIdentityProviderDTO() *IdentityProviderDTO`

NewIdentityProviderDTO instantiates a new IdentityProviderDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityProviderDTOWithDefaults

`func NewIdentityProviderDTOWithDefaults() *IdentityProviderDTO`

NewIdentityProviderDTOWithDefaults instantiates a new IdentityProviderDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveBindings

`func (o *IdentityProviderDTO) GetActiveBindings() []string`

GetActiveBindings returns the ActiveBindings field if non-nil, zero value otherwise.

### GetActiveBindingsOk

`func (o *IdentityProviderDTO) GetActiveBindingsOk() (*[]string, bool)`

GetActiveBindingsOk returns a tuple with the ActiveBindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveBindings

`func (o *IdentityProviderDTO) SetActiveBindings(v []string)`

SetActiveBindings sets ActiveBindings field to given value.

### HasActiveBindings

`func (o *IdentityProviderDTO) HasActiveBindings() bool`

HasActiveBindings returns a boolean if a field has been set.

### GetActiveProfiles

`func (o *IdentityProviderDTO) GetActiveProfiles() []string`

GetActiveProfiles returns the ActiveProfiles field if non-nil, zero value otherwise.

### GetActiveProfilesOk

`func (o *IdentityProviderDTO) GetActiveProfilesOk() (*[]string, bool)`

GetActiveProfilesOk returns a tuple with the ActiveProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveProfiles

`func (o *IdentityProviderDTO) SetActiveProfiles(v []string)`

SetActiveProfiles sets ActiveProfiles field to given value.

### HasActiveProfiles

`func (o *IdentityProviderDTO) HasActiveProfiles() bool`

HasActiveProfiles returns a boolean if a field has been set.

### GetAttributeProfile

`func (o *IdentityProviderDTO) GetAttributeProfile() AttributeProfileDTO`

GetAttributeProfile returns the AttributeProfile field if non-nil, zero value otherwise.

### GetAttributeProfileOk

`func (o *IdentityProviderDTO) GetAttributeProfileOk() (*AttributeProfileDTO, bool)`

GetAttributeProfileOk returns a tuple with the AttributeProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeProfile

`func (o *IdentityProviderDTO) SetAttributeProfile(v AttributeProfileDTO)`

SetAttributeProfile sets AttributeProfile field to given value.

### HasAttributeProfile

`func (o *IdentityProviderDTO) HasAttributeProfile() bool`

HasAttributeProfile returns a boolean if a field has been set.

### GetAuthenticationContract

`func (o *IdentityProviderDTO) GetAuthenticationContract() AuthenticationContractDTO`

GetAuthenticationContract returns the AuthenticationContract field if non-nil, zero value otherwise.

### GetAuthenticationContractOk

`func (o *IdentityProviderDTO) GetAuthenticationContractOk() (*AuthenticationContractDTO, bool)`

GetAuthenticationContractOk returns a tuple with the AuthenticationContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationContract

`func (o *IdentityProviderDTO) SetAuthenticationContract(v AuthenticationContractDTO)`

SetAuthenticationContract sets AuthenticationContract field to given value.

### HasAuthenticationContract

`func (o *IdentityProviderDTO) HasAuthenticationContract() bool`

HasAuthenticationContract returns a boolean if a field has been set.

### GetAuthenticationMechanisms

`func (o *IdentityProviderDTO) GetAuthenticationMechanisms() []AuthenticationMechanismDTO`

GetAuthenticationMechanisms returns the AuthenticationMechanisms field if non-nil, zero value otherwise.

### GetAuthenticationMechanismsOk

`func (o *IdentityProviderDTO) GetAuthenticationMechanismsOk() (*[]AuthenticationMechanismDTO, bool)`

GetAuthenticationMechanismsOk returns a tuple with the AuthenticationMechanisms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMechanisms

`func (o *IdentityProviderDTO) SetAuthenticationMechanisms(v []AuthenticationMechanismDTO)`

SetAuthenticationMechanisms sets AuthenticationMechanisms field to given value.

### HasAuthenticationMechanisms

`func (o *IdentityProviderDTO) HasAuthenticationMechanisms() bool`

HasAuthenticationMechanisms returns a boolean if a field has been set.

### GetConfig

`func (o *IdentityProviderDTO) GetConfig() ProviderConfigDTO`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *IdentityProviderDTO) GetConfigOk() (*ProviderConfigDTO, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *IdentityProviderDTO) SetConfig(v ProviderConfigDTO)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *IdentityProviderDTO) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetDashboardUrl

`func (o *IdentityProviderDTO) GetDashboardUrl() string`

GetDashboardUrl returns the DashboardUrl field if non-nil, zero value otherwise.

### GetDashboardUrlOk

`func (o *IdentityProviderDTO) GetDashboardUrlOk() (*string, bool)`

GetDashboardUrlOk returns a tuple with the DashboardUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardUrl

`func (o *IdentityProviderDTO) SetDashboardUrl(v string)`

SetDashboardUrl sets DashboardUrl field to given value.

### HasDashboardUrl

`func (o *IdentityProviderDTO) HasDashboardUrl() bool`

HasDashboardUrl returns a boolean if a field has been set.

### GetDelegatedAuthentications

`func (o *IdentityProviderDTO) GetDelegatedAuthentications() []DelegatedAuthenticationDTO`

GetDelegatedAuthentications returns the DelegatedAuthentications field if non-nil, zero value otherwise.

### GetDelegatedAuthenticationsOk

`func (o *IdentityProviderDTO) GetDelegatedAuthenticationsOk() (*[]DelegatedAuthenticationDTO, bool)`

GetDelegatedAuthenticationsOk returns a tuple with the DelegatedAuthentications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatedAuthentications

`func (o *IdentityProviderDTO) SetDelegatedAuthentications(v []DelegatedAuthenticationDTO)`

SetDelegatedAuthentications sets DelegatedAuthentications field to given value.

### HasDelegatedAuthentications

`func (o *IdentityProviderDTO) HasDelegatedAuthentications() bool`

HasDelegatedAuthentications returns a boolean if a field has been set.

### GetDescription

`func (o *IdentityProviderDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IdentityProviderDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IdentityProviderDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IdentityProviderDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDestroyPreviousSession

`func (o *IdentityProviderDTO) GetDestroyPreviousSession() bool`

GetDestroyPreviousSession returns the DestroyPreviousSession field if non-nil, zero value otherwise.

### GetDestroyPreviousSessionOk

`func (o *IdentityProviderDTO) GetDestroyPreviousSessionOk() (*bool, bool)`

GetDestroyPreviousSessionOk returns a tuple with the DestroyPreviousSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestroyPreviousSession

`func (o *IdentityProviderDTO) SetDestroyPreviousSession(v bool)`

SetDestroyPreviousSession sets DestroyPreviousSession field to given value.

### HasDestroyPreviousSession

`func (o *IdentityProviderDTO) HasDestroyPreviousSession() bool`

HasDestroyPreviousSession returns a boolean if a field has been set.

### GetDisplayName

`func (o *IdentityProviderDTO) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IdentityProviderDTO) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IdentityProviderDTO) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IdentityProviderDTO) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetElementId

`func (o *IdentityProviderDTO) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *IdentityProviderDTO) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *IdentityProviderDTO) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *IdentityProviderDTO) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetEmissionPolicy

`func (o *IdentityProviderDTO) GetEmissionPolicy() AuthenticationAssertionEmissionPolicyDTO`

GetEmissionPolicy returns the EmissionPolicy field if non-nil, zero value otherwise.

### GetEmissionPolicyOk

`func (o *IdentityProviderDTO) GetEmissionPolicyOk() (*AuthenticationAssertionEmissionPolicyDTO, bool)`

GetEmissionPolicyOk returns a tuple with the EmissionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmissionPolicy

`func (o *IdentityProviderDTO) SetEmissionPolicy(v AuthenticationAssertionEmissionPolicyDTO)`

SetEmissionPolicy sets EmissionPolicy field to given value.

### HasEmissionPolicy

`func (o *IdentityProviderDTO) HasEmissionPolicy() bool`

HasEmissionPolicy returns a boolean if a field has been set.

### GetEnableMetadataEndpoint

`func (o *IdentityProviderDTO) GetEnableMetadataEndpoint() bool`

GetEnableMetadataEndpoint returns the EnableMetadataEndpoint field if non-nil, zero value otherwise.

### GetEnableMetadataEndpointOk

`func (o *IdentityProviderDTO) GetEnableMetadataEndpointOk() (*bool, bool)`

GetEnableMetadataEndpointOk returns a tuple with the EnableMetadataEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMetadataEndpoint

`func (o *IdentityProviderDTO) SetEnableMetadataEndpoint(v bool)`

SetEnableMetadataEndpoint sets EnableMetadataEndpoint field to given value.

### HasEnableMetadataEndpoint

`func (o *IdentityProviderDTO) HasEnableMetadataEndpoint() bool`

HasEnableMetadataEndpoint returns a boolean if a field has been set.

### GetEncryptAssertion

`func (o *IdentityProviderDTO) GetEncryptAssertion() bool`

GetEncryptAssertion returns the EncryptAssertion field if non-nil, zero value otherwise.

### GetEncryptAssertionOk

`func (o *IdentityProviderDTO) GetEncryptAssertionOk() (*bool, bool)`

GetEncryptAssertionOk returns a tuple with the EncryptAssertion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptAssertion

`func (o *IdentityProviderDTO) SetEncryptAssertion(v bool)`

SetEncryptAssertion sets EncryptAssertion field to given value.

### HasEncryptAssertion

`func (o *IdentityProviderDTO) HasEncryptAssertion() bool`

HasEncryptAssertion returns a boolean if a field has been set.

### GetEncryptAssertionAlgorithm

`func (o *IdentityProviderDTO) GetEncryptAssertionAlgorithm() string`

GetEncryptAssertionAlgorithm returns the EncryptAssertionAlgorithm field if non-nil, zero value otherwise.

### GetEncryptAssertionAlgorithmOk

`func (o *IdentityProviderDTO) GetEncryptAssertionAlgorithmOk() (*string, bool)`

GetEncryptAssertionAlgorithmOk returns a tuple with the EncryptAssertionAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptAssertionAlgorithm

`func (o *IdentityProviderDTO) SetEncryptAssertionAlgorithm(v string)`

SetEncryptAssertionAlgorithm sets EncryptAssertionAlgorithm field to given value.

### HasEncryptAssertionAlgorithm

`func (o *IdentityProviderDTO) HasEncryptAssertionAlgorithm() bool`

HasEncryptAssertionAlgorithm returns a boolean if a field has been set.

### GetErrorBinding

`func (o *IdentityProviderDTO) GetErrorBinding() string`

GetErrorBinding returns the ErrorBinding field if non-nil, zero value otherwise.

### GetErrorBindingOk

`func (o *IdentityProviderDTO) GetErrorBindingOk() (*string, bool)`

GetErrorBindingOk returns a tuple with the ErrorBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorBinding

`func (o *IdentityProviderDTO) SetErrorBinding(v string)`

SetErrorBinding sets ErrorBinding field to given value.

### HasErrorBinding

`func (o *IdentityProviderDTO) HasErrorBinding() bool`

HasErrorBinding returns a boolean if a field has been set.

### GetExternallyHostedIdentityConfirmationTokenService

`func (o *IdentityProviderDTO) GetExternallyHostedIdentityConfirmationTokenService() bool`

GetExternallyHostedIdentityConfirmationTokenService returns the ExternallyHostedIdentityConfirmationTokenService field if non-nil, zero value otherwise.

### GetExternallyHostedIdentityConfirmationTokenServiceOk

`func (o *IdentityProviderDTO) GetExternallyHostedIdentityConfirmationTokenServiceOk() (*bool, bool)`

GetExternallyHostedIdentityConfirmationTokenServiceOk returns a tuple with the ExternallyHostedIdentityConfirmationTokenService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternallyHostedIdentityConfirmationTokenService

`func (o *IdentityProviderDTO) SetExternallyHostedIdentityConfirmationTokenService(v bool)`

SetExternallyHostedIdentityConfirmationTokenService sets ExternallyHostedIdentityConfirmationTokenService field to given value.

### HasExternallyHostedIdentityConfirmationTokenService

`func (o *IdentityProviderDTO) HasExternallyHostedIdentityConfirmationTokenService() bool`

HasExternallyHostedIdentityConfirmationTokenService returns a boolean if a field has been set.

### GetFederatedConnectionsA

`func (o *IdentityProviderDTO) GetFederatedConnectionsA() []FederatedConnectionDTO`

GetFederatedConnectionsA returns the FederatedConnectionsA field if non-nil, zero value otherwise.

### GetFederatedConnectionsAOk

`func (o *IdentityProviderDTO) GetFederatedConnectionsAOk() (*[]FederatedConnectionDTO, bool)`

GetFederatedConnectionsAOk returns a tuple with the FederatedConnectionsA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederatedConnectionsA

`func (o *IdentityProviderDTO) SetFederatedConnectionsA(v []FederatedConnectionDTO)`

SetFederatedConnectionsA sets FederatedConnectionsA field to given value.

### HasFederatedConnectionsA

`func (o *IdentityProviderDTO) HasFederatedConnectionsA() bool`

HasFederatedConnectionsA returns a boolean if a field has been set.

### GetFederatedConnectionsB

`func (o *IdentityProviderDTO) GetFederatedConnectionsB() []FederatedConnectionDTO`

GetFederatedConnectionsB returns the FederatedConnectionsB field if non-nil, zero value otherwise.

### GetFederatedConnectionsBOk

`func (o *IdentityProviderDTO) GetFederatedConnectionsBOk() (*[]FederatedConnectionDTO, bool)`

GetFederatedConnectionsBOk returns a tuple with the FederatedConnectionsB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederatedConnectionsB

`func (o *IdentityProviderDTO) SetFederatedConnectionsB(v []FederatedConnectionDTO)`

SetFederatedConnectionsB sets FederatedConnectionsB field to given value.

### HasFederatedConnectionsB

`func (o *IdentityProviderDTO) HasFederatedConnectionsB() bool`

HasFederatedConnectionsB returns a boolean if a field has been set.

### GetId

`func (o *IdentityProviderDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentityProviderDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentityProviderDTO) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *IdentityProviderDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentityLookups

`func (o *IdentityProviderDTO) GetIdentityLookups() []IdentityLookupDTO`

GetIdentityLookups returns the IdentityLookups field if non-nil, zero value otherwise.

### GetIdentityLookupsOk

`func (o *IdentityProviderDTO) GetIdentityLookupsOk() (*[]IdentityLookupDTO, bool)`

GetIdentityLookupsOk returns a tuple with the IdentityLookups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityLookups

`func (o *IdentityProviderDTO) SetIdentityLookups(v []IdentityLookupDTO)`

SetIdentityLookups sets IdentityLookups field to given value.

### HasIdentityLookups

`func (o *IdentityProviderDTO) HasIdentityLookups() bool`

HasIdentityLookups returns a boolean if a field has been set.

### GetIdentityAppliance

`func (o *IdentityProviderDTO) GetIdentityAppliance() IdentityApplianceDefinitionDTO`

GetIdentityAppliance returns the IdentityAppliance field if non-nil, zero value otherwise.

### GetIdentityApplianceOk

`func (o *IdentityProviderDTO) GetIdentityApplianceOk() (*IdentityApplianceDefinitionDTO, bool)`

GetIdentityApplianceOk returns a tuple with the IdentityAppliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityAppliance

`func (o *IdentityProviderDTO) SetIdentityAppliance(v IdentityApplianceDefinitionDTO)`

SetIdentityAppliance sets IdentityAppliance field to given value.

### HasIdentityAppliance

`func (o *IdentityProviderDTO) HasIdentityAppliance() bool`

HasIdentityAppliance returns a boolean if a field has been set.

### GetIdentityConfirmationEnabled

`func (o *IdentityProviderDTO) GetIdentityConfirmationEnabled() bool`

GetIdentityConfirmationEnabled returns the IdentityConfirmationEnabled field if non-nil, zero value otherwise.

### GetIdentityConfirmationEnabledOk

`func (o *IdentityProviderDTO) GetIdentityConfirmationEnabledOk() (*bool, bool)`

GetIdentityConfirmationEnabledOk returns a tuple with the IdentityConfirmationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityConfirmationEnabled

`func (o *IdentityProviderDTO) SetIdentityConfirmationEnabled(v bool)`

SetIdentityConfirmationEnabled sets IdentityConfirmationEnabled field to given value.

### HasIdentityConfirmationEnabled

`func (o *IdentityProviderDTO) HasIdentityConfirmationEnabled() bool`

HasIdentityConfirmationEnabled returns a boolean if a field has been set.

### GetIdentityConfirmationOAuth2AuthorizationServerEndpoint

`func (o *IdentityProviderDTO) GetIdentityConfirmationOAuth2AuthorizationServerEndpoint() string`

GetIdentityConfirmationOAuth2AuthorizationServerEndpoint returns the IdentityConfirmationOAuth2AuthorizationServerEndpoint field if non-nil, zero value otherwise.

### GetIdentityConfirmationOAuth2AuthorizationServerEndpointOk

`func (o *IdentityProviderDTO) GetIdentityConfirmationOAuth2AuthorizationServerEndpointOk() (*string, bool)`

GetIdentityConfirmationOAuth2AuthorizationServerEndpointOk returns a tuple with the IdentityConfirmationOAuth2AuthorizationServerEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityConfirmationOAuth2AuthorizationServerEndpoint

`func (o *IdentityProviderDTO) SetIdentityConfirmationOAuth2AuthorizationServerEndpoint(v string)`

SetIdentityConfirmationOAuth2AuthorizationServerEndpoint sets IdentityConfirmationOAuth2AuthorizationServerEndpoint field to given value.

### HasIdentityConfirmationOAuth2AuthorizationServerEndpoint

`func (o *IdentityProviderDTO) HasIdentityConfirmationOAuth2AuthorizationServerEndpoint() bool`

HasIdentityConfirmationOAuth2AuthorizationServerEndpoint returns a boolean if a field has been set.

### GetIdentityConfirmationOAuth2ClientId

`func (o *IdentityProviderDTO) GetIdentityConfirmationOAuth2ClientId() string`

GetIdentityConfirmationOAuth2ClientId returns the IdentityConfirmationOAuth2ClientId field if non-nil, zero value otherwise.

### GetIdentityConfirmationOAuth2ClientIdOk

`func (o *IdentityProviderDTO) GetIdentityConfirmationOAuth2ClientIdOk() (*string, bool)`

GetIdentityConfirmationOAuth2ClientIdOk returns a tuple with the IdentityConfirmationOAuth2ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityConfirmationOAuth2ClientId

`func (o *IdentityProviderDTO) SetIdentityConfirmationOAuth2ClientId(v string)`

SetIdentityConfirmationOAuth2ClientId sets IdentityConfirmationOAuth2ClientId field to given value.

### HasIdentityConfirmationOAuth2ClientId

`func (o *IdentityProviderDTO) HasIdentityConfirmationOAuth2ClientId() bool`

HasIdentityConfirmationOAuth2ClientId returns a boolean if a field has been set.

### GetIdentityConfirmationOAuth2ClientSecret

`func (o *IdentityProviderDTO) GetIdentityConfirmationOAuth2ClientSecret() string`

GetIdentityConfirmationOAuth2ClientSecret returns the IdentityConfirmationOAuth2ClientSecret field if non-nil, zero value otherwise.

### GetIdentityConfirmationOAuth2ClientSecretOk

`func (o *IdentityProviderDTO) GetIdentityConfirmationOAuth2ClientSecretOk() (*string, bool)`

GetIdentityConfirmationOAuth2ClientSecretOk returns a tuple with the IdentityConfirmationOAuth2ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityConfirmationOAuth2ClientSecret

`func (o *IdentityProviderDTO) SetIdentityConfirmationOAuth2ClientSecret(v string)`

SetIdentityConfirmationOAuth2ClientSecret sets IdentityConfirmationOAuth2ClientSecret field to given value.

### HasIdentityConfirmationOAuth2ClientSecret

`func (o *IdentityProviderDTO) HasIdentityConfirmationOAuth2ClientSecret() bool`

HasIdentityConfirmationOAuth2ClientSecret returns a boolean if a field has been set.

### GetIdentityConfirmationPolicy

`func (o *IdentityProviderDTO) GetIdentityConfirmationPolicy() ExtensionDTO`

GetIdentityConfirmationPolicy returns the IdentityConfirmationPolicy field if non-nil, zero value otherwise.

### GetIdentityConfirmationPolicyOk

`func (o *IdentityProviderDTO) GetIdentityConfirmationPolicyOk() (*ExtensionDTO, bool)`

GetIdentityConfirmationPolicyOk returns a tuple with the IdentityConfirmationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityConfirmationPolicy

`func (o *IdentityProviderDTO) SetIdentityConfirmationPolicy(v ExtensionDTO)`

SetIdentityConfirmationPolicy sets IdentityConfirmationPolicy field to given value.

### HasIdentityConfirmationPolicy

`func (o *IdentityProviderDTO) HasIdentityConfirmationPolicy() bool`

HasIdentityConfirmationPolicy returns a boolean if a field has been set.

### GetIgnoreRequestedNameIDPolicy

`func (o *IdentityProviderDTO) GetIgnoreRequestedNameIDPolicy() bool`

GetIgnoreRequestedNameIDPolicy returns the IgnoreRequestedNameIDPolicy field if non-nil, zero value otherwise.

### GetIgnoreRequestedNameIDPolicyOk

`func (o *IdentityProviderDTO) GetIgnoreRequestedNameIDPolicyOk() (*bool, bool)`

GetIgnoreRequestedNameIDPolicyOk returns a tuple with the IgnoreRequestedNameIDPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreRequestedNameIDPolicy

`func (o *IdentityProviderDTO) SetIgnoreRequestedNameIDPolicy(v bool)`

SetIgnoreRequestedNameIDPolicy sets IgnoreRequestedNameIDPolicy field to given value.

### HasIgnoreRequestedNameIDPolicy

`func (o *IdentityProviderDTO) HasIgnoreRequestedNameIDPolicy() bool`

HasIgnoreRequestedNameIDPolicy returns a boolean if a field has been set.

### GetIsRemote

`func (o *IdentityProviderDTO) GetIsRemote() bool`

GetIsRemote returns the IsRemote field if non-nil, zero value otherwise.

### GetIsRemoteOk

`func (o *IdentityProviderDTO) GetIsRemoteOk() (*bool, bool)`

GetIsRemoteOk returns a tuple with the IsRemote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemote

`func (o *IdentityProviderDTO) SetIsRemote(v bool)`

SetIsRemote sets IsRemote field to given value.

### HasIsRemote

`func (o *IdentityProviderDTO) HasIsRemote() bool`

HasIsRemote returns a boolean if a field has been set.

### GetLocation

`func (o *IdentityProviderDTO) GetLocation() LocationDTO`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *IdentityProviderDTO) GetLocationOk() (*LocationDTO, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *IdentityProviderDTO) SetLocation(v LocationDTO)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *IdentityProviderDTO) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetMaxSessionsPerUser

`func (o *IdentityProviderDTO) GetMaxSessionsPerUser() int32`

GetMaxSessionsPerUser returns the MaxSessionsPerUser field if non-nil, zero value otherwise.

### GetMaxSessionsPerUserOk

`func (o *IdentityProviderDTO) GetMaxSessionsPerUserOk() (*int32, bool)`

GetMaxSessionsPerUserOk returns a tuple with the MaxSessionsPerUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSessionsPerUser

`func (o *IdentityProviderDTO) SetMaxSessionsPerUser(v int32)`

SetMaxSessionsPerUser sets MaxSessionsPerUser field to given value.

### HasMaxSessionsPerUser

`func (o *IdentityProviderDTO) HasMaxSessionsPerUser() bool`

HasMaxSessionsPerUser returns a boolean if a field has been set.

### GetMessageTtl

`func (o *IdentityProviderDTO) GetMessageTtl() int32`

GetMessageTtl returns the MessageTtl field if non-nil, zero value otherwise.

### GetMessageTtlOk

`func (o *IdentityProviderDTO) GetMessageTtlOk() (*int32, bool)`

GetMessageTtlOk returns a tuple with the MessageTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageTtl

`func (o *IdentityProviderDTO) SetMessageTtl(v int32)`

SetMessageTtl sets MessageTtl field to given value.

### HasMessageTtl

`func (o *IdentityProviderDTO) HasMessageTtl() bool`

HasMessageTtl returns a boolean if a field has been set.

### GetMessageTtlTolerance

`func (o *IdentityProviderDTO) GetMessageTtlTolerance() int32`

GetMessageTtlTolerance returns the MessageTtlTolerance field if non-nil, zero value otherwise.

### GetMessageTtlToleranceOk

`func (o *IdentityProviderDTO) GetMessageTtlToleranceOk() (*int32, bool)`

GetMessageTtlToleranceOk returns a tuple with the MessageTtlTolerance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageTtlTolerance

`func (o *IdentityProviderDTO) SetMessageTtlTolerance(v int32)`

SetMessageTtlTolerance sets MessageTtlTolerance field to given value.

### HasMessageTtlTolerance

`func (o *IdentityProviderDTO) HasMessageTtlTolerance() bool`

HasMessageTtlTolerance returns a boolean if a field has been set.

### GetMetadata

`func (o *IdentityProviderDTO) GetMetadata() ResourceDTO`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IdentityProviderDTO) GetMetadataOk() (*ResourceDTO, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IdentityProviderDTO) SetMetadata(v ResourceDTO)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IdentityProviderDTO) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *IdentityProviderDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdentityProviderDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdentityProviderDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IdentityProviderDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOauth2Clients

`func (o *IdentityProviderDTO) GetOauth2Clients() []OAuth2ClientDTO`

GetOauth2Clients returns the Oauth2Clients field if non-nil, zero value otherwise.

### GetOauth2ClientsOk

`func (o *IdentityProviderDTO) GetOauth2ClientsOk() (*[]OAuth2ClientDTO, bool)`

GetOauth2ClientsOk returns a tuple with the Oauth2Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2Clients

`func (o *IdentityProviderDTO) SetOauth2Clients(v []OAuth2ClientDTO)`

SetOauth2Clients sets Oauth2Clients field to given value.

### HasOauth2Clients

`func (o *IdentityProviderDTO) HasOauth2Clients() bool`

HasOauth2Clients returns a boolean if a field has been set.

### GetOauth2ClientsConfig

`func (o *IdentityProviderDTO) GetOauth2ClientsConfig() string`

GetOauth2ClientsConfig returns the Oauth2ClientsConfig field if non-nil, zero value otherwise.

### GetOauth2ClientsConfigOk

`func (o *IdentityProviderDTO) GetOauth2ClientsConfigOk() (*string, bool)`

GetOauth2ClientsConfigOk returns a tuple with the Oauth2ClientsConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2ClientsConfig

`func (o *IdentityProviderDTO) SetOauth2ClientsConfig(v string)`

SetOauth2ClientsConfig sets Oauth2ClientsConfig field to given value.

### HasOauth2ClientsConfig

`func (o *IdentityProviderDTO) HasOauth2ClientsConfig() bool`

HasOauth2ClientsConfig returns a boolean if a field has been set.

### GetOauth2Enabled

`func (o *IdentityProviderDTO) GetOauth2Enabled() bool`

GetOauth2Enabled returns the Oauth2Enabled field if non-nil, zero value otherwise.

### GetOauth2EnabledOk

`func (o *IdentityProviderDTO) GetOauth2EnabledOk() (*bool, bool)`

GetOauth2EnabledOk returns a tuple with the Oauth2Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2Enabled

`func (o *IdentityProviderDTO) SetOauth2Enabled(v bool)`

SetOauth2Enabled sets Oauth2Enabled field to given value.

### HasOauth2Enabled

`func (o *IdentityProviderDTO) HasOauth2Enabled() bool`

HasOauth2Enabled returns a boolean if a field has been set.

### GetOauth2Key

`func (o *IdentityProviderDTO) GetOauth2Key() string`

GetOauth2Key returns the Oauth2Key field if non-nil, zero value otherwise.

### GetOauth2KeyOk

`func (o *IdentityProviderDTO) GetOauth2KeyOk() (*string, bool)`

GetOauth2KeyOk returns a tuple with the Oauth2Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2Key

`func (o *IdentityProviderDTO) SetOauth2Key(v string)`

SetOauth2Key sets Oauth2Key field to given value.

### HasOauth2Key

`func (o *IdentityProviderDTO) HasOauth2Key() bool`

HasOauth2Key returns a boolean if a field has been set.

### GetOauth2RememberMeTokenValidity

`func (o *IdentityProviderDTO) GetOauth2RememberMeTokenValidity() int64`

GetOauth2RememberMeTokenValidity returns the Oauth2RememberMeTokenValidity field if non-nil, zero value otherwise.

### GetOauth2RememberMeTokenValidityOk

`func (o *IdentityProviderDTO) GetOauth2RememberMeTokenValidityOk() (*int64, bool)`

GetOauth2RememberMeTokenValidityOk returns a tuple with the Oauth2RememberMeTokenValidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2RememberMeTokenValidity

`func (o *IdentityProviderDTO) SetOauth2RememberMeTokenValidity(v int64)`

SetOauth2RememberMeTokenValidity sets Oauth2RememberMeTokenValidity field to given value.

### HasOauth2RememberMeTokenValidity

`func (o *IdentityProviderDTO) HasOauth2RememberMeTokenValidity() bool`

HasOauth2RememberMeTokenValidity returns a boolean if a field has been set.

### GetOauth2TokenValidity

`func (o *IdentityProviderDTO) GetOauth2TokenValidity() int64`

GetOauth2TokenValidity returns the Oauth2TokenValidity field if non-nil, zero value otherwise.

### GetOauth2TokenValidityOk

`func (o *IdentityProviderDTO) GetOauth2TokenValidityOk() (*int64, bool)`

GetOauth2TokenValidityOk returns a tuple with the Oauth2TokenValidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2TokenValidity

`func (o *IdentityProviderDTO) SetOauth2TokenValidity(v int64)`

SetOauth2TokenValidity sets Oauth2TokenValidity field to given value.

### HasOauth2TokenValidity

`func (o *IdentityProviderDTO) HasOauth2TokenValidity() bool`

HasOauth2TokenValidity returns a boolean if a field has been set.

### GetOidcAccessTokenTimeToLive

`func (o *IdentityProviderDTO) GetOidcAccessTokenTimeToLive() int32`

GetOidcAccessTokenTimeToLive returns the OidcAccessTokenTimeToLive field if non-nil, zero value otherwise.

### GetOidcAccessTokenTimeToLiveOk

`func (o *IdentityProviderDTO) GetOidcAccessTokenTimeToLiveOk() (*int32, bool)`

GetOidcAccessTokenTimeToLiveOk returns a tuple with the OidcAccessTokenTimeToLive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcAccessTokenTimeToLive

`func (o *IdentityProviderDTO) SetOidcAccessTokenTimeToLive(v int32)`

SetOidcAccessTokenTimeToLive sets OidcAccessTokenTimeToLive field to given value.

### HasOidcAccessTokenTimeToLive

`func (o *IdentityProviderDTO) HasOidcAccessTokenTimeToLive() bool`

HasOidcAccessTokenTimeToLive returns a boolean if a field has been set.

### GetOidcAuthzCodeTimeToLive

`func (o *IdentityProviderDTO) GetOidcAuthzCodeTimeToLive() int32`

GetOidcAuthzCodeTimeToLive returns the OidcAuthzCodeTimeToLive field if non-nil, zero value otherwise.

### GetOidcAuthzCodeTimeToLiveOk

`func (o *IdentityProviderDTO) GetOidcAuthzCodeTimeToLiveOk() (*int32, bool)`

GetOidcAuthzCodeTimeToLiveOk returns a tuple with the OidcAuthzCodeTimeToLive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcAuthzCodeTimeToLive

`func (o *IdentityProviderDTO) SetOidcAuthzCodeTimeToLive(v int32)`

SetOidcAuthzCodeTimeToLive sets OidcAuthzCodeTimeToLive field to given value.

### HasOidcAuthzCodeTimeToLive

`func (o *IdentityProviderDTO) HasOidcAuthzCodeTimeToLive() bool`

HasOidcAuthzCodeTimeToLive returns a boolean if a field has been set.

### GetOidcIdTokenTimeToLive

`func (o *IdentityProviderDTO) GetOidcIdTokenTimeToLive() int32`

GetOidcIdTokenTimeToLive returns the OidcIdTokenTimeToLive field if non-nil, zero value otherwise.

### GetOidcIdTokenTimeToLiveOk

`func (o *IdentityProviderDTO) GetOidcIdTokenTimeToLiveOk() (*int32, bool)`

GetOidcIdTokenTimeToLiveOk returns a tuple with the OidcIdTokenTimeToLive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcIdTokenTimeToLive

`func (o *IdentityProviderDTO) SetOidcIdTokenTimeToLive(v int32)`

SetOidcIdTokenTimeToLive sets OidcIdTokenTimeToLive field to given value.

### HasOidcIdTokenTimeToLive

`func (o *IdentityProviderDTO) HasOidcIdTokenTimeToLive() bool`

HasOidcIdTokenTimeToLive returns a boolean if a field has been set.

### GetOpenIdEnabled

`func (o *IdentityProviderDTO) GetOpenIdEnabled() bool`

GetOpenIdEnabled returns the OpenIdEnabled field if non-nil, zero value otherwise.

### GetOpenIdEnabledOk

`func (o *IdentityProviderDTO) GetOpenIdEnabledOk() (*bool, bool)`

GetOpenIdEnabledOk returns a tuple with the OpenIdEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenIdEnabled

`func (o *IdentityProviderDTO) SetOpenIdEnabled(v bool)`

SetOpenIdEnabled sets OpenIdEnabled field to given value.

### HasOpenIdEnabled

`func (o *IdentityProviderDTO) HasOpenIdEnabled() bool`

HasOpenIdEnabled returns a boolean if a field has been set.

### GetPwdlessAuthnEnabled

`func (o *IdentityProviderDTO) GetPwdlessAuthnEnabled() bool`

GetPwdlessAuthnEnabled returns the PwdlessAuthnEnabled field if non-nil, zero value otherwise.

### GetPwdlessAuthnEnabledOk

`func (o *IdentityProviderDTO) GetPwdlessAuthnEnabledOk() (*bool, bool)`

GetPwdlessAuthnEnabledOk returns a tuple with the PwdlessAuthnEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPwdlessAuthnEnabled

`func (o *IdentityProviderDTO) SetPwdlessAuthnEnabled(v bool)`

SetPwdlessAuthnEnabled sets PwdlessAuthnEnabled field to given value.

### HasPwdlessAuthnEnabled

`func (o *IdentityProviderDTO) HasPwdlessAuthnEnabled() bool`

HasPwdlessAuthnEnabled returns a boolean if a field has been set.

### GetPwdlessAuthnFrom

`func (o *IdentityProviderDTO) GetPwdlessAuthnFrom() string`

GetPwdlessAuthnFrom returns the PwdlessAuthnFrom field if non-nil, zero value otherwise.

### GetPwdlessAuthnFromOk

`func (o *IdentityProviderDTO) GetPwdlessAuthnFromOk() (*string, bool)`

GetPwdlessAuthnFromOk returns a tuple with the PwdlessAuthnFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPwdlessAuthnFrom

`func (o *IdentityProviderDTO) SetPwdlessAuthnFrom(v string)`

SetPwdlessAuthnFrom sets PwdlessAuthnFrom field to given value.

### HasPwdlessAuthnFrom

`func (o *IdentityProviderDTO) HasPwdlessAuthnFrom() bool`

HasPwdlessAuthnFrom returns a boolean if a field has been set.

### GetPwdlessAuthnSubject

`func (o *IdentityProviderDTO) GetPwdlessAuthnSubject() string`

GetPwdlessAuthnSubject returns the PwdlessAuthnSubject field if non-nil, zero value otherwise.

### GetPwdlessAuthnSubjectOk

`func (o *IdentityProviderDTO) GetPwdlessAuthnSubjectOk() (*string, bool)`

GetPwdlessAuthnSubjectOk returns a tuple with the PwdlessAuthnSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPwdlessAuthnSubject

`func (o *IdentityProviderDTO) SetPwdlessAuthnSubject(v string)`

SetPwdlessAuthnSubject sets PwdlessAuthnSubject field to given value.

### HasPwdlessAuthnSubject

`func (o *IdentityProviderDTO) HasPwdlessAuthnSubject() bool`

HasPwdlessAuthnSubject returns a boolean if a field has been set.

### GetPwdlessAuthnTemplate

`func (o *IdentityProviderDTO) GetPwdlessAuthnTemplate() string`

GetPwdlessAuthnTemplate returns the PwdlessAuthnTemplate field if non-nil, zero value otherwise.

### GetPwdlessAuthnTemplateOk

`func (o *IdentityProviderDTO) GetPwdlessAuthnTemplateOk() (*string, bool)`

GetPwdlessAuthnTemplateOk returns a tuple with the PwdlessAuthnTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPwdlessAuthnTemplate

`func (o *IdentityProviderDTO) SetPwdlessAuthnTemplate(v string)`

SetPwdlessAuthnTemplate sets PwdlessAuthnTemplate field to given value.

### HasPwdlessAuthnTemplate

`func (o *IdentityProviderDTO) HasPwdlessAuthnTemplate() bool`

HasPwdlessAuthnTemplate returns a boolean if a field has been set.

### GetPwdlessAuthnTo

`func (o *IdentityProviderDTO) GetPwdlessAuthnTo() string`

GetPwdlessAuthnTo returns the PwdlessAuthnTo field if non-nil, zero value otherwise.

### GetPwdlessAuthnToOk

`func (o *IdentityProviderDTO) GetPwdlessAuthnToOk() (*string, bool)`

GetPwdlessAuthnToOk returns a tuple with the PwdlessAuthnTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPwdlessAuthnTo

`func (o *IdentityProviderDTO) SetPwdlessAuthnTo(v string)`

SetPwdlessAuthnTo sets PwdlessAuthnTo field to given value.

### HasPwdlessAuthnTo

`func (o *IdentityProviderDTO) HasPwdlessAuthnTo() bool`

HasPwdlessAuthnTo returns a boolean if a field has been set.

### GetRemote

`func (o *IdentityProviderDTO) GetRemote() bool`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *IdentityProviderDTO) GetRemoteOk() (*bool, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *IdentityProviderDTO) SetRemote(v bool)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *IdentityProviderDTO) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### GetRole

`func (o *IdentityProviderDTO) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *IdentityProviderDTO) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *IdentityProviderDTO) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *IdentityProviderDTO) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSessionManagerFactory

`func (o *IdentityProviderDTO) GetSessionManagerFactory() SessionManagerFactoryDTO`

GetSessionManagerFactory returns the SessionManagerFactory field if non-nil, zero value otherwise.

### GetSessionManagerFactoryOk

`func (o *IdentityProviderDTO) GetSessionManagerFactoryOk() (*SessionManagerFactoryDTO, bool)`

GetSessionManagerFactoryOk returns a tuple with the SessionManagerFactory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionManagerFactory

`func (o *IdentityProviderDTO) SetSessionManagerFactory(v SessionManagerFactoryDTO)`

SetSessionManagerFactory sets SessionManagerFactory field to given value.

### HasSessionManagerFactory

`func (o *IdentityProviderDTO) HasSessionManagerFactory() bool`

HasSessionManagerFactory returns a boolean if a field has been set.

### GetSignRequests

`func (o *IdentityProviderDTO) GetSignRequests() bool`

GetSignRequests returns the SignRequests field if non-nil, zero value otherwise.

### GetSignRequestsOk

`func (o *IdentityProviderDTO) GetSignRequestsOk() (*bool, bool)`

GetSignRequestsOk returns a tuple with the SignRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignRequests

`func (o *IdentityProviderDTO) SetSignRequests(v bool)`

SetSignRequests sets SignRequests field to given value.

### HasSignRequests

`func (o *IdentityProviderDTO) HasSignRequests() bool`

HasSignRequests returns a boolean if a field has been set.

### GetSignatureHash

`func (o *IdentityProviderDTO) GetSignatureHash() string`

GetSignatureHash returns the SignatureHash field if non-nil, zero value otherwise.

### GetSignatureHashOk

`func (o *IdentityProviderDTO) GetSignatureHashOk() (*string, bool)`

GetSignatureHashOk returns a tuple with the SignatureHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureHash

`func (o *IdentityProviderDTO) SetSignatureHash(v string)`

SetSignatureHash sets SignatureHash field to given value.

### HasSignatureHash

`func (o *IdentityProviderDTO) HasSignatureHash() bool`

HasSignatureHash returns a boolean if a field has been set.

### GetSsoSessionTimeout

`func (o *IdentityProviderDTO) GetSsoSessionTimeout() int32`

GetSsoSessionTimeout returns the SsoSessionTimeout field if non-nil, zero value otherwise.

### GetSsoSessionTimeoutOk

`func (o *IdentityProviderDTO) GetSsoSessionTimeoutOk() (*int32, bool)`

GetSsoSessionTimeoutOk returns a tuple with the SsoSessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoSessionTimeout

`func (o *IdentityProviderDTO) SetSsoSessionTimeout(v int32)`

SetSsoSessionTimeout sets SsoSessionTimeout field to given value.

### HasSsoSessionTimeout

`func (o *IdentityProviderDTO) HasSsoSessionTimeout() bool`

HasSsoSessionTimeout returns a boolean if a field has been set.

### GetSubjectAuthnPolicies

`func (o *IdentityProviderDTO) GetSubjectAuthnPolicies() []SubjectAuthenticationPolicyDTO`

GetSubjectAuthnPolicies returns the SubjectAuthnPolicies field if non-nil, zero value otherwise.

### GetSubjectAuthnPoliciesOk

`func (o *IdentityProviderDTO) GetSubjectAuthnPoliciesOk() (*[]SubjectAuthenticationPolicyDTO, bool)`

GetSubjectAuthnPoliciesOk returns a tuple with the SubjectAuthnPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectAuthnPolicies

`func (o *IdentityProviderDTO) SetSubjectAuthnPolicies(v []SubjectAuthenticationPolicyDTO)`

SetSubjectAuthnPolicies sets SubjectAuthnPolicies field to given value.

### HasSubjectAuthnPolicies

`func (o *IdentityProviderDTO) HasSubjectAuthnPolicies() bool`

HasSubjectAuthnPolicies returns a boolean if a field has been set.

### GetSubjectNameIDPolicy

`func (o *IdentityProviderDTO) GetSubjectNameIDPolicy() SubjectNameIdentifierPolicyDTO`

GetSubjectNameIDPolicy returns the SubjectNameIDPolicy field if non-nil, zero value otherwise.

### GetSubjectNameIDPolicyOk

`func (o *IdentityProviderDTO) GetSubjectNameIDPolicyOk() (*SubjectNameIdentifierPolicyDTO, bool)`

GetSubjectNameIDPolicyOk returns a tuple with the SubjectNameIDPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectNameIDPolicy

`func (o *IdentityProviderDTO) SetSubjectNameIDPolicy(v SubjectNameIdentifierPolicyDTO)`

SetSubjectNameIDPolicy sets SubjectNameIDPolicy field to given value.

### HasSubjectNameIDPolicy

`func (o *IdentityProviderDTO) HasSubjectNameIDPolicy() bool`

HasSubjectNameIDPolicy returns a boolean if a field has been set.

### GetUserDashboardBranding

`func (o *IdentityProviderDTO) GetUserDashboardBranding() string`

GetUserDashboardBranding returns the UserDashboardBranding field if non-nil, zero value otherwise.

### GetUserDashboardBrandingOk

`func (o *IdentityProviderDTO) GetUserDashboardBrandingOk() (*string, bool)`

GetUserDashboardBrandingOk returns a tuple with the UserDashboardBranding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDashboardBranding

`func (o *IdentityProviderDTO) SetUserDashboardBranding(v string)`

SetUserDashboardBranding sets UserDashboardBranding field to given value.

### HasUserDashboardBranding

`func (o *IdentityProviderDTO) HasUserDashboardBranding() bool`

HasUserDashboardBranding returns a boolean if a field has been set.

### GetWantAuthnRequestsSigned

`func (o *IdentityProviderDTO) GetWantAuthnRequestsSigned() bool`

GetWantAuthnRequestsSigned returns the WantAuthnRequestsSigned field if non-nil, zero value otherwise.

### GetWantAuthnRequestsSignedOk

`func (o *IdentityProviderDTO) GetWantAuthnRequestsSignedOk() (*bool, bool)`

GetWantAuthnRequestsSignedOk returns a tuple with the WantAuthnRequestsSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWantAuthnRequestsSigned

`func (o *IdentityProviderDTO) SetWantAuthnRequestsSigned(v bool)`

SetWantAuthnRequestsSigned sets WantAuthnRequestsSigned field to given value.

### HasWantAuthnRequestsSigned

`func (o *IdentityProviderDTO) HasWantAuthnRequestsSigned() bool`

HasWantAuthnRequestsSigned returns a boolean if a field has been set.

### GetWantSignedRequests

`func (o *IdentityProviderDTO) GetWantSignedRequests() bool`

GetWantSignedRequests returns the WantSignedRequests field if non-nil, zero value otherwise.

### GetWantSignedRequestsOk

`func (o *IdentityProviderDTO) GetWantSignedRequestsOk() (*bool, bool)`

GetWantSignedRequestsOk returns a tuple with the WantSignedRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWantSignedRequests

`func (o *IdentityProviderDTO) SetWantSignedRequests(v bool)`

SetWantSignedRequests sets WantSignedRequests field to given value.

### HasWantSignedRequests

`func (o *IdentityProviderDTO) HasWantSignedRequests() bool`

HasWantSignedRequests returns a boolean if a field has been set.

### GetX

`func (o *IdentityProviderDTO) GetX() float64`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *IdentityProviderDTO) GetXOk() (*float64, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *IdentityProviderDTO) SetX(v float64)`

SetX sets X field to given value.

### HasX

`func (o *IdentityProviderDTO) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *IdentityProviderDTO) GetY() float64`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *IdentityProviderDTO) GetYOk() (*float64, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *IdentityProviderDTO) SetY(v float64)`

SetY sets Y field to given value.

### HasY

`func (o *IdentityProviderDTO) HasY() bool`

HasY returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


