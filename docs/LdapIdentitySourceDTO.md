# LdapIdentitySourceDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CredentialQueryString** | Pointer to **string** |  | [optional] 
**CustomClass** | Pointer to [**CustomClassDTO**](CustomClassDTO.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**IncludeOperationalAttributes** | Pointer to **bool** |  | [optional] 
**InitialContextFactory** | Pointer to **string** |  | [optional] 
**LdapSearchScope** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PrincipalUidAttributeID** | Pointer to **string** |  | [optional] 
**ProviderUrl** | Pointer to **string** |  | [optional] 
**Referrals** | Pointer to **string** |  | [optional] 
**RoleAttributeID** | Pointer to **string** |  | [optional] 
**RoleMatchingMode** | Pointer to **string** |  | [optional] 
**RolesCtxDN** | Pointer to **string** |  | [optional] 
**SecurityAuthentication** | Pointer to **string** |  | [optional] 
**SecurityCredential** | Pointer to **string** |  | [optional] 
**SecurityPrincipal** | Pointer to **string** |  | [optional] 
**UidAttributeID** | Pointer to **string** |  | [optional] 
**UpdatePasswordEnabled** | Pointer to **bool** |  | [optional] 
**UpdateableCredentialAttribute** | Pointer to **string** |  | [optional] 
**UserPropertiesQueryString** | Pointer to **string** |  | [optional] 
**UsersCtxDN** | Pointer to **string** |  | [optional] 
**X** | Pointer to **float64** |  | [optional] 
**Y** | Pointer to **float64** |  | [optional] 

## Methods

### NewLdapIdentitySourceDTO

`func NewLdapIdentitySourceDTO() *LdapIdentitySourceDTO`

NewLdapIdentitySourceDTO instantiates a new LdapIdentitySourceDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapIdentitySourceDTOWithDefaults

`func NewLdapIdentitySourceDTOWithDefaults() *LdapIdentitySourceDTO`

NewLdapIdentitySourceDTOWithDefaults instantiates a new LdapIdentitySourceDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentialQueryString

`func (o *LdapIdentitySourceDTO) GetCredentialQueryString() string`

GetCredentialQueryString returns the CredentialQueryString field if non-nil, zero value otherwise.

### GetCredentialQueryStringOk

`func (o *LdapIdentitySourceDTO) GetCredentialQueryStringOk() (*string, bool)`

GetCredentialQueryStringOk returns a tuple with the CredentialQueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialQueryString

`func (o *LdapIdentitySourceDTO) SetCredentialQueryString(v string)`

SetCredentialQueryString sets CredentialQueryString field to given value.

### HasCredentialQueryString

`func (o *LdapIdentitySourceDTO) HasCredentialQueryString() bool`

HasCredentialQueryString returns a boolean if a field has been set.

### GetCustomClass

`func (o *LdapIdentitySourceDTO) GetCustomClass() CustomClassDTO`

GetCustomClass returns the CustomClass field if non-nil, zero value otherwise.

### GetCustomClassOk

`func (o *LdapIdentitySourceDTO) GetCustomClassOk() (*CustomClassDTO, bool)`

GetCustomClassOk returns a tuple with the CustomClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomClass

`func (o *LdapIdentitySourceDTO) SetCustomClass(v CustomClassDTO)`

SetCustomClass sets CustomClass field to given value.

### HasCustomClass

`func (o *LdapIdentitySourceDTO) HasCustomClass() bool`

HasCustomClass returns a boolean if a field has been set.

### GetDescription

`func (o *LdapIdentitySourceDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LdapIdentitySourceDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LdapIdentitySourceDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LdapIdentitySourceDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetElementId

`func (o *LdapIdentitySourceDTO) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *LdapIdentitySourceDTO) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *LdapIdentitySourceDTO) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *LdapIdentitySourceDTO) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetId

