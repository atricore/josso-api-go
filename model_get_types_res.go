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

// GetTypesRes struct for GetTypesRes
type GetTypesRes struct {
	BasicAuthnMechanism *BasicAuthenticationDTO `json:"basicAuthnMechanism,omitempty"`
	Idpc *IdentityProviderChannelDTO `json:"idpc,omitempty"`
	Spc *InternalSaml2ServiceProviderChannelDTO `json:"spc,omitempty"`
	TotpAuthnSvc *TOTPAuthenticationServiceDTO `json:"totpAuthnSvc,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetTypesRes GetTypesRes

// NewGetTypesRes instantiates a new GetTypesRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTypesRes() *GetTypesRes {
	this := GetTypesRes{}
	return &this
}

// NewGetTypesResWithDefaults instantiates a new GetTypesRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTypesResWithDefaults() *GetTypesRes {
	this := GetTypesRes{}
	return &this
}

// GetBasicAuthnMechanism returns the BasicAuthnMechanism field value if set, zero value otherwise.
func (o *GetTypesRes) GetBasicAuthnMechanism() BasicAuthenticationDTO {
	if o == nil || o.BasicAuthnMechanism == nil {
		var ret BasicAuthenticationDTO
		return ret
	}
	return *o.BasicAuthnMechanism
}

// GetBasicAuthnMechanismOk returns a tuple with the BasicAuthnMechanism field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTypesRes) GetBasicAuthnMechanismOk() (*BasicAuthenticationDTO, bool) {
	if o == nil || o.BasicAuthnMechanism == nil {
		return nil, false
	}
	return o.BasicAuthnMechanism, true
}

// HasBasicAuthnMechanism returns a boolean if a field has been set.
func (o *GetTypesRes) HasBasicAuthnMechanism() bool {
	if o != nil && o.BasicAuthnMechanism != nil {
		return true
	}

	return false
}

// SetBasicAuthnMechanism gets a reference to the given BasicAuthenticationDTO and assigns it to the BasicAuthnMechanism field.
func (o *GetTypesRes) SetBasicAuthnMechanism(v BasicAuthenticationDTO) {
	o.BasicAuthnMechanism = &v
}

// GetIdpc returns the Idpc field value if set, zero value otherwise.
func (o *GetTypesRes) GetIdpc() IdentityProviderChannelDTO {
	if o == nil || o.Idpc == nil {
		var ret IdentityProviderChannelDTO
		return ret
	}
	return *o.Idpc
}

// GetIdpcOk returns a tuple with the Idpc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTypesRes) GetIdpcOk() (*IdentityProviderChannelDTO, bool) {
	if o == nil || o.Idpc == nil {
		return nil, false
	}
	return o.Idpc, true
}

// HasIdpc returns a boolean if a field has been set.
func (o *GetTypesRes) HasIdpc() bool {
	if o != nil && o.Idpc != nil {
		return true
	}

	return false
}

// SetIdpc gets a reference to the given IdentityProviderChannelDTO and assigns it to the Idpc field.
func (o *GetTypesRes) SetIdpc(v IdentityProviderChannelDTO) {
	o.Idpc = &v
}

// GetSpc returns the Spc field value if set, zero value otherwise.
func (o *GetTypesRes) GetSpc() InternalSaml2ServiceProviderChannelDTO {
	if o == nil || o.Spc == nil {
		var ret InternalSaml2ServiceProviderChannelDTO
		return ret
	}
	return *o.Spc
}

// GetSpcOk returns a tuple with the Spc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTypesRes) GetSpcOk() (*InternalSaml2ServiceProviderChannelDTO, bool) {
	if o == nil || o.Spc == nil {
		return nil, false
	}
	return o.Spc, true
}

// HasSpc returns a boolean if a field has been set.
func (o *GetTypesRes) HasSpc() bool {
	if o != nil && o.Spc != nil {
		return true
	}

	return false
}

// SetSpc gets a reference to the given InternalSaml2ServiceProviderChannelDTO and assigns it to the Spc field.
func (o *GetTypesRes) SetSpc(v InternalSaml2ServiceProviderChannelDTO) {
	o.Spc = &v
}

// GetTotpAuthnSvc returns the TotpAuthnSvc field value if set, zero value otherwise.
func (o *GetTypesRes) GetTotpAuthnSvc() TOTPAuthenticationServiceDTO {
	if o == nil || o.TotpAuthnSvc == nil {
		var ret TOTPAuthenticationServiceDTO
		return ret
	}
	return *o.TotpAuthnSvc
}

// GetTotpAuthnSvcOk returns a tuple with the TotpAuthnSvc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTypesRes) GetTotpAuthnSvcOk() (*TOTPAuthenticationServiceDTO, bool) {
	if o == nil || o.TotpAuthnSvc == nil {
		return nil, false
	}
	return o.TotpAuthnSvc, true
}

// HasTotpAuthnSvc returns a boolean if a field has been set.
func (o *GetTypesRes) HasTotpAuthnSvc() bool {
	if o != nil && o.TotpAuthnSvc != nil {
		return true
	}

	return false
}

// SetTotpAuthnSvc gets a reference to the given TOTPAuthenticationServiceDTO and assigns it to the TotpAuthnSvc field.
func (o *GetTypesRes) SetTotpAuthnSvc(v TOTPAuthenticationServiceDTO) {
	o.TotpAuthnSvc = &v
}

func (o GetTypesRes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BasicAuthnMechanism != nil {
		toSerialize["basicAuthnMechanism"] = o.BasicAuthnMechanism
	}
	if o.Idpc != nil {
		toSerialize["idpc"] = o.Idpc
	}
	if o.Spc != nil {
		toSerialize["spc"] = o.Spc
	}
	if o.TotpAuthnSvc != nil {
		toSerialize["totpAuthnSvc"] = o.TotpAuthnSvc
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GetTypesRes) UnmarshalJSON(bytes []byte) (err error) {
	varGetTypesRes := _GetTypesRes{}

	if err = json.Unmarshal(bytes, &varGetTypesRes); err == nil {
		*o = GetTypesRes(varGetTypesRes)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "basicAuthnMechanism")
		delete(additionalProperties, "idpc")
		delete(additionalProperties, "spc")
		delete(additionalProperties, "totpAuthnSvc")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetTypesRes struct {
	value *GetTypesRes
	isSet bool
}

func (v NullableGetTypesRes) Get() *GetTypesRes {
	return v.value
}

func (v *NullableGetTypesRes) Set(val *GetTypesRes) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTypesRes) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTypesRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTypesRes(val *GetTypesRes) *NullableGetTypesRes {
	return &NullableGetTypesRes{value: val, isSet: true}
}

func (v NullableGetTypesRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTypesRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


