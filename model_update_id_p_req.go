/*
 * Atricore Console :: Remote : API
 *
 * # Atricore Console API
 *
 * API version: 1.4.3-SNAPSHOT
 * Contact: sgonzalez@atricore.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package jossoappi

import (
	"encoding/json"
)

// UpdateIdPReq struct for UpdateIdPReq
type UpdateIdPReq struct {
	IdaName *string `json:"idaName,omitempty"`
	Idp *IdentityProviderDTO `json:"idp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateIdPReq UpdateIdPReq

// NewUpdateIdPReq instantiates a new UpdateIdPReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateIdPReq() *UpdateIdPReq {
	this := UpdateIdPReq{}
	return &this
}

// NewUpdateIdPReqWithDefaults instantiates a new UpdateIdPReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateIdPReqWithDefaults() *UpdateIdPReq {
	this := UpdateIdPReq{}
	return &this
}

// GetIdaName returns the IdaName field value if set, zero value otherwise.
func (o *UpdateIdPReq) GetIdaName() string {
	if o == nil || o.IdaName == nil {
		var ret string
		return ret
	}
	return *o.IdaName
}

// GetIdaNameOk returns a tuple with the IdaName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIdPReq) GetIdaNameOk() (*string, bool) {
	if o == nil || o.IdaName == nil {
		return nil, false
	}
	return o.IdaName, true
}

// HasIdaName returns a boolean if a field has been set.
func (o *UpdateIdPReq) HasIdaName() bool {
	if o != nil && o.IdaName != nil {
		return true
	}

	return false
}

// SetIdaName gets a reference to the given string and assigns it to the IdaName field.
func (o *UpdateIdPReq) SetIdaName(v string) {
	o.IdaName = &v
}

// GetIdp returns the Idp field value if set, zero value otherwise.
func (o *UpdateIdPReq) GetIdp() IdentityProviderDTO {
	if o == nil || o.Idp == nil {
		var ret IdentityProviderDTO
		return ret
	}
	return *o.Idp
}

// GetIdpOk returns a tuple with the Idp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIdPReq) GetIdpOk() (*IdentityProviderDTO, bool) {
	if o == nil || o.Idp == nil {
		return nil, false
	}
	return o.Idp, true
}

// HasIdp returns a boolean if a field has been set.
func (o *UpdateIdPReq) HasIdp() bool {
	if o != nil && o.Idp != nil {
		return true
	}

	return false
}

// SetIdp gets a reference to the given IdentityProviderDTO and assigns it to the Idp field.
func (o *UpdateIdPReq) SetIdp(v IdentityProviderDTO) {
	o.Idp = &v
}

func (o UpdateIdPReq) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IdaName != nil {
		toSerialize["idaName"] = o.IdaName
	}
	if o.Idp != nil {
		toSerialize["idp"] = o.Idp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UpdateIdPReq) UnmarshalJSON(bytes []byte) (err error) {
	varUpdateIdPReq := _UpdateIdPReq{}

	if err = json.Unmarshal(bytes, &varUpdateIdPReq); err == nil {
		*o = UpdateIdPReq(varUpdateIdPReq)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "idaName")
		delete(additionalProperties, "idp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateIdPReq struct {
	value *UpdateIdPReq
	isSet bool
}

func (v NullableUpdateIdPReq) Get() *UpdateIdPReq {
	return v.value
}

func (v *NullableUpdateIdPReq) Set(val *UpdateIdPReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateIdPReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateIdPReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateIdPReq(val *UpdateIdPReq) *NullableUpdateIdPReq {
	return &NullableUpdateIdPReq{value: val, isSet: true}
}

func (v NullableUpdateIdPReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateIdPReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


