# OIDCSignOnRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**Secret** | Pointer to **string** |  | [optional] 
**Server** | Pointer to [**ServerContext**](ServerContext.md) |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewOIDCSignOnRequest

`func NewOIDCSignOnRequest() *OIDCSignOnRequest`

NewOIDCSignOnRequest instantiates a new OIDCSignOnRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOIDCSignOnRequestWithDefaults

`func NewOIDCSignOnRequestWithDefaults() *OIDCSignOnRequest`

NewOIDCSignOnRequestWithDefaults instantiates a new OIDCSignOnRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *OIDCSignOnRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OIDCSignOnRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OIDCSignOnRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *OIDCSignOnRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetPassword

`func (o *OIDCSignOnRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *OIDCSignOnRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *OIDCSignOnRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *OIDCSignOnRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetSecret

`func (o *OIDCSignOnRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *OIDCSignOnRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *OIDCSignOnRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *OIDCSignOnRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetServer

`func (o *OIDCSignOnRequest) GetServer() ServerContext`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *OIDCSignOnRequest) GetServerOk() (*ServerContext, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *OIDCSignOnRequest) SetServer(v ServerContext)`

SetServer sets Server field to given value.

### HasServer

`func (o *OIDCSignOnRequest) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetUsername

`func (o *OIDCSignOnRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *OIDCSignOnRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *OIDCSignOnRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *OIDCSignOnRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


