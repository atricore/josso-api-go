# GetTomcatExecEnvRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**TomcatExecEnv** | Pointer to [**TomcatExecutionEnvironmentDTO**](TomcatExecutionEnvironmentDTO.md) |  | [optional] 
**ValidationErrors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetTomcatExecEnvRes

`func NewGetTomcatExecEnvRes() *GetTomcatExecEnvRes`

NewGetTomcatExecEnvRes instantiates a new GetTomcatExecEnvRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTomcatExecEnvResWithDefaults

`func NewGetTomcatExecEnvResWithDefaults() *GetTomcatExecEnvRes`

NewGetTomcatExecEnvResWithDefaults instantiates a new GetTomcatExecEnvRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *GetTomcatExecEnvRes) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetTomcatExecEnvRes) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetTomcatExecEnvRes) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetTomcatExecEnvRes) HasError() bool`

HasError returns a boolean if a field has been set.

### GetTomcatExecEnv

`func (o *GetTomcatExecEnvRes) GetTomcatExecEnv() TomcatExecutionEnvironmentDTO`

GetTomcatExecEnv returns the TomcatExecEnv field if non-nil, zero value otherwise.

### GetTomcatExecEnvOk

`func (o *GetTomcatExecEnvRes) GetTomcatExecEnvOk() (*TomcatExecutionEnvironmentDTO, bool)`

GetTomcatExecEnvOk returns a tuple with the TomcatExecEnv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTomcatExecEnv

`func (o *GetTomcatExecEnvRes) SetTomcatExecEnv(v TomcatExecutionEnvironmentDTO)`

SetTomcatExecEnv sets TomcatExecEnv field to given value.

### HasTomcatExecEnv

`func (o *GetTomcatExecEnvRes) HasTomcatExecEnv() bool`

HasTomcatExecEnv returns a boolean if a field has been set.

### GetValidationErrors

`func (o *GetTomcatExecEnvRes) GetValidationErrors() []string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *GetTomcatExecEnvRes) GetValidationErrorsOk() (*[]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *GetTomcatExecEnvRes) SetValidationErrors(v []string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *GetTomcatExecEnvRes) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


