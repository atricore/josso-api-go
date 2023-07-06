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

// GetSharepointRsRes struct for GetSharepointRsRes
type GetSharepointRsRes struct {
	Error *string `json:"error,omitempty"`
	Resource *SharepointResourceDTO `json:"resource,omitempty"`
	ValidationErrors []string `json:"validationErrors,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetSharepointRsRes GetSharepointRsRes

// NewGetSharepointRsRes instantiates a new GetSharepointRsRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSharepointRsRes() *GetSharepointRsRes {
	this := GetSharepointRsRes{}
	return &this
}

// NewGetSharepointRsResWithDefaults instantiates a new GetSharepointRsRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSharepointRsResWithDefaults() *GetSharepointRsRes {
	this := GetSharepointRsRes{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *GetSharepointRsRes) GetError() string {
	if o == nil || isNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSharepointRsRes) GetErrorOk() (*string, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *GetSharepointRsRes) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *GetSharepointRsRes) SetError(v string) {
	o.Error = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *GetSharepointRsRes) GetResource() SharepointResourceDTO {
	if o == nil || isNil(o.Resource) {
		var ret SharepointResourceDTO
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSharepointRsRes) GetResourceOk() (*SharepointResourceDTO, bool) {
	if o == nil || isNil(o.Resource) {
    return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *GetSharepointRsRes) HasResource() bool {
	if o != nil && !isNil(o.Resource) {
		return true
	}

	return false
}

// SetResource gets a reference to the given SharepointResourceDTO and assigns it to the Resource field.
func (o *GetSharepointRsRes) SetResource(v SharepointResourceDTO) {
	o.Resource = &v
}

// GetValidationErrors returns the ValidationErrors field value if set, zero value otherwise.
func (o *GetSharepointRsRes) GetValidationErrors() []string {
	if o == nil || isNil(o.ValidationErrors) {
		var ret []string
		return ret
	}
	return o.ValidationErrors
}

// GetValidationErrorsOk returns a tuple with the ValidationErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSharepointRsRes) GetValidationErrorsOk() ([]string, bool) {
	if o == nil || isNil(o.ValidationErrors) {
    return nil, false
	}
	return o.ValidationErrors, true
}

// HasValidationErrors returns a boolean if a field has been set.
func (o *GetSharepointRsRes) HasValidationErrors() bool {
	if o != nil && !isNil(o.ValidationErrors) {
		return true
	}

	return false
}

// SetValidationErrors gets a reference to the given []string and assigns it to the ValidationErrors field.
func (o *GetSharepointRsRes) SetValidationErrors(v []string) {
	o.ValidationErrors = v
}

func (o GetSharepointRsRes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !isNil(o.Resource) {
		toSerialize["resource"] = o.Resource
	}
	if !isNil(o.ValidationErrors) {
		toSerialize["validationErrors"] = o.ValidationErrors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GetSharepointRsRes) UnmarshalJSON(bytes []byte) (err error) {
	varGetSharepointRsRes := _GetSharepointRsRes{}

	if err = json.Unmarshal(bytes, &varGetSharepointRsRes); err == nil {
		*o = GetSharepointRsRes(varGetSharepointRsRes)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "error")
		delete(additionalProperties, "resource")
		delete(additionalProperties, "validationErrors")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetSharepointRsRes struct {
	value *GetSharepointRsRes
	isSet bool
}

func (v NullableGetSharepointRsRes) Get() *GetSharepointRsRes {
	return v.value
}

func (v *NullableGetSharepointRsRes) Set(val *GetSharepointRsRes) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSharepointRsRes) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSharepointRsRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSharepointRsRes(val *GetSharepointRsRes) *NullableGetSharepointRsRes {
	return &NullableGetSharepointRsRes{value: val, isSet: true}
}

func (v NullableGetSharepointRsRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSharepointRsRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


