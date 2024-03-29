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

// BundleDescr struct for BundleDescr
type BundleDescr struct {
	Id *int64 `json:"id,omitempty"`
	Level *string `json:"level,omitempty"`
	Location *string `json:"location,omitempty"`
	Name *string `json:"name,omitempty"`
	State *string `json:"state,omitempty"`
	SymbolicName *string `json:"symbolicName,omitempty"`
	UpdateLocation *string `json:"updateLocation,omitempty"`
	Version *string `json:"version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BundleDescr BundleDescr

// NewBundleDescr instantiates a new BundleDescr object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBundleDescr() *BundleDescr {
	this := BundleDescr{}
	return &this
}

// NewBundleDescrWithDefaults instantiates a new BundleDescr object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBundleDescrWithDefaults() *BundleDescr {
	this := BundleDescr{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BundleDescr) GetId() int64 {
	if o == nil || isNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleDescr) GetIdOk() (*int64, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BundleDescr) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *BundleDescr) SetId(v int64) {
	o.Id = &v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *BundleDescr) GetLevel() string {
	if o == nil || isNil(o.Level) {
		var ret string
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleDescr) GetLevelOk() (*string, bool) {
	if o == nil || isNil(o.Level) {
    return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *BundleDescr) HasLevel() bool {
	if o != nil && !isNil(o.Level) {
		return true
	}

	return false
}

// SetLevel gets a reference to the given string and assigns it to the Level field.
func (o *BundleDescr) SetLevel(v string) {
	o.Level = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *BundleDescr) GetLocation() string {
	if o == nil || isNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleDescr) GetLocationOk() (*string, bool) {
	if o == nil || isNil(o.Location) {
    return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *BundleDescr) HasLocation() bool {
	if o != nil && !isNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *BundleDescr) SetLocation(v string) {
	o.Location = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BundleDescr) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleDescr) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BundleDescr) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BundleDescr) SetName(v string) {
	o.Name = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *BundleDescr) GetState() string {
	if o == nil || isNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleDescr) GetStateOk() (*string, bool) {
	if o == nil || isNil(o.State) {
    return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *BundleDescr) HasState() bool {
	if o != nil && !isNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *BundleDescr) SetState(v string) {
	o.State = &v
}

// GetSymbolicName returns the SymbolicName field value if set, zero value otherwise.
func (o *BundleDescr) GetSymbolicName() string {
	if o == nil || isNil(o.SymbolicName) {
		var ret string
		return ret
	}
	return *o.SymbolicName
}

// GetSymbolicNameOk returns a tuple with the SymbolicName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleDescr) GetSymbolicNameOk() (*string, bool) {
	if o == nil || isNil(o.SymbolicName) {
    return nil, false
	}
	return o.SymbolicName, true
}

// HasSymbolicName returns a boolean if a field has been set.
func (o *BundleDescr) HasSymbolicName() bool {
	if o != nil && !isNil(o.SymbolicName) {
		return true
	}

	return false
}

// SetSymbolicName gets a reference to the given string and assigns it to the SymbolicName field.
func (o *BundleDescr) SetSymbolicName(v string) {
	o.SymbolicName = &v
}

// GetUpdateLocation returns the UpdateLocation field value if set, zero value otherwise.
func (o *BundleDescr) GetUpdateLocation() string {
	if o == nil || isNil(o.UpdateLocation) {
		var ret string
		return ret
	}
	return *o.UpdateLocation
}

// GetUpdateLocationOk returns a tuple with the UpdateLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleDescr) GetUpdateLocationOk() (*string, bool) {
	if o == nil || isNil(o.UpdateLocation) {
    return nil, false
	}
	return o.UpdateLocation, true
}

// HasUpdateLocation returns a boolean if a field has been set.
func (o *BundleDescr) HasUpdateLocation() bool {
	if o != nil && !isNil(o.UpdateLocation) {
		return true
	}

	return false
}

// SetUpdateLocation gets a reference to the given string and assigns it to the UpdateLocation field.
func (o *BundleDescr) SetUpdateLocation(v string) {
	o.UpdateLocation = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *BundleDescr) GetVersion() string {
	if o == nil || isNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleDescr) GetVersionOk() (*string, bool) {
	if o == nil || isNil(o.Version) {
    return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *BundleDescr) HasVersion() bool {
	if o != nil && !isNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *BundleDescr) SetVersion(v string) {
	o.Version = &v
}

func (o BundleDescr) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Level) {
		toSerialize["level"] = o.Level
	}
	if !isNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !isNil(o.SymbolicName) {
		toSerialize["symbolicName"] = o.SymbolicName
	}
	if !isNil(o.UpdateLocation) {
		toSerialize["updateLocation"] = o.UpdateLocation
	}
	if !isNil(o.Version) {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BundleDescr) UnmarshalJSON(bytes []byte) (err error) {
	varBundleDescr := _BundleDescr{}

	if err = json.Unmarshal(bytes, &varBundleDescr); err == nil {
		*o = BundleDescr(varBundleDescr)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "level")
		delete(additionalProperties, "location")
		delete(additionalProperties, "name")
		delete(additionalProperties, "state")
		delete(additionalProperties, "symbolicName")
		delete(additionalProperties, "updateLocation")
		delete(additionalProperties, "version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBundleDescr struct {
	value *BundleDescr
	isSet bool
}

func (v NullableBundleDescr) Get() *BundleDescr {
	return v.value
}

func (v *NullableBundleDescr) Set(val *BundleDescr) {
	v.value = val
	v.isSet = true
}

func (v NullableBundleDescr) IsSet() bool {
	return v.isSet
}

func (v *NullableBundleDescr) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBundleDescr(val *BundleDescr) *NullableBundleDescr {
	return &NullableBundleDescr{value: val, isSet: true}
}

func (v NullableBundleDescr) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBundleDescr) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


