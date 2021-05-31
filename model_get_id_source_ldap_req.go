/*
 * Atricore Console :: Remote : API
 *
 * # Atricore Console API
 *
 * API version: 1.4.3-SNAPSHOT
 * Contact: sgonzalez@atricore.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package jossoappi

import (
	"encoding/json"
)

// GetIdSourceLdapReq struct for GetIdSourceLdapReq
type GetIdSourceLdapReq struct {
	IdaName *string `json:"idaName,omitempty"`
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetIdSourceLdapReq GetIdSourceLdapReq

// NewGetIdSourceLdapReq instantiates a new GetIdSourceLdapReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetIdSourceLdapReq() *GetIdSourceLdapReq {
	this := GetIdSourceLdapReq{}
	return &this
}

// NewGetIdSourceLdapReqWithDefaults instantiates a new GetIdSourceLdapReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetIdSourceLdapReqWithDefaults() *GetIdSourceLdapReq {
	this := GetIdSourceLdapReq{}
	return &this
}

// GetIdaName returns the IdaName field value if set, zero value otherwise.
func (o *GetIdSourceLdapReq) GetIdaName() string {
	if o == nil || o.IdaName == nil {
		var ret string
		return ret
	}
	return *o.IdaName
}

// GetIdaNameOk returns a tuple with the IdaName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIdSourceLdapReq) GetIdaNameOk() (*string, bool) {
	if o == nil || o.IdaName == nil {
		return nil, false
	}
	return o.IdaName, true
}

// HasIdaName returns a boolean if a field has been set.
func (o *GetIdSourceLdapReq) HasIdaName() bool {
	if o != nil && o.IdaName != nil {
		return true
	}

	return false
}

// SetIdaName gets a reference to the given string and assigns it to the IdaName field.
func (o *GetIdSourceLdapReq) SetIdaName(v string) {
	o.IdaName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetIdSourceLdapReq) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIdSourceLdapReq) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetIdSourceLdapReq) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetIdSourceLdapReq) SetName(v string) {
	o.Name = &v
}

func (o GetIdSourceLdapReq) MarshalJSON() ([]byte, error) {
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

func (o *GetIdSourceLdapReq) UnmarshalJSON(bytes []byte) (err error) {
	varGetIdSourceLdapReq := _GetIdSourceLdapReq{}

	if err = json.Unmarshal(bytes, &varGetIdSourceLdapReq); err == nil {
		*o = GetIdSourceLdapReq(varGetIdSourceLdapReq)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "idaName")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetIdSourceLdapReq struct {
	value *GetIdSourceLdapReq
	isSet bool
}

func (v NullableGetIdSourceLdapReq) Get() *GetIdSourceLdapReq {
	return v.value
}

func (v *NullableGetIdSourceLdapReq) Set(val *GetIdSourceLdapReq) {
	v.value = val
	v.isSet = true
}

func (v NullableGetIdSourceLdapReq) IsSet() bool {
	return v.isSet
}

func (v *NullableGetIdSourceLdapReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetIdSourceLdapReq(val *GetIdSourceLdapReq) *NullableGetIdSourceLdapReq {
	return &NullableGetIdSourceLdapReq{value: val, isSet: true}
}

func (v NullableGetIdSourceLdapReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetIdSourceLdapReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

