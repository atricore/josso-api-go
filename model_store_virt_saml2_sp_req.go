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

// StoreVirtSaml2SpReq struct for StoreVirtSaml2SpReq
type StoreVirtSaml2SpReq struct {
	IdOrName *string `json:"idOrName,omitempty"`
	Sp *VirtualSaml2ServiceProviderDTO `json:"sp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StoreVirtSaml2SpReq StoreVirtSaml2SpReq

// NewStoreVirtSaml2SpReq instantiates a new StoreVirtSaml2SpReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoreVirtSaml2SpReq() *StoreVirtSaml2SpReq {
	this := StoreVirtSaml2SpReq{}
	return &this
}

// NewStoreVirtSaml2SpReqWithDefaults instantiates a new StoreVirtSaml2SpReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoreVirtSaml2SpReqWithDefaults() *StoreVirtSaml2SpReq {
	this := StoreVirtSaml2SpReq{}
	return &this
}

// GetIdOrName returns the IdOrName field value if set, zero value otherwise.
func (o *StoreVirtSaml2SpReq) GetIdOrName() string {
	if o == nil || isNil(o.IdOrName) {
		var ret string
		return ret
	}
	return *o.IdOrName
}

// GetIdOrNameOk returns a tuple with the IdOrName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreVirtSaml2SpReq) GetIdOrNameOk() (*string, bool) {
	if o == nil || isNil(o.IdOrName) {
    return nil, false
	}
	return o.IdOrName, true
}

// HasIdOrName returns a boolean if a field has been set.
func (o *StoreVirtSaml2SpReq) HasIdOrName() bool {
	if o != nil && !isNil(o.IdOrName) {
		return true
	}

	return false
}

// SetIdOrName gets a reference to the given string and assigns it to the IdOrName field.
func (o *StoreVirtSaml2SpReq) SetIdOrName(v string) {
	o.IdOrName = &v
}

// GetSp returns the Sp field value if set, zero value otherwise.
func (o *StoreVirtSaml2SpReq) GetSp() VirtualSaml2ServiceProviderDTO {
	if o == nil || isNil(o.Sp) {
		var ret VirtualSaml2ServiceProviderDTO
		return ret
	}
	return *o.Sp
}

// GetSpOk returns a tuple with the Sp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreVirtSaml2SpReq) GetSpOk() (*VirtualSaml2ServiceProviderDTO, bool) {
	if o == nil || isNil(o.Sp) {
    return nil, false
	}
	return o.Sp, true
}

// HasSp returns a boolean if a field has been set.
func (o *StoreVirtSaml2SpReq) HasSp() bool {
	if o != nil && !isNil(o.Sp) {
		return true
	}

	return false
}

// SetSp gets a reference to the given VirtualSaml2ServiceProviderDTO and assigns it to the Sp field.
func (o *StoreVirtSaml2SpReq) SetSp(v VirtualSaml2ServiceProviderDTO) {
	o.Sp = &v
}

func (o StoreVirtSaml2SpReq) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.IdOrName) {
		toSerialize["idOrName"] = o.IdOrName
	}
	if !isNil(o.Sp) {
		toSerialize["sp"] = o.Sp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StoreVirtSaml2SpReq) UnmarshalJSON(bytes []byte) (err error) {
	varStoreVirtSaml2SpReq := _StoreVirtSaml2SpReq{}

	if err = json.Unmarshal(bytes, &varStoreVirtSaml2SpReq); err == nil {
		*o = StoreVirtSaml2SpReq(varStoreVirtSaml2SpReq)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "idOrName")
		delete(additionalProperties, "sp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStoreVirtSaml2SpReq struct {
	value *StoreVirtSaml2SpReq
	isSet bool
}

func (v NullableStoreVirtSaml2SpReq) Get() *StoreVirtSaml2SpReq {
	return v.value
}

func (v *NullableStoreVirtSaml2SpReq) Set(val *StoreVirtSaml2SpReq) {
	v.value = val
	v.isSet = true
}

func (v NullableStoreVirtSaml2SpReq) IsSet() bool {
	return v.isSet
}

func (v *NullableStoreVirtSaml2SpReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoreVirtSaml2SpReq(val *StoreVirtSaml2SpReq) *NullableStoreVirtSaml2SpReq {
	return &NullableStoreVirtSaml2SpReq{value: val, isSet: true}
}

func (v NullableStoreVirtSaml2SpReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoreVirtSaml2SpReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


