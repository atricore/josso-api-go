/*
Atricore Console :: Remote : API

# Atricore Console API

API version: 1.5.1-SNAPSHOT
Contact: sgonzalez@atricore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package jossoappi

import (
	"encoding/json"
)

// StoreIdpGoogleReq struct for StoreIdpGoogleReq
type StoreIdpGoogleReq struct {
	IdOrName *string `json:"idOrName,omitempty"`
	Idp *GoogleOpenIDConnectIdentityProviderDTO `json:"idp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StoreIdpGoogleReq StoreIdpGoogleReq

// NewStoreIdpGoogleReq instantiates a new StoreIdpGoogleReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoreIdpGoogleReq() *StoreIdpGoogleReq {
	this := StoreIdpGoogleReq{}
	return &this
}

// NewStoreIdpGoogleReqWithDefaults instantiates a new StoreIdpGoogleReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoreIdpGoogleReqWithDefaults() *StoreIdpGoogleReq {
	this := StoreIdpGoogleReq{}
	return &this
}

// GetIdOrName returns the IdOrName field value if set, zero value otherwise.
func (o *StoreIdpGoogleReq) GetIdOrName() string {
	if o == nil || isNil(o.IdOrName) {
		var ret string
		return ret
	}
	return *o.IdOrName
}

// GetIdOrNameOk returns a tuple with the IdOrName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreIdpGoogleReq) GetIdOrNameOk() (*string, bool) {
	if o == nil || isNil(o.IdOrName) {
    return nil, false
	}
	return o.IdOrName, true
}

// HasIdOrName returns a boolean if a field has been set.
func (o *StoreIdpGoogleReq) HasIdOrName() bool {
	if o != nil && !isNil(o.IdOrName) {
		return true
	}

	return false
}

// SetIdOrName gets a reference to the given string and assigns it to the IdOrName field.
func (o *StoreIdpGoogleReq) SetIdOrName(v string) {
	o.IdOrName = &v
}

// GetIdp returns the Idp field value if set, zero value otherwise.
func (o *StoreIdpGoogleReq) GetIdp() GoogleOpenIDConnectIdentityProviderDTO {
	if o == nil || isNil(o.Idp) {
		var ret GoogleOpenIDConnectIdentityProviderDTO
		return ret
	}
	return *o.Idp
}

// GetIdpOk returns a tuple with the Idp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreIdpGoogleReq) GetIdpOk() (*GoogleOpenIDConnectIdentityProviderDTO, bool) {
	if o == nil || isNil(o.Idp) {
    return nil, false
	}
	return o.Idp, true
}

// HasIdp returns a boolean if a field has been set.
func (o *StoreIdpGoogleReq) HasIdp() bool {
	if o != nil && !isNil(o.Idp) {
		return true
	}

	return false
}

// SetIdp gets a reference to the given GoogleOpenIDConnectIdentityProviderDTO and assigns it to the Idp field.
func (o *StoreIdpGoogleReq) SetIdp(v GoogleOpenIDConnectIdentityProviderDTO) {
	o.Idp = &v
}

func (o StoreIdpGoogleReq) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.IdOrName) {
		toSerialize["idOrName"] = o.IdOrName
	}
	if !isNil(o.Idp) {
		toSerialize["idp"] = o.Idp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StoreIdpGoogleReq) UnmarshalJSON(bytes []byte) (err error) {
	varStoreIdpGoogleReq := _StoreIdpGoogleReq{}

	if err = json.Unmarshal(bytes, &varStoreIdpGoogleReq); err == nil {
		*o = StoreIdpGoogleReq(varStoreIdpGoogleReq)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "idOrName")
		delete(additionalProperties, "idp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStoreIdpGoogleReq struct {
	value *StoreIdpGoogleReq
	isSet bool
}

func (v NullableStoreIdpGoogleReq) Get() *StoreIdpGoogleReq {
	return v.value
}

func (v *NullableStoreIdpGoogleReq) Set(val *StoreIdpGoogleReq) {
	v.value = val
	v.isSet = true
}

func (v NullableStoreIdpGoogleReq) IsSet() bool {
	return v.isSet
}

func (v *NullableStoreIdpGoogleReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoreIdpGoogleReq(val *StoreIdpGoogleReq) *NullableStoreIdpGoogleReq {
	return &NullableStoreIdpGoogleReq{value: val, isSet: true}
}

func (v NullableStoreIdpGoogleReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoreIdpGoogleReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


