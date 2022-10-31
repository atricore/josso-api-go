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

// BrandingDefinitionDTO struct for BrandingDefinitionDTO
type BrandingDefinitionDTO struct {
	DefaultLocale *string `json:"defaultLocale,omitempty"`
	Description *string `json:"description,omitempty"`
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
	WebBrandingId *string `json:"webBrandingId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BrandingDefinitionDTO BrandingDefinitionDTO

// NewBrandingDefinitionDTO instantiates a new BrandingDefinitionDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrandingDefinitionDTO() *BrandingDefinitionDTO {
	this := BrandingDefinitionDTO{}
	return &this
}

// NewBrandingDefinitionDTOWithDefaults instantiates a new BrandingDefinitionDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrandingDefinitionDTOWithDefaults() *BrandingDefinitionDTO {
	this := BrandingDefinitionDTO{}
	return &this
}

// GetDefaultLocale returns the DefaultLocale field value if set, zero value otherwise.
func (o *BrandingDefinitionDTO) GetDefaultLocale() string {
	if o == nil || o.DefaultLocale == nil {
		var ret string
		return ret
	}
	return *o.DefaultLocale
}

// GetDefaultLocaleOk returns a tuple with the DefaultLocale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandingDefinitionDTO) GetDefaultLocaleOk() (*string, bool) {
	if o == nil || o.DefaultLocale == nil {
		return nil, false
	}
	return o.DefaultLocale, true
}

// HasDefaultLocale returns a boolean if a field has been set.
func (o *BrandingDefinitionDTO) HasDefaultLocale() bool {
	if o != nil && o.DefaultLocale != nil {
		return true
	}

	return false
}

// SetDefaultLocale gets a reference to the given string and assigns it to the DefaultLocale field.
func (o *BrandingDefinitionDTO) SetDefaultLocale(v string) {
	o.DefaultLocale = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BrandingDefinitionDTO) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandingDefinitionDTO) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BrandingDefinitionDTO) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BrandingDefinitionDTO) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BrandingDefinitionDTO) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandingDefinitionDTO) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BrandingDefinitionDTO) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *BrandingDefinitionDTO) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BrandingDefinitionDTO) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandingDefinitionDTO) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BrandingDefinitionDTO) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BrandingDefinitionDTO) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BrandingDefinitionDTO) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandingDefinitionDTO) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BrandingDefinitionDTO) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BrandingDefinitionDTO) SetType(v string) {
	o.Type = &v
}

// GetWebBrandingId returns the WebBrandingId field value if set, zero value otherwise.
func (o *BrandingDefinitionDTO) GetWebBrandingId() string {
	if o == nil || o.WebBrandingId == nil {
		var ret string
		return ret
	}
	return *o.WebBrandingId
}

// GetWebBrandingIdOk returns a tuple with the WebBrandingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandingDefinitionDTO) GetWebBrandingIdOk() (*string, bool) {
	if o == nil || o.WebBrandingId == nil {
		return nil, false
	}
	return o.WebBrandingId, true
}

// HasWebBrandingId returns a boolean if a field has been set.
func (o *BrandingDefinitionDTO) HasWebBrandingId() bool {
	if o != nil && o.WebBrandingId != nil {
		return true
	}

	return false
}

// SetWebBrandingId gets a reference to the given string and assigns it to the WebBrandingId field.
func (o *BrandingDefinitionDTO) SetWebBrandingId(v string) {
	o.WebBrandingId = &v
}

func (o BrandingDefinitionDTO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DefaultLocale != nil {
		toSerialize["defaultLocale"] = o.DefaultLocale
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.WebBrandingId != nil {
		toSerialize["webBrandingId"] = o.WebBrandingId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BrandingDefinitionDTO) UnmarshalJSON(bytes []byte) (err error) {
	varBrandingDefinitionDTO := _BrandingDefinitionDTO{}

	if err = json.Unmarshal(bytes, &varBrandingDefinitionDTO); err == nil {
		*o = BrandingDefinitionDTO(varBrandingDefinitionDTO)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "defaultLocale")
		delete(additionalProperties, "description")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "webBrandingId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBrandingDefinitionDTO struct {
	value *BrandingDefinitionDTO
	isSet bool
}

func (v NullableBrandingDefinitionDTO) Get() *BrandingDefinitionDTO {
	return v.value
}

func (v *NullableBrandingDefinitionDTO) Set(val *BrandingDefinitionDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableBrandingDefinitionDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableBrandingDefinitionDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrandingDefinitionDTO(val *BrandingDefinitionDTO) *NullableBrandingDefinitionDTO {
	return &NullableBrandingDefinitionDTO{value: val, isSet: true}
}

func (v NullableBrandingDefinitionDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBrandingDefinitionDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


