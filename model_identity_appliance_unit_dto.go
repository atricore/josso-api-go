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

// IdentityApplianceUnitDTO struct for IdentityApplianceUnitDTO
type IdentityApplianceUnitDTO struct {
	BundleName *string `json:"bundleName,omitempty"`
	Description *string `json:"description,omitempty"`
	ElementId *string `json:"elementId,omitempty"`
	Group *string `json:"group,omitempty"`
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Providers []ProviderDTO `json:"providers,omitempty"`
	Type *string `json:"type,omitempty"`
	Version *string `json:"version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentityApplianceUnitDTO IdentityApplianceUnitDTO

// NewIdentityApplianceUnitDTO instantiates a new IdentityApplianceUnitDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityApplianceUnitDTO() *IdentityApplianceUnitDTO {
	this := IdentityApplianceUnitDTO{}
	return &this
}

// NewIdentityApplianceUnitDTOWithDefaults instantiates a new IdentityApplianceUnitDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityApplianceUnitDTOWithDefaults() *IdentityApplianceUnitDTO {
	this := IdentityApplianceUnitDTO{}
	return &this
}

// GetBundleName returns the BundleName field value if set, zero value otherwise.
func (o *IdentityApplianceUnitDTO) GetBundleName() string {
	if o == nil || isNil(o.BundleName) {
		var ret string
		return ret
	}
	return *o.BundleName
}

// GetBundleNameOk returns a tuple with the BundleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityApplianceUnitDTO) GetBundleNameOk() (*string, bool) {
	if o == nil || isNil(o.BundleName) {
    return nil, false
	}
	return o.BundleName, true
}

// HasBundleName returns a boolean if a field has been set.
func (o *IdentityApplianceUnitDTO) HasBundleName() bool {
	if o != nil && !isNil(o.BundleName) {
		return true
	}

	return false
}

// SetBundleName gets a reference to the given string and assigns it to the BundleName field.
func (o *IdentityApplianceUnitDTO) SetBundleName(v string) {
	o.BundleName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IdentityApplianceUnitDTO) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityApplianceUnitDTO) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IdentityApplianceUnitDTO) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IdentityApplianceUnitDTO) SetDescription(v string) {
	o.Description = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *IdentityApplianceUnitDTO) GetElementId() string {
	if o == nil || isNil(o.ElementId) {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityApplianceUnitDTO) GetElementIdOk() (*string, bool) {
	if o == nil || isNil(o.ElementId) {
    return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *IdentityApplianceUnitDTO) HasElementId() bool {
	if o != nil && !isNil(o.ElementId) {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *IdentityApplianceUnitDTO) SetElementId(v string) {
	o.ElementId = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *IdentityApplianceUnitDTO) GetGroup() string {
	if o == nil || isNil(o.Group) {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityApplianceUnitDTO) GetGroupOk() (*string, bool) {
	if o == nil || isNil(o.Group) {
    return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *IdentityApplianceUnitDTO) HasGroup() bool {
	if o != nil && !isNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *IdentityApplianceUnitDTO) SetGroup(v string) {
	o.Group = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IdentityApplianceUnitDTO) GetId() int64 {
	if o == nil || isNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityApplianceUnitDTO) GetIdOk() (*int64, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IdentityApplianceUnitDTO) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *IdentityApplianceUnitDTO) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IdentityApplianceUnitDTO) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityApplianceUnitDTO) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IdentityApplianceUnitDTO) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IdentityApplianceUnitDTO) SetName(v string) {
	o.Name = &v
}

// GetProviders returns the Providers field value if set, zero value otherwise.
func (o *IdentityApplianceUnitDTO) GetProviders() []ProviderDTO {
	if o == nil || isNil(o.Providers) {
		var ret []ProviderDTO
		return ret
	}
	return o.Providers
}

// GetProvidersOk returns a tuple with the Providers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityApplianceUnitDTO) GetProvidersOk() ([]ProviderDTO, bool) {
	if o == nil || isNil(o.Providers) {
    return nil, false
	}
	return o.Providers, true
}

// HasProviders returns a boolean if a field has been set.
func (o *IdentityApplianceUnitDTO) HasProviders() bool {
	if o != nil && !isNil(o.Providers) {
		return true
	}

	return false
}

// SetProviders gets a reference to the given []ProviderDTO and assigns it to the Providers field.
func (o *IdentityApplianceUnitDTO) SetProviders(v []ProviderDTO) {
	o.Providers = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IdentityApplianceUnitDTO) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityApplianceUnitDTO) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IdentityApplianceUnitDTO) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *IdentityApplianceUnitDTO) SetType(v string) {
	o.Type = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *IdentityApplianceUnitDTO) GetVersion() string {
	if o == nil || isNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityApplianceUnitDTO) GetVersionOk() (*string, bool) {
	if o == nil || isNil(o.Version) {
    return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *IdentityApplianceUnitDTO) HasVersion() bool {
	if o != nil && !isNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *IdentityApplianceUnitDTO) SetVersion(v string) {
	o.Version = &v
}

func (o IdentityApplianceUnitDTO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.BundleName) {
		toSerialize["bundleName"] = o.BundleName
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.ElementId) {
		toSerialize["elementId"] = o.ElementId
	}
	if !isNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Providers) {
		toSerialize["providers"] = o.Providers
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Version) {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IdentityApplianceUnitDTO) UnmarshalJSON(bytes []byte) (err error) {
	varIdentityApplianceUnitDTO := _IdentityApplianceUnitDTO{}

	if err = json.Unmarshal(bytes, &varIdentityApplianceUnitDTO); err == nil {
		*o = IdentityApplianceUnitDTO(varIdentityApplianceUnitDTO)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "bundleName")
		delete(additionalProperties, "description")
		delete(additionalProperties, "elementId")
		delete(additionalProperties, "group")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "providers")
		delete(additionalProperties, "type")
		delete(additionalProperties, "version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentityApplianceUnitDTO struct {
	value *IdentityApplianceUnitDTO
	isSet bool
}

func (v NullableIdentityApplianceUnitDTO) Get() *IdentityApplianceUnitDTO {
	return v.value
}

func (v *NullableIdentityApplianceUnitDTO) Set(val *IdentityApplianceUnitDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityApplianceUnitDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityApplianceUnitDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityApplianceUnitDTO(val *IdentityApplianceUnitDTO) *NullableIdentityApplianceUnitDTO {
	return &NullableIdentityApplianceUnitDTO{value: val, isSet: true}
}

func (v NullableIdentityApplianceUnitDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityApplianceUnitDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


