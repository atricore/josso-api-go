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

// GetIdVaultReq struct for GetIdVaultReq
type GetIdVaultReq struct {
	IdaName *string `json:"idaName,omitempty"`
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetIdVaultReq GetIdVaultReq

// NewGetIdVaultReq instantiates a new GetIdVaultReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetIdVaultReq() *GetIdVaultReq {
	this := GetIdVaultReq{}
	return &this
}

// NewGetIdVaultReqWithDefaults instantiates a new GetIdVaultReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetIdVaultReqWithDefaults() *GetIdVaultReq {
	this := GetIdVaultReq{}
	return &this
}

// GetIdaName returns the IdaName field value if set, zero value otherwise.
func (o *GetIdVaultReq) GetIdaName() string {
	if o == nil || o.IdaName == nil {
		var ret string
		return ret
	}
	return *o.IdaName
}

// GetIdaNameOk returns a tuple with the IdaName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIdVaultReq) GetIdaNameOk() (*string, bool) {
	if o == nil || o.IdaName == nil {
		return nil, false
	}
	return o.IdaName, true
}

// HasIdaName returns a boolean if a field has been set.
func (o *GetIdVaultReq) HasIdaName() bool {
	if o != nil && o.IdaName != nil {
		return true
	}

	return false
}

// SetIdaName gets a reference to the given string and assigns it to the IdaName field.
func (o *GetIdVaultReq) SetIdaName(v string) {
	o.IdaName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetIdVaultReq) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIdVaultReq) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetIdVaultReq) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetIdVaultReq) SetName(v string) {
	o.Name = &v
}

func (o GetIdVaultReq) MarshalJSON() ([]byte, error) {
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

func (o *GetIdVaultReq) UnmarshalJSON(bytes []byte) (err error) {
	varGetIdVaultReq := _GetIdVaultReq{}

	if err = json.Unmarshal(bytes, &varGetIdVaultReq); err == nil {
		*o = GetIdVaultReq(varGetIdVaultReq)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "idaName")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetIdVaultReq struct {
	value *GetIdVaultReq
	isSet bool
}

func (v NullableGetIdVaultReq) Get() *GetIdVaultReq {
	return v.value
}

func (v *NullableGetIdVaultReq) Set(val *GetIdVaultReq) {
	v.value = val
	v.isSet = true
}

func (v NullableGetIdVaultReq) IsSet() bool {
	return v.isSet
}

func (v *NullableGetIdVaultReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetIdVaultReq(val *GetIdVaultReq) *NullableGetIdVaultReq {
	return &NullableGetIdVaultReq{value: val, isSet: true}
}

func (v NullableGetIdVaultReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetIdVaultReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


