# FacebookOpenIDConnectIdentityProviderDTO

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

### NewFacebookOpenIDConnectIdentityProviderDTO

`func NewFacebookOpenIDConnectIdentityProviderDTO() *FacebookOpenIDConnectIdentityProviderDTO`

NewFacebookOpenIDConnectIdentityProviderDTO instantiates a new FacebookOpenIDConnectIdentityProviderDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFacebookOpenIDConnectIdentityProviderDTOWithDefaults

`func NewFacebookOpenIDConnectIdentityProviderDTOWithDefaults() *FacebookOpenIDConnectIdentityProviderDTO`

NewFacebookOpenIDConnectIdentityProviderDTOWithDefaults instantiates a new FacebookOpenIDConnectIdentityProviderDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessTokenService

`func (o *FacebookOpenIDConnectIdentityProviderDTO) GetAccessTokenService() LocationDTO`

GetAccessTokenService returns the AccessTokenService field if non-nil, zero value otherwise.

### GetAccessTokenServiceOk

`func (o *FacebookOpenIDConnectIdentityProviderDTO) GetAccessTokenServiceOk() (*LocationDTO, bool)`

GetAccessTokenServiceOk returns a tuple with the AccessTokenService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenService

`func (o *FacebookOpenIDConnectIdentityProviderDTO) SetAccessTokenService(v LocationDTO)`

SetAccessTokenService sets AccessTokenService field to given value.

### HasAccessTokenService

`func (o *FacebookOpenIDConnectIdentityProviderDTO) HasAccessTokenService() bool`

HasAccessTokenService returns a boolean if a field has been set.

### GetActiveBindings

`func (o *FacebookOpenIDConnectIdentityProviderDTO) GetActiveBindings() []string`

GetActiveBindings returns the ActiveBindings field if non-nil, zero value otherwise.

### GetActiveBindingsOk

`func (o *FacebookOpenIDConnectIdentityProviderDTO) GetActiveBindingsOk() (*[]string, bool)`

GetActiveBindingsOk returns a tuple with the ActiveBindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveBindings

`func (o *FacebookOpenIDConnectIdentityProviderDTO) SetActiveBindings(v []string)`

SetActiveBindings sets ActiveBindings field to given value.

### HasActiveBindings

`func (o *FacebookOpenIDConnectIdentityProviderDTO) HasActiveBindings() bool`

HasActiveBindings returns a boolean if a field has been set.

### GetActiveProfiles

`func (o *FacebookOpenIDConnectIdentityProviderDTO) GetActiveProfiles() []string`

GetActiveProfiles returns the ActiveProfiles field if non-nil, zero value otherwise.

### GetActiveProfilesOk

`func (o *FacebookOpenIDConnectIdentityProviderDTO) GetActiveProfilesOk() (*[]string, bool)`

GetActiveProfilesOk returns a tuple with the ActiveProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveProfiles

`func (o *FacebookOpenIDConnectIdentityProviderDTO) SetActiveProfiles(v []string)`

SetActiveProfiles sets ActiveProfiles field to given value.

### HasActiveProfiles

`func (o *FacebookOpenIDConnectIdentityProviderDTO) HasActiveProfiles() bool`

HasActiveProfiles returns a boolean if a field has been set.

### GetAuthzTokenService

`func (o *FacebookOpenIDConnectIdentityProviderDTO) GetAuthzTokenService() LocationDTO`

GetAuthzTokenService returns the AuthzTokenService field if non-nil, zero value otherwise.

### GetAuthzTokenServiceOk

`func (o *FacebookOpenIDConnectIdentityProviderDTO) GetAuthzTokenServiceOk() (*LocationDTO, bool)`

GetAuthzTokenServiceOk returns a tuple with the AuthzTokenService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthzTokenService

`func (o *FacebookOpenIDConnectIdentityProviderDTO) SetAuthzTokenService(v LocationDTO)`

SetAuthzTokenService sets AuthzTokenService field to given value.

### HasAuthzTokenService

