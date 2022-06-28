# OAuth2PreAuthenticationServiceDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthnService** | Pointer to **string** |  | [optional] 
**DelegatedAuthentications** | Pointer to [**[]DelegatedAuthenticationDTO**](DelegatedAuthenticationDTO.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**ExternalAuth** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**RememberMe** | Pointer to **bool** |  | [optional] 
**X** | Pointer to **float64** |  | [optional] 
**Y** | Pointer to **float64** |  | [optional] 

## Methods

### NewOAuth2PreAuthenticationServiceDTO

`func NewOAuth2PreAuthenticationServiceDTO() *OAuth2PreAuthenticationServiceDTO`

NewOAuth2PreAuthenticationServiceDTO instantiates a new OAuth2PreAuthenticationServiceDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2PreAuthenticationServiceDTOWithDefaults

`func NewOAuth2PreAuthenticationServiceDTOWithDefaults() *OAuth2PreAuthenticationServiceDTO`

NewOAuth2PreAuthenticationServiceDTOWithDefaults instantiates a new OAuth2PreAuthenticationServiceDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthnService

`func (o *OAuth2PreAuthenticationServiceDTO) GetAuthnService() string`

GetAuthnService returns the AuthnService field if non-nil, zero value otherwise.

### GetAuthnServiceOk

`func (o *OAuth2PreAuthenticationServiceDTO) GetAuthnServiceOk() (*string, bool)`

GetAuthnServiceOk returns a tuple with the AuthnService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthnService

`func (o *OAuth2PreAuthenticationServiceDTO) SetAuthnService(v string)`

SetAuthnService sets AuthnService field to given value.

### HasAuthnService

`func (o *OAuth2PreAuthenticationServiceDTO) HasAuthnService() bool`

HasAuthnService returns a boolean if a field has been set.

### GetDelegatedAuthentications

`func (o *OAuth2PreAuthenticationServiceDTO) GetDelegatedAuthentications() []DelegatedAuthenticationDTO`

GetDelegatedAuthentications returns the DelegatedAuthentications field if non-nil, zero value otherwise.

### GetDelegatedAuthenticationsOk

`func (o *OAuth2PreAuthenticationServiceDTO) GetDelegatedAuthenticationsOk() (*[]DelegatedAuthenticationDTO, bool)`

GetDelegatedAuthenticationsOk returns a tuple with the DelegatedAuthentications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatedAuthentications

`func (o *OAuth2PreAuthenticationServiceDTO) SetDelegatedAuthentications(v []DelegatedAuthenticationDTO)`

SetDelegatedAuthentications sets DelegatedAuthentications field to given value.

### HasDelegatedAuthentications

`func (o *OAuth2PreAuthenticationServiceDTO) HasDelegatedAuthentications() bool`

HasDelegatedAuthentications returns a boolean if a field has been set.

### GetDescription

`func (o *OAuth2PreAuthenticationServiceDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OAuth2PreAuthenticationServiceDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OAuth2PreAuthenticationServiceDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OAuth2PreAuthenticationServiceDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *OAuth2PreAuthenticationServiceDTO) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *OAuth2PreAuthenticationServiceDTO) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *OAuth2PreAuthenticationServiceDTO) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *OAuth2PreAuthenticationServiceDTO) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetElementId

`func (o *OAuth2PreAuthenticationServiceDTO) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *OAuth2PreAuthenticationServiceDTO) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *OAuth2PreAuthenticationServiceDTO) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *OAuth2PreAuthenticationServiceDTO) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetExternalAuth

`func (o *OAuth2PreAuthenticationServiceDTO) GetExternalAuth() bool`

GetExternalAuth returns the ExternalAuth field if non-nil, zero value otherwise.

### GetExternalAuthOk

`func (o *OAuth2PreAuthenticationServiceDTO) GetExternalAuthOk() (*bool, bool)`

GetExternalAuthOk returns a tuple with the ExternalAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalAuth

`func (o *OAuth2PreAuthenticationServiceDTO) SetExternalAuth(v bool)`

SetExternalAuth sets ExternalAuth field to given value.

### HasExternalAuth

`func (o *OAuth2PreAuthenticationServiceDTO) HasExternalAuth() bool`

HasExternalAuth returns a boolean if a field has been set.

### GetId

`func (o *OAuth2PreAuthenticationServiceDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OAuth2PreAuthenticationServiceDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OAuth2PreAuthenticationServiceDTO) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *OAuth2PreAuthenticationServiceDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *OAuth2PreAuthenticationServiceDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OAuth2PreAuthenticationServiceDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OAuth2PreAuthenticationServiceDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OAuth2PreAuthenticationServiceDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRememberMe

`func (o *OAuth2PreAuthenticationServiceDTO) GetRememberMe() bool`

GetRememberMe returns the RememberMe field if non-nil, zero value otherwise.

### GetRememberMeOk

`func (o *OAuth2PreAuthenticationServiceDTO) GetRememberMeOk() (*bool, bool)`

GetRememberMeOk returns a tuple with the RememberMe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRememberMe

`func (o *OAuth2PreAuthenticationServiceDTO) SetRememberMe(v bool)`

SetRememberMe sets RememberMe field to given value.

### HasRememberMe

`func (o *OAuth2PreAuthenticationServiceDTO) HasRememberMe() bool`

HasRememberMe returns a boolean if a field has been set.

### GetX

`func (o *OAuth2PreAuthenticationServiceDTO) GetX() float64`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *OAuth2PreAuthenticationServiceDTO) GetXOk() (*float64, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *OAuth2PreAuthenticationServiceDTO) SetX(v float64)`

SetX sets X field to given value.

### HasX

`func (o *OAuth2PreAuthenticationServiceDTO) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *OAuth2PreAuthenticationServiceDTO) GetY() float64`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *OAuth2PreAuthenticationServiceDTO) GetYOk() (*float64, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *OAuth2PreAuthenticationServiceDTO) SetY(v float64)`

SetY sets Y field to given value.

### HasY

`func (o *OAuth2PreAuthenticationServiceDTO) HasY() bool`

HasY returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


