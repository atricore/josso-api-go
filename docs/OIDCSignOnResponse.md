# OIDCSignOnResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | Pointer to **string** |  | [optional] 
**AuthenticatedUser** | Pointer to [**UserDTO**](UserDTO.md) |  | [optional] 
**IdToken** | Pointer to **string** |  | [optional] 
**RefreshToken** | Pointer to **string** |  | [optional] 
**Server** | Pointer to [**ServerContext**](ServerContext.md) |  | [optional] 
**ValidationErrors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewOIDCSignOnResponse

`func NewOIDCSignOnResponse() *OIDCSignOnResponse`

NewOIDCSignOnResponse instantiates a new OIDCSignOnResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOIDCSignOnResponseWithDefaults

`func NewOIDCSignOnResponseWithDefaults() *OIDCSignOnResponse`

NewOIDCSignOnResponseWithDefaults instantiates a new OIDCSignOnResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *OIDCSignOnResponse) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *OIDCSignOnResponse) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *OIDCSignOnResponse) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *OIDCSignOnResponse) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetAuthenticatedUser

`func (o *OIDCSignOnResponse) GetAuthenticatedUser() UserDTO`

GetAuthenticatedUser returns the AuthenticatedUser field if non-nil, zero value otherwise.

### GetAuthenticatedUserOk

`func (o *OIDCSignOnResponse) GetAuthenticatedUserOk() (*UserDTO, bool)`

GetAuthenticatedUserOk returns a tuple with the AuthenticatedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticatedUser

`func (o *OIDCSignOnResponse) SetAuthenticatedUser(v UserDTO)`

SetAuthenticatedUser sets AuthenticatedUser field to given value.

### HasAuthenticatedUser

`func (o *OIDCSignOnResponse) HasAuthenticatedUser() bool`

HasAuthenticatedUser returns a boolean if a field has been set.

### GetIdToken

`func (o *OIDCSignOnResponse) GetIdToken() string`

GetIdToken returns the IdToken field if non-nil, zero value otherwise.

### GetIdTokenOk

`func (o *OIDCSignOnResponse) GetIdTokenOk() (*string, bool)`

GetIdTokenOk returns a tuple with the IdToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdToken

`func (o *OIDCSignOnResponse) SetIdToken(v string)`

SetIdToken sets IdToken field to given value.

### HasIdToken

`func (o *OIDCSignOnResponse) HasIdToken() bool`

HasIdToken returns a boolean if a field has been set.

### GetRefreshToken

`func (o *OIDCSignOnResponse) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *OIDCSignOnResponse) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *OIDCSignOnResponse) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *OIDCSignOnResponse) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetServer

`func (o *OIDCSignOnResponse) GetServer() ServerContext`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *OIDCSignOnResponse) GetServerOk() (*ServerContext, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *OIDCSignOnResponse) SetServer(v ServerContext)`

SetServer sets Server field to given value.

### HasServer

`func (o *OIDCSignOnResponse) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetValidationErrors

`func (o *OIDCSignOnResponse) GetValidationErrors() []string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *OIDCSignOnResponse) GetValidationErrorsOk() (*[]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *OIDCSignOnResponse) SetValidationErrors(v []string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *OIDCSignOnResponse) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


