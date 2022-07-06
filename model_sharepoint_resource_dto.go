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

// SharepointResourceDTO struct for SharepointResourceDTO
type SharepointResourceDTO struct {
	Activation *ActivationDTO `json:"activation,omitempty"`
	DefaultResource *string `json:"defaultResource,omitempty"`
	Description *string `json:"description,omitempty"`
	ElementId *string `json:"elementId,omitempty"`
	Id *int64 `json:"id,omitempty"`
	IgnoredWebResources []string `json:"ignoredWebResources,omitempty"`
	Name *string `json:"name,omitempty"`
	PartnerAppLocation *LocationDTO `json:"partnerAppLocation,omitempty"`
	RelayingPartyLocation *LocationDTO `json:"relayingPartyLocation,omitempty"`
	ServiceConnection *ServiceConnectionDTO `json:"serviceConnection,omitempty"`
	SloLocation *LocationDTO `json:"sloLocation,omitempty"`
	SloLocationEnabled *bool `json:"sloLocationEnabled,omitempty"`
	StsEncryptingCertSubject *string `json:"stsEncryptingCertSubject,omitempty"`
	StsMetadata *ResourceDTO `json:"stsMetadata,omitempty"`
	StsSigningCertSubject *string `json:"stsSigningCertSubject,omitempty"`
	X *float64 `json:"x,omitempty"`
	Y *float64 `json:"y,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SharepointResourceDTO SharepointResourceDTO

// NewSharepointResourceDTO instantiates a new SharepointResourceDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSharepointResourceDTO() *SharepointResourceDTO {
	this := SharepointResourceDTO{}
	return &this
}

// NewSharepointResourceDTOWithDefaults instantiates a new SharepointResourceDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSharepointResourceDTOWithDefaults() *SharepointResourceDTO {
	this := SharepointResourceDTO{}
	return &this
}

// GetActivation returns the Activation field value if set, zero value otherwise.
func (o *SharepointResourceDTO) GetActivation() ActivationDTO {
	if o == nil || o.Activation == nil {
		var ret ActivationDTO
		return ret
	}
	return *o.Activation
}

// GetActivationOk returns a tuple with the Activation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharepointResourceDTO) GetActivationOk() (*ActivationDTO, bool) {
	if o == nil || o.Activation == nil {
		return nil, false
	}
	return o.Activation, true
}

// HasActivation returns a boolean if a field has been set.
func (o *SharepointResourceDTO) HasActivation() bool {
	if o != nil && o.Activation != nil {
		return true
	}

	return false
}

// SetActivation gets a reference to the given ActivationDTO and assigns it to the Activation field.
func (o *SharepointResourceDTO) SetActivation(v ActivationDTO) {
	o.Activation = &v
}

// GetDefaultResource returns the DefaultResource field value if set, zero value otherwise.
func (o *SharepointResourceDTO) GetDefaultResource() string {
	if o == nil || o.DefaultResource == nil {
		var ret string
		return ret
	}
	return *o.DefaultResource
}

// GetDefaultResourceOk returns a tuple with the DefaultResource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharepointResourceDTO) GetDefaultResourceOk() (*string, bool) {
	if o == nil || o.DefaultResource == nil {
		return nil, false
	}
	return o.DefaultResource, true
}

// HasDefaultResource returns a boolean if a field has been set.
func (o *SharepointResourceDTO) HasDefaultResource() bool {
	if o != nil && o.DefaultResource != nil {
		return true
	}

	return false
}

// SetDefaultResource gets a reference to the given string and assigns it to the DefaultResource field.
func (o *SharepointResourceDTO) SetDefaultResource(v string) {
	o.DefaultResource = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SharepointResourceDTO) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharepointResourceDTO) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SharepointResourceDTO) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SharepointResourceDTO) SetDescription(v string) {
	o.Description = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *SharepointResourceDTO) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharepointResourceDTO) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *SharepointResourceDTO) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *SharepointResourceDTO) SetElementId(v string) {
	o.ElementId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SharepointResourceDTO) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharepointResourceDTO) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SharepointResourceDTO) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *SharepointResourceDTO) SetId(v int64) {
	o.Id = &v
}

// GetIgnoredWebResources returns the IgnoredWebResources field value if set, zero value otherwise.
func (o *SharepointResourceDTO) GetIgnoredWebResources() []string {
	if o == nil || o.IgnoredWebResources == nil {
		var ret []string
		return ret
	}
	return o.IgnoredWebResources
}

// GetIgnoredWebResourcesOk returns a tuple with the IgnoredWebResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharepointResourceDTO) GetIgnoredWebResourcesOk() ([]string, bool) {
	if o == nil || o.IgnoredWebResources == nil {
		return nil, false
	}
	return o.IgnoredWebResources, true
}

// HasIgnoredWebResources returns a boolean if a field has been set.
func (o *SharepointResourceDTO) HasIgnoredWebResources() bool {
	if o != nil && o.IgnoredWebResources != nil {
		return true
	}

	return false
}

// SetIgnoredWebResources gets a reference to the given []string and assigns it to the IgnoredWebResources field.
func (o *SharepointResourceDTO) SetIgnoredWebResources(v []string) {
	o.IgnoredWebResources = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SharepointResourceDTO) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharepointResourceDTO) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SharepointResourceDTO) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SharepointResourceDTO) SetName(v string) {
	o.Name = &v
}

// GetPartnerAppLocation returns the PartnerAppLocation field value if set, zero value otherwise.
func (o *SharepointResourceDTO) GetPartnerAppLocation() LocationDTO {
	if o == nil || o.PartnerAppLocation == nil {
		var ret LocationDTO
		return ret
	}
	return *o.PartnerAppLocation
}

// GetPartnerAppLocationOk returns a tuple with the PartnerAppLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharepointResourceDTO) GetPartnerAppLocationOk() (*LocationDTO, bool) {
	if o == nil || o.PartnerAppLocation == nil {
		return nil, false
	}
	return o.PartnerAppLocation, true
}

// HasPartnerAppLocation returns a boolean if a field has been set.
func (o *SharepointResourceDTO) HasPartnerAppLocation() bool {
	if o != nil && o.PartnerAppLocation != nil {
		return true
	}

	return false
}

// SetPartnerAppLocation gets a reference to the given LocationDTO and assigns it to the PartnerAppLocation field.
func (o *SharepointResourceDTO) SetPartnerAppLocation(v LocationDTO) {
	o.PartnerAppLocation = &v
}

// GetRelayingPartyLocation returns the RelayingPartyLocation field value if set, zero value otherwise.
func (o *SharepointResourceDTO) GetRelayingPartyLocation() LocationDTO {
	if o == nil || o.RelayingPartyLocation == nil {
		var ret LocationDTO
		return ret
	}
	return *o.RelayingPartyLocation
}

// GetRelayingPartyLocationOk returns a tuple with the RelayingPartyLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharepointResourceDTO) GetRelayingPartyLocationOk() (*LocationDTO, bool) {
	if o == nil || o.RelayingPartyLocation == nil {
		return nil, false
	}
	return o.RelayingPartyLocation, true
}

// HasRelayingPartyLocation returns a boolean if a field has been set.
func (o *SharepointResourceDTO) HasRelayingPartyLocation() bool {
	if o != nil && o.RelayingPartyLocation != nil {
		return true
	}

	return false
}

// SetRelayingPartyLocation gets a reference to the given LocationDTO and assigns it to the RelayingPartyLocation field.
func (o *SharepointResourceDTO) SetRelayingPartyLocation(v LocationDTO) {
	o.RelayingPartyLocation = &v
}

// GetServiceConnection returns the ServiceConnection field value if set, zero value otherwise.
func (o *SharepointResourceDTO) GetServiceConnection() ServiceConnectionDTO {
	if o == nil || o.ServiceConnection == nil {
		var ret ServiceConnectionDTO
		return ret
	}
	return *o.ServiceConnection
}

// GetServiceConnectionOk returns a tuple with the ServiceConnection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharepointResourceDTO) GetServiceConnectionOk() (*ServiceConnectionDTO, bool) {
	if o == nil || o.ServiceConnection == nil {
		return nil, false
	}
	return o.ServiceConnection, true
}

// HasServiceConnection returns a boolean if a field has been set.
func (o *SharepointResourceDTO) HasServiceConnection() bool {
	if o != nil && o.ServiceConnection != nil {
		return true
	}

	return false
}

// SetServiceConnection gets a reference to the given ServiceConnectionDTO and assigns it to the ServiceConnection field.
func (o *SharepointResourceDTO) SetServiceConnection(v ServiceConnectionDTO) {
	o.ServiceConnection = &v
}

// GetSloLocation returns the SloLocation field value if set, zero value otherwise.
func (o *SharepointResourceDTO) GetSloLocation() LocationDTO {
	if o == nil || o.SloLocation == nil {
		var ret LocationDTO
		return ret
	}
	return *o.SloLocation
}

// GetSloLocationOk returns a tuple with the SloLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharepointResourceDTO) GetSloLocationOk() (*LocationDTO, bool) {
	if o == nil || o.SloLocation == nil {
		return nil, false
	}
	return o.SloLocation, true
}

// HasSloLocation returns a boolean if a field has been set.
func (o *SharepointResourceDTO) HasSloLocation() bool {
	if o != nil && o.SloLocation != nil {
		return true
	}

	return false
}

// SetSloLocation gets a reference to the given LocationDTO and assigns it to the SloLocation field.
func (o *SharepointResourceDTO) SetSloLocation(v LocationDTO) {
	o.SloLocation = &v
}

// GetSloLocationEnabled returns the SloLocationEnabled field value if set, zero value otherwise.
func (o *SharepointResourceDTO) GetSloLocationEnabled() bool {
	if o == nil || o.SloLocationEnabled == nil {
		var ret bool
		return ret
	}
	return *o.SloLocationEnabled
}

// GetSloLocationEnabledOk returns a tuple with the SloLocationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharepointResourceDTO) GetSloLocationEnabledOk() (*bool, bool) {
	if o == nil || o.SloLocationEnabled == nil {
		return nil, false
	}
	return o.SloLocationEnabled, true
}

// HasSloLocationEnabled returns a boolean if a field has been set.
func (o *SharepointResourceDTO) HasSloLocationEnabled() bool {
	if o != nil && o.SloLocationEnabled != nil {
		return true
	}

	return false
}

// SetSloLocationEnabled gets a reference to the given bool and assigns it to the SloLocationEnabled field.
func (o *SharepointResourceDTO) SetSloLocationEnabled(v bool) {
	o.SloLocationEnabled = &v
}

// GetStsEncryptingCertSubject returns the StsEncryptingCertSubject field value if set, zero value otherwise.
func (o *SharepointResourceDTO) GetStsEncryptingCertSubject() string {
	if o == nil || o.StsEncryptingCertSubject == nil {
		var ret string
		return ret
	}
	return *o.StsEncryptingCertSubject
}

// GetStsEncryptingCertSubjectOk returns a tuple with the StsEncryptingCertSubject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharepointResourceDTO) GetStsEncryptingCertSubjectOk() (*string, bool) {
	if o == nil || o.StsEncryptingCertSubject == nil {
		return nil, false
	}
	return o.StsEncryptingCertSubject, true
}

// HasStsEncryptingCertSubject returns a boolean if a field has been set.
func (o *SharepointResourceDTO) HasStsEncryptingCertSubject() bool {
	if o != nil && o.StsEncryptingCertSubject != nil {
		return true
	}

	return false
}

// SetStsEncryptingCertSubject gets a reference to the given string and assigns it to the StsEncryptingCertSubject field.
func (o *SharepointResourceDTO) SetStsEncryptingCertSubject(v string) {
	o.StsEncryptingCertSubject = &v
}

// GetStsMetadata returns the StsMetadata field value if set, zero value otherwise.
func (o *SharepointResourceDTO) GetStsMetadata() ResourceDTO {
	if o == nil || o.StsMetadata == nil {
		var ret ResourceDTO
		return ret
	}
	return *o.StsMetadata
}

// GetStsMetadataOk returns a tuple with the StsMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharepointResourceDTO) GetStsMetadataOk() (*ResourceDTO, bool) {
	if o == nil || o.StsMetadata == nil {
		return nil, false
	}
	return o.StsMetadata, true
}

// HasStsMetadata returns a boolean if a field has been set.
func (o *SharepointResourceDTO) HasStsMetadata() bool {
	if o != nil && o.StsMetadata != nil {
		return true
	}

	return false
}

// SetStsMetadata gets a reference to the given ResourceDTO and assigns it to the StsMetadata field.
func (o *SharepointResourceDTO) SetStsMetadata(v ResourceDTO) {
	o.StsMetadata = &v
}

// GetStsSigningCertSubject returns the StsSigningCertSubject field value if set, zero value otherwise.
func (o *SharepointResourceDTO) GetStsSigningCertSubject() string {
	if o == nil || o.StsSigningCertSubject == nil {
		var ret string
		return ret
	}
	return *o.StsSigningCertSubject
}

// GetStsSigningCertSubjectOk returns a tuple with the StsSigningCertSubject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharepointResourceDTO) GetStsSigningCertSubjectOk() (*string, bool) {
	if o == nil || o.StsSigningCertSubject == nil {
		return nil, false
	}
	return o.StsSigningCertSubject, true
}

// HasStsSigningCertSubject returns a boolean if a field has been set.
func (o *SharepointResourceDTO) HasStsSigningCertSubject() bool {
	if o != nil && o.StsSigningCertSubject != nil {
		return true
	}

	return false
}

// SetStsSigningCertSubject gets a reference to the given string and assigns it to the StsSigningCertSubject field.
func (o *SharepointResourceDTO) SetStsSigningCertSubject(v string) {
	o.StsSigningCertSubject = &v
}

// GetX returns the X field value if set, zero value otherwise.
func (o *SharepointResourceDTO) GetX() float64 {
	if o == nil || o.X == nil {
		var ret float64
		return ret
	}
	return *o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharepointResourceDTO) GetXOk() (*float64, bool) {
	if o == nil || o.X == nil {
		return nil, false
	}
	return o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *SharepointResourceDTO) HasX() bool {
	if o != nil && o.X != nil {
		return true
	}

	return false
}

// SetX gets a reference to the given float64 and assigns it to the X field.
func (o *SharepointResourceDTO) SetX(v float64) {
	o.X = &v
}

// GetY returns the Y field value if set, zero value otherwise.
func (o *SharepointResourceDTO) GetY() float64 {
	if o == nil || o.Y == nil {
		var ret float64
		return ret
	}
	return *o.Y
}

// GetYOk returns a tuple with the Y field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharepointResourceDTO) GetYOk() (*float64, bool) {
	if o == nil || o.Y == nil {
		return nil, false
	}
	return o.Y, true
}

// HasY returns a boolean if a field has been set.
func (o *SharepointResourceDTO) HasY() bool {
	if o != nil && o.Y != nil {
		return true
	}

	return false
}

// SetY gets a reference to the given float64 and assigns it to the Y field.
func (o *SharepointResourceDTO) SetY(v float64) {
	o.Y = &v
}

func (o SharepointResourceDTO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Activation != nil {
		toSerialize["activation"] = o.Activation
	}
	if o.DefaultResource != nil {
		toSerialize["defaultResource"] = o.DefaultResource
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
	if o.IgnoredWebResources != nil {
		toSerialize["ignoredWebResources"] = o.IgnoredWebResources
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.PartnerAppLocation != nil {
		toSerialize["partnerAppLocation"] = o.PartnerAppLocation
	}
	if o.RelayingPartyLocation != nil {
		toSerialize["relayingPartyLocation"] = o.RelayingPartyLocation
	}
	if o.ServiceConnection != nil {
		toSerialize["serviceConnection"] = o.ServiceConnection
	}
	if o.SloLocation != nil {
		toSerialize["sloLocation"] = o.SloLocation
	}
	if o.SloLocationEnabled != nil {
		toSerialize["sloLocationEnabled"] = o.SloLocationEnabled
	}
	if o.StsEncryptingCertSubject != nil {
		toSerialize["stsEncryptingCertSubject"] = o.StsEncryptingCertSubject
	}
	if o.StsMetadata != nil {
		toSerialize["stsMetadata"] = o.StsMetadata
	}
	if o.StsSigningCertSubject != nil {
		toSerialize["stsSigningCertSubject"] = o.StsSigningCertSubject
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

func (o *SharepointResourceDTO) UnmarshalJSON(bytes []byte) (err error) {
	varSharepointResourceDTO := _SharepointResourceDTO{}

	if err = json.Unmarshal(bytes, &varSharepointResourceDTO); err == nil {
		*o = SharepointResourceDTO(varSharepointResourceDTO)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "activation")
		delete(additionalProperties, "defaultResource")
		delete(additionalProperties, "description")
		delete(additionalProperties, "elementId")
		delete(additionalProperties, "id")
		delete(additionalProperties, "ignoredWebResources")
		delete(additionalProperties, "name")
		delete(additionalProperties, "partnerAppLocation")
		delete(additionalProperties, "relayingPartyLocation")
		delete(additionalProperties, "serviceConnection")
		delete(additionalProperties, "sloLocation")
		delete(additionalProperties, "sloLocationEnabled")
		delete(additionalProperties, "stsEncryptingCertSubject")
		delete(additionalProperties, "stsMetadata")
		delete(additionalProperties, "stsSigningCertSubject")
		delete(additionalProperties, "x")
		delete(additionalProperties, "y")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSharepointResourceDTO struct {
	value *SharepointResourceDTO
	isSet bool
}

func (v NullableSharepointResourceDTO) Get() *SharepointResourceDTO {
	return v.value
}

func (v *NullableSharepointResourceDTO) Set(val *SharepointResourceDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableSharepointResourceDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableSharepointResourceDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSharepointResourceDTO(val *SharepointResourceDTO) *NullableSharepointResourceDTO {
	return &NullableSharepointResourceDTO{value: val, isSet: true}
}

func (v NullableSharepointResourceDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSharepointResourceDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