`func (o *FacebookOpenIDConnectIdentityProviderDTO) HasAuthzTokenService() bool`

HasAuthzTokenService returns a boolean if a field has been set.

### GetClientId

`func (o *FacebookOpenIDConnectIdentityProviderDTO) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *FacebookOpenIDConnectIdentityProviderDTO) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *FacebookOpenIDConnectIdentityProviderDTO) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *FacebookOpenIDConnectIdentityProviderDTO) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *FacebookOpenIDConnectIdentityProviderDTO) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *FacebookOpenIDConnectIdentityProviderDTO) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *FacebookOpenIDConnectIdentityProviderDTO) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *FacebookOpenIDConnectIdentityProviderDTO) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetConfig

`func (o *FacebookOpenIDConnectIdentityProviderDTO) GetConfig() ProviderConfigDTO`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *FacebookOpenIDConnectIdentityProviderDTO) GetConfigOk() (*ProviderConfigDTO, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *FacebookOpenIDConnectIdentityProviderDTO) SetConfig(v ProviderConfigDTO)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *FacebookOpenIDConnectIdentityProviderDTO) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetDescription

`func (o *FacebookOpenIDConnectIdentityProviderDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FacebookOpenIDConnectIdentityProviderDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FacebookOpenIDConnectIdentityProviderDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FacebookOpenIDConnectIdentityProviderDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *FacebookOpenIDConnectIdentityProviderDTO) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *FacebookOpenIDConnectIdentityProviderDTO) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *FacebookOpenIDConnectIdentityProviderDTO) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *FacebookOpenIDConnectIdentityProviderDTO) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetElementId

`func (o *FacebookOpenIDConnectIdentityProviderDTO) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *FacebookOpenIDConnectIdentityProviderDTO) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *FacebookOpenIDConnectIdentityProviderDTO) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *FacebookOpenIDConnectIdentityProviderDTO) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetFederatedConnectionsA

`func (o *FacebookOpenIDConnectIdentityProviderDTO) GetFederatedConnectionsA() []FederatedConnectionDTO`

GetFederatedConnectionsA returns the FederatedConnectionsA field if non-nil, zero value otherwise.

### GetFederatedConnectionsAOk

`func (o *FacebookOpenIDConnectIdentityProviderDTO) GetFederatedConnectionsAOk() (*[]FederatedConnectionDTO, bool)`

GetFederatedConnectionsAOk returns a tuple with the FederatedConnectionsA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederatedConnectionsA

`func (o *FacebookOpenIDConnectIdentityProviderDTO) SetFederatedConnectionsA(v []FederatedConnectionDTO)`

SetFederatedConnectionsA sets FederatedConnectionsA field to given value.

### HasFederatedConnectionsA

`func (o *FacebookOpenIDConnectIdentityProviderDTO) HasFederatedConnectionsA() bool`

HasFederatedConnectionsA returns a boolean if a field has been set.

### GetFederatedConnectionsB

`func (o *FacebookOpenIDConnectIdentityProviderDTO) GetFederatedConnectionsB() []FederatedConnectionDTO`

GetFederatedConnectionsB returns the FederatedConnectionsB field if non-nil, zero value otherwise.

### GetFederatedConnectionsBOk

`func (o *FacebookOpenIDConnectIdentityProviderDTO) GetFederatedConnectionsBOk() (*[]FederatedConnectionDTO, bool)`

GetFederatedConnectionsBOk returns a tuple with the FederatedConnectionsB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederatedConnectionsB

`func (o *FacebookOpenIDConnectIdentityProviderDTO) SetFederatedConnectionsB(v []FederatedConnectionDTO)`

SetFederatedConnectionsB sets FederatedConnectionsB field to given value.

### HasFederatedConnectionsB

`func (o *FacebookOpenIDConnectIdentityProviderDTO) HasFederatedConnectionsB() bool`

HasFederatedConnectionsB returns a boolean if a field has been set.

### GetId

`func (o *FacebookOpenIDConnectIdentityProviderDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FacebookOpenIDConnectIdentityProviderDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FacebookOpenIDConnectIdentityProviderDTO) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *FacebookOpenIDConnectIdentityProviderDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentityAppliance

