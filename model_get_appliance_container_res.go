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

// GetApplianceContainerRes struct for GetApplianceContainerRes
type GetApplianceContainerRes struct {
	Appliance *IdentityApplianceContainerDTO `json:"appliance,omitempty"`
	Error *string `json:"error,omitempty"`
	ValidationErrors []string `json:"validationErrors,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetApplianceContainerRes GetApplianceContainerRes

// NewGetApplianceContainerRes instantiates a new GetApplianceContainerRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetApplianceContainerRes() *GetApplianceContainerRes {
	this := GetApplianceContainerRes{}
	return &this
}

// NewGetApplianceContainerResWithDefaults instantiates a new GetApplianceContainerRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetApplianceContainerResWithDefaults() *GetApplianceContainerRes {
	this := GetApplianceContainerRes{}
	return &this
}

// GetAppliance returns the Appliance field value if set, zero value otherwise.
func (o *GetApplianceContainerRes) GetAppliance() IdentityApplianceContainerDTO {
	if o == nil || isNil(o.Appliance) {
		var ret IdentityApplianceContainerDTO
		return ret
	}
	return *o.Appliance
}

// GetApplianceOk returns a tuple with the Appliance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetApplianceContainerRes) GetApplianceOk() (*IdentityApplianceContainerDTO, bool) {
	if o == nil || isNil(o.Appliance) {
    return nil, false
	}
	return o.Appliance, true
}

// HasAppliance returns a boolean if a field has been set.
func (o *GetApplianceContainerRes) HasAppliance() bool {
	if o != nil && !isNil(o.Appliance) {
		return true
	}

	return false
}

// SetAppliance gets a reference to the given IdentityApplianceContainerDTO and assigns it to the Appliance field.
func (o *GetApplianceContainerRes) SetAppliance(v IdentityApplianceContainerDTO) {
	o.Appliance = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *GetApplianceContainerRes) GetError() string {
	if o == nil || isNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetApplianceContainerRes) GetErrorOk() (*string, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *GetApplianceContainerRes) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *GetApplianceContainerRes) SetError(v string) {
	o.Error = &v
}

// GetValidationErrors returns the ValidationErrors field value if set, zero value otherwise.
func (o *GetApplianceContainerRes) GetValidationErrors() []string {
	if o == nil || isNil(o.ValidationErrors) {
		var ret []string
		return ret
	}
	return o.ValidationErrors
}

// GetValidationErrorsOk returns a tuple with the ValidationErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetApplianceContainerRes) GetValidationErrorsOk() ([]string, bool) {
	if o == nil || isNil(o.ValidationErrors) {
    return nil, false
	}
	return o.ValidationErrors, true
}

// HasValidationErrors returns a boolean if a field has been set.
func (o *GetApplianceContainerRes) HasValidationErrors() bool {
	if o != nil && !isNil(o.ValidationErrors) {
		return true
	}

	return false
}

// SetValidationErrors gets a reference to the given []string and assigns it to the ValidationErrors field.
func (o *GetApplianceContainerRes) SetValidationErrors(v []string) {
	o.ValidationErrors = v
}

func (o GetApplianceContainerRes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Appliance) {
		toSerialize["appliance"] = o.Appliance
	}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !isNil(o.ValidationErrors) {
		toSerialize["validationErrors"] = o.ValidationErrors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GetApplianceContainerRes) UnmarshalJSON(bytes []byte) (err error) {
	varGetApplianceContainerRes := _GetApplianceContainerRes{}

	if err = json.Unmarshal(bytes, &varGetApplianceContainerRes); err == nil {
		*o = GetApplianceContainerRes(varGetApplianceContainerRes)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "appliance")
		delete(additionalProperties, "error")
		delete(additionalProperties, "validationErrors")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetApplianceContainerRes struct {
	value *GetApplianceContainerRes
	isSet bool
}

func (v NullableGetApplianceContainerRes) Get() *GetApplianceContainerRes {
	return v.value
}

func (v *NullableGetApplianceContainerRes) Set(val *GetApplianceContainerRes) {
	v.value = val
	v.isSet = true
}

func (v NullableGetApplianceContainerRes) IsSet() bool {
	return v.isSet
}

func (v *NullableGetApplianceContainerRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetApplianceContainerRes(val *GetApplianceContainerRes) *NullableGetApplianceContainerRes {
	return &NullableGetApplianceContainerRes{value: val, isSet: true}
}

func (v NullableGetApplianceContainerRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetApplianceContainerRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


