# StorePhpExecEnvReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdOrName** | Pointer to **string** |  | [optional] 
**PhpExecEnv** | Pointer to [**PHPExecutionEnvironmentDTO**](PHPExecutionEnvironmentDTO.md) |  | [optional] 

## Methods

### NewStorePhpExecEnvReq

`func NewStorePhpExecEnvReq() *StorePhpExecEnvReq`

NewStorePhpExecEnvReq instantiates a new StorePhpExecEnvReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorePhpExecEnvReqWithDefaults

`func NewStorePhpExecEnvReqWithDefaults() *StorePhpExecEnvReq`

NewStorePhpExecEnvReqWithDefaults instantiates a new StorePhpExecEnvReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdOrName

`func (o *StorePhpExecEnvReq) GetIdOrName() string`

GetIdOrName returns the IdOrName field if non-nil, zero value otherwise.

### GetIdOrNameOk

`func (o *StorePhpExecEnvReq) GetIdOrNameOk() (*string, bool)`

GetIdOrNameOk returns a tuple with the IdOrName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdOrName

`func (o *StorePhpExecEnvReq) SetIdOrName(v string)`

SetIdOrName sets IdOrName field to given value.

### HasIdOrName

`func (o *StorePhpExecEnvReq) HasIdOrName() bool`

HasIdOrName returns a boolean if a field has been set.

### GetPhpExecEnv

`func (o *StorePhpExecEnvReq) GetPhpExecEnv() PHPExecutionEnvironmentDTO`

GetPhpExecEnv returns the PhpExecEnv field if non-nil, zero value otherwise.

### GetPhpExecEnvOk

`func (o *StorePhpExecEnvReq) GetPhpExecEnvOk() (*PHPExecutionEnvironmentDTO, bool)`

GetPhpExecEnvOk returns a tuple with the PhpExecEnv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhpExecEnv

`func (o *StorePhpExecEnvReq) SetPhpExecEnv(v PHPExecutionEnvironmentDTO)`

SetPhpExecEnv sets PhpExecEnv field to given value.

### HasPhpExecEnv

`func (o *StorePhpExecEnvReq) HasPhpExecEnv() bool`

HasPhpExecEnv returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


