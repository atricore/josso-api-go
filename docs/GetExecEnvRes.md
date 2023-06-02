# GetExecEnvRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**ExecEnv** | Pointer to [**ExecEnvContainerDTO**](ExecEnvContainerDTO.md) |  | [optional] 
**ValidationErrors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetExecEnvRes

`func NewGetExecEnvRes() *GetExecEnvRes`

NewGetExecEnvRes instantiates a new GetExecEnvRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetExecEnvResWithDefaults

`func NewGetExecEnvResWithDefaults() *GetExecEnvRes`

NewGetExecEnvResWithDefaults instantiates a new GetExecEnvRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *GetExecEnvRes) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetExecEnvRes) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetExecEnvRes) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetExecEnvRes) HasError() bool`

HasError returns a boolean if a field has been set.

### GetExecEnv

`func (o *GetExecEnvRes) GetExecEnv() ExecEnvContainerDTO`

GetExecEnv returns the ExecEnv field if non-nil, zero value otherwise.

### GetExecEnvOk

`func (o *GetExecEnvRes) GetExecEnvOk() (*ExecEnvContainerDTO, bool)`

GetExecEnvOk returns a tuple with the ExecEnv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecEnv

`func (o *GetExecEnvRes) SetExecEnv(v ExecEnvContainerDTO)`

SetExecEnv sets ExecEnv field to given value.

### HasExecEnv

`func (o *GetExecEnvRes) HasExecEnv() bool`

HasExecEnv returns a boolean if a field has been set.

### GetValidationErrors

`func (o *GetExecEnvRes) GetValidationErrors() []string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *GetExecEnvRes) GetValidationErrorsOk() (*[]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *GetExecEnvRes) SetValidationErrors(v []string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *GetExecEnvRes) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


