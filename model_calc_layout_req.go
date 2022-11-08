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

// CalcLayoutReq struct for CalcLayoutReq
type CalcLayoutReq struct {
	Export *bool `json:"export,omitempty"`
	IdOrName *string `json:"idOrName,omitempty"`
	IdaName *string `json:"idaName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CalcLayoutReq CalcLayoutReq

// NewCalcLayoutReq instantiates a new CalcLayoutReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCalcLayoutReq() *CalcLayoutReq {
	this := CalcLayoutReq{}
	return &this
}

// NewCalcLayoutReqWithDefaults instantiates a new CalcLayoutReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCalcLayoutReqWithDefaults() *CalcLayoutReq {
	this := CalcLayoutReq{}
	return &this
}

// GetExport returns the Export field value if set, zero value otherwise.
func (o *CalcLayoutReq) GetExport() bool {
	if o == nil || isNil(o.Export) {
		var ret bool
		return ret
	}
	return *o.Export
}

// GetExportOk returns a tuple with the Export field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalcLayoutReq) GetExportOk() (*bool, bool) {
	if o == nil || isNil(o.Export) {
    return nil, false
	}
	return o.Export, true
}

// HasExport returns a boolean if a field has been set.
func (o *CalcLayoutReq) HasExport() bool {
	if o != nil && !isNil(o.Export) {
		return true
	}

	return false
}

// SetExport gets a reference to the given bool and assigns it to the Export field.
func (o *CalcLayoutReq) SetExport(v bool) {
	o.Export = &v
}

// GetIdOrName returns the IdOrName field value if set, zero value otherwise.
func (o *CalcLayoutReq) GetIdOrName() string {
	if o == nil || isNil(o.IdOrName) {
		var ret string
		return ret
	}
	return *o.IdOrName
}

// GetIdOrNameOk returns a tuple with the IdOrName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalcLayoutReq) GetIdOrNameOk() (*string, bool) {
	if o == nil || isNil(o.IdOrName) {
    return nil, false
	}
	return o.IdOrName, true
}

// HasIdOrName returns a boolean if a field has been set.
func (o *CalcLayoutReq) HasIdOrName() bool {
	if o != nil && !isNil(o.IdOrName) {
		return true
	}

	return false
}

// SetIdOrName gets a reference to the given string and assigns it to the IdOrName field.
func (o *CalcLayoutReq) SetIdOrName(v string) {
	o.IdOrName = &v
}

// GetIdaName returns the IdaName field value if set, zero value otherwise.
func (o *CalcLayoutReq) GetIdaName() string {
	if o == nil || isNil(o.IdaName) {
		var ret string
		return ret
	}
	return *o.IdaName
}

// GetIdaNameOk returns a tuple with the IdaName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalcLayoutReq) GetIdaNameOk() (*string, bool) {
	if o == nil || isNil(o.IdaName) {
    return nil, false
	}
	return o.IdaName, true
}

// HasIdaName returns a boolean if a field has been set.
func (o *CalcLayoutReq) HasIdaName() bool {
	if o != nil && !isNil(o.IdaName) {
		return true
	}

	return false
}

// SetIdaName gets a reference to the given string and assigns it to the IdaName field.
func (o *CalcLayoutReq) SetIdaName(v string) {
	o.IdaName = &v
}

func (o CalcLayoutReq) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Export) {
		toSerialize["export"] = o.Export
	}
	if !isNil(o.IdOrName) {
		toSerialize["idOrName"] = o.IdOrName
	}
	if !isNil(o.IdaName) {
		toSerialize["idaName"] = o.IdaName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CalcLayoutReq) UnmarshalJSON(bytes []byte) (err error) {
	varCalcLayoutReq := _CalcLayoutReq{}

	if err = json.Unmarshal(bytes, &varCalcLayoutReq); err == nil {
		*o = CalcLayoutReq(varCalcLayoutReq)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "export")
		delete(additionalProperties, "idOrName")
		delete(additionalProperties, "idaName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCalcLayoutReq struct {
	value *CalcLayoutReq
	isSet bool
}

func (v NullableCalcLayoutReq) Get() *CalcLayoutReq {
	return v.value
}

func (v *NullableCalcLayoutReq) Set(val *CalcLayoutReq) {
	v.value = val
	v.isSet = true
}

func (v NullableCalcLayoutReq) IsSet() bool {
	return v.isSet
}

func (v *NullableCalcLayoutReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCalcLayoutReq(val *CalcLayoutReq) *NullableCalcLayoutReq {
	return &NullableCalcLayoutReq{value: val, isSet: true}
}

func (v NullableCalcLayoutReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCalcLayoutReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


