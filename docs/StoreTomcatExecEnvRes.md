# StoreTomcatExecEnvRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**TomcatExecEnv** | Pointer to [**TomcatExecutionEnvironmentDTO**](TomcatExecutionEnvironmentDTO.md) |  | [optional] 
**ValidationErrors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewStoreTomcatExecEnvRes

`func NewStoreTomcatExecEnvRes() *StoreTomcatExecEnvRes`

NewStoreTomcatExecEnvRes instantiates a new StoreTomcatExecEnvRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreTomcatExecEnvResWithDefaults

`func NewStoreTomcatExecEnvResWithDefaults() *StoreTomcatExecEnvRes`

NewStoreTomcatExecEnvResWithDefaults instantiates a new StoreTomcatExecEnvRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *StoreTomcatExecEnvRes) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *StoreTomcatExecEnvRes) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *StoreTomcatExecEnvRes) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *StoreTomcatExecEnvRes) HasError() bool`

HasError returns a boolean if a field has been set.

### GetTomcatExecEnv

`func (o *StoreTomcatExecEnvRes) GetTomcatExecEnv() TomcatExecutionEnvironmentDTO`

GetTomcatExecEnv returns the TomcatExecEnv field if non-nil, zero value otherwise.

### GetTomcatExecEnvOk

`func (o *StoreTomcatExecEnvRes) GetTomcatExecEnvOk() (*TomcatExecutionEnvironmentDTO, bool)`

GetTomcatExecEnvOk returns a tuple with the TomcatExecEnv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTomcatExecEnv

`func (o *StoreTomcatExecEnvRes) SetTomcatExecEnv(v TomcatExecutionEnvironmentDTO)`

SetTomcatExecEnv sets TomcatExecEnv field to given value.

### HasTomcatExecEnv

`func (o *StoreTomcatExecEnvRes) HasTomcatExecEnv() bool`

HasTomcatExecEnv returns a boolean if a field has been set.

### GetValidationErrors

`func (o *StoreTomcatExecEnvRes) GetValidationErrors() []string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *StoreTomcatExecEnvRes) GetValidationErrorsOk() (*[]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *StoreTomcatExecEnvRes) SetValidationErrors(v []string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *StoreTomcatExecEnvRes) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


