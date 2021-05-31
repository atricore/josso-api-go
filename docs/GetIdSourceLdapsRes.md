# GetIdSourceLdapsRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**IdSourceLdaps** | Pointer to [**[]LdapIdentitySourceDTO**](LdapIdentitySourceDTO.md) |  | [optional] 
**ValidationErrors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetIdSourceLdapsRes

`func NewGetIdSourceLdapsRes() *GetIdSourceLdapsRes`

NewGetIdSourceLdapsRes instantiates a new GetIdSourceLdapsRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIdSourceLdapsResWithDefaults

`func NewGetIdSourceLdapsResWithDefaults() *GetIdSourceLdapsRes`

NewGetIdSourceLdapsResWithDefaults instantiates a new GetIdSourceLdapsRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *GetIdSourceLdapsRes) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetIdSourceLdapsRes) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetIdSourceLdapsRes) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetIdSourceLdapsRes) HasError() bool`

HasError returns a boolean if a field has been set.

### GetIdSourceLdaps

`func (o *GetIdSourceLdapsRes) GetIdSourceLdaps() []LdapIdentitySourceDTO`

GetIdSourceLdaps returns the IdSourceLdaps field if non-nil, zero value otherwise.

### GetIdSourceLdapsOk

`func (o *GetIdSourceLdapsRes) GetIdSourceLdapsOk() (*[]LdapIdentitySourceDTO, bool)`

GetIdSourceLdapsOk returns a tuple with the IdSourceLdaps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdSourceLdaps

`func (o *GetIdSourceLdapsRes) SetIdSourceLdaps(v []LdapIdentitySourceDTO)`

SetIdSourceLdaps sets IdSourceLdaps field to given value.

### HasIdSourceLdaps

`func (o *GetIdSourceLdapsRes) HasIdSourceLdaps() bool`

HasIdSourceLdaps returns a boolean if a field has been set.

### GetValidationErrors

`func (o *GetIdSourceLdapsRes) GetValidationErrors() []string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *GetIdSourceLdapsRes) GetValidationErrorsOk() (*[]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *GetIdSourceLdapsRes) SetValidationErrors(v []string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *GetIdSourceLdapsRes) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


