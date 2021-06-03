# StoreOidcRpRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**OidcRp** | Pointer to [**ExternalOpenIDConnectRelayingPartyDTO**](ExternalOpenIDConnectRelayingPartyDTO.md) |  | [optional] 
**ValidationErrors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewStoreOidcRpRes

`func NewStoreOidcRpRes() *StoreOidcRpRes`

NewStoreOidcRpRes instantiates a new StoreOidcRpRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreOidcRpResWithDefaults

`func NewStoreOidcRpResWithDefaults() *StoreOidcRpRes`

NewStoreOidcRpResWithDefaults instantiates a new StoreOidcRpRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *StoreOidcRpRes) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *StoreOidcRpRes) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *StoreOidcRpRes) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *StoreOidcRpRes) HasError() bool`

HasError returns a boolean if a field has been set.

### GetOidcRp

`func (o *StoreOidcRpRes) GetOidcRp() ExternalOpenIDConnectRelayingPartyDTO`

GetOidcRp returns the OidcRp field if non-nil, zero value otherwise.

### GetOidcRpOk

`func (o *StoreOidcRpRes) GetOidcRpOk() (*ExternalOpenIDConnectRelayingPartyDTO, bool)`

GetOidcRpOk returns a tuple with the OidcRp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcRp

`func (o *StoreOidcRpRes) SetOidcRp(v ExternalOpenIDConnectRelayingPartyDTO)`

SetOidcRp sets OidcRp field to given value.

### HasOidcRp

`func (o *StoreOidcRpRes) HasOidcRp() bool`

HasOidcRp returns a boolean if a field has been set.

### GetValidationErrors

`func (o *StoreOidcRpRes) GetValidationErrors() []string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *StoreOidcRpRes) GetValidationErrorsOk() (*[]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *StoreOidcRpRes) SetValidationErrors(v []string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *StoreOidcRpRes) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


