# ServerVersionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Server** | Pointer to [**ServerContext**](ServerContext.md) |  | [optional] 

## Methods

### NewServerVersionRequest

`func NewServerVersionRequest() *ServerVersionRequest`

NewServerVersionRequest instantiates a new ServerVersionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerVersionRequestWithDefaults

`func NewServerVersionRequestWithDefaults() *ServerVersionRequest`

NewServerVersionRequestWithDefaults instantiates a new ServerVersionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServer

`func (o *ServerVersionRequest) GetServer() ServerContext`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *ServerVersionRequest) GetServerOk() (*ServerContext, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *ServerVersionRequest) SetServer(v ServerContext)`

SetServer sets Server field to given value.

### HasServer

`func (o *ServerVersionRequest) HasServer() bool`

HasServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


