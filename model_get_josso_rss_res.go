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

// GetJossoRssRes struct for GetJossoRssRes
type GetJossoRssRes struct {
	Error *string `json:"error,omitempty"`
	Resources []JOSSO1ResourceDTO `json:"resources,omitempty"`
	ValidationErrors []string `json:"validationErrors,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetJossoRssRes GetJossoRssRes

// NewGetJossoRssRes instantiates a new GetJossoRssRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetJossoRssRes() *GetJossoRssRes {
	this := GetJossoRssRes{}
	return &this
}

// NewGetJossoRssResWithDefaults instantiates a new GetJossoRssRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetJossoRssResWithDefaults() *GetJossoRssRes {
	this := GetJossoRssRes{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *GetJossoRssRes) GetError() string {
	if o == nil || isNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetJossoRssRes) GetErrorOk() (*string, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *GetJossoRssRes) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *GetJossoRssRes) SetError(v string) {
	o.Error = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *GetJossoRssRes) GetResources() []JOSSO1ResourceDTO {
	if o == nil || isNil(o.Resources) {
		var ret []JOSSO1ResourceDTO
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetJossoRssRes) GetResourcesOk() ([]JOSSO1ResourceDTO, bool) {
	if o == nil || isNil(o.Resources) {
    return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *GetJossoRssRes) HasResources() bool {
	if o != nil && !isNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given []JOSSO1ResourceDTO and assigns it to the Resources field.
func (o *GetJossoRssRes) SetResources(v []JOSSO1ResourceDTO) {
	o.Resources = v
}

// GetValidationErrors returns the ValidationErrors field value if set, zero value otherwise.
func (o *GetJossoRssRes) GetValidationErrors() []string {
	if o == nil || isNil(o.ValidationErrors) {
		var ret []string
		return ret
	}
	return o.ValidationErrors
}

// GetValidationErrorsOk returns a tuple with the ValidationErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetJossoRssRes) GetValidationErrorsOk() ([]string, bool) {
	if o == nil || isNil(o.ValidationErrors) {
    return nil, false
	}
	return o.ValidationErrors, true
}

// HasValidationErrors returns a boolean if a field has been set.
func (o *GetJossoRssRes) HasValidationErrors() bool {
	if o != nil && !isNil(o.ValidationErrors) {
		return true
	}

	return false
}

// SetValidationErrors gets a reference to the given []string and assigns it to the ValidationErrors field.
func (o *GetJossoRssRes) SetValidationErrors(v []string) {
	o.ValidationErrors = v
}

func (o GetJossoRssRes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !isNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}
	if !isNil(o.ValidationErrors) {
		toSerialize["validationErrors"] = o.ValidationErrors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GetJossoRssRes) UnmarshalJSON(bytes []byte) (err error) {
	varGetJossoRssRes := _GetJossoRssRes{}

	if err = json.Unmarshal(bytes, &varGetJossoRssRes); err == nil {
		*o = GetJossoRssRes(varGetJossoRssRes)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "error")
		delete(additionalProperties, "resources")
		delete(additionalProperties, "validationErrors")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetJossoRssRes struct {
	value *GetJossoRssRes
	isSet bool
}

func (v NullableGetJossoRssRes) Get() *GetJossoRssRes {
	return v.value
}

func (v *NullableGetJossoRssRes) Set(val *GetJossoRssRes) {
	v.value = val
	v.isSet = true
}

func (v NullableGetJossoRssRes) IsSet() bool {
	return v.isSet
}

func (v *NullableGetJossoRssRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetJossoRssRes(val *GetJossoRssRes) *NullableGetJossoRssRes {
	return &NullableGetJossoRssRes{value: val, isSet: true}
}

func (v NullableGetJossoRssRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetJossoRssRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


