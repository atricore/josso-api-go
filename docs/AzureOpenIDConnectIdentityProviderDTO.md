# AzureOpenIDConnectIdentityProviderDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessTokenService** | Pointer to [**LocationDTO**](LocationDTO.md) |  | [optional] 
**ActiveBindings** | Pointer to **[]string** |  | [optional] 
**ActiveProfiles** | Pointer to **[]string** |  | [optional] 
**AuthzTokenService** | Pointer to [**LocationDTO**](LocationDTO.md) |  | [optional] 
**AzObjectId** | Pointer to **string** |  | [optional] 
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

### NewAzureOpenIDConnectIdentityProviderDTO

`func NewAzureOpenIDConnectIdentityProviderDTO() *AzureOpenIDConnectIdentityProviderDTO`

NewAzureOpenIDConnectIdentityProviderDTO instantiates a new AzureOpenIDConnectIdentityProviderDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureOpenIDConnectIdentityProviderDTOWithDefaults

`func NewAzureOpenIDConnectIdentityProviderDTOWithDefaults() *AzureOpenIDConnectIdentityProviderDTO`

NewAzureOpenIDConnectIdentityProviderDTOWithDefaults instantiates a new AzureOpenIDConnectIdentityProviderDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessTokenService

`func (o *AzureOpenIDConnectIdentityProviderDTO) GetAccessTokenService() LocationDTO`

GetAccessTokenService returns the AccessTokenService field if non-nil, zero value otherwise.

### GetAccessTokenServiceOk

`func (o *AzureOpenIDConnectIdentityProviderDTO) GetAccessTokenServiceOk() (*LocationDTO, bool)`

GetAccessTokenServiceOk returns a tuple with the AccessTokenService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenService

`func (o *AzureOpenIDConnectIdentityProviderDTO) SetAccessTokenService(v LocationDTO)`

SetAccessTokenService sets AccessTokenService field to given value.

### HasAccessTokenService

`func (o *AzureOpenIDConnectIdentityProviderDTO) HasAccessTokenService() bool`

HasAccessTokenService returns a boolean if a field has been set.

### GetActiveBindings

`func (o *AzureOpenIDConnectIdentityProviderDTO) GetActiveBindings() []string`

GetActiveBindings returns the ActiveBindings field if non-nil, zero value otherwise.

### GetActiveBindingsOk

`func (o *AzureOpenIDConnectIdentityProviderDTO) GetActiveBindingsOk() (*[]string, bool)`

GetActiveBindingsOk returns a tuple with the ActiveBindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveBindings

`func (o *AzureOpenIDConnectIdentityProviderDTO) SetActiveBindings(v []string)`

SetActiveBindings sets ActiveBindings field to given value.

### HasActiveBindings

`func (o *AzureOpenIDConnectIdentityProviderDTO) HasActiveBindings() bool`

HasActiveBindings returns a boolean if a field has been set.

### GetActiveProfiles

`func (o *AzureOpenIDConnectIdentityProviderDTO) GetActiveProfiles() []string`

GetActiveProfiles returns the ActiveProfiles field if non-nil, zero value otherwise.

### GetActiveProfilesOk

`func (o *AzureOpenIDConnectIdentityProviderDTO) GetActiveProfilesOk() (*[]string, bool)`

GetActiveProfilesOk returns a tuple with the ActiveProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveProfiles

`func (o *AzureOpenIDConnectIdentityProviderDTO) SetActiveProfiles(v []string)`

SetActiveProfiles sets ActiveProfiles field to given value.

### HasActiveProfiles

`func (o *AzureOpenIDConnectIdentityProviderDTO) HasActiveProfiles() bool`

HasActiveProfiles returns a boolean if a field has been set.

### GetAuthzTokenService

`func (o *AzureOpenIDConnectIdentityProviderDTO) GetAuthzTokenService() LocationDTO`

GetAuthzTokenService returns the AuthzTokenService field if non-nil, zero value otherwise.

### GetAuthzTokenServiceOk

`func (o *AzureOpenIDConnectIdentityProviderDTO) GetAuthzTokenServiceOk() (*LocationDTO, bool)`

GetAuthzTokenServiceOk returns a tuple with the AuthzTokenService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthzTokenService

