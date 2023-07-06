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

// GetApplianceReq struct for GetApplianceReq
type GetApplianceReq struct {
	IdOrName *string `json:"idOrName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetApplianceReq GetApplianceReq

// NewGetApplianceReq instantiates a new GetApplianceReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetApplianceReq() *GetApplianceReq {
	this := GetApplianceReq{}
	return &this
}

// NewGetApplianceReqWithDefaults instantiates a new GetApplianceReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetApplianceReqWithDefaults() *GetApplianceReq {
	this := GetApplianceReq{}
	return &this
}

// GetIdOrName returns the IdOrName field value if set, zero value otherwise.
func (o *GetApplianceReq) GetIdOrName() string {
	if o == nil || isNil(o.IdOrName) {
		var ret string
		return ret
	}
	return *o.IdOrName
}

// GetIdOrNameOk returns a tuple with the IdOrName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetApplianceReq) GetIdOrNameOk() (*string, bool) {
	if o == nil || isNil(o.IdOrName) {
    return nil, false
	}
	return o.IdOrName, true
}

// HasIdOrName returns a boolean if a field has been set.
func (o *GetApplianceReq) HasIdOrName() bool {
	if o != nil && !isNil(o.IdOrName) {
		return true
	}

	return false
}

// SetIdOrName gets a reference to the given string and assigns it to the IdOrName field.
func (o *GetApplianceReq) SetIdOrName(v string) {
	o.IdOrName = &v
}

func (o GetApplianceReq) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.IdOrName) {
		toSerialize["idOrName"] = o.IdOrName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GetApplianceReq) UnmarshalJSON(bytes []byte) (err error) {
	varGetApplianceReq := _GetApplianceReq{}

	if err = json.Unmarshal(bytes, &varGetApplianceReq); err == nil {
		*o = GetApplianceReq(varGetApplianceReq)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "idOrName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetApplianceReq struct {
	value *GetApplianceReq
	isSet bool
}

func (v NullableGetApplianceReq) Get() *GetApplianceReq {
	return v.value
}

func (v *NullableGetApplianceReq) Set(val *GetApplianceReq) {
	v.value = val
	v.isSet = true
}

func (v NullableGetApplianceReq) IsSet() bool {
	return v.isSet
}

func (v *NullableGetApplianceReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetApplianceReq(val *GetApplianceReq) *NullableGetApplianceReq {
	return &NullableGetApplianceReq{value: val, isSet: true}
}

func (v NullableGetApplianceReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetApplianceReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


