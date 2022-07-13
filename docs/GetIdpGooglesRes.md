# GetIdpGooglesRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**Idps** | Pointer to [**[]GoogleOpenIDConnectIdentityProviderDTO**](GoogleOpenIDConnectIdentityProviderDTO.md) |  | [optional] 
**ValidationErrors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetIdpGooglesRes

`func NewGetIdpGooglesRes() *GetIdpGooglesRes`

NewGetIdpGooglesRes instantiates a new GetIdpGooglesRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIdpGooglesResWithDefaults

`func NewGetIdpGooglesResWithDefaults() *GetIdpGooglesRes`

NewGetIdpGooglesResWithDefaults instantiates a new GetIdpGooglesRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *GetIdpGooglesRes) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetIdpGooglesRes) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetIdpGooglesRes) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetIdpGooglesRes) HasError() bool`

HasError returns a boolean if a field has been set.

### GetIdps

`func (o *GetIdpGooglesRes) GetIdps() []GoogleOpenIDConnectIdentityProviderDTO`

GetIdps returns the Idps field if non-nil, zero value otherwise.

### GetIdpsOk

`func (o *GetIdpGooglesRes) GetIdpsOk() (*[]GoogleOpenIDConnectIdentityProviderDTO, bool)`

GetIdpsOk returns a tuple with the Idps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdps

`func (o *GetIdpGooglesRes) SetIdps(v []GoogleOpenIDConnectIdentityProviderDTO)`

SetIdps sets Idps field to given value.

### HasIdps

`func (o *GetIdpGooglesRes) HasIdps() bool`

HasIdps returns a boolean if a field has been set.

### GetValidationErrors

`func (o *GetIdpGooglesRes) GetValidationErrors() []string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *GetIdpGooglesRes) GetValidationErrorsOk() (*[]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *GetIdpGooglesRes) SetValidationErrors(v []string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *GetIdpGooglesRes) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


