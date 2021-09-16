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

// SetApplianceStateReq struct for SetApplianceStateReq
type SetApplianceStateReq struct {
	IdOrName *string `json:"idOrName,omitempty"`
	IdaName *string `json:"idaName,omitempty"`
	State *string `json:"state,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SetApplianceStateReq SetApplianceStateReq

// NewSetApplianceStateReq instantiates a new SetApplianceStateReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetApplianceStateReq() *SetApplianceStateReq {
	this := SetApplianceStateReq{}
	return &this
}

// NewSetApplianceStateReqWithDefaults instantiates a new SetApplianceStateReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetApplianceStateReqWithDefaults() *SetApplianceStateReq {
	this := SetApplianceStateReq{}
	return &this
}

// GetIdOrName returns the IdOrName field value if set, zero value otherwise.
func (o *SetApplianceStateReq) GetIdOrName() string {
	if o == nil || o.IdOrName == nil {
		var ret string
		return ret
	}
	return *o.IdOrName
}

// GetIdOrNameOk returns a tuple with the IdOrName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetApplianceStateReq) GetIdOrNameOk() (*string, bool) {
	if o == nil || o.IdOrName == nil {
		return nil, false
	}
	return o.IdOrName, true
}

// HasIdOrName returns a boolean if a field has been set.
func (o *SetApplianceStateReq) HasIdOrName() bool {
	if o != nil && o.IdOrName != nil {
		return true
	}

	return false
}

// SetIdOrName gets a reference to the given string and assigns it to the IdOrName field.
func (o *SetApplianceStateReq) SetIdOrName(v string) {
	o.IdOrName = &v
}

// GetIdaName returns the IdaName field value if set, zero value otherwise.
func (o *SetApplianceStateReq) GetIdaName() string {
	if o == nil || o.IdaName == nil {
		var ret string
		return ret
	}
	return *o.IdaName
}

// GetIdaNameOk returns a tuple with the IdaName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetApplianceStateReq) GetIdaNameOk() (*string, bool) {
	if o == nil || o.IdaName == nil {
		return nil, false
	}
	return o.IdaName, true
}

// HasIdaName returns a boolean if a field has been set.
func (o *SetApplianceStateReq) HasIdaName() bool {
	if o != nil && o.IdaName != nil {
		return true
	}

	return false
}

// SetIdaName gets a reference to the given string and assigns it to the IdaName field.
func (o *SetApplianceStateReq) SetIdaName(v string) {
	o.IdaName = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *SetApplianceStateReq) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetApplianceStateReq) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *SetApplianceStateReq) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *SetApplianceStateReq) SetState(v string) {
	o.State = &v
}

func (o SetApplianceStateReq) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IdOrName != nil {
		toSerialize["idOrName"] = o.IdOrName
	}
	if o.IdaName != nil {
		toSerialize["idaName"] = o.IdaName
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SetApplianceStateReq) UnmarshalJSON(bytes []byte) (err error) {
	varSetApplianceStateReq := _SetApplianceStateReq{}

	if err = json.Unmarshal(bytes, &varSetApplianceStateReq); err == nil {
		*o = SetApplianceStateReq(varSetApplianceStateReq)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "idOrName")
		delete(additionalProperties, "idaName")
		delete(additionalProperties, "state")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSetApplianceStateReq struct {
	value *SetApplianceStateReq
	isSet bool
}

func (v NullableSetApplianceStateReq) Get() *SetApplianceStateReq {
	return v.value
}

func (v *NullableSetApplianceStateReq) Set(val *SetApplianceStateReq) {
	v.value = val
	v.isSet = true
}

func (v NullableSetApplianceStateReq) IsSet() bool {
	return v.isSet
}

func (v *NullableSetApplianceStateReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetApplianceStateReq(val *SetApplianceStateReq) *NullableSetApplianceStateReq {
	return &NullableSetApplianceStateReq{value: val, isSet: true}
}

func (v NullableSetApplianceStateReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetApplianceStateReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

