# StoreIisExecEnvReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdOrName** | Pointer to **string** |  | [optional] 
**IisExecEnv** | Pointer to [**WindowsIISExecutionEnvironmentDTO**](WindowsIISExecutionEnvironmentDTO.md) |  | [optional] 

## Methods

### NewStoreIisExecEnvReq

`func NewStoreIisExecEnvReq() *StoreIisExecEnvReq`

NewStoreIisExecEnvReq instantiates a new StoreIisExecEnvReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreIisExecEnvReqWithDefaults

`func NewStoreIisExecEnvReqWithDefaults() *StoreIisExecEnvReq`

NewStoreIisExecEnvReqWithDefaults instantiates a new StoreIisExecEnvReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdOrName

`func (o *StoreIisExecEnvReq) GetIdOrName() string`

GetIdOrName returns the IdOrName field if non-nil, zero value otherwise.

### GetIdOrNameOk

`func (o *StoreIisExecEnvReq) GetIdOrNameOk() (*string, bool)`

GetIdOrNameOk returns a tuple with the IdOrName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdOrName

`func (o *StoreIisExecEnvReq) SetIdOrName(v string)`

SetIdOrName sets IdOrName field to given value.

### HasIdOrName

`func (o *StoreIisExecEnvReq) HasIdOrName() bool`

HasIdOrName returns a boolean if a field has been set.

### GetIisExecEnv

`func (o *StoreIisExecEnvReq) GetIisExecEnv() WindowsIISExecutionEnvironmentDTO`

GetIisExecEnv returns the IisExecEnv field if non-nil, zero value otherwise.

### GetIisExecEnvOk

`func (o *StoreIisExecEnvReq) GetIisExecEnvOk() (*WindowsIISExecutionEnvironmentDTO, bool)`

GetIisExecEnvOk returns a tuple with the IisExecEnv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIisExecEnv

`func (o *StoreIisExecEnvReq) SetIisExecEnv(v WindowsIISExecutionEnvironmentDTO)`

SetIisExecEnv sets IisExecEnv field to given value.

### HasIisExecEnv

`func (o *StoreIisExecEnvReq) HasIisExecEnv() bool`

HasIisExecEnv returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


