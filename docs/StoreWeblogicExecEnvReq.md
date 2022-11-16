# StoreWeblogicExecEnvReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExecEnv** | Pointer to [**WeblogicExecutionEnvironmentDTO**](WeblogicExecutionEnvironmentDTO.md) |  | [optional] 
**IdOrName** | Pointer to **string** |  | [optional] 

## Methods

### NewStoreWeblogicExecEnvReq

`func NewStoreWeblogicExecEnvReq() *StoreWeblogicExecEnvReq`

NewStoreWeblogicExecEnvReq instantiates a new StoreWeblogicExecEnvReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreWeblogicExecEnvReqWithDefaults

`func NewStoreWeblogicExecEnvReqWithDefaults() *StoreWeblogicExecEnvReq`

NewStoreWeblogicExecEnvReqWithDefaults instantiates a new StoreWeblogicExecEnvReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExecEnv

`func (o *StoreWeblogicExecEnvReq) GetExecEnv() WeblogicExecutionEnvironmentDTO`

GetExecEnv returns the ExecEnv field if non-nil, zero value otherwise.

### GetExecEnvOk

`func (o *StoreWeblogicExecEnvReq) GetExecEnvOk() (*WeblogicExecutionEnvironmentDTO, bool)`

GetExecEnvOk returns a tuple with the ExecEnv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecEnv

`func (o *StoreWeblogicExecEnvReq) SetExecEnv(v WeblogicExecutionEnvironmentDTO)`

SetExecEnv sets ExecEnv field to given value.

### HasExecEnv

`func (o *StoreWeblogicExecEnvReq) HasExecEnv() bool`

HasExecEnv returns a boolean if a field has been set.

### GetIdOrName

`func (o *StoreWeblogicExecEnvReq) GetIdOrName() string`

GetIdOrName returns the IdOrName field if non-nil, zero value otherwise.

### GetIdOrNameOk

`func (o *StoreWeblogicExecEnvReq) GetIdOrNameOk() (*string, bool)`

GetIdOrNameOk returns a tuple with the IdOrName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdOrName

`func (o *StoreWeblogicExecEnvReq) SetIdOrName(v string)`

SetIdOrName sets IdOrName field to given value.

### HasIdOrName

`func (o *StoreWeblogicExecEnvReq) HasIdOrName() bool`

HasIdOrName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


