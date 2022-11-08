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

// GetIdpGoogleReq struct for GetIdpGoogleReq
type GetIdpGoogleReq struct {
	IdaName *string `json:"idaName,omitempty"`
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetIdpGoogleReq GetIdpGoogleReq

// NewGetIdpGoogleReq instantiates a new GetIdpGoogleReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetIdpGoogleReq() *GetIdpGoogleReq {
	this := GetIdpGoogleReq{}
	return &this
}

// NewGetIdpGoogleReqWithDefaults instantiates a new GetIdpGoogleReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetIdpGoogleReqWithDefaults() *GetIdpGoogleReq {
	this := GetIdpGoogleReq{}
	return &this
}

// GetIdaName returns the IdaName field value if set, zero value otherwise.
func (o *GetIdpGoogleReq) GetIdaName() string {
	if o == nil || isNil(o.IdaName) {
		var ret string
		return ret
	}
	return *o.IdaName
}

// GetIdaNameOk returns a tuple with the IdaName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIdpGoogleReq) GetIdaNameOk() (*string, bool) {
	if o == nil || isNil(o.IdaName) {
    return nil, false
	}
	return o.IdaName, true
}

// HasIdaName returns a boolean if a field has been set.
func (o *GetIdpGoogleReq) HasIdaName() bool {
	if o != nil && !isNil(o.IdaName) {
		return true
	}

	return false
}

// SetIdaName gets a reference to the given string and assigns it to the IdaName field.
func (o *GetIdpGoogleReq) SetIdaName(v string) {
	o.IdaName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetIdpGoogleReq) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIdpGoogleReq) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetIdpGoogleReq) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetIdpGoogleReq) SetName(v string) {
	o.Name = &v
}

func (o GetIdpGoogleReq) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.IdaName) {
		toSerialize["idaName"] = o.IdaName
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GetIdpGoogleReq) UnmarshalJSON(bytes []byte) (err error) {
	varGetIdpGoogleReq := _GetIdpGoogleReq{}

	if err = json.Unmarshal(bytes, &varGetIdpGoogleReq); err == nil {
		*o = GetIdpGoogleReq(varGetIdpGoogleReq)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "idaName")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetIdpGoogleReq struct {
	value *GetIdpGoogleReq
	isSet bool
}

func (v NullableGetIdpGoogleReq) Get() *GetIdpGoogleReq {
	return v.value
}

func (v *NullableGetIdpGoogleReq) Set(val *GetIdpGoogleReq) {
	v.value = val
	v.isSet = true
}

func (v NullableGetIdpGoogleReq) IsSet() bool {
	return v.isSet
}

func (v *NullableGetIdpGoogleReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetIdpGoogleReq(val *GetIdpGoogleReq) *NullableGetIdpGoogleReq {
	return &NullableGetIdpGoogleReq{value: val, isSet: true}
}

func (v NullableGetIdpGoogleReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetIdpGoogleReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


