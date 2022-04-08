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

// GetApplianceStateRes struct for GetApplianceStateRes
type GetApplianceStateRes struct {
	Error *string `json:"error,omitempty"`
	State *string `json:"state,omitempty"`
	ValidationErrors []string `json:"validationErrors,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetApplianceStateRes GetApplianceStateRes

// NewGetApplianceStateRes instantiates a new GetApplianceStateRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetApplianceStateRes() *GetApplianceStateRes {
	this := GetApplianceStateRes{}
	return &this
}

// NewGetApplianceStateResWithDefaults instantiates a new GetApplianceStateRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetApplianceStateResWithDefaults() *GetApplianceStateRes {
	this := GetApplianceStateRes{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *GetApplianceStateRes) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetApplianceStateRes) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *GetApplianceStateRes) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *GetApplianceStateRes) SetError(v string) {
	o.Error = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *GetApplianceStateRes) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetApplianceStateRes) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *GetApplianceStateRes) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *GetApplianceStateRes) SetState(v string) {
	o.State = &v
}

// GetValidationErrors returns the ValidationErrors field value if set, zero value otherwise.
func (o *GetApplianceStateRes) GetValidationErrors() []string {
	if o == nil || o.ValidationErrors == nil {
		var ret []string
		return ret
	}
	return o.ValidationErrors
}

// GetValidationErrorsOk returns a tuple with the ValidationErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetApplianceStateRes) GetValidationErrorsOk() ([]string, bool) {
	if o == nil || o.ValidationErrors == nil {
		return nil, false
	}
	return o.ValidationErrors, true
}

// HasValidationErrors returns a boolean if a field has been set.
func (o *GetApplianceStateRes) HasValidationErrors() bool {
	if o != nil && o.ValidationErrors != nil {
		return true
	}

	return false
}

// SetValidationErrors gets a reference to the given []string and assigns it to the ValidationErrors field.
func (o *GetApplianceStateRes) SetValidationErrors(v []string) {
	o.ValidationErrors = v
}

func (o GetApplianceStateRes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.ValidationErrors != nil {
		toSerialize["validationErrors"] = o.ValidationErrors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GetApplianceStateRes) UnmarshalJSON(bytes []byte) (err error) {
	varGetApplianceStateRes := _GetApplianceStateRes{}

	if err = json.Unmarshal(bytes, &varGetApplianceStateRes); err == nil {
		*o = GetApplianceStateRes(varGetApplianceStateRes)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "error")
		delete(additionalProperties, "state")
		delete(additionalProperties, "validationErrors")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetApplianceStateRes struct {
	value *GetApplianceStateRes
	isSet bool
}

func (v NullableGetApplianceStateRes) Get() *GetApplianceStateRes {
	return v.value
}

func (v *NullableGetApplianceStateRes) Set(val *GetApplianceStateRes) {
	v.value = val
	v.isSet = true
}

func (v NullableGetApplianceStateRes) IsSet() bool {
	return v.isSet
}

func (v *NullableGetApplianceStateRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetApplianceStateRes(val *GetApplianceStateRes) *NullableGetApplianceStateRes {
	return &NullableGetApplianceStateRes{value: val, isSet: true}
}

func (v NullableGetApplianceStateRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetApplianceStateRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


