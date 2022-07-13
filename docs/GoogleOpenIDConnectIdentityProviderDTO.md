# GoogleOpenIDConnectIdentityProviderDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessTokenService** | Pointer to [**LocationDTO**](LocationDTO.md) |  | [optional] 
**ActiveBindings** | Pointer to **[]string** |  | [optional] 
**ActiveProfiles** | Pointer to **[]string** |  | [optional] 
**AuthzTokenService** | Pointer to [**LocationDTO**](LocationDTO.md) |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**ClientSecret** | Pointer to **string** |  | [optional] 
**Config** | Pointer to [**ProviderConfigDTO**](ProviderConfigDTO.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**FederatedConnectionsA** | Pointer to [**[]FederatedConnectionDTO**](FederatedConnectionDTO.md) |  | [optional] 
**FederatedConnectionsB** | Pointer to [**[]FederatedConnectionDTO**](FederatedConnectionDTO.md) |  | [optional] 
**GoogleAppsDomain** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**IdentityAppliance** | Pointer to [**IdentityApplianceDefinitionDTO**](IdentityApplianceDefinitionDTO.md) |  | [optional] 
**IdentityLookups** | Pointer to [**[]IdentityLookupDTO**](IdentityLookupDTO.md) |  | [optional] 
**IsRemote** | Pointer to **bool** |  | [optional] 
**Location** | Pointer to [**LocationDTO**](LocationDTO.md) |  | [optional] 
**Metadata** | Pointer to [**ResourceDTO**](ResourceDTO.md) |  | [optional] 
**MobileAuthzTokenService** | Pointer to [**LocationDTO**](LocationDTO.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Remote** | Pointer to **bool** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**Scopes** | Pointer to **string** |  | [optional] 
**ServerKey** | Pointer to **string** |  | [optional] 
**UserFields** | Pointer to **string** |  | [optional] 
**X** | Pointer to **float64** |  | [optional] 
**Y** | Pointer to **float64** |  | [optional] 

## Methods

### NewGoogleOpenIDConnectIdentityProviderDTO

`func NewGoogleOpenIDConnectIdentityProviderDTO() *GoogleOpenIDConnectIdentityProviderDTO`

NewGoogleOpenIDConnectIdentityProviderDTO instantiates a new GoogleOpenIDConnectIdentityProviderDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleOpenIDConnectIdentityProviderDTOWithDefaults

`func NewGoogleOpenIDConnectIdentityProviderDTOWithDefaults() *GoogleOpenIDConnectIdentityProviderDTO`

NewGoogleOpenIDConnectIdentityProviderDTOWithDefaults instantiates a new GoogleOpenIDConnectIdentityProviderDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessTokenService

`func (o *GoogleOpenIDConnectIdentityProviderDTO) GetAccessTokenService() LocationDTO`

GetAccessTokenService returns the AccessTokenService field if non-nil, zero value otherwise.

### GetAccessTokenServiceOk

`func (o *GoogleOpenIDConnectIdentityProviderDTO) GetAccessTokenServiceOk() (*LocationDTO, bool)`

GetAccessTokenServiceOk returns a tuple with the AccessTokenService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenService

`func (o *GoogleOpenIDConnectIdentityProviderDTO) SetAccessTokenService(v LocationDTO)`

SetAccessTokenService sets AccessTokenService field to given value.

### HasAccessTokenService

`func (o *GoogleOpenIDConnectIdentityProviderDTO) HasAccessTokenService() bool`

HasAccessTokenService returns a boolean if a field has been set.

### GetActiveBindings

`func (o *GoogleOpenIDConnectIdentityProviderDTO) GetActiveBindings() []string`

GetActiveBindings returns the ActiveBindings field if non-nil, zero value otherwise.

### GetActiveBindingsOk

`func (o *GoogleOpenIDConnectIdentityProviderDTO) GetActiveBindingsOk() (*[]string, bool)`

GetActiveBindingsOk returns a tuple with the ActiveBindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveBindings

`func (o *GoogleOpenIDConnectIdentityProviderDTO) SetActiveBindings(v []string)`

SetActiveBindings sets ActiveBindings field to given value.

### HasActiveBindings

`func (o *GoogleOpenIDConnectIdentityProviderDTO) HasActiveBindings() bool`

HasActiveBindings returns a boolean if a field has been set.

### GetActiveProfiles

`func (o *GoogleOpenIDConnectIdentityProviderDTO) GetActiveProfiles() []string`

GetActiveProfiles returns the ActiveProfiles field if non-nil, zero value otherwise.

### GetActiveProfilesOk

`func (o *GoogleOpenIDConnectIdentityProviderDTO) GetActiveProfilesOk() (*[]string, bool)`

GetActiveProfilesOk returns a tuple with the ActiveProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveProfiles

`func (o *GoogleOpenIDConnectIdentityProviderDTO) SetActiveProfiles(v []string)`

SetActiveProfiles sets ActiveProfiles field to given value.

### HasActiveProfiles

`func (o *GoogleOpenIDConnectIdentityProviderDTO) HasActiveProfiles() bool`

HasActiveProfiles returns a boolean if a field has been set.

### GetAuthzTokenService

`func (o *GoogleOpenIDConnectIdentityProviderDTO) GetAuthzTokenService() LocationDTO`

GetAuthzTokenService returns the AuthzTokenService field if non-nil, zero value otherwise.

### GetAuthzTokenServiceOk

`func (o *GoogleOpenIDConnectIdentityProviderDTO) GetAuthzTokenServiceOk() (*LocationDTO, bool)`

GetAuthzTokenServiceOk returns a tuple with the AuthzTokenService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthzTokenService

`func (o *GoogleOpenIDConnectIdentityProviderDTO) SetAuthzTokenService(v LocationDTO)`

SetAuthzTokenService sets AuthzTokenService field to given value.

### HasAuthzTokenService

`func (o *GoogleOpenIDConnectIdentityProviderDTO) HasAuthzTokenService() bool`

HasAuthzTokenService returns a boolean if a field has been set.

### GetClientId

`func (o *GoogleOpenIDConnectIdentityProviderDTO) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *GoogleOpenIDConnectIdentityProviderDTO) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *GoogleOpenIDConnectIdentityProviderDTO) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *GoogleOpenIDConnectIdentityProviderDTO) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *GoogleOpenIDConnectIdentityProviderDTO) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *GoogleOpenIDConnectIdentityProviderDTO) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *GoogleOpenIDConnectIdentityProviderDTO) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *GoogleOpenIDConnectIdentityProviderDTO) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetConfig

`func (o *GoogleOpenIDConnectIdentityProviderDTO) GetConfig() ProviderConfigDTO`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GoogleOpenIDConnectIdentityProviderDTO) GetConfigOk() (*ProviderConfigDTO, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GoogleOpenIDConnectIdentityProviderDTO) SetConfig(v ProviderConfigDTO)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GoogleOpenIDConnectIdentityProviderDTO) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetDescription

`func (o *GoogleOpenIDConnectIdentityProviderDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GoogleOpenIDConnectIdentityProviderDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GoogleOpenIDConnectIdentityProviderDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GoogleOpenIDConnectIdentityProviderDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *GoogleOpenIDConnectIdentityProviderDTO) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GoogleOpenIDConnectIdentityProviderDTO) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GoogleOpenIDConnectIdentityProviderDTO) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GoogleOpenIDConnectIdentityProviderDTO) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetElementId

