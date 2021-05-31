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

// StoreIdPRes struct for StoreIdPRes
type StoreIdPRes struct {
	Error *string `json:"error,omitempty"`
	Idp *IdentityProviderDTO `json:"idp,omitempty"`
	ValidationErrors *[]string `json:"validationErrors,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StoreIdPRes StoreIdPRes

// NewStoreIdPRes instantiates a new StoreIdPRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoreIdPRes() *StoreIdPRes {
	this := StoreIdPRes{}
	return &this
}

// NewStoreIdPResWithDefaults instantiates a new StoreIdPRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoreIdPResWithDefaults() *StoreIdPRes {
	this := StoreIdPRes{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *StoreIdPRes) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreIdPRes) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *StoreIdPRes) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *StoreIdPRes) SetError(v string) {
	o.Error = &v
}

// GetIdp returns the Idp field value if set, zero value otherwise.
func (o *StoreIdPRes) GetIdp() IdentityProviderDTO {
	if o == nil || o.Idp == nil {
		var ret IdentityProviderDTO
		return ret
	}
	return *o.Idp
}

// GetIdpOk returns a tuple with the Idp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreIdPRes) GetIdpOk() (*IdentityProviderDTO, bool) {
	if o == nil || o.Idp == nil {
		return nil, false
	}
	return o.Idp, true
}

// HasIdp returns a boolean if a field has been set.
func (o *StoreIdPRes) HasIdp() bool {
	if o != nil && o.Idp != nil {
		return true
	}

	return false
}

// SetIdp gets a reference to the given IdentityProviderDTO and assigns it to the Idp field.
func (o *StoreIdPRes) SetIdp(v IdentityProviderDTO) {
	o.Idp = &v
}

// GetValidationErrors returns the ValidationErrors field value if set, zero value otherwise.
func (o *StoreIdPRes) GetValidationErrors() []string {
	if o == nil || o.ValidationErrors == nil {
		var ret []string
		return ret
	}
	return *o.ValidationErrors
}

// GetValidationErrorsOk returns a tuple with the ValidationErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreIdPRes) GetValidationErrorsOk() (*[]string, bool) {
	if o == nil || o.ValidationErrors == nil {
		return nil, false
	}
	return o.ValidationErrors, true
}

// HasValidationErrors returns a boolean if a field has been set.
func (o *StoreIdPRes) HasValidationErrors() bool {
	if o != nil && o.ValidationErrors != nil {
		return true
	}

	return false
}

// SetValidationErrors gets a reference to the given []string and assigns it to the ValidationErrors field.
func (o *StoreIdPRes) SetValidationErrors(v []string) {
	o.ValidationErrors = &v
}

func (o StoreIdPRes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.Idp != nil {
		toSerialize["idp"] = o.Idp
	}
	if o.ValidationErrors != nil {
		toSerialize["validationErrors"] = o.ValidationErrors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StoreIdPRes) UnmarshalJSON(bytes []byte) (err error) {
	varStoreIdPRes := _StoreIdPRes{}

	if err = json.Unmarshal(bytes, &varStoreIdPRes); err == nil {
		*o = StoreIdPRes(varStoreIdPRes)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "error")
		delete(additionalProperties, "idp")
		delete(additionalProperties, "validationErrors")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStoreIdPRes struct {
	value *StoreIdPRes
	isSet bool
}

func (v NullableStoreIdPRes) Get() *StoreIdPRes {
	return v.value
}

func (v *NullableStoreIdPRes) Set(val *StoreIdPRes) {
	v.value = val
	v.isSet = true
}

func (v NullableStoreIdPRes) IsSet() bool {
	return v.isSet
}

func (v *NullableStoreIdPRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoreIdPRes(val *StoreIdPRes) *NullableStoreIdPRes {
	return &NullableStoreIdPRes{value: val, isSet: true}
}

func (v NullableStoreIdPRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoreIdPRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

