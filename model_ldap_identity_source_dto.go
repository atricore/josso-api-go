/*
Atricore Console :: Remote : API

# Atricore Console API

API version: 1.5.0-SNAPSHOT
Contact: sgonzalez@atricore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package jossoappi

import (
	"encoding/json"
)

// LdapIdentitySourceDTO struct for LdapIdentitySourceDTO
type LdapIdentitySourceDTO struct {
	CredentialQueryString *string `json:"credentialQueryString,omitempty"`
	CustomClass *CustomClassDTO `json:"customClass,omitempty"`
	Description *string `json:"description,omitempty"`
	ElementId *string `json:"elementId,omitempty"`
	Id *int64 `json:"id,omitempty"`
	IncludeOperationalAttributes *bool `json:"includeOperationalAttributes,omitempty"`
	InitialContextFactory *string `json:"initialContextFactory,omitempty"`
	LdapSearchScope *string `json:"ldapSearchScope,omitempty"`
	Name *string `json:"name,omitempty"`
	PrincipalUidAttributeID *string `json:"principalUidAttributeID,omitempty"`
	ProviderUrl *string `json:"providerUrl,omitempty"`
	Referrals *string `json:"referrals,omitempty"`
	RoleAttributeID *string `json:"roleAttributeID,omitempty"`
	RoleMatchingMode *string `json:"roleMatchingMode,omitempty"`
	RolesCtxDN *string `json:"rolesCtxDN,omitempty"`
	SecurityAuthentication *string `json:"securityAuthentication,omitempty"`
	SecurityCredential *string `json:"securityCredential,omitempty"`
	SecurityPrincipal *string `json:"securityPrincipal,omitempty"`
	UidAttributeID *string `json:"uidAttributeID,omitempty"`
	UpdatePasswordEnabled *bool `json:"updatePasswordEnabled,omitempty"`
	UpdateableCredentialAttribute *string `json:"updateableCredentialAttribute,omitempty"`
	UserPropertiesQueryString *string `json:"userPropertiesQueryString,omitempty"`
	UsersCtxDN *string `json:"usersCtxDN,omitempty"`
	X *float64 `json:"x,omitempty"`
	Y *float64 `json:"y,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LdapIdentitySourceDTO LdapIdentitySourceDTO

// NewLdapIdentitySourceDTO instantiates a new LdapIdentitySourceDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLdapIdentitySourceDTO() *LdapIdentitySourceDTO {
	this := LdapIdentitySourceDTO{}
	return &this
}

// NewLdapIdentitySourceDTOWithDefaults instantiates a new LdapIdentitySourceDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLdapIdentitySourceDTOWithDefaults() *LdapIdentitySourceDTO {
	this := LdapIdentitySourceDTO{}
	return &this
}

// GetCredentialQueryString returns the CredentialQueryString field value if set, zero value otherwise.
func (o *LdapIdentitySourceDTO) GetCredentialQueryString() string {
	if o == nil || isNil(o.CredentialQueryString) {
		var ret string
		return ret
	}
	return *o.CredentialQueryString
}

// GetCredentialQueryStringOk returns a tuple with the CredentialQueryString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapIdentitySourceDTO) GetCredentialQueryStringOk() (*string, bool) {
	if o == nil || isNil(o.CredentialQueryString) {
    return nil, false
	}
	return o.CredentialQueryString, true
}

// HasCredentialQueryString returns a boolean if a field has been set.
func (o *LdapIdentitySourceDTO) HasCredentialQueryString() bool {
	if o != nil && !isNil(o.CredentialQueryString) {
		return true
	}

	return false
}

// SetCredentialQueryString gets a reference to the given string and assigns it to the CredentialQueryString field.
func (o *LdapIdentitySourceDTO) SetCredentialQueryString(v string) {
	o.CredentialQueryString = &v
}

// GetCustomClass returns the CustomClass field value if set, zero value otherwise.
func (o *LdapIdentitySourceDTO) GetCustomClass() CustomClassDTO {
	if o == nil || isNil(o.CustomClass) {
		var ret CustomClassDTO
		return ret
	}
	return *o.CustomClass
}

// GetCustomClassOk returns a tuple with the CustomClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapIdentitySourceDTO) GetCustomClassOk() (*CustomClassDTO, bool) {
	if o == nil || isNil(o.CustomClass) {
    return nil, false
	}
	return o.CustomClass, true
}

// HasCustomClass returns a boolean if a field has been set.
func (o *LdapIdentitySourceDTO) HasCustomClass() bool {
	if o != nil && !isNil(o.CustomClass) {
		return true
	}

	return false
}

// SetCustomClass gets a reference to the given CustomClassDTO and assigns it to the CustomClass field.
func (o *LdapIdentitySourceDTO) SetCustomClass(v CustomClassDTO) {
	o.CustomClass = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *LdapIdentitySourceDTO) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapIdentitySourceDTO) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *LdapIdentitySourceDTO) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *LdapIdentitySourceDTO) SetDescription(v string) {
	o.Description = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *LdapIdentitySourceDTO) GetElementId() string {
	if o == nil || isNil(o.ElementId) {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapIdentitySourceDTO) GetElementIdOk() (*string, bool) {
	if o == nil || isNil(o.ElementId) {
    return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *LdapIdentitySourceDTO) HasElementId() bool {
	if o != nil && !isNil(o.ElementId) {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *LdapIdentitySourceDTO) SetElementId(v string) {
	o.ElementId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LdapIdentitySourceDTO) GetId() int64 {
	if o == nil || isNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapIdentitySourceDTO) GetIdOk() (*int64, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LdapIdentitySourceDTO) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *LdapIdentitySourceDTO) SetId(v int64) {
	o.Id = &v
}

// GetIncludeOperationalAttributes returns the IncludeOperationalAttributes field value if set, zero value otherwise.
func (o *LdapIdentitySourceDTO) GetIncludeOperationalAttributes() bool {
	if o == nil || isNil(o.IncludeOperationalAttributes) {
		var ret bool
		return ret
	}
	return *o.IncludeOperationalAttributes
}

// GetIncludeOperationalAttributesOk returns a tuple with the IncludeOperationalAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapIdentitySourceDTO) GetIncludeOperationalAttributesOk() (*bool, bool) {
	if o == nil || isNil(o.IncludeOperationalAttributes) {
    return nil, false
	}
	return o.IncludeOperationalAttributes, true
}

// HasIncludeOperationalAttributes returns a boolean if a field has been set.
func (o *LdapIdentitySourceDTO) HasIncludeOperationalAttributes() bool {
	if o != nil && !isNil(o.IncludeOperationalAttributes) {
		return true
	}

	return false
}

// SetIncludeOperationalAttributes gets a reference to the given bool and assigns it to the IncludeOperationalAttributes field.
func (o *LdapIdentitySourceDTO) SetIncludeOperationalAttributes(v bool) {
	o.IncludeOperationalAttributes = &v
}

// GetInitialContextFactory returns the InitialContextFactory field value if set, zero value otherwise.
func (o *LdapIdentitySourceDTO) GetInitialContextFactory() string {
	if o == nil || isNil(o.InitialContextFactory) {
		var ret string
		return ret
	}
	return *o.InitialContextFactory
}

// GetInitialContextFactoryOk returns a tuple with the InitialContextFactory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapIdentitySourceDTO) GetInitialContextFactoryOk() (*string, bool) {
	if o == nil || isNil(o.InitialContextFactory) {
    return nil, false
	}
	return o.InitialContextFactory, true
}

// HasInitialContextFactory returns a boolean if a field has been set.
func (o *LdapIdentitySourceDTO) HasInitialContextFactory() bool {
	if o != nil && !isNil(o.InitialContextFactory) {
		return true
	}

	return false
}

// SetInitialContextFactory gets a reference to the given string and assigns it to the InitialContextFactory field.
func (o *LdapIdentitySourceDTO) SetInitialContextFactory(v string) {
	o.InitialContextFactory = &v
}

// GetLdapSearchScope returns the LdapSearchScope field value if set, zero value otherwise.
func (o *LdapIdentitySourceDTO) GetLdapSearchScope() string {
	if o == nil || isNil(o.LdapSearchScope) {
		var ret string
		return ret
	}
	return *o.LdapSearchScope
}

// GetLdapSearchScopeOk returns a tuple with the LdapSearchScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapIdentitySourceDTO) GetLdapSearchScopeOk() (*string, bool) {
	if o == nil || isNil(o.LdapSearchScope) {
    return nil, false
	}
	return o.LdapSearchScope, true
}

// HasLdapSearchScope returns a boolean if a field has been set.
func (o *LdapIdentitySourceDTO) HasLdapSearchScope() bool {
	if o != nil && !isNil(o.LdapSearchScope) {
		return true
	}

	return false
}

// SetLdapSearchScope gets a reference to the given string and assigns it to the LdapSearchScope field.
func (o *LdapIdentitySourceDTO) SetLdapSearchScope(v string) {
	o.LdapSearchScope = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LdapIdentitySourceDTO) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapIdentitySourceDTO) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LdapIdentitySourceDTO) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LdapIdentitySourceDTO) SetName(v string) {
	o.Name = &v
}

// GetPrincipalUidAttributeID returns the PrincipalUidAttributeID field value if set, zero value otherwise.
func (o *LdapIdentitySourceDTO) GetPrincipalUidAttributeID() string {
	if o == nil || isNil(o.PrincipalUidAttributeID) {
		var ret string
		return ret
	}
	return *o.PrincipalUidAttributeID
}

// GetPrincipalUidAttributeIDOk returns a tuple with the PrincipalUidAttributeID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapIdentitySourceDTO) GetPrincipalUidAttributeIDOk() (*string, bool) {
	if o == nil || isNil(o.PrincipalUidAttributeID) {
    return nil, false
	}
	return o.PrincipalUidAttributeID, true
}

// HasPrincipalUidAttributeID returns a boolean if a field has been set.
func (o *LdapIdentitySourceDTO) HasPrincipalUidAttributeID() bool {
	if o != nil && !isNil(o.PrincipalUidAttributeID) {
		return true
	}

	return false
}

// SetPrincipalUidAttributeID gets a reference to the given string and assigns it to the PrincipalUidAttributeID field.
func (o *LdapIdentitySourceDTO) SetPrincipalUidAttributeID(v string) {
	o.PrincipalUidAttributeID = &v
}

// GetProviderUrl returns the ProviderUrl field value if set, zero value otherwise.
func (o *LdapIdentitySourceDTO) GetProviderUrl() string {
	if o == nil || isNil(o.ProviderUrl) {
		var ret string
		return ret
	}
	return *o.ProviderUrl
}

// GetProviderUrlOk returns a tuple with the ProviderUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapIdentitySourceDTO) GetProviderUrlOk() (*string, bool) {
	if o == nil || isNil(o.ProviderUrl) {
    return nil, false
	}
	return o.ProviderUrl, true
}

// HasProviderUrl returns a boolean if a field has been set.
func (o *LdapIdentitySourceDTO) HasProviderUrl() bool {
	if o != nil && !isNil(o.ProviderUrl) {
		return true
	}

	return false
}

// SetProviderUrl gets a reference to the given string and assigns it to the ProviderUrl field.
func (o *LdapIdentitySourceDTO) SetProviderUrl(v string) {
	o.ProviderUrl = &v
}

// GetReferrals returns the Referrals field value if set, zero value otherwise.
func (o *LdapIdentitySourceDTO) GetReferrals() string {
	if o == nil || isNil(o.Referrals) {
		var ret string
		return ret
	}
	return *o.Referrals
}

// GetReferralsOk returns a tuple with the Referrals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapIdentitySourceDTO) GetReferralsOk() (*string, bool) {
	if o == nil || isNil(o.Referrals) {
    return nil, false
	}
	return o.Referrals, true
}

// HasReferrals returns a boolean if a field has been set.
func (o *LdapIdentitySourceDTO) HasReferrals() bool {
	if o != nil && !isNil(o.Referrals) {
		return true
	}

	return false
}

// SetReferrals gets a reference to the given string and assigns it to the Referrals field.
func (o *LdapIdentitySourceDTO) SetReferrals(v string) {
	o.Referrals = &v
}

// GetRoleAttributeID returns the RoleAttributeID field value if set, zero value otherwise.
func (o *LdapIdentitySourceDTO) GetRoleAttributeID() string {
	if o == nil || isNil(o.RoleAttributeID) {
		var ret string
		return ret
	}
	return *o.RoleAttributeID
}

// GetRoleAttributeIDOk returns a tuple with the RoleAttributeID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapIdentitySourceDTO) GetRoleAttributeIDOk() (*string, bool) {
	if o == nil || isNil(o.RoleAttributeID) {
    return nil, false
	}
	return o.RoleAttributeID, true
}

// HasRoleAttributeID returns a boolean if a field has been set.
func (o *LdapIdentitySourceDTO) HasRoleAttributeID() bool {
	if o != nil && !isNil(o.RoleAttributeID) {
		return true
	}

	return false
}

// SetRoleAttributeID gets a reference to the given string and assigns it to the RoleAttributeID field.
func (o *LdapIdentitySourceDTO) SetRoleAttributeID(v string) {
	o.RoleAttributeID = &v
}

// GetRoleMatchingMode returns the RoleMatchingMode field value if set, zero value otherwise.
func (o *LdapIdentitySourceDTO) GetRoleMatchingMode() string {
	if o == nil || isNil(o.RoleMatchingMode) {
		var ret string
		return ret
	}
	return *o.RoleMatchingMode
}

// GetRoleMatchingModeOk returns a tuple with the RoleMatchingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapIdentitySourceDTO) GetRoleMatchingModeOk() (*string, bool) {
	if o == nil || isNil(o.RoleMatchingMode) {
    return nil, false
	}
	return o.RoleMatchingMode, true
}

// HasRoleMatchingMode returns a boolean if a field has been set.
func (o *LdapIdentitySourceDTO) HasRoleMatchingMode() bool {
	if o != nil && !isNil(o.RoleMatchingMode) {
		return true
	}

	return false
}

// SetRoleMatchingMode gets a reference to the given string and assigns it to the RoleMatchingMode field.
func (o *LdapIdentitySourceDTO) SetRoleMatchingMode(v string) {
	o.RoleMatchingMode = &v
}

// GetRolesCtxDN returns the RolesCtxDN field value if set, zero value otherwise.
func (o *LdapIdentitySourceDTO) GetRolesCtxDN() string {
	if o == nil || isNil(o.RolesCtxDN) {
		var ret string
		return ret
	}
	return *o.RolesCtxDN
}

// GetRolesCtxDNOk returns a tuple with the RolesCtxDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapIdentitySourceDTO) GetRolesCtxDNOk() (*string, bool) {
	if o == nil || isNil(o.RolesCtxDN) {
    return nil, false
	}
	return o.RolesCtxDN, true
}

// HasRolesCtxDN returns a boolean if a field has been set.
func (o *LdapIdentitySourceDTO) HasRolesCtxDN() bool {
	if o != nil && !isNil(o.RolesCtxDN) {
		return true
	}

	return false
}

// SetRolesCtxDN gets a reference to the given string and assigns it to the RolesCtxDN field.
func (o *LdapIdentitySourceDTO) SetRolesCtxDN(v string) {
	o.RolesCtxDN = &v
}

// GetSecurityAuthentication returns the SecurityAuthentication field value if set, zero value otherwise.
func (o *LdapIdentitySourceDTO) GetSecurityAuthentication() string {
	if o == nil || isNil(o.SecurityAuthentication) {
		var ret string
		return ret
	}
	return *o.SecurityAuthentication
}

// GetSecurityAuthenticationOk returns a tuple with the SecurityAuthentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapIdentitySourceDTO) GetSecurityAuthenticationOk() (*string, bool) {
	if o == nil || isNil(o.SecurityAuthentication) {
    return nil, false
	}
	return o.SecurityAuthentication, true
}

// HasSecurityAuthentication returns a boolean if a field has been set.
func (o *LdapIdentitySourceDTO) HasSecurityAuthentication() bool {
	if o != nil && !isNil(o.SecurityAuthentication) {
		return true
	}

	return false
}

// SetSecurityAuthentication gets a reference to the given string and assigns it to the SecurityAuthentication field.
func (o *LdapIdentitySourceDTO) SetSecurityAuthentication(v string) {
	o.SecurityAuthentication = &v
}

// GetSecurityCredential returns the SecurityCredential field value if set, zero value otherwise.
func (o *LdapIdentitySourceDTO) GetSecurityCredential() string {
	if o == nil || isNil(o.SecurityCredential) {
		var ret string
		return ret
	}
	return *o.SecurityCredential
}

// GetSecurityCredentialOk returns a tuple with the SecurityCredential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapIdentitySourceDTO) GetSecurityCredentialOk() (*string, bool) {
	if o == nil || isNil(o.SecurityCredential) {
    return nil, false
	}
	return o.SecurityCredential, true
}

// HasSecurityCredential returns a boolean if a field has been set.
func (o *LdapIdentitySourceDTO) HasSecurityCredential() bool {
	if o != nil && !isNil(o.SecurityCredential) {
		return true
	}

	return false
}

// SetSecurityCredential gets a reference to the given string and assigns it to the SecurityCredential field.
func (o *LdapIdentitySourceDTO) SetSecurityCredential(v string) {
	o.SecurityCredential = &v
}

// GetSecurityPrincipal returns the SecurityPrincipal field value if set, zero value otherwise.
func (o *LdapIdentitySourceDTO) GetSecurityPrincipal() string {
	if o == nil || isNil(o.SecurityPrincipal) {
		var ret string
		return ret
	}
	return *o.SecurityPrincipal
}

// GetSecurityPrincipalOk returns a tuple with the SecurityPrincipal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapIdentitySourceDTO) GetSecurityPrincipalOk() (*string, bool) {
	if o == nil || isNil(o.SecurityPrincipal) {
    return nil, false
	}
	return o.SecurityPrincipal, true
}

// HasSecurityPrincipal returns a boolean if a field has been set.
func (o *LdapIdentitySourceDTO) HasSecurityPrincipal() bool {
	if o != nil && !isNil(o.SecurityPrincipal) {
		return true
	}

	return false
}

// SetSecurityPrincipal gets a reference to the given string and assigns it to the SecurityPrincipal field.
func (o *LdapIdentitySourceDTO) SetSecurityPrincipal(v string) {
	o.SecurityPrincipal = &v
}

// GetUidAttributeID returns the UidAttributeID field value if set, zero value otherwise.
func (o *LdapIdentitySourceDTO) GetUidAttributeID() string {
	if o == nil || isNil(o.UidAttributeID) {
		var ret string
		return ret
	}
	return *o.UidAttributeID
}

// GetUidAttributeIDOk returns a tuple with the UidAttributeID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapIdentitySourceDTO) GetUidAttributeIDOk() (*string, bool) {
	if o == nil || isNil(o.UidAttributeID) {
    return nil, false
	}
	return o.UidAttributeID, true
}

// HasUidAttributeID returns a boolean if a field has been set.
func (o *LdapIdentitySourceDTO) HasUidAttributeID() bool {
	if o != nil && !isNil(o.UidAttributeID) {
		return true
	}

	return false
}

// SetUidAttributeID gets a reference to the given string and assigns it to the UidAttributeID field.
func (o *LdapIdentitySourceDTO) SetUidAttributeID(v string) {
	o.UidAttributeID = &v
}

// GetUpdatePasswordEnabled returns the UpdatePasswordEnabled field value if set, zero value otherwise.
func (o *LdapIdentitySourceDTO) GetUpdatePasswordEnabled() bool {
	if o == nil || isNil(o.UpdatePasswordEnabled) {
		var ret bool
		return ret
	}
	return *o.UpdatePasswordEnabled
}

// GetUpdatePasswordEnabledOk returns a tuple with the UpdatePasswordEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapIdentitySourceDTO) GetUpdatePasswordEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.UpdatePasswordEnabled) {
    return nil, false
	}
	return o.UpdatePasswordEnabled, true
}

// HasUpdatePasswordEnabled returns a boolean if a field has been set.
func (o *LdapIdentitySourceDTO) HasUpdatePasswordEnabled() bool {
	if o != nil && !isNil(o.UpdatePasswordEnabled) {
		return true
	}

	return false
}

// SetUpdatePasswordEnabled gets a reference to the given bool and assigns it to the UpdatePasswordEnabled field.
func (o *LdapIdentitySourceDTO) SetUpdatePasswordEnabled(v bool) {
	o.UpdatePasswordEnabled = &v
}

// GetUpdateableCredentialAttribute returns the UpdateableCredentialAttribute field value if set, zero value otherwise.
func (o *LdapIdentitySourceDTO) GetUpdateableCredentialAttribute() string {
	if o == nil || isNil(o.UpdateableCredentialAttribute) {
		var ret string
		return ret
	}
	return *o.UpdateableCredentialAttribute
}

// GetUpdateableCredentialAttributeOk returns a tuple with the UpdateableCredentialAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapIdentitySourceDTO) GetUpdateableCredentialAttributeOk() (*string, bool) {
	if o == nil || isNil(o.UpdateableCredentialAttribute) {
    return nil, false
	}
	return o.UpdateableCredentialAttribute, true
}

// HasUpdateableCredentialAttribute returns a boolean if a field has been set.
func (o *LdapIdentitySourceDTO) HasUpdateableCredentialAttribute() bool {
	if o != nil && !isNil(o.UpdateableCredentialAttribute) {
		return true
	}

	return false
}

// SetUpdateableCredentialAttribute gets a reference to the given string and assigns it to the UpdateableCredentialAttribute field.
func (o *LdapIdentitySourceDTO) SetUpdateableCredentialAttribute(v string) {
	o.UpdateableCredentialAttribute = &v
}

// GetUserPropertiesQueryString returns the UserPropertiesQueryString field value if set, zero value otherwise.
func (o *LdapIdentitySourceDTO) GetUserPropertiesQueryString() string {
	if o == nil || isNil(o.UserPropertiesQueryString) {
		var ret string
		return ret
	}
	return *o.UserPropertiesQueryString
}

// GetUserPropertiesQueryStringOk returns a tuple with the UserPropertiesQueryString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapIdentitySourceDTO) GetUserPropertiesQueryStringOk() (*string, bool) {
	if o == nil || isNil(o.UserPropertiesQueryString) {
    return nil, false
	}
	return o.UserPropertiesQueryString, true
}

// HasUserPropertiesQueryString returns a boolean if a field has been set.
func (o *LdapIdentitySourceDTO) HasUserPropertiesQueryString() bool {
	if o != nil && !isNil(o.UserPropertiesQueryString) {
		return true
	}

	return false
}

// SetUserPropertiesQueryString gets a reference to the given string and assigns it to the UserPropertiesQueryString field.
func (o *LdapIdentitySourceDTO) SetUserPropertiesQueryString(v string) {
	o.UserPropertiesQueryString = &v
}

// GetUsersCtxDN returns the UsersCtxDN field value if set, zero value otherwise.
func (o *LdapIdentitySourceDTO) GetUsersCtxDN() string {
	if o == nil || isNil(o.UsersCtxDN) {
		var ret string
		return ret
	}
	return *o.UsersCtxDN
}

// GetUsersCtxDNOk returns a tuple with the UsersCtxDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapIdentitySourceDTO) GetUsersCtxDNOk() (*string, bool) {
	if o == nil || isNil(o.UsersCtxDN) {
    return nil, false
	}
	return o.UsersCtxDN, true
}

// HasUsersCtxDN returns a boolean if a field has been set.
func (o *LdapIdentitySourceDTO) HasUsersCtxDN() bool {
	if o != nil && !isNil(o.UsersCtxDN) {
		return true
	}

	return false
}

// SetUsersCtxDN gets a reference to the given string and assigns it to the UsersCtxDN field.
func (o *LdapIdentitySourceDTO) SetUsersCtxDN(v string) {
	o.UsersCtxDN = &v
}

// GetX returns the X field value if set, zero value otherwise.
func (o *LdapIdentitySourceDTO) GetX() float64 {
	if o == nil || isNil(o.X) {
		var ret float64
		return ret
	}
	return *o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapIdentitySourceDTO) GetXOk() (*float64, bool) {
	if o == nil || isNil(o.X) {
    return nil, false
	}
	return o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *LdapIdentitySourceDTO) HasX() bool {
	if o != nil && !isNil(o.X) {
		return true
	}

	return false
}

// SetX gets a reference to the given float64 and assigns it to the X field.
func (o *LdapIdentitySourceDTO) SetX(v float64) {
	o.X = &v
}

// GetY returns the Y field value if set, zero value otherwise.
func (o *LdapIdentitySourceDTO) GetY() float64 {
	if o == nil || isNil(o.Y) {
		var ret float64
		return ret
	}
	return *o.Y
}

// GetYOk returns a tuple with the Y field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapIdentitySourceDTO) GetYOk() (*float64, bool) {
	if o == nil || isNil(o.Y) {
    return nil, false
	}
	return o.Y, true
}

// HasY returns a boolean if a field has been set.
func (o *LdapIdentitySourceDTO) HasY() bool {
	if o != nil && !isNil(o.Y) {
		return true
	}

	return false
}

// SetY gets a reference to the given float64 and assigns it to the Y field.
func (o *LdapIdentitySourceDTO) SetY(v float64) {
	o.Y = &v
}

func (o LdapIdentitySourceDTO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CredentialQueryString) {
		toSerialize["credentialQueryString"] = o.CredentialQueryString
	}
	if !isNil(o.CustomClass) {
		toSerialize["customClass"] = o.CustomClass
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.ElementId) {
		toSerialize["elementId"] = o.ElementId
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.IncludeOperationalAttributes) {
		toSerialize["includeOperationalAttributes"] = o.IncludeOperationalAttributes
	}
	if !isNil(o.InitialContextFactory) {
		toSerialize["initialContextFactory"] = o.InitialContextFactory
	}
	if !isNil(o.LdapSearchScope) {
		toSerialize["ldapSearchScope"] = o.LdapSearchScope
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.PrincipalUidAttributeID) {
		toSerialize["principalUidAttributeID"] = o.PrincipalUidAttributeID
	}
	if !isNil(o.ProviderUrl) {
		toSerialize["providerUrl"] = o.ProviderUrl
	}
	if !isNil(o.Referrals) {
		toSerialize["referrals"] = o.Referrals
	}
	if !isNil(o.RoleAttributeID) {
		toSerialize["roleAttributeID"] = o.RoleAttributeID
	}
	if !isNil(o.RoleMatchingMode) {
		toSerialize["roleMatchingMode"] = o.RoleMatchingMode
	}
	if !isNil(o.RolesCtxDN) {
		toSerialize["rolesCtxDN"] = o.RolesCtxDN
	}
	if !isNil(o.SecurityAuthentication) {
		toSerialize["securityAuthentication"] = o.SecurityAuthentication
	}
	if !isNil(o.SecurityCredential) {
		toSerialize["securityCredential"] = o.SecurityCredential
	}
	if !isNil(o.SecurityPrincipal) {
		toSerialize["securityPrincipal"] = o.SecurityPrincipal
	}
	if !isNil(o.UidAttributeID) {
		toSerialize["uidAttributeID"] = o.UidAttributeID
	}
	if !isNil(o.UpdatePasswordEnabled) {
		toSerialize["updatePasswordEnabled"] = o.UpdatePasswordEnabled
	}
	if !isNil(o.UpdateableCredentialAttribute) {
		toSerialize["updateableCredentialAttribute"] = o.UpdateableCredentialAttribute
	}
	if !isNil(o.UserPropertiesQueryString) {
		toSerialize["userPropertiesQueryString"] = o.UserPropertiesQueryString
	}
	if !isNil(o.UsersCtxDN) {
		toSerialize["usersCtxDN"] = o.UsersCtxDN
	}
	if !isNil(o.X) {
		toSerialize["x"] = o.X
	}
	if !isNil(o.Y) {
		toSerialize["y"] = o.Y
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LdapIdentitySourceDTO) UnmarshalJSON(bytes []byte) (err error) {
	varLdapIdentitySourceDTO := _LdapIdentitySourceDTO{}

	if err = json.Unmarshal(bytes, &varLdapIdentitySourceDTO); err == nil {
		*o = LdapIdentitySourceDTO(varLdapIdentitySourceDTO)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "credentialQueryString")
		delete(additionalProperties, "customClass")
		delete(additionalProperties, "description")
		delete(additionalProperties, "elementId")
		delete(additionalProperties, "id")
		delete(additionalProperties, "includeOperationalAttributes")
		delete(additionalProperties, "initialContextFactory")
		delete(additionalProperties, "ldapSearchScope")
		delete(additionalProperties, "name")
		delete(additionalProperties, "principalUidAttributeID")
		delete(additionalProperties, "providerUrl")
		delete(additionalProperties, "referrals")
		delete(additionalProperties, "roleAttributeID")
		delete(additionalProperties, "roleMatchingMode")
		delete(additionalProperties, "rolesCtxDN")
		delete(additionalProperties, "securityAuthentication")
		delete(additionalProperties, "securityCredential")
		delete(additionalProperties, "securityPrincipal")
		delete(additionalProperties, "uidAttributeID")
		delete(additionalProperties, "updatePasswordEnabled")
		delete(additionalProperties, "updateableCredentialAttribute")
		delete(additionalProperties, "userPropertiesQueryString")
		delete(additionalProperties, "usersCtxDN")
		delete(additionalProperties, "x")
		delete(additionalProperties, "y")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLdapIdentitySourceDTO struct {
	value *LdapIdentitySourceDTO
	isSet bool
}

func (v NullableLdapIdentitySourceDTO) Get() *LdapIdentitySourceDTO {
	return v.value
}

func (v *NullableLdapIdentitySourceDTO) Set(val *LdapIdentitySourceDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableLdapIdentitySourceDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableLdapIdentitySourceDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLdapIdentitySourceDTO(val *LdapIdentitySourceDTO) *NullableLdapIdentitySourceDTO {
	return &NullableLdapIdentitySourceDTO{value: val, isSet: true}
}

func (v NullableLdapIdentitySourceDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLdapIdentitySourceDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


