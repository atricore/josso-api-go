# GetServerInfoRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**NodeId** | Pointer to **string** |  | [optional] 
**ValidationErrors** | Pointer to **[]string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewGetServerInfoRes

`func NewGetServerInfoRes() *GetServerInfoRes`

NewGetServerInfoRes instantiates a new GetServerInfoRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetServerInfoResWithDefaults

`func NewGetServerInfoResWithDefaults() *GetServerInfoRes`

NewGetServerInfoResWithDefaults instantiates a new GetServerInfoRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *GetServerInfoRes) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetServerInfoRes) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetServerInfoRes) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetServerInfoRes) HasError() bool`

HasError returns a boolean if a field has been set.

### GetNodeId

`func (o *GetServerInfoRes) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *GetServerInfoRes) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *GetServerInfoRes) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *GetServerInfoRes) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetValidationErrors

`func (o *GetServerInfoRes) GetValidationErrors() []string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *GetServerInfoRes) GetValidationErrorsOk() (*[]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *GetServerInfoRes) SetValidationErrors(v []string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *GetServerInfoRes) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.

### GetVersion

`func (o *GetServerInfoRes) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetServerInfoRes) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetServerInfoRes) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GetServerInfoRes) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


