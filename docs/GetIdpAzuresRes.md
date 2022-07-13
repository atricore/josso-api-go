# GetIdpAzuresRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**Idps** | Pointer to [**[]AzureOpenIDConnectIdentityProviderDTO**](AzureOpenIDConnectIdentityProviderDTO.md) |  | [optional] 
**ValidationErrors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetIdpAzuresRes

`func NewGetIdpAzuresRes() *GetIdpAzuresRes`

NewGetIdpAzuresRes instantiates a new GetIdpAzuresRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIdpAzuresResWithDefaults

`func NewGetIdpAzuresResWithDefaults() *GetIdpAzuresRes`

NewGetIdpAzuresResWithDefaults instantiates a new GetIdpAzuresRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *GetIdpAzuresRes) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetIdpAzuresRes) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetIdpAzuresRes) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetIdpAzuresRes) HasError() bool`

HasError returns a boolean if a field has been set.

### GetIdps

`func (o *GetIdpAzuresRes) GetIdps() []AzureOpenIDConnectIdentityProviderDTO`

GetIdps returns the Idps field if non-nil, zero value otherwise.

### GetIdpsOk

`func (o *GetIdpAzuresRes) GetIdpsOk() (*[]AzureOpenIDConnectIdentityProviderDTO, bool)`

GetIdpsOk returns a tuple with the Idps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdps

`func (o *GetIdpAzuresRes) SetIdps(v []AzureOpenIDConnectIdentityProviderDTO)`

SetIdps sets Idps field to given value.

### HasIdps

`func (o *GetIdpAzuresRes) HasIdps() bool`

HasIdps returns a boolean if a field has been set.

### GetValidationErrors

`func (o *GetIdpAzuresRes) GetValidationErrors() []string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *GetIdpAzuresRes) GetValidationErrorsOk() (*[]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *GetIdpAzuresRes) SetValidationErrors(v []string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *GetIdpAzuresRes) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