`func (o *GoogleOpenIDConnectIdentityProviderDTO) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *GoogleOpenIDConnectIdentityProviderDTO) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *GoogleOpenIDConnectIdentityProviderDTO) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *GoogleOpenIDConnectIdentityProviderDTO) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetFederatedConnectionsA

`func (o *GoogleOpenIDConnectIdentityProviderDTO) GetFederatedConnectionsA() []FederatedConnectionDTO`

GetFederatedConnectionsA returns the FederatedConnectionsA field if non-nil, zero value otherwise.

### GetFederatedConnectionsAOk

`func (o *GoogleOpenIDConnectIdentityProviderDTO) GetFederatedConnectionsAOk() (*[]FederatedConnectionDTO, bool)`

GetFederatedConnectionsAOk returns a tuple with the FederatedConnectionsA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederatedConnectionsA

`func (o *GoogleOpenIDConnectIdentityProviderDTO) SetFederatedConnectionsA(v []FederatedConnectionDTO)`

SetFederatedConnectionsA sets FederatedConnectionsA field to given value.

### HasFederatedConnectionsA

`func (o *GoogleOpenIDConnectIdentityProviderDTO) HasFederatedConnectionsA() bool`

HasFederatedConnectionsA returns a boolean if a field has been set.

### GetFederatedConnectionsB

`func (o *GoogleOpenIDConnectIdentityProviderDTO) GetFederatedConnectionsB() []FederatedConnectionDTO`

GetFederatedConnectionsB returns the FederatedConnectionsB field if non-nil, zero value otherwise.

### GetFederatedConnectionsBOk

`func (o *GoogleOpenIDConnectIdentityProviderDTO) GetFederatedConnectionsBOk() (*[]FederatedConnectionDTO, bool)`

GetFederatedConnectionsBOk returns a tuple with the FederatedConnectionsB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederatedConnectionsB

`func (o *GoogleOpenIDConnectIdentityProviderDTO) SetFederatedConnectionsB(v []FederatedConnectionDTO)`

SetFederatedConnectionsB sets FederatedConnectionsB field to given value.

### HasFederatedConnectionsB

`func (o *GoogleOpenIDConnectIdentityProviderDTO) HasFederatedConnectionsB() bool`

HasFederatedConnectionsB returns a boolean if a field has been set.

### GetGoogleAppsDomain

`func (o *GoogleOpenIDConnectIdentityProviderDTO) GetGoogleAppsDomain() string`

GetGoogleAppsDomain returns the GoogleAppsDomain field if non-nil, zero value otherwise.

### GetGoogleAppsDomainOk

`func (o *GoogleOpenIDConnectIdentityProviderDTO) GetGoogleAppsDomainOk() (*string, bool)`

GetGoogleAppsDomainOk returns a tuple with the GoogleAppsDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleAppsDomain

`func (o *GoogleOpenIDConnectIdentityProviderDTO) SetGoogleAppsDomain(v string)`

SetGoogleAppsDomain sets GoogleAppsDomain field to given value.

### HasGoogleAppsDomain

`func (o *GoogleOpenIDConnectIdentityProviderDTO) HasGoogleAppsDomain() bool`

HasGoogleAppsDomain returns a boolean if a field has been set.

### GetId

`func (o *GoogleOpenIDConnectIdentityProviderDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GoogleOpenIDConnectIdentityProviderDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GoogleOpenIDConnectIdentityProviderDTO) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GoogleOpenIDConnectIdentityProviderDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentityAppliance

