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

// GetIdPsRes struct for GetIdPsRes
type GetIdPsRes struct {
	Error *string `json:"error,omitempty"`
	Idps *[]IdentityProviderDTO `json:"idps,omitempty"`
	ValidationErrors *[]string `json:"validationErrors,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetIdPsRes GetIdPsRes

// NewGetIdPsRes instantiates a new GetIdPsRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetIdPsRes() *GetIdPsRes {
	this := GetIdPsRes{}
	return &this
}

// NewGetIdPsResWithDefaults instantiates a new GetIdPsRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetIdPsResWithDefaults() *GetIdPsRes {
	this := GetIdPsRes{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *GetIdPsRes) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIdPsRes) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *GetIdPsRes) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *GetIdPsRes) SetError(v string) {
	o.Error = &v
}

// GetIdps returns the Idps field value if set, zero value otherwise.
func (o *GetIdPsRes) GetIdps() []IdentityProviderDTO {
	if o == nil || o.Idps == nil {
		var ret []IdentityProviderDTO
		return ret
	}
	return *o.Idps
}

// GetIdpsOk returns a tuple with the Idps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIdPsRes) GetIdpsOk() (*[]IdentityProviderDTO, bool) {
	if o == nil || o.Idps == nil {
		return nil, false
	}
	return o.Idps, true
}

// HasIdps returns a boolean if a field has been set.
func (o *GetIdPsRes) HasIdps() bool {
	if o != nil && o.Idps != nil {
		return true
	}

	return false
}

// SetIdps gets a reference to the given []IdentityProviderDTO and assigns it to the Idps field.
func (o *GetIdPsRes) SetIdps(v []IdentityProviderDTO) {
	o.Idps = &v
}

// GetValidationErrors returns the ValidationErrors field value if set, zero value otherwise.
func (o *GetIdPsRes) GetValidationErrors() []string {
	if o == nil || o.ValidationErrors == nil {
		var ret []string
		return ret
	}
	return *o.ValidationErrors
}

// GetValidationErrorsOk returns a tuple with the ValidationErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIdPsRes) GetValidationErrorsOk() (*[]string, bool) {
	if o == nil || o.ValidationErrors == nil {
		return nil, false
	}
	return o.ValidationErrors, true
}

// HasValidationErrors returns a boolean if a field has been set.
func (o *GetIdPsRes) HasValidationErrors() bool {
	if o != nil && o.ValidationErrors != nil {
		return true
	}

	return false
}

// SetValidationErrors gets a reference to the given []string and assigns it to the ValidationErrors field.
func (o *GetIdPsRes) SetValidationErrors(v []string) {
	o.ValidationErrors = &v
}

func (o GetIdPsRes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.Idps != nil {
		toSerialize["idps"] = o.Idps
	}
	if o.ValidationErrors != nil {
		toSerialize["validationErrors"] = o.ValidationErrors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GetIdPsRes) UnmarshalJSON(bytes []byte) (err error) {
	varGetIdPsRes := _GetIdPsRes{}

	if err = json.Unmarshal(bytes, &varGetIdPsRes); err == nil {
		*o = GetIdPsRes(varGetIdPsRes)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "error")
		delete(additionalProperties, "idps")
		delete(additionalProperties, "validationErrors")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetIdPsRes struct {
	value *GetIdPsRes
	isSet bool
}

func (v NullableGetIdPsRes) Get() *GetIdPsRes {
	return v.value
}

func (v *NullableGetIdPsRes) Set(val *GetIdPsRes) {
	v.value = val
	v.isSet = true
}

func (v NullableGetIdPsRes) IsSet() bool {
	return v.isSet
}

func (v *NullableGetIdPsRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetIdPsRes(val *GetIdPsRes) *NullableGetIdPsRes {
	return &NullableGetIdPsRes{value: val, isSet: true}
}

func (v NullableGetIdPsRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetIdPsRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