`func (o *FacebookOpenIDConnectIdentityProviderDTO) GetIdentityAppliance() IdentityApplianceDefinitionDTO`

GetIdentityAppliance returns the IdentityAppliance field if non-nil, zero value otherwise.

### GetIdentityApplianceOk

`func (o *FacebookOpenIDConnectIdentityProviderDTO) GetIdentityApplianceOk() (*IdentityApplianceDefinitionDTO, bool)`

GetIdentityApplianceOk returns a tuple with the IdentityAppliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityAppliance

`func (o *FacebookOpenIDConnectIdentityProviderDTO) SetIdentityAppliance(v IdentityApplianceDefinitionDTO)`

SetIdentityAppliance sets IdentityAppliance field to given value.

### HasIdentityAppliance

`func (o *FacebookOpenIDConnectIdentityProviderDTO) HasIdentityAppliance() bool`

HasIdentityAppliance returns a boolean if a field has been set.

### GetIdentityLookups

`func (o *FacebookOpenIDConnectIdentityProviderDTO) GetIdentityLookups() []IdentityLookupDTO`

GetIdentityLookups returns the IdentityLookups field if non-nil, zero value otherwise.

### GetIdentityLookupsOk

`func (o *FacebookOpenIDConnectIdentityProviderDTO) GetIdentityLookupsOk() (*[]IdentityLookupDTO, bool)`

GetIdentityLookupsOk returns a tuple with the IdentityLookups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityLookups

`func (o *FacebookOpenIDConnectIdentityProviderDTO) SetIdentityLookups(v []IdentityLookupDTO)`

SetIdentityLookups sets IdentityLookups field to given value.

### HasIdentityLookups

`func (o *FacebookOpenIDConnectIdentityProviderDTO) HasIdentityLookups() bool`

HasIdentityLookups returns a boolean if a field has been set.

### GetIsRemote

`func (o *FacebookOpenIDConnectIdentityProviderDTO) GetIsRemote() bool`

GetIsRemote returns the IsRemote field if non-nil, zero value otherwise.

### GetIsRemoteOk

`func (o *FacebookOpenIDConnectIdentityProviderDTO) GetIsRemoteOk() (*bool, bool)`

GetIsRemoteOk returns a tuple with the IsRemote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemote

`func (o *FacebookOpenIDConnectIdentityProviderDTO) SetIsRemote(v bool)`

SetIsRemote sets IsRemote field to given value.

### HasIsRemote

`func (o *FacebookOpenIDConnectIdentityProviderDTO) HasIsRemote() bool`

HasIsRemote returns a boolean if a field has been set.

### GetLocation

`func (o *FacebookOpenIDConnectIdentityProviderDTO) GetLocation() LocationDTO`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *FacebookOpenIDConnectIdentityProviderDTO) GetLocationOk() (*LocationDTO, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *FacebookOpenIDConnectIdentityProviderDTO) SetLocation(v LocationDTO)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *FacebookOpenIDConnectIdentityProviderDTO) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetMetadata

`func (o *FacebookOpenIDConnectIdentityProviderDTO) GetMetadata() ResourceDTO`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FacebookOpenIDConnectIdentityProviderDTO) GetMetadataOk() (*ResourceDTO, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FacebookOpenIDConnectIdentityProviderDTO) SetMetadata(v ResourceDTO)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *FacebookOpenIDConnectIdentityProviderDTO) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMobileAuthzTokenService

`func (o *FacebookOpenIDConnectIdentityProviderDTO) GetMobileAuthzTokenService() LocationDTO`

GetMobileAuthzTokenService returns the MobileAuthzTokenService field if non-nil, zero value otherwise.

### GetMobileAuthzTokenServiceOk

`func (o *FacebookOpenIDConnectIdentityProviderDTO) GetMobileAuthzTokenServiceOk() (*LocationDTO, bool)`

