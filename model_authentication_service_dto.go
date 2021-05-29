/*
 * Atricore Console :: Remote : API
 *
 * # Atricore Console API
 *
 * API version: 1.4.3-SNAPSHOT
 * Contact: sgonzalez@atricore.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package jossoappi

import (
	"encoding/json"
)

// AuthenticationServiceDTO struct for AuthenticationServiceDTO
type AuthenticationServiceDTO struct {
	DelegatedAuthentications *[]DelegatedAuthenticationDTO `json:"delegatedAuthentications,omitempty"`
	Description *string `json:"description,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	ElementId *string `json:"elementId,omitempty"`
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	X *float64 `json:"x,omitempty"`
	Y *float64 `json:"y,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticationServiceDTO AuthenticationServiceDTO

// NewAuthenticationServiceDTO instantiates a new AuthenticationServiceDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticationServiceDTO() *AuthenticationServiceDTO {
	this := AuthenticationServiceDTO{}
	return &this
}

// NewAuthenticationServiceDTOWithDefaults instantiates a new AuthenticationServiceDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticationServiceDTOWithDefaults() *AuthenticationServiceDTO {
	this := AuthenticationServiceDTO{}
	return &this
}

// GetDelegatedAuthentications returns the DelegatedAuthentications field value if set, zero value otherwise.
func (o *AuthenticationServiceDTO) GetDelegatedAuthentications() []DelegatedAuthenticationDTO {
	if o == nil || o.DelegatedAuthentications == nil {
		var ret []DelegatedAuthenticationDTO
		return ret
	}
	return *o.DelegatedAuthentications
}

// GetDelegatedAuthenticationsOk returns a tuple with the DelegatedAuthentications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationServiceDTO) GetDelegatedAuthenticationsOk() (*[]DelegatedAuthenticationDTO, bool) {
	if o == nil || o.DelegatedAuthentications == nil {
		return nil, false
	}
	return o.DelegatedAuthentications, true
}

// HasDelegatedAuthentications returns a boolean if a field has been set.
func (o *AuthenticationServiceDTO) HasDelegatedAuthentications() bool {
	if o != nil && o.DelegatedAuthentications != nil {
		return true
	}

	return false
}

// SetDelegatedAuthentications gets a reference to the given []DelegatedAuthenticationDTO and assigns it to the DelegatedAuthentications field.
func (o *AuthenticationServiceDTO) SetDelegatedAuthentications(v []DelegatedAuthenticationDTO) {
	o.DelegatedAuthentications = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AuthenticationServiceDTO) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationServiceDTO) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AuthenticationServiceDTO) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AuthenticationServiceDTO) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *AuthenticationServiceDTO) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationServiceDTO) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *AuthenticationServiceDTO) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *AuthenticationServiceDTO) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *AuthenticationServiceDTO) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationServiceDTO) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *AuthenticationServiceDTO) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *AuthenticationServiceDTO) SetElementId(v string) {
	o.ElementId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AuthenticationServiceDTO) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationServiceDTO) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AuthenticationServiceDTO) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *AuthenticationServiceDTO) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AuthenticationServiceDTO) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationServiceDTO) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AuthenticationServiceDTO) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AuthenticationServiceDTO) SetName(v string) {
	o.Name = &v
}

// GetX returns the X field value if set, zero value otherwise.
func (o *AuthenticationServiceDTO) GetX() float64 {
	if o == nil || o.X == nil {
		var ret float64
		return ret
	}
	return *o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationServiceDTO) GetXOk() (*float64, bool) {
	if o == nil || o.X == nil {
		return nil, false
	}
	return o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *AuthenticationServiceDTO) HasX() bool {
	if o != nil && o.X != nil {
		return true
	}

	return false
}

// SetX gets a reference to the given float64 and assigns it to the X field.
func (o *AuthenticationServiceDTO) SetX(v float64) {
	o.X = &v
}

// GetY returns the Y field value if set, zero value otherwise.
func (o *AuthenticationServiceDTO) GetY() float64 {
	if o == nil || o.Y == nil {
		var ret float64
		return ret
	}
	return *o.Y
}

// GetYOk returns a tuple with the Y field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationServiceDTO) GetYOk() (*float64, bool) {
	if o == nil || o.Y == nil {
		return nil, false
	}
	return o.Y, true
}

// HasY returns a boolean if a field has been set.
func (o *AuthenticationServiceDTO) HasY() bool {
	if o != nil && o.Y != nil {
		return true
	}

	return false
}

// SetY gets a reference to the given float64 and assigns it to the Y field.
func (o *AuthenticationServiceDTO) SetY(v float64) {
	o.Y = &v
}

func (o AuthenticationServiceDTO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DelegatedAuthentications != nil {
		toSerialize["delegatedAuthentications"] = o.DelegatedAuthentications
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.ElementId != nil {
		toSerialize["elementId"] = o.ElementId
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.X != nil {
		toSerialize["x"] = o.X
	}
	if o.Y != nil {
		toSerialize["y"] = o.Y
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AuthenticationServiceDTO) UnmarshalJSON(bytes []byte) (err error) {
	varAuthenticationServiceDTO := _AuthenticationServiceDTO{}

	if err = json.Unmarshal(bytes, &varAuthenticationServiceDTO); err == nil {
		*o = AuthenticationServiceDTO(varAuthenticationServiceDTO)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "delegatedAuthentications")
		delete(additionalProperties, "description")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "elementId")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "x")
		delete(additionalProperties, "y")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuthenticationServiceDTO struct {
	value *AuthenticationServiceDTO
	isSet bool
}

func (v NullableAuthenticationServiceDTO) Get() *AuthenticationServiceDTO {
	return v.value
}

func (v *NullableAuthenticationServiceDTO) Set(val *AuthenticationServiceDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticationServiceDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticationServiceDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticationServiceDTO(val *AuthenticationServiceDTO) *NullableAuthenticationServiceDTO {
	return &NullableAuthenticationServiceDTO{value: val, isSet: true}
}

func (v NullableAuthenticationServiceDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticationServiceDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