`func (o *GoogleOpenIDConnectIdentityProviderDTO) GetIdentityAppliance() IdentityApplianceDefinitionDTO`

GetIdentityAppliance returns the IdentityAppliance field if non-nil, zero value otherwise.

### GetIdentityApplianceOk

`func (o *GoogleOpenIDConnectIdentityProviderDTO) GetIdentityApplianceOk() (*IdentityApplianceDefinitionDTO, bool)`

GetIdentityApplianceOk returns a tuple with the IdentityAppliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityAppliance

`func (o *GoogleOpenIDConnectIdentityProviderDTO) SetIdentityAppliance(v IdentityApplianceDefinitionDTO)`

SetIdentityAppliance sets IdentityAppliance field to given value.

### HasIdentityAppliance

`func (o *GoogleOpenIDConnectIdentityProviderDTO) HasIdentityAppliance() bool`

HasIdentityAppliance returns a boolean if a field has been set.

### GetIdentityLookups

`func (o *GoogleOpenIDConnectIdentityProviderDTO) GetIdentityLookups() []IdentityLookupDTO`

GetIdentityLookups returns the IdentityLookups field if non-nil, zero value otherwise.

### GetIdentityLookupsOk

`func (o *GoogleOpenIDConnectIdentityProviderDTO) GetIdentityLookupsOk() (*[]IdentityLookupDTO, bool)`

GetIdentityLookupsOk returns a tuple with the IdentityLookups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityLookups

`func (o *GoogleOpenIDConnectIdentityProviderDTO) SetIdentityLookups(v []IdentityLookupDTO)`

SetIdentityLookups sets IdentityLookups field to given value.

### HasIdentityLookups

`func (o *GoogleOpenIDConnectIdentityProviderDTO) HasIdentityLookups() bool`

HasIdentityLookups returns a boolean if a field has been set.

### GetIsRemote

`func (o *GoogleOpenIDConnectIdentityProviderDTO) GetIsRemote() bool`

GetIsRemote returns the IsRemote field if non-nil, zero value otherwise.

### GetIsRemoteOk

`func (o *GoogleOpenIDConnectIdentityProviderDTO) GetIsRemoteOk() (*bool, bool)`

GetIsRemoteOk returns a tuple with the IsRemote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemote

`func (o *GoogleOpenIDConnectIdentityProviderDTO) SetIsRemote(v bool)`

SetIsRemote sets IsRemote field to given value.

### HasIsRemote

`func (o *GoogleOpenIDConnectIdentityProviderDTO) HasIsRemote() bool`

HasIsRemote returns a boolean if a field has been set.

### GetLocation

`func (o *GoogleOpenIDConnectIdentityProviderDTO) GetLocation() LocationDTO`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *GoogleOpenIDConnectIdentityProviderDTO) GetLocationOk() (*LocationDTO, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *GoogleOpenIDConnectIdentityProviderDTO) SetLocation(v LocationDTO)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *GoogleOpenIDConnectIdentityProviderDTO) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetMetadata

