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

// GetIdpFacebookRes struct for GetIdpFacebookRes
type GetIdpFacebookRes struct {
	Error *string `json:"error,omitempty"`
	Idp *FacebookOpenIDConnectIdentityProviderDTO `json:"idp,omitempty"`
	ValidationErrors []string `json:"validationErrors,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetIdpFacebookRes GetIdpFacebookRes

// NewGetIdpFacebookRes instantiates a new GetIdpFacebookRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetIdpFacebookRes() *GetIdpFacebookRes {
	this := GetIdpFacebookRes{}
	return &this
}

// NewGetIdpFacebookResWithDefaults instantiates a new GetIdpFacebookRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetIdpFacebookResWithDefaults() *GetIdpFacebookRes {
	this := GetIdpFacebookRes{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *GetIdpFacebookRes) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIdpFacebookRes) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *GetIdpFacebookRes) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *GetIdpFacebookRes) SetError(v string) {
	o.Error = &v
}

// GetIdp returns the Idp field value if set, zero value otherwise.
func (o *GetIdpFacebookRes) GetIdp() FacebookOpenIDConnectIdentityProviderDTO {
	if o == nil || o.Idp == nil {
		var ret FacebookOpenIDConnectIdentityProviderDTO
		return ret
	}
	return *o.Idp
}

// GetIdpOk returns a tuple with the Idp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIdpFacebookRes) GetIdpOk() (*FacebookOpenIDConnectIdentityProviderDTO, bool) {
	if o == nil || o.Idp == nil {
		return nil, false
	}
	return o.Idp, true
}

// HasIdp returns a boolean if a field has been set.
func (o *GetIdpFacebookRes) HasIdp() bool {
	if o != nil && o.Idp != nil {
		return true
	}

	return false
}

// SetIdp gets a reference to the given FacebookOpenIDConnectIdentityProviderDTO and assigns it to the Idp field.
func (o *GetIdpFacebookRes) SetIdp(v FacebookOpenIDConnectIdentityProviderDTO) {
	o.Idp = &v
}

// GetValidationErrors returns the ValidationErrors field value if set, zero value otherwise.
func (o *GetIdpFacebookRes) GetValidationErrors() []string {
	if o == nil || o.ValidationErrors == nil {
		var ret []string
		return ret
	}
	return o.ValidationErrors
}

// GetValidationErrorsOk returns a tuple with the ValidationErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIdpFacebookRes) GetValidationErrorsOk() ([]string, bool) {
	if o == nil || o.ValidationErrors == nil {
		return nil, false
	}
	return o.ValidationErrors, true
}

// HasValidationErrors returns a boolean if a field has been set.
func (o *GetIdpFacebookRes) HasValidationErrors() bool {
	if o != nil && o.ValidationErrors != nil {
		return true
	}

	return false
}

// SetValidationErrors gets a reference to the given []string and assigns it to the ValidationErrors field.
func (o *GetIdpFacebookRes) SetValidationErrors(v []string) {
	o.ValidationErrors = v
}

func (o GetIdpFacebookRes) MarshalJSON() ([]byte, error) {
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

func (o *GetIdpFacebookRes) UnmarshalJSON(bytes []byte) (err error) {
	varGetIdpFacebookRes := _GetIdpFacebookRes{}

	if err = json.Unmarshal(bytes, &varGetIdpFacebookRes); err == nil {
		*o = GetIdpFacebookRes(varGetIdpFacebookRes)
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

type NullableGetIdpFacebookRes struct {
	value *GetIdpFacebookRes
	isSet bool
}

func (v NullableGetIdpFacebookRes) Get() *GetIdpFacebookRes {
	return v.value
}

func (v *NullableGetIdpFacebookRes) Set(val *GetIdpFacebookRes) {
	v.value = val
	v.isSet = true
}

func (v NullableGetIdpFacebookRes) IsSet() bool {
	return v.isSet
}

func (v *NullableGetIdpFacebookRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetIdpFacebookRes(val *GetIdpFacebookRes) *NullableGetIdpFacebookRes {
	return &NullableGetIdpFacebookRes{value: val, isSet: true}
}

func (v NullableGetIdpFacebookRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetIdpFacebookRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

