/*
Atricore Console :: Remote : API

# Atricore Console API

API version: 1.4.3-SNAPSHOT
Contact: sgonzalez@atricore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package jossoappi

import (
	"encoding/json"
)

// InternalSaml2ServiceProviderChannelDTO struct for InternalSaml2ServiceProviderChannelDTO
type InternalSaml2ServiceProviderChannelDTO struct {
	ActiveBindings *[]string `json:"activeBindings,omitempty"`
	ActiveProfiles *[]string `json:"activeProfiles,omitempty"`
	AttributeProfile *AttributeProfileDTO `json:"attributeProfile,omitempty"`
	AuthenticationContract *AuthenticationContractDTO `json:"authenticationContract,omitempty"`
	AuthenticationMechanism *AuthenticationMechanismDTO `json:"authenticationMechanism,omitempty"`
	ConnectionA *FederatedConnectionDTO `json:"connectionA,omitempty"`
	ConnectionB *FederatedConnectionDTO `json:"connectionB,omitempty"`
	Description *string `json:"description,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	ElementId *string `json:"elementId,omitempty"`
	EmissionPolicy *AuthenticationAssertionEmissionPolicyDTO `json:"emissionPolicy,omitempty"`
	EncryptAssertion *bool `json:"encryptAssertion,omitempty"`
	EncryptAssertionAlgorithm *string `json:"encryptAssertionAlgorithm,omitempty"`
	Id *int64 `json:"id,omitempty"`
	IgnoreRequestedNameIDPolicy *bool `json:"ignoreRequestedNameIDPolicy,omitempty"`
	Location *LocationDTO `json:"location,omitempty"`
	MessageTtl *int32 `json:"messageTtl,omitempty"`
	MessageTtlTolerance *int32 `json:"messageTtlTolerance,omitempty"`
	Name *string `json:"name,omitempty"`
	OverrideProviderSetup *bool `json:"overrideProviderSetup,omitempty"`
	RequiredRoles *[]string `json:"requiredRoles,omitempty"`
	RequiredRolesMatchMode *int32 `json:"requiredRolesMatchMode,omitempty"`
	RestrictedRoles *[]string `json:"restrictedRoles,omitempty"`
	RestrictedRolesMatchMode *int32 `json:"restrictedRolesMatchMode,omitempty"`
	SignatureHash *string `json:"signatureHash,omitempty"`
	SubjectNameIDPolicy *SubjectNameIdentifierPolicyDTO `json:"subjectNameIDPolicy,omitempty"`
	WantAuthnRequestsSigned *bool `json:"wantAuthnRequestsSigned,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InternalSaml2ServiceProviderChannelDTO InternalSaml2ServiceProviderChannelDTO

// NewInternalSaml2ServiceProviderChannelDTO instantiates a new InternalSaml2ServiceProviderChannelDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalSaml2ServiceProviderChannelDTO() *InternalSaml2ServiceProviderChannelDTO {
	this := InternalSaml2ServiceProviderChannelDTO{}
	return &this
}

// NewInternalSaml2ServiceProviderChannelDTOWithDefaults instantiates a new InternalSaml2ServiceProviderChannelDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalSaml2ServiceProviderChannelDTOWithDefaults() *InternalSaml2ServiceProviderChannelDTO {
	this := InternalSaml2ServiceProviderChannelDTO{}
	return &this
}

// GetActiveBindings returns the ActiveBindings field value if set, zero value otherwise.
func (o *InternalSaml2ServiceProviderChannelDTO) GetActiveBindings() []string {
	if o == nil || o.ActiveBindings == nil {
		var ret []string
		return ret
	}
	return *o.ActiveBindings
}

// GetActiveBindingsOk returns a tuple with the ActiveBindings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalSaml2ServiceProviderChannelDTO) GetActiveBindingsOk() (*[]string, bool) {
	if o == nil || o.ActiveBindings == nil {
		return nil, false
	}
	return o.ActiveBindings, true
}

// HasActiveBindings returns a boolean if a field has been set.
func (o *InternalSaml2ServiceProviderChannelDTO) HasActiveBindings() bool {
	if o != nil && o.ActiveBindings != nil {
		return true
	}

	return false
}

// SetActiveBindings gets a reference to the given []string and assigns it to the ActiveBindings field.
func (o *InternalSaml2ServiceProviderChannelDTO) SetActiveBindings(v []string) {
	o.ActiveBindings = &v
}

// GetActiveProfiles returns the ActiveProfiles field value if set, zero value otherwise.
func (o *InternalSaml2ServiceProviderChannelDTO) GetActiveProfiles() []string {
	if o == nil || o.ActiveProfiles == nil {
		var ret []string
		return ret
	}
	return *o.ActiveProfiles
}

// GetActiveProfilesOk returns a tuple with the ActiveProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalSaml2ServiceProviderChannelDTO) GetActiveProfilesOk() (*[]string, bool) {
	if o == nil || o.ActiveProfiles == nil {
		return nil, false
	}
	return o.ActiveProfiles, true
}

// HasActiveProfiles returns a boolean if a field has been set.
func (o *InternalSaml2ServiceProviderChannelDTO) HasActiveProfiles() bool {
	if o != nil && o.ActiveProfiles != nil {
		return true
	}

	return false
}

// SetActiveProfiles gets a reference to the given []string and assigns it to the ActiveProfiles field.
func (o *InternalSaml2ServiceProviderChannelDTO) SetActiveProfiles(v []string) {
	o.ActiveProfiles = &v
}

// GetAttributeProfile returns the AttributeProfile field value if set, zero value otherwise.
func (o *InternalSaml2ServiceProviderChannelDTO) GetAttributeProfile() AttributeProfileDTO {
	if o == nil || o.AttributeProfile == nil {
		var ret AttributeProfileDTO
		return ret
	}
	return *o.AttributeProfile
}

// GetAttributeProfileOk returns a tuple with the AttributeProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalSaml2ServiceProviderChannelDTO) GetAttributeProfileOk() (*AttributeProfileDTO, bool) {
	if o == nil || o.AttributeProfile == nil {
		return nil, false
	}
	return o.AttributeProfile, true
}

// HasAttributeProfile returns a boolean if a field has been set.
func (o *InternalSaml2ServiceProviderChannelDTO) HasAttributeProfile() bool {
	if o != nil && o.AttributeProfile != nil {
		return true
	}

	return false
}

// SetAttributeProfile gets a reference to the given AttributeProfileDTO and assigns it to the AttributeProfile field.
func (o *InternalSaml2ServiceProviderChannelDTO) SetAttributeProfile(v AttributeProfileDTO) {
	o.AttributeProfile = &v
}

// GetAuthenticationContract returns the AuthenticationContract field value if set, zero value otherwise.
func (o *InternalSaml2ServiceProviderChannelDTO) GetAuthenticationContract() AuthenticationContractDTO {
	if o == nil || o.AuthenticationContract == nil {
		var ret AuthenticationContractDTO
		return ret
	}
	return *o.AuthenticationContract
}

// GetAuthenticationContractOk returns a tuple with the AuthenticationContract field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalSaml2ServiceProviderChannelDTO) GetAuthenticationContractOk() (*AuthenticationContractDTO, bool) {
	if o == nil || o.AuthenticationContract == nil {
		return nil, false
	}
	return o.AuthenticationContract, true
}

// HasAuthenticationContract returns a boolean if a field has been set.
func (o *InternalSaml2ServiceProviderChannelDTO) HasAuthenticationContract() bool {
	if o != nil && o.AuthenticationContract != nil {
		return true
	}

	return false
}

// SetAuthenticationContract gets a reference to the given AuthenticationContractDTO and assigns it to the AuthenticationContract field.
func (o *InternalSaml2ServiceProviderChannelDTO) SetAuthenticationContract(v AuthenticationContractDTO) {
	o.AuthenticationContract = &v
}

// GetAuthenticationMechanism returns the AuthenticationMechanism field value if set, zero value otherwise.
func (o *InternalSaml2ServiceProviderChannelDTO) GetAuthenticationMechanism() AuthenticationMechanismDTO {
	if o == nil || o.AuthenticationMechanism == nil {
		var ret AuthenticationMechanismDTO
		return ret
	}
	return *o.AuthenticationMechanism
}

// GetAuthenticationMechanismOk returns a tuple with the AuthenticationMechanism field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalSaml2ServiceProviderChannelDTO) GetAuthenticationMechanismOk() (*AuthenticationMechanismDTO, bool) {
	if o == nil || o.AuthenticationMechanism == nil {
		return nil, false
	}
	return o.AuthenticationMechanism, true
}

// HasAuthenticationMechanism returns a boolean if a field has been set.
func (o *InternalSaml2ServiceProviderChannelDTO) HasAuthenticationMechanism() bool {
	if o != nil && o.AuthenticationMechanism != nil {
		return true
	}

	return false
}

// SetAuthenticationMechanism gets a reference to the given AuthenticationMechanismDTO and assigns it to the AuthenticationMechanism field.
func (o *InternalSaml2ServiceProviderChannelDTO) SetAuthenticationMechanism(v AuthenticationMechanismDTO) {
	o.AuthenticationMechanism = &v
}

// GetConnectionA returns the ConnectionA field value if set, zero value otherwise.
func (o *InternalSaml2ServiceProviderChannelDTO) GetConnectionA() FederatedConnectionDTO {
	if o == nil || o.ConnectionA == nil {
		var ret FederatedConnectionDTO
		return ret
	}
	return *o.ConnectionA
}

// GetConnectionAOk returns a tuple with the ConnectionA field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalSaml2ServiceProviderChannelDTO) GetConnectionAOk() (*FederatedConnectionDTO, bool) {
	if o == nil || o.ConnectionA == nil {
		return nil, false
	}
	return o.ConnectionA, true
}

// HasConnectionA returns a boolean if a field has been set.
func (o *InternalSaml2ServiceProviderChannelDTO) HasConnectionA() bool {
	if o != nil && o.ConnectionA != nil {
		return true
	}

	return false
}

// SetConnectionA gets a reference to the given FederatedConnectionDTO and assigns it to the ConnectionA field.
func (o *InternalSaml2ServiceProviderChannelDTO) SetConnectionA(v FederatedConnectionDTO) {
	o.ConnectionA = &v
}

// GetConnectionB returns the ConnectionB field value if set, zero value otherwise.
func (o *InternalSaml2ServiceProviderChannelDTO) GetConnectionB() FederatedConnectionDTO {
	if o == nil || o.ConnectionB == nil {
		var ret FederatedConnectionDTO
		return ret
	}
	return *o.ConnectionB
}

// GetConnectionBOk returns a tuple with the ConnectionB field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalSaml2ServiceProviderChannelDTO) GetConnectionBOk() (*FederatedConnectionDTO, bool) {
	if o == nil || o.ConnectionB == nil {
		return nil, false
	}
	return o.ConnectionB, true
}

// HasConnectionB returns a boolean if a field has been set.
func (o *InternalSaml2ServiceProviderChannelDTO) HasConnectionB() bool {
	if o != nil && o.ConnectionB != nil {
		return true
	}

	return false
}

// SetConnectionB gets a reference to the given FederatedConnectionDTO and assigns it to the ConnectionB field.
func (o *InternalSaml2ServiceProviderChannelDTO) SetConnectionB(v FederatedConnectionDTO) {
	o.ConnectionB = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InternalSaml2ServiceProviderChannelDTO) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalSaml2ServiceProviderChannelDTO) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InternalSaml2ServiceProviderChannelDTO) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InternalSaml2ServiceProviderChannelDTO) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *InternalSaml2ServiceProviderChannelDTO) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalSaml2ServiceProviderChannelDTO) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *InternalSaml2ServiceProviderChannelDTO) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *InternalSaml2ServiceProviderChannelDTO) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *InternalSaml2ServiceProviderChannelDTO) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalSaml2ServiceProviderChannelDTO) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *InternalSaml2ServiceProviderChannelDTO) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *InternalSaml2ServiceProviderChannelDTO) SetElementId(v string) {
	o.ElementId = &v
}

// GetEmissionPolicy returns the EmissionPolicy field value if set, zero value otherwise.
func (o *InternalSaml2ServiceProviderChannelDTO) GetEmissionPolicy() AuthenticationAssertionEmissionPolicyDTO {
	if o == nil || o.EmissionPolicy == nil {
		var ret AuthenticationAssertionEmissionPolicyDTO
		return ret
	}
	return *o.EmissionPolicy
}

// GetEmissionPolicyOk returns a tuple with the EmissionPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalSaml2ServiceProviderChannelDTO) GetEmissionPolicyOk() (*AuthenticationAssertionEmissionPolicyDTO, bool) {
	if o == nil || o.EmissionPolicy == nil {
		return nil, false
	}
	return o.EmissionPolicy, true
}

// HasEmissionPolicy returns a boolean if a field has been set.
func (o *InternalSaml2ServiceProviderChannelDTO) HasEmissionPolicy() bool {
	if o != nil && o.EmissionPolicy != nil {
		return true
	}

	return false
}

// SetEmissionPolicy gets a reference to the given AuthenticationAssertionEmissionPolicyDTO and assigns it to the EmissionPolicy field.
func (o *InternalSaml2ServiceProviderChannelDTO) SetEmissionPolicy(v AuthenticationAssertionEmissionPolicyDTO) {
	o.EmissionPolicy = &v
}

// GetEncryptAssertion returns the EncryptAssertion field value if set, zero value otherwise.
func (o *InternalSaml2ServiceProviderChannelDTO) GetEncryptAssertion() bool {
	if o == nil || o.EncryptAssertion == nil {
		var ret bool
		return ret
	}
	return *o.EncryptAssertion
}

// GetEncryptAssertionOk returns a tuple with the EncryptAssertion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalSaml2ServiceProviderChannelDTO) GetEncryptAssertionOk() (*bool, bool) {
	if o == nil || o.EncryptAssertion == nil {
		return nil, false
	}
	return o.EncryptAssertion, true
}

// HasEncryptAssertion returns a boolean if a field has been set.
func (o *InternalSaml2ServiceProviderChannelDTO) HasEncryptAssertion() bool {
	if o != nil && o.EncryptAssertion != nil {
		return true
	}

	return false
}

// SetEncryptAssertion gets a reference to the given bool and assigns it to the EncryptAssertion field.
func (o *InternalSaml2ServiceProviderChannelDTO) SetEncryptAssertion(v bool) {
	o.EncryptAssertion = &v
}

// GetEncryptAssertionAlgorithm returns the EncryptAssertionAlgorithm field value if set, zero value otherwise.
func (o *InternalSaml2ServiceProviderChannelDTO) GetEncryptAssertionAlgorithm() string {
	if o == nil || o.EncryptAssertionAlgorithm == nil {
		var ret string
		return ret
	}
	return *o.EncryptAssertionAlgorithm
}

// GetEncryptAssertionAlgorithmOk returns a tuple with the EncryptAssertionAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalSaml2ServiceProviderChannelDTO) GetEncryptAssertionAlgorithmOk() (*string, bool) {
	if o == nil || o.EncryptAssertionAlgorithm == nil {
		return nil, false
	}
	return o.EncryptAssertionAlgorithm, true
}

// HasEncryptAssertionAlgorithm returns a boolean if a field has been set.
func (o *InternalSaml2ServiceProviderChannelDTO) HasEncryptAssertionAlgorithm() bool {
	if o != nil && o.EncryptAssertionAlgorithm != nil {
		return true
	}

	return false
}

// SetEncryptAssertionAlgorithm gets a reference to the given string and assigns it to the EncryptAssertionAlgorithm field.
func (o *InternalSaml2ServiceProviderChannelDTO) SetEncryptAssertionAlgorithm(v string) {
	o.EncryptAssertionAlgorithm = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InternalSaml2ServiceProviderChannelDTO) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalSaml2ServiceProviderChannelDTO) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InternalSaml2ServiceProviderChannelDTO) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *InternalSaml2ServiceProviderChannelDTO) SetId(v int64) {
	o.Id = &v
}

// GetIgnoreRequestedNameIDPolicy returns the IgnoreRequestedNameIDPolicy field value if set, zero value otherwise.
func (o *InternalSaml2ServiceProviderChannelDTO) GetIgnoreRequestedNameIDPolicy() bool {
	if o == nil || o.IgnoreRequestedNameIDPolicy == nil {
		var ret bool
		return ret
	}
	return *o.IgnoreRequestedNameIDPolicy
}

// GetIgnoreRequestedNameIDPolicyOk returns a tuple with the IgnoreRequestedNameIDPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalSaml2ServiceProviderChannelDTO) GetIgnoreRequestedNameIDPolicyOk() (*bool, bool) {
	if o == nil || o.IgnoreRequestedNameIDPolicy == nil {
		return nil, false
	}
	return o.IgnoreRequestedNameIDPolicy, true
}

// HasIgnoreRequestedNameIDPolicy returns a boolean if a field has been set.
func (o *InternalSaml2ServiceProviderChannelDTO) HasIgnoreRequestedNameIDPolicy() bool {
	if o != nil && o.IgnoreRequestedNameIDPolicy != nil {
		return true
	}

	return false
}

// SetIgnoreRequestedNameIDPolicy gets a reference to the given bool and assigns it to the IgnoreRequestedNameIDPolicy field.
func (o *InternalSaml2ServiceProviderChannelDTO) SetIgnoreRequestedNameIDPolicy(v bool) {
	o.IgnoreRequestedNameIDPolicy = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *InternalSaml2ServiceProviderChannelDTO) GetLocation() LocationDTO {
	if o == nil || o.Location == nil {
		var ret LocationDTO
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalSaml2ServiceProviderChannelDTO) GetLocationOk() (*LocationDTO, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *InternalSaml2ServiceProviderChannelDTO) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given LocationDTO and assigns it to the Location field.
func (o *InternalSaml2ServiceProviderChannelDTO) SetLocation(v LocationDTO) {
	o.Location = &v
}

// GetMessageTtl returns the MessageTtl field value if set, zero value otherwise.
func (o *InternalSaml2ServiceProviderChannelDTO) GetMessageTtl() int32 {
	if o == nil || o.MessageTtl == nil {
		var ret int32
		return ret
	}
	return *o.MessageTtl
}

// GetMessageTtlOk returns a tuple with the MessageTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalSaml2ServiceProviderChannelDTO) GetMessageTtlOk() (*int32, bool) {
	if o == nil || o.MessageTtl == nil {
		return nil, false
	}
	return o.MessageTtl, true
}

// HasMessageTtl returns a boolean if a field has been set.
func (o *InternalSaml2ServiceProviderChannelDTO) HasMessageTtl() bool {
	if o != nil && o.MessageTtl != nil {
		return true
	}

	return false
}

// SetMessageTtl gets a reference to the given int32 and assigns it to the MessageTtl field.
func (o *InternalSaml2ServiceProviderChannelDTO) SetMessageTtl(v int32) {
	o.MessageTtl = &v
}

// GetMessageTtlTolerance returns the MessageTtlTolerance field value if set, zero value otherwise.
func (o *InternalSaml2ServiceProviderChannelDTO) GetMessageTtlTolerance() int32 {
	if o == nil || o.MessageTtlTolerance == nil {
		var ret int32
		return ret
	}
	return *o.MessageTtlTolerance
}

// GetMessageTtlToleranceOk returns a tuple with the MessageTtlTolerance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalSaml2ServiceProviderChannelDTO) GetMessageTtlToleranceOk() (*int32, bool) {
	if o == nil || o.MessageTtlTolerance == nil {
		return nil, false
	}
	return o.MessageTtlTolerance, true
}

// HasMessageTtlTolerance returns a boolean if a field has been set.
func (o *InternalSaml2ServiceProviderChannelDTO) HasMessageTtlTolerance() bool {
	if o != nil && o.MessageTtlTolerance != nil {
		return true
	}

	return false
}

// SetMessageTtlTolerance gets a reference to the given int32 and assigns it to the MessageTtlTolerance field.
func (o *InternalSaml2ServiceProviderChannelDTO) SetMessageTtlTolerance(v int32) {
	o.MessageTtlTolerance = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InternalSaml2ServiceProviderChannelDTO) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalSaml2ServiceProviderChannelDTO) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InternalSaml2ServiceProviderChannelDTO) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InternalSaml2ServiceProviderChannelDTO) SetName(v string) {
	o.Name = &v
}

// GetOverrideProviderSetup returns the OverrideProviderSetup field value if set, zero value otherwise.
func (o *InternalSaml2ServiceProviderChannelDTO) GetOverrideProviderSetup() bool {
	if o == nil || o.OverrideProviderSetup == nil {
		var ret bool
		return ret
	}
	return *o.OverrideProviderSetup
}

// GetOverrideProviderSetupOk returns a tuple with the OverrideProviderSetup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalSaml2ServiceProviderChannelDTO) GetOverrideProviderSetupOk() (*bool, bool) {
	if o == nil || o.OverrideProviderSetup == nil {
		return nil, false
	}
	return o.OverrideProviderSetup, true
}

// HasOverrideProviderSetup returns a boolean if a field has been set.
func (o *InternalSaml2ServiceProviderChannelDTO) HasOverrideProviderSetup() bool {
	if o != nil && o.OverrideProviderSetup != nil {
		return true
	}

	return false
}

// SetOverrideProviderSetup gets a reference to the given bool and assigns it to the OverrideProviderSetup field.
func (o *InternalSaml2ServiceProviderChannelDTO) SetOverrideProviderSetup(v bool) {
	o.OverrideProviderSetup = &v
}

// GetRequiredRoles returns the RequiredRoles field value if set, zero value otherwise.
func (o *InternalSaml2ServiceProviderChannelDTO) GetRequiredRoles() []string {
	if o == nil || o.RequiredRoles == nil {
		var ret []string
		return ret
	}
	return *o.RequiredRoles
}

// GetRequiredRolesOk returns a tuple with the RequiredRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalSaml2ServiceProviderChannelDTO) GetRequiredRolesOk() (*[]string, bool) {
	if o == nil || o.RequiredRoles == nil {
		return nil, false
	}
	return o.RequiredRoles, true
}

// HasRequiredRoles returns a boolean if a field has been set.
func (o *InternalSaml2ServiceProviderChannelDTO) HasRequiredRoles() bool {
	if o != nil && o.RequiredRoles != nil {
		return true
	}

	return false
}

// SetRequiredRoles gets a reference to the given []string and assigns it to the RequiredRoles field.
func (o *InternalSaml2ServiceProviderChannelDTO) SetRequiredRoles(v []string) {
	o.RequiredRoles = &v
}

// GetRequiredRolesMatchMode returns the RequiredRolesMatchMode field value if set, zero value otherwise.
func (o *InternalSaml2ServiceProviderChannelDTO) GetRequiredRolesMatchMode() int32 {
	if o == nil || o.RequiredRolesMatchMode == nil {
		var ret int32
		return ret
	}
	return *o.RequiredRolesMatchMode
}

// GetRequiredRolesMatchModeOk returns a tuple with the RequiredRolesMatchMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalSaml2ServiceProviderChannelDTO) GetRequiredRolesMatchModeOk() (*int32, bool) {
	if o == nil || o.RequiredRolesMatchMode == nil {
		return nil, false
	}
	return o.RequiredRolesMatchMode, true
}

// HasRequiredRolesMatchMode returns a boolean if a field has been set.
func (o *InternalSaml2ServiceProviderChannelDTO) HasRequiredRolesMatchMode() bool {
	if o != nil && o.RequiredRolesMatchMode != nil {
		return true
	}

	return false
}

// SetRequiredRolesMatchMode gets a reference to the given int32 and assigns it to the RequiredRolesMatchMode field.
func (o *InternalSaml2ServiceProviderChannelDTO) SetRequiredRolesMatchMode(v int32) {
	o.RequiredRolesMatchMode = &v
}

// GetRestrictedRoles returns the RestrictedRoles field value if set, zero value otherwise.
func (o *InternalSaml2ServiceProviderChannelDTO) GetRestrictedRoles() []string {
	if o == nil || o.RestrictedRoles == nil {
		var ret []string
		return ret
	}
	return *o.RestrictedRoles
}

// GetRestrictedRolesOk returns a tuple with the RestrictedRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalSaml2ServiceProviderChannelDTO) GetRestrictedRolesOk() (*[]string, bool) {
	if o == nil || o.RestrictedRoles == nil {
		return nil, false
	}
	return o.RestrictedRoles, true
}

// HasRestrictedRoles returns a boolean if a field has been set.
func (o *InternalSaml2ServiceProviderChannelDTO) HasRestrictedRoles() bool {
	if o != nil && o.RestrictedRoles != nil {
		return true
	}

	return false
}

// SetRestrictedRoles gets a reference to the given []string and assigns it to the RestrictedRoles field.
func (o *InternalSaml2ServiceProviderChannelDTO) SetRestrictedRoles(v []string) {
	o.RestrictedRoles = &v
}

// GetRestrictedRolesMatchMode returns the RestrictedRolesMatchMode field value if set, zero value otherwise.
func (o *InternalSaml2ServiceProviderChannelDTO) GetRestrictedRolesMatchMode() int32 {
	if o == nil || o.RestrictedRolesMatchMode == nil {
		var ret int32
		return ret
	}
	return *o.RestrictedRolesMatchMode
}

// GetRestrictedRolesMatchModeOk returns a tuple with the RestrictedRolesMatchMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalSaml2ServiceProviderChannelDTO) GetRestrictedRolesMatchModeOk() (*int32, bool) {
	if o == nil || o.RestrictedRolesMatchMode == nil {
		return nil, false
	}
	return o.RestrictedRolesMatchMode, true
}

// HasRestrictedRolesMatchMode returns a boolean if a field has been set.
func (o *InternalSaml2ServiceProviderChannelDTO) HasRestrictedRolesMatchMode() bool {
	if o != nil && o.RestrictedRolesMatchMode != nil {
		return true
	}

	return false
}

// SetRestrictedRolesMatchMode gets a reference to the given int32 and assigns it to the RestrictedRolesMatchMode field.
func (o *InternalSaml2ServiceProviderChannelDTO) SetRestrictedRolesMatchMode(v int32) {
	o.RestrictedRolesMatchMode = &v
}

// GetSignatureHash returns the SignatureHash field value if set, zero value otherwise.
func (o *InternalSaml2ServiceProviderChannelDTO) GetSignatureHash() string {
	if o == nil || o.SignatureHash == nil {
		var ret string
		return ret
	}
	return *o.SignatureHash
}

// GetSignatureHashOk returns a tuple with the SignatureHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalSaml2ServiceProviderChannelDTO) GetSignatureHashOk() (*string, bool) {
	if o == nil || o.SignatureHash == nil {
		return nil, false
	}
	return o.SignatureHash, true
}

// HasSignatureHash returns a boolean if a field has been set.
func (o *InternalSaml2ServiceProviderChannelDTO) HasSignatureHash() bool {
	if o != nil && o.SignatureHash != nil {
		return true
	}

	return false
}

// SetSignatureHash gets a reference to the given string and assigns it to the SignatureHash field.
func (o *InternalSaml2ServiceProviderChannelDTO) SetSignatureHash(v string) {
	o.SignatureHash = &v
}

// GetSubjectNameIDPolicy returns the SubjectNameIDPolicy field value if set, zero value otherwise.
func (o *InternalSaml2ServiceProviderChannelDTO) GetSubjectNameIDPolicy() SubjectNameIdentifierPolicyDTO {
	if o == nil || o.SubjectNameIDPolicy == nil {
		var ret SubjectNameIdentifierPolicyDTO
		return ret
	}
	return *o.SubjectNameIDPolicy
}

// GetSubjectNameIDPolicyOk returns a tuple with the SubjectNameIDPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalSaml2ServiceProviderChannelDTO) GetSubjectNameIDPolicyOk() (*SubjectNameIdentifierPolicyDTO, bool) {
	if o == nil || o.SubjectNameIDPolicy == nil {
		return nil, false
	}
	return o.SubjectNameIDPolicy, true
}

// HasSubjectNameIDPolicy returns a boolean if a field has been set.
func (o *InternalSaml2ServiceProviderChannelDTO) HasSubjectNameIDPolicy() bool {
	if o != nil && o.SubjectNameIDPolicy != nil {
		return true
	}

	return false
}

// SetSubjectNameIDPolicy gets a reference to the given SubjectNameIdentifierPolicyDTO and assigns it to the SubjectNameIDPolicy field.
func (o *InternalSaml2ServiceProviderChannelDTO) SetSubjectNameIDPolicy(v SubjectNameIdentifierPolicyDTO) {
	o.SubjectNameIDPolicy = &v
}

// GetWantAuthnRequestsSigned returns the WantAuthnRequestsSigned field value if set, zero value otherwise.
func (o *InternalSaml2ServiceProviderChannelDTO) GetWantAuthnRequestsSigned() bool {
	if o == nil || o.WantAuthnRequestsSigned == nil {
		var ret bool
		return ret
	}
	return *o.WantAuthnRequestsSigned
}

// GetWantAuthnRequestsSignedOk returns a tuple with the WantAuthnRequestsSigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalSaml2ServiceProviderChannelDTO) GetWantAuthnRequestsSignedOk() (*bool, bool) {
	if o == nil || o.WantAuthnRequestsSigned == nil {
		return nil, false
	}
	return o.WantAuthnRequestsSigned, true
}

// HasWantAuthnRequestsSigned returns a boolean if a field has been set.
func (o *InternalSaml2ServiceProviderChannelDTO) HasWantAuthnRequestsSigned() bool {
	if o != nil && o.WantAuthnRequestsSigned != nil {
		return true
	}

	return false
}

// SetWantAuthnRequestsSigned gets a reference to the given bool and assigns it to the WantAuthnRequestsSigned field.
func (o *InternalSaml2ServiceProviderChannelDTO) SetWantAuthnRequestsSigned(v bool) {
	o.WantAuthnRequestsSigned = &v
}

func (o InternalSaml2ServiceProviderChannelDTO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ActiveBindings != nil {
		toSerialize["activeBindings"] = o.ActiveBindings
	}
	if o.ActiveProfiles != nil {
		toSerialize["activeProfiles"] = o.ActiveProfiles
	}
	if o.AttributeProfile != nil {
		toSerialize["attributeProfile"] = o.AttributeProfile
	}
	if o.AuthenticationContract != nil {
		toSerialize["authenticationContract"] = o.AuthenticationContract
	}
	if o.AuthenticationMechanism != nil {
		toSerialize["authenticationMechanism"] = o.AuthenticationMechanism
	}
	if o.ConnectionA != nil {
		toSerialize["connectionA"] = o.ConnectionA
	}
	if o.ConnectionB != nil {
		toSerialize["connectionB"] = o.ConnectionB
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.ElementId != nil {
		toSerialize["elementId"] = o.ElementId
	}
	if o.EmissionPolicy != nil {
		toSerialize["emissionPolicy"] = o.EmissionPolicy
	}
	if o.EncryptAssertion != nil {
		toSerialize["encryptAssertion"] = o.EncryptAssertion
	}
	if o.EncryptAssertionAlgorithm != nil {
		toSerialize["encryptAssertionAlgorithm"] = o.EncryptAssertionAlgorithm
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IgnoreRequestedNameIDPolicy != nil {
		toSerialize["ignoreRequestedNameIDPolicy"] = o.IgnoreRequestedNameIDPolicy
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	if o.MessageTtl != nil {
		toSerialize["messageTtl"] = o.MessageTtl
	}
	if o.MessageTtlTolerance != nil {
		toSerialize["messageTtlTolerance"] = o.MessageTtlTolerance
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.OverrideProviderSetup != nil {
		toSerialize["overrideProviderSetup"] = o.OverrideProviderSetup
	}
	if o.RequiredRoles != nil {
		toSerialize["requiredRoles"] = o.RequiredRoles
	}
	if o.RequiredRolesMatchMode != nil {
		toSerialize["requiredRolesMatchMode"] = o.RequiredRolesMatchMode
	}
	if o.RestrictedRoles != nil {
		toSerialize["restrictedRoles"] = o.RestrictedRoles
	}
	if o.RestrictedRolesMatchMode != nil {
		toSerialize["restrictedRolesMatchMode"] = o.RestrictedRolesMatchMode
	}
	if o.SignatureHash != nil {
		toSerialize["signatureHash"] = o.SignatureHash
	}
	if o.SubjectNameIDPolicy != nil {
		toSerialize["subjectNameIDPolicy"] = o.SubjectNameIDPolicy
	}
	if o.WantAuthnRequestsSigned != nil {
		toSerialize["wantAuthnRequestsSigned"] = o.WantAuthnRequestsSigned
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InternalSaml2ServiceProviderChannelDTO) UnmarshalJSON(bytes []byte) (err error) {
	varInternalSaml2ServiceProviderChannelDTO := _InternalSaml2ServiceProviderChannelDTO{}

	if err = json.Unmarshal(bytes, &varInternalSaml2ServiceProviderChannelDTO); err == nil {
		*o = InternalSaml2ServiceProviderChannelDTO(varInternalSaml2ServiceProviderChannelDTO)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "activeBindings")
		delete(additionalProperties, "activeProfiles")
		delete(additionalProperties, "attributeProfile")
		delete(additionalProperties, "authenticationContract")
		delete(additionalProperties, "authenticationMechanism")
		delete(additionalProperties, "connectionA")
		delete(additionalProperties, "connectionB")
		delete(additionalProperties, "description")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "elementId")
		delete(additionalProperties, "emissionPolicy")
		delete(additionalProperties, "encryptAssertion")
		delete(additionalProperties, "encryptAssertionAlgorithm")
		delete(additionalProperties, "id")
		delete(additionalProperties, "ignoreRequestedNameIDPolicy")
		delete(additionalProperties, "location")
		delete(additionalProperties, "messageTtl")
		delete(additionalProperties, "messageTtlTolerance")
		delete(additionalProperties, "name")
		delete(additionalProperties, "overrideProviderSetup")
		delete(additionalProperties, "requiredRoles")
		delete(additionalProperties, "requiredRolesMatchMode")
		delete(additionalProperties, "restrictedRoles")
		delete(additionalProperties, "restrictedRolesMatchMode")
		delete(additionalProperties, "signatureHash")
		delete(additionalProperties, "subjectNameIDPolicy")
		delete(additionalProperties, "wantAuthnRequestsSigned")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInternalSaml2ServiceProviderChannelDTO struct {
	value *InternalSaml2ServiceProviderChannelDTO
	isSet bool
}

func (v NullableInternalSaml2ServiceProviderChannelDTO) Get() *InternalSaml2ServiceProviderChannelDTO {
	return v.value
}

func (v *NullableInternalSaml2ServiceProviderChannelDTO) Set(val *InternalSaml2ServiceProviderChannelDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalSaml2ServiceProviderChannelDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalSaml2ServiceProviderChannelDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalSaml2ServiceProviderChannelDTO(val *InternalSaml2ServiceProviderChannelDTO) *NullableInternalSaml2ServiceProviderChannelDTO {
	return &NullableInternalSaml2ServiceProviderChannelDTO{value: val, isSet: true}
}

func (v NullableInternalSaml2ServiceProviderChannelDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalSaml2ServiceProviderChannelDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


