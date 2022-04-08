# BasicAuthenticationDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DelegatedAuthentication** | Pointer to [**DelegatedAuthenticationDTO**](DelegatedAuthenticationDTO.md) |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**HashAlgorithm** | Pointer to **string** |  | [optional] 
**HashEncoding** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**IgnorePasswordCase** | Pointer to **bool** |  | [optional] 
**IgnoreUsernameCase** | Pointer to **bool** |  | [optional] 
**ImpersonateUserPolicy** | Pointer to [**ImpersonateUserPolicyDTO**](ImpersonateUserPolicyDTO.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**SaltLength** | Pointer to **int32** |  | [optional] 
**SaltPrefix** | Pointer to **string** |  | [optional] 
**SaltSuffix** | Pointer to **string** |  | [optional] 
**SimpleAuthnSaml2AuthnCtxClass** | Pointer to **string** |  | [optional] 
**X** | Pointer to **float64** |  | [optional] 
**Y** | Pointer to **float64** |  | [optional] 

## Methods

### NewBasicAuthenticationDTO

`func NewBasicAuthenticationDTO() *BasicAuthenticationDTO`

NewBasicAuthenticationDTO instantiates a new BasicAuthenticationDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBasicAuthenticationDTOWithDefaults

`func NewBasicAuthenticationDTOWithDefaults() *BasicAuthenticationDTO`

NewBasicAuthenticationDTOWithDefaults instantiates a new BasicAuthenticationDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelegatedAuthentication

`func (o *BasicAuthenticationDTO) GetDelegatedAuthentication() DelegatedAuthenticationDTO`

GetDelegatedAuthentication returns the DelegatedAuthentication field if non-nil, zero value otherwise.

### GetDelegatedAuthenticationOk

`func (o *BasicAuthenticationDTO) GetDelegatedAuthenticationOk() (*DelegatedAuthenticationDTO, bool)`

GetDelegatedAuthenticationOk returns a tuple with the DelegatedAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatedAuthentication

`func (o *BasicAuthenticationDTO) SetDelegatedAuthentication(v DelegatedAuthenticationDTO)`

SetDelegatedAuthentication sets DelegatedAuthentication field to given value.

### HasDelegatedAuthentication

`func (o *BasicAuthenticationDTO) HasDelegatedAuthentication() bool`

HasDelegatedAuthentication returns a boolean if a field has been set.

### GetDisplayName

`func (o *BasicAuthenticationDTO) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *BasicAuthenticationDTO) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *BasicAuthenticationDTO) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *BasicAuthenticationDTO) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetElementId

`func (o *BasicAuthenticationDTO) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *BasicAuthenticationDTO) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *BasicAuthenticationDTO) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *BasicAuthenticationDTO) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetEnabled

`func (o *BasicAuthenticationDTO) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BasicAuthenticationDTO) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BasicAuthenticationDTO) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *BasicAuthenticationDTO) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetHashAlgorithm

`func (o *BasicAuthenticationDTO) GetHashAlgorithm() string`

GetHashAlgorithm returns the HashAlgorithm field if non-nil, zero value otherwise.

### GetHashAlgorithmOk

`func (o *BasicAuthenticationDTO) GetHashAlgorithmOk() (*string, bool)`

GetHashAlgorithmOk returns a tuple with the HashAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashAlgorithm

`func (o *BasicAuthenticationDTO) SetHashAlgorithm(v string)`

SetHashAlgorithm sets HashAlgorithm field to given value.

### HasHashAlgorithm

`func (o *BasicAuthenticationDTO) HasHashAlgorithm() bool`

HasHashAlgorithm returns a boolean if a field has been set.

### GetHashEncoding

`func (o *BasicAuthenticationDTO) GetHashEncoding() string`

GetHashEncoding returns the HashEncoding field if non-nil, zero value otherwise.

### GetHashEncodingOk

`func (o *BasicAuthenticationDTO) GetHashEncodingOk() (*string, bool)`

GetHashEncodingOk returns a tuple with the HashEncoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashEncoding

`func (o *BasicAuthenticationDTO) SetHashEncoding(v string)`

SetHashEncoding sets HashEncoding field to given value.

### HasHashEncoding

`func (o *BasicAuthenticationDTO) HasHashEncoding() bool`

HasHashEncoding returns a boolean if a field has been set.

### GetId

`func (o *BasicAuthenticationDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BasicAuthenticationDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BasicAuthenticationDTO) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *BasicAuthenticationDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIgnorePasswordCase

`func (o *BasicAuthenticationDTO) GetIgnorePasswordCase() bool`

GetIgnorePasswordCase returns the IgnorePasswordCase field if non-nil, zero value otherwise.

### GetIgnorePasswordCaseOk

`func (o *BasicAuthenticationDTO) GetIgnorePasswordCaseOk() (*bool, bool)`

GetIgnorePasswordCaseOk returns a tuple with the IgnorePasswordCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnorePasswordCase

`func (o *BasicAuthenticationDTO) SetIgnorePasswordCase(v bool)`

SetIgnorePasswordCase sets IgnorePasswordCase field to given value.

### HasIgnorePasswordCase

`func (o *BasicAuthenticationDTO) HasIgnorePasswordCase() bool`

HasIgnorePasswordCase returns a boolean if a field has been set.

### GetIgnoreUsernameCase

`func (o *BasicAuthenticationDTO) GetIgnoreUsernameCase() bool`

GetIgnoreUsernameCase returns the IgnoreUsernameCase field if non-nil, zero value otherwise.

### GetIgnoreUsernameCaseOk

`func (o *BasicAuthenticationDTO) GetIgnoreUsernameCaseOk() (*bool, bool)`

