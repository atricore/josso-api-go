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

// StoreIdpAzureReq struct for StoreIdpAzureReq
type StoreIdpAzureReq struct {
	IdOrName *string `json:"idOrName,omitempty"`
	Idp *AzureOpenIDConnectIdentityProviderDTO `json:"idp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StoreIdpAzureReq StoreIdpAzureReq

// NewStoreIdpAzureReq instantiates a new StoreIdpAzureReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoreIdpAzureReq() *StoreIdpAzureReq {
	this := StoreIdpAzureReq{}
	return &this
}

// NewStoreIdpAzureReqWithDefaults instantiates a new StoreIdpAzureReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoreIdpAzureReqWithDefaults() *StoreIdpAzureReq {
	this := StoreIdpAzureReq{}
	return &this
}

// GetIdOrName returns the IdOrName field value if set, zero value otherwise.
func (o *StoreIdpAzureReq) GetIdOrName() string {
	if o == nil || isNil(o.IdOrName) {
		var ret string
		return ret
	}
	return *o.IdOrName
}

// GetIdOrNameOk returns a tuple with the IdOrName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreIdpAzureReq) GetIdOrNameOk() (*string, bool) {
	if o == nil || isNil(o.IdOrName) {
    return nil, false
	}
	return o.IdOrName, true
}

// HasIdOrName returns a boolean if a field has been set.
func (o *StoreIdpAzureReq) HasIdOrName() bool {
	if o != nil && !isNil(o.IdOrName) {
		return true
	}

	return false
}

// SetIdOrName gets a reference to the given string and assigns it to the IdOrName field.
func (o *StoreIdpAzureReq) SetIdOrName(v string) {
	o.IdOrName = &v
}

// GetIdp returns the Idp field value if set, zero value otherwise.
func (o *StoreIdpAzureReq) GetIdp() AzureOpenIDConnectIdentityProviderDTO {
	if o == nil || isNil(o.Idp) {
		var ret AzureOpenIDConnectIdentityProviderDTO
		return ret
	}
	return *o.Idp
}

// GetIdpOk returns a tuple with the Idp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreIdpAzureReq) GetIdpOk() (*AzureOpenIDConnectIdentityProviderDTO, bool) {
	if o == nil || isNil(o.Idp) {
    return nil, false
	}
	return o.Idp, true
}

// HasIdp returns a boolean if a field has been set.
func (o *StoreIdpAzureReq) HasIdp() bool {
	if o != nil && !isNil(o.Idp) {
		return true
	}

	return false
}

// SetIdp gets a reference to the given AzureOpenIDConnectIdentityProviderDTO and assigns it to the Idp field.
func (o *StoreIdpAzureReq) SetIdp(v AzureOpenIDConnectIdentityProviderDTO) {
	o.Idp = &v
}

func (o StoreIdpAzureReq) MarshalJSON() ([]byte, error) {
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

func (o *StoreIdpAzureReq) UnmarshalJSON(bytes []byte) (err error) {
	varStoreIdpAzureReq := _StoreIdpAzureReq{}

	if err = json.Unmarshal(bytes, &varStoreIdpAzureReq); err == nil {
		*o = StoreIdpAzureReq(varStoreIdpAzureReq)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "idOrName")
		delete(additionalProperties, "idp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStoreIdpAzureReq struct {
	value *StoreIdpAzureReq
	isSet bool
}

func (v NullableStoreIdpAzureReq) Get() *StoreIdpAzureReq {
	return v.value
}

func (v *NullableStoreIdpAzureReq) Set(val *StoreIdpAzureReq) {
	v.value = val
	v.isSet = true
}

func (v NullableStoreIdpAzureReq) IsSet() bool {
	return v.isSet
}

func (v *NullableStoreIdpAzureReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoreIdpAzureReq(val *StoreIdpAzureReq) *NullableStoreIdpAzureReq {
	return &NullableStoreIdpAzureReq{value: val, isSet: true}
}

func (v NullableStoreIdpAzureReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoreIdpAzureReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


