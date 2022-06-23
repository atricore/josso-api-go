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

// AccountLinkagePolicyDTO struct for AccountLinkagePolicyDTO
type AccountLinkagePolicyDTO struct {
	CustomLinkEmitter *string `json:"customLinkEmitter,omitempty"`
	ElementId *string `json:"elementId,omitempty"`
	Id *int64 `json:"id,omitempty"`
	LinkEmitterType *string `json:"linkEmitterType,omitempty"`
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountLinkagePolicyDTO AccountLinkagePolicyDTO

// NewAccountLinkagePolicyDTO instantiates a new AccountLinkagePolicyDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountLinkagePolicyDTO() *AccountLinkagePolicyDTO {
	this := AccountLinkagePolicyDTO{}
	return &this
}

// NewAccountLinkagePolicyDTOWithDefaults instantiates a new AccountLinkagePolicyDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountLinkagePolicyDTOWithDefaults() *AccountLinkagePolicyDTO {
	this := AccountLinkagePolicyDTO{}
	return &this
}

// GetCustomLinkEmitter returns the CustomLinkEmitter field value if set, zero value otherwise.
func (o *AccountLinkagePolicyDTO) GetCustomLinkEmitter() string {
	if o == nil || o.CustomLinkEmitter == nil {
		var ret string
		return ret
	}
	return *o.CustomLinkEmitter
}

// GetCustomLinkEmitterOk returns a tuple with the CustomLinkEmitter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountLinkagePolicyDTO) GetCustomLinkEmitterOk() (*string, bool) {
	if o == nil || o.CustomLinkEmitter == nil {
		return nil, false
	}
	return o.CustomLinkEmitter, true
}

// HasCustomLinkEmitter returns a boolean if a field has been set.
func (o *AccountLinkagePolicyDTO) HasCustomLinkEmitter() bool {
	if o != nil && o.CustomLinkEmitter != nil {
		return true
	}

	return false
}

// SetCustomLinkEmitter gets a reference to the given string and assigns it to the CustomLinkEmitter field.
func (o *AccountLinkagePolicyDTO) SetCustomLinkEmitter(v string) {
	o.CustomLinkEmitter = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *AccountLinkagePolicyDTO) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountLinkagePolicyDTO) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *AccountLinkagePolicyDTO) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *AccountLinkagePolicyDTO) SetElementId(v string) {
	o.ElementId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AccountLinkagePolicyDTO) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountLinkagePolicyDTO) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AccountLinkagePolicyDTO) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *AccountLinkagePolicyDTO) SetId(v int64) {
	o.Id = &v
}

// GetLinkEmitterType returns the LinkEmitterType field value if set, zero value otherwise.
func (o *AccountLinkagePolicyDTO) GetLinkEmitterType() string {
	if o == nil || o.LinkEmitterType == nil {
		var ret string
		return ret
	}
	return *o.LinkEmitterType
}

// GetLinkEmitterTypeOk returns a tuple with the LinkEmitterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountLinkagePolicyDTO) GetLinkEmitterTypeOk() (*string, bool) {
	if o == nil || o.LinkEmitterType == nil {
		return nil, false
	}
	return o.LinkEmitterType, true
}

// HasLinkEmitterType returns a boolean if a field has been set.
func (o *AccountLinkagePolicyDTO) HasLinkEmitterType() bool {
	if o != nil && o.LinkEmitterType != nil {
		return true
	}

	return false
}

// SetLinkEmitterType gets a reference to the given string and assigns it to the LinkEmitterType field.
func (o *AccountLinkagePolicyDTO) SetLinkEmitterType(v string) {
	o.LinkEmitterType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AccountLinkagePolicyDTO) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountLinkagePolicyDTO) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AccountLinkagePolicyDTO) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AccountLinkagePolicyDTO) SetName(v string) {
	o.Name = &v
}

func (o AccountLinkagePolicyDTO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CustomLinkEmitter != nil {
		toSerialize["customLinkEmitter"] = o.CustomLinkEmitter
	}
	if o.ElementId != nil {
		toSerialize["elementId"] = o.ElementId
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LinkEmitterType != nil {
		toSerialize["linkEmitterType"] = o.LinkEmitterType
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AccountLinkagePolicyDTO) UnmarshalJSON(bytes []byte) (err error) {
	varAccountLinkagePolicyDTO := _AccountLinkagePolicyDTO{}

	if err = json.Unmarshal(bytes, &varAccountLinkagePolicyDTO); err == nil {
		*o = AccountLinkagePolicyDTO(varAccountLinkagePolicyDTO)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "customLinkEmitter")
		delete(additionalProperties, "elementId")
		delete(additionalProperties, "id")
		delete(additionalProperties, "linkEmitterType")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountLinkagePolicyDTO struct {
	value *AccountLinkagePolicyDTO
	isSet bool
}

func (v NullableAccountLinkagePolicyDTO) Get() *AccountLinkagePolicyDTO {
	return v.value
}

func (v *NullableAccountLinkagePolicyDTO) Set(val *AccountLinkagePolicyDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountLinkagePolicyDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountLinkagePolicyDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountLinkagePolicyDTO(val *AccountLinkagePolicyDTO) *NullableAccountLinkagePolicyDTO {
	return &NullableAccountLinkagePolicyDTO{value: val, isSet: true}
}

func (v NullableAccountLinkagePolicyDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountLinkagePolicyDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