`func (o *AzureOpenIDConnectIdentityProviderDTO) SetAuthzTokenService(v LocationDTO)`

SetAuthzTokenService sets AuthzTokenService field to given value.

### HasAuthzTokenService

`func (o *AzureOpenIDConnectIdentityProviderDTO) HasAuthzTokenService() bool`

HasAuthzTokenService returns a boolean if a field has been set.

### GetAzObjectId

`func (o *AzureOpenIDConnectIdentityProviderDTO) GetAzObjectId() string`

GetAzObjectId returns the AzObjectId field if non-nil, zero value otherwise.

### GetAzObjectIdOk

`func (o *AzureOpenIDConnectIdentityProviderDTO) GetAzObjectIdOk() (*string, bool)`

GetAzObjectIdOk returns a tuple with the AzObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzObjectId

`func (o *AzureOpenIDConnectIdentityProviderDTO) SetAzObjectId(v string)`

SetAzObjectId sets AzObjectId field to given value.

### HasAzObjectId

`func (o *AzureOpenIDConnectIdentityProviderDTO) HasAzObjectId() bool`

HasAzObjectId returns a boolean if a field has been set.

### GetClientId

`func (o *AzureOpenIDConnectIdentityProviderDTO) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *AzureOpenIDConnectIdentityProviderDTO) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *AzureOpenIDConnectIdentityProviderDTO) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *AzureOpenIDConnectIdentityProviderDTO) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *AzureOpenIDConnectIdentityProviderDTO) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *AzureOpenIDConnectIdentityProviderDTO) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *AzureOpenIDConnectIdentityProviderDTO) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *AzureOpenIDConnectIdentityProviderDTO) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetConfig

`func (o *AzureOpenIDConnectIdentityProviderDTO) GetConfig() ProviderConfigDTO`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AzureOpenIDConnectIdentityProviderDTO) GetConfigOk() (*ProviderConfigDTO, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AzureOpenIDConnectIdentityProviderDTO) SetConfig(v ProviderConfigDTO)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AzureOpenIDConnectIdentityProviderDTO) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetDescription

`func (o *AzureOpenIDConnectIdentityProviderDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AzureOpenIDConnectIdentityProviderDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AzureOpenIDConnectIdentityProviderDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AzureOpenIDConnectIdentityProviderDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *AzureOpenIDConnectIdentityProviderDTO) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AzureOpenIDConnectIdentityProviderDTO) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AzureOpenIDConnectIdentityProviderDTO) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AzureOpenIDConnectIdentityProviderDTO) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetElementId

`func (o *AzureOpenIDConnectIdentityProviderDTO) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *AzureOpenIDConnectIdentityProviderDTO) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *AzureOpenIDConnectIdentityProviderDTO) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *AzureOpenIDConnectIdentityProviderDTO) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetFederatedConnectionsA

`func (o *AzureOpenIDConnectIdentityProviderDTO) GetFederatedConnectionsA() []FederatedConnectionDTO`

GetFederatedConnectionsA returns the FederatedConnectionsA field if non-nil, zero value otherwise.

### GetFederatedConnectionsAOk

`func (o *AzureOpenIDConnectIdentityProviderDTO) GetFederatedConnectionsAOk() (*[]FederatedConnectionDTO, bool)`

GetFederatedConnectionsAOk returns a tuple with the FederatedConnectionsA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederatedConnectionsA

`func (o *AzureOpenIDConnectIdentityProviderDTO) SetFederatedConnectionsA(v []FederatedConnectionDTO)`

SetFederatedConnectionsA sets FederatedConnectionsA field to given value.

### HasFederatedConnectionsA

`func (o *AzureOpenIDConnectIdentityProviderDTO) HasFederatedConnectionsA() bool`

HasFederatedConnectionsA returns a boolean if a field has been set.

### GetFederatedConnectionsB

`func (o *AzureOpenIDConnectIdentityProviderDTO) GetFederatedConnectionsB() []FederatedConnectionDTO`

GetFederatedConnectionsB returns the FederatedConnectionsB field if non-nil, zero value otherwise.

### GetFederatedConnectionsBOk

`func (o *AzureOpenIDConnectIdentityProviderDTO) GetFederatedConnectionsBOk() (*[]FederatedConnectionDTO, bool)`

GetFederatedConnectionsBOk returns a tuple with the FederatedConnectionsB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederatedConnectionsB

`func (o *AzureOpenIDConnectIdentityProviderDTO) SetFederatedConnectionsB(v []FederatedConnectionDTO)`

SetFederatedConnectionsB sets FederatedConnectionsB field to given value.

### HasFederatedConnectionsB

`func (o *AzureOpenIDConnectIdentityProviderDTO) HasFederatedConnectionsB() bool`

HasFederatedConnectionsB returns a boolean if a field has been set.

### GetId

`func (o *AzureOpenIDConnectIdentityProviderDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AzureOpenIDConnectIdentityProviderDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AzureOpenIDConnectIdentityProviderDTO) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AzureOpenIDConnectIdentityProviderDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentityAppliance

