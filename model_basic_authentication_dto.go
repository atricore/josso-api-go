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

// BasicAuthenticationDTO struct for BasicAuthenticationDTO
type BasicAuthenticationDTO struct {
	DelegatedAuthentication *DelegatedAuthenticationDTO `json:"delegatedAuthentication,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	ElementId *string `json:"elementId,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	HashAlgorithm *string `json:"hashAlgorithm,omitempty"`
	HashEncoding *string `json:"hashEncoding,omitempty"`
	Id *int64 `json:"id,omitempty"`
	IgnorePasswordCase *bool `json:"ignorePasswordCase,omitempty"`
	IgnoreUsernameCase *bool `json:"ignoreUsernameCase,omitempty"`
	ImpersonateUserPolicy *ImpersonateUserPolicyDTO `json:"impersonateUserPolicy,omitempty"`
	Name *string `json:"name,omitempty"`
	Priority *int32 `json:"priority,omitempty"`
	SaltLength *int32 `json:"saltLength,omitempty"`
	SaltPrefix *string `json:"saltPrefix,omitempty"`
	SaltSuffix *string `json:"saltSuffix,omitempty"`
	SimpleAuthnSaml2AuthnCtxClass *string `json:"simpleAuthnSaml2AuthnCtxClass,omitempty"`
	X *float64 `json:"x,omitempty"`
	Y *float64 `json:"y,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BasicAuthenticationDTO BasicAuthenticationDTO

// NewBasicAuthenticationDTO instantiates a new BasicAuthenticationDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBasicAuthenticationDTO() *BasicAuthenticationDTO {
	this := BasicAuthenticationDTO{}
	return &this
}

// NewBasicAuthenticationDTOWithDefaults instantiates a new BasicAuthenticationDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBasicAuthenticationDTOWithDefaults() *BasicAuthenticationDTO {
	this := BasicAuthenticationDTO{}
	return &this
}

// GetDelegatedAuthentication returns the DelegatedAuthentication field value if set, zero value otherwise.
func (o *BasicAuthenticationDTO) GetDelegatedAuthentication() DelegatedAuthenticationDTO {
	if o == nil || o.DelegatedAuthentication == nil {
		var ret DelegatedAuthenticationDTO
		return ret
	}
	return *o.DelegatedAuthentication
}

// GetDelegatedAuthenticationOk returns a tuple with the DelegatedAuthentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicAuthenticationDTO) GetDelegatedAuthenticationOk() (*DelegatedAuthenticationDTO, bool) {
	if o == nil || o.DelegatedAuthentication == nil {
		return nil, false
	}
	return o.DelegatedAuthentication, true
}

// HasDelegatedAuthentication returns a boolean if a field has been set.
func (o *BasicAuthenticationDTO) HasDelegatedAuthentication() bool {
	if o != nil && o.DelegatedAuthentication != nil {
		return true
	}

	return false
}

// SetDelegatedAuthentication gets a reference to the given DelegatedAuthenticationDTO and assigns it to the DelegatedAuthentication field.
func (o *BasicAuthenticationDTO) SetDelegatedAuthentication(v DelegatedAuthenticationDTO) {
	o.DelegatedAuthentication = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *BasicAuthenticationDTO) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicAuthenticationDTO) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *BasicAuthenticationDTO) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *BasicAuthenticationDTO) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *BasicAuthenticationDTO) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicAuthenticationDTO) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *BasicAuthenticationDTO) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *BasicAuthenticationDTO) SetElementId(v string) {
	o.ElementId = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *BasicAuthenticationDTO) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicAuthenticationDTO) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *BasicAuthenticationDTO) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *BasicAuthenticationDTO) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetHashAlgorithm returns the HashAlgorithm field value if set, zero value otherwise.
func (o *BasicAuthenticationDTO) GetHashAlgorithm() string {
	if o == nil || o.HashAlgorithm == nil {
		var ret string
		return ret
	}
	return *o.HashAlgorithm
}

// GetHashAlgorithmOk returns a tuple with the HashAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicAuthenticationDTO) GetHashAlgorithmOk() (*string, bool) {
	if o == nil || o.HashAlgorithm == nil {
		return nil, false
	}
	return o.HashAlgorithm, true
}

// HasHashAlgorithm returns a boolean if a field has been set.
func (o *BasicAuthenticationDTO) HasHashAlgorithm() bool {
	if o != nil && o.HashAlgorithm != nil {
		return true
	}

	return false
}

// SetHashAlgorithm gets a reference to the given string and assigns it to the HashAlgorithm field.
func (o *BasicAuthenticationDTO) SetHashAlgorithm(v string) {
	o.HashAlgorithm = &v
}

// GetHashEncoding returns the HashEncoding field value if set, zero value otherwise.
func (o *BasicAuthenticationDTO) GetHashEncoding() string {
	if o == nil || o.HashEncoding == nil {
		var ret string
		return ret
	}
	return *o.HashEncoding
}

// GetHashEncodingOk returns a tuple with the HashEncoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicAuthenticationDTO) GetHashEncodingOk() (*string, bool) {
	if o == nil || o.HashEncoding == nil {
		return nil, false
	}
	return o.HashEncoding, true
}

// HasHashEncoding returns a boolean if a field has been set.
func (o *BasicAuthenticationDTO) HasHashEncoding() bool {
	if o != nil && o.HashEncoding != nil {
		return true
	}

	return false
}

// SetHashEncoding gets a reference to the given string and assigns it to the HashEncoding field.
func (o *BasicAuthenticationDTO) SetHashEncoding(v string) {
	o.HashEncoding = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BasicAuthenticationDTO) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicAuthenticationDTO) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BasicAuthenticationDTO) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *BasicAuthenticationDTO) SetId(v int64) {
	o.Id = &v
}

// GetIgnorePasswordCase returns the IgnorePasswordCase field value if set, zero value otherwise.
func (o *BasicAuthenticationDTO) GetIgnorePasswordCase() bool {
	if o == nil || o.IgnorePasswordCase == nil {
		var ret bool
		return ret
	}
	return *o.IgnorePasswordCase
}

// GetIgnorePasswordCaseOk returns a tuple with the IgnorePasswordCase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicAuthenticationDTO) GetIgnorePasswordCaseOk() (*bool, bool) {
	if o == nil || o.IgnorePasswordCase == nil {
		return nil, false
	}
	return o.IgnorePasswordCase, true
}

// HasIgnorePasswordCase returns a boolean if a field has been set.
func (o *BasicAuthenticationDTO) HasIgnorePasswordCase() bool {
	if o != nil && o.IgnorePasswordCase != nil {
		return true
	}

	return false
}

// SetIgnorePasswordCase gets a reference to the given bool and assigns it to the IgnorePasswordCase field.
func (o *BasicAuthenticationDTO) SetIgnorePasswordCase(v bool) {
	o.IgnorePasswordCase = &v
}

// GetIgnoreUsernameCase returns the IgnoreUsernameCase field value if set, zero value otherwise.
func (o *BasicAuthenticationDTO) GetIgnoreUsernameCase() bool {
	if o == nil || o.IgnoreUsernameCase == nil {
		var ret bool
		return ret
	}
	return *o.IgnoreUsernameCase
}

// GetIgnoreUsernameCaseOk returns a tuple with the IgnoreUsernameCase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicAuthenticationDTO) GetIgnoreUsernameCaseOk() (*bool, bool) {
	if o == nil || o.IgnoreUsernameCase == nil {
		return nil, false
	}
	return o.IgnoreUsernameCase, true
}

// HasIgnoreUsernameCase returns a boolean if a field has been set.
func (o *BasicAuthenticationDTO) HasIgnoreUsernameCase() bool {
	if o != nil && o.IgnoreUsernameCase != nil {
		return true
	}

	return false
}

// SetIgnoreUsernameCase gets a reference to the given bool and assigns it to the IgnoreUsernameCase field.
func (o *BasicAuthenticationDTO) SetIgnoreUsernameCase(v bool) {
	o.IgnoreUsernameCase = &v
}

// GetImpersonateUserPolicy returns the ImpersonateUserPolicy field value if set, zero value otherwise.
func (o *BasicAuthenticationDTO) GetImpersonateUserPolicy() ImpersonateUserPolicyDTO {
	if o == nil || o.ImpersonateUserPolicy == nil {
		var ret ImpersonateUserPolicyDTO
		return ret
	}
	return *o.ImpersonateUserPolicy
}

// GetImpersonateUserPolicyOk returns a tuple with the ImpersonateUserPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicAuthenticationDTO) GetImpersonateUserPolicyOk() (*ImpersonateUserPolicyDTO, bool) {
	if o == nil || o.ImpersonateUserPolicy == nil {
		return nil, false
	}
	return o.ImpersonateUserPolicy, true
}

// HasImpersonateUserPolicy returns a boolean if a field has been set.
func (o *BasicAuthenticationDTO) HasImpersonateUserPolicy() bool {
	if o != nil && o.ImpersonateUserPolicy != nil {
		return true
	}

	return false
}

// SetImpersonateUserPolicy gets a reference to the given ImpersonateUserPolicyDTO and assigns it to the ImpersonateUserPolicy field.
func (o *BasicAuthenticationDTO) SetImpersonateUserPolicy(v ImpersonateUserPolicyDTO) {
	o.ImpersonateUserPolicy = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BasicAuthenticationDTO) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicAuthenticationDTO) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BasicAuthenticationDTO) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BasicAuthenticationDTO) SetName(v string) {
	o.Name = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *BasicAuthenticationDTO) GetPriority() int32 {
	if o == nil || o.Priority == nil {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicAuthenticationDTO) GetPriorityOk() (*int32, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *BasicAuthenticationDTO) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *BasicAuthenticationDTO) SetPriority(v int32) {
	o.Priority = &v
}

// GetSaltLength returns the SaltLength field value if set, zero value otherwise.
func (o *BasicAuthenticationDTO) GetSaltLength() int32 {
	if o == nil || o.SaltLength == nil {
		var ret int32
		return ret
	}
	return *o.SaltLength
}

// GetSaltLengthOk returns a tuple with the SaltLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicAuthenticationDTO) GetSaltLengthOk() (*int32, bool) {
	if o == nil || o.SaltLength == nil {
		return nil, false
	}
	return o.SaltLength, true
}

// HasSaltLength returns a boolean if a field has been set.
func (o *BasicAuthenticationDTO) HasSaltLength() bool {
	if o != nil && o.SaltLength != nil {
		return true
	}

	return false
}

// SetSaltLength gets a reference to the given int32 and assigns it to the SaltLength field.
func (o *BasicAuthenticationDTO) SetSaltLength(v int32) {
	o.SaltLength = &v
}

// GetSaltPrefix returns the SaltPrefix field value if set, zero value otherwise.
func (o *BasicAuthenticationDTO) GetSaltPrefix() string {
	if o == nil || o.SaltPrefix == nil {
		var ret string
		return ret
	}
	return *o.SaltPrefix
}

// GetSaltPrefixOk returns a tuple with the SaltPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicAuthenticationDTO) GetSaltPrefixOk() (*string, bool) {
	if o == nil || o.SaltPrefix == nil {
		return nil, false
	}
	return o.SaltPrefix, true
}

// HasSaltPrefix returns a boolean if a field has been set.
func (o *BasicAuthenticationDTO) HasSaltPrefix() bool {
	if o != nil && o.SaltPrefix != nil {
		return true
	}

	return false
}

// SetSaltPrefix gets a reference to the given string and assigns it to the SaltPrefix field.
func (o *BasicAuthenticationDTO) SetSaltPrefix(v string) {
	o.SaltPrefix = &v
}

// GetSaltSuffix returns the SaltSuffix field value if set, zero value otherwise.
func (o *BasicAuthenticationDTO) GetSaltSuffix() string {
	if o == nil || o.SaltSuffix == nil {
		var ret string
		return ret
	}
	return *o.SaltSuffix
}

// GetSaltSuffixOk returns a tuple with the SaltSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicAuthenticationDTO) GetSaltSuffixOk() (*string, bool) {
	if o == nil || o.SaltSuffix == nil {
		return nil, false
	}
	return o.SaltSuffix, true
}

// HasSaltSuffix returns a boolean if a field has been set.
func (o *BasicAuthenticationDTO) HasSaltSuffix() bool {
	if o != nil && o.SaltSuffix != nil {
		return true
	}

	return false
}

// SetSaltSuffix gets a reference to the given string and assigns it to the SaltSuffix field.
func (o *BasicAuthenticationDTO) SetSaltSuffix(v string) {
	o.SaltSuffix = &v
}

// GetSimpleAuthnSaml2AuthnCtxClass returns the SimpleAuthnSaml2AuthnCtxClass field value if set, zero value otherwise.
func (o *BasicAuthenticationDTO) GetSimpleAuthnSaml2AuthnCtxClass() string {
	if o == nil || o.SimpleAuthnSaml2AuthnCtxClass == nil {
		var ret string
		return ret
	}
	return *o.SimpleAuthnSaml2AuthnCtxClass
}

// GetSimpleAuthnSaml2AuthnCtxClassOk returns a tuple with the SimpleAuthnSaml2AuthnCtxClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicAuthenticationDTO) GetSimpleAuthnSaml2AuthnCtxClassOk() (*string, bool) {
	if o == nil || o.SimpleAuthnSaml2AuthnCtxClass == nil {
		return nil, false
	}
	return o.SimpleAuthnSaml2AuthnCtxClass, true
}

// HasSimpleAuthnSaml2AuthnCtxClass returns a boolean if a field has been set.
func (o *BasicAuthenticationDTO) HasSimpleAuthnSaml2AuthnCtxClass() bool {
	if o != nil && o.SimpleAuthnSaml2AuthnCtxClass != nil {
		return true
	}

	return false
}

// SetSimpleAuthnSaml2AuthnCtxClass gets a reference to the given string and assigns it to the SimpleAuthnSaml2AuthnCtxClass field.
func (o *BasicAuthenticationDTO) SetSimpleAuthnSaml2AuthnCtxClass(v string) {
	o.SimpleAuthnSaml2AuthnCtxClass = &v
}

// GetX returns the X field value if set, zero value otherwise.
func (o *BasicAuthenticationDTO) GetX() float64 {
	if o == nil || o.X == nil {
		var ret float64
		return ret
	}
	return *o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicAuthenticationDTO) GetXOk() (*float64, bool) {
	if o == nil || o.X == nil {
		return nil, false
	}
	return o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *BasicAuthenticationDTO) HasX() bool {
	if o != nil && o.X != nil {
		return true
	}

	return false
}

// SetX gets a reference to the given float64 and assigns it to the X field.
func (o *BasicAuthenticationDTO) SetX(v float64) {
	o.X = &v
}

// GetY returns the Y field value if set, zero value otherwise.
func (o *BasicAuthenticationDTO) GetY() float64 {
	if o == nil || o.Y == nil {
		var ret float64
		return ret
	}
	return *o.Y
}

// GetYOk returns a tuple with the Y field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicAuthenticationDTO) GetYOk() (*float64, bool) {
	if o == nil || o.Y == nil {
		return nil, false
	}
	return o.Y, true
}

// HasY returns a boolean if a field has been set.
func (o *BasicAuthenticationDTO) HasY() bool {
	if o != nil && o.Y != nil {
		return true
	}

	return false
}

// SetY gets a reference to the given float64 and assigns it to the Y field.
func (o *BasicAuthenticationDTO) SetY(v float64) {
	o.Y = &v
}

func (o BasicAuthenticationDTO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DelegatedAuthentication != nil {
		toSerialize["delegatedAuthentication"] = o.DelegatedAuthentication
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.ElementId != nil {
		toSerialize["elementId"] = o.ElementId
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.HashAlgorithm != nil {
		toSerialize["hashAlgorithm"] = o.HashAlgorithm
	}
	if o.HashEncoding != nil {
		toSerialize["hashEncoding"] = o.HashEncoding
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IgnorePasswordCase != nil {
		toSerialize["ignorePasswordCase"] = o.IgnorePasswordCase
	}
	if o.IgnoreUsernameCase != nil {
		toSerialize["ignoreUsernameCase"] = o.IgnoreUsernameCase
	}
	if o.ImpersonateUserPolicy != nil {
		toSerialize["impersonateUserPolicy"] = o.ImpersonateUserPolicy
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	if o.SaltLength != nil {
		toSerialize["saltLength"] = o.SaltLength
	}
	if o.SaltPrefix != nil {
		toSerialize["saltPrefix"] = o.SaltPrefix
	}
	if o.SaltSuffix != nil {
		toSerialize["saltSuffix"] = o.SaltSuffix
	}
	if o.SimpleAuthnSaml2AuthnCtxClass != nil {
		toSerialize["simpleAuthnSaml2AuthnCtxClass"] = o.SimpleAuthnSaml2AuthnCtxClass
	}
	if o.X != nil {
		toSerialize["x"] = o.X
	}
	if o.Y != nil {
		toSerialize["y"] = o.Y
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BasicAuthenticationDTO) UnmarshalJSON(bytes []byte) (err error) {
	varBasicAuthenticationDTO := _BasicAuthenticationDTO{}

	if err = json.Unmarshal(bytes, &varBasicAuthenticationDTO); err == nil {
		*o = BasicAuthenticationDTO(varBasicAuthenticationDTO)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "delegatedAuthentication")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "elementId")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "hashAlgorithm")
		delete(additionalProperties, "hashEncoding")
		delete(additionalProperties, "id")
		delete(additionalProperties, "ignorePasswordCase")
		delete(additionalProperties, "ignoreUsernameCase")
		delete(additionalProperties, "impersonateUserPolicy")
		delete(additionalProperties, "name")
		delete(additionalProperties, "priority")
		delete(additionalProperties, "saltLength")
		delete(additionalProperties, "saltPrefix")
		delete(additionalProperties, "saltSuffix")
		delete(additionalProperties, "simpleAuthnSaml2AuthnCtxClass")
		delete(additionalProperties, "x")
		delete(additionalProperties, "y")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBasicAuthenticationDTO struct {
	value *BasicAuthenticationDTO
	isSet bool
}

func (v NullableBasicAuthenticationDTO) Get() *BasicAuthenticationDTO {
	return v.value
}

func (v *NullableBasicAuthenticationDTO) Set(val *BasicAuthenticationDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableBasicAuthenticationDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableBasicAuthenticationDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBasicAuthenticationDTO(val *BasicAuthenticationDTO) *NullableBasicAuthenticationDTO {
	return &NullableBasicAuthenticationDTO{value: val, isSet: true}
}

func (v NullableBasicAuthenticationDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBasicAuthenticationDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


