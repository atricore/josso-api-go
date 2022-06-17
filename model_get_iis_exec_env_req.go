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

// GetIisExecEnvReq struct for GetIisExecEnvReq
type GetIisExecEnvReq struct {
	IdaName *string `json:"idaName,omitempty"`
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetIisExecEnvReq GetIisExecEnvReq

// NewGetIisExecEnvReq instantiates a new GetIisExecEnvReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetIisExecEnvReq() *GetIisExecEnvReq {
	this := GetIisExecEnvReq{}
	return &this
}

// NewGetIisExecEnvReqWithDefaults instantiates a new GetIisExecEnvReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetIisExecEnvReqWithDefaults() *GetIisExecEnvReq {
	this := GetIisExecEnvReq{}
	return &this
}

// GetIdaName returns the IdaName field value if set, zero value otherwise.
func (o *GetIisExecEnvReq) GetIdaName() string {
	if o == nil || o.IdaName == nil {
		var ret string
		return ret
	}
	return *o.IdaName
}

// GetIdaNameOk returns a tuple with the IdaName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIisExecEnvReq) GetIdaNameOk() (*string, bool) {
	if o == nil || o.IdaName == nil {
		return nil, false
	}
	return o.IdaName, true
}

// HasIdaName returns a boolean if a field has been set.
func (o *GetIisExecEnvReq) HasIdaName() bool {
	if o != nil && o.IdaName != nil {
		return true
	}

	return false
}

// SetIdaName gets a reference to the given string and assigns it to the IdaName field.
func (o *GetIisExecEnvReq) SetIdaName(v string) {
	o.IdaName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetIisExecEnvReq) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIisExecEnvReq) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetIisExecEnvReq) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetIisExecEnvReq) SetName(v string) {
	o.Name = &v
}

func (o GetIisExecEnvReq) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IdaName != nil {
		toSerialize["idaName"] = o.IdaName
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GetIisExecEnvReq) UnmarshalJSON(bytes []byte) (err error) {
	varGetIisExecEnvReq := _GetIisExecEnvReq{}

	if err = json.Unmarshal(bytes, &varGetIisExecEnvReq); err == nil {
		*o = GetIisExecEnvReq(varGetIisExecEnvReq)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "idaName")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetIisExecEnvReq struct {
	value *GetIisExecEnvReq
	isSet bool
}

func (v NullableGetIisExecEnvReq) Get() *GetIisExecEnvReq {
	return v.value
}

func (v *NullableGetIisExecEnvReq) Set(val *GetIisExecEnvReq) {
	v.value = val
	v.isSet = true
}

func (v NullableGetIisExecEnvReq) IsSet() bool {
	return v.isSet
}

func (v *NullableGetIisExecEnvReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetIisExecEnvReq(val *GetIisExecEnvReq) *NullableGetIisExecEnvReq {
	return &NullableGetIisExecEnvReq{value: val, isSet: true}
}

func (v NullableGetIisExecEnvReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetIisExecEnvReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