`func (o *AzureOpenIDConnectIdentityProviderDTO) GetIdentityAppliance() IdentityApplianceDefinitionDTO`

GetIdentityAppliance returns the IdentityAppliance field if non-nil, zero value otherwise.

### GetIdentityApplianceOk

`func (o *AzureOpenIDConnectIdentityProviderDTO) GetIdentityApplianceOk() (*IdentityApplianceDefinitionDTO, bool)`

GetIdentityApplianceOk returns a tuple with the IdentityAppliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityAppliance

`func (o *AzureOpenIDConnectIdentityProviderDTO) SetIdentityAppliance(v IdentityApplianceDefinitionDTO)`

SetIdentityAppliance sets IdentityAppliance field to given value.

### HasIdentityAppliance

`func (o *AzureOpenIDConnectIdentityProviderDTO) HasIdentityAppliance() bool`

HasIdentityAppliance returns a boolean if a field has been set.

### GetIdentityLookups

`func (o *AzureOpenIDConnectIdentityProviderDTO) GetIdentityLookups() []IdentityLookupDTO`

GetIdentityLookups returns the IdentityLookups field if non-nil, zero value otherwise.

### GetIdentityLookupsOk

`func (o *AzureOpenIDConnectIdentityProviderDTO) GetIdentityLookupsOk() (*[]IdentityLookupDTO, bool)`

GetIdentityLookupsOk returns a tuple with the IdentityLookups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityLookups

`func (o *AzureOpenIDConnectIdentityProviderDTO) SetIdentityLookups(v []IdentityLookupDTO)`

SetIdentityLookups sets IdentityLookups field to given value.

### HasIdentityLookups

`func (o *AzureOpenIDConnectIdentityProviderDTO) HasIdentityLookups() bool`

HasIdentityLookups returns a boolean if a field has been set.

### GetIsRemote

`func (o *AzureOpenIDConnectIdentityProviderDTO) GetIsRemote() bool`

GetIsRemote returns the IsRemote field if non-nil, zero value otherwise.

### GetIsRemoteOk

`func (o *AzureOpenIDConnectIdentityProviderDTO) GetIsRemoteOk() (*bool, bool)`

GetIsRemoteOk returns a tuple with the IsRemote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemote

`func (o *AzureOpenIDConnectIdentityProviderDTO) SetIsRemote(v bool)`

SetIsRemote sets IsRemote field to given value.

### HasIsRemote

`func (o *AzureOpenIDConnectIdentityProviderDTO) HasIsRemote() bool`

HasIsRemote returns a boolean if a field has been set.

### GetLocation

`func (o *AzureOpenIDConnectIdentityProviderDTO) GetLocation() LocationDTO`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *AzureOpenIDConnectIdentityProviderDTO) GetLocationOk() (*LocationDTO, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *AzureOpenIDConnectIdentityProviderDTO) SetLocation(v LocationDTO)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *AzureOpenIDConnectIdentityProviderDTO) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetMetadata

`func (o *AzureOpenIDConnectIdentityProviderDTO) GetMetadata() ResourceDTO`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AzureOpenIDConnectIdentityProviderDTO) GetMetadataOk() (*ResourceDTO, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AzureOpenIDConnectIdentityProviderDTO) SetMetadata(v ResourceDTO)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AzureOpenIDConnectIdentityProviderDTO) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMobileAuthzTokenService

`func (o *AzureOpenIDConnectIdentityProviderDTO) GetMobileAuthzTokenService() LocationDTO`

