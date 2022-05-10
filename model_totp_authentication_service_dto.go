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

// TOTPAuthenticationServiceDTO struct for TOTPAuthenticationServiceDTO
type TOTPAuthenticationServiceDTO struct {
	Crypto *string `json:"crypto,omitempty"`
	DelegatedAuthentications []DelegatedAuthenticationDTO `json:"delegatedAuthentications,omitempty"`
	Description *string `json:"description,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	ElementId *string `json:"elementId,omitempty"`
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	ReturnDigits *int32 `json:"returnDigits,omitempty"`
	SecretCredential *string `json:"secretCredential,omitempty"`
	X *float64 `json:"x,omitempty"`
	Y *float64 `json:"y,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TOTPAuthenticationServiceDTO TOTPAuthenticationServiceDTO

// NewTOTPAuthenticationServiceDTO instantiates a new TOTPAuthenticationServiceDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTOTPAuthenticationServiceDTO() *TOTPAuthenticationServiceDTO {
	this := TOTPAuthenticationServiceDTO{}
	return &this
}

// NewTOTPAuthenticationServiceDTOWithDefaults instantiates a new TOTPAuthenticationServiceDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTOTPAuthenticationServiceDTOWithDefaults() *TOTPAuthenticationServiceDTO {
	this := TOTPAuthenticationServiceDTO{}
	return &this
}

// GetCrypto returns the Crypto field value if set, zero value otherwise.
func (o *TOTPAuthenticationServiceDTO) GetCrypto() string {
	if o == nil || o.Crypto == nil {
		var ret string
		return ret
	}
	return *o.Crypto
}

// GetCryptoOk returns a tuple with the Crypto field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TOTPAuthenticationServiceDTO) GetCryptoOk() (*string, bool) {
	if o == nil || o.Crypto == nil {
		return nil, false
	}
	return o.Crypto, true
}

// HasCrypto returns a boolean if a field has been set.
func (o *TOTPAuthenticationServiceDTO) HasCrypto() bool {
	if o != nil && o.Crypto != nil {
		return true
	}

	return false
}

// SetCrypto gets a reference to the given string and assigns it to the Crypto field.
func (o *TOTPAuthenticationServiceDTO) SetCrypto(v string) {
	o.Crypto = &v
}

// GetDelegatedAuthentications returns the DelegatedAuthentications field value if set, zero value otherwise.
func (o *TOTPAuthenticationServiceDTO) GetDelegatedAuthentications() []DelegatedAuthenticationDTO {
	if o == nil || o.DelegatedAuthentications == nil {
		var ret []DelegatedAuthenticationDTO
		return ret
	}
	return o.DelegatedAuthentications
}

// GetDelegatedAuthenticationsOk returns a tuple with the DelegatedAuthentications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TOTPAuthenticationServiceDTO) GetDelegatedAuthenticationsOk() ([]DelegatedAuthenticationDTO, bool) {
	if o == nil || o.DelegatedAuthentications == nil {
		return nil, false
	}
	return o.DelegatedAuthentications, true
}

// HasDelegatedAuthentications returns a boolean if a field has been set.
func (o *TOTPAuthenticationServiceDTO) HasDelegatedAuthentications() bool {
	if o != nil && o.DelegatedAuthentications != nil {
		return true
	}

	return false
}

// SetDelegatedAuthentications gets a reference to the given []DelegatedAuthenticationDTO and assigns it to the DelegatedAuthentications field.
func (o *TOTPAuthenticationServiceDTO) SetDelegatedAuthentications(v []DelegatedAuthenticationDTO) {
	o.DelegatedAuthentications = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TOTPAuthenticationServiceDTO) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TOTPAuthenticationServiceDTO) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TOTPAuthenticationServiceDTO) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TOTPAuthenticationServiceDTO) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *TOTPAuthenticationServiceDTO) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TOTPAuthenticationServiceDTO) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *TOTPAuthenticationServiceDTO) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *TOTPAuthenticationServiceDTO) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *TOTPAuthenticationServiceDTO) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TOTPAuthenticationServiceDTO) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *TOTPAuthenticationServiceDTO) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *TOTPAuthenticationServiceDTO) SetElementId(v string) {
	o.ElementId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TOTPAuthenticationServiceDTO) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TOTPAuthenticationServiceDTO) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TOTPAuthenticationServiceDTO) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *TOTPAuthenticationServiceDTO) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TOTPAuthenticationServiceDTO) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TOTPAuthenticationServiceDTO) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TOTPAuthenticationServiceDTO) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TOTPAuthenticationServiceDTO) SetName(v string) {
	o.Name = &v
}

// GetReturnDigits returns the ReturnDigits field value if set, zero value otherwise.
func (o *TOTPAuthenticationServiceDTO) GetReturnDigits() int32 {
	if o == nil || o.ReturnDigits == nil {
		var ret int32
		return ret
	}
	return *o.ReturnDigits
}

