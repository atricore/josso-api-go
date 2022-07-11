# GetIdpFacebooksRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**Idps** | Pointer to [**[]FacebookOpenIDConnectIdentityProviderDTO**](FacebookOpenIDConnectIdentityProviderDTO.md) |  | [optional] 
**ValidationErrors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetIdpFacebooksRes

`func NewGetIdpFacebooksRes() *GetIdpFacebooksRes`

NewGetIdpFacebooksRes instantiates a new GetIdpFacebooksRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIdpFacebooksResWithDefaults

`func NewGetIdpFacebooksResWithDefaults() *GetIdpFacebooksRes`

NewGetIdpFacebooksResWithDefaults instantiates a new GetIdpFacebooksRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *GetIdpFacebooksRes) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetIdpFacebooksRes) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetIdpFacebooksRes) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetIdpFacebooksRes) HasError() bool`

HasError returns a boolean if a field has been set.

### GetIdps

`func (o *GetIdpFacebooksRes) GetIdps() []FacebookOpenIDConnectIdentityProviderDTO`

GetIdps returns the Idps field if non-nil, zero value otherwise.

### GetIdpsOk

`func (o *GetIdpFacebooksRes) GetIdpsOk() (*[]FacebookOpenIDConnectIdentityProviderDTO, bool)`

GetIdpsOk returns a tuple with the Idps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdps

`func (o *GetIdpFacebooksRes) SetIdps(v []FacebookOpenIDConnectIdentityProviderDTO)`

SetIdps sets Idps field to given value.

### HasIdps

`func (o *GetIdpFacebooksRes) HasIdps() bool`

HasIdps returns a boolean if a field has been set.

### GetValidationErrors

`func (o *GetIdpFacebooksRes) GetValidationErrors() []string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *GetIdpFacebooksRes) GetValidationErrorsOk() (*[]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *GetIdpFacebooksRes) SetValidationErrors(v []string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *GetIdpFacebooksRes) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