`func (o *LdapIdentitySourceDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LdapIdentitySourceDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LdapIdentitySourceDTO) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *LdapIdentitySourceDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIncludeOperationalAttributes

`func (o *LdapIdentitySourceDTO) GetIncludeOperationalAttributes() bool`

GetIncludeOperationalAttributes returns the IncludeOperationalAttributes field if non-nil, zero value otherwise.

### GetIncludeOperationalAttributesOk

`func (o *LdapIdentitySourceDTO) GetIncludeOperationalAttributesOk() (*bool, bool)`

GetIncludeOperationalAttributesOk returns a tuple with the IncludeOperationalAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeOperationalAttributes

`func (o *LdapIdentitySourceDTO) SetIncludeOperationalAttributes(v bool)`

SetIncludeOperationalAttributes sets IncludeOperationalAttributes field to given value.

### HasIncludeOperationalAttributes

`func (o *LdapIdentitySourceDTO) HasIncludeOperationalAttributes() bool`

HasIncludeOperationalAttributes returns a boolean if a field has been set.

### GetInitialContextFactory

`func (o *LdapIdentitySourceDTO) GetInitialContextFactory() string`

GetInitialContextFactory returns the InitialContextFactory field if non-nil, zero value otherwise.

### GetInitialContextFactoryOk

`func (o *LdapIdentitySourceDTO) GetInitialContextFactoryOk() (*string, bool)`

GetInitialContextFactoryOk returns a tuple with the InitialContextFactory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialContextFactory

`func (o *LdapIdentitySourceDTO) SetInitialContextFactory(v string)`

SetInitialContextFactory sets InitialContextFactory field to given value.

### HasInitialContextFactory

`func (o *LdapIdentitySourceDTO) HasInitialContextFactory() bool`

HasInitialContextFactory returns a boolean if a field has been set.

### GetLdapSearchScope

`func (o *LdapIdentitySourceDTO) GetLdapSearchScope() string`

GetLdapSearchScope returns the LdapSearchScope field if non-nil, zero value otherwise.

### GetLdapSearchScopeOk

`func (o *LdapIdentitySourceDTO) GetLdapSearchScopeOk() (*string, bool)`

GetLdapSearchScopeOk returns a tuple with the LdapSearchScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapSearchScope

`func (o *LdapIdentitySourceDTO) SetLdapSearchScope(v string)`

SetLdapSearchScope sets LdapSearchScope field to given value.

### HasLdapSearchScope

`func (o *LdapIdentitySourceDTO) HasLdapSearchScope() bool`

HasLdapSearchScope returns a boolean if a field has been set.

### GetName

`func (o *LdapIdentitySourceDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LdapIdentitySourceDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LdapIdentitySourceDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LdapIdentitySourceDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrincipalUidAttributeID

`func (o *LdapIdentitySourceDTO) GetPrincipalUidAttributeID() string`

GetPrincipalUidAttributeID returns the PrincipalUidAttributeID field if non-nil, zero value otherwise.

### GetPrincipalUidAttributeIDOk

`func (o *LdapIdentitySourceDTO) GetPrincipalUidAttributeIDOk() (*string, bool)`

GetPrincipalUidAttributeIDOk returns a tuple with the PrincipalUidAttributeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalUidAttributeID

`func (o *LdapIdentitySourceDTO) SetPrincipalUidAttributeID(v string)`

SetPrincipalUidAttributeID sets PrincipalUidAttributeID field to given value.

### HasPrincipalUidAttributeID

`func (o *LdapIdentitySourceDTO) HasPrincipalUidAttributeID() bool`

HasPrincipalUidAttributeID returns a boolean if a field has been set.

### GetProviderUrl

`func (o *LdapIdentitySourceDTO) GetProviderUrl() string`

GetProviderUrl returns the ProviderUrl field if non-nil, zero value otherwise.

### GetProviderUrlOk

`func (o *LdapIdentitySourceDTO) GetProviderUrlOk() (*string, bool)`

GetProviderUrlOk returns a tuple with the ProviderUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderUrl

`func (o *LdapIdentitySourceDTO) SetProviderUrl(v string)`

SetProviderUrl sets ProviderUrl field to given value.

### HasProviderUrl

`func (o *LdapIdentitySourceDTO) HasProviderUrl() bool`

HasProviderUrl returns a boolean if a field has been set.

### GetReferrals

`func (o *LdapIdentitySourceDTO) GetReferrals() string`

GetReferrals returns the Referrals field if non-nil, zero value otherwise.

### GetReferralsOk

`func (o *LdapIdentitySourceDTO) GetReferralsOk() (*string, bool)`

GetReferralsOk returns a tuple with the Referrals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferrals

`func (o *LdapIdentitySourceDTO) SetReferrals(v string)`

SetReferrals sets Referrals field to given value.

### HasReferrals

`func (o *LdapIdentitySourceDTO) HasReferrals() bool`

HasReferrals returns a boolean if a field has been set.

### GetRoleAttributeID

`func (o *LdapIdentitySourceDTO) GetRoleAttributeID() string`

GetRoleAttributeID returns the RoleAttributeID field if non-nil, zero value otherwise.

### GetRoleAttributeIDOk

`func (o *LdapIdentitySourceDTO) GetRoleAttributeIDOk() (*string, bool)`

GetRoleAttributeIDOk returns a tuple with the RoleAttributeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleAttributeID

`func (o *LdapIdentitySourceDTO) SetRoleAttributeID(v string)`

SetRoleAttributeID sets RoleAttributeID field to given value.

### HasRoleAttributeID

`func (o *LdapIdentitySourceDTO) HasRoleAttributeID() bool`

HasRoleAttributeID returns a boolean if a field has been set.

### GetRoleMatchingMode

`func (o *LdapIdentitySourceDTO) GetRoleMatchingMode() string`

GetRoleMatchingMode returns the RoleMatchingMode field if non-nil, zero value otherwise.

### GetRoleMatchingModeOk

`func (o *LdapIdentitySourceDTO) GetRoleMatchingModeOk() (*string, bool)`

GetRoleMatchingModeOk returns a tuple with the RoleMatchingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleMatchingMode

`func (o *LdapIdentitySourceDTO) SetRoleMatchingMode(v string)`

SetRoleMatchingMode sets RoleMatchingMode field to given value.

### HasRoleMatchingMode

`func (o *LdapIdentitySourceDTO) HasRoleMatchingMode() bool`

HasRoleMatchingMode returns a boolean if a field has been set.

### GetRolesCtxDN

`func (o *LdapIdentitySourceDTO) GetRolesCtxDN() string`

GetRolesCtxDN returns the RolesCtxDN field if non-nil, zero value otherwise.

### GetRolesCtxDNOk

`func (o *LdapIdentitySourceDTO) GetRolesCtxDNOk() (*string, bool)`

GetRolesCtxDNOk returns a tuple with the RolesCtxDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolesCtxDN

`func (o *LdapIdentitySourceDTO) SetRolesCtxDN(v string)`

SetRolesCtxDN sets RolesCtxDN field to given value.

### HasRolesCtxDN

`func (o *LdapIdentitySourceDTO) HasRolesCtxDN() bool`

HasRolesCtxDN returns a boolean if a field has been set.

### GetSecurityAuthentication

`func (o *LdapIdentitySourceDTO) GetSecurityAuthentication() string`

GetSecurityAuthentication returns the SecurityAuthentication field if non-nil, zero value otherwise.

### GetSecurityAuthenticationOk

`func (o *LdapIdentitySourceDTO) GetSecurityAuthenticationOk() (*string, bool)`

GetSecurityAuthenticationOk returns a tuple with the SecurityAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityAuthentication

`func (o *LdapIdentitySourceDTO) SetSecurityAuthentication(v string)`

SetSecurityAuthentication sets SecurityAuthentication field to given value.

### HasSecurityAuthentication

`func (o *LdapIdentitySourceDTO) HasSecurityAuthentication() bool`

HasSecurityAuthentication returns a boolean if a field has been set.

### GetSecurityCredential

`func (o *LdapIdentitySourceDTO) GetSecurityCredential() string`

GetSecurityCredential returns the SecurityCredential field if non-nil, zero value otherwise.

### GetSecurityCredentialOk

`func (o *LdapIdentitySourceDTO) GetSecurityCredentialOk() (*string, bool)`

GetSecurityCredentialOk returns a tuple with the SecurityCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityCredential

`func (o *LdapIdentitySourceDTO) SetSecurityCredential(v string)`

SetSecurityCredential sets SecurityCredential field to given value.

### HasSecurityCredential

`func (o *LdapIdentitySourceDTO) HasSecurityCredential() bool`

HasSecurityCredential returns a boolean if a field has been set.

### GetSecurityPrincipal

`func (o *LdapIdentitySourceDTO) GetSecurityPrincipal() string`

GetSecurityPrincipal returns the SecurityPrincipal field if non-nil, zero value otherwise.

### GetSecurityPrincipalOk

`func (o *LdapIdentitySourceDTO) GetSecurityPrincipalOk() (*string, bool)`

GetSecurityPrincipalOk returns a tuple with the SecurityPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityPrincipal

`func (o *LdapIdentitySourceDTO) SetSecurityPrincipal(v string)`

SetSecurityPrincipal sets SecurityPrincipal field to given value.

### HasSecurityPrincipal

`func (o *LdapIdentitySourceDTO) HasSecurityPrincipal() bool`

HasSecurityPrincipal returns a boolean if a field has been set.

### GetUidAttributeID

`func (o *LdapIdentitySourceDTO) GetUidAttributeID() string`

GetUidAttributeID returns the UidAttributeID field if non-nil, zero value otherwise.

### GetUidAttributeIDOk

`func (o *LdapIdentitySourceDTO) GetUidAttributeIDOk() (*string, bool)`

GetUidAttributeIDOk returns a tuple with the UidAttributeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidAttributeID

`func (o *LdapIdentitySourceDTO) SetUidAttributeID(v string)`

SetUidAttributeID sets UidAttributeID field to given value.

### HasUidAttributeID

`func (o *LdapIdentitySourceDTO) HasUidAttributeID() bool`

HasUidAttributeID returns a boolean if a field has been set.

### GetUpdatePasswordEnabled

`func (o *LdapIdentitySourceDTO) GetUpdatePasswordEnabled() bool`

GetUpdatePasswordEnabled returns the UpdatePasswordEnabled field if non-nil, zero value otherwise.

### GetUpdatePasswordEnabledOk

`func (o *LdapIdentitySourceDTO) GetUpdatePasswordEnabledOk() (*bool, bool)`

GetUpdatePasswordEnabledOk returns a tuple with the UpdatePasswordEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatePasswordEnabled

`func (o *LdapIdentitySourceDTO) SetUpdatePasswordEnabled(v bool)`

SetUpdatePasswordEnabled sets UpdatePasswordEnabled field to given value.

### HasUpdatePasswordEnabled

`func (o *LdapIdentitySourceDTO) HasUpdatePasswordEnabled() bool`

HasUpdatePasswordEnabled returns a boolean if a field has been set.

### GetUpdateableCredentialAttribute

`func (o *LdapIdentitySourceDTO) GetUpdateableCredentialAttribute() string`

GetUpdateableCredentialAttribute returns the UpdateableCredentialAttribute field if non-nil, zero value otherwise.

### GetUpdateableCredentialAttributeOk

`func (o *LdapIdentitySourceDTO) GetUpdateableCredentialAttributeOk() (*string, bool)`

GetUpdateableCredentialAttributeOk returns a tuple with the UpdateableCredentialAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateableCredentialAttribute

`func (o *LdapIdentitySourceDTO) SetUpdateableCredentialAttribute(v string)`

SetUpdateableCredentialAttribute sets UpdateableCredentialAttribute field to given value.

### HasUpdateableCredentialAttribute

`func (o *LdapIdentitySourceDTO) HasUpdateableCredentialAttribute() bool`

HasUpdateableCredentialAttribute returns a boolean if a field has been set.

### GetUserPropertiesQueryString

`func (o *LdapIdentitySourceDTO) GetUserPropertiesQueryString() string`

GetUserPropertiesQueryString returns the UserPropertiesQueryString field if non-nil, zero value otherwise.

### GetUserPropertiesQueryStringOk

`func (o *LdapIdentitySourceDTO) GetUserPropertiesQueryStringOk() (*string, bool)`

GetUserPropertiesQueryStringOk returns a tuple with the UserPropertiesQueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPropertiesQueryString

`func (o *LdapIdentitySourceDTO) SetUserPropertiesQueryString(v string)`

SetUserPropertiesQueryString sets UserPropertiesQueryString field to given value.

### HasUserPropertiesQueryString

`func (o *LdapIdentitySourceDTO) HasUserPropertiesQueryString() bool`

HasUserPropertiesQueryString returns a boolean if a field has been set.

### GetUsersCtxDN

`func (o *LdapIdentitySourceDTO) GetUsersCtxDN() string`

GetUsersCtxDN returns the UsersCtxDN field if non-nil, zero value otherwise.

### GetUsersCtxDNOk

`func (o *LdapIdentitySourceDTO) GetUsersCtxDNOk() (*string, bool)`

GetUsersCtxDNOk returns a tuple with the UsersCtxDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersCtxDN

`func (o *LdapIdentitySourceDTO) SetUsersCtxDN(v string)`

SetUsersCtxDN sets UsersCtxDN field to given value.

### HasUsersCtxDN

`func (o *LdapIdentitySourceDTO) HasUsersCtxDN() bool`

HasUsersCtxDN returns a boolean if a field has been set.

### GetX

`func (o *LdapIdentitySourceDTO) GetX() float64`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *LdapIdentitySourceDTO) GetXOk() (*float64, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *LdapIdentitySourceDTO) SetX(v float64)`

SetX sets X field to given value.

### HasX

`func (o *LdapIdentitySourceDTO) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *LdapIdentitySourceDTO) GetY() float64`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *LdapIdentitySourceDTO) GetYOk() (*float64, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *LdapIdentitySourceDTO) SetY(v float64)`

SetY sets Y field to given value.

### HasY

`func (o *LdapIdentitySourceDTO) HasY() bool`

HasY returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


