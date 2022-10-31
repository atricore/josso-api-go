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

// StoreBrandingReq struct for StoreBrandingReq
type StoreBrandingReq struct {
	BundleArtifact *string `json:"bundleArtifact,omitempty"`
	BundleGroup *string `json:"bundleGroup,omitempty"`
	BundleVersion *string `json:"bundleVersion,omitempty"`
	CustomOpenIdAppClazz *string `json:"customOpenIdAppClazz,omitempty"`
	CustomSsoAppClazz *string `json:"customSsoAppClazz,omitempty"`
	CustomSsoIdPAppClazz *string `json:"customSsoIdPAppClazz,omitempty"`
	Description *string `json:"description,omitempty"`
	Name *string `json:"name,omitempty"`
	Resource *string `json:"resource,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StoreBrandingReq StoreBrandingReq

// NewStoreBrandingReq instantiates a new StoreBrandingReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoreBrandingReq() *StoreBrandingReq {
	this := StoreBrandingReq{}
	return &this
}

// NewStoreBrandingReqWithDefaults instantiates a new StoreBrandingReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoreBrandingReqWithDefaults() *StoreBrandingReq {
	this := StoreBrandingReq{}
	return &this
}

// GetBundleArtifact returns the BundleArtifact field value if set, zero value otherwise.
func (o *StoreBrandingReq) GetBundleArtifact() string {
	if o == nil || o.BundleArtifact == nil {
		var ret string
		return ret
	}
	return *o.BundleArtifact
}

// GetBundleArtifactOk returns a tuple with the BundleArtifact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreBrandingReq) GetBundleArtifactOk() (*string, bool) {
	if o == nil || o.BundleArtifact == nil {
		return nil, false
	}
	return o.BundleArtifact, true
}

// HasBundleArtifact returns a boolean if a field has been set.
func (o *StoreBrandingReq) HasBundleArtifact() bool {
	if o != nil && o.BundleArtifact != nil {
		return true
	}

	return false
}

// SetBundleArtifact gets a reference to the given string and assigns it to the BundleArtifact field.
func (o *StoreBrandingReq) SetBundleArtifact(v string) {
	o.BundleArtifact = &v
}

// GetBundleGroup returns the BundleGroup field value if set, zero value otherwise.
func (o *StoreBrandingReq) GetBundleGroup() string {
	if o == nil || o.BundleGroup == nil {
		var ret string
		return ret
	}
	return *o.BundleGroup
}

// GetBundleGroupOk returns a tuple with the BundleGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreBrandingReq) GetBundleGroupOk() (*string, bool) {
	if o == nil || o.BundleGroup == nil {
		return nil, false
	}
	return o.BundleGroup, true
}

// HasBundleGroup returns a boolean if a field has been set.
func (o *StoreBrandingReq) HasBundleGroup() bool {
	if o != nil && o.BundleGroup != nil {
		return true
	}

	return false
}

// SetBundleGroup gets a reference to the given string and assigns it to the BundleGroup field.
func (o *StoreBrandingReq) SetBundleGroup(v string) {
	o.BundleGroup = &v
}

// GetBundleVersion returns the BundleVersion field value if set, zero value otherwise.
func (o *StoreBrandingReq) GetBundleVersion() string {
	if o == nil || o.BundleVersion == nil {
		var ret string
		return ret
	}
	return *o.BundleVersion
}

// GetBundleVersionOk returns a tuple with the BundleVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreBrandingReq) GetBundleVersionOk() (*string, bool) {
	if o == nil || o.BundleVersion == nil {
		return nil, false
	}
	return o.BundleVersion, true
}

// HasBundleVersion returns a boolean if a field has been set.
func (o *StoreBrandingReq) HasBundleVersion() bool {
	if o != nil && o.BundleVersion != nil {
		return true
	}

	return false
}

// SetBundleVersion gets a reference to the given string and assigns it to the BundleVersion field.
func (o *StoreBrandingReq) SetBundleVersion(v string) {
	o.BundleVersion = &v
}

// GetCustomOpenIdAppClazz returns the CustomOpenIdAppClazz field value if set, zero value otherwise.
func (o *StoreBrandingReq) GetCustomOpenIdAppClazz() string {
	if o == nil || o.CustomOpenIdAppClazz == nil {
		var ret string
		return ret
	}
	return *o.CustomOpenIdAppClazz
}

// GetCustomOpenIdAppClazzOk returns a tuple with the CustomOpenIdAppClazz field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreBrandingReq) GetCustomOpenIdAppClazzOk() (*string, bool) {
	if o == nil || o.CustomOpenIdAppClazz == nil {
		return nil, false
	}
	return o.CustomOpenIdAppClazz, true
}

// HasCustomOpenIdAppClazz returns a boolean if a field has been set.
func (o *StoreBrandingReq) HasCustomOpenIdAppClazz() bool {
	if o != nil && o.CustomOpenIdAppClazz != nil {
		return true
	}

	return false
}

// SetCustomOpenIdAppClazz gets a reference to the given string and assigns it to the CustomOpenIdAppClazz field.
func (o *StoreBrandingReq) SetCustomOpenIdAppClazz(v string) {
	o.CustomOpenIdAppClazz = &v
}

// GetCustomSsoAppClazz returns the CustomSsoAppClazz field value if set, zero value otherwise.
func (o *StoreBrandingReq) GetCustomSsoAppClazz() string {
	if o == nil || o.CustomSsoAppClazz == nil {
		var ret string
		return ret
	}
	return *o.CustomSsoAppClazz
}

// GetCustomSsoAppClazzOk returns a tuple with the CustomSsoAppClazz field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreBrandingReq) GetCustomSsoAppClazzOk() (*string, bool) {
	if o == nil || o.CustomSsoAppClazz == nil {
		return nil, false
	}
	return o.CustomSsoAppClazz, true
}

// HasCustomSsoAppClazz returns a boolean if a field has been set.
func (o *StoreBrandingReq) HasCustomSsoAppClazz() bool {
	if o != nil && o.CustomSsoAppClazz != nil {
		return true
	}

	return false
}

// SetCustomSsoAppClazz gets a reference to the given string and assigns it to the CustomSsoAppClazz field.
func (o *StoreBrandingReq) SetCustomSsoAppClazz(v string) {
	o.CustomSsoAppClazz = &v
}

// GetCustomSsoIdPAppClazz returns the CustomSsoIdPAppClazz field value if set, zero value otherwise.
func (o *StoreBrandingReq) GetCustomSsoIdPAppClazz() string {
	if o == nil || o.CustomSsoIdPAppClazz == nil {
		var ret string
		return ret
	}
	return *o.CustomSsoIdPAppClazz
}

// GetCustomSsoIdPAppClazzOk returns a tuple with the CustomSsoIdPAppClazz field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreBrandingReq) GetCustomSsoIdPAppClazzOk() (*string, bool) {
	if o == nil || o.CustomSsoIdPAppClazz == nil {
		return nil, false
	}
	return o.CustomSsoIdPAppClazz, true
}

// HasCustomSsoIdPAppClazz returns a boolean if a field has been set.
func (o *StoreBrandingReq) HasCustomSsoIdPAppClazz() bool {
	if o != nil && o.CustomSsoIdPAppClazz != nil {
		return true
	}

	return false
}

// SetCustomSsoIdPAppClazz gets a reference to the given string and assigns it to the CustomSsoIdPAppClazz field.
func (o *StoreBrandingReq) SetCustomSsoIdPAppClazz(v string) {
	o.CustomSsoIdPAppClazz = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *StoreBrandingReq) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreBrandingReq) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *StoreBrandingReq) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *StoreBrandingReq) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StoreBrandingReq) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreBrandingReq) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StoreBrandingReq) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StoreBrandingReq) SetName(v string) {
	o.Name = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *StoreBrandingReq) GetResource() string {
	if o == nil || o.Resource == nil {
		var ret string
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreBrandingReq) GetResourceOk() (*string, bool) {
	if o == nil || o.Resource == nil {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *StoreBrandingReq) HasResource() bool {
	if o != nil && o.Resource != nil {
		return true
	}

	return false
}

// SetResource gets a reference to the given string and assigns it to the Resource field.
func (o *StoreBrandingReq) SetResource(v string) {
	o.Resource = &v
}

func (o StoreBrandingReq) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BundleArtifact != nil {
		toSerialize["bundleArtifact"] = o.BundleArtifact
	}
	if o.BundleGroup != nil {
		toSerialize["bundleGroup"] = o.BundleGroup
	}
	if o.BundleVersion != nil {
		toSerialize["bundleVersion"] = o.BundleVersion
	}
	if o.CustomOpenIdAppClazz != nil {
		toSerialize["customOpenIdAppClazz"] = o.CustomOpenIdAppClazz
	}
	if o.CustomSsoAppClazz != nil {
		toSerialize["customSsoAppClazz"] = o.CustomSsoAppClazz
	}
	if o.CustomSsoIdPAppClazz != nil {
		toSerialize["customSsoIdPAppClazz"] = o.CustomSsoIdPAppClazz
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Resource != nil {
		toSerialize["resource"] = o.Resource
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StoreBrandingReq) UnmarshalJSON(bytes []byte) (err error) {
	varStoreBrandingReq := _StoreBrandingReq{}

	if err = json.Unmarshal(bytes, &varStoreBrandingReq); err == nil {
		*o = StoreBrandingReq(varStoreBrandingReq)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "bundleArtifact")
		delete(additionalProperties, "bundleGroup")
		delete(additionalProperties, "bundleVersion")
		delete(additionalProperties, "customOpenIdAppClazz")
		delete(additionalProperties, "customSsoAppClazz")
		delete(additionalProperties, "customSsoIdPAppClazz")
		delete(additionalProperties, "description")
		delete(additionalProperties, "name")
		delete(additionalProperties, "resource")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStoreBrandingReq struct {
	value *StoreBrandingReq
	isSet bool
}

func (v NullableStoreBrandingReq) Get() *StoreBrandingReq {
	return v.value
}

func (v *NullableStoreBrandingReq) Set(val *StoreBrandingReq) {
	v.value = val
	v.isSet = true
}

func (v NullableStoreBrandingReq) IsSet() bool {
	return v.isSet
}

func (v *NullableStoreBrandingReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoreBrandingReq(val *StoreBrandingReq) *NullableStoreBrandingReq {
	return &NullableStoreBrandingReq{value: val, isSet: true}
}

func (v NullableStoreBrandingReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoreBrandingReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


