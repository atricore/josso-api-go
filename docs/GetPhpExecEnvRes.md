# GetPhpExecEnvRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**PhpExecEnv** | Pointer to [**PHPExecutionEnvironmentDTO**](PHPExecutionEnvironmentDTO.md) |  | [optional] 
**ValidationErrors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetPhpExecEnvRes

`func NewGetPhpExecEnvRes() *GetPhpExecEnvRes`

NewGetPhpExecEnvRes instantiates a new GetPhpExecEnvRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPhpExecEnvResWithDefaults

`func NewGetPhpExecEnvResWithDefaults() *GetPhpExecEnvRes`

NewGetPhpExecEnvResWithDefaults instantiates a new GetPhpExecEnvRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *GetPhpExecEnvRes) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetPhpExecEnvRes) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetPhpExecEnvRes) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetPhpExecEnvRes) HasError() bool`

HasError returns a boolean if a field has been set.

### GetPhpExecEnv

`func (o *GetPhpExecEnvRes) GetPhpExecEnv() PHPExecutionEnvironmentDTO`

GetPhpExecEnv returns the PhpExecEnv field if non-nil, zero value otherwise.

### GetPhpExecEnvOk

`func (o *GetPhpExecEnvRes) GetPhpExecEnvOk() (*PHPExecutionEnvironmentDTO, bool)`

GetPhpExecEnvOk returns a tuple with the PhpExecEnv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhpExecEnv

`func (o *GetPhpExecEnvRes) SetPhpExecEnv(v PHPExecutionEnvironmentDTO)`

SetPhpExecEnv sets PhpExecEnv field to given value.

### HasPhpExecEnv

`func (o *GetPhpExecEnvRes) HasPhpExecEnv() bool`

HasPhpExecEnv returns a boolean if a field has been set.

### GetValidationErrors

`func (o *GetPhpExecEnvRes) GetValidationErrors() []string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *GetPhpExecEnvRes) GetValidationErrorsOk() (*[]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *GetPhpExecEnvRes) SetValidationErrors(v []string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *GetPhpExecEnvRes) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