// GetReturnDigitsOk returns a tuple with the ReturnDigits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TOTPAuthenticationServiceDTO) GetReturnDigitsOk() (*int32, bool) {
	if o == nil || o.ReturnDigits == nil {
		return nil, false
	}
	return o.ReturnDigits, true
}

// HasReturnDigits returns a boolean if a field has been set.
func (o *TOTPAuthenticationServiceDTO) HasReturnDigits() bool {
	if o != nil && o.ReturnDigits != nil {
		return true
	}

	return false
}

// SetReturnDigits gets a reference to the given int32 and assigns it to the ReturnDigits field.
func (o *TOTPAuthenticationServiceDTO) SetReturnDigits(v int32) {
	o.ReturnDigits = &v
}

// GetSecretCredential returns the SecretCredential field value if set, zero value otherwise.
func (o *TOTPAuthenticationServiceDTO) GetSecretCredential() string {
	if o == nil || o.SecretCredential == nil {
		var ret string
		return ret
	}
	return *o.SecretCredential
}

// GetSecretCredentialOk returns a tuple with the SecretCredential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TOTPAuthenticationServiceDTO) GetSecretCredentialOk() (*string, bool) {
	if o == nil || o.SecretCredential == nil {
		return nil, false
	}
	return o.SecretCredential, true
}

// HasSecretCredential returns a boolean if a field has been set.
func (o *TOTPAuthenticationServiceDTO) HasSecretCredential() bool {
	if o != nil && o.SecretCredential != nil {
		return true
	}

	return false
}

// SetSecretCredential gets a reference to the given string and assigns it to the SecretCredential field.
func (o *TOTPAuthenticationServiceDTO) SetSecretCredential(v string) {
	o.SecretCredential = &v
}

// GetX returns the X field value if set, zero value otherwise.
func (o *TOTPAuthenticationServiceDTO) GetX() float64 {
	if o == nil || o.X == nil {
		var ret float64
		return ret
	}
	return *o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TOTPAuthenticationServiceDTO) GetXOk() (*float64, bool) {
	if o == nil || o.X == nil {
		return nil, false
	}
	return o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *TOTPAuthenticationServiceDTO) HasX() bool {
	if o != nil && o.X != nil {
		return true
	}

	return false
}

// SetX gets a reference to the given float64 and assigns it to the X field.
func (o *TOTPAuthenticationServiceDTO) SetX(v float64) {
	o.X = &v
}

// GetY returns the Y field value if set, zero value otherwise.
func (o *TOTPAuthenticationServiceDTO) GetY() float64 {
	if o == nil || o.Y == nil {
		var ret float64
		return ret
	}
	return *o.Y
}

// GetYOk returns a tuple with the Y field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TOTPAuthenticationServiceDTO) GetYOk() (*float64, bool) {
	if o == nil || o.Y == nil {
		return nil, false
	}
	return o.Y, true
}

// HasY returns a boolean if a field has been set.
func (o *TOTPAuthenticationServiceDTO) HasY() bool {
	if o != nil && o.Y != nil {
		return true
	}

	return false
}

// SetY gets a reference to the given float64 and assigns it to the Y field.
func (o *TOTPAuthenticationServiceDTO) SetY(v float64) {
	o.Y = &v
}

func (o TOTPAuthenticationServiceDTO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Crypto != nil {
		toSerialize["crypto"] = o.Crypto
	}
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
	if o.ReturnDigits != nil {
		toSerialize["returnDigits"] = o.ReturnDigits
	}
	if o.SecretCredential != nil {
		toSerialize["secretCredential"] = o.SecretCredential
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

func (o *TOTPAuthenticationServiceDTO) UnmarshalJSON(bytes []byte) (err error) {
	varTOTPAuthenticationServiceDTO := _TOTPAuthenticationServiceDTO{}

	if err = json.Unmarshal(bytes, &varTOTPAuthenticationServiceDTO); err == nil {
		*o = TOTPAuthenticationServiceDTO(varTOTPAuthenticationServiceDTO)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "crypto")
		delete(additionalProperties, "delegatedAuthentications")
		delete(additionalProperties, "description")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "elementId")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "returnDigits")
		delete(additionalProperties, "secretCredential")
		delete(additionalProperties, "x")
		delete(additionalProperties, "y")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTOTPAuthenticationServiceDTO struct {
	value *TOTPAuthenticationServiceDTO
	isSet bool
}

func (v NullableTOTPAuthenticationServiceDTO) Get() *TOTPAuthenticationServiceDTO {
	return v.value
}

func (v *NullableTOTPAuthenticationServiceDTO) Set(val *TOTPAuthenticationServiceDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableTOTPAuthenticationServiceDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableTOTPAuthenticationServiceDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTOTPAuthenticationServiceDTO(val *TOTPAuthenticationServiceDTO) *NullableTOTPAuthenticationServiceDTO {
	return &NullableTOTPAuthenticationServiceDTO{value: val, isSet: true}
}

func (v NullableTOTPAuthenticationServiceDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTOTPAuthenticationServiceDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

