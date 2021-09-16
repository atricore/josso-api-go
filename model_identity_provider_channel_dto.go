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

// IdentityProviderChannelDTO struct for IdentityProviderChannelDTO
type IdentityProviderChannelDTO struct {
	AccountLinkagePolicy *AccountLinkagePolicyDTO `json:"accountLinkagePolicy,omitempty"`
	ActiveBindings *[]string `json:"activeBindings,omitempty"`
	ActiveProfiles *[]string `json:"activeProfiles,omitempty"`
	AuthenticationContract *AuthenticationContractDTO `json:"authenticationContract,omitempty"`
	AuthenticationMechanism *AuthenticationMechanismDTO `json:"authenticationMechanism,omitempty"`
	ConnectionA *FederatedConnectionDTO `json:"connectionA,omitempty"`
	ConnectionB *FederatedConnectionDTO `json:"connectionB,omitempty"`
	Description *string `json:"description,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	ElementId *string `json:"elementId,omitempty"`
	EnableProxyExtension *bool `json:"enableProxyExtension,omitempty"`
	Id *int64 `json:"id,omitempty"`
	IdentityMappingPolicy *IdentityMappingPolicyDTO `json:"identityMappingPolicy,omitempty"`
	Location *LocationDTO `json:"location,omitempty"`
	MessageTtl *int32 `json:"messageTtl,omitempty"`
	MessageTtlTolerance *int32 `json:"messageTtlTolerance,omitempty"`
	Name *string `json:"name,omitempty"`
	OverrideProviderSetup *bool `json:"overrideProviderSetup,omitempty"`
	Preferred *bool `json:"preferred,omitempty"`
	SignAuthenticationRequests *bool `json:"signAuthenticationRequests,omitempty"`
	SignatureHash *string `json:"signatureHash,omitempty"`
	WantAssertionSigned *bool `json:"wantAssertionSigned,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentityProviderChannelDTO IdentityProviderChannelDTO

// NewIdentityProviderChannelDTO instantiates a new IdentityProviderChannelDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityProviderChannelDTO() *IdentityProviderChannelDTO {
	this := IdentityProviderChannelDTO{}
	return &this
}

// NewIdentityProviderChannelDTOWithDefaults instantiates a new IdentityProviderChannelDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityProviderChannelDTOWithDefaults() *IdentityProviderChannelDTO {
	this := IdentityProviderChannelDTO{}
	return &this
}

// GetAccountLinkagePolicy returns the AccountLinkagePolicy field value if set, zero value otherwise.
func (o *IdentityProviderChannelDTO) GetAccountLinkagePolicy() AccountLinkagePolicyDTO {
	if o == nil || o.AccountLinkagePolicy == nil {
		var ret AccountLinkagePolicyDTO
		return ret
	}
	return *o.AccountLinkagePolicy
}

// GetAccountLinkagePolicyOk returns a tuple with the AccountLinkagePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderChannelDTO) GetAccountLinkagePolicyOk() (*AccountLinkagePolicyDTO, bool) {
	if o == nil || o.AccountLinkagePolicy == nil {
		return nil, false
	}
	return o.AccountLinkagePolicy, true
}

// HasAccountLinkagePolicy returns a boolean if a field has been set.
func (o *IdentityProviderChannelDTO) HasAccountLinkagePolicy() bool {
	if o != nil && o.AccountLinkagePolicy != nil {
		return true
	}

	return false
}

// SetAccountLinkagePolicy gets a reference to the given AccountLinkagePolicyDTO and assigns it to the AccountLinkagePolicy field.
func (o *IdentityProviderChannelDTO) SetAccountLinkagePolicy(v AccountLinkagePolicyDTO) {
	o.AccountLinkagePolicy = &v
}

// GetActiveBindings returns the ActiveBindings field value if set, zero value otherwise.
func (o *IdentityProviderChannelDTO) GetActiveBindings() []string {
	if o == nil || o.ActiveBindings == nil {
		var ret []string
		return ret
	}
	return *o.ActiveBindings
}

