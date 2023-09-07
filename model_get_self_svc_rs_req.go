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

// GetSelfSvcRsReq struct for GetSelfSvcRsReq
type GetSelfSvcRsReq struct {
	IdOrName *string `json:"idOrName,omitempty"`
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetSelfSvcRsReq GetSelfSvcRsReq

// NewGetSelfSvcRsReq instantiates a new GetSelfSvcRsReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSelfSvcRsReq() *GetSelfSvcRsReq {
	this := GetSelfSvcRsReq{}
	return &this
}

// NewGetSelfSvcRsReqWithDefaults instantiates a new GetSelfSvcRsReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSelfSvcRsReqWithDefaults() *GetSelfSvcRsReq {
	this := GetSelfSvcRsReq{}
	return &this
}

// GetIdOrName returns the IdOrName field value if set, zero value otherwise.
func (o *GetSelfSvcRsReq) GetIdOrName() string {
	if o == nil || isNil(o.IdOrName) {
		var ret string
		return ret
	}
	return *o.IdOrName
}

// GetIdOrNameOk returns a tuple with the IdOrName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSelfSvcRsReq) GetIdOrNameOk() (*string, bool) {
	if o == nil || isNil(o.IdOrName) {
    return nil, false
	}
	return o.IdOrName, true
}

// HasIdOrName returns a boolean if a field has been set.
func (o *GetSelfSvcRsReq) HasIdOrName() bool {
	if o != nil && !isNil(o.IdOrName) {
		return true
	}

	return false
}

// SetIdOrName gets a reference to the given string and assigns it to the IdOrName field.
func (o *GetSelfSvcRsReq) SetIdOrName(v string) {
	o.IdOrName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetSelfSvcRsReq) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSelfSvcRsReq) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetSelfSvcRsReq) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetSelfSvcRsReq) SetName(v string) {
	o.Name = &v
}

func (o GetSelfSvcRsReq) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.IdOrName) {
		toSerialize["idOrName"] = o.IdOrName
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GetSelfSvcRsReq) UnmarshalJSON(bytes []byte) (err error) {
	varGetSelfSvcRsReq := _GetSelfSvcRsReq{}

	if err = json.Unmarshal(bytes, &varGetSelfSvcRsReq); err == nil {
		*o = GetSelfSvcRsReq(varGetSelfSvcRsReq)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "idOrName")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetSelfSvcRsReq struct {
	value *GetSelfSvcRsReq
	isSet bool
}

func (v NullableGetSelfSvcRsReq) Get() *GetSelfSvcRsReq {
	return v.value
}

func (v *NullableGetSelfSvcRsReq) Set(val *GetSelfSvcRsReq) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSelfSvcRsReq) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSelfSvcRsReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSelfSvcRsReq(val *GetSelfSvcRsReq) *NullableGetSelfSvcRsReq {
	return &NullableGetSelfSvcRsReq{value: val, isSet: true}
}

func (v NullableGetSelfSvcRsReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSelfSvcRsReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

