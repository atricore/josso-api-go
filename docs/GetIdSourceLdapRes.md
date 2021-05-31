# GetIdSourceLdapRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**IdSourceLdap** | Pointer to [**LdapIdentitySourceDTO**](LdapIdentitySourceDTO.md) |  | [optional] 
**ValidationErrors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetIdSourceLdapRes

`func NewGetIdSourceLdapRes() *GetIdSourceLdapRes`

NewGetIdSourceLdapRes instantiates a new GetIdSourceLdapRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIdSourceLdapResWithDefaults

`func NewGetIdSourceLdapResWithDefaults() *GetIdSourceLdapRes`

NewGetIdSourceLdapResWithDefaults instantiates a new GetIdSourceLdapRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *GetIdSourceLdapRes) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetIdSourceLdapRes) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetIdSourceLdapRes) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetIdSourceLdapRes) HasError() bool`

HasError returns a boolean if a field has been set.

### GetIdSourceLdap

`func (o *GetIdSourceLdapRes) GetIdSourceLdap() LdapIdentitySourceDTO`

GetIdSourceLdap returns the IdSourceLdap field if non-nil, zero value otherwise.

### GetIdSourceLdapOk

`func (o *GetIdSourceLdapRes) GetIdSourceLdapOk() (*LdapIdentitySourceDTO, bool)`

GetIdSourceLdapOk returns a tuple with the IdSourceLdap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdSourceLdap

`func (o *GetIdSourceLdapRes) SetIdSourceLdap(v LdapIdentitySourceDTO)`

SetIdSourceLdap sets IdSourceLdap field to given value.

### HasIdSourceLdap

`func (o *GetIdSourceLdapRes) HasIdSourceLdap() bool`

HasIdSourceLdap returns a boolean if a field has been set.

### GetValidationErrors

`func (o *GetIdSourceLdapRes) GetValidationErrors() []string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *GetIdSourceLdapRes) GetValidationErrorsOk() (*[]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *GetIdSourceLdapRes) SetValidationErrors(v []string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *GetIdSourceLdapRes) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


