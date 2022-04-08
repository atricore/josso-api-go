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

// GetIdSourceLdapRes struct for GetIdSourceLdapRes
type GetIdSourceLdapRes struct {
	Error *string `json:"error,omitempty"`
	IdSourceLdap *LdapIdentitySourceDTO `json:"idSourceLdap,omitempty"`
	ValidationErrors []string `json:"validationErrors,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetIdSourceLdapRes GetIdSourceLdapRes

// NewGetIdSourceLdapRes instantiates a new GetIdSourceLdapRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetIdSourceLdapRes() *GetIdSourceLdapRes {
	this := GetIdSourceLdapRes{}
	return &this
}

// NewGetIdSourceLdapResWithDefaults instantiates a new GetIdSourceLdapRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetIdSourceLdapResWithDefaults() *GetIdSourceLdapRes {
	this := GetIdSourceLdapRes{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *GetIdSourceLdapRes) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIdSourceLdapRes) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *GetIdSourceLdapRes) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *GetIdSourceLdapRes) SetError(v string) {
	o.Error = &v
}

// GetIdSourceLdap returns the IdSourceLdap field value if set, zero value otherwise.
func (o *GetIdSourceLdapRes) GetIdSourceLdap() LdapIdentitySourceDTO {
	if o == nil || o.IdSourceLdap == nil {
		var ret LdapIdentitySourceDTO
		return ret
	}
	return *o.IdSourceLdap
}

// GetIdSourceLdapOk returns a tuple with the IdSourceLdap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIdSourceLdapRes) GetIdSourceLdapOk() (*LdapIdentitySourceDTO, bool) {
	if o == nil || o.IdSourceLdap == nil {
		return nil, false
	}
	return o.IdSourceLdap, true
}

// HasIdSourceLdap returns a boolean if a field has been set.
func (o *GetIdSourceLdapRes) HasIdSourceLdap() bool {
	if o != nil && o.IdSourceLdap != nil {
		return true
	}

	return false
}

// SetIdSourceLdap gets a reference to the given LdapIdentitySourceDTO and assigns it to the IdSourceLdap field.
func (o *GetIdSourceLdapRes) SetIdSourceLdap(v LdapIdentitySourceDTO) {
	o.IdSourceLdap = &v
}

// GetValidationErrors returns the ValidationErrors field value if set, zero value otherwise.
func (o *GetIdSourceLdapRes) GetValidationErrors() []string {
	if o == nil || o.ValidationErrors == nil {
		var ret []string
		return ret
	}
	return o.ValidationErrors
}

// GetValidationErrorsOk returns a tuple with the ValidationErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIdSourceLdapRes) GetValidationErrorsOk() ([]string, bool) {
	if o == nil || o.ValidationErrors == nil {
		return nil, false
	}
	return o.ValidationErrors, true
}

// HasValidationErrors returns a boolean if a field has been set.
func (o *GetIdSourceLdapRes) HasValidationErrors() bool {
	if o != nil && o.ValidationErrors != nil {
		return true
	}

	return false
}

// SetValidationErrors gets a reference to the given []string and assigns it to the ValidationErrors field.
func (o *GetIdSourceLdapRes) SetValidationErrors(v []string) {
	o.ValidationErrors = v
}

func (o GetIdSourceLdapRes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.IdSourceLdap != nil {
		toSerialize["idSourceLdap"] = o.IdSourceLdap
	}
	if o.ValidationErrors != nil {
		toSerialize["validationErrors"] = o.ValidationErrors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GetIdSourceLdapRes) UnmarshalJSON(bytes []byte) (err error) {
	varGetIdSourceLdapRes := _GetIdSourceLdapRes{}

	if err = json.Unmarshal(bytes, &varGetIdSourceLdapRes); err == nil {
		*o = GetIdSourceLdapRes(varGetIdSourceLdapRes)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "error")
		delete(additionalProperties, "idSourceLdap")
		delete(additionalProperties, "validationErrors")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetIdSourceLdapRes struct {
	value *GetIdSourceLdapRes
	isSet bool
}

func (v NullableGetIdSourceLdapRes) Get() *GetIdSourceLdapRes {
	return v.value
}

func (v *NullableGetIdSourceLdapRes) Set(val *GetIdSourceLdapRes) {
	v.value = val
	v.isSet = true
}

func (v NullableGetIdSourceLdapRes) IsSet() bool {
	return v.isSet
}

func (v *NullableGetIdSourceLdapRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetIdSourceLdapRes(val *GetIdSourceLdapRes) *NullableGetIdSourceLdapRes {
	return &NullableGetIdSourceLdapRes{value: val, isSet: true}
}

func (v NullableGetIdSourceLdapRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetIdSourceLdapRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


