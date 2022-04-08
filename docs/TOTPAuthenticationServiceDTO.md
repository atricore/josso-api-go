# TOTPAuthenticationServiceDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Crypto** | Pointer to **string** |  | [optional] 
**DelegatedAuthentications** | Pointer to [**[]DelegatedAuthenticationDTO**](DelegatedAuthenticationDTO.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ReturnDigits** | Pointer to **int32** |  | [optional] 
**SecretCredential** | Pointer to **string** |  | [optional] 
**X** | Pointer to **float64** |  | [optional] 
**Y** | Pointer to **float64** |  | [optional] 

## Methods

### NewTOTPAuthenticationServiceDTO

`func NewTOTPAuthenticationServiceDTO() *TOTPAuthenticationServiceDTO`

NewTOTPAuthenticationServiceDTO instantiates a new TOTPAuthenticationServiceDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTOTPAuthenticationServiceDTOWithDefaults

`func NewTOTPAuthenticationServiceDTOWithDefaults() *TOTPAuthenticationServiceDTO`

NewTOTPAuthenticationServiceDTOWithDefaults instantiates a new TOTPAuthenticationServiceDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCrypto

`func (o *TOTPAuthenticationServiceDTO) GetCrypto() string`

GetCrypto returns the Crypto field if non-nil, zero value otherwise.

### GetCryptoOk

`func (o *TOTPAuthenticationServiceDTO) GetCryptoOk() (*string, bool)`

GetCryptoOk returns a tuple with the Crypto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrypto

`func (o *TOTPAuthenticationServiceDTO) SetCrypto(v string)`

SetCrypto sets Crypto field to given value.

### HasCrypto

`func (o *TOTPAuthenticationServiceDTO) HasCrypto() bool`

HasCrypto returns a boolean if a field has been set.

### GetDelegatedAuthentications

`func (o *TOTPAuthenticationServiceDTO) GetDelegatedAuthentications() []DelegatedAuthenticationDTO`

GetDelegatedAuthentications returns the DelegatedAuthentications field if non-nil, zero value otherwise.

### GetDelegatedAuthenticationsOk

`func (o *TOTPAuthenticationServiceDTO) GetDelegatedAuthenticationsOk() (*[]DelegatedAuthenticationDTO, bool)`

GetDelegatedAuthenticationsOk returns a tuple with the DelegatedAuthentications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatedAuthentications

`func (o *TOTPAuthenticationServiceDTO) SetDelegatedAuthentications(v []DelegatedAuthenticationDTO)`

SetDelegatedAuthentications sets DelegatedAuthentications field to given value.

### HasDelegatedAuthentications

`func (o *TOTPAuthenticationServiceDTO) HasDelegatedAuthentications() bool`

HasDelegatedAuthentications returns a boolean if a field has been set.

### GetDescription

`func (o *TOTPAuthenticationServiceDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TOTPAuthenticationServiceDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TOTPAuthenticationServiceDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TOTPAuthenticationServiceDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *TOTPAuthenticationServiceDTO) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *TOTPAuthenticationServiceDTO) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *TOTPAuthenticationServiceDTO) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *TOTPAuthenticationServiceDTO) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetElementId

`func (o *TOTPAuthenticationServiceDTO) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *TOTPAuthenticationServiceDTO) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *TOTPAuthenticationServiceDTO) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *TOTPAuthenticationServiceDTO) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetId

`func (o *TOTPAuthenticationServiceDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TOTPAuthenticationServiceDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TOTPAuthenticationServiceDTO) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *TOTPAuthenticationServiceDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TOTPAuthenticationServiceDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TOTPAuthenticationServiceDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TOTPAuthenticationServiceDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TOTPAuthenticationServiceDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReturnDigits

`func (o *TOTPAuthenticationServiceDTO) GetReturnDigits() int32`

GetReturnDigits returns the ReturnDigits field if non-nil, zero value otherwise.

### GetReturnDigitsOk

`func (o *TOTPAuthenticationServiceDTO) GetReturnDigitsOk() (*int32, bool)`

GetReturnDigitsOk returns a tuple with the ReturnDigits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnDigits

`func (o *TOTPAuthenticationServiceDTO) SetReturnDigits(v int32)`

SetReturnDigits sets ReturnDigits field to given value.

### HasReturnDigits

`func (o *TOTPAuthenticationServiceDTO) HasReturnDigits() bool`

HasReturnDigits returns a boolean if a field has been set.

### GetSecretCredential

`func (o *TOTPAuthenticationServiceDTO) GetSecretCredential() string`

GetSecretCredential returns the SecretCredential field if non-nil, zero value otherwise.

### GetSecretCredentialOk

`func (o *TOTPAuthenticationServiceDTO) GetSecretCredentialOk() (*string, bool)`

GetSecretCredentialOk returns a tuple with the SecretCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretCredential

`func (o *TOTPAuthenticationServiceDTO) SetSecretCredential(v string)`

SetSecretCredential sets SecretCredential field to given value.

### HasSecretCredential

`func (o *TOTPAuthenticationServiceDTO) HasSecretCredential() bool`

HasSecretCredential returns a boolean if a field has been set.

### GetX

`func (o *TOTPAuthenticationServiceDTO) GetX() float64`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *TOTPAuthenticationServiceDTO) GetXOk() (*float64, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *TOTPAuthenticationServiceDTO) SetX(v float64)`

SetX sets X field to given value.

### HasX

`func (o *TOTPAuthenticationServiceDTO) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *TOTPAuthenticationServiceDTO) GetY() float64`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *TOTPAuthenticationServiceDTO) GetYOk() (*float64, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *TOTPAuthenticationServiceDTO) SetY(v float64)`

SetY sets Y field to given value.

### HasY

`func (o *TOTPAuthenticationServiceDTO) HasY() bool`

HasY returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


