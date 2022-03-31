/*
Atricore Console :: Remote : API

# Atricore Console API

API version: 1.4.3-SNAPSHOT
Contact: sgonzalez@atricore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package jossoappi

import (
	"encoding/json"
	"time"
)

// IdentityApplianceDeploymentDTO struct for IdentityApplianceDeploymentDTO
type IdentityApplianceDeploymentDTO struct {
	DeployedRevision *int32 `json:"deployedRevision,omitempty"`
	DeploymentTime *time.Time `json:"deploymentTime,omitempty"`
	Description *string `json:"description,omitempty"`
	ElementId *string `json:"elementId,omitempty"`
	FeatureName *string `json:"featureName,omitempty"`
	FeatureUri *string `json:"featureUri,omitempty"`
	Id *int64 `json:"id,omitempty"`
	Idaus []IdentityApplianceUnitDTO `json:"idaus,omitempty"`
	State *string `json:"state,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentityApplianceDeploymentDTO IdentityApplianceDeploymentDTO

// NewIdentityApplianceDeploymentDTO instantiates a new IdentityApplianceDeploymentDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityApplianceDeploymentDTO() *IdentityApplianceDeploymentDTO {
	this := IdentityApplianceDeploymentDTO{}
	return &this
}

// NewIdentityApplianceDeploymentDTOWithDefaults instantiates a new IdentityApplianceDeploymentDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityApplianceDeploymentDTOWithDefaults() *IdentityApplianceDeploymentDTO {
	this := IdentityApplianceDeploymentDTO{}
	return &this
}

// GetDeployedRevision returns the DeployedRevision field value if set, zero value otherwise.
func (o *IdentityApplianceDeploymentDTO) GetDeployedRevision() int32 {
	if o == nil || o.DeployedRevision == nil {
		var ret int32
		return ret
	}
	return *o.DeployedRevision
}

// GetDeployedRevisionOk returns a tuple with the DeployedRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityApplianceDeploymentDTO) GetDeployedRevisionOk() (*int32, bool) {
	if o == nil || o.DeployedRevision == nil {
		return nil, false
	}
	return o.DeployedRevision, true
}

// HasDeployedRevision returns a boolean if a field has been set.
func (o *IdentityApplianceDeploymentDTO) HasDeployedRevision() bool {
	if o != nil && o.DeployedRevision != nil {
		return true
	}

	return false
}

// SetDeployedRevision gets a reference to the given int32 and assigns it to the DeployedRevision field.
func (o *IdentityApplianceDeploymentDTO) SetDeployedRevision(v int32) {
	o.DeployedRevision = &v
}

// GetDeploymentTime returns the DeploymentTime field value if set, zero value otherwise.
func (o *IdentityApplianceDeploymentDTO) GetDeploymentTime() time.Time {
	if o == nil || o.DeploymentTime == nil {
		var ret time.Time
		return ret
	}
	return *o.DeploymentTime
}

// GetDeploymentTimeOk returns a tuple with the DeploymentTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityApplianceDeploymentDTO) GetDeploymentTimeOk() (*time.Time, bool) {
	if o == nil || o.DeploymentTime == nil {
		return nil, false
	}
	return o.DeploymentTime, true
}

// HasDeploymentTime returns a boolean if a field has been set.
func (o *IdentityApplianceDeploymentDTO) HasDeploymentTime() bool {
	if o != nil && o.DeploymentTime != nil {
		return true
	}

	return false
}

// SetDeploymentTime gets a reference to the given time.Time and assigns it to the DeploymentTime field.
func (o *IdentityApplianceDeploymentDTO) SetDeploymentTime(v time.Time) {
	o.DeploymentTime = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IdentityApplianceDeploymentDTO) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityApplianceDeploymentDTO) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IdentityApplianceDeploymentDTO) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IdentityApplianceDeploymentDTO) SetDescription(v string) {
	o.Description = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *IdentityApplianceDeploymentDTO) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityApplianceDeploymentDTO) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *IdentityApplianceDeploymentDTO) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *IdentityApplianceDeploymentDTO) SetElementId(v string) {
	o.ElementId = &v
}

// GetFeatureName returns the FeatureName field value if set, zero value otherwise.
func (o *IdentityApplianceDeploymentDTO) GetFeatureName() string {
	if o == nil || o.FeatureName == nil {
		var ret string
		return ret
	}
	return *o.FeatureName
}

// GetFeatureNameOk returns a tuple with the FeatureName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityApplianceDeploymentDTO) GetFeatureNameOk() (*string, bool) {
	if o == nil || o.FeatureName == nil {
		return nil, false
	}
	return o.FeatureName, true
}

// HasFeatureName returns a boolean if a field has been set.
func (o *IdentityApplianceDeploymentDTO) HasFeatureName() bool {
	if o != nil && o.FeatureName != nil {
		return true
	}

	return false
}

// SetFeatureName gets a reference to the given string and assigns it to the FeatureName field.
func (o *IdentityApplianceDeploymentDTO) SetFeatureName(v string) {
	o.FeatureName = &v
}

// GetFeatureUri returns the FeatureUri field value if set, zero value otherwise.
func (o *IdentityApplianceDeploymentDTO) GetFeatureUri() string {
	if o == nil || o.FeatureUri == nil {
		var ret string
		return ret
	}
	return *o.FeatureUri
}

// GetFeatureUriOk returns a tuple with the FeatureUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityApplianceDeploymentDTO) GetFeatureUriOk() (*string, bool) {
	if o == nil || o.FeatureUri == nil {
		return nil, false
	}
	return o.FeatureUri, true
}

// HasFeatureUri returns a boolean if a field has been set.
func (o *IdentityApplianceDeploymentDTO) HasFeatureUri() bool {
	if o != nil && o.FeatureUri != nil {
		return true
	}

	return false
}

// SetFeatureUri gets a reference to the given string and assigns it to the FeatureUri field.
func (o *IdentityApplianceDeploymentDTO) SetFeatureUri(v string) {
	o.FeatureUri = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IdentityApplianceDeploymentDTO) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityApplianceDeploymentDTO) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IdentityApplianceDeploymentDTO) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *IdentityApplianceDeploymentDTO) SetId(v int64) {
	o.Id = &v
}

// GetIdaus returns the Idaus field value if set, zero value otherwise.
func (o *IdentityApplianceDeploymentDTO) GetIdaus() []IdentityApplianceUnitDTO {
	if o == nil || o.Idaus == nil {
		var ret []IdentityApplianceUnitDTO
		return ret
	}
	return o.Idaus
}

// GetIdausOk returns a tuple with the Idaus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityApplianceDeploymentDTO) GetIdausOk() ([]IdentityApplianceUnitDTO, bool) {
	if o == nil || o.Idaus == nil {
		return nil, false
	}
	return o.Idaus, true
}

// HasIdaus returns a boolean if a field has been set.
func (o *IdentityApplianceDeploymentDTO) HasIdaus() bool {
	if o != nil && o.Idaus != nil {
		return true
	}

	return false
}

// SetIdaus gets a reference to the given []IdentityApplianceUnitDTO and assigns it to the Idaus field.
func (o *IdentityApplianceDeploymentDTO) SetIdaus(v []IdentityApplianceUnitDTO) {
	o.Idaus = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *IdentityApplianceDeploymentDTO) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityApplianceDeploymentDTO) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *IdentityApplianceDeploymentDTO) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *IdentityApplianceDeploymentDTO) SetState(v string) {
	o.State = &v
}

func (o IdentityApplianceDeploymentDTO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeployedRevision != nil {
		toSerialize["deployedRevision"] = o.DeployedRevision
	}
	if o.DeploymentTime != nil {
		toSerialize["deploymentTime"] = o.DeploymentTime
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ElementId != nil {
		toSerialize["elementId"] = o.ElementId
	}
	if o.FeatureName != nil {
		toSerialize["featureName"] = o.FeatureName
	}
	if o.FeatureUri != nil {
		toSerialize["featureUri"] = o.FeatureUri
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Idaus != nil {
		toSerialize["idaus"] = o.Idaus
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IdentityApplianceDeploymentDTO) UnmarshalJSON(bytes []byte) (err error) {
	varIdentityApplianceDeploymentDTO := _IdentityApplianceDeploymentDTO{}

	if err = json.Unmarshal(bytes, &varIdentityApplianceDeploymentDTO); err == nil {
		*o = IdentityApplianceDeploymentDTO(varIdentityApplianceDeploymentDTO)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "deployedRevision")
		delete(additionalProperties, "deploymentTime")
		delete(additionalProperties, "description")
		delete(additionalProperties, "elementId")
		delete(additionalProperties, "featureName")
		delete(additionalProperties, "featureUri")
		delete(additionalProperties, "id")
		delete(additionalProperties, "idaus")
		delete(additionalProperties, "state")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentityApplianceDeploymentDTO struct {
	value *IdentityApplianceDeploymentDTO
	isSet bool
}

func (v NullableIdentityApplianceDeploymentDTO) Get() *IdentityApplianceDeploymentDTO {
	return v.value
}

func (v *NullableIdentityApplianceDeploymentDTO) Set(val *IdentityApplianceDeploymentDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityApplianceDeploymentDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityApplianceDeploymentDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityApplianceDeploymentDTO(val *IdentityApplianceDeploymentDTO) *NullableIdentityApplianceDeploymentDTO {
	return &NullableIdentityApplianceDeploymentDTO{value: val, isSet: true}
}

func (v NullableIdentityApplianceDeploymentDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityApplianceDeploymentDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