// GetActiveBindingsOk returns a tuple with the ActiveBindings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderChannelDTO) GetActiveBindingsOk() (*[]string, bool) {
	if o == nil || o.ActiveBindings == nil {
		return nil, false
	}
	return o.ActiveBindings, true
}

// HasActiveBindings returns a boolean if a field has been set.
func (o *IdentityProviderChannelDTO) HasActiveBindings() bool {
	if o != nil && o.ActiveBindings != nil {
		return true
	}

	return false
}

// SetActiveBindings gets a reference to the given []string and assigns it to the ActiveBindings field.
func (o *IdentityProviderChannelDTO) SetActiveBindings(v []string) {
	o.ActiveBindings = &v
}

// GetActiveProfiles returns the ActiveProfiles field value if set, zero value otherwise.
func (o *IdentityProviderChannelDTO) GetActiveProfiles() []string {
	if o == nil || o.ActiveProfiles == nil {
		var ret []string
		return ret
	}
	return *o.ActiveProfiles
}

// GetActiveProfilesOk returns a tuple with the ActiveProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderChannelDTO) GetActiveProfilesOk() (*[]string, bool) {
	if o == nil || o.ActiveProfiles == nil {
		return nil, false
	}
	return o.ActiveProfiles, true
}

// HasActiveProfiles returns a boolean if a field has been set.
func (o *IdentityProviderChannelDTO) HasActiveProfiles() bool {
	if o != nil && o.ActiveProfiles != nil {
		return true
	}

	return false
}

// SetActiveProfiles gets a reference to the given []string and assigns it to the ActiveProfiles field.
func (o *IdentityProviderChannelDTO) SetActiveProfiles(v []string) {
	o.ActiveProfiles = &v
}

// GetAuthenticationContract returns the AuthenticationContract field value if set, zero value otherwise.
func (o *IdentityProviderChannelDTO) GetAuthenticationContract() AuthenticationContractDTO {
	if o == nil || o.AuthenticationContract == nil {
		var ret AuthenticationContractDTO
		return ret
	}
	return *o.AuthenticationContract
}

// GetAuthenticationContractOk returns a tuple with the AuthenticationContract field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderChannelDTO) GetAuthenticationContractOk() (*AuthenticationContractDTO, bool) {
	if o == nil || o.AuthenticationContract == nil {
		return nil, false
	}
	return o.AuthenticationContract, true
}

// HasAuthenticationContract returns a boolean if a field has been set.
func (o *IdentityProviderChannelDTO) HasAuthenticationContract() bool {
	if o != nil && o.AuthenticationContract != nil {
		return true
	}

	return false
}

// SetAuthenticationContract gets a reference to the given AuthenticationContractDTO and assigns it to the AuthenticationContract field.
func (o *IdentityProviderChannelDTO) SetAuthenticationContract(v AuthenticationContractDTO) {
	o.AuthenticationContract = &v
}

// GetAuthenticationMechanism returns the AuthenticationMechanism field value if set, zero value otherwise.
func (o *IdentityProviderChannelDTO) GetAuthenticationMechanism() AuthenticationMechanismDTO {
	if o == nil || o.AuthenticationMechanism == nil {
		var ret AuthenticationMechanismDTO
		return ret
	}
	return *o.AuthenticationMechanism
}

// GetAuthenticationMechanismOk returns a tuple with the AuthenticationMechanism field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderChannelDTO) GetAuthenticationMechanismOk() (*AuthenticationMechanismDTO, bool) {
	if o == nil || o.AuthenticationMechanism == nil {
		return nil, false
	}
	return o.AuthenticationMechanism, true
}

// HasAuthenticationMechanism returns a boolean if a field has been set.
func (o *IdentityProviderChannelDTO) HasAuthenticationMechanism() bool {
	if o != nil && o.AuthenticationMechanism != nil {
		return true
	}

	return false
}

// SetAuthenticationMechanism gets a reference to the given AuthenticationMechanismDTO and assigns it to the AuthenticationMechanism field.
func (o *IdentityProviderChannelDTO) SetAuthenticationMechanism(v AuthenticationMechanismDTO) {
	o.AuthenticationMechanism = &v
}

