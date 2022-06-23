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

// GetIntSaml2SpsRes struct for GetIntSaml2SpsRes
type GetIntSaml2SpsRes struct {
	Error *string `json:"error,omitempty"`
	Sps []InternalSaml2ServiceProviderDTO `json:"sps,omitempty"`
	ValidationErrors []string `json:"validationErrors,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetIntSaml2SpsRes GetIntSaml2SpsRes

// NewGetIntSaml2SpsRes instantiates a new GetIntSaml2SpsRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetIntSaml2SpsRes() *GetIntSaml2SpsRes {
	this := GetIntSaml2SpsRes{}
	return &this
}

// NewGetIntSaml2SpsResWithDefaults instantiates a new GetIntSaml2SpsRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetIntSaml2SpsResWithDefaults() *GetIntSaml2SpsRes {
	this := GetIntSaml2SpsRes{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *GetIntSaml2SpsRes) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIntSaml2SpsRes) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *GetIntSaml2SpsRes) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *GetIntSaml2SpsRes) SetError(v string) {
	o.Error = &v
}

// GetSps returns the Sps field value if set, zero value otherwise.
func (o *GetIntSaml2SpsRes) GetSps() []InternalSaml2ServiceProviderDTO {
	if o == nil || o.Sps == nil {
		var ret []InternalSaml2ServiceProviderDTO
		return ret
	}
	return o.Sps
}

// GetSpsOk returns a tuple with the Sps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIntSaml2SpsRes) GetSpsOk() ([]InternalSaml2ServiceProviderDTO, bool) {
	if o == nil || o.Sps == nil {
		return nil, false
	}
	return o.Sps, true
}

// HasSps returns a boolean if a field has been set.
func (o *GetIntSaml2SpsRes) HasSps() bool {
	if o != nil && o.Sps != nil {
		return true
	}

	return false
}

// SetSps gets a reference to the given []InternalSaml2ServiceProviderDTO and assigns it to the Sps field.
func (o *GetIntSaml2SpsRes) SetSps(v []InternalSaml2ServiceProviderDTO) {
	o.Sps = v
}

// GetValidationErrors returns the ValidationErrors field value if set, zero value otherwise.
func (o *GetIntSaml2SpsRes) GetValidationErrors() []string {
	if o == nil || o.ValidationErrors == nil {
		var ret []string
		return ret
	}
	return o.ValidationErrors
}

// GetValidationErrorsOk returns a tuple with the ValidationErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIntSaml2SpsRes) GetValidationErrorsOk() ([]string, bool) {
	if o == nil || o.ValidationErrors == nil {
		return nil, false
	}
	return o.ValidationErrors, true
}

// HasValidationErrors returns a boolean if a field has been set.
func (o *GetIntSaml2SpsRes) HasValidationErrors() bool {
	if o != nil && o.ValidationErrors != nil {
		return true
	}

	return false
}

// SetValidationErrors gets a reference to the given []string and assigns it to the ValidationErrors field.
func (o *GetIntSaml2SpsRes) SetValidationErrors(v []string) {
	o.ValidationErrors = v
}

func (o GetIntSaml2SpsRes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.Sps != nil {
		toSerialize["sps"] = o.Sps
	}
	if o.ValidationErrors != nil {
		toSerialize["validationErrors"] = o.ValidationErrors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GetIntSaml2SpsRes) UnmarshalJSON(bytes []byte) (err error) {
	varGetIntSaml2SpsRes := _GetIntSaml2SpsRes{}

	if err = json.Unmarshal(bytes, &varGetIntSaml2SpsRes); err == nil {
		*o = GetIntSaml2SpsRes(varGetIntSaml2SpsRes)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "error")
		delete(additionalProperties, "sps")
		delete(additionalProperties, "validationErrors")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetIntSaml2SpsRes struct {
	value *GetIntSaml2SpsRes
	isSet bool
}

func (v NullableGetIntSaml2SpsRes) Get() *GetIntSaml2SpsRes {
	return v.value
}

func (v *NullableGetIntSaml2SpsRes) Set(val *GetIntSaml2SpsRes) {
	v.value = val
	v.isSet = true
}

func (v NullableGetIntSaml2SpsRes) IsSet() bool {
	return v.isSet
}

func (v *NullableGetIntSaml2SpsRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetIntSaml2SpsRes(val *GetIntSaml2SpsRes) *NullableGetIntSaml2SpsRes {
	return &NullableGetIntSaml2SpsRes{value: val, isSet: true}
}

func (v NullableGetIntSaml2SpsRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetIntSaml2SpsRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