GetIgnoreUsernameCaseOk returns a tuple with the IgnoreUsernameCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreUsernameCase

`func (o *BasicAuthenticationDTO) SetIgnoreUsernameCase(v bool)`

SetIgnoreUsernameCase sets IgnoreUsernameCase field to given value.

### HasIgnoreUsernameCase

`func (o *BasicAuthenticationDTO) HasIgnoreUsernameCase() bool`

HasIgnoreUsernameCase returns a boolean if a field has been set.

### GetImpersonateUserPolicy

`func (o *BasicAuthenticationDTO) GetImpersonateUserPolicy() ImpersonateUserPolicyDTO`

GetImpersonateUserPolicy returns the ImpersonateUserPolicy field if non-nil, zero value otherwise.

### GetImpersonateUserPolicyOk

`func (o *BasicAuthenticationDTO) GetImpersonateUserPolicyOk() (*ImpersonateUserPolicyDTO, bool)`

GetImpersonateUserPolicyOk returns a tuple with the ImpersonateUserPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpersonateUserPolicy

`func (o *BasicAuthenticationDTO) SetImpersonateUserPolicy(v ImpersonateUserPolicyDTO)`

SetImpersonateUserPolicy sets ImpersonateUserPolicy field to given value.

### HasImpersonateUserPolicy

`func (o *BasicAuthenticationDTO) HasImpersonateUserPolicy() bool`

HasImpersonateUserPolicy returns a boolean if a field has been set.

### GetName

`func (o *BasicAuthenticationDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BasicAuthenticationDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BasicAuthenticationDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BasicAuthenticationDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPriority

`func (o *BasicAuthenticationDTO) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *BasicAuthenticationDTO) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *BasicAuthenticationDTO) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *BasicAuthenticationDTO) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetSaltLength

`func (o *BasicAuthenticationDTO) GetSaltLength() int32`

GetSaltLength returns the SaltLength field if non-nil, zero value otherwise.

### GetSaltLengthOk

`func (o *BasicAuthenticationDTO) GetSaltLengthOk() (*int32, bool)`

GetSaltLengthOk returns a tuple with the SaltLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaltLength

`func (o *BasicAuthenticationDTO) SetSaltLength(v int32)`

SetSaltLength sets SaltLength field to given value.

### HasSaltLength

`func (o *BasicAuthenticationDTO) HasSaltLength() bool`

HasSaltLength returns a boolean if a field has been set.

### GetSaltPrefix

`func (o *BasicAuthenticationDTO) GetSaltPrefix() string`

GetSaltPrefix returns the SaltPrefix field if non-nil, zero value otherwise.

### GetSaltPrefixOk

`func (o *BasicAuthenticationDTO) GetSaltPrefixOk() (*string, bool)`

GetSaltPrefixOk returns a tuple with the SaltPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaltPrefix

`func (o *BasicAuthenticationDTO) SetSaltPrefix(v string)`

SetSaltPrefix sets SaltPrefix field to given value.

### HasSaltPrefix

`func (o *BasicAuthenticationDTO) HasSaltPrefix() bool`

HasSaltPrefix returns a boolean if a field has been set.

### GetSaltSuffix

`func (o *BasicAuthenticationDTO) GetSaltSuffix() string`

GetSaltSuffix returns the SaltSuffix field if non-nil, zero value otherwise.

### GetSaltSuffixOk

`func (o *BasicAuthenticationDTO) GetSaltSuffixOk() (*string, bool)`

GetSaltSuffixOk returns a tuple with the SaltSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaltSuffix

`func (o *BasicAuthenticationDTO) SetSaltSuffix(v string)`

SetSaltSuffix sets SaltSuffix field to given value.

### HasSaltSuffix

`func (o *BasicAuthenticationDTO) HasSaltSuffix() bool`

HasSaltSuffix returns a boolean if a field has been set.

### GetSimpleAuthnSaml2AuthnCtxClass

`func (o *BasicAuthenticationDTO) GetSimpleAuthnSaml2AuthnCtxClass() string`

GetSimpleAuthnSaml2AuthnCtxClass returns the SimpleAuthnSaml2AuthnCtxClass field if non-nil, zero value otherwise.

### GetSimpleAuthnSaml2AuthnCtxClassOk

`func (o *BasicAuthenticationDTO) GetSimpleAuthnSaml2AuthnCtxClassOk() (*string, bool)`

GetSimpleAuthnSaml2AuthnCtxClassOk returns a tuple with the SimpleAuthnSaml2AuthnCtxClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimpleAuthnSaml2AuthnCtxClass

`func (o *BasicAuthenticationDTO) SetSimpleAuthnSaml2AuthnCtxClass(v string)`

SetSimpleAuthnSaml2AuthnCtxClass sets SimpleAuthnSaml2AuthnCtxClass field to given value.

### HasSimpleAuthnSaml2AuthnCtxClass

`func (o *BasicAuthenticationDTO) HasSimpleAuthnSaml2AuthnCtxClass() bool`

HasSimpleAuthnSaml2AuthnCtxClass returns a boolean if a field has been set.

### GetX

`func (o *BasicAuthenticationDTO) GetX() float64`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *BasicAuthenticationDTO) GetXOk() (*float64, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *BasicAuthenticationDTO) SetX(v float64)`

SetX sets X field to given value.

### HasX

`func (o *BasicAuthenticationDTO) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *BasicAuthenticationDTO) GetY() float64`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *BasicAuthenticationDTO) GetYOk() (*float64, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *BasicAuthenticationDTO) SetY(v float64)`

SetY sets Y field to given value.

### HasY

`func (o *BasicAuthenticationDTO) HasY() bool`

HasY returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


