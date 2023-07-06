# GetIdPSaml2sRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**Idps** | Pointer to [**[]ExternalSaml2IdentityProviderDTO**](ExternalSaml2IdentityProviderDTO.md) |  | [optional] 
**ValidationErrors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetIdPSaml2sRes

`func NewGetIdPSaml2sRes() *GetIdPSaml2sRes`

NewGetIdPSaml2sRes instantiates a new GetIdPSaml2sRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIdPSaml2sResWithDefaults

`func NewGetIdPSaml2sResWithDefaults() *GetIdPSaml2sRes`

NewGetIdPSaml2sResWithDefaults instantiates a new GetIdPSaml2sRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *GetIdPSaml2sRes) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetIdPSaml2sRes) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetIdPSaml2sRes) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetIdPSaml2sRes) HasError() bool`

HasError returns a boolean if a field has been set.

### GetIdps

`func (o *GetIdPSaml2sRes) GetIdps() []ExternalSaml2IdentityProviderDTO`

GetIdps returns the Idps field if non-nil, zero value otherwise.

### GetIdpsOk

`func (o *GetIdPSaml2sRes) GetIdpsOk() (*[]ExternalSaml2IdentityProviderDTO, bool)`

GetIdpsOk returns a tuple with the Idps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdps

`func (o *GetIdPSaml2sRes) SetIdps(v []ExternalSaml2IdentityProviderDTO)`

SetIdps sets Idps field to given value.

### HasIdps

`func (o *GetIdPSaml2sRes) HasIdps() bool`

HasIdps returns a boolean if a field has been set.

### GetValidationErrors

`func (o *GetIdPSaml2sRes) GetValidationErrors() []string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *GetIdPSaml2sRes) GetValidationErrorsOk() (*[]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *GetIdPSaml2sRes) SetValidationErrors(v []string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *GetIdPSaml2sRes) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