GetMobileAuthzTokenService returns the MobileAuthzTokenService field if non-nil, zero value otherwise.

### GetMobileAuthzTokenServiceOk

`func (o *AzureOpenIDConnectIdentityProviderDTO) GetMobileAuthzTokenServiceOk() (*LocationDTO, bool)`

GetMobileAuthzTokenServiceOk returns a tuple with the MobileAuthzTokenService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileAuthzTokenService

`func (o *AzureOpenIDConnectIdentityProviderDTO) SetMobileAuthzTokenService(v LocationDTO)`

SetMobileAuthzTokenService sets MobileAuthzTokenService field to given value.

### HasMobileAuthzTokenService

`func (o *AzureOpenIDConnectIdentityProviderDTO) HasMobileAuthzTokenService() bool`

HasMobileAuthzTokenService returns a boolean if a field has been set.

### GetName

`func (o *AzureOpenIDConnectIdentityProviderDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AzureOpenIDConnectIdentityProviderDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AzureOpenIDConnectIdentityProviderDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AzureOpenIDConnectIdentityProviderDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRemote

`func (o *AzureOpenIDConnectIdentityProviderDTO) GetRemote() bool`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *AzureOpenIDConnectIdentityProviderDTO) GetRemoteOk() (*bool, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *AzureOpenIDConnectIdentityProviderDTO) SetRemote(v bool)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *AzureOpenIDConnectIdentityProviderDTO) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### GetRole

`func (o *AzureOpenIDConnectIdentityProviderDTO) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AzureOpenIDConnectIdentityProviderDTO) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AzureOpenIDConnectIdentityProviderDTO) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *AzureOpenIDConnectIdentityProviderDTO) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetScopes

`func (o *AzureOpenIDConnectIdentityProviderDTO) GetScopes() string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *AzureOpenIDConnectIdentityProviderDTO) GetScopesOk() (*string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *AzureOpenIDConnectIdentityProviderDTO) SetScopes(v string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *AzureOpenIDConnectIdentityProviderDTO) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetServerKey

`func (o *AzureOpenIDConnectIdentityProviderDTO) GetServerKey() string`

GetServerKey returns the ServerKey field if non-nil, zero value otherwise.

### GetServerKeyOk

`func (o *AzureOpenIDConnectIdentityProviderDTO) GetServerKeyOk() (*string, bool)`

GetServerKeyOk returns a tuple with the ServerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerKey

`func (o *AzureOpenIDConnectIdentityProviderDTO) SetServerKey(v string)`

SetServerKey sets ServerKey field to given value.

### HasServerKey

`func (o *AzureOpenIDConnectIdentityProviderDTO) HasServerKey() bool`

HasServerKey returns a boolean if a field has been set.

### GetUserFields

`func (o *AzureOpenIDConnectIdentityProviderDTO) GetUserFields() string`

GetUserFields returns the UserFields field if non-nil, zero value otherwise.

### GetUserFieldsOk

`func (o *AzureOpenIDConnectIdentityProviderDTO) GetUserFieldsOk() (*string, bool)`

GetUserFieldsOk returns a tuple with the UserFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserFields

`func (o *AzureOpenIDConnectIdentityProviderDTO) SetUserFields(v string)`

SetUserFields sets UserFields field to given value.

### HasUserFields

`func (o *AzureOpenIDConnectIdentityProviderDTO) HasUserFields() bool`

HasUserFields returns a boolean if a field has been set.

### GetX

`func (o *AzureOpenIDConnectIdentityProviderDTO) GetX() float64`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *AzureOpenIDConnectIdentityProviderDTO) GetXOk() (*float64, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *AzureOpenIDConnectIdentityProviderDTO) SetX(v float64)`

SetX sets X field to given value.

### HasX

`func (o *AzureOpenIDConnectIdentityProviderDTO) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *AzureOpenIDConnectIdentityProviderDTO) GetY() float64`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *AzureOpenIDConnectIdentityProviderDTO) GetYOk() (*float64, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *AzureOpenIDConnectIdentityProviderDTO) SetY(v float64)`

SetY sets Y field to given value.

### HasY

`func (o *AzureOpenIDConnectIdentityProviderDTO) HasY() bool`

HasY returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


