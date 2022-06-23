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

// StoreExtSaml2SpReq struct for StoreExtSaml2SpReq
type StoreExtSaml2SpReq struct {
	IdaName *string `json:"idaName,omitempty"`
	Sp *ExternalSaml2ServiceProviderDTO `json:"sp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StoreExtSaml2SpReq StoreExtSaml2SpReq

// NewStoreExtSaml2SpReq instantiates a new StoreExtSaml2SpReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoreExtSaml2SpReq() *StoreExtSaml2SpReq {
	this := StoreExtSaml2SpReq{}
	return &this
}

// NewStoreExtSaml2SpReqWithDefaults instantiates a new StoreExtSaml2SpReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoreExtSaml2SpReqWithDefaults() *StoreExtSaml2SpReq {
	this := StoreExtSaml2SpReq{}
	return &this
}

// GetIdaName returns the IdaName field value if set, zero value otherwise.
func (o *StoreExtSaml2SpReq) GetIdaName() string {
	if o == nil || o.IdaName == nil {
		var ret string
		return ret
	}
	return *o.IdaName
}

// GetIdaNameOk returns a tuple with the IdaName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreExtSaml2SpReq) GetIdaNameOk() (*string, bool) {
	if o == nil || o.IdaName == nil {
		return nil, false
	}
	return o.IdaName, true
}

// HasIdaName returns a boolean if a field has been set.
func (o *StoreExtSaml2SpReq) HasIdaName() bool {
	if o != nil && o.IdaName != nil {
		return true
	}

	return false
}

// SetIdaName gets a reference to the given string and assigns it to the IdaName field.
func (o *StoreExtSaml2SpReq) SetIdaName(v string) {
	o.IdaName = &v
}

// GetSp returns the Sp field value if set, zero value otherwise.
func (o *StoreExtSaml2SpReq) GetSp() ExternalSaml2ServiceProviderDTO {
	if o == nil || o.Sp == nil {
		var ret ExternalSaml2ServiceProviderDTO
		return ret
	}
	return *o.Sp
}

// GetSpOk returns a tuple with the Sp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreExtSaml2SpReq) GetSpOk() (*ExternalSaml2ServiceProviderDTO, bool) {
	if o == nil || o.Sp == nil {
		return nil, false
	}
	return o.Sp, true
}

// HasSp returns a boolean if a field has been set.
func (o *StoreExtSaml2SpReq) HasSp() bool {
	if o != nil && o.Sp != nil {
		return true
	}

	return false
}

// SetSp gets a reference to the given ExternalSaml2ServiceProviderDTO and assigns it to the Sp field.
func (o *StoreExtSaml2SpReq) SetSp(v ExternalSaml2ServiceProviderDTO) {
	o.Sp = &v
}

func (o StoreExtSaml2SpReq) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IdaName != nil {
		toSerialize["idaName"] = o.IdaName
	}
	if o.Sp != nil {
		toSerialize["sp"] = o.Sp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StoreExtSaml2SpReq) UnmarshalJSON(bytes []byte) (err error) {
	varStoreExtSaml2SpReq := _StoreExtSaml2SpReq{}

	if err = json.Unmarshal(bytes, &varStoreExtSaml2SpReq); err == nil {
		*o = StoreExtSaml2SpReq(varStoreExtSaml2SpReq)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "idaName")
		delete(additionalProperties, "sp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStoreExtSaml2SpReq struct {
	value *StoreExtSaml2SpReq
	isSet bool
}

func (v NullableStoreExtSaml2SpReq) Get() *StoreExtSaml2SpReq {
	return v.value
}

func (v *NullableStoreExtSaml2SpReq) Set(val *StoreExtSaml2SpReq) {
	v.value = val
	v.isSet = true
}

func (v NullableStoreExtSaml2SpReq) IsSet() bool {
	return v.isSet
}

func (v *NullableStoreExtSaml2SpReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoreExtSaml2SpReq(val *StoreExtSaml2SpReq) *NullableStoreExtSaml2SpReq {
	return &NullableStoreExtSaml2SpReq{value: val, isSet: true}
}

func (v NullableStoreExtSaml2SpReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoreExtSaml2SpReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


