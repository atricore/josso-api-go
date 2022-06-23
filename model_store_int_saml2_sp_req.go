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

// StoreIntSaml2SpReq struct for StoreIntSaml2SpReq
type StoreIntSaml2SpReq struct {
	IdaName *string `json:"idaName,omitempty"`
	Sp *InternalSaml2ServiceProviderDTO `json:"sp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StoreIntSaml2SpReq StoreIntSaml2SpReq

// NewStoreIntSaml2SpReq instantiates a new StoreIntSaml2SpReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoreIntSaml2SpReq() *StoreIntSaml2SpReq {
	this := StoreIntSaml2SpReq{}
	return &this
}

// NewStoreIntSaml2SpReqWithDefaults instantiates a new StoreIntSaml2SpReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoreIntSaml2SpReqWithDefaults() *StoreIntSaml2SpReq {
	this := StoreIntSaml2SpReq{}
	return &this
}

// GetIdaName returns the IdaName field value if set, zero value otherwise.
func (o *StoreIntSaml2SpReq) GetIdaName() string {
	if o == nil || o.IdaName == nil {
		var ret string
		return ret
	}
	return *o.IdaName
}

// GetIdaNameOk returns a tuple with the IdaName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreIntSaml2SpReq) GetIdaNameOk() (*string, bool) {
	if o == nil || o.IdaName == nil {
		return nil, false
	}
	return o.IdaName, true
}

// HasIdaName returns a boolean if a field has been set.
func (o *StoreIntSaml2SpReq) HasIdaName() bool {
	if o != nil && o.IdaName != nil {
		return true
	}

	return false
}

// SetIdaName gets a reference to the given string and assigns it to the IdaName field.
func (o *StoreIntSaml2SpReq) SetIdaName(v string) {
	o.IdaName = &v
}

// GetSp returns the Sp field value if set, zero value otherwise.
func (o *StoreIntSaml2SpReq) GetSp() InternalSaml2ServiceProviderDTO {
	if o == nil || o.Sp == nil {
		var ret InternalSaml2ServiceProviderDTO
		return ret
	}
	return *o.Sp
}

// GetSpOk returns a tuple with the Sp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreIntSaml2SpReq) GetSpOk() (*InternalSaml2ServiceProviderDTO, bool) {
	if o == nil || o.Sp == nil {
		return nil, false
	}
	return o.Sp, true
}

// HasSp returns a boolean if a field has been set.
func (o *StoreIntSaml2SpReq) HasSp() bool {
	if o != nil && o.Sp != nil {
		return true
	}

	return false
}

// SetSp gets a reference to the given InternalSaml2ServiceProviderDTO and assigns it to the Sp field.
func (o *StoreIntSaml2SpReq) SetSp(v InternalSaml2ServiceProviderDTO) {
	o.Sp = &v
}

func (o StoreIntSaml2SpReq) MarshalJSON() ([]byte, error) {
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

func (o *StoreIntSaml2SpReq) UnmarshalJSON(bytes []byte) (err error) {
	varStoreIntSaml2SpReq := _StoreIntSaml2SpReq{}

	if err = json.Unmarshal(bytes, &varStoreIntSaml2SpReq); err == nil {
		*o = StoreIntSaml2SpReq(varStoreIntSaml2SpReq)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "idaName")
		delete(additionalProperties, "sp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStoreIntSaml2SpReq struct {
	value *StoreIntSaml2SpReq
	isSet bool
}

func (v NullableStoreIntSaml2SpReq) Get() *StoreIntSaml2SpReq {
	return v.value
}

func (v *NullableStoreIntSaml2SpReq) Set(val *StoreIntSaml2SpReq) {
	v.value = val
	v.isSet = true
}

func (v NullableStoreIntSaml2SpReq) IsSet() bool {
	return v.isSet
}

func (v *NullableStoreIntSaml2SpReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoreIntSaml2SpReq(val *StoreIntSaml2SpReq) *NullableStoreIntSaml2SpReq {
	return &NullableStoreIntSaml2SpReq{value: val, isSet: true}
}

func (v NullableStoreIntSaml2SpReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoreIntSaml2SpReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


