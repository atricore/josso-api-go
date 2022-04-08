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

// GetJossoRsRes struct for GetJossoRsRes
type GetJossoRsRes struct {
	Error *string `json:"error,omitempty"`
	Resource *JOSSO1ResourceDTO `json:"resource,omitempty"`
	ValidationErrors *[]string `json:"validationErrors,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetJossoRsRes GetJossoRsRes

// NewGetJossoRsRes instantiates a new GetJossoRsRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetJossoRsRes() *GetJossoRsRes {
	this := GetJossoRsRes{}
	return &this
}

// NewGetJossoRsResWithDefaults instantiates a new GetJossoRsRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetJossoRsResWithDefaults() *GetJossoRsRes {
	this := GetJossoRsRes{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *GetJossoRsRes) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetJossoRsRes) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *GetJossoRsRes) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *GetJossoRsRes) SetError(v string) {
	o.Error = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *GetJossoRsRes) GetResource() JOSSO1ResourceDTO {
	if o == nil || o.Resource == nil {
		var ret JOSSO1ResourceDTO
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetJossoRsRes) GetResourceOk() (*JOSSO1ResourceDTO, bool) {
	if o == nil || o.Resource == nil {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *GetJossoRsRes) HasResource() bool {
	if o != nil && o.Resource != nil {
		return true
	}

	return false
}

// SetResource gets a reference to the given JOSSO1ResourceDTO and assigns it to the Resource field.
func (o *GetJossoRsRes) SetResource(v JOSSO1ResourceDTO) {
	o.Resource = &v
}

// GetValidationErrors returns the ValidationErrors field value if set, zero value otherwise.
func (o *GetJossoRsRes) GetValidationErrors() []string {
	if o == nil || o.ValidationErrors == nil {
		var ret []string
		return ret
	}
	return *o.ValidationErrors
}

// GetValidationErrorsOk returns a tuple with the ValidationErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetJossoRsRes) GetValidationErrorsOk() (*[]string, bool) {
	if o == nil || o.ValidationErrors == nil {
		return nil, false
	}
	return o.ValidationErrors, true
}

// HasValidationErrors returns a boolean if a field has been set.
func (o *GetJossoRsRes) HasValidationErrors() bool {
	if o != nil && o.ValidationErrors != nil {
		return true
	}

	return false
}

// SetValidationErrors gets a reference to the given []string and assigns it to the ValidationErrors field.
func (o *GetJossoRsRes) SetValidationErrors(v []string) {
	o.ValidationErrors = &v
}

func (o GetJossoRsRes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.Resource != nil {
		toSerialize["resource"] = o.Resource
	}
	if o.ValidationErrors != nil {
		toSerialize["validationErrors"] = o.ValidationErrors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GetJossoRsRes) UnmarshalJSON(bytes []byte) (err error) {
	varGetJossoRsRes := _GetJossoRsRes{}

	if err = json.Unmarshal(bytes, &varGetJossoRsRes); err == nil {
		*o = GetJossoRsRes(varGetJossoRsRes)
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

type NullableGetJossoRsRes struct {
	value *GetJossoRsRes
	isSet bool
}

func (v NullableGetJossoRsRes) Get() *GetJossoRsRes {
	return v.value
}

func (v *NullableGetJossoRsRes) Set(val *GetJossoRsRes) {
	v.value = val
	v.isSet = true
}

func (v NullableGetJossoRsRes) IsSet() bool {
	return v.isSet
}

func (v *NullableGetJossoRsRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetJossoRsRes(val *GetJossoRsRes) *NullableGetJossoRsRes {
	return &NullableGetJossoRsRes{value: val, isSet: true}
}

func (v NullableGetJossoRsRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetJossoRsRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