`func (o *GoogleOpenIDConnectIdentityProviderDTO) GetMetadata() ResourceDTO`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GoogleOpenIDConnectIdentityProviderDTO) GetMetadataOk() (*ResourceDTO, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GoogleOpenIDConnectIdentityProviderDTO) SetMetadata(v ResourceDTO)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GoogleOpenIDConnectIdentityProviderDTO) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMobileAuthzTokenService

`func (o *GoogleOpenIDConnectIdentityProviderDTO) GetMobileAuthzTokenService() LocationDTO`

GetMobileAuthzTokenService returns the MobileAuthzTokenService field if non-nil, zero value otherwise.

### GetMobileAuthzTokenServiceOk

`func (o *GoogleOpenIDConnectIdentityProviderDTO) GetMobileAuthzTokenServiceOk() (*LocationDTO, bool)`

GetMobileAuthzTokenServiceOk returns a tuple with the MobileAuthzTokenService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileAuthzTokenService

`func (o *GoogleOpenIDConnectIdentityProviderDTO) SetMobileAuthzTokenService(v LocationDTO)`

SetMobileAuthzTokenService sets MobileAuthzTokenService field to given value.

### HasMobileAuthzTokenService

`func (o *GoogleOpenIDConnectIdentityProviderDTO) HasMobileAuthzTokenService() bool`

HasMobileAuthzTokenService returns a boolean if a field has been set.

### GetName

`func (o *GoogleOpenIDConnectIdentityProviderDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GoogleOpenIDConnectIdentityProviderDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GoogleOpenIDConnectIdentityProviderDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GoogleOpenIDConnectIdentityProviderDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRemote

`func (o *GoogleOpenIDConnectIdentityProviderDTO) GetRemote() bool`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *GoogleOpenIDConnectIdentityProviderDTO) GetRemoteOk() (*bool, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *GoogleOpenIDConnectIdentityProviderDTO) SetRemote(v bool)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *GoogleOpenIDConnectIdentityProviderDTO) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### GetRole

`func (o *GoogleOpenIDConnectIdentityProviderDTO) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *GoogleOpenIDConnectIdentityProviderDTO) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *GoogleOpenIDConnectIdentityProviderDTO) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *GoogleOpenIDConnectIdentityProviderDTO) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetScopes

`func (o *GoogleOpenIDConnectIdentityProviderDTO) GetScopes() string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *GoogleOpenIDConnectIdentityProviderDTO) GetScopesOk() (*string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *GoogleOpenIDConnectIdentityProviderDTO) SetScopes(v string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *GoogleOpenIDConnectIdentityProviderDTO) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetServerKey

`func (o *GoogleOpenIDConnectIdentityProviderDTO) GetServerKey() string`

GetServerKey returns the ServerKey field if non-nil, zero value otherwise.

### GetServerKeyOk

`func (o *GoogleOpenIDConnectIdentityProviderDTO) GetServerKeyOk() (*string, bool)`

GetServerKeyOk returns a tuple with the ServerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerKey

`func (o *GoogleOpenIDConnectIdentityProviderDTO) SetServerKey(v string)`

SetServerKey sets ServerKey field to given value.

### HasServerKey

`func (o *GoogleOpenIDConnectIdentityProviderDTO) HasServerKey() bool`

HasServerKey returns a boolean if a field has been set.

### GetUserFields

`func (o *GoogleOpenIDConnectIdentityProviderDTO) GetUserFields() string`

GetUserFields returns the UserFields field if non-nil, zero value otherwise.

### GetUserFieldsOk

`func (o *GoogleOpenIDConnectIdentityProviderDTO) GetUserFieldsOk() (*string, bool)`

GetUserFieldsOk returns a tuple with the UserFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserFields

`func (o *GoogleOpenIDConnectIdentityProviderDTO) SetUserFields(v string)`

SetUserFields sets UserFields field to given value.

### HasUserFields

`func (o *GoogleOpenIDConnectIdentityProviderDTO) HasUserFields() bool`

HasUserFields returns a boolean if a field has been set.

### GetX

`func (o *GoogleOpenIDConnectIdentityProviderDTO) GetX() float64`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *GoogleOpenIDConnectIdentityProviderDTO) GetXOk() (*float64, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *GoogleOpenIDConnectIdentityProviderDTO) SetX(v float64)`

SetX sets X field to given value.

### HasX

`func (o *GoogleOpenIDConnectIdentityProviderDTO) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *GoogleOpenIDConnectIdentityProviderDTO) GetY() float64`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *GoogleOpenIDConnectIdentityProviderDTO) GetYOk() (*float64, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *GoogleOpenIDConnectIdentityProviderDTO) SetY(v float64)`

SetY sets Y field to given value.

### HasY

`func (o *GoogleOpenIDConnectIdentityProviderDTO) HasY() bool`

HasY returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