// GetConnectionA returns the ConnectionA field value if set, zero value otherwise.
func (o *IdentityProviderChannelDTO) GetConnectionA() FederatedConnectionDTO {
	if o == nil || o.ConnectionA == nil {
		var ret FederatedConnectionDTO
		return ret
	}
	return *o.ConnectionA
}

// GetConnectionAOk returns a tuple with the ConnectionA field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderChannelDTO) GetConnectionAOk() (*FederatedConnectionDTO, bool) {
	if o == nil || o.ConnectionA == nil {
		return nil, false
	}
	return o.ConnectionA, true
}

// HasConnectionA returns a boolean if a field has been set.
func (o *IdentityProviderChannelDTO) HasConnectionA() bool {
	if o != nil && o.ConnectionA != nil {
		return true
	}

	return false
}

// SetConnectionA gets a reference to the given FederatedConnectionDTO and assigns it to the ConnectionA field.
func (o *IdentityProviderChannelDTO) SetConnectionA(v FederatedConnectionDTO) {
	o.ConnectionA = &v
}

// GetConnectionB returns the ConnectionB field value if set, zero value otherwise.
func (o *IdentityProviderChannelDTO) GetConnectionB() FederatedConnectionDTO {
	if o == nil || o.ConnectionB == nil {
		var ret FederatedConnectionDTO
		return ret
	}
	return *o.ConnectionB
}

// GetConnectionBOk returns a tuple with the ConnectionB field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderChannelDTO) GetConnectionBOk() (*FederatedConnectionDTO, bool) {
	if o == nil || o.ConnectionB == nil {
		return nil, false
	}
	return o.ConnectionB, true
}

// HasConnectionB returns a boolean if a field has been set.
func (o *IdentityProviderChannelDTO) HasConnectionB() bool {
	if o != nil && o.ConnectionB != nil {
		return true
	}

	return false
}

// SetConnectionB gets a reference to the given FederatedConnectionDTO and assigns it to the ConnectionB field.
func (o *IdentityProviderChannelDTO) SetConnectionB(v FederatedConnectionDTO) {
	o.ConnectionB = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IdentityProviderChannelDTO) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderChannelDTO) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IdentityProviderChannelDTO) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IdentityProviderChannelDTO) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *IdentityProviderChannelDTO) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderChannelDTO) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *IdentityProviderChannelDTO) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *IdentityProviderChannelDTO) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *IdentityProviderChannelDTO) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderChannelDTO) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *IdentityProviderChannelDTO) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *IdentityProviderChannelDTO) SetElementId(v string) {
	o.ElementId = &v
}

// GetEnableProxyExtension returns the EnableProxyExtension field value if set, zero value otherwise.
func (o *IdentityProviderChannelDTO) GetEnableProxyExtension() bool {
	if o == nil || o.EnableProxyExtension == nil {
		var ret bool
		return ret
	}
	return *o.EnableProxyExtension
}

// GetEnableProxyExtensionOk returns a tuple with the EnableProxyExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderChannelDTO) GetEnableProxyExtensionOk() (*bool, bool) {
	if o == nil || o.EnableProxyExtension == nil {
		return nil, false
	}
	return o.EnableProxyExtension, true
}

// HasEnableProxyExtension returns a boolean if a field has been set.
func (o *IdentityProviderChannelDTO) HasEnableProxyExtension() bool {
	if o != nil && o.EnableProxyExtension != nil {
		return true
	}

	return false
}

// SetEnableProxyExtension gets a reference to the given bool and assigns it to the EnableProxyExtension field.
func (o *IdentityProviderChannelDTO) SetEnableProxyExtension(v bool) {
	o.EnableProxyExtension = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IdentityProviderChannelDTO) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderChannelDTO) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IdentityProviderChannelDTO) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *IdentityProviderChannelDTO) SetId(v int64) {
	o.Id = &v
}