GetMobileAuthzTokenServiceOk returns a tuple with the MobileAuthzTokenService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileAuthzTokenService

`func (o *FacebookOpenIDConnectIdentityProviderDTO) SetMobileAuthzTokenService(v LocationDTO)`

SetMobileAuthzTokenService sets MobileAuthzTokenService field to given value.

### HasMobileAuthzTokenService

`func (o *FacebookOpenIDConnectIdentityProviderDTO) HasMobileAuthzTokenService() bool`

HasMobileAuthzTokenService returns a boolean if a field has been set.

### GetName

`func (o *FacebookOpenIDConnectIdentityProviderDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FacebookOpenIDConnectIdentityProviderDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FacebookOpenIDConnectIdentityProviderDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FacebookOpenIDConnectIdentityProviderDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRemote

`func (o *FacebookOpenIDConnectIdentityProviderDTO) GetRemote() bool`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *FacebookOpenIDConnectIdentityProviderDTO) GetRemoteOk() (*bool, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *FacebookOpenIDConnectIdentityProviderDTO) SetRemote(v bool)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *FacebookOpenIDConnectIdentityProviderDTO) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### GetRole

`func (o *FacebookOpenIDConnectIdentityProviderDTO) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *FacebookOpenIDConnectIdentityProviderDTO) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *FacebookOpenIDConnectIdentityProviderDTO) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *FacebookOpenIDConnectIdentityProviderDTO) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetScopes

`func (o *FacebookOpenIDConnectIdentityProviderDTO) GetScopes() string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *FacebookOpenIDConnectIdentityProviderDTO) GetScopesOk() (*string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *FacebookOpenIDConnectIdentityProviderDTO) SetScopes(v string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *FacebookOpenIDConnectIdentityProviderDTO) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetServerKey

`func (o *FacebookOpenIDConnectIdentityProviderDTO) GetServerKey() string`

GetServerKey returns the ServerKey field if non-nil, zero value otherwise.

### GetServerKeyOk

`func (o *FacebookOpenIDConnectIdentityProviderDTO) GetServerKeyOk() (*string, bool)`

GetServerKeyOk returns a tuple with the ServerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerKey

`func (o *FacebookOpenIDConnectIdentityProviderDTO) SetServerKey(v string)`

SetServerKey sets ServerKey field to given value.

### HasServerKey

`func (o *FacebookOpenIDConnectIdentityProviderDTO) HasServerKey() bool`

HasServerKey returns a boolean if a field has been set.

### GetUserFields

`func (o *FacebookOpenIDConnectIdentityProviderDTO) GetUserFields() string`

GetUserFields returns the UserFields field if non-nil, zero value otherwise.

### GetUserFieldsOk

`func (o *FacebookOpenIDConnectIdentityProviderDTO) GetUserFieldsOk() (*string, bool)`

GetUserFieldsOk returns a tuple with the UserFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserFields

`func (o *FacebookOpenIDConnectIdentityProviderDTO) SetUserFields(v string)`

SetUserFields sets UserFields field to given value.

### HasUserFields

`func (o *FacebookOpenIDConnectIdentityProviderDTO) HasUserFields() bool`

HasUserFields returns a boolean if a field has been set.

### GetX

`func (o *FacebookOpenIDConnectIdentityProviderDTO) GetX() float64`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *FacebookOpenIDConnectIdentityProviderDTO) GetXOk() (*float64, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *FacebookOpenIDConnectIdentityProviderDTO) SetX(v float64)`

SetX sets X field to given value.

### HasX

`func (o *FacebookOpenIDConnectIdentityProviderDTO) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *FacebookOpenIDConnectIdentityProviderDTO) GetY() float64`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *FacebookOpenIDConnectIdentityProviderDTO) GetYOk() (*float64, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *FacebookOpenIDConnectIdentityProviderDTO) SetY(v float64)`

SetY sets Y field to given value.

### HasY

`func (o *FacebookOpenIDConnectIdentityProviderDTO) HasY() bool`

HasY returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


