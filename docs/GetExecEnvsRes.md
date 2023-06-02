# GetExecEnvsRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**ExecEnvs** | Pointer to [**[]ExecEnvContainerDTO**](ExecEnvContainerDTO.md) |  | [optional] 
**ValidationErrors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetExecEnvsRes

`func NewGetExecEnvsRes() *GetExecEnvsRes`

NewGetExecEnvsRes instantiates a new GetExecEnvsRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetExecEnvsResWithDefaults

`func NewGetExecEnvsResWithDefaults() *GetExecEnvsRes`

NewGetExecEnvsResWithDefaults instantiates a new GetExecEnvsRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *GetExecEnvsRes) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetExecEnvsRes) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetExecEnvsRes) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetExecEnvsRes) HasError() bool`

HasError returns a boolean if a field has been set.

### GetExecEnvs

`func (o *GetExecEnvsRes) GetExecEnvs() []ExecEnvContainerDTO`

GetExecEnvs returns the ExecEnvs field if non-nil, zero value otherwise.

### GetExecEnvsOk

`func (o *GetExecEnvsRes) GetExecEnvsOk() (*[]ExecEnvContainerDTO, bool)`

GetExecEnvsOk returns a tuple with the ExecEnvs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecEnvs

`func (o *GetExecEnvsRes) SetExecEnvs(v []ExecEnvContainerDTO)`

SetExecEnvs sets ExecEnvs field to given value.

### HasExecEnvs

`func (o *GetExecEnvsRes) HasExecEnvs() bool`

HasExecEnvs returns a boolean if a field has been set.

### GetValidationErrors

`func (o *GetExecEnvsRes) GetValidationErrors() []string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *GetExecEnvsRes) GetValidationErrorsOk() (*[]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *GetExecEnvsRes) SetValidationErrors(v []string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *GetExecEnvsRes) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


