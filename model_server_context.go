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

// ServerContext struct for ServerContext
type ServerContext struct {
	AccessToken *string `json:"accessToken,omitempty"`
	RefreshToken *string `json:"refreshToken,omitempty"`
	ServerId *string `json:"serverId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServerContext ServerContext

// NewServerContext instantiates a new ServerContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerContext() *ServerContext {
	this := ServerContext{}
	return &this
}

// NewServerContextWithDefaults instantiates a new ServerContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerContextWithDefaults() *ServerContext {
	this := ServerContext{}
	return &this
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise.
func (o *ServerContext) GetAccessToken() string {
	if o == nil || o.AccessToken == nil {
		var ret string
		return ret
	}
	return *o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerContext) GetAccessTokenOk() (*string, bool) {
	if o == nil || o.AccessToken == nil {
		return nil, false
	}
	return o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *ServerContext) HasAccessToken() bool {
	if o != nil && o.AccessToken != nil {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given string and assigns it to the AccessToken field.
func (o *ServerContext) SetAccessToken(v string) {
	o.AccessToken = &v
}

// GetRefreshToken returns the RefreshToken field value if set, zero value otherwise.
func (o *ServerContext) GetRefreshToken() string {
	if o == nil || o.RefreshToken == nil {
		var ret string
		return ret
	}
	return *o.RefreshToken
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerContext) GetRefreshTokenOk() (*string, bool) {
	if o == nil || o.RefreshToken == nil {
		return nil, false
	}
	return o.RefreshToken, true
}

// HasRefreshToken returns a boolean if a field has been set.
func (o *ServerContext) HasRefreshToken() bool {
	if o != nil && o.RefreshToken != nil {
		return true
	}

	return false
}

// SetRefreshToken gets a reference to the given string and assigns it to the RefreshToken field.
func (o *ServerContext) SetRefreshToken(v string) {
	o.RefreshToken = &v
}

// GetServerId returns the ServerId field value if set, zero value otherwise.
func (o *ServerContext) GetServerId() string {
	if o == nil || o.ServerId == nil {
		var ret string
		return ret
	}
	return *o.ServerId
}

// GetServerIdOk returns a tuple with the ServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerContext) GetServerIdOk() (*string, bool) {
	if o == nil || o.ServerId == nil {
		return nil, false
	}
	return o.ServerId, true
}

// HasServerId returns a boolean if a field has been set.
func (o *ServerContext) HasServerId() bool {
	if o != nil && o.ServerId != nil {
		return true
	}

	return false
}

// SetServerId gets a reference to the given string and assigns it to the ServerId field.
func (o *ServerContext) SetServerId(v string) {
	o.ServerId = &v
}

func (o ServerContext) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessToken != nil {
		toSerialize["accessToken"] = o.AccessToken
	}
	if o.RefreshToken != nil {
		toSerialize["refreshToken"] = o.RefreshToken
	}
	if o.ServerId != nil {
		toSerialize["serverId"] = o.ServerId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ServerContext) UnmarshalJSON(bytes []byte) (err error) {
	varServerContext := _ServerContext{}

	if err = json.Unmarshal(bytes, &varServerContext); err == nil {
		*o = ServerContext(varServerContext)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "accessToken")
		delete(additionalProperties, "refreshToken")
		delete(additionalProperties, "serverId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServerContext struct {
	value *ServerContext
	isSet bool
}

func (v NullableServerContext) Get() *ServerContext {
	return v.value
}

func (v *NullableServerContext) Set(val *ServerContext) {
	v.value = val
	v.isSet = true
}

func (v NullableServerContext) IsSet() bool {
	return v.isSet
}

func (v *NullableServerContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerContext(val *ServerContext) *NullableServerContext {
	return &NullableServerContext{value: val, isSet: true}
}

func (v NullableServerContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


