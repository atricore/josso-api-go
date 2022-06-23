# DirectoryAuthenticationServiceDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomClass** | Pointer to [**CustomClassDTO**](CustomClassDTO.md) |  | [optional] 
**DelegatedAuthentications** | Pointer to [**[]DelegatedAuthenticationDTO**](DelegatedAuthenticationDTO.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**IncludeOperationalAttributes** | Pointer to **bool** |  | [optional] 
**InitialContextFactory** | Pointer to **string** |  | [optional] 
**LdapSearchScope** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PasswordPolicy** | Pointer to **string** |  | [optional] 
**PerformDnSearch** | Pointer to **bool** |  | [optional] 
**PrincipalUidAttributeID** | Pointer to **string** |  | [optional] 
**ProviderUrl** | Pointer to **string** |  | [optional] 
**Referrals** | Pointer to **string** |  | [optional] 
**SecurityAuthentication** | Pointer to **string** |  | [optional] 
**SecurityCredential** | Pointer to **string** |  | [optional] 
**SecurityPrincipal** | Pointer to **string** |  | [optional] 
**SimpleAuthnSaml2AuthnCtxClass** | Pointer to **string** |  | [optional] 
**UsersCtxDN** | Pointer to **string** |  | [optional] 
**X** | Pointer to **float64** |  | [optional] 
**Y** | Pointer to **float64** |  | [optional] 

## Methods

### NewDirectoryAuthenticationServiceDTO

`func NewDirectoryAuthenticationServiceDTO() *DirectoryAuthenticationServiceDTO`

NewDirectoryAuthenticationServiceDTO instantiates a new DirectoryAuthenticationServiceDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectoryAuthenticationServiceDTOWithDefaults

`func NewDirectoryAuthenticationServiceDTOWithDefaults() *DirectoryAuthenticationServiceDTO`

NewDirectoryAuthenticationServiceDTOWithDefaults instantiates a new DirectoryAuthenticationServiceDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomClass

`func (o *DirectoryAuthenticationServiceDTO) GetCustomClass() CustomClassDTO`

GetCustomClass returns the CustomClass field if non-nil, zero value otherwise.

### GetCustomClassOk

`func (o *DirectoryAuthenticationServiceDTO) GetCustomClassOk() (*CustomClassDTO, bool)`

GetCustomClassOk returns a tuple with the CustomClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomClass

`func (o *DirectoryAuthenticationServiceDTO) SetCustomClass(v CustomClassDTO)`

SetCustomClass sets CustomClass field to given value.

### HasCustomClass

`func (o *DirectoryAuthenticationServiceDTO) HasCustomClass() bool`

HasCustomClass returns a boolean if a field has been set.

### GetDelegatedAuthentications

`func (o *DirectoryAuthenticationServiceDTO) GetDelegatedAuthentications() []DelegatedAuthenticationDTO`

GetDelegatedAuthentications returns the DelegatedAuthentications field if non-nil, zero value otherwise.

### GetDelegatedAuthenticationsOk

`func (o *DirectoryAuthenticationServiceDTO) GetDelegatedAuthenticationsOk() (*[]DelegatedAuthenticationDTO, bool)`

GetDelegatedAuthenticationsOk returns a tuple with the DelegatedAuthentications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatedAuthentications

`func (o *DirectoryAuthenticationServiceDTO) SetDelegatedAuthentications(v []DelegatedAuthenticationDTO)`

SetDelegatedAuthentications sets DelegatedAuthentications field to given value.

### HasDelegatedAuthentications

`func (o *DirectoryAuthenticationServiceDTO) HasDelegatedAuthentications() bool`

HasDelegatedAuthentications returns a boolean if a field has been set.

### GetDescription

`func (o *DirectoryAuthenticationServiceDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DirectoryAuthenticationServiceDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DirectoryAuthenticationServiceDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DirectoryAuthenticationServiceDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *DirectoryAuthenticationServiceDTO) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DirectoryAuthenticationServiceDTO) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DirectoryAuthenticationServiceDTO) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *DirectoryAuthenticationServiceDTO) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetElementId

`func (o *DirectoryAuthenticationServiceDTO) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *DirectoryAuthenticationServiceDTO) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *DirectoryAuthenticationServiceDTO) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *DirectoryAuthenticationServiceDTO) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetId

`func (o *DirectoryAuthenticationServiceDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DirectoryAuthenticationServiceDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DirectoryAuthenticationServiceDTO) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DirectoryAuthenticationServiceDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIncludeOperationalAttributes

`func (o *DirectoryAuthenticationServiceDTO) GetIncludeOperationalAttributes() bool`

GetIncludeOperationalAttributes returns the IncludeOperationalAttributes field if non-nil, zero value otherwise.

### GetIncludeOperationalAttributesOk

`func (o *DirectoryAuthenticationServiceDTO) GetIncludeOperationalAttributesOk() (*bool, bool)`

GetIncludeOperationalAttributesOk returns a tuple with the IncludeOperationalAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeOperationalAttributes

`func (o *DirectoryAuthenticationServiceDTO) SetIncludeOperationalAttributes(v bool)`

SetIncludeOperationalAttributes sets IncludeOperationalAttributes field to given value.

### HasIncludeOperationalAttributes

`func (o *DirectoryAuthenticationServiceDTO) HasIncludeOperationalAttributes() bool`

HasIncludeOperationalAttributes returns a boolean if a field has been set.

### GetInitialContextFactory

`func (o *DirectoryAuthenticationServiceDTO) GetInitialContextFactory() string`

GetInitialContextFactory returns the InitialContextFactory field if non-nil, zero value otherwise.

### GetInitialContextFactoryOk

`func (o *DirectoryAuthenticationServiceDTO) GetInitialContextFactoryOk() (*string, bool)`

GetInitialContextFactoryOk returns a tuple with the InitialContextFactory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialContextFactory

`func (o *DirectoryAuthenticationServiceDTO) SetInitialContextFactory(v string)`

SetInitialContextFactory sets InitialContextFactory field to given value.

### HasInitialContextFactory

`func (o *DirectoryAuthenticationServiceDTO) HasInitialContextFactory() bool`

HasInitialContextFactory returns a boolean if a field has been set.

### GetLdapSearchScope

`func (o *DirectoryAuthenticationServiceDTO) GetLdapSearchScope() string`

GetLdapSearchScope returns the LdapSearchScope field if non-nil, zero value otherwise.

### GetLdapSearchScopeOk

`func (o *DirectoryAuthenticationServiceDTO) GetLdapSearchScopeOk() (*string, bool)`

GetLdapSearchScopeOk returns a tuple with the LdapSearchScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapSearchScope

`func (o *DirectoryAuthenticationServiceDTO) SetLdapSearchScope(v string)`

SetLdapSearchScope sets LdapSearchScope field to given value.

### HasLdapSearchScope

`func (o *DirectoryAuthenticationServiceDTO) HasLdapSearchScope() bool`

HasLdapSearchScope returns a boolean if a field has been set.

### GetName

`func (o *DirectoryAuthenticationServiceDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DirectoryAuthenticationServiceDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DirectoryAuthenticationServiceDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DirectoryAuthenticationServiceDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPasswordPolicy

`func (o *DirectoryAuthenticationServiceDTO) GetPasswordPolicy() string`

GetPasswordPolicy returns the PasswordPolicy field if non-nil, zero value otherwise.

### GetPasswordPolicyOk

`func (o *DirectoryAuthenticationServiceDTO) GetPasswordPolicyOk() (*string, bool)`

GetPasswordPolicyOk returns a tuple with the PasswordPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordPolicy

`func (o *DirectoryAuthenticationServiceDTO) SetPasswordPolicy(v string)`

SetPasswordPolicy sets PasswordPolicy field to given value.

### HasPasswordPolicy

`func (o *DirectoryAuthenticationServiceDTO) HasPasswordPolicy() bool`

HasPasswordPolicy returns a boolean if a field has been set.

### GetPerformDnSearch

`func (o *DirectoryAuthenticationServiceDTO) GetPerformDnSearch() bool`

GetPerformDnSearch returns the PerformDnSearch field if non-nil, zero value otherwise.

### GetPerformDnSearchOk

`func (o *DirectoryAuthenticationServiceDTO) GetPerformDnSearchOk() (*bool, bool)`

GetPerformDnSearchOk returns a tuple with the PerformDnSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformDnSearch

`func (o *DirectoryAuthenticationServiceDTO) SetPerformDnSearch(v bool)`

SetPerformDnSearch sets PerformDnSearch field to given value.

### HasPerformDnSearch

`func (o *DirectoryAuthenticationServiceDTO) HasPerformDnSearch() bool`

HasPerformDnSearch returns a boolean if a field has been set.

### GetPrincipalUidAttributeID

`func (o *DirectoryAuthenticationServiceDTO) GetPrincipalUidAttributeID() string`

GetPrincipalUidAttributeID returns the PrincipalUidAttributeID field if non-nil, zero value otherwise.

### GetPrincipalUidAttributeIDOk

`func (o *DirectoryAuthenticationServiceDTO) GetPrincipalUidAttributeIDOk() (*string, bool)`

GetPrincipalUidAttributeIDOk returns a tuple with the PrincipalUidAttributeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalUidAttributeID

`func (o *DirectoryAuthenticationServiceDTO) SetPrincipalUidAttributeID(v string)`

SetPrincipalUidAttributeID sets PrincipalUidAttributeID field to given value.

### HasPrincipalUidAttributeID

`func (o *DirectoryAuthenticationServiceDTO) HasPrincipalUidAttributeID() bool`

HasPrincipalUidAttributeID returns a boolean if a field has been set.

### GetProviderUrl

`func (o *DirectoryAuthenticationServiceDTO) GetProviderUrl() string`

GetProviderUrl returns the ProviderUrl field if non-nil, zero value otherwise.

### GetProviderUrlOk

`func (o *DirectoryAuthenticationServiceDTO) GetProviderUrlOk() (*string, bool)`

GetProviderUrlOk returns a tuple with the ProviderUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderUrl

`func (o *DirectoryAuthenticationServiceDTO) SetProviderUrl(v string)`

SetProviderUrl sets ProviderUrl field to given value.

### HasProviderUrl

`func (o *DirectoryAuthenticationServiceDTO) HasProviderUrl() bool`

HasProviderUrl returns a boolean if a field has been set.

### GetReferrals

`func (o *DirectoryAuthenticationServiceDTO) GetReferrals() string`

GetReferrals returns the Referrals field if non-nil, zero value otherwise.

### GetReferralsOk

`func (o *DirectoryAuthenticationServiceDTO) GetReferralsOk() (*string, bool)`

GetReferralsOk returns a tuple with the Referrals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferrals

`func (o *DirectoryAuthenticationServiceDTO) SetReferrals(v string)`

SetReferrals sets Referrals field to given value.

### HasReferrals

`func (o *DirectoryAuthenticationServiceDTO) HasReferrals() bool`

HasReferrals returns a boolean if a field has been set.

### GetSecurityAuthentication

`func (o *DirectoryAuthenticationServiceDTO) GetSecurityAuthentication() string`

GetSecurityAuthentication returns the SecurityAuthentication field if non-nil, zero value otherwise.

### GetSecurityAuthenticationOk

`func (o *DirectoryAuthenticationServiceDTO) GetSecurityAuthenticationOk() (*string, bool)`

GetSecurityAuthenticationOk returns a tuple with the SecurityAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityAuthentication

`func (o *DirectoryAuthenticationServiceDTO) SetSecurityAuthentication(v string)`

SetSecurityAuthentication sets SecurityAuthentication field to given value.

### HasSecurityAuthentication

`func (o *DirectoryAuthenticationServiceDTO) HasSecurityAuthentication() bool`

HasSecurityAuthentication returns a boolean if a field has been set.

### GetSecurityCredential

`func (o *DirectoryAuthenticationServiceDTO) GetSecurityCredential() string`

GetSecurityCredential returns the SecurityCredential field if non-nil, zero value otherwise.

### GetSecurityCredentialOk

`func (o *DirectoryAuthenticationServiceDTO) GetSecurityCredentialOk() (*string, bool)`

GetSecurityCredentialOk returns a tuple with the SecurityCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityCredential

`func (o *DirectoryAuthenticationServiceDTO) SetSecurityCredential(v string)`

SetSecurityCredential sets SecurityCredential field to given value.

### HasSecurityCredential

`func (o *DirectoryAuthenticationServiceDTO) HasSecurityCredential() bool`

HasSecurityCredential returns a boolean if a field has been set.

### GetSecurityPrincipal

`func (o *DirectoryAuthenticationServiceDTO) GetSecurityPrincipal() string`

GetSecurityPrincipal returns the SecurityPrincipal field if non-nil, zero value otherwise.

### GetSecurityPrincipalOk

`func (o *DirectoryAuthenticationServiceDTO) GetSecurityPrincipalOk() (*string, bool)`

GetSecurityPrincipalOk returns a tuple with the SecurityPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityPrincipal

`func (o *DirectoryAuthenticationServiceDTO) SetSecurityPrincipal(v string)`

SetSecurityPrincipal sets SecurityPrincipal field to given value.

### HasSecurityPrincipal

`func (o *DirectoryAuthenticationServiceDTO) HasSecurityPrincipal() bool`

HasSecurityPrincipal returns a boolean if a field has been set.

### GetSimpleAuthnSaml2AuthnCtxClass

`func (o *DirectoryAuthenticationServiceDTO) GetSimpleAuthnSaml2AuthnCtxClass() string`

GetSimpleAuthnSaml2AuthnCtxClass returns the SimpleAuthnSaml2AuthnCtxClass field if non-nil, zero value otherwise.

### GetSimpleAuthnSaml2AuthnCtxClassOk

`func (o *DirectoryAuthenticationServiceDTO) GetSimpleAuthnSaml2AuthnCtxClassOk() (*string, bool)`

GetSimpleAuthnSaml2AuthnCtxClassOk returns a tuple with the SimpleAuthnSaml2AuthnCtxClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimpleAuthnSaml2AuthnCtxClass

`func (o *DirectoryAuthenticationServiceDTO) SetSimpleAuthnSaml2AuthnCtxClass(v string)`

SetSimpleAuthnSaml2AuthnCtxClass sets SimpleAuthnSaml2AuthnCtxClass field to given value.

### HasSimpleAuthnSaml2AuthnCtxClass

`func (o *DirectoryAuthenticationServiceDTO) HasSimpleAuthnSaml2AuthnCtxClass() bool`

HasSimpleAuthnSaml2AuthnCtxClass returns a boolean if a field has been set.

### GetUsersCtxDN

`func (o *DirectoryAuthenticationServiceDTO) GetUsersCtxDN() string`

GetUsersCtxDN returns the UsersCtxDN field if non-nil, zero value otherwise.

### GetUsersCtxDNOk

`func (o *DirectoryAuthenticationServiceDTO) GetUsersCtxDNOk() (*string, bool)`

GetUsersCtxDNOk returns a tuple with the UsersCtxDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersCtxDN

`func (o *DirectoryAuthenticationServiceDTO) SetUsersCtxDN(v string)`

SetUsersCtxDN sets UsersCtxDN field to given value.

### HasUsersCtxDN

`func (o *DirectoryAuthenticationServiceDTO) HasUsersCtxDN() bool`

HasUsersCtxDN returns a boolean if a field has been set.

### GetX

`func (o *DirectoryAuthenticationServiceDTO) GetX() float64`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *DirectoryAuthenticationServiceDTO) GetXOk() (*float64, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *DirectoryAuthenticationServiceDTO) SetX(v float64)`

SetX sets X field to given value.

### HasX

`func (o *DirectoryAuthenticationServiceDTO) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *DirectoryAuthenticationServiceDTO) GetY() float64`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *DirectoryAuthenticationServiceDTO) GetYOk() (*float64, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *DirectoryAuthenticationServiceDTO) SetY(v float64)`

SetY sets Y field to given value.

### HasY

`func (o *DirectoryAuthenticationServiceDTO) HasY() bool`

HasY returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


