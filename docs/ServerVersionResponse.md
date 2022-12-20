# ServerVersionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Server** | Pointer to [**ServerContext**](ServerContext.md) |  | [optional] 
**ValidationErrors** | Pointer to **[]string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewServerVersionResponse

`func NewServerVersionResponse() *ServerVersionResponse`

NewServerVersionResponse instantiates a new ServerVersionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerVersionResponseWithDefaults

`func NewServerVersionResponseWithDefaults() *ServerVersionResponse`

NewServerVersionResponseWithDefaults instantiates a new ServerVersionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServer

`func (o *ServerVersionResponse) GetServer() ServerContext`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *ServerVersionResponse) GetServerOk() (*ServerContext, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *ServerVersionResponse) SetServer(v ServerContext)`

SetServer sets Server field to given value.

### HasServer

`func (o *ServerVersionResponse) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetValidationErrors

`func (o *ServerVersionResponse) GetValidationErrors() []string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *ServerVersionResponse) GetValidationErrorsOk() (*[]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *ServerVersionResponse) SetValidationErrors(v []string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *ServerVersionResponse) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.

### GetVersion

`func (o *ServerVersionResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ServerVersionResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ServerVersionResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ServerVersionResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


