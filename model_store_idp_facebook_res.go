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

// StoreIdpFacebookRes struct for StoreIdpFacebookRes
type StoreIdpFacebookRes struct {
	Error *string `json:"error,omitempty"`
	Idp *FacebookOpenIDConnectIdentityProviderDTO `json:"idp,omitempty"`
	ValidationErrors []string `json:"validationErrors,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StoreIdpFacebookRes StoreIdpFacebookRes

// NewStoreIdpFacebookRes instantiates a new StoreIdpFacebookRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoreIdpFacebookRes() *StoreIdpFacebookRes {
	this := StoreIdpFacebookRes{}
	return &this
}

// NewStoreIdpFacebookResWithDefaults instantiates a new StoreIdpFacebookRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoreIdpFacebookResWithDefaults() *StoreIdpFacebookRes {
	this := StoreIdpFacebookRes{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *StoreIdpFacebookRes) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreIdpFacebookRes) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *StoreIdpFacebookRes) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *StoreIdpFacebookRes) SetError(v string) {
	o.Error = &v
}

// GetIdp returns the Idp field value if set, zero value otherwise.
func (o *StoreIdpFacebookRes) GetIdp() FacebookOpenIDConnectIdentityProviderDTO {
	if o == nil || o.Idp == nil {
		var ret FacebookOpenIDConnectIdentityProviderDTO
		return ret
	}
	return *o.Idp
}

// GetIdpOk returns a tuple with the Idp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreIdpFacebookRes) GetIdpOk() (*FacebookOpenIDConnectIdentityProviderDTO, bool) {
	if o == nil || o.Idp == nil {
		return nil, false
	}
	return o.Idp, true
}

// HasIdp returns a boolean if a field has been set.
func (o *StoreIdpFacebookRes) HasIdp() bool {
	if o != nil && o.Idp != nil {
		return true
	}

	return false
}

// SetIdp gets a reference to the given FacebookOpenIDConnectIdentityProviderDTO and assigns it to the Idp field.
func (o *StoreIdpFacebookRes) SetIdp(v FacebookOpenIDConnectIdentityProviderDTO) {
	o.Idp = &v
}

// GetValidationErrors returns the ValidationErrors field value if set, zero value otherwise.
func (o *StoreIdpFacebookRes) GetValidationErrors() []string {
	if o == nil || o.ValidationErrors == nil {
		var ret []string
		return ret
	}
	return o.ValidationErrors
}

// GetValidationErrorsOk returns a tuple with the ValidationErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreIdpFacebookRes) GetValidationErrorsOk() ([]string, bool) {
	if o == nil || o.ValidationErrors == nil {
		return nil, false
	}
	return o.ValidationErrors, true
}

// HasValidationErrors returns a boolean if a field has been set.
func (o *StoreIdpFacebookRes) HasValidationErrors() bool {
	if o != nil && o.ValidationErrors != nil {
		return true
	}

	return false
}

// SetValidationErrors gets a reference to the given []string and assigns it to the ValidationErrors field.
func (o *StoreIdpFacebookRes) SetValidationErrors(v []string) {
	o.ValidationErrors = v
}

func (o StoreIdpFacebookRes) MarshalJSON() ([]byte, error) {
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

func (o *StoreIdpFacebookRes) UnmarshalJSON(bytes []byte) (err error) {
	varStoreIdpFacebookRes := _StoreIdpFacebookRes{}

	if err = json.Unmarshal(bytes, &varStoreIdpFacebookRes); err == nil {
		*o = StoreIdpFacebookRes(varStoreIdpFacebookRes)
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

type NullableStoreIdpFacebookRes struct {
	value *StoreIdpFacebookRes
	isSet bool
}

func (v NullableStoreIdpFacebookRes) Get() *StoreIdpFacebookRes {
	return v.value
}

func (v *NullableStoreIdpFacebookRes) Set(val *StoreIdpFacebookRes) {
	v.value = val
	v.isSet = true
}

func (v NullableStoreIdpFacebookRes) IsSet() bool {
	return v.isSet
}

func (v *NullableStoreIdpFacebookRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoreIdpFacebookRes(val *StoreIdpFacebookRes) *NullableStoreIdpFacebookRes {
	return &NullableStoreIdpFacebookRes{value: val, isSet: true}
}

func (v NullableStoreIdpFacebookRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoreIdpFacebookRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


