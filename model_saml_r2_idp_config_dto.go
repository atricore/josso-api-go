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

// SamlR2IDPConfigDTO struct for SamlR2IDPConfigDTO
type SamlR2IDPConfigDTO struct {
	Description *string `json:"description,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	ElementId *string `json:"elementId,omitempty"`
	Encrypter *KeystoreDTO `json:"encrypter,omitempty"`
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Signer *KeystoreDTO `json:"signer,omitempty"`
	UseSampleStore *bool `json:"useSampleStore,omitempty"`
	UseSystemStore *bool `json:"useSystemStore,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SamlR2IDPConfigDTO SamlR2IDPConfigDTO

// NewSamlR2IDPConfigDTO instantiates a new SamlR2IDPConfigDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSamlR2IDPConfigDTO() *SamlR2IDPConfigDTO {
	this := SamlR2IDPConfigDTO{}
	return &this
}

// NewSamlR2IDPConfigDTOWithDefaults instantiates a new SamlR2IDPConfigDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSamlR2IDPConfigDTOWithDefaults() *SamlR2IDPConfigDTO {
	this := SamlR2IDPConfigDTO{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SamlR2IDPConfigDTO) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlR2IDPConfigDTO) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SamlR2IDPConfigDTO) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SamlR2IDPConfigDTO) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *SamlR2IDPConfigDTO) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlR2IDPConfigDTO) GetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.DisplayName) {
    return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *SamlR2IDPConfigDTO) HasDisplayName() bool {
	if o != nil && !isNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *SamlR2IDPConfigDTO) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *SamlR2IDPConfigDTO) GetElementId() string {
	if o == nil || isNil(o.ElementId) {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlR2IDPConfigDTO) GetElementIdOk() (*string, bool) {
	if o == nil || isNil(o.ElementId) {
    return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *SamlR2IDPConfigDTO) HasElementId() bool {
	if o != nil && !isNil(o.ElementId) {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *SamlR2IDPConfigDTO) SetElementId(v string) {
	o.ElementId = &v
}

// GetEncrypter returns the Encrypter field value if set, zero value otherwise.
func (o *SamlR2IDPConfigDTO) GetEncrypter() KeystoreDTO {
	if o == nil || isNil(o.Encrypter) {
		var ret KeystoreDTO
		return ret
	}
	return *o.Encrypter
}

// GetEncrypterOk returns a tuple with the Encrypter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlR2IDPConfigDTO) GetEncrypterOk() (*KeystoreDTO, bool) {
	if o == nil || isNil(o.Encrypter) {
    return nil, false
	}
	return o.Encrypter, true
}

// HasEncrypter returns a boolean if a field has been set.
func (o *SamlR2IDPConfigDTO) HasEncrypter() bool {
	if o != nil && !isNil(o.Encrypter) {
		return true
	}

	return false
}

// SetEncrypter gets a reference to the given KeystoreDTO and assigns it to the Encrypter field.
func (o *SamlR2IDPConfigDTO) SetEncrypter(v KeystoreDTO) {
	o.Encrypter = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SamlR2IDPConfigDTO) GetId() int64 {
	if o == nil || isNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlR2IDPConfigDTO) GetIdOk() (*int64, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SamlR2IDPConfigDTO) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *SamlR2IDPConfigDTO) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SamlR2IDPConfigDTO) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlR2IDPConfigDTO) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SamlR2IDPConfigDTO) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SamlR2IDPConfigDTO) SetName(v string) {
	o.Name = &v
}

// GetSigner returns the Signer field value if set, zero value otherwise.
func (o *SamlR2IDPConfigDTO) GetSigner() KeystoreDTO {
	if o == nil || isNil(o.Signer) {
		var ret KeystoreDTO
		return ret
	}
	return *o.Signer
}

// GetSignerOk returns a tuple with the Signer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlR2IDPConfigDTO) GetSignerOk() (*KeystoreDTO, bool) {
	if o == nil || isNil(o.Signer) {
    return nil, false
	}
	return o.Signer, true
}

// HasSigner returns a boolean if a field has been set.
func (o *SamlR2IDPConfigDTO) HasSigner() bool {
	if o != nil && !isNil(o.Signer) {
		return true
	}

	return false
}

