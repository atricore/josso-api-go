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

// LocationDTO struct for LocationDTO
type LocationDTO struct {
	Context *string `json:"context,omitempty"`
	ElementId *string `json:"elementId,omitempty"`
	Host *string `json:"host,omitempty"`
	Id *int64 `json:"id,omitempty"`
	LocationAsString *string `json:"locationAsString,omitempty"`
	Port *int32 `json:"port,omitempty"`
	Protocol *string `json:"protocol,omitempty"`
	Uri *string `json:"uri,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LocationDTO LocationDTO

// NewLocationDTO instantiates a new LocationDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationDTO() *LocationDTO {
	this := LocationDTO{}
	return &this
}

// NewLocationDTOWithDefaults instantiates a new LocationDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationDTOWithDefaults() *LocationDTO {
	this := LocationDTO{}
	return &this
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *LocationDTO) GetContext() string {
	if o == nil || o.Context == nil {
		var ret string
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationDTO) GetContextOk() (*string, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *LocationDTO) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given string and assigns it to the Context field.
func (o *LocationDTO) SetContext(v string) {
	o.Context = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *LocationDTO) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationDTO) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *LocationDTO) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *LocationDTO) SetElementId(v string) {
	o.ElementId = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *LocationDTO) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationDTO) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *LocationDTO) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *LocationDTO) SetHost(v string) {
	o.Host = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LocationDTO) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationDTO) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LocationDTO) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *LocationDTO) SetId(v int64) {
	o.Id = &v
}

// GetLocationAsString returns the LocationAsString field value if set, zero value otherwise.
func (o *LocationDTO) GetLocationAsString() string {
	if o == nil || o.LocationAsString == nil {
		var ret string
		return ret
	}
	return *o.LocationAsString
}

// GetLocationAsStringOk returns a tuple with the LocationAsString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationDTO) GetLocationAsStringOk() (*string, bool) {
	if o == nil || o.LocationAsString == nil {
		return nil, false
	}
	return o.LocationAsString, true
}

// HasLocationAsString returns a boolean if a field has been set.
func (o *LocationDTO) HasLocationAsString() bool {
	if o != nil && o.LocationAsString != nil {
		return true
	}

	return false
}

// SetLocationAsString gets a reference to the given string and assigns it to the LocationAsString field.
func (o *LocationDTO) SetLocationAsString(v string) {
	o.LocationAsString = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *LocationDTO) GetPort() int32 {
	if o == nil || o.Port == nil {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationDTO) GetPortOk() (*int32, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *LocationDTO) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *LocationDTO) SetPort(v int32) {
	o.Port = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *LocationDTO) GetProtocol() string {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationDTO) GetProtocolOk() (*string, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *LocationDTO) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *LocationDTO) SetProtocol(v string) {
	o.Protocol = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *LocationDTO) GetUri() string {
	if o == nil || o.Uri == nil {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationDTO) GetUriOk() (*string, bool) {
	if o == nil || o.Uri == nil {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *LocationDTO) HasUri() bool {
	if o != nil && o.Uri != nil {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *LocationDTO) SetUri(v string) {
	o.Uri = &v
}

func (o LocationDTO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}
	if o.ElementId != nil {
		toSerialize["elementId"] = o.ElementId
	}
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LocationAsString != nil {
		toSerialize["locationAsString"] = o.LocationAsString
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if o.Protocol != nil {
		toSerialize["protocol"] = o.Protocol
	}
	if o.Uri != nil {
		toSerialize["uri"] = o.Uri
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LocationDTO) UnmarshalJSON(bytes []byte) (err error) {
	varLocationDTO := _LocationDTO{}

	if err = json.Unmarshal(bytes, &varLocationDTO); err == nil {
		*o = LocationDTO(varLocationDTO)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "context")
		delete(additionalProperties, "elementId")
		delete(additionalProperties, "host")
		delete(additionalProperties, "id")
		delete(additionalProperties, "locationAsString")
		delete(additionalProperties, "port")
		delete(additionalProperties, "protocol")
		delete(additionalProperties, "uri")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLocationDTO struct {
	value *LocationDTO
	isSet bool
}

func (v NullableLocationDTO) Get() *LocationDTO {
	return v.value
}

func (v *NullableLocationDTO) Set(val *LocationDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationDTO(val *LocationDTO) *NullableLocationDTO {
	return &NullableLocationDTO{value: val, isSet: true}
}

func (v NullableLocationDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