// GetIdentityMappingPolicy returns the IdentityMappingPolicy field value if set, zero value otherwise.
func (o *IdentityProviderChannelDTO) GetIdentityMappingPolicy() IdentityMappingPolicyDTO {
	if o == nil || o.IdentityMappingPolicy == nil {
		var ret IdentityMappingPolicyDTO
		return ret
	}
	return *o.IdentityMappingPolicy
}

// GetIdentityMappingPolicyOk returns a tuple with the IdentityMappingPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderChannelDTO) GetIdentityMappingPolicyOk() (*IdentityMappingPolicyDTO, bool) {
	if o == nil || o.IdentityMappingPolicy == nil {
		return nil, false
	}
	return o.IdentityMappingPolicy, true
}

// HasIdentityMappingPolicy returns a boolean if a field has been set.
func (o *IdentityProviderChannelDTO) HasIdentityMappingPolicy() bool {
	if o != nil && o.IdentityMappingPolicy != nil {
		return true
	}

	return false
}

// SetIdentityMappingPolicy gets a reference to the given IdentityMappingPolicyDTO and assigns it to the IdentityMappingPolicy field.
func (o *IdentityProviderChannelDTO) SetIdentityMappingPolicy(v IdentityMappingPolicyDTO) {
	o.IdentityMappingPolicy = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *IdentityProviderChannelDTO) GetLocation() LocationDTO {
	if o == nil || o.Location == nil {
		var ret LocationDTO
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderChannelDTO) GetLocationOk() (*LocationDTO, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *IdentityProviderChannelDTO) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given LocationDTO and assigns it to the Location field.
func (o *IdentityProviderChannelDTO) SetLocation(v LocationDTO) {
	o.Location = &v
}

// GetMessageTtl returns the MessageTtl field value if set, zero value otherwise.
func (o *IdentityProviderChannelDTO) GetMessageTtl() int32 {
	if o == nil || o.MessageTtl == nil {
		var ret int32
		return ret
	}
	return *o.MessageTtl
}

// GetMessageTtlOk returns a tuple with the MessageTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderChannelDTO) GetMessageTtlOk() (*int32, bool) {
	if o == nil || o.MessageTtl == nil {
		return nil, false
	}
	return o.MessageTtl, true
}

// HasMessageTtl returns a boolean if a field has been set.
func (o *IdentityProviderChannelDTO) HasMessageTtl() bool {
	if o != nil && o.MessageTtl != nil {
		return true
	}

	return false
}

// SetMessageTtl gets a reference to the given int32 and assigns it to the MessageTtl field.
func (o *IdentityProviderChannelDTO) SetMessageTtl(v int32) {
	o.MessageTtl = &v
}

// GetMessageTtlTolerance returns the MessageTtlTolerance field value if set, zero value otherwise.
func (o *IdentityProviderChannelDTO) GetMessageTtlTolerance() int32 {
	if o == nil || o.MessageTtlTolerance == nil {
		var ret int32
		return ret
	}
	return *o.MessageTtlTolerance
}

// GetMessageTtlToleranceOk returns a tuple with the MessageTtlTolerance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderChannelDTO) GetMessageTtlToleranceOk() (*int32, bool) {
	if o == nil || o.MessageTtlTolerance == nil {
		return nil, false
	}
	return o.MessageTtlTolerance, true
}

// HasMessageTtlTolerance returns a boolean if a field has been set.
func (o *IdentityProviderChannelDTO) HasMessageTtlTolerance() bool {
	if o != nil && o.MessageTtlTolerance != nil {
		return true
	}

	return false
}

// SetMessageTtlTolerance gets a reference to the given int32 and assigns it to the MessageTtlTolerance field.
func (o *IdentityProviderChannelDTO) SetMessageTtlTolerance(v int32) {
	o.MessageTtlTolerance = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IdentityProviderChannelDTO) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderChannelDTO) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IdentityProviderChannelDTO) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IdentityProviderChannelDTO) SetName(v string) {
	o.Name = &v
}

// GetOverrideProviderSetup returns the OverrideProviderSetup field value if set, zero value otherwise.
func (o *IdentityProviderChannelDTO) GetOverrideProviderSetup() bool {
	if o == nil || o.OverrideProviderSetup == nil {
		var ret bool
		return ret
	}
	return *o.OverrideProviderSetup
}

