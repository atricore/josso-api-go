# GetOidcRpsRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**OidcRps** | Pointer to [**[]ExternalOpenIDConnectRelayingPartyDTO**](ExternalOpenIDConnectRelayingPartyDTO.md) |  | [optional] 
**ValidationErrors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetOidcRpsRes

`func NewGetOidcRpsRes() *GetOidcRpsRes`

NewGetOidcRpsRes instantiates a new GetOidcRpsRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOidcRpsResWithDefaults

`func NewGetOidcRpsResWithDefaults() *GetOidcRpsRes`

NewGetOidcRpsResWithDefaults instantiates a new GetOidcRpsRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *GetOidcRpsRes) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetOidcRpsRes) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetOidcRpsRes) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetOidcRpsRes) HasError() bool`

HasError returns a boolean if a field has been set.

### GetOidcRps

`func (o *GetOidcRpsRes) GetOidcRps() []ExternalOpenIDConnectRelayingPartyDTO`

GetOidcRps returns the OidcRps field if non-nil, zero value otherwise.

### GetOidcRpsOk

`func (o *GetOidcRpsRes) GetOidcRpsOk() (*[]ExternalOpenIDConnectRelayingPartyDTO, bool)`

GetOidcRpsOk returns a tuple with the OidcRps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcRps

`func (o *GetOidcRpsRes) SetOidcRps(v []ExternalOpenIDConnectRelayingPartyDTO)`

SetOidcRps sets OidcRps field to given value.

### HasOidcRps

`func (o *GetOidcRpsRes) HasOidcRps() bool`

HasOidcRps returns a boolean if a field has been set.

### GetValidationErrors

`func (o *GetOidcRpsRes) GetValidationErrors() []string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *GetOidcRpsRes) GetValidationErrorsOk() (*[]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *GetOidcRpsRes) SetValidationErrors(v []string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *GetOidcRpsRes) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


