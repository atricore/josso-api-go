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

// GetSharepointRsReq struct for GetSharepointRsReq
type GetSharepointRsReq struct {
	IdaName *string `json:"idaName,omitempty"`
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetSharepointRsReq GetSharepointRsReq

// NewGetSharepointRsReq instantiates a new GetSharepointRsReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSharepointRsReq() *GetSharepointRsReq {
	this := GetSharepointRsReq{}
	return &this
}

// NewGetSharepointRsReqWithDefaults instantiates a new GetSharepointRsReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSharepointRsReqWithDefaults() *GetSharepointRsReq {
	this := GetSharepointRsReq{}
	return &this
}

// GetIdaName returns the IdaName field value if set, zero value otherwise.
func (o *GetSharepointRsReq) GetIdaName() string {
	if o == nil || o.IdaName == nil {
		var ret string
		return ret
	}
	return *o.IdaName
}

// GetIdaNameOk returns a tuple with the IdaName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSharepointRsReq) GetIdaNameOk() (*string, bool) {
	if o == nil || o.IdaName == nil {
		return nil, false
	}
	return o.IdaName, true
}

// HasIdaName returns a boolean if a field has been set.
func (o *GetSharepointRsReq) HasIdaName() bool {
	if o != nil && o.IdaName != nil {
		return true
	}

	return false
}

// SetIdaName gets a reference to the given string and assigns it to the IdaName field.
func (o *GetSharepointRsReq) SetIdaName(v string) {
	o.IdaName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetSharepointRsReq) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSharepointRsReq) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetSharepointRsReq) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetSharepointRsReq) SetName(v string) {
	o.Name = &v
}

func (o GetSharepointRsReq) MarshalJSON() ([]byte, error) {
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

func (o *GetSharepointRsReq) UnmarshalJSON(bytes []byte) (err error) {
	varGetSharepointRsReq := _GetSharepointRsReq{}

	if err = json.Unmarshal(bytes, &varGetSharepointRsReq); err == nil {
		*o = GetSharepointRsReq(varGetSharepointRsReq)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "idaName")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetSharepointRsReq struct {
	value *GetSharepointRsReq
	isSet bool
}

func (v NullableGetSharepointRsReq) Get() *GetSharepointRsReq {
	return v.value
}

func (v *NullableGetSharepointRsReq) Set(val *GetSharepointRsReq) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSharepointRsReq) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSharepointRsReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSharepointRsReq(val *GetSharepointRsReq) *NullableGetSharepointRsReq {
	return &NullableGetSharepointRsReq{value: val, isSet: true}
}

func (v NullableGetSharepointRsReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSharepointRsReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

