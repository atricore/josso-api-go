# GetOidcRpRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**OidcRp** | Pointer to [**ExternalOpenIDConnectRelayingPartyDTO**](ExternalOpenIDConnectRelayingPartyDTO.md) |  | [optional] 
**ValidationErrors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetOidcRpRes

`func NewGetOidcRpRes() *GetOidcRpRes`

NewGetOidcRpRes instantiates a new GetOidcRpRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOidcRpResWithDefaults

`func NewGetOidcRpResWithDefaults() *GetOidcRpRes`

NewGetOidcRpResWithDefaults instantiates a new GetOidcRpRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *GetOidcRpRes) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetOidcRpRes) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetOidcRpRes) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetOidcRpRes) HasError() bool`

HasError returns a boolean if a field has been set.

### GetOidcRp

`func (o *GetOidcRpRes) GetOidcRp() ExternalOpenIDConnectRelayingPartyDTO`

GetOidcRp returns the OidcRp field if non-nil, zero value otherwise.

### GetOidcRpOk

`func (o *GetOidcRpRes) GetOidcRpOk() (*ExternalOpenIDConnectRelayingPartyDTO, bool)`

GetOidcRpOk returns a tuple with the OidcRp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcRp

`func (o *GetOidcRpRes) SetOidcRp(v ExternalOpenIDConnectRelayingPartyDTO)`

SetOidcRp sets OidcRp field to given value.

### HasOidcRp

`func (o *GetOidcRpRes) HasOidcRp() bool`

HasOidcRp returns a boolean if a field has been set.

### GetValidationErrors

`func (o *GetOidcRpRes) GetValidationErrors() []string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *GetOidcRpRes) GetValidationErrorsOk() (*[]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *GetOidcRpRes) SetValidationErrors(v []string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *GetOidcRpRes) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


