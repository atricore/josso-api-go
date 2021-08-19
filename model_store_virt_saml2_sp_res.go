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

// StoreVirtSaml2SpRes struct for StoreVirtSaml2SpRes
type StoreVirtSaml2SpRes struct {
	Error *string `json:"error,omitempty"`
	Sp *VirtualSaml2ServiceProviderDTO `json:"sp,omitempty"`
	ValidationErrors *[]string `json:"validationErrors,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StoreVirtSaml2SpRes StoreVirtSaml2SpRes

// NewStoreVirtSaml2SpRes instantiates a new StoreVirtSaml2SpRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoreVirtSaml2SpRes() *StoreVirtSaml2SpRes {
	this := StoreVirtSaml2SpRes{}
	return &this
}

// NewStoreVirtSaml2SpResWithDefaults instantiates a new StoreVirtSaml2SpRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoreVirtSaml2SpResWithDefaults() *StoreVirtSaml2SpRes {
	this := StoreVirtSaml2SpRes{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *StoreVirtSaml2SpRes) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreVirtSaml2SpRes) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *StoreVirtSaml2SpRes) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *StoreVirtSaml2SpRes) SetError(v string) {
	o.Error = &v
}

// GetSp returns the Sp field value if set, zero value otherwise.
func (o *StoreVirtSaml2SpRes) GetSp() VirtualSaml2ServiceProviderDTO {
	if o == nil || o.Sp == nil {
		var ret VirtualSaml2ServiceProviderDTO
		return ret
	}
	return *o.Sp
}

// GetSpOk returns a tuple with the Sp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreVirtSaml2SpRes) GetSpOk() (*VirtualSaml2ServiceProviderDTO, bool) {
	if o == nil || o.Sp == nil {
		return nil, false
	}
	return o.Sp, true
}

// HasSp returns a boolean if a field has been set.
func (o *StoreVirtSaml2SpRes) HasSp() bool {
	if o != nil && o.Sp != nil {
		return true
	}

	return false
}

// SetSp gets a reference to the given VirtualSaml2ServiceProviderDTO and assigns it to the Sp field.
func (o *StoreVirtSaml2SpRes) SetSp(v VirtualSaml2ServiceProviderDTO) {
	o.Sp = &v
}

// GetValidationErrors returns the ValidationErrors field value if set, zero value otherwise.
func (o *StoreVirtSaml2SpRes) GetValidationErrors() []string {
	if o == nil || o.ValidationErrors == nil {
		var ret []string
		return ret
	}
	return *o.ValidationErrors
}

// GetValidationErrorsOk returns a tuple with the ValidationErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreVirtSaml2SpRes) GetValidationErrorsOk() (*[]string, bool) {
	if o == nil || o.ValidationErrors == nil {
		return nil, false
	}
	return o.ValidationErrors, true
}

// HasValidationErrors returns a boolean if a field has been set.
func (o *StoreVirtSaml2SpRes) HasValidationErrors() bool {
	if o != nil && o.ValidationErrors != nil {
		return true
	}

	return false
}

// SetValidationErrors gets a reference to the given []string and assigns it to the ValidationErrors field.
func (o *StoreVirtSaml2SpRes) SetValidationErrors(v []string) {
	o.ValidationErrors = &v
}

func (o StoreVirtSaml2SpRes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.Sp != nil {
		toSerialize["sp"] = o.Sp
	}
	if o.ValidationErrors != nil {
		toSerialize["validationErrors"] = o.ValidationErrors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StoreVirtSaml2SpRes) UnmarshalJSON(bytes []byte) (err error) {
	varStoreVirtSaml2SpRes := _StoreVirtSaml2SpRes{}

	if err = json.Unmarshal(bytes, &varStoreVirtSaml2SpRes); err == nil {
		*o = StoreVirtSaml2SpRes(varStoreVirtSaml2SpRes)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "error")
		delete(additionalProperties, "sp")
		delete(additionalProperties, "validationErrors")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStoreVirtSaml2SpRes struct {
	value *StoreVirtSaml2SpRes
	isSet bool
}

func (v NullableStoreVirtSaml2SpRes) Get() *StoreVirtSaml2SpRes {
	return v.value
}

func (v *NullableStoreVirtSaml2SpRes) Set(val *StoreVirtSaml2SpRes) {
	v.value = val
	v.isSet = true
}

func (v NullableStoreVirtSaml2SpRes) IsSet() bool {
	return v.isSet
}

func (v *NullableStoreVirtSaml2SpRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoreVirtSaml2SpRes(val *StoreVirtSaml2SpRes) *NullableStoreVirtSaml2SpRes {
	return &NullableStoreVirtSaml2SpRes{value: val, isSet: true}
}

func (v NullableStoreVirtSaml2SpRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoreVirtSaml2SpRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


