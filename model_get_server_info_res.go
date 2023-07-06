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

// GetServerInfoRes struct for GetServerInfoRes
type GetServerInfoRes struct {
	Error *string `json:"error,omitempty"`
	NodeId *string `json:"nodeId,omitempty"`
	ValidationErrors []string `json:"validationErrors,omitempty"`
	Version *string `json:"version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetServerInfoRes GetServerInfoRes

// NewGetServerInfoRes instantiates a new GetServerInfoRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetServerInfoRes() *GetServerInfoRes {
	this := GetServerInfoRes{}
	return &this
}

// NewGetServerInfoResWithDefaults instantiates a new GetServerInfoRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetServerInfoResWithDefaults() *GetServerInfoRes {
	this := GetServerInfoRes{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *GetServerInfoRes) GetError() string {
	if o == nil || isNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetServerInfoRes) GetErrorOk() (*string, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *GetServerInfoRes) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *GetServerInfoRes) SetError(v string) {
	o.Error = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *GetServerInfoRes) GetNodeId() string {
	if o == nil || isNil(o.NodeId) {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetServerInfoRes) GetNodeIdOk() (*string, bool) {
	if o == nil || isNil(o.NodeId) {
    return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *GetServerInfoRes) HasNodeId() bool {
	if o != nil && !isNil(o.NodeId) {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *GetServerInfoRes) SetNodeId(v string) {
	o.NodeId = &v
}

// GetValidationErrors returns the ValidationErrors field value if set, zero value otherwise.
func (o *GetServerInfoRes) GetValidationErrors() []string {
	if o == nil || isNil(o.ValidationErrors) {
		var ret []string
		return ret
	}
	return o.ValidationErrors
}

// GetValidationErrorsOk returns a tuple with the ValidationErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetServerInfoRes) GetValidationErrorsOk() ([]string, bool) {
	if o == nil || isNil(o.ValidationErrors) {
    return nil, false
	}
	return o.ValidationErrors, true
}

// HasValidationErrors returns a boolean if a field has been set.
func (o *GetServerInfoRes) HasValidationErrors() bool {
	if o != nil && !isNil(o.ValidationErrors) {
		return true
	}

	return false
}

// SetValidationErrors gets a reference to the given []string and assigns it to the ValidationErrors field.
func (o *GetServerInfoRes) SetValidationErrors(v []string) {
	o.ValidationErrors = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *GetServerInfoRes) GetVersion() string {
	if o == nil || isNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetServerInfoRes) GetVersionOk() (*string, bool) {
	if o == nil || isNil(o.Version) {
    return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *GetServerInfoRes) HasVersion() bool {
	if o != nil && !isNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *GetServerInfoRes) SetVersion(v string) {
	o.Version = &v
}

func (o GetServerInfoRes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !isNil(o.NodeId) {
		toSerialize["nodeId"] = o.NodeId
	}
	if !isNil(o.ValidationErrors) {
		toSerialize["validationErrors"] = o.ValidationErrors
	}
	if !isNil(o.Version) {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GetServerInfoRes) UnmarshalJSON(bytes []byte) (err error) {
	varGetServerInfoRes := _GetServerInfoRes{}

	if err = json.Unmarshal(bytes, &varGetServerInfoRes); err == nil {
		*o = GetServerInfoRes(varGetServerInfoRes)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "error")
		delete(additionalProperties, "nodeId")
		delete(additionalProperties, "validationErrors")
		delete(additionalProperties, "version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetServerInfoRes struct {
	value *GetServerInfoRes
	isSet bool
}

func (v NullableGetServerInfoRes) Get() *GetServerInfoRes {
	return v.value
}

func (v *NullableGetServerInfoRes) Set(val *GetServerInfoRes) {
	v.value = val
	v.isSet = true
}

func (v NullableGetServerInfoRes) IsSet() bool {
	return v.isSet
}

func (v *NullableGetServerInfoRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetServerInfoRes(val *GetServerInfoRes) *NullableGetServerInfoRes {
	return &NullableGetServerInfoRes{value: val, isSet: true}
}

func (v NullableGetServerInfoRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetServerInfoRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


