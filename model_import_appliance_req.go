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

// ImportApplianceReq struct for ImportApplianceReq
type ImportApplianceReq struct {
	Base64Json *string `json:"base64Json,omitempty"`
	IdaName *string `json:"idaName,omitempty"`
	Modify *bool `json:"modify,omitempty"`
	NewDescription *string `json:"newDescription,omitempty"`
	NewLocation *string `json:"newLocation,omitempty"`
	NewName *string `json:"newName,omitempty"`
	NewNamespace *string `json:"newNamespace,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ImportApplianceReq ImportApplianceReq

// NewImportApplianceReq instantiates a new ImportApplianceReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportApplianceReq() *ImportApplianceReq {
	this := ImportApplianceReq{}
	return &this
}

// NewImportApplianceReqWithDefaults instantiates a new ImportApplianceReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportApplianceReqWithDefaults() *ImportApplianceReq {
	this := ImportApplianceReq{}
	return &this
}

// GetBase64Json returns the Base64Json field value if set, zero value otherwise.
func (o *ImportApplianceReq) GetBase64Json() string {
	if o == nil || o.Base64Json == nil {
		var ret string
		return ret
	}
	return *o.Base64Json
}

// GetBase64JsonOk returns a tuple with the Base64Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportApplianceReq) GetBase64JsonOk() (*string, bool) {
	if o == nil || o.Base64Json == nil {
		return nil, false
	}
	return o.Base64Json, true
}

// HasBase64Json returns a boolean if a field has been set.
func (o *ImportApplianceReq) HasBase64Json() bool {
	if o != nil && o.Base64Json != nil {
		return true
	}

	return false
}

// SetBase64Json gets a reference to the given string and assigns it to the Base64Json field.
func (o *ImportApplianceReq) SetBase64Json(v string) {
	o.Base64Json = &v
}

// GetIdaName returns the IdaName field value if set, zero value otherwise.
func (o *ImportApplianceReq) GetIdaName() string {
	if o == nil || o.IdaName == nil {
		var ret string
		return ret
	}
	return *o.IdaName
}

// GetIdaNameOk returns a tuple with the IdaName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportApplianceReq) GetIdaNameOk() (*string, bool) {
	if o == nil || o.IdaName == nil {
		return nil, false
	}
	return o.IdaName, true
}

// HasIdaName returns a boolean if a field has been set.
func (o *ImportApplianceReq) HasIdaName() bool {
	if o != nil && o.IdaName != nil {
		return true
	}

	return false
}

// SetIdaName gets a reference to the given string and assigns it to the IdaName field.
func (o *ImportApplianceReq) SetIdaName(v string) {
	o.IdaName = &v
}

// GetModify returns the Modify field value if set, zero value otherwise.
func (o *ImportApplianceReq) GetModify() bool {
	if o == nil || o.Modify == nil {
		var ret bool
		return ret
	}
	return *o.Modify
}

// GetModifyOk returns a tuple with the Modify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportApplianceReq) GetModifyOk() (*bool, bool) {
	if o == nil || o.Modify == nil {
		return nil, false
	}
	return o.Modify, true
}

// HasModify returns a boolean if a field has been set.
func (o *ImportApplianceReq) HasModify() bool {
	if o != nil && o.Modify != nil {
		return true
	}

	return false
}

// SetModify gets a reference to the given bool and assigns it to the Modify field.
func (o *ImportApplianceReq) SetModify(v bool) {
	o.Modify = &v
}

// GetNewDescription returns the NewDescription field value if set, zero value otherwise.
func (o *ImportApplianceReq) GetNewDescription() string {
	if o == nil || o.NewDescription == nil {
		var ret string
		return ret
	}
	return *o.NewDescription
}

// GetNewDescriptionOk returns a tuple with the NewDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportApplianceReq) GetNewDescriptionOk() (*string, bool) {
	if o == nil || o.NewDescription == nil {
		return nil, false
	}
	return o.NewDescription, true
}

// HasNewDescription returns a boolean if a field has been set.
func (o *ImportApplianceReq) HasNewDescription() bool {
	if o != nil && o.NewDescription != nil {
		return true
	}

	return false
}

// SetNewDescription gets a reference to the given string and assigns it to the NewDescription field.
func (o *ImportApplianceReq) SetNewDescription(v string) {
	o.NewDescription = &v
}

// GetNewLocation returns the NewLocation field value if set, zero value otherwise.
func (o *ImportApplianceReq) GetNewLocation() string {
	if o == nil || o.NewLocation == nil {
		var ret string
		return ret
	}
	return *o.NewLocation
}

// GetNewLocationOk returns a tuple with the NewLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportApplianceReq) GetNewLocationOk() (*string, bool) {
	if o == nil || o.NewLocation == nil {
		return nil, false
	}
	return o.NewLocation, true
}

// HasNewLocation returns a boolean if a field has been set.
func (o *ImportApplianceReq) HasNewLocation() bool {
	if o != nil && o.NewLocation != nil {
		return true
	}

	return false
}

// SetNewLocation gets a reference to the given string and assigns it to the NewLocation field.
func (o *ImportApplianceReq) SetNewLocation(v string) {
	o.NewLocation = &v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *ImportApplianceReq) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportApplianceReq) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *ImportApplianceReq) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *ImportApplianceReq) SetNewName(v string) {
	o.NewName = &v
}

// GetNewNamespace returns the NewNamespace field value if set, zero value otherwise.
func (o *ImportApplianceReq) GetNewNamespace() string {
	if o == nil || o.NewNamespace == nil {
		var ret string
		return ret
	}
	return *o.NewNamespace
}

// GetNewNamespaceOk returns a tuple with the NewNamespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportApplianceReq) GetNewNamespaceOk() (*string, bool) {
	if o == nil || o.NewNamespace == nil {
		return nil, false
	}
	return o.NewNamespace, true
}

// HasNewNamespace returns a boolean if a field has been set.
func (o *ImportApplianceReq) HasNewNamespace() bool {
	if o != nil && o.NewNamespace != nil {
		return true
	}

	return false
}

// SetNewNamespace gets a reference to the given string and assigns it to the NewNamespace field.
func (o *ImportApplianceReq) SetNewNamespace(v string) {
	o.NewNamespace = &v
}

func (o ImportApplianceReq) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Base64Json != nil {
		toSerialize["base64Json"] = o.Base64Json
	}
	if o.IdaName != nil {
		toSerialize["idaName"] = o.IdaName
	}
	if o.Modify != nil {
		toSerialize["modify"] = o.Modify
	}
	if o.NewDescription != nil {
		toSerialize["newDescription"] = o.NewDescription
	}
	if o.NewLocation != nil {
		toSerialize["newLocation"] = o.NewLocation
	}
	if o.NewName != nil {
		toSerialize["newName"] = o.NewName
	}
	if o.NewNamespace != nil {
		toSerialize["newNamespace"] = o.NewNamespace
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ImportApplianceReq) UnmarshalJSON(bytes []byte) (err error) {
	varImportApplianceReq := _ImportApplianceReq{}

	if err = json.Unmarshal(bytes, &varImportApplianceReq); err == nil {
		*o = ImportApplianceReq(varImportApplianceReq)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "base64Json")
		delete(additionalProperties, "idaName")
		delete(additionalProperties, "modify")
		delete(additionalProperties, "newDescription")
		delete(additionalProperties, "newLocation")
		delete(additionalProperties, "newName")
		delete(additionalProperties, "newNamespace")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableImportApplianceReq struct {
	value *ImportApplianceReq
	isSet bool
}

func (v NullableImportApplianceReq) Get() *ImportApplianceReq {
	return v.value
}

func (v *NullableImportApplianceReq) Set(val *ImportApplianceReq) {
	v.value = val
	v.isSet = true
}

func (v NullableImportApplianceReq) IsSet() bool {
	return v.isSet
}

func (v *NullableImportApplianceReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportApplianceReq(val *ImportApplianceReq) *NullableImportApplianceReq {
	return &NullableImportApplianceReq{value: val, isSet: true}
}

func (v NullableImportApplianceReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportApplianceReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


