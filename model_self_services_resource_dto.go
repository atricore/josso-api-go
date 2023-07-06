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

// SelfServicesResourceDTO struct for SelfServicesResourceDTO
type SelfServicesResourceDTO struct {
	Activation *ActivationDTO `json:"activation,omitempty"`
	Description *string `json:"description,omitempty"`
	ElementId *string `json:"elementId,omitempty"`
	Id *int64 `json:"id,omitempty"`
	Location *LocationDTO `json:"location,omitempty"`
	Name *string `json:"name,omitempty"`
	Secret *string `json:"secret,omitempty"`
	ServiceConnection *ServiceConnectionDTO `json:"serviceConnection,omitempty"`
	X *float64 `json:"x,omitempty"`
	Y *float64 `json:"y,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SelfServicesResourceDTO SelfServicesResourceDTO

// NewSelfServicesResourceDTO instantiates a new SelfServicesResourceDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelfServicesResourceDTO() *SelfServicesResourceDTO {
	this := SelfServicesResourceDTO{}
	return &this
}

// NewSelfServicesResourceDTOWithDefaults instantiates a new SelfServicesResourceDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelfServicesResourceDTOWithDefaults() *SelfServicesResourceDTO {
	this := SelfServicesResourceDTO{}
	return &this
}

// GetActivation returns the Activation field value if set, zero value otherwise.
func (o *SelfServicesResourceDTO) GetActivation() ActivationDTO {
	if o == nil || isNil(o.Activation) {
		var ret ActivationDTO
		return ret
	}
	return *o.Activation
}

// GetActivationOk returns a tuple with the Activation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServicesResourceDTO) GetActivationOk() (*ActivationDTO, bool) {
	if o == nil || isNil(o.Activation) {
    return nil, false
	}
	return o.Activation, true
}

// HasActivation returns a boolean if a field has been set.
func (o *SelfServicesResourceDTO) HasActivation() bool {
	if o != nil && !isNil(o.Activation) {
		return true
	}

	return false
}

// SetActivation gets a reference to the given ActivationDTO and assigns it to the Activation field.
func (o *SelfServicesResourceDTO) SetActivation(v ActivationDTO) {
	o.Activation = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SelfServicesResourceDTO) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServicesResourceDTO) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SelfServicesResourceDTO) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SelfServicesResourceDTO) SetDescription(v string) {
	o.Description = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *SelfServicesResourceDTO) GetElementId() string {
	if o == nil || isNil(o.ElementId) {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServicesResourceDTO) GetElementIdOk() (*string, bool) {
	if o == nil || isNil(o.ElementId) {
    return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *SelfServicesResourceDTO) HasElementId() bool {
	if o != nil && !isNil(o.ElementId) {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *SelfServicesResourceDTO) SetElementId(v string) {
	o.ElementId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SelfServicesResourceDTO) GetId() int64 {
	if o == nil || isNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServicesResourceDTO) GetIdOk() (*int64, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SelfServicesResourceDTO) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *SelfServicesResourceDTO) SetId(v int64) {
	o.Id = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *SelfServicesResourceDTO) GetLocation() LocationDTO {
	if o == nil || isNil(o.Location) {
		var ret LocationDTO
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServicesResourceDTO) GetLocationOk() (*LocationDTO, bool) {
	if o == nil || isNil(o.Location) {
    return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *SelfServicesResourceDTO) HasLocation() bool {
	if o != nil && !isNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given LocationDTO and assigns it to the Location field.
func (o *SelfServicesResourceDTO) SetLocation(v LocationDTO) {
	o.Location = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SelfServicesResourceDTO) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServicesResourceDTO) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SelfServicesResourceDTO) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SelfServicesResourceDTO) SetName(v string) {
	o.Name = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *SelfServicesResourceDTO) GetSecret() string {
	if o == nil || isNil(o.Secret) {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServicesResourceDTO) GetSecretOk() (*string, bool) {
	if o == nil || isNil(o.Secret) {
    return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *SelfServicesResourceDTO) HasSecret() bool {
	if o != nil && !isNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *SelfServicesResourceDTO) SetSecret(v string) {
	o.Secret = &v
}

// GetServiceConnection returns the ServiceConnection field value if set, zero value otherwise.
func (o *SelfServicesResourceDTO) GetServiceConnection() ServiceConnectionDTO {
	if o == nil || isNil(o.ServiceConnection) {
		var ret ServiceConnectionDTO
		return ret
	}
	return *o.ServiceConnection
}

// GetServiceConnectionOk returns a tuple with the ServiceConnection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServicesResourceDTO) GetServiceConnectionOk() (*ServiceConnectionDTO, bool) {
	if o == nil || isNil(o.ServiceConnection) {
    return nil, false
	}
	return o.ServiceConnection, true
}

// HasServiceConnection returns a boolean if a field has been set.
func (o *SelfServicesResourceDTO) HasServiceConnection() bool {
	if o != nil && !isNil(o.ServiceConnection) {
		return true
	}

	return false
}

// SetServiceConnection gets a reference to the given ServiceConnectionDTO and assigns it to the ServiceConnection field.
func (o *SelfServicesResourceDTO) SetServiceConnection(v ServiceConnectionDTO) {
	o.ServiceConnection = &v
}

// GetX returns the X field value if set, zero value otherwise.
func (o *SelfServicesResourceDTO) GetX() float64 {
	if o == nil || isNil(o.X) {
		var ret float64
		return ret
	}
	return *o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServicesResourceDTO) GetXOk() (*float64, bool) {
	if o == nil || isNil(o.X) {
    return nil, false
	}
	return o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *SelfServicesResourceDTO) HasX() bool {
	if o != nil && !isNil(o.X) {
		return true
	}

	return false
}

// SetX gets a reference to the given float64 and assigns it to the X field.
func (o *SelfServicesResourceDTO) SetX(v float64) {
	o.X = &v
}

// GetY returns the Y field value if set, zero value otherwise.
func (o *SelfServicesResourceDTO) GetY() float64 {
	if o == nil || isNil(o.Y) {
		var ret float64
		return ret
	}
	return *o.Y
}

// GetYOk returns a tuple with the Y field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServicesResourceDTO) GetYOk() (*float64, bool) {
	if o == nil || isNil(o.Y) {
    return nil, false
	}
	return o.Y, true
}

// HasY returns a boolean if a field has been set.
func (o *SelfServicesResourceDTO) HasY() bool {
	if o != nil && !isNil(o.Y) {
		return true
	}

	return false
}

// SetY gets a reference to the given float64 and assigns it to the Y field.
func (o *SelfServicesResourceDTO) SetY(v float64) {
	o.Y = &v
}

func (o SelfServicesResourceDTO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Activation) {
		toSerialize["activation"] = o.Activation
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.ElementId) {
		toSerialize["elementId"] = o.ElementId
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Secret) {
		toSerialize["secret"] = o.Secret
	}
	if !isNil(o.ServiceConnection) {
		toSerialize["serviceConnection"] = o.ServiceConnection
	}
	if !isNil(o.X) {
		toSerialize["x"] = o.X
	}
	if !isNil(o.Y) {
		toSerialize["y"] = o.Y
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SelfServicesResourceDTO) UnmarshalJSON(bytes []byte) (err error) {
	varSelfServicesResourceDTO := _SelfServicesResourceDTO{}

	if err = json.Unmarshal(bytes, &varSelfServicesResourceDTO); err == nil {
		*o = SelfServicesResourceDTO(varSelfServicesResourceDTO)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "activation")
		delete(additionalProperties, "description")
		delete(additionalProperties, "elementId")
		delete(additionalProperties, "id")
		delete(additionalProperties, "location")
		delete(additionalProperties, "name")
		delete(additionalProperties, "secret")
		delete(additionalProperties, "serviceConnection")
		delete(additionalProperties, "x")
		delete(additionalProperties, "y")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSelfServicesResourceDTO struct {
	value *SelfServicesResourceDTO
	isSet bool
}

func (v NullableSelfServicesResourceDTO) Get() *SelfServicesResourceDTO {
	return v.value
}

func (v *NullableSelfServicesResourceDTO) Set(val *SelfServicesResourceDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableSelfServicesResourceDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableSelfServicesResourceDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelfServicesResourceDTO(val *SelfServicesResourceDTO) *NullableSelfServicesResourceDTO {
	return &NullableSelfServicesResourceDTO{value: val, isSet: true}
}

func (v NullableSelfServicesResourceDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelfServicesResourceDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


