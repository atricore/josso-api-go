# GetWeblogicExecEnvsRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**ExecEnv** | Pointer to [**[]WeblogicExecutionEnvironmentDTO**](WeblogicExecutionEnvironmentDTO.md) |  | [optional] 
**ValidationErrors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetWeblogicExecEnvsRes

`func NewGetWeblogicExecEnvsRes() *GetWeblogicExecEnvsRes`

NewGetWeblogicExecEnvsRes instantiates a new GetWeblogicExecEnvsRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWeblogicExecEnvsResWithDefaults

`func NewGetWeblogicExecEnvsResWithDefaults() *GetWeblogicExecEnvsRes`

NewGetWeblogicExecEnvsResWithDefaults instantiates a new GetWeblogicExecEnvsRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *GetWeblogicExecEnvsRes) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetWeblogicExecEnvsRes) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetWeblogicExecEnvsRes) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetWeblogicExecEnvsRes) HasError() bool`

HasError returns a boolean if a field has been set.

### GetExecEnv

`func (o *GetWeblogicExecEnvsRes) GetExecEnv() []WeblogicExecutionEnvironmentDTO`

GetExecEnv returns the ExecEnv field if non-nil, zero value otherwise.

### GetExecEnvOk

`func (o *GetWeblogicExecEnvsRes) GetExecEnvOk() (*[]WeblogicExecutionEnvironmentDTO, bool)`

GetExecEnvOk returns a tuple with the ExecEnv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecEnv

`func (o *GetWeblogicExecEnvsRes) SetExecEnv(v []WeblogicExecutionEnvironmentDTO)`

SetExecEnv sets ExecEnv field to given value.

### HasExecEnv

`func (o *GetWeblogicExecEnvsRes) HasExecEnv() bool`

HasExecEnv returns a boolean if a field has been set.

### GetValidationErrors

`func (o *GetWeblogicExecEnvsRes) GetValidationErrors() []string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *GetWeblogicExecEnvsRes) GetValidationErrorsOk() (*[]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *GetWeblogicExecEnvsRes) SetValidationErrors(v []string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *GetWeblogicExecEnvsRes) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


