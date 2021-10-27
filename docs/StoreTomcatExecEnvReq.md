# StoreTomcatExecEnvReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdaName** | Pointer to **string** |  | [optional] 
**TomcatExecEnv** | Pointer to [**TomcatExecutionEnvironmentDTO**](TomcatExecutionEnvironmentDTO.md) |  | [optional] 

## Methods

### NewStoreTomcatExecEnvReq

`func NewStoreTomcatExecEnvReq() *StoreTomcatExecEnvReq`

NewStoreTomcatExecEnvReq instantiates a new StoreTomcatExecEnvReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreTomcatExecEnvReqWithDefaults

`func NewStoreTomcatExecEnvReqWithDefaults() *StoreTomcatExecEnvReq`

NewStoreTomcatExecEnvReqWithDefaults instantiates a new StoreTomcatExecEnvReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdaName

`func (o *StoreTomcatExecEnvReq) GetIdaName() string`

GetIdaName returns the IdaName field if non-nil, zero value otherwise.

### GetIdaNameOk

`func (o *StoreTomcatExecEnvReq) GetIdaNameOk() (*string, bool)`

GetIdaNameOk returns a tuple with the IdaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdaName

`func (o *StoreTomcatExecEnvReq) SetIdaName(v string)`

SetIdaName sets IdaName field to given value.

### HasIdaName

`func (o *StoreTomcatExecEnvReq) HasIdaName() bool`

HasIdaName returns a boolean if a field has been set.

### GetTomcatExecEnv

`func (o *StoreTomcatExecEnvReq) GetTomcatExecEnv() TomcatExecutionEnvironmentDTO`

GetTomcatExecEnv returns the TomcatExecEnv field if non-nil, zero value otherwise.

### GetTomcatExecEnvOk

`func (o *StoreTomcatExecEnvReq) GetTomcatExecEnvOk() (*TomcatExecutionEnvironmentDTO, bool)`

GetTomcatExecEnvOk returns a tuple with the TomcatExecEnv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTomcatExecEnv

`func (o *StoreTomcatExecEnvReq) SetTomcatExecEnv(v TomcatExecutionEnvironmentDTO)`

SetTomcatExecEnv sets TomcatExecEnv field to given value.

### HasTomcatExecEnv

`func (o *StoreTomcatExecEnvReq) HasTomcatExecEnv() bool`

HasTomcatExecEnv returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