// GetOverrideProviderSetupOk returns a tuple with the OverrideProviderSetup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderChannelDTO) GetOverrideProviderSetupOk() (*bool, bool) {
	if o == nil || o.OverrideProviderSetup == nil {
		return nil, false
	}
	return o.OverrideProviderSetup, true
}

// HasOverrideProviderSetup returns a boolean if a field has been set.
func (o *IdentityProviderChannelDTO) HasOverrideProviderSetup() bool {
	if o != nil && o.OverrideProviderSetup != nil {
		return true
	}

	return false
}

// SetOverrideProviderSetup gets a reference to the given bool and assigns it to the OverrideProviderSetup field.
func (o *IdentityProviderChannelDTO) SetOverrideProviderSetup(v bool) {
	o.OverrideProviderSetup = &v
}

// GetPreferred returns the Preferred field value if set, zero value otherwise.
func (o *IdentityProviderChannelDTO) GetPreferred() bool {
	if o == nil || o.Preferred == nil {
		var ret bool
		return ret
	}
	return *o.Preferred
}

// GetPreferredOk returns a tuple with the Preferred field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderChannelDTO) GetPreferredOk() (*bool, bool) {
	if o == nil || o.Preferred == nil {
		return nil, false
	}
	return o.Preferred, true
}

// HasPreferred returns a boolean if a field has been set.
func (o *IdentityProviderChannelDTO) HasPreferred() bool {
	if o != nil && o.Preferred != nil {
		return true
	}

	return false
}

// SetPreferred gets a reference to the given bool and assigns it to the Preferred field.
func (o *IdentityProviderChannelDTO) SetPreferred(v bool) {
	o.Preferred = &v
}

// GetSignAuthenticationRequests returns the SignAuthenticationRequests field value if set, zero value otherwise.
func (o *IdentityProviderChannelDTO) GetSignAuthenticationRequests() bool {
	if o == nil || o.SignAuthenticationRequests == nil {
		var ret bool
		return ret
	}
	return *o.SignAuthenticationRequests
}

// GetSignAuthenticationRequestsOk returns a tuple with the SignAuthenticationRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderChannelDTO) GetSignAuthenticationRequestsOk() (*bool, bool) {
	if o == nil || o.SignAuthenticationRequests == nil {
		return nil, false
	}
	return o.SignAuthenticationRequests, true
}

// HasSignAuthenticationRequests returns a boolean if a field has been set.
func (o *IdentityProviderChannelDTO) HasSignAuthenticationRequests() bool {
	if o != nil && o.SignAuthenticationRequests != nil {
		return true
	}

	return false
}

// SetSignAuthenticationRequests gets a reference to the given bool and assigns it to the SignAuthenticationRequests field.
func (o *IdentityProviderChannelDTO) SetSignAuthenticationRequests(v bool) {
	o.SignAuthenticationRequests = &v
}

// GetSignatureHash returns the SignatureHash field value if set, zero value otherwise.
func (o *IdentityProviderChannelDTO) GetSignatureHash() string {
	if o == nil || o.SignatureHash == nil {
		var ret string
		return ret
	}
	return *o.SignatureHash
}

// GetSignatureHashOk returns a tuple with the SignatureHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderChannelDTO) GetSignatureHashOk() (*string, bool) {
	if o == nil || o.SignatureHash == nil {
		return nil, false
	}
	return o.SignatureHash, true
}

// HasSignatureHash returns a boolean if a field has been set.
func (o *IdentityProviderChannelDTO) HasSignatureHash() bool {
	if o != nil && o.SignatureHash != nil {
		return true
	}

	return false
}

// SetSignatureHash gets a reference to the given string and assigns it to the SignatureHash field.
func (o *IdentityProviderChannelDTO) SetSignatureHash(v string) {
	o.SignatureHash = &v
}

