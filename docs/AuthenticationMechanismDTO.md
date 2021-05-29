# AuthenticationMechanismDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DelegatedAuthentication** | Pointer to [**DelegatedAuthenticationDTO**](DelegatedAuthenticationDTO.md) |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 

## Methods

### NewAuthenticationMechanismDTO

`func NewAuthenticationMechanismDTO() *AuthenticationMechanismDTO`

NewAuthenticationMechanismDTO instantiates a new AuthenticationMechanismDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationMechanismDTOWithDefaults

`func NewAuthenticationMechanismDTOWithDefaults() *AuthenticationMechanismDTO`

NewAuthenticationMechanismDTOWithDefaults instantiates a new AuthenticationMechanismDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelegatedAuthentication

`func (o *AuthenticationMechanismDTO) GetDelegatedAuthentication() DelegatedAuthenticationDTO`

GetDelegatedAuthentication returns the DelegatedAuthentication field if non-nil, zero value otherwise.

### GetDelegatedAuthenticationOk

`func (o *AuthenticationMechanismDTO) GetDelegatedAuthenticationOk() (*DelegatedAuthenticationDTO, bool)`

GetDelegatedAuthenticationOk returns a tuple with the DelegatedAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatedAuthentication

`func (o *AuthenticationMechanismDTO) SetDelegatedAuthentication(v DelegatedAuthenticationDTO)`

SetDelegatedAuthentication sets DelegatedAuthentication field to given value.

### HasDelegatedAuthentication

`func (o *AuthenticationMechanismDTO) HasDelegatedAuthentication() bool`

HasDelegatedAuthentication returns a boolean if a field has been set.

### GetDisplayName

`func (o *AuthenticationMechanismDTO) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AuthenticationMechanismDTO) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AuthenticationMechanismDTO) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AuthenticationMechanismDTO) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetElementId

`func (o *AuthenticationMechanismDTO) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *AuthenticationMechanismDTO) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *AuthenticationMechanismDTO) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *AuthenticationMechanismDTO) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetId

`func (o *AuthenticationMechanismDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthenticationMechanismDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthenticationMechanismDTO) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AuthenticationMechanismDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AuthenticationMechanismDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthenticationMechanismDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthenticationMechanismDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuthenticationMechanismDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPriority

`func (o *AuthenticationMechanismDTO) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *AuthenticationMechanismDTO) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *AuthenticationMechanismDTO) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *AuthenticationMechanismDTO) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


