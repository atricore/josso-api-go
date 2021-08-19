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

// GetDbIdVaultsRes struct for GetDbIdVaultsRes
type GetDbIdVaultsRes struct {
	DbIdVaults *[]DbIdentityVaultDTO `json:"dbIdVaults,omitempty"`
	Error *string `json:"error,omitempty"`
	ValidationErrors *[]string `json:"validationErrors,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetDbIdVaultsRes GetDbIdVaultsRes

// NewGetDbIdVaultsRes instantiates a new GetDbIdVaultsRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDbIdVaultsRes() *GetDbIdVaultsRes {
	this := GetDbIdVaultsRes{}
	return &this
}

// NewGetDbIdVaultsResWithDefaults instantiates a new GetDbIdVaultsRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDbIdVaultsResWithDefaults() *GetDbIdVaultsRes {
	this := GetDbIdVaultsRes{}
	return &this
}

// GetDbIdVaults returns the DbIdVaults field value if set, zero value otherwise.
func (o *GetDbIdVaultsRes) GetDbIdVaults() []DbIdentityVaultDTO {
	if o == nil || o.DbIdVaults == nil {
		var ret []DbIdentityVaultDTO
		return ret
	}
	return *o.DbIdVaults
}

// GetDbIdVaultsOk returns a tuple with the DbIdVaults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDbIdVaultsRes) GetDbIdVaultsOk() (*[]DbIdentityVaultDTO, bool) {
	if o == nil || o.DbIdVaults == nil {
		return nil, false
	}
	return o.DbIdVaults, true
}

// HasDbIdVaults returns a boolean if a field has been set.
func (o *GetDbIdVaultsRes) HasDbIdVaults() bool {
	if o != nil && o.DbIdVaults != nil {
		return true
	}

	return false
}

// SetDbIdVaults gets a reference to the given []DbIdentityVaultDTO and assigns it to the DbIdVaults field.
func (o *GetDbIdVaultsRes) SetDbIdVaults(v []DbIdentityVaultDTO) {
	o.DbIdVaults = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *GetDbIdVaultsRes) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDbIdVaultsRes) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *GetDbIdVaultsRes) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *GetDbIdVaultsRes) SetError(v string) {
	o.Error = &v
}

// GetValidationErrors returns the ValidationErrors field value if set, zero value otherwise.
func (o *GetDbIdVaultsRes) GetValidationErrors() []string {
	if o == nil || o.ValidationErrors == nil {
		var ret []string
		return ret
	}
	return *o.ValidationErrors
}

// GetValidationErrorsOk returns a tuple with the ValidationErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDbIdVaultsRes) GetValidationErrorsOk() (*[]string, bool) {
	if o == nil || o.ValidationErrors == nil {
		return nil, false
	}
	return o.ValidationErrors, true
}

// HasValidationErrors returns a boolean if a field has been set.
func (o *GetDbIdVaultsRes) HasValidationErrors() bool {
	if o != nil && o.ValidationErrors != nil {
		return true
	}

	return false
}

// SetValidationErrors gets a reference to the given []string and assigns it to the ValidationErrors field.
func (o *GetDbIdVaultsRes) SetValidationErrors(v []string) {
	o.ValidationErrors = &v
}

func (o GetDbIdVaultsRes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DbIdVaults != nil {
		toSerialize["dbIdVaults"] = o.DbIdVaults
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.ValidationErrors != nil {
		toSerialize["validationErrors"] = o.ValidationErrors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GetDbIdVaultsRes) UnmarshalJSON(bytes []byte) (err error) {
	varGetDbIdVaultsRes := _GetDbIdVaultsRes{}

	if err = json.Unmarshal(bytes, &varGetDbIdVaultsRes); err == nil {
		*o = GetDbIdVaultsRes(varGetDbIdVaultsRes)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "dbIdVaults")
		delete(additionalProperties, "error")
		delete(additionalProperties, "validationErrors")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetDbIdVaultsRes struct {
	value *GetDbIdVaultsRes
	isSet bool
}

func (v NullableGetDbIdVaultsRes) Get() *GetDbIdVaultsRes {
	return v.value
}

func (v *NullableGetDbIdVaultsRes) Set(val *GetDbIdVaultsRes) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDbIdVaultsRes) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDbIdVaultsRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDbIdVaultsRes(val *GetDbIdVaultsRes) *NullableGetDbIdVaultsRes {
	return &NullableGetDbIdVaultsRes{value: val, isSet: true}
}

func (v NullableGetDbIdVaultsRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDbIdVaultsRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