// GetWantAssertionSigned returns the WantAssertionSigned field value if set, zero value otherwise.
func (o *IdentityProviderChannelDTO) GetWantAssertionSigned() bool {
	if o == nil || o.WantAssertionSigned == nil {
		var ret bool
		return ret
	}
	return *o.WantAssertionSigned
}

// GetWantAssertionSignedOk returns a tuple with the WantAssertionSigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderChannelDTO) GetWantAssertionSignedOk() (*bool, bool) {
	if o == nil || o.WantAssertionSigned == nil {
		return nil, false
	}
	return o.WantAssertionSigned, true
}

// HasWantAssertionSigned returns a boolean if a field has been set.
func (o *IdentityProviderChannelDTO) HasWantAssertionSigned() bool {
	if o != nil && o.WantAssertionSigned != nil {
		return true
	}

	return false
}

// SetWantAssertionSigned gets a reference to the given bool and assigns it to the WantAssertionSigned field.
func (o *IdentityProviderChannelDTO) SetWantAssertionSigned(v bool) {
	o.WantAssertionSigned = &v
}

func (o IdentityProviderChannelDTO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountLinkagePolicy != nil {
		toSerialize["accountLinkagePolicy"] = o.AccountLinkagePolicy
	}
	if o.ActiveBindings != nil {
		toSerialize["activeBindings"] = o.ActiveBindings
	}
	if o.ActiveProfiles != nil {
		toSerialize["activeProfiles"] = o.ActiveProfiles
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
	if o.EnableProxyExtension != nil {
		toSerialize["enableProxyExtension"] = o.EnableProxyExtension
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IdentityMappingPolicy != nil {
		toSerialize["identityMappingPolicy"] = o.IdentityMappingPolicy
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
	if o.Preferred != nil {
		toSerialize["preferred"] = o.Preferred
	}
	if o.SignAuthenticationRequests != nil {
		toSerialize["signAuthenticationRequests"] = o.SignAuthenticationRequests
	}
	if o.SignatureHash != nil {
		toSerialize["signatureHash"] = o.SignatureHash
	}
	if o.WantAssertionSigned != nil {
		toSerialize["wantAssertionSigned"] = o.WantAssertionSigned
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IdentityProviderChannelDTO) UnmarshalJSON(bytes []byte) (err error) {
	varIdentityProviderChannelDTO := _IdentityProviderChannelDTO{}

	if err = json.Unmarshal(bytes, &varIdentityProviderChannelDTO); err == nil {
		*o = IdentityProviderChannelDTO(varIdentityProviderChannelDTO)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "accountLinkagePolicy")
		delete(additionalProperties, "activeBindings")
		delete(additionalProperties, "activeProfiles")
		delete(additionalProperties, "authenticationContract")
		delete(additionalProperties, "authenticationMechanism")
		delete(additionalProperties, "connectionA")
		delete(additionalProperties, "connectionB")
		delete(additionalProperties, "description")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "elementId")
		delete(additionalProperties, "enableProxyExtension")
		delete(additionalProperties, "id")
		delete(additionalProperties, "identityMappingPolicy")
		delete(additionalProperties, "location")
		delete(additionalProperties, "messageTtl")
		delete(additionalProperties, "messageTtlTolerance")
		delete(additionalProperties, "name")
		delete(additionalProperties, "overrideProviderSetup")
		delete(additionalProperties, "preferred")
		delete(additionalProperties, "signAuthenticationRequests")
		delete(additionalProperties, "signatureHash")
		delete(additionalProperties, "wantAssertionSigned")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentityProviderChannelDTO struct {
	value *IdentityProviderChannelDTO
	isSet bool
}

func (v NullableIdentityProviderChannelDTO) Get() *IdentityProviderChannelDTO {
	return v.value
}

func (v *NullableIdentityProviderChannelDTO) Set(val *IdentityProviderChannelDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityProviderChannelDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityProviderChannelDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityProviderChannelDTO(val *IdentityProviderChannelDTO) *NullableIdentityProviderChannelDTO {
	return &NullableIdentityProviderChannelDTO{value: val, isSet: true}
}

func (v NullableIdentityProviderChannelDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityProviderChannelDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

