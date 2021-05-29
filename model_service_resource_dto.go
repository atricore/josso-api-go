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

// ServiceResourceDTO struct for ServiceResourceDTO
type ServiceResourceDTO struct {
	Activation *ActivationDTO `json:"activation,omitempty"`
	Description *string `json:"description,omitempty"`
	ElementId *string `json:"elementId,omitempty"`
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	ServiceConnection *ServiceConnectionDTO `json:"serviceConnection,omitempty"`
	X *float64 `json:"x,omitempty"`
	Y *float64 `json:"y,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServiceResourceDTO ServiceResourceDTO

// NewServiceResourceDTO instantiates a new ServiceResourceDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceResourceDTO() *ServiceResourceDTO {
	this := ServiceResourceDTO{}
	return &this
}

// NewServiceResourceDTOWithDefaults instantiates a new ServiceResourceDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceResourceDTOWithDefaults() *ServiceResourceDTO {
	this := ServiceResourceDTO{}
	return &this
}

// GetActivation returns the Activation field value if set, zero value otherwise.
func (o *ServiceResourceDTO) GetActivation() ActivationDTO {
	if o == nil || o.Activation == nil {
		var ret ActivationDTO
		return ret
	}
	return *o.Activation
}

// GetActivationOk returns a tuple with the Activation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceResourceDTO) GetActivationOk() (*ActivationDTO, bool) {
	if o == nil || o.Activation == nil {
		return nil, false
	}
	return o.Activation, true
}

// HasActivation returns a boolean if a field has been set.
func (o *ServiceResourceDTO) HasActivation() bool {
	if o != nil && o.Activation != nil {
		return true
	}

	return false
}

// SetActivation gets a reference to the given ActivationDTO and assigns it to the Activation field.
func (o *ServiceResourceDTO) SetActivation(v ActivationDTO) {
	o.Activation = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ServiceResourceDTO) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceResourceDTO) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ServiceResourceDTO) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ServiceResourceDTO) SetDescription(v string) {
	o.Description = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *ServiceResourceDTO) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceResourceDTO) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *ServiceResourceDTO) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *ServiceResourceDTO) SetElementId(v string) {
	o.ElementId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ServiceResourceDTO) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceResourceDTO) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ServiceResourceDTO) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ServiceResourceDTO) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ServiceResourceDTO) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceResourceDTO) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ServiceResourceDTO) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ServiceResourceDTO) SetName(v string) {
	o.Name = &v
}

// GetServiceConnection returns the ServiceConnection field value if set, zero value otherwise.
func (o *ServiceResourceDTO) GetServiceConnection() ServiceConnectionDTO {
	if o == nil || o.ServiceConnection == nil {
		var ret ServiceConnectionDTO
		return ret
	}
	return *o.ServiceConnection
}

// GetServiceConnectionOk returns a tuple with the ServiceConnection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceResourceDTO) GetServiceConnectionOk() (*ServiceConnectionDTO, bool) {
	if o == nil || o.ServiceConnection == nil {
		return nil, false
	}
	return o.ServiceConnection, true
}

// HasServiceConnection returns a boolean if a field has been set.
func (o *ServiceResourceDTO) HasServiceConnection() bool {
	if o != nil && o.ServiceConnection != nil {
		return true
	}

	return false
}

// SetServiceConnection gets a reference to the given ServiceConnectionDTO and assigns it to the ServiceConnection field.
func (o *ServiceResourceDTO) SetServiceConnection(v ServiceConnectionDTO) {
	o.ServiceConnection = &v
}

// GetX returns the X field value if set, zero value otherwise.
func (o *ServiceResourceDTO) GetX() float64 {
	if o == nil || o.X == nil {
		var ret float64
		return ret
	}
	return *o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceResourceDTO) GetXOk() (*float64, bool) {
	if o == nil || o.X == nil {
		return nil, false
	}
	return o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *ServiceResourceDTO) HasX() bool {
	if o != nil && o.X != nil {
		return true
	}

	return false
}

// SetX gets a reference to the given float64 and assigns it to the X field.
func (o *ServiceResourceDTO) SetX(v float64) {
	o.X = &v
}

// GetY returns the Y field value if set, zero value otherwise.
func (o *ServiceResourceDTO) GetY() float64 {
	if o == nil || o.Y == nil {
		var ret float64
		return ret
	}
	return *o.Y
}

// GetYOk returns a tuple with the Y field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceResourceDTO) GetYOk() (*float64, bool) {
	if o == nil || o.Y == nil {
		return nil, false
	}
	return o.Y, true
}

// HasY returns a boolean if a field has been set.
func (o *ServiceResourceDTO) HasY() bool {
	if o != nil && o.Y != nil {
		return true
	}

	return false
}

// SetY gets a reference to the given float64 and assigns it to the Y field.
func (o *ServiceResourceDTO) SetY(v float64) {
	o.Y = &v
}

func (o ServiceResourceDTO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Activation != nil {
		toSerialize["activation"] = o.Activation
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
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
	if o.ServiceConnection != nil {
		toSerialize["serviceConnection"] = o.ServiceConnection
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

func (o *ServiceResourceDTO) UnmarshalJSON(bytes []byte) (err error) {
	varServiceResourceDTO := _ServiceResourceDTO{}

	if err = json.Unmarshal(bytes, &varServiceResourceDTO); err == nil {
		*o = ServiceResourceDTO(varServiceResourceDTO)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "activation")
		delete(additionalProperties, "description")
		delete(additionalProperties, "elementId")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "serviceConnection")
		delete(additionalProperties, "x")
		delete(additionalProperties, "y")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServiceResourceDTO struct {
	value *ServiceResourceDTO
	isSet bool
}

func (v NullableServiceResourceDTO) Get() *ServiceResourceDTO {
	return v.value
}

func (v *NullableServiceResourceDTO) Set(val *ServiceResourceDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceResourceDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceResourceDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceResourceDTO(val *ServiceResourceDTO) *NullableServiceResourceDTO {
	return &NullableServiceResourceDTO{value: val, isSet: true}
}

func (v NullableServiceResourceDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceResourceDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


