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
)

// StoreOidcRpRes struct for StoreOidcRpRes
type StoreOidcRpRes struct {
	Error *string `json:"error,omitempty"`
	OidcRp *ExternalOpenIDConnectRelayingPartyDTO `json:"oidcRp,omitempty"`
	ValidationErrors *[]string `json:"validationErrors,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StoreOidcRpRes StoreOidcRpRes

// NewStoreOidcRpRes instantiates a new StoreOidcRpRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoreOidcRpRes() *StoreOidcRpRes {
	this := StoreOidcRpRes{}
	return &this
}

// NewStoreOidcRpResWithDefaults instantiates a new StoreOidcRpRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoreOidcRpResWithDefaults() *StoreOidcRpRes {
	this := StoreOidcRpRes{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *StoreOidcRpRes) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreOidcRpRes) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *StoreOidcRpRes) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *StoreOidcRpRes) SetError(v string) {
	o.Error = &v
}

// GetOidcRp returns the OidcRp field value if set, zero value otherwise.
func (o *StoreOidcRpRes) GetOidcRp() ExternalOpenIDConnectRelayingPartyDTO {
	if o == nil || o.OidcRp == nil {
		var ret ExternalOpenIDConnectRelayingPartyDTO
		return ret
	}
	return *o.OidcRp
}

// GetOidcRpOk returns a tuple with the OidcRp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreOidcRpRes) GetOidcRpOk() (*ExternalOpenIDConnectRelayingPartyDTO, bool) {
	if o == nil || o.OidcRp == nil {
		return nil, false
	}
	return o.OidcRp, true
}

// HasOidcRp returns a boolean if a field has been set.
func (o *StoreOidcRpRes) HasOidcRp() bool {
	if o != nil && o.OidcRp != nil {
		return true
	}

	return false
}

// SetOidcRp gets a reference to the given ExternalOpenIDConnectRelayingPartyDTO and assigns it to the OidcRp field.
func (o *StoreOidcRpRes) SetOidcRp(v ExternalOpenIDConnectRelayingPartyDTO) {
	o.OidcRp = &v
}

// GetValidationErrors returns the ValidationErrors field value if set, zero value otherwise.
func (o *StoreOidcRpRes) GetValidationErrors() []string {
	if o == nil || o.ValidationErrors == nil {
		var ret []string
		return ret
	}
	return *o.ValidationErrors
}

// GetValidationErrorsOk returns a tuple with the ValidationErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreOidcRpRes) GetValidationErrorsOk() (*[]string, bool) {
	if o == nil || o.ValidationErrors == nil {
		return nil, false
	}
	return o.ValidationErrors, true
}

// HasValidationErrors returns a boolean if a field has been set.
func (o *StoreOidcRpRes) HasValidationErrors() bool {
	if o != nil && o.ValidationErrors != nil {
		return true
	}

	return false
}

// SetValidationErrors gets a reference to the given []string and assigns it to the ValidationErrors field.
func (o *StoreOidcRpRes) SetValidationErrors(v []string) {
	o.ValidationErrors = &v
}

func (o StoreOidcRpRes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.OidcRp != nil {
		toSerialize["oidcRp"] = o.OidcRp
	}
	if o.ValidationErrors != nil {
		toSerialize["validationErrors"] = o.ValidationErrors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StoreOidcRpRes) UnmarshalJSON(bytes []byte) (err error) {
	varStoreOidcRpRes := _StoreOidcRpRes{}

	if err = json.Unmarshal(bytes, &varStoreOidcRpRes); err == nil {
		*o = StoreOidcRpRes(varStoreOidcRpRes)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "error")
		delete(additionalProperties, "oidcRp")
		delete(additionalProperties, "validationErrors")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStoreOidcRpRes struct {
	value *StoreOidcRpRes
	isSet bool
}

func (v NullableStoreOidcRpRes) Get() *StoreOidcRpRes {
	return v.value
}

func (v *NullableStoreOidcRpRes) Set(val *StoreOidcRpRes) {
	v.value = val
	v.isSet = true
}

func (v NullableStoreOidcRpRes) IsSet() bool {
	return v.isSet
}

func (v *NullableStoreOidcRpRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoreOidcRpRes(val *StoreOidcRpRes) *NullableStoreOidcRpRes {
	return &NullableStoreOidcRpRes{value: val, isSet: true}
}

func (v NullableStoreOidcRpRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoreOidcRpRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