// SetSigner gets a reference to the given KeystoreDTO and assigns it to the Signer field.
func (o *SamlR2IDPConfigDTO) SetSigner(v KeystoreDTO) {
	o.Signer = &v
}

// GetUseSampleStore returns the UseSampleStore field value if set, zero value otherwise.
func (o *SamlR2IDPConfigDTO) GetUseSampleStore() bool {
	if o == nil || isNil(o.UseSampleStore) {
		var ret bool
		return ret
	}
	return *o.UseSampleStore
}

// GetUseSampleStoreOk returns a tuple with the UseSampleStore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlR2IDPConfigDTO) GetUseSampleStoreOk() (*bool, bool) {
	if o == nil || isNil(o.UseSampleStore) {
    return nil, false
	}
	return o.UseSampleStore, true
}

// HasUseSampleStore returns a boolean if a field has been set.
func (o *SamlR2IDPConfigDTO) HasUseSampleStore() bool {
	if o != nil && !isNil(o.UseSampleStore) {
		return true
	}

	return false
}

// SetUseSampleStore gets a reference to the given bool and assigns it to the UseSampleStore field.
func (o *SamlR2IDPConfigDTO) SetUseSampleStore(v bool) {
	o.UseSampleStore = &v
}

// GetUseSystemStore returns the UseSystemStore field value if set, zero value otherwise.
func (o *SamlR2IDPConfigDTO) GetUseSystemStore() bool {
	if o == nil || isNil(o.UseSystemStore) {
		var ret bool
		return ret
	}
	return *o.UseSystemStore
}

// GetUseSystemStoreOk returns a tuple with the UseSystemStore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlR2IDPConfigDTO) GetUseSystemStoreOk() (*bool, bool) {
	if o == nil || isNil(o.UseSystemStore) {
    return nil, false
	}
	return o.UseSystemStore, true
}

// HasUseSystemStore returns a boolean if a field has been set.
func (o *SamlR2IDPConfigDTO) HasUseSystemStore() bool {
	if o != nil && !isNil(o.UseSystemStore) {
		return true
	}

	return false
}

// SetUseSystemStore gets a reference to the given bool and assigns it to the UseSystemStore field.
func (o *SamlR2IDPConfigDTO) SetUseSystemStore(v bool) {
	o.UseSystemStore = &v
}

func (o SamlR2IDPConfigDTO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !isNil(o.ElementId) {
		toSerialize["elementId"] = o.ElementId
	}
	if !isNil(o.Encrypter) {
		toSerialize["encrypter"] = o.Encrypter
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Signer) {
		toSerialize["signer"] = o.Signer
	}
	if !isNil(o.UseSampleStore) {
		toSerialize["useSampleStore"] = o.UseSampleStore
	}
	if !isNil(o.UseSystemStore) {
		toSerialize["useSystemStore"] = o.UseSystemStore
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SamlR2IDPConfigDTO) UnmarshalJSON(bytes []byte) (err error) {
	varSamlR2IDPConfigDTO := _SamlR2IDPConfigDTO{}

	if err = json.Unmarshal(bytes, &varSamlR2IDPConfigDTO); err == nil {
		*o = SamlR2IDPConfigDTO(varSamlR2IDPConfigDTO)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "elementId")
		delete(additionalProperties, "encrypter")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "signer")
		delete(additionalProperties, "useSampleStore")
		delete(additionalProperties, "useSystemStore")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSamlR2IDPConfigDTO struct {
	value *SamlR2IDPConfigDTO
	isSet bool
}

func (v NullableSamlR2IDPConfigDTO) Get() *SamlR2IDPConfigDTO {
	return v.value
}

func (v *NullableSamlR2IDPConfigDTO) Set(val *SamlR2IDPConfigDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableSamlR2IDPConfigDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableSamlR2IDPConfigDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSamlR2IDPConfigDTO(val *SamlR2IDPConfigDTO) *NullableSamlR2IDPConfigDTO {
	return &NullableSamlR2IDPConfigDTO{value: val, isSet: true}
}

func (v NullableSamlR2IDPConfigDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSamlR2IDPConfigDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


