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

// GetIdPSaml2Req struct for GetIdPSaml2Req
type GetIdPSaml2Req struct {
	IdOrName *string `json:"idOrName,omitempty"`
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetIdPSaml2Req GetIdPSaml2Req

// NewGetIdPSaml2Req instantiates a new GetIdPSaml2Req object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetIdPSaml2Req() *GetIdPSaml2Req {
	this := GetIdPSaml2Req{}
	return &this
}

// NewGetIdPSaml2ReqWithDefaults instantiates a new GetIdPSaml2Req object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetIdPSaml2ReqWithDefaults() *GetIdPSaml2Req {
	this := GetIdPSaml2Req{}
	return &this
}

// GetIdOrName returns the IdOrName field value if set, zero value otherwise.
func (o *GetIdPSaml2Req) GetIdOrName() string {
	if o == nil || isNil(o.IdOrName) {
		var ret string
		return ret
	}
	return *o.IdOrName
}

// GetIdOrNameOk returns a tuple with the IdOrName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIdPSaml2Req) GetIdOrNameOk() (*string, bool) {
	if o == nil || isNil(o.IdOrName) {
    return nil, false
	}
	return o.IdOrName, true
}

// HasIdOrName returns a boolean if a field has been set.
func (o *GetIdPSaml2Req) HasIdOrName() bool {
	if o != nil && !isNil(o.IdOrName) {
		return true
	}

	return false
}

// SetIdOrName gets a reference to the given string and assigns it to the IdOrName field.
func (o *GetIdPSaml2Req) SetIdOrName(v string) {
	o.IdOrName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetIdPSaml2Req) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIdPSaml2Req) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetIdPSaml2Req) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetIdPSaml2Req) SetName(v string) {
	o.Name = &v
}

func (o GetIdPSaml2Req) MarshalJSON() ([]byte, error) {
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

func (o *GetIdPSaml2Req) UnmarshalJSON(bytes []byte) (err error) {
	varGetIdPSaml2Req := _GetIdPSaml2Req{}

	if err = json.Unmarshal(bytes, &varGetIdPSaml2Req); err == nil {
		*o = GetIdPSaml2Req(varGetIdPSaml2Req)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "idOrName")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetIdPSaml2Req struct {
	value *GetIdPSaml2Req
	isSet bool
}

func (v NullableGetIdPSaml2Req) Get() *GetIdPSaml2Req {
	return v.value
}

func (v *NullableGetIdPSaml2Req) Set(val *GetIdPSaml2Req) {
	v.value = val
	v.isSet = true
}

func (v NullableGetIdPSaml2Req) IsSet() bool {
	return v.isSet
}

func (v *NullableGetIdPSaml2Req) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetIdPSaml2Req(val *GetIdPSaml2Req) *NullableGetIdPSaml2Req {
	return &NullableGetIdPSaml2Req{value: val, isSet: true}
}

func (v NullableGetIdPSaml2Req) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetIdPSaml2Req) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


